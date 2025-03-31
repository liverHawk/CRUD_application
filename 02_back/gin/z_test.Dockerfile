FROM golang:latest

WORKDIR /test

RUN go install github.com/smartystreets/goconvey@latest

CMD ["goconvey", "-host", "0.0.0.0", "-workDir", "/test", "-watchedSuffixes", "_test.go"]