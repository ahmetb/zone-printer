FROM golang:1.15-alpine as compiler
RUN apk add --no-cache git
WORKDIR /go/src/zoneprinter
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app

FROM alpine:latest
WORKDIR /app
COPY --from=compiler /bin/app ./server
COPY ./templates ./templates
COPY ./static ./static
ENTRYPOINT /app/server
