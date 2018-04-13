package main

import (
	"flag"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	var region string
	flag.StringVar(&region, "region", "us-south", "Bluemix region")

	flag.Parse()

	if org == "" || space == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New(&bluemix.Config{Region: region, Debug: true})
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
	log.Println(myorg.GUID, myspace.GUID)

	quotaAPI := client.SpaceQuotas()

	createRequest := mccpv2.SpaceQuotaCreateRequest{
		Name:                    "test2",
		OrgGUID:                 myorg.GUID,
		MemoryLimitInMB:         1024,
		InstanceMemoryLimitInMB: 1024,
		RoutesLimit:             50,
		ServicesLimit:           150,
		NonBasicServicesAllowed: false,
	}

	myquota, err := quotaAPI.Create(createRequest)
	if err != nil {
		log.Fatal(err)
	}

	spaceCreateRequest := mccpv2.SpaceCreateRequest{
		Name:           "test",
		OrgGUID:        myorg.GUID,
		SpaceQuotaGUID: myquota.Metadata.GUID,
	}
	newspace, err := spaceAPI.Create(spaceCreateRequest)
	if err != nil {
		log.Fatal(err)
	}

	newspace, err = spaceAPI.Get(newspace.Metadata.GUID)
	if err != nil {
		log.Fatal(err)
	}

	spaceUpdateRequest := mccpv2.SpaceUpdateRequest{
		Name: helpers.String("testupdate"),
	}
	newspace, err = spaceAPI.Update(newspace.Metadata.GUID, spaceUpdateRequest)
	if err != nil {
		log.Fatal(err)
	}

	err = spaceAPI.Delete(newspace.Metadata.GUID)
	if err != nil {
		log.Fatal(err)
	}

	err = quotaAPI.Delete(myquota.Metadata.GUID)
	if err != nil {
		log.Fatal(err)
	}
}
