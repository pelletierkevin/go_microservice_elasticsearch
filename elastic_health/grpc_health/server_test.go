package grpc_health


import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)


func TestGrpcGetClusterStatusWithClientE2E(t *testing.T) {
	log.Printf("- TestGrpcSayHello - ")

	grpcPort := "9000"
	testHostname := "127.0.0.1"
	testPort := "49200"

	StartAsyncGRPCServer(grpcPort, testHostname, testPort)

	c, conn:= CreateGrpcClient(grpcPort)
	defer conn.Close()

	expectedJsonResponse := "[{\"epoch\" : \"1609867421\",\"timestamp\" : \"17:23:41\",\"cluster\" : \"test_cluster\",\"status\" : \"yellow\",\"node.total\" : \"1\",\"node.data\" : \"1\",\"shards\" : \"3\",\"pri\" : \"3\",\"relo\" : \"0\",\"init\" : \"0\",\"unassign\" : \"6\",\"pending_tasks\" : \"0\",\"max_task_wait_time\" : \"-\",\"active_shards_percent\" : \"33.3%\"}]"

	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	var clusterinfo *ClusterInfo
	clusterinfo, err := c.GetClusterStatus(context.Background(), &Message{Body: "Asking Cluster status"})
	if err != nil {
		log.Fatalf("Error when calling GetClusterStatus: %s", err)
	}

	if( clusterinfo.Name != "test_cluster") {
		t.Errorf("Error the expected cluster name is %s but received %s ", "test_cluster", clusterinfo.Name)
	}
	if( clusterinfo.Status != "yellow") {
		t.Errorf("Error the expected cluster status is %s but received %s ", "yellow", clusterinfo.Status)
	}
	if( clusterinfo.Nodes != "1") {
		t.Errorf("Error the expected cluster number of nodes is %s but received %s ", "1", clusterinfo.Nodes)
	}
}


func StartAsyncGRPCServer(grpcPort string, clusterHostname string, clusterPort string) {
	
		log.Printf("Start gRPC server")
	
		lis, err := net.Listen("tcp", ":" + grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen on port %s :  %v", grpcPort, err)
		}
	
		s := Server{clusterHostname, clusterPort}
	
		grpcServer := grpc.NewServer()
	
		RegisterElasticServiceServer(grpcServer, &s)
	
		go grpcServer.Serve(lis)
		
}

func CreateGrpcClient(grpcPort string) (ElasticServiceClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error, did not connect to the gRPC server: %s", err)
	} else {
		log.Printf("Connected to gRPC server - Health/Status for Elasticsearch cluster")
	}

	c := NewElasticServiceClient(conn)

	return c, conn
}

// The same method is defined in health_test.go. Needs to define a test_util class with this method. 
func MockHttpResponseFromCluster(expectedJsonResponse string, clusterHostname string, clusterPort string) (*httptest.Server) {

	// Start a local HTTP server
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(expectedJsonResponse))
	}) 

	// create a listener with the desired port.
	l, err := net.Listen("tcp", clusterHostname +":"+ clusterPort)
	if err != nil {
		log.Fatal(err)
	}

	testserver := httptest.NewUnstartedServer(handler)

	// NewUnstartedServer creates a listener. Close that listener and replace 
	// with the one we created.
	testserver.Listener.Close()
	testserver.Listener = l

	// Start the server.
	testserver.Start()

	return testserver
}