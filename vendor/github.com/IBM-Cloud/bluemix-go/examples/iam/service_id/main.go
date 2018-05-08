package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/utils"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var serviceID string
	flag.StringVar(&serviceID, "serviceID", "", "Bluemix service id name")

	flag.Parse()
	if org == "" || serviceID == "" {
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
	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, sess.Config.Region)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	regionAPI := client.Regions()
	region, err := regionAPI.FindRegionByName(sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	serviceIdAPI := iamClient.ServiceIds()

	boundTo := utils.GenerateBoundToCRN(*region, myAccount.GUID).String()

	data := models.ServiceID{
		Name:    serviceID,
		BoundTo: boundTo,
	}
	sID, err := serviceIdAPI.Create(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sID)

	sID, err = serviceIdAPI.Get(sID.UUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sID)

	err = serviceIdAPI.Delete(sID.UUID)
	if err != nil {
		log.Fatal(err)
	}

}
