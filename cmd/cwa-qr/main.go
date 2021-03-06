package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/axiomhq/pkg/version"
	"github.com/skip2/go-qrcode"

	cwaqr "github.com/lukasmalkmus/cwa-qr"
)

var inputValidatorOpt = survey.WithValidator(survey.ComposeValidators(
	survey.MinLength(3),
	survey.MaxLength(cwaqr.CharacterLimit),
))

var (
	description string
	address     string
	duration    time.Duration
	startStr    string
	endStr      string
	showVersion = flag.Bool("version", false, "print the version")
)

func init() {
	flag.StringVar(&description, "description", "", "description of the event, limited to 100 characters")
	flag.StringVar(&address, "address", "", "address of the event, limited to 100 characters")
	flag.DurationVar(&duration, "duration", 0, "average stay duration at an event")
	flag.StringVar(&startStr, "start", "", "start time of a temporary event specified using RFC822")
	flag.StringVar(&endStr, "end", "", "end time of a temporary event specified using RFC822")
}

func main() {
	log.SetFlags(0)

	flag.Parse()

	if *showVersion {
		log.Printf("Corona Warn App QR-Code Generator version %s\n", version.Release())
		os.Exit(0)
	}

	outputFile := flag.Arg(0)
	if outputFile == "" {
		log.Printf("Missing output file!\n\n")
		log.Printf("Usage: %s [flags] <output-file>\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if description == "" {
		if err := survey.AskOne(&survey.Input{
			Message: "What is the description of the event?",
		}, &description, inputValidatorOpt); err != nil {
			log.Fatal(err)
		}
	}

	if address == "" {
		if err := survey.AskOne(&survey.Input{
			Message: "What is the address of the event?",
		}, &address, inputValidatorOpt); err != nil {
			log.Fatal(err)
		}
	}

	var (
		eventType    cwaqr.EventType
		eventTypeStr string
	)
	if err := survey.AskOne(&survey.Select{
		Message: "What kind of event is taking place?",
		Options: eventTypeStrs(),
	}, &eventTypeStr); err != nil {
		log.Fatal(err)
	}

	for _, et := range eventTypes() {
		if et.String() == eventTypeStr {
			eventType = et
		}
	}

	if eventType == cwaqr.Unknown {
		log.Fatalf("invalid event type %s", eventTypeStr)
	}

	if duration == 0 {
		var durationStr string
		if err := survey.AskOne(&survey.Input{
			Message: "How long is the average stay at the event?",
		}, &durationStr); err != nil {
			log.Fatal(err)
		}

		var err error
		if duration, err = time.ParseDuration(durationStr); err != nil {
			log.Fatal(err)
		}
	}

	if eventType.IsTemporary() {
		if startStr == "" {
			if err := survey.AskOne(&survey.Input{
				Message: "When does the event start?",
			}, &startStr); err != nil {
				log.Fatal(err)
			}
		}

		if endStr == "" {
			if err := survey.AskOne(&survey.Input{
				Message: "When does the event end?",
			}, &endStr); err != nil {
				log.Fatal(err)
			}
		}
	}

	var start time.Time
	if startStr != "" {
		var err error
		if start, err = time.Parse(time.RFC822, startStr); err != nil {
			log.Fatal(err)
		}
	}

	var end time.Time
	if endStr != "" {
		var err error
		if end, err = time.Parse(time.RFC822, startStr); err != nil {
			log.Fatal(err)
		}
	}

	qrCodeURL, err := cwaqr.GenerateURL(cwaqr.Event{
		Description: description,
		Address:     address,
		Type:        eventType,
		Duration:    duration,
		Start:       start,
		End:         end,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = qrcode.WriteFile(qrCodeURL.String(), qrcode.Low, 512, outputFile); err != nil {
		log.Fatal(err)
	}

	log.Printf("Created QR-Code %q!\n", outputFile)
}

func eventTypes() []cwaqr.EventType {
	var res []cwaqr.EventType
	for i := cwaqr.OtherPermanent; i <= cwaqr.WorshipService; i++ {
		res = append(res, cwaqr.EventType(i))
	}
	return res
}

func eventTypeStrs() []string {
	var res []string
	for _, eventType := range eventTypes() {
		res = append(res, eventType.String())
	}
	return res
}
