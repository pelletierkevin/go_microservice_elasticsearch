module elasticsearch_health

go 1.15

replace github.com/pelletierkevin/go_microservice_elasticsearch/grpc => ./grpc

replace github.com/pelletierkevin/go_microservice_elasticsearch/elasticsearch => ./elasticsearch

replace github.com/tutorialedge/go-grpc-beginners-tutorial/chat => ../chat

require (
	github.com/pelletierkevin/go_microservice_elasticsearch/elasticsearch v0.0.0-00010101000000-000000000000
	github.com/pelletierkevin/go_microservice_elasticsearch/grpc v0.0.0-00010101000000-000000000000
	github.com/tutorialedge/go-grpc-beginners-tutorial/chat v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
)
