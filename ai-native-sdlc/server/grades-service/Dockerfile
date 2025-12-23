FROM golang:1.21-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/grades-service ./cmd/grades-service

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app/cmd/grades-service /usr/local/bin/grades-service
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/grades-service"]
