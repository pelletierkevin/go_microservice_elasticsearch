module github.com/pelletierkevin/go_microservice_elasticsearch/grpc

go 1.15

replace github.com/tutorialedge/go-grpc-beginners-tutorial/chat => ../../chat

replace github.com/pelletierkevin/go_microservice_elasticsearch/elasticsearch => ../elasticsearch

require (
	github.com/tutorialedge/go-grpc-beginners-tutorial/chat v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.34.0
)
