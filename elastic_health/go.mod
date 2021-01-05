module elasticsearch_health

go 1.15

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch => ./elasticsearch

replace github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health => ./grpc_health

require (
	github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch v0.0.0-00010101000000-000000000000
	github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.34.0
)
