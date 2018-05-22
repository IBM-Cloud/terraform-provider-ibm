package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
)

//Example to list all supported versions of the kube in the IBM Container Service
func main() {
	c := new(bluemix.Config)
	flag.StringVar(&c.IBMID, "ibmid", "", "The IBM ID. You can also source it from env IBMID")
	flag.StringVar(&c.IBMIDPassword, "ibmidpass", "", "The IBMID Password. You can also source it from IBMID_PASSWORD")
	flag.StringVar(&c.Region, "region", "us-south", "The Bluemix region. You can source it from env BM_REGION or BLUEMIX_REGION")
	flag.BoolVar(&c.Debug, "debug", false, "Show full trace if on")

	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	flag.Parse()

	if org == "" || space == "" {
		flag.Usage()
		os.Exit(1)
	}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	client, err := mccpv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	region := sess.Config.Region
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
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, c.Region)
	if err != nil {
		log.Fatal(err)
	}

	target := v1.ClusterTargetHeader{
		OrgID:     myorg.GUID,
		SpaceID:   myspace.GUID,
		AccountID: myAccount.GUID,
	}

	clusterClient, err := v1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.KubeVersions()

	out, err := clustersAPI.List(target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
