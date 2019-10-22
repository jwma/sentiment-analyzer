# Build stage
FROM golang:1.12 as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.io
COPY go.mod go.sum ./
RUN go mod download
COPY webapp.go ./
ENV CGO_ENABLED=0
RUN go build -o webapp .

# Final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/webapp /app/
EXPOSE 8080
CMD ["./webapp"]
