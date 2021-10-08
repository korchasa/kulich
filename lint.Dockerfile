FROM golang:1.17.1-alpine3.14 as build

ENV \
    TERM=xterm-color \
    TIME_ZONE="UTC" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64 \
    GOLANGCI_VERSION="1.42.1" \
    GOLANGCI_HASHSUM="0fbb58f36933b502bc841f8b28a5c609ac030d3a843fe1ea2dce2cee3a2b0d10"

RUN \
    echo "## Prepare timezone" && \
    apk add --no-cache --update tzdata coreutils && \
    cp /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime && \
    echo "${TIME_ZONE}" > /etc/timezone && date

RUN echo "## Install golangci"
ADD https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCI_VERSION}/golangci-lint-${GOLANGCI_VERSION}-linux-${GOARCH}.tar.gz ./golangci-lint.tar.gz
RUN echo "${GOLANGCI_HASHSUM}  golangci-lint.tar.gz" | sha256sum -c -
RUN tar -xzf golangci-lint.tar.gz
RUN cp ./golangci-lint-${GOLANGCI_VERSION}-linux-${GOARCH}/golangci-lint /usr/bin/
RUN golangci-lint --version

WORKDIR /work
ADD go.* .
ADD pkg pkg
ADD vendor vendor
ADD .golangci.yml .golangci.yml

CMD golangci-lint run --timeout 3m --color always --verbose --out-format colored-line-number
