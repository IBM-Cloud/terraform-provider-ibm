package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var resourcegrp string
	flag.StringVar(&resourcegrp, "name", "", "Name of the group")

	flag.Parse()

	if resourcegrp == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	client, err := management.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	resGrpAPI := client.ResourceGroup()

	resourceGroupQuery := management.ResourceGroupQuery{
		Default: true,
	}

	grpList, err := resGrpAPI.List(&resourceGroupQuery)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group default Details :", grpList)

	var grpInfo = models.ResourceGroup{
		Name: resourcegrp,
	}

	grp, err := resGrpAPI.Create(grpInfo)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group create :", grp)

	grps, err := resGrpAPI.FindByName(nil, resourcegrp)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group Details :", grps[0])

	grp, err = resGrpAPI.Get(grp.ID)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group Details by ID:", grp)

	var updateGrpInfo = management.ResourceGroupUpdateRequest{
		Name: "default",
	}

	grp, err = resGrpAPI.Update(grp.ID, &updateGrpInfo)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group update :", grp)

	err = resGrpAPI.Delete(grp.ID)
	if err != nil {
		log.Fatal(err)
	}

}
