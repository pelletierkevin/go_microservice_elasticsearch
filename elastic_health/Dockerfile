FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o elastic_health elastic_health.go

EXPOSE 9000

CMD ["./elastic_health"]