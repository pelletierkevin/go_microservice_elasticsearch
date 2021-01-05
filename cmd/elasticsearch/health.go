package elasticsearch

import (
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


type ClusterInfo struct {
	Epoch string `json:"epoch"`
	Timestamp string `json:"timestamp"`
	Name string `json:"cluster"`
	Status string `json:"status"`
	Node_total string `json:"node.total"`
	Node_data string `json:"node.data"`
	Shards string `json:"shards"`
	Pri string `json:"pri"`
	Relo string `json:"relo"`
	Init string `json:"init"`
	Unassign string `json:"unassign"`
	Pending_tasks string `json:"pending_tasks"`
	Max_task_wait_time string `json:"max_task_wait_time"`
	Active_shards_percent string `json:"active_shards_percent"`

}

type IndiceInfo struct {
	Health string `json:"health"`
	Status string `json:"status"`
	Index string `json:"index"`
	Uuid string `json:"uuid"`
	Pri string `json:"pri"`
	Rep string `json:"rep"`
	Docs_count string `json:"docs.count"`
	Docs_deleted string `json:"docs.deleted"`
	Store_size string `json:"store.size"`
	Pri_store_size string `json:"pri.store.size"`
}


func GetClusterHealth(hostname string, port string) (ClusterInfo, error) {
	// If no name was given, return an error with a message.
	var clusterInfo ClusterInfo

    if (hostname == "" || port ==  "") {
        return clusterInfo, errors.New("Hostname or Port not defined")
    }
    // Create a message using a random format.
	healthURL := "http://" + hostname + ":" + port + "/_cat/health?v&pretty"

	responseHealth, err := HttpGetWithJson(healthURL)
	if err != nil {
        return clusterInfo, err
	}

	defer responseHealth.Body.Close()

	body, err := ioutil.ReadAll(responseHealth.Body)
	if err != nil {
        return clusterInfo, err
	}

	fmt.Println(string(body))

	var generic []ClusterInfo
	err = json.Unmarshal(body, &generic)
	if err != nil {
		return clusterInfo, err
	}

	clusterInfo = generic[0]

	fmt.Println("Struct :")
	fmt.Println(clusterInfo)

	fmt.Println("Get STATUS of Cluster : " + string(clusterInfo.Status) )

	// fmt.Println("Get STATUS of Cluster : " + string(jsonStruct) )

	// return string(body), nil
	return clusterInfo, nil
}

func GetClusterIndices(hostname string, port string) ([]IndiceInfo, error) {
	
	// Proto requires to use an array of pointers
	var listIndices []IndiceInfo
	
	// If no name was given, return an error with a message.
    if (hostname == "" || port ==  "") {
        return listIndices, errors.New("Hostname or Port not defined")
    }
    // Create a message using a random format.
	healthURL := "http://" + hostname + ":" + port + "/_cat/indices?v"

	responseHealth, err := HttpGetWithJson(healthURL)
	if err != nil {
        return listIndices, err
	}

	defer responseHealth.Body.Close()

	body, err := ioutil.ReadAll(responseHealth.Body)
	if err != nil {
        return listIndices, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &listIndices)
	if err != nil {
		return listIndices, err
	}

	fmt.Println("All indices :")
	fmt.Println(listIndices)
	

	return listIndices, nil
}

func GetHealthOfIndice(hostname string, port string, indicename string) (IndiceInfo, error) {
   
	var indiceInfo IndiceInfo
	
	// If no name was given, return an error with a message.
    if (hostname == "" || port ==  "") {
        return indiceInfo, errors.New("Hostname or Port not defined")
    }
    // Create a message using a random format.
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


	fmt.Println(string(body))

	var generic []IndiceInfo
	err = json.Unmarshal(body, &generic)
	if err != nil {
		return indiceInfo, err
	}

	indiceInfo = generic[0]

	fmt.Println("Struct :")
	fmt.Println(indiceInfo)

	fmt.Println("Get STATUS of Cluster : " + string(indiceInfo.Status) )
	

	return indiceInfo, nil
}

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



