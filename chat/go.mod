module github.com/tutorialedge/go-grpc-beginners-tutorial/chat

go 1.15

replace github.com/pelletierkevin/go_microservice_elasticsearch/elasticsearch => ../cmd/elasticsearch

require (
	github.com/golang/protobuf v1.4.3
	github.com/pelletierkevin/go_microservice_elasticsearch/elasticsearch v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
