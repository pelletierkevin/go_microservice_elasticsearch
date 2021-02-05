// Package elasticsearch implements methods to call Elasticsearch API, retrieve and cast results in proper return values.
// The methods will essentially deliver informations about the health, status and the indices of an Elasticsearch cluster.
package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetClusterIndices returns an array of IndiceInfo which are currently available in the given
// Elasticsearch cluster.
//
// GetClusterIndices("127.0.0.1", "9200")
//
func GetClusterIndices(hostname string, port string) ([]IndiceInfo, error) {

	var listIndices []IndiceInfo

	// Url to get all the indices of the Elasticsearch cluster
	healthURL := "http://" + hostname + ":" + port + "/_cat/indices?v"

	responseHealth, err := HttpGetWithJson(healthURL)
	if err != nil {
		return listIndices, err
	} else {
		log.Printf("Successfuly retrieved the indices of the cluster")
	}

	defer responseHealth.Body.Close()

	body, err := ioutil.ReadAll(responseHealth.Body)
	if err != nil {
		return listIndices, err
	}

	err = json.Unmarshal(body, &listIndices)
	if err != nil {
		return listIndices, err
	}

	return listIndices, nil
}

// CreateIndexInCluster will create an index in the specified elasticsearch cluster. For the moment
// the index will be created using default values. It returns a string which represent the response of
// the request to the Elasticsearch API.
//
// CreateIndexInCluster("127.0.0.1", "9200", "mynewindex")
//
func CreateIndexInCluster(hostname string, port string, indicename string) (string, error) {

	// Url to get all the indices of the Elasticsearch cluster
	indexURL := "http://" + hostname + ":" + port + "/" + indicename

	client := &http.Client{}
	// TODO Here we can change the 3rd parameters set to nil, to some data to specify the number of replicas and others field.
	req, err := http.NewRequest("PUT", indexURL, nil)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	log.Printf("Delete index request response : %s", string(body))

	return string(body), err

}

// DeleteIndexInCluster will delete an index in the specified elasticsearch cluster.
// It returns a string which represent the response of the request to the Elasticsearch API.
//
// DeleteIndexInCluster("127.0.0.1", "9200", "mynewindex")
//
func DeleteIndexInCluster(hostname string, port string, indicename string) (string, error) {

	// Url to get all the indices of the Elasticsearch cluster
	indexURL := "http://" + hostname + ":" + port + "/" + indicename

	client := &http.Client{}
	// TODO Here we can change the 3rd parameters set to nil, to some data to specify the number of replicas and others field.
	req, err := http.NewRequest("DELETE", indexURL, nil)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	log.Printf("Delete index request response : %s", string(body))

	return string(body), err

}
