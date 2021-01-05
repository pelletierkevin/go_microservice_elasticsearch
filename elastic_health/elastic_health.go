package main

import (
    "fmt"
    "log"
    "os"
    "net"
    "google.golang.org/grpc"
    "github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health"
)

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

    s := grpc_health.Server{hostnameCluster, portCluster}

    grpcServer := grpc.NewServer()

    grpc_health.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
    }
    
    // ----------------------------


}
