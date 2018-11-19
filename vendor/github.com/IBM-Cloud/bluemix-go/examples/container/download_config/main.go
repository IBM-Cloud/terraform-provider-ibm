package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var clusterName string
	flag.StringVar(&clusterName, "clustername", "", "The cluster whose config will be downloaded")

	var path string
	flag.StringVar(&path, "path", "", "The Path where the config will be downloaded")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	var region string
	flag.StringVar(&region, "region", "us-south", "Bluemix region")

	var admin bool
	flag.BoolVar(&admin, "admin", false, "If true download the admin config")

	var network bool
	flag.BoolVar(&network, "network", false, "If true download the calico network config")

	flag.Parse()
	trace.Logger = trace.NewLogger("true")
	if org == "" || space == "" || clusterName == "" || path == "" {
		flag.Usage()
		os.Exit(1)
	}

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mccpv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, region)

	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, region)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, region)
	if err != nil {
		log.Fatal(err)
	}

	target := v1.ClusterTargetHeader{
		OrgID:     myorg.GUID,
		SpaceID:   myspace.GUID,
		AccountID: myAccount.GUID,
		Region:    region,
	}

	clusterClient, err := v1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.Clusters()
	var configPath, kubeConfig string
	if network {
		kubeConfig, configPath, err = clustersAPI.StoreConfig(clusterName, path, admin, network, target)
	} else {
		configPath, err = clustersAPI.GetClusterConfig(clusterName, path, admin, target)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kubeConfig)
	fmt.Println(configPath)
}
