FROM golang:1.22-alpine
LABEL authors="Akavi"

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]