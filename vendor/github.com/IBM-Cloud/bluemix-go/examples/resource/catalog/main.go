package main

import (
	"flag"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var service string
	flag.StringVar(&service, "service", "", "Name of the service offering")

	flag.Parse()

	if service == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New(&bluemix.Config{Debug: true})
	if err != nil {
		log.Fatal(err)
	}

	catalogClient, err := catalog.New(sess)

	if err != nil {
		log.Fatal(err)
	}
	resCatalogAPI := catalogClient.ResourceCatalog()

	services, err := resCatalogAPI.GetServices()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(services)

	serviceRes, err := resCatalogAPI.FindByName(service, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(serviceRes)

	plans, err := resCatalogAPI.GetServicePlans(serviceRes[0])
	if err != nil {
		log.Fatal(err)
	}

	log.Println(plans)

}
