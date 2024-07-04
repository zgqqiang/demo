FROM golang:1.21.4-alpine AS builder
WORKDIR /app
ADD . .
ENV CGO_ENABLED=0
RUN go build -o k8s-service main.go

FROM alpine:latest
workdir /test
COPY --from=builder /app/k8s-service . /test
RUN chmod +x k8s-service
ENTRYPOINT ["./k8s-service"]