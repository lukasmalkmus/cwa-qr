# Corona Warn App QR-Code Generator

> Generate QR-Codes for checking into events using the official Corona Warn App.

[![Go Reference][gopkg_badge]][gopkg]
[![Go Workflow][go_workflow_badge]][go_workflow]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]
[![Docker][docker_badge]][docker]

---

## Table of Contents

1. [Introduction](#introduction)
1. [Installation](#installation)
1. [Usage](#usage)
1. [Contributing](#contributing)
1. [License](#license)

## Introduction

_Corona Warn App QR-Code Generator_ is a Go client library and command line
application for creating QR-Codes which users of the official Corona Warn App
can use to check into events.

It is an implementation of the Protocol used to generate event and location QR
codes for the Corona Warn App as described in Coronatheir [documentation][1].

This is not an official implementation! Use it at your own risk!

  [1]: https://github.com/corona-warn-app/cwa-documentation/blob/master/event_registration.md

### State

Currently, rotating QR-Codes is not implemented. The seed is random every time.
This needs to be addressed. Unfortunaly, the documentation is missing context.

## Installation

### Download and install the pre-compiled binary manually

Binary releases are evailable on [GitHub Releases][2].

  [2]: https://github.com/lukasmalkmus/cwa-qr/releases/latest

### Install using [Homebrew][3]

```shell
$ brew tap lukasmalkmus/tap
$ brew install cwa-qr
```

To update:

```shell
$ brew upgrade cwa-qr
```

  [3]: https://brew.sh

### Install using `go get`

```shell
$ go get github.com/lukasmalkmus/cwa-qr
```

### Install from source

```shell
$ git clone https://github.com/lukasmalkmus/cwa-qr.git
$ cd cwa-qr
$ make install # Build and install binary into $GOPATH
```

### Run the Docker image

Docker images are available on [DockerHub][docker].

```shell
$ docker pull lukasmalkmus/cwa-qr
$ docker run lukasmalkmus/cwa-qr
```

### Validate installation

In all cases the installation can be validated by running `cwa-qr -version` in the
terminal:

```shell
Corona Warn App QR-Code Generator version 1.0.0
```

## Usage

```shell
$ cwa-qr [flags] <output-file>
```

Most basic usage:

```shell
$ cwa-qr event.png
```

Flags are optional. If not provided, the application questions for input on the
command line.

### Library usage

```go
import cwaqr "github.com/lukasmalkmus/cwa-qr"

// ...

qrCode, err := cwaqr.GenerateQRCode(cwaqr.Event{
	// ...
})

// Write qrCode to file, etc.
```

## Contributing

Feel free to submit PRs or to fill issues. Every kind of help is appreciated.

Before committing, `make` should run without any issues.

## License

&copy; Lukas Malkmus, 2021

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.

[![License Status][license_status_badge]][license_status]

<!-- Badges -->

[gopkg]: https://pkg.go.dev/github.com/lukasmalkmus/cwa-qr
[gopkg_badge]: https://img.shields.io/badge/doc-reference-007d9c?logo=go&logoColor=white&style=flat-square
[go_workflow]: https://github.com/lukasmalkmus/cwa-qr/actions?query=workflow%3Ago
[go_workflow_badge]: https://img.shields.io/github/workflow/status/lukasmalkmus/cwa-qr/go?style=flat-square&ghcache=unused
[coverage]: https://codecov.io/gh/lukasmalkmus/cwa-qr
[coverage_badge]: https://img.shields.io/codecov/c/github/lukasmalkmus/cwa-qr.svg?style=flat-square&ghcache=unused
[report]: https://goreportcard.com/report/github.com/lukasmalkmus/cwa-qr
[report_badge]: https://goreportcard.com/badge/github.com/lukasmalkmus/cwa-qr?style=flat-square&ghcache=unused
[release]: https://github.com/lukasmalkmus/cwa-qr/releases/latest
[release_badge]: https://img.shields.io/github/release/lukasmalkmus/cwa-qr.svg?style=flat-square&ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/lukasmalkmus/cwa-qr.svg?color=blue&style=flat-square&ghcache=unused
[license_status]: https://app.fossa.com/projects/git%2Bgithub.com%2Flukasmalkmus%2Fcwa-qr
[license_status_badge]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Flukasmalkmus%2Fcwa-qr.svg?type=large&ghcache=unused
[docker]: https://hub.docker.com/r/lukasmalkmus/cwa-qr
[docker_badge]: https://img.shields.io/docker/pulls/lukasmalkmus/cwa-qr.svg?style=flat-square&ghcache=unused
