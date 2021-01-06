FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o elastic_health elastic_health.go

ENV clusterhostname localhost

ENV clusterport 9200

EXPOSE 9000

ENTRYPOINT ["./elastic_health"]