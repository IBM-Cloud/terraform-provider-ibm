package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var resourceID string
	var tagin string
	flag.StringVar(&resourceID, "id", "", "CRN string")
	flag.StringVar(&tagin, "tags", "", "List of comma separated tags")
	flag.Parse()

	if resourceID == "" || tagin == "" {
		flag.Usage()
		os.Exit(1)
	}

	taglist := strings.Split(tagin, ",")

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	gtClient, err := globaltaggingv3.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	tagUpdateResult, err := gtClient.Tags().AttachTags(resourceID, taglist)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result: ", tagUpdateResult)

}
