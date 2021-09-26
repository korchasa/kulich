FROM centos:7

ENV GO_VERSION 1.17
ENV GO_ARCH arm64

WORKDIR /

ADD https://dl.google.com/go/go$GO_VERSION.linux-$GO_ARCH.tar.gz /tmp/
RUN tar -xzf /tmp/go$GO_VERSION.linux-$GO_ARCH.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go$GO_VERSION.linux-$GO_ARCH.tar.gz \
    && ln -s /usr/local/go/bin/go /bin/go

RUN yum update -y \
    && yum -y install gcc

WORKDIR /work

ADD . .
CMD go test ./...