# Multistage docker file
FROM golang:1.16 AS builder

# Create App directory
RUN mkdir /app
# Add everything in the current directory to the app dir
ADD . /app
# Set work directory
WORKDIR /app

# run the build command operating system is linux
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go



FROM alpine:latest AS production
# copy from builder image the app directory to current directory
COPY --from=builder /app .
# kick-off application
CMD ["./app"]