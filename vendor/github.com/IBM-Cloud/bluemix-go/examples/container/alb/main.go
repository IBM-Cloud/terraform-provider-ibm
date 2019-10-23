package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM-Cloud/bluemix-go/session"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var albID string
	flag.StringVar(&albID, "albID", "", "Alb Id")

	var enable bool
	flag.BoolVar(&enable, "enable", false, "enable alb")

	var clusterID string
	flag.StringVar(&clusterID, "clusterNameOrID", "", "cluster name or ID")

	var region string
	flag.StringVar(&region, "region", "us-south", "region of cluster")

	flag.Parse()

	trace.Logger = trace.NewLogger("true")
	if albID == "" || clusterID == "" {
		flag.Usage()
		os.Exit(1)
	}

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	albClient, err := v1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	albAPI := albClient.Albs()
	target := v1.ClusterTargetHeader{
		Region: region,
	}
	//List All Albs
	albs, err := albAPI.ListClusterALBs(clusterID, target)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listing the albs bounded to cluster ", clusterID)

	for _, alb := range albs {
		fmt.Println(alb.ALBID, alb.ALBIP, alb.State)
	}

	//Enable the Alb
	var albConfig = v1.ALBConfig{
		ALBID:     albID,
		Enable:    enable,
		ClusterID: clusterID,
	}
	err = albAPI.ConfigureALB(albID, albConfig, false, target)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	alb, err := albAPI.GetALB(albID, target)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The status of alb is ", alb.State)

}
