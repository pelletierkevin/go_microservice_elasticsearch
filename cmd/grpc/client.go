package main

import (
	"log"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error, did not connect: %s", err)
	} else {
		log.Printf("Connected to gRPC server - Health/Status for Elasticsearch cluster")
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	modeChoice := os.Args[1]
	var response *chat.Message

	if( modeChoice == "1") {
		response, err = c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling SayHello or SayBonjour: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
	} else if( modeChoice == "2" ){
		response, err = c.SayBonjour(context.Background(), &chat.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling SayHello or SayBonjour: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
	} else if ( modeChoice == "3" ) {
		var clusterinfo *chat.ClusterInfo
		clusterinfo, err = c.GetClusterStatus(context.Background(), &chat.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling GetClusterStatus: %s", err)
		}
		log.Printf("Response from server: %s", clusterinfo.Name)
		log.Printf("Cluster name: %s", clusterinfo.Name)
		log.Printf("Cluster status: %s", clusterinfo.Status)
		log.Printf("Cluster nb nodes: %s", clusterinfo.Nodes)
	} else if ( modeChoice == "4" ) {
		indiceName := os.Args[2]
		var indiceInfo *chat.IndiceInfo
		indiceInfo, err = c.GetIndiceStatus(context.Background(), &chat.IndiceName{Indicename: indiceName})

		if err != nil {
			log.Fatalf("Error when calling GetClusterStatus: %s", err)
		}
		log.Printf("Response from server:")
		log.Printf("Indice name: %s", indiceInfo.Indicename)
		log.Printf("Indice status: %s", indiceInfo.Status)
		log.Printf("Indice health: %s", indiceInfo.Health)
		log.Printf("Indice uuid: %s", indiceInfo.Uuid)
	} else if ( modeChoice == "5" ) {
		var listIndices *chat.ListIndices
		listIndices, err = c.GetIndicesList(context.Background(), &chat.Message{Body: "Hello From Client! I'll wait for your reply!!!"})

		if err != nil {
			log.Fatalf("Error when calling GetIndicesList: %s", err)
		}
		log.Printf("Response from server: ")
		log.Printf("Nb indices : %s", listIndices.NbIndices)


		log.Printf("Check first value of listindices : ")
		log.Printf("Indice name: %s", listIndices.Indicelist[0].Indicename)
		log.Printf("Indice status: %s", listIndices.Indicelist[0].Status)

	}


}