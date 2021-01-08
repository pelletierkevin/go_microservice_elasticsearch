// Package grpc_health implements methods to start a gRPC server and implements the different
// service endpoints defined in the server.proto file.
package grpc_health

import (
	"log"
	"net"
	"strconv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch"
)

type Server struct {
	ClusterHostname string 
	ClusterPort string
}

// StartGrpcServerOnPort will prepare and start the gRPC server exposing the 
// different services endpoints. Note that this is a blocking function and will never end unless 
// an error occur.
//
// StartGrpcServerOnPort("9000", "127.0.0.1", "9200")
//
func StartGrpcServerOnPort(grpcPort string, clusterHostname string, clusterPort string) {
	
	log.Printf("Start gRPC server")

	lis, err := net.Listen("tcp", ":" + grpcPort)
    if err != nil {
        log.Fatalf("Failed to listen on port %s :  %v", grpcPort, err)
    }

    s := Server{clusterHostname, clusterPort}

    grpcServer := grpc.NewServer()

    RegisterElasticServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server : %s", err)
    }
    
}

// SayHello implements the gRPC service endpoint defined in the proto configuration.
// It is mainly a check to verify the gRPC server is reachable from the client. 
//
// SayHello(context.Background(), &Message{Body: "My message"})
//
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server! "}, nil
}

// GetClusterStatus implements the gRPC service endpoint defined in the proto configuration.
// It will return a ClusterInfo struct defined in server.pb.go. Note that the ClusterInfo defined in server.pb.go
// only contains essential informations and is different than the one defined in the "elasticsearch/health.go" module.
// The in *Message argument is not important here.
//
// GetClusterStatus(context.Background(), &Message{Body: "My message"})
//
func (s *Server) GetClusterStatus(ctx context.Context, in *Message) (*ClusterInfo, error) {
	log.Printf("Received GetClusterStatus request")

	resp, err := elasticsearch.GetClusterHealth(s.ClusterHostname, s.ClusterPort)
    if err != nil {
        log.Fatal(err)
    }
	
	var clusterInfo *ClusterInfo
	clusterInfo = new(ClusterInfo)
	clusterInfo.Name = resp.Name
	clusterInfo.Status = resp.Status
	clusterInfo.Nodes = resp.Node_total

	return clusterInfo, nil
}

// GetIndiceStatus implements the gRPC service endpoint defined in the proto configuration.
// It will return a IndiceInfo struct defined in server.pb.go. Note that the IndiceInfo defined in server.pb.go
// only contains essential informations and is different than the one defined in the "elasticsearch/health.go" module.
//
// GetIndiceStatus(context.Background(), &IndiceName{Indicename: "myindex"})
//
func (s *Server) GetIndiceStatus(ctx context.Context, in *IndiceName) (*IndiceInfo, error) {
	log.Printf("Received GetIndiceStatus request")

	resp, err := elasticsearch.GetHealthOfIndice(s.ClusterHostname, s.ClusterPort, in.Indicename)
    if err != nil {
        log.Fatal(err)
    }
	
	var indiceInfo *IndiceInfo
	indiceInfo = new(IndiceInfo)
	indiceInfo.Indicename = resp.Index
	indiceInfo.Status = resp.Status
	indiceInfo.Health = resp.Health
	indiceInfo.Uuid = resp.Uuid

	return indiceInfo, nil
}

// GetIndicesList implements the gRPC service endpoint defined in the proto configuration.
// It will return an array of IndiceInfo struct defined in server.pb.go. Note that the IndiceInfo defined in server.pb.go
// only contains essential informations and is different than the one defined in the "elasticsearch/health.go" module.
// The in *Message argument is not important here.
//
// GetIndicesList(context.Message(), &Message{Body: "My message"})
//
func (s *Server) GetIndicesList(ctx context.Context, in *Message) (*ListIndices, error) {
	log.Printf("Received GetIndicesList request")

	resp, err := elasticsearch.GetClusterIndices(s.ClusterHostname, s.ClusterPort)
    if err != nil {
        log.Fatal(err)
    }
	
	var listIndice *ListIndices
	listIndice = new(ListIndices)
	listIndice.NbIndices = strconv.Itoa(len(resp))
	
	var listOfPointers []*IndiceInfo
	for _, element := range resp {
		var indiceInfo *IndiceInfo
		indiceInfo = new(IndiceInfo)
		indiceInfo.Indicename = element.Index
		indiceInfo.Status = element.Status
		indiceInfo.Health = element.Health
		indiceInfo.Uuid = element.Uuid

		listOfPointers = append(listOfPointers, indiceInfo)

	}

	listIndice.Indicelist = listOfPointers

	return listIndice, nil
}

// CreateIndexInCluster implements the gRPC service endpoint defined in the proto configuration.
// It will create an index in the elasticsearch cluster. It returns a Message which will contain the 
// response of the request.
//
// CreateIndexInCluster(context.Message(), &IndiceName{Indicename: "myindex"})
//
func (s *Server) CreateIndexInCluster(ctx context.Context, in *IndiceName) (*Message, error) {
	log.Printf("Received CreateIndexInCluster request")

	response, err := elasticsearch.CreateIndexInCluster(s.ClusterHostname, s.ClusterPort, in.Indicename)
    if err != nil {
        log.Fatal(err)
    }
	return &Message{Body: "Index succesfully created. " + response}, nil
}

// DeleteIndexInCluster implements the gRPC service endpoint defined in the proto configuration.
// It will delete an index in the elasticsearch cluster. It returns a Message which will contain the 
// response of the request.
//
// DeleteIndexInCluster(context.Message(), &IndiceName{Indicename: "myindex"})
//
func (s *Server) DeleteIndexInCluster(ctx context.Context, in *IndiceName) (*Message, error) {
	log.Printf("Received DeleteIndexInCluster request")

	response, err := elasticsearch.DeleteIndexInCluster(s.ClusterHostname, s.ClusterPort, in.Indicename)
    if err != nil {
        log.Fatal(err)
    }
	return &Message{Body: "Index succesfully deleted. " + response}, nil
}
