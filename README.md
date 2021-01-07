# ElasticHealth Go Microservice

Microservice written in Go to retrieve and deliver the health, status and the indices of an Elasticsearch cluster using gRPC endpoints.

![alt text](https://github.com/pelletierkevin/go_microservice_elasticsearch/blob/main/elastichealth_schema.png?raw=true)

# The project

The project is a microservice written in Golang exposing gRPC endpoints. Those endpoints allow a client program to retrieve informations about an Elasticsearch cluster like health, status and its indices. The global implementation is wrapped around a Docker container. The docker image can be found in Docker hub here : https://hub.docker.com/repository/docker/kevinplltr/elastic-health. Additionally, the project can be used in a Kubernetes/Openshift environment, as it provides a usable Helm chart running this docker image. 

The microservice is launched by passing the elasticsearch cluster hostname:port parameter. There is a one-to-one relationship between the microservice and an elasticsearch cluster. In order to use the microservice on another elasticsearch cluster you need to execute a 2nd instance of the elastichealth microservice.

---

# Features

Currently, the microservice allows a client to call his gRPC endpoints and offers the following functionalities : 

- Get the elasticsearch cluster health/status, name and the number of nodes running.
- Get the list of indices associated with this elasticsearch cluster.
- Get the health, status and uuid of an indice of the cluster. (Given the index name)
- Create a new index in the elasticsearch cluster
- Delete an index in the elasticsearch cluster

The gRPC server will use the port 9000. 

![Client example caption="Example of client calling the elastichealth gRPC microservice to get the cluster status"](https://github.com/pelletierkevin/go_microservice_elasticsearch/blob/main/client_example1.png?raw=true)

---

# Requirements

- Elasticsearch installed. 
    - Install Elasticsearch here : https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html
    - Start an elastic search cluster (i.e. run the `elasticsearch`command). Get the hostname and port of this cluster. 

- To modify the project you'll have to install Go and have a Go environment set up.
    - Install Go here : https://golang.org/doc/install

- Docker installed
    - Install Docker here : https://docs.docker.com/get-docker/

- Helm installed
    - Install Helm here : https://helm.sh/docs/intro/install/

- Protoc installed (protobuf)
    - Get the latest version here : https://github.com/protocolbuffers/protobuf/tags and download the binary.
    - Specify the `protoc` executable in yourr `$PATH` 

---

# How to run

In folder elastic_health there is the executable for the microservice

First of all start an elasticsearch cluster. You will need the hostname and the port (usually 9200) of your cluster. 

## Run the elastichealth microservice

###     - Using the Go executable directly
  Go the the `elastic_health` folder
  `./elasticsearch <cluster hostname> <cluster port>`

  Example : `./elasticsearch 127.0.0.1 9200`

###     - Using the Docker image
  - You can either pull the docker image from dockerhub : `docker pull kevinplltr/elastic-health:0.1.0`
  - Or build the Dockerfile : `docker build -t kevinplltr/elastic-health:0.1.0 .`

  Then run the docker container : `docker run -p 9000:9000 kevinplltr/elastic-health:0.0.1 <cluster hostname> <cluster port>`

  Example : `docker run -p 9000:9000 kevinplltr/elastic-health:0.0.1 host.docker.internal 9200` if the cluster is running on the host machine.

###     - Using the Helm chart
  - First make sure you can connect to your Kubernetes cluster. 
  - Execute the command : `helm install -n helm-chart-elastic ./helm-chart-elastic`
  - Verify the pod is running : `kubectl get pods`

## Run the gRPC client

Once you have an elasticsearch cluster and the elastichealth microservice running you can launch a client to call the gRPC endpoints. 
In the `grpc_client` folder a client executable is ready to be used. To use it : 
  - `./client <grpc hostname> clusterhealth`
  - `./client <grpc hostname> listindices`
  - `./client <grpc hostname> indexhealth <index name>`
  - `./client <grpc hostname> createindex <index name>`
  - `./client <grpc hostname> deleteindex <index name>`

The gRPC hostname will be the IP address where the microservice is running. 
Example : `./client 127.0.0.1 clusterhealth`

---

# How to build
Go build

---

# Components

## - I) Elastichealth Microservice
elastic_health folder

### - I.1) Go Project

This folder contains the Go project running the microservice and the Dockerfile definition. 
There are 2 modules in this project. The first one `elasticsearch` is responsible of making the calls to the elasticsearch API. The second one `grpc_health` is defining the 
gRPC server and its different services/endpoints. Those 2 modules are called by the `elastic_health.go` main interface. 

The project tries to use the standard Go packages : https://golang.org/pkg/ instead of using fancy modules. (exception for gRPC module still provided by golang)

#### elasticsearch
- **health.go** (Retrieve the health of a cluster, and the health of a specified index)
- **indices.go** (List all the indexes in a cluster, create an index, delete an index)
- **health_test.go** (Unittests, covering the different methods calling the elasticsearch API)
#### grpc_health
- **server.go** (create the gRPC server, define the different endpoints)
- **server.pb.go** (auto-generated by protoc)
- **server_test.go** (Unittests, covering an end to end simulation of a gRPC client, making calls to the gRPC server to get elasticsearch informations.)
#### Main file elastic_health.go
- **elastic_health.go** (main method, which will call the gRPC module)
- **elastic_health_test.go** (Unittests, covering the parsing of arguments, hostname and port)
- **elastic_health** (executable, built from elastic_health.go)

### - I.2) gRPC protobuf
- **server.proto**
This file defines the different structure, services and endpoints provided by the gRPC module.

To compile and regenerate the `grpc_health/server.pb.go` file use this command : 
`protoc --go_out=plugins=grpc:grpc_health server.proto`

### - I.3) Dockerfile
- **Dockerfile**
The dockerfile is copying the Golang project, building it and set the `elastic_health` executable as an entrypoint. When running the docker image, you'll have to specify the arguments, like when you run the executable. 

Build the docker image : 
`docker build -t kevinplltr/elastic-health:0.1.0 .`

Run the docker image : 
`docker run -p 9000:9000 kevinplltr/elastic-health:0.0.1 <cluster hostname> <cluster port >`

The docker image is published in Dockerhub : 
https://hub.docker.com/repository/docker/kevinplltr/elastic-health
Last tag version : 0.1.0

## - II) gRPC Client executable written in Go
grpc_client folder
This folder is composed by the `client.go` file which is basically able to call the different gRPC endpoints of our microservice given its hostname. 
- **client.go** (Composed of a main method which will make different endpoints calls based on the arguments)
  - `./client <grpc hostname> clusterhealth`
  - `./client <grpc hostname> listindices`
  - `./client <grpc hostname> indexhealth <index name>`
  - `./client <grpc hostname> createindex <index name>`
  - `./client <grpc hostname> deleteindex <index name>`

## - III) Helm chart
helm-chart-elastic folder
This folder defines the different files required to use a Helm Chart which will run the microservice (using the docker image) on Kubernetes or Openshift.

In values.yaml we define the docker image which is the one stored in dockerhub `kevinplltr/elastic-health` and we specify the gRPC port which is 9000.

---

# Improvements & Ideas

- Log/Store real time info, + previous info (1hour, 1day)
- Alert message (email, slack) when the health cluster is red
- Make periodic calls (monitor every 30 seconds for example)
- Enable encrypted gRPC communication (using certificate)
- Enable uses of an https elasticsearch cluster
- Implement a middleware to cache and log the requests
- Implement a background job checking the existing indices and deleting the old ones using the timestamp. 
- Add the timestamp information in the gRPC requests
- Implement a middleware to check and validate the arguments given for a request
- Implement additional functionalities using the elasticsearch API 
- Give a specific Json when creating an Index to define the number of replicas and other settings.

---

# Sources

## - gRPC
- gRPC package in Golang : https://godoc.org/google.golang.org/grpc
- Writing a microservice in Golang using gRPC : https://bitbucket.org/blog/writing-a-microservice-in-golang-which-communicates-over-grpc
- gRPC in Go Tutorial : https://bitbucket.org/blog/writing-a-microservice-in-golang-which-communicates-over-grpc

## - Elasticsearch
- Get Elasticsearch up and running : https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started-install.html
- Cluster health API : https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
- Create index API : https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
- Delete index API : https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
- Cluster indices API : https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html

## - Helm
- Deploy a Go application with Helm : https://docs.bitnami.com/tutorials/deploy-go-application-kubernetes-helm/
- Create Helm Chart : https://docs.bitnami.com/tutorials/create-your-first-helm-chart




