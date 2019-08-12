FROM golang:1.12

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY $PWD/webapp.go ./
RUN go build -o webapp .
EXPOSE 8080

CMD ["./webapp"]
