package main

import (
	"fmt"
	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	c := new(bluemix.Config)

	trace.Logger = trace.NewLogger("true")

	target := v2.ClusterTargetHeader{}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	clusterClient, err := v2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.Clusters()

	out, err := clustersAPI.GetCluster("bm64u3ed02o93vv36hb0", target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ouyt=", out)
}
