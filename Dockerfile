FROM golang:1.22-alpine
LABEL authors="Akavi"

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o main cmd/main.go

CMD ["/app/main"]