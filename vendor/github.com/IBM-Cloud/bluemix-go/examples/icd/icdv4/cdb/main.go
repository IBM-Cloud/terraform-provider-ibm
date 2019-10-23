package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var icdId string
	flag.StringVar(&icdId, "icdId", "", "CRN of the IBM Cloud Database service instance")
	flag.Parse()

	if icdId == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	icdClient, err := icdv4.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	cdbAPI := icdClient.Cdbs()
	cdb, err := cdbAPI.GetCdb(icdId)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Cdb :", cdb)

}
