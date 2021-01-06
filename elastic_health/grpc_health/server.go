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

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server! "}, nil
}

func (s *Server) GetClusterStatus(ctx context.Context, in *Message) (*ClusterInfo, error) {
	log.Printf("Receive GetClusterStatus request")

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

func (s *Server) GetIndiceStatus(ctx context.Context, in *IndiceName) (*IndiceInfo, error) {
	log.Printf("Receive GetClusterStatus request")

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

func (s *Server) GetIndicesList(ctx context.Context, in *Message) (*ListIndices, error) {
	log.Printf("Receive GetIndicesList request")

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
