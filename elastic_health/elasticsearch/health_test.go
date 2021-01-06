package elasticsearch

import (
	"net"
	"net/http"
	"net/http/httptest"
	"log"
	"testing"
)


func TestGetClusterHealthWithCorrectResponse(t *testing.T) {
	log.Printf("- TestGetClusterHealthWithCorrectResponse - ")
	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "[{\"epoch\" : \"1609867421\",\"timestamp\" : \"17:23:41\",\"cluster\" : \"test_cluster\",\"status\" : \"yellow\",\"node.total\" : \"1\",\"node.data\" : \"1\",\"shards\" : \"3\",\"pri\" : \"3\",\"relo\" : \"0\",\"init\" : \"0\",\"unassign\" : \"6\",\"pending_tasks\" : \"0\",\"max_task_wait_time\" : \"-\",\"active_shards_percent\" : \"33.3%\"}]"

	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	CanConnectToElasticSearchCluster(testHostname, testPort)
	clusterinfo, err := GetClusterHealth(testHostname, testPort)
	if err != nil {
        log.Fatal(err)
	} 

	if( clusterinfo.Name != "test_cluster") {
		t.Errorf("Error the expected cluster name is %s but received %s ", "test_cluster", clusterinfo.Name)
	}

}

func TestGetClusterHealthWithWrongResponse(t *testing.T) {
	log.Printf("- TestGetClusterHealthWithWrongResponse - ")

	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "fake_response"
	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	_, err := GetClusterHealth(testHostname, testPort)
	if err == nil {
        t.Errorf("Was expecting an error. It should have thrown an error because of the wrong response sent by the mocked http cluster.")
	} 

}

func TestGetClusterIndicesWithCorrectResponse(t *testing.T) {
	log.Printf("- TestGetClusterIndicesWithCorrectResponse - ")
	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "[{\"health\":\"yellow\",\"status\":\"open\",\"index\":\"my-index-000001\",\"uuid\":\"Q7cE68VoSAa_S_NIFgeZow\",\"pri\":\"3\",\"rep\":\"2\",\"docs.count\":\"0\",\"docs.deleted\":\"0\",\"store.size\":\"624b\",\"pri.store.size\":\"624b\"}]"

	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	listIndices, err := GetClusterIndices(testHostname, testPort)
	if err != nil {
        log.Fatal(err)
	} 

	if( listIndices[0].Health != "yellow") {
		t.Errorf("Error the expected cluster health is %s but received %s ", "yellow", listIndices[0].Health)
	}

}

func TestGetClusterIndicesWithWrongResponse(t *testing.T) {
	log.Printf("- TestGetClusterIndicesWithWrongResponse - ")
	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "fake_response"
	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	_, err := GetClusterIndices(testHostname, testPort)
	if err == nil {
        t.Errorf("Was expecting an error. It should have thrown an error because of the wrong response sent by the mocked http cluster.")
	} 

}

func TestGetHealthOfIndiceWithCorrectResponse(t *testing.T) {
	log.Printf("- TestGetHealthOfIndiceWithCorrectResponse - ")
	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "[{\"health\":\"yellow\",\"status\":\"open\",\"index\":\"my-index-000001\",\"uuid\":\"Q7cE68VoSAa_S_NIFgeZow\",\"pri\":\"3\",\"rep\":\"2\",\"docs.count\":\"0\",\"docs.deleted\":\"0\",\"store.size\":\"624b\",\"pri.store.size\":\"624b\"}]"

	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	indiceInfo, err := GetHealthOfIndice(testHostname, testPort, "my-index-000001")
	if err != nil {
        log.Fatal(err)
	} 

	if( indiceInfo.Health != "yellow") {
		t.Errorf("Error the expected cluster name is %s but received %s ", "yellow", indiceInfo.Health)
	}

}

func TestGetHealthOfIndiceWithWrongResponse(t *testing.T) {
	log.Printf("- TestGetHealthOfIndiceWithWrongResponse - ")
	testHostname := "127.0.0.1"
	testPort := "49200"

	expectedJsonResponse := "fake_response"

	testserver := MockHttpResponseFromCluster(expectedJsonResponse, testHostname, testPort)
	defer testserver.Close()

	_, err := GetHealthOfIndice(testHostname, testPort, "my-index-000001")
	if err == nil {
        t.Errorf("Was expecting an error. It should have thrown an error because of the wrong response sent by the mocked http cluster.")
	} 

}

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