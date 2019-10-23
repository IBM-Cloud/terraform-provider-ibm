package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var resourceID string
	flag.StringVar(&resourceID, "id", "", "CRN string")
	flag.Parse()

	if resourceID == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	gtClient, err := globaltaggingv3.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	taggingResult, err := gtClient.Tags().GetTags(resourceID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result: ", taggingResult)

	log.Println("Items: ", taggingResult.Items)

}
