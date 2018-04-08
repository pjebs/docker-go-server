FROM golang:1.8

COPY main.go .
RUN go build -o server

ENTRYPOINT ["/go/server"]