FROM centos:centos7.9.2009

ENV \
    TERM=xterm-color \
    TIME_ZONE="UTC" \
    CGO_ENABLED=0 \
    GO_VERSION=1.17.1 \
    GOOS=linux \
    GOARCH=arm64 \
    GOFLAGS="-mod=vendor"

ADD https://dl.google.com/go/go$GO_VERSION.linux-$GOARCH.tar.gz /tmp/
RUN tar -xzf /tmp/go$GO_VERSION.linux-$GOARCH.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go$GO_VERSION.linux-$GOARCH.tar.gz \
    && ln -s /usr/local/go/bin/go /bin/go

WORKDIR /work
ADD go.* .
ADD pkg .
ADD vendor .

CMD go test ./...