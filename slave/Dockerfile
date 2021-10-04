FROM golang:1.17 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o file-proc-slave

FROM alpine:3.14
WORKDIR /app
# RUN apk add --update --no-cache ca-certificates curl
COPY --from=builder /app/file-proc-slave .

ENTRYPOINT ["/app/file-proc-slave"]
