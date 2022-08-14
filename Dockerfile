FROM golang:alpine
RUN apk add build-base

WORKDIR /app
COPY . .
RUN go build -o sms

EXPOSE 8090
