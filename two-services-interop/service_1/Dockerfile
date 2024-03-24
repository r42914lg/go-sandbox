# syntax=docker/dockerfile:1

FROM golang:1.22.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY go.sum ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /service-1-binary

EXPOSE 8091

CMD ["/service-1-binary"]