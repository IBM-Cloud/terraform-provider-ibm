package main

import (
	"log"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mccpv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	regionAPI := client.Regions()

	regions, err := regionAPI.Regions()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Region Details :", regions)

	regions, err = regionAPI.PublicRegions()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Region Details :", regions)

}
