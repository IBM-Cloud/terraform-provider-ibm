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
	// Delete tag. Tag must be detached from all resources prior to deletion.
	var tag string
	flag.StringVar(&tag, "tag", "", "Tag to delete, after it has been detached")
	flag.Parse()

	if tag == "" {
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

	tagUpdateResult, err := gtClient.Tags().DeleteTag(tag)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result: ", tagUpdateResult)

}
