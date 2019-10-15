package main

import (
	"fmt"
	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"

	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	c := new(bluemix.Config)

	trace.Logger = trace.NewLogger("true")

	var clusterInfo = v2.ClusterCreateRequest{
		DisablePublicServiceEndpoint: true,
		KubeVersion:                  "1.15.3_1515",
		Name:                         "mycluscretaed12",
		PodSubnet:                    "172.30.0.0/16",
		Provider:                     "vpc-classic",
		ServiceSubnet:                "172.21.0.0/16",
		WorkerPools: v2.WorkerPoolConfig{
			DiskEncryption: true,
			Flavor:         "c2.2x4",
			Isolation:      "dedicated",
			Name:           "mywork1",
			VpcID:          "6015365a-9d93-4bb4-8248-79ae0db2dc26",
			WorkerCount:    1,
			Zones: []v2.Zone{
				{
					ID:       "us-south-1",
					SubnetID: "015ffb8b-efb1-4c03-8757-29335a07493b",
				},
			},
		},
	}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	target := v2.ClusterTargetHeader{}

	clusterClient, err := v2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.Clusters()

	out, err := clustersAPI.Create(clusterInfo, target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ouyt=", out)
}
