package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var spacequota string
	flag.StringVar(&spacequota, "spacequota", "", "Bluemix Space Quota Definition")

	var newspacequota string
	flag.StringVar(&newspacequota, "newspacequota", "", "Bluemix Space Quota Definition")

	flag.Parse()

	if org == "" || spacequota == "" || newspacequota == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
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

	createRequest := mccpv2.SpaceQuotaCreateRequest{
		Name:                    spacequota,
		OrgGUID:                 myorg.GUID,
		MemoryLimitInMB:         1024,
		InstanceMemoryLimitInMB: 1024,
		RoutesLimit:             50,
		ServicesLimit:           150,
		NonBasicServicesAllowed: false,
	}

	spaceQuotaAPI := client.SpaceQuotas()
	_, err = spaceQuotaAPI.Create(createRequest)

	if err != nil {
		log.Fatal(err)
	}

	quota, err := spaceQuotaAPI.FindByName(spacequota, myorg.GUID)

	if err != nil {
		log.Fatal(err)
	}

	updateRequest := mccpv2.SpaceQuotaUpdateRequest{
		Name: newspacequota,
	}

	_, err = spaceQuotaAPI.Update(updateRequest, quota.GUID)

	if err != nil {
		log.Fatal(err)
	}

	quota, err = spaceQuotaAPI.FindByName(newspacequota, myorg.GUID)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(quota.GUID, myorg.GUID)

	err = spaceQuotaAPI.Delete(quota.GUID)

	if err != nil {
		log.Fatal(err)
	}
}
