## Build
FROM golang:1.21.3-bullseye as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

## Deploy
FROM alpine:3.15
WORKDIR /
COPY --from=build /app/main /usr/bin/
COPY templates /templates
ENTRYPOINT ["main"]
