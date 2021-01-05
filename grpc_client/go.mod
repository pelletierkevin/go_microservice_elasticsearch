module github.com/pelletierkevin/go_microservice_elasticsearch/grpc_client

go 1.15

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health => ../elastic_health/grpc_health

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch => ../elastic_health/elasticsearch

require (
	github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
)
