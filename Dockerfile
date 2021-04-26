# Production image based on alpine.
FROM alpine
LABEL maintainer="Lukas Malkmus <mial@lukasmalkmus.com>"

# Upgrade packages and install ca-certificates.
RUN apk update --no-cache
RUN apk upgrade --no-cache
RUN apk add --no-cache ca-certificates

# Copy binary into image.
COPY cwa-qr /usr/bin/cwa-qr

# Use the project name as working directory.
WORKDIR /cwa-qr

# Set the binary as entrypoint.
ENTRYPOINT [ "/usr/bin/cwa-qr" ]
