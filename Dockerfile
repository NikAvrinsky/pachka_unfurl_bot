FROM golang:1.23.2 as builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o apiserver -v ./cmd
FROM alpine as runner
WORKDIR /app
COPY --from=builder /go/apiserver .
RUN chmod +x /app/apiserver
ENTRYPOINT ["/app/apiserver"]
EXPOSE 8080