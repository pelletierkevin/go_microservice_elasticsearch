package elasticsearch

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

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