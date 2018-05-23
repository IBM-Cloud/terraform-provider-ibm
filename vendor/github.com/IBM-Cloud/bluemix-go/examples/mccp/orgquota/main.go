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

	var orgquota string
	flag.StringVar(&orgquota, "orgquota", "", "Bluemix Org Quota Definition")

	flag.Parse()

	if orgquota == "" {
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

	orgQuotaAPI := client.OrgQuotas()

	quotas, err := orgQuotaAPI.List()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(quotas)

	quota, err := orgQuotaAPI.FindByName(orgquota)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Quota Defination Details :", quota)

}
