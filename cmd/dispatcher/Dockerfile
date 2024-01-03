###################################################################################################
########## GO Apps Build Stage ##########
# This build creates a binary from the go packages in /cmd

# Start from the latest golang base image
FROM golang:alpine as builder

WORKDIR /app
# TODO copy only the files we need
COPY ../.. . 

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go apps
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/dispatcher ./cmd/dispatcher


###################################################################################################
####### Dispatcher Deploy Stage #######
FROM alpine:latest as dispatcher

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/dispatcher .

# Set the SECRET environment variable
# need a better way to inject our secrets here. TOOD: use a secret manager
ENV SECRET="ABC" 

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["./dispatcher"] 