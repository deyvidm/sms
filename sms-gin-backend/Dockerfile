# stage 1 - build
# we pull all necessary tools to build our Go binary
FROM golang:1.20-alpine AS build
RUN apk add --no-cache git # absolutely wild. git should come with golang image imo

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/sms


# stage 2 - production
# we pull only what's necessary to run our Go binary
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/sms .

CMD ["./sms"]

EXPOSE 8080
