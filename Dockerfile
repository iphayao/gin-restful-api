FROM golang:1.15.4-alpine3.12 as build_base

RUN apk add --no-cache git

# Set the current working directory
WORKDIR /build

# Populate modules cache based on go.{mod,sum} files
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from current direcotyr to PWD
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start fresh from a smaller image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build_base /build/app .

# This container expose port 8080 to outside
EXPOSE 8080

# RUn the binary application
CMD ["./app"]