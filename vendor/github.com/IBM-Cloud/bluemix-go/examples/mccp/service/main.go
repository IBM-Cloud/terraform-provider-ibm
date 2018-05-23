package main

import (
	"flag"
	"log"
	"os"

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

	var skipDeletion bool
	flag.BoolVar(&skipDeletion, "no-delete", true, "If provided will delete the resources created")

	flag.Parse()

	if org == "" || space == "" {
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

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, region)

	if err != nil {
		log.Fatal(err)
	}

	serviceOfferingAPI := client.ServiceOfferings()
	myserviceOff, err := serviceOfferingAPI.FindByLabel("cloud-object-storage")
	if err != nil {
		log.Fatal(err)
	}
	servicePlanAPI := client.ServicePlans()
	plan, err := servicePlanAPI.FindPlanInServiceOffering(myserviceOff.GUID, "Lite")
	if err != nil {
		log.Fatal(err)
	}

	serviceInstanceAPI := client.ServiceInstances()
	myService, err := serviceInstanceAPI.Create(mccpv2.ServiceInstanceCreateRequest{
		Name:      "myservice",
		PlanGUID:  plan.GUID,
		SpaceGUID: myspace.GUID,
	})
	if err != nil {
		log.Fatal(err)
	}

	updatedInstance, err := serviceInstanceAPI.Update(myService.Metadata.GUID, mccpv2.ServiceInstanceUpdateRequest{
		Name: helpers.String("New instance"),
	})
	if err != nil {
		log.Fatal(err)
	}

	serviceKeys := client.ServiceKeys()
	mykeys, err := serviceKeys.Create(updatedInstance.Metadata.GUID, "mykey", nil)
	if err != nil {
		log.Fatal(err)
	}

	myRetrievedKeys, err := serviceKeys.FindByName(myService.Metadata.GUID, "mykey")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(myorg.GUID, myspace.GUID, plan.GUID, myService.Metadata.GUID, mykeys.Metadata.GUID, myRetrievedKeys)

	if !skipDeletion {
		err = serviceKeys.Delete(myRetrievedKeys.GUID)
		if err != nil {
			log.Fatal(err)
		}

		err = serviceInstanceAPI.Delete(myService.Metadata.GUID)
		if err != nil {
			log.Fatal(err)
		}

	}

}
