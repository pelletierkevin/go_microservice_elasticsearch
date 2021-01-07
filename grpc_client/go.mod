module github.com/pelletierkevin/go_microservice_elasticsearch/grpc_client

go 1.15


replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch => ../elastic_health/elasticsearch

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health => ../elastic_health/grpc_health



require (
	github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health v0.0.0-20210106181153-86c3c23d5b66
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
)
