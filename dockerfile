# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

# The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, 
# COPY and ADD instructions that follow it in the Dockerfile. If the WORKDIR doesn’t 
# exist, it will be created even if it’s not used in any subsequent Dockerfile instruction.
WORKDIR /app

# Copy all files to builder workdir
COPY . .

# Build goapp
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goapp cmd/goapp/main.go

##
## Deploy
##
FROM golang:1.16-alpine

WORKDIR /

# Copy app from builder to deploymennt container
COPY --from=build /app/goapp /goapp
COPY --from=build /app/configs/secrets.json /configs/secrets.json
COPY --from=build /app/schema /schema
# Export port8080
EXPOSE 8080

ENTRYPOINT ["/goapp"]