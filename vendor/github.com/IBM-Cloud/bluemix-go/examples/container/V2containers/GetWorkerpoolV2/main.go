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

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	target := v2.ClusterTargetHeader{}

	var cluster_id = "bm64u3ed02o93vv36hb0"
	var workerpool_id = "bm64u3ed02o93vv36hb0-0dc20a0"

	clusterClient, err := v2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	workerpoolAPI := clusterClient.WorkerPools()

	out, err := workerpoolAPI.GetWorkerPool(cluster_id, workerpool_id, target)

	fmt.Println("out=", out)
}
