FROM golang:1.20-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/asynq

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/asynq .

CMD ["./asynq"]