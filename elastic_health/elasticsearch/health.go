// Package elasticsearch implements methods to call Elasticsearch API, retrieve and cast results in proper return values.
// The methods will essentially deliver informations about the health, status and the indices of an Elasticsearch cluster.
package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// The ClusterInfo struct is representing all the fields returned by the Elasticsearch API when
// asking for information about a cluster
type ClusterInfo struct {
	Epoch                 string `json:"epoch"`
	Timestamp             string `json:"timestamp"`
	Name                  string `json:"cluster"`
	Status                string `json:"status"`
	Node_total            string `json:"node.total"`
	Node_data             string `json:"node.data"`
	Shards                string `json:"shards"`
	Pri                   string `json:"pri"`
	Relo                  string `json:"relo"`
	Init                  string `json:"init"`
	Unassign              string `json:"unassign"`
	Pending_tasks         string `json:"pending_tasks"`
	Max_task_wait_time    string `json:"max_task_wait_time"`
	Active_shards_percent string `json:"active_shards_percent"`
}

// The IndiceInfo struct is representing all the fields returned by the Elasticsearch API when
// asking for information about an index
type IndiceInfo struct {
	Health         string `json:"health"`
	Status         string `json:"status"`
	Index          string `json:"index"`
	Uuid           string `json:"uuid"`
	Pri            string `json:"pri"`
	Rep            string `json:"rep"`
	Docs_count     string `json:"docs.count"`
	Docs_deleted   string `json:"docs.deleted"`
	Store_size     string `json:"store.size"`
	Pri_store_size string `json:"pri.store.size"`
}

// CanConnectToElasticSearchCluster will simply check if the Elasticsearch cluster is reachable.
// If it can connect, it will simply pass, and if it can't it will thrown an error.
//
// CanConnectToElasticSearchCluster("127.0.0.1", "9200")
//
func CanConnectToElasticSearchCluster(hostname string, port string) {
	// Url to get the health of the Elasticsearch cluster
	healthURL := "http://" + hostname + ":" + port + "/_cat/health?v&pretty"

	_, err := HttpGetWithJson(healthURL)
	if err != nil {
		log.Fatalf("Failed to connect to the cluster: %v", err)
	} else {
		log.Printf("Successfuly reached the Elasticsearch cluster")
	}
}

// GetClusterHealth returns a ClusterInfo struct describing the health, status and
// other informations about the specified cluster.
//
// GetClusterHealth("127.0.0.1", "9200")
//
func GetClusterHealth(hostname string, port string) (ClusterInfo, error) {

	var clusterInfo ClusterInfo

	// Url to get the health of the Elasticsearch cluster
	healthURL := "http://" + hostname + ":" + port + "/_cat/health?v&pretty"

	responseHealth, err := HttpGetWithJson(healthURL)
	if err != nil {
		return clusterInfo, err
	} else {
		log.Printf("Successfuly retrieved the health status of the cluster")
	}

	defer responseHealth.Body.Close()

	body, err := ioutil.ReadAll(responseHealth.Body)
	if err != nil {
		return clusterInfo, err
	}

	var clusterInfoToBeFilled []ClusterInfo
	err = json.Unmarshal(body, &clusterInfoToBeFilled)
	if err != nil {
		return clusterInfo, err
	}

	// The result is encapsulated in an array. The cluster is always at index 0.
	clusterInfo = clusterInfoToBeFilled[0]

	return clusterInfo, nil
}

// GetHealthOfIndice returns a IndiceInfo struct describing the health, status and
// other informations about the specified index in the cluster.
//
// GetHealthOfIndice("127.0.0.1", "9200", "myindex")
//
func GetHealthOfIndice(hostname string, port string, indicename string) (IndiceInfo, error) {

	var indiceInfo IndiceInfo

	// Url to get the health status of the given index of the Elasticsearch cluster
	healthURL := "http://" + hostname + ":" + port + "/_cat/indices/" + indicename + "?v"

	responseHealth, err := HttpGetWithJson(healthURL)
	if err != nil {
		return indiceInfo, err
	}

	defer responseHealth.Body.Close()

	body, err := ioutil.ReadAll(responseHealth.Body)
	if err != nil {
		return indiceInfo, err
	}

	var indiceInfoToBeFilled []IndiceInfo
	err = json.Unmarshal(body, &indiceInfoToBeFilled)
	if err != nil {
		return indiceInfo, err
	}

	// The result is encapsulated in an array. The cluster is always at index 0.
	indiceInfo = indiceInfoToBeFilled[0]

	return indiceInfo, nil
}

// HttpGetWithJson is an helper function to send Http GET request and
// accept Json format as responses.
func HttpGetWithJson(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
