# syntax=docker/dockerfile:1

FROM golang:1.22.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /http-simple-binary

EXPOSE 8090

CMD ["/http-simple-binary"]