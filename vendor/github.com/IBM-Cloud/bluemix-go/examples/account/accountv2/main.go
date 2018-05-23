package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	var ownerUserID string
	flag.StringVar(&org, "org", "", "Bluemix Organization")
	flag.StringVar(&ownerUserID, "owner_id", "", "Owner user id, for example - abc@c.com")

	flag.Parse()

	if org == "" || ownerUserID == "" {
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
	myorg, err := orgAPI.FindByName(org, "us-south")

	if err != nil {
		log.Fatal(err)
	}
	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, "us-south")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(myAccount.Name, myAccount.CountryCode, myAccount.OwnerUserID, myAccount.GUID)

	myAccount, err = accountAPI.Get(myAccount.GUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(myAccount.Name, myAccount.CountryCode, myAccount.OwnerUserID, myAccount.GUID)

	myAccount, err = accountAPI.FindByOwner(ownerUserID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(myAccount.Name, myAccount.CountryCode, myAccount.OwnerUserID, myAccount.GUID)

	allAccounts, err := accountAPI.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, acc := range allAccounts {
		fmt.Println(acc.OwnerUserID, acc.GUID)
	}
}
