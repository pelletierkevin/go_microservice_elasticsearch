package main

import (
	"log"
	"os"
	"strconv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health"
)

func main() {

	if (len(os.Args) != 3 && len(os.Args) != 4) {
		ErrorCommandArguments()
		return
	}

	grpcHostname := os.Args[1]
	grpcPort := "9000"

	c, conn := CreateGrpcConnection(grpcHostname, grpcPort)
	defer conn.Close()

	modeChoice := os.Args[2]
	var response *grpc_health.Message

	var err error 
	switch modeChoice {

	case "sayhello":
		response, err = c.SayHello(context.Background(), &grpc_health.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling SayHello or SayBonjour: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
	
	case "clusterhealth":
		var clusterinfo *grpc_health.ClusterInfo
		clusterinfo, err = c.GetClusterStatus(context.Background(), &grpc_health.Message{Body: "Asking Cluster status"})

		if err != nil {
			log.Fatalf("Error when calling GetClusterStatus: %s", err)
		}
		log.Printf("Response from server: %s", clusterinfo.Name)
		log.Printf("Cluster name: %s", clusterinfo.Name)
		log.Printf("Cluster status: %s", clusterinfo.Status)
		log.Printf("Cluster nb nodes: %s", clusterinfo.Nodes)
	
	case "indexhealth":
		if (len(os.Args) != 4) {
			ErrorCommandArguments()
			return
		} 
		
		indiceName := os.Args[3]
		var indiceInfo *grpc_health.IndiceInfo
		indiceInfo, err = c.GetIndiceStatus(context.Background(), &grpc_health.IndiceName{Indicename: indiceName})

		if err != nil {
			log.Fatalf("Error when calling GetClusterStatus: %s", err)
		}
		log.Printf("Response from server:")
		log.Printf("Indice name: %s", indiceInfo.Indicename)
		log.Printf("Indice status: %s", indiceInfo.Status)
		log.Printf("Indice health: %s", indiceInfo.Health)
		log.Printf("Indice uuid: %s", indiceInfo.Uuid)
	
	case "listindices":
		var listIndices *grpc_health.ListIndices
		listIndices, err = c.GetIndicesList(context.Background(), &grpc_health.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling GetIndicesList: %s", err)
		}
		log.Printf("Nb indices : %s", listIndices.NbIndices)

		nbIndices, err := strconv.Atoi(listIndices.NbIndices)
		if err != nil {
			log.Fatalf("Error when converting number of indices: %s", err)
		}

		var indiceInfo *grpc_health.IndiceInfo
		for i := 0; i < nbIndices; i++ {
			indiceInfo = listIndices.Indicelist[i]
			log.Printf("Index [ %s ] - Status : %s - Health : %s - Uuid : %s ", indiceInfo.Indicename, indiceInfo.Status, indiceInfo.Health, indiceInfo.Uuid)
		}
	
	case "createindex":
		if (len(os.Args) != 4) {
			ErrorCommandArguments()
			return
		} 
		indiceName := os.Args[3]
		response, err = c.CreateIndexInCluster(context.Background(), &grpc_health.IndiceName{Indicename: indiceName})
		if err != nil {
			log.Fatalf("Create index request response : %s", err)
		}
		log.Printf("Create index request succeeded. %s", response.Body)

	case "deleteindex":
		if (len(os.Args) != 4) {
			ErrorCommandArguments()
			return
		} 
		
		indiceName := os.Args[3]
		response, err = c.DeleteIndexInCluster(context.Background(), &grpc_health.IndiceName{Indicename: indiceName})
		if err != nil {
			log.Fatalf("Error when deleting index in Cluster: %s", err)
		}
		log.Printf("Delete index request response : %s", response.Body)
	
	default:
		ErrorCommandArguments()
		return

	}

}

func ErrorCommandArguments() {
	log.Printf("Command error. Please use one of the following commmand :")
	log.Printf("- %s %s %s", os.Args[0], "<grpc hostname>", "clusterhealth")
	log.Printf("- %s %s %s", os.Args[0], "<grpc hostname>", "listindices")
	log.Printf("- %s %s %s %s", os.Args[0], "<grpc hostname>", "indexhealth", "<index name>")
	log.Printf("- %s %s %s %s", os.Args[0], "<grpc hostname>", "createindex", "<index name>")
	log.Printf("- %s %s %s %s", os.Args[0], "<grpc hostname>", "deleteindex", "<index name>")
}

func CreateGrpcConnection(grpcHostname string, grpcPort string) (grpc_health.ElasticServiceClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHostname+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error, did not connect to the gRPC server: %s", err)
	} else {
		log.Printf("Connected to gRPC server %s:%s - Health/Status for Elasticsearch cluster", grpcHostname, grpcPort)
	}

	c := grpc_health.NewElasticServiceClient(conn)

	return c, conn
}