# ElasticHealth Go Microservice

Microservice written in Go to retrieve and deliver the health, status and the indices of an Elasticsearch cluster using gRPC endpoints.

![alt text](https://github.com/pelletierkevin/go_microservice_elasticsearch/blob/main/elastichealth_schema.png?raw=true)

# The project

The project is a microservice written in Golang exposing gRPC endpoints. Those endpoints allow a client program to retrieve informations about an Elasticsearch cluster like health, status and its indices. The global implementation is wrapped around a Docker container. The docker image can be found in Docker hub here : https://hub.docker.com/repository/docker/kevinplltr/elastic-health. Additionally, the project can be used in a Kubernetes/Openshift environment, as it provides a usable Helm chart running this docker image. 

The microservice is launched by passing the elasticsearch cluster hostname:port parameter. There is a one-to-one relationship between the microservice and an elasticsearch cluster. In order to use the microservice on another elasticsearch cluster you need to execute a 2nd instance of the elastichealth microservice.

# Features

Currently, the microservice allows a client to call his gRPC endpoints and offers the following functionalities : 

- Get the elasticsearch cluster health/status, name and the number of nodes running.
- Get the list of indices associated with this elasticsearch cluster.
- Get the health, status and uuid of an indice of the cluster. (Given the index name)
- Create a new index in the elasticsearch cluster
- Delete an index in the elasticsearch cluster

![alt text](https://github.com/pelletierkevin/go_microservice_elasticsearch/blob/main/client_example1.png?raw=true)


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



# Sources
Articles used to develop this project. 
gRPC
Elastic search
Go
Helm