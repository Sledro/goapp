# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, 
# COPY and ADD instructions that follow it in the Dockerfile. If the WORKDIR doesn’t 
# exist, it will be created even if it’s not used in any subsequent Dockerfile instruction.
WORKDIR /

# Copy all files to builder workdir
COPY go.mod ./
COPY go.sum ./

# Download necessary Go modules
RUN go mod download

COPY . .

# Build goapp
RUN go build -o build/goapp cmd/goapp/main.go

##
## Deploy
##
FROM golang:1.16-alpine

WORKDIR /

# Copy app from builder to deploymennt container
COPY --from=build /goapp /goapp

# Export port8080
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/goapp"]