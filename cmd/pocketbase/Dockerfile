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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/pocketbase ./cmd/pocketbase

###################################################################################################
######### PocketBase Deploy Stage  #######
FROM alpine:latest as pocketbase

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/pocketbase .
# uncomment to copy the local pb_migrations dir into the image
# COPY ./pb_migrations /pb/pb_migrations
# uncomment to copy the local pb_hooks dir into the image
# COPY ./pb_hooks /pb/pb_hooks

# Set the SECRET environment variable
ENV SECRET=my_secret_value

# Expose port 8090 to the outside
EXPOSE 8090

VOLUME /app/pb_data

# Command to run the executable
CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090"]