package main

import (
	"fmt"
	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	//v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"

	"github.com/IBM-Cloud/bluemix-go/trace"
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

	clusterClient, err := v2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.Clusters()

	var name string
	fmt.Print("Enter cluster name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("out=", name)
	err1 := clustersAPI.Delete(name, target)
	//out,err=

	fmt.Println("err=", err1)

}
