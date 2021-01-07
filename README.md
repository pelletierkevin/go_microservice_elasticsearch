# ElasticHealth Go Microservice

Microservice written in Go to retrieve and deliver the health, status and the indices of an Elasticsearch cluster using gRPC endpoints.

![alt text](https://github.com/pelletierkevin/go_microservice_elasticsearch/blob/main/elastichealth_schema.png?raw=true)

# The project

The project will create a gRPC server exposing mutliple endpoints. Those endpoints allow a client program to : 
- Get the cluster health
- Get the indices in the cluster
- Get the health of an indice of the cluster
- Create an index in the cluster
- Delete an index in the cluster

- elastic_health
- grpc_client
- helm-chart-elastic
- 
## elastic_health

## grpc_client

## helm-chart-elastic

# Requirements
Elasticsearch installed. Deploy an elastic search cluster. Explanations

Use directly the executable

Install Go if you want to modify the project. 

Docker installed

Helm installed

Protoc installed (protobuf)

 # How to build
Go build


# Docker image 

https://hub.docker.com/repository/docker/kevinplltr/elastic-health


# Sources
Articles used to develop this project. 
gRPC
Elastic search
Go
Helm