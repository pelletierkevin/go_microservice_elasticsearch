package main

import (
    "os"
    "log"
    "errors"
    "strconv"
    "github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/grpc_health"
    "github.com/pelletierkevin/go_microservice_elasticsearch/elastic_health/elasticsearch"
)

func main() {

    clusterHostname, clusterPort, err := GetClusterHostnameAndPortFromArguments()
    if(err != nil) {
        log.Fatalf("Command error. Please use the commmand like this : %s %s %s", os.Args[0], "<cluster hostname>", "<cluster port>")
    }

    log.Printf(" ------- Go gRPC Elasticsearch Health Microservice! ------- ")
  
    elasticsearch.CanConnectToElasticSearchCluster(clusterHostname, clusterPort)

    grpc_health.StartGrpcServerOnPort("9000", clusterHostname, clusterPort)

}

func GetClusterHostnameAndPortFromArguments() (string, string, error)  {
    if len(os.Args) != 3 {
        return "", "", errors.New("Command error. Missing arguments or more arguments than required.")
    } else {
        _, err := strconv.Atoi(os.Args[2])
        if(err != nil) {
            return os.Args[1], os.Args[2], errors.New("Command error. The second argument should represent the port number and can only contains digits.")
        }
    }
    return os.Args[1], os.Args[2], nil
}
