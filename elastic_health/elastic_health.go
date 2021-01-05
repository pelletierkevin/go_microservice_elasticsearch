package main

import (
    "fmt"
    "log"
    "os"
    "net"
    "google.golang.org/grpc"
    "github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch"
    "github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health"
)

func CallElasticSearchHealthFunc(hostnameCluster string, portCluster string) {
    fmt.Println("Hostname of Elastic Search cluster : " + hostnameCluster)
    fmt.Println("Port of Elastic Search cluster : " + portCluster)

    // Request greeting messages for the names.
    clusterInfo, err := elasticsearch.GetClusterHealth(hostnameCluster, portCluster)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(clusterInfo)

    // Request greeting messages for the names.
    resp, err := elasticsearch.GetClusterIndices(hostnameCluster, portCluster)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    hostnameCluster := os.Args[1]
    portCluster := os.Args[2]

    fmt.Println("Go gRPC Elasticsearch Health Microservice!")
  
    // GRPC ---------------
    lis, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc_health.Server{"localhost", "9200"}

    grpcServer := grpc.NewServer()

    grpc_health.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
    }
    
    // ----------------------------

    if (len(os.Args) > 3) {
        indiceName := os.Args[3]

        // Request greeting messages for the names.
        resp, err := elasticsearch.GetHealthOfIndice(hostnameCluster, portCluster, indiceName)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(resp)
    }


}
