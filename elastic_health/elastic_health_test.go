package main

import ( 
	"testing"
	"os"
	"log"

)

func TestCorrectArguments(t *testing.T) {
	log.Printf("- TestCorrectArguments- ")

	os.Args = []string{"./elastic_health", "localhost", "9200"}

	clusterHostname, clusterPort, err := GetClusterHostnameAndPortFromArguments()

	if(err != nil || clusterHostname != "localhost" || clusterPort != "9200") {
		t.Errorf("Error arguments.")
	}


}

func TestCorrectArgumentsWithIP(t *testing.T) {
	log.Printf("- TestCorrectArgumentsWithIP- ")


	os.Args = []string{"./elastic_health", "9.3.12.128", "9200"}

	clusterHostname, clusterPort, err := GetClusterHostnameAndPortFromArguments()

	if(err != nil || clusterHostname != "9.3.12.128" || clusterPort != "9200") {
		t.Errorf("Error arguments.")
	}


}

func TestArgumentsWithWrongPort(t *testing.T) {
	log.Printf("- TestArgumentsWithWrongPort- ")


	os.Args = []string{"./elastic_health", "9.3.12.128", "localhost"}

	_, _, err := GetClusterHostnameAndPortFromArguments()

	if( err == nil) {
		t.Errorf("Error expected but the arguments check did not throw any error for having a port value not equal to a number.")
	}


}

func TestMissingArguments(t *testing.T) {
	log.Printf("- TestMissingArguments- ")

	os.Args = []string{"./elastic_health", "localhost"}

	_, _, err := GetClusterHostnameAndPortFromArguments()

	if( err == nil) {
		t.Errorf("Error expected but the arguments check did not throw any error.")
	}

}

func TestNoArguments(t *testing.T) {
	log.Printf("- TestNoArguments- ")

	os.Args = []string{"./elastic_health"}

	_, _, err := GetClusterHostnameAndPortFromArguments()

	if( err == nil) {
		t.Errorf("Error expected but the arguments check did not throw any error.")
	}

}

func TestMoreArgumentsThanRequired(t *testing.T) {
	log.Printf("- TestMoreArgumentsThanRequired- ")

	os.Args = []string{"./elastic_health", "localhost", "9200", "xyz"}

	_, _, err := GetClusterHostnameAndPortFromArguments()

	if( err == nil) {
		t.Errorf("Error expected but the arguments check did not throw any error.")
	}

}