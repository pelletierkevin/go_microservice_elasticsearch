module github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health

go 1.15

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch => ../elasticsearch

require (
	github.com/golang/protobuf v1.4.3
	github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch v0.0.0-20210106143910-bf4457c629a0
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
