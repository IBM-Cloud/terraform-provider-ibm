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

	var groupType string
	flag.StringVar(&groupType, "groupType", "", "Type of ICD database for defaultGroup")
	flag.Parse()

	if groupType == "" {
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
	groupAPI := icdClient.Groups()

	defaultGroup, err := groupAPI.GetDefaultGroups(groupType)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DefaultGroup :", defaultGroup)

}
