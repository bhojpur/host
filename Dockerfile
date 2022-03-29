FROM golang:1.17

RUN apt-get update && apt-get install -y --no-install-recommends \
                openssh-client \
                rsync \
                fuse \
                sshfs \
        && rm -rf /var/lib/apt/lists/*

RUN go get  github.com/golang/lint/golint \
            github.com/mattn/goveralls \
            golang.org/x/tools/cover

ENV USER root
WORKDIR /go/src/github.com/bhojpur/host

COPY . ./
RUN mkdir bin