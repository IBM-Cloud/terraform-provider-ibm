package main

import (
	"flag"
	"log"
	"os"

	"github.com/softlayer/softlayer-go/sl"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var neworg string
	flag.StringVar(&neworg, "neworg", "", "Bluemix Organization")

	flag.Parse()

	if org == "" || neworg == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	region := sess.Config.Region
	client, err := mccpv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgAPI := client.Organizations()

	payload := mccpv2.OrgCreateRequest{

		Name: org,
	}

	orgDetails, err := orgAPI.Create(payload)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Org Details:", orgDetails)

	myorg, err := orgAPI.FindByName(org, region)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(myorg.GUID, myorg.Name)

	updatedPayload := mccpv2.OrgUpdateRequest{
		Name: sl.String(neworg),
	}

	updatedOrgDetails, err := orgAPI.Update(myorg.GUID, updatedPayload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Org Details after update:", updatedOrgDetails)

	updatedOrg, err := orgAPI.FindByName(neworg, region)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(updatedOrg.GUID, updatedOrg.Name)

	getOrgByGUID, err := orgAPI.Get(updatedOrg.GUID)
	if err != nil {
		log.Fatal(err)
	}

	if updatedOrg.Name != getOrgByGUID.Entity.Name {
		log.Fatalf("Org obtained from FindByName and Get doesn't  match %s != %s", updatedOrg.GUID, getOrgByGUID.Metadata.GUID)
	}

	err = orgAPI.DeleteByRegion(updatedOrg.GUID, region, true)
	if err != nil {
		log.Fatal(err)
	}

	_, err = orgAPI.List(region)
	if err != nil {
		log.Fatal(err)
	}
}
