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

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error, did not connect to the gRPC server: %s", err)
	} else {
		log.Printf("Connected to gRPC server - Health/Status for Elasticsearch cluster")
	}
	defer conn.Close()

	c := grpc_health.NewElasticServiceClient(conn)

	modeChoice := os.Args[1]
	var response *grpc_health.Message

	if( modeChoice == "sayhello") {
		response, err = c.SayHello(context.Background(), &grpc_health.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling SayHello or SayBonjour: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
	} else if ( modeChoice == "clusterhealth" ) {
		var clusterinfo *grpc_health.ClusterInfo
		clusterinfo, err = c.GetClusterStatus(context.Background(), &grpc_health.Message{Body: "Asking Cluster status"})

		if err != nil {
			log.Fatalf("Error when calling GetClusterStatus: %s", err)
		}
		log.Printf("Response from server: %s", clusterinfo.Name)
		log.Printf("Cluster name: %s", clusterinfo.Name)
		log.Printf("Cluster status: %s", clusterinfo.Status)
		log.Printf("Cluster nb nodes: %s", clusterinfo.Nodes)
	} else if ( modeChoice == "indicehealth" ) {
		indiceName := os.Args[2]
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
	} else if ( modeChoice == "listindices" ) {
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

	} else if ( modeChoice == "createindex" ) {
		indiceName := os.Args[2]
		response, err = c.CreateIndexInCluster(context.Background(), &grpc_health.IndiceName{Indicename: indiceName})
		if err != nil {
			log.Fatalf("Create index request response : %s", err)
		}
		log.Printf("Create index request succeeded. %s", response.Body)
	} else if ( modeChoice == "deleteindex" ) {
		indiceName := os.Args[2]
		response, err = c.DeleteIndexInCluster(context.Background(), &grpc_health.IndiceName{Indicename: indiceName})
		if err != nil {
			log.Fatalf("Error when deleting index in Cluster: %s", err)
		}
		log.Printf("Delete index request response : %s", response.Body)
	} 

}