package main

import (
	"log"

	"github.com/IBM-Cloud/bluemix-go/api/cse/csev2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	trace.Logger = trace.NewLogger("false")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	cseClient, err := csev2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	seAPI := cseClient.ServiceEndpoints()

	payload := csev2.SeCreateData{
		ServiceName:      "test-terrafor-11",
		CustomerName:     "test-customer-11",
		ServiceAddresses: []string{"10.102.33.131", "10.102.33.133"},
		Region:           "us-south",
		DataCenters:      []string{"dal10"},
		TCPPorts:         []int{8080, 80},
	}

	// create a serviceendpoint
	log.Println("create a serviceendpoint with ", payload)
	newSrvId, err := seAPI.CreateServiceEndpoint(payload)
	if err != nil {
		log.Fatal(err)
	}

	// query the serviceendpoint
	log.Println("query the serviceendpoint ", newSrvId)
	srvObj, err := seAPI.GetServiceEndpoint(newSrvId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Srvid=", srvObj.Service.Srvid)

	// delete serviceendpoint
	log.Println("delete the service endpoint ", newSrvId)
	err = seAPI.DeleteServiceEndpoint(newSrvId)
	if err != nil {
		log.Fatal(err)
	}
}
