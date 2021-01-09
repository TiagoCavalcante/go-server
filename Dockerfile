FROM golang:latest

WORKDIR /go
COPY src ./src

RUN go build src/main.go

CMD ["./main"]