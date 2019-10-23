package main

import (
	"fmt"
	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func main() {

	c := new(bluemix.Config)

	trace.Logger = trace.NewLogger("true")

	var workerpoolid = "bm64u3ed02o93vv36hb0-a627b81"
	var clusterId = "bm64u3ed02o93vv36hb0"
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

	WorkerAPI := clusterClient.Workers()

	workerInfo, err2 := WorkerAPI.ListByWorkerPool(clusterId, workerpoolid, true, target)

	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println("workerout=", workerInfo)
}
