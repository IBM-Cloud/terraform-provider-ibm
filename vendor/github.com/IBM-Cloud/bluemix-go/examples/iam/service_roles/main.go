package main

import (
	"flag"
	"log"

	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var service string
	flag.StringVar(&service, "service", "", "Name of the service offering")

	flag.Parse()
	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	serviceRolesAPI := iamClient.ServiceRoles()
	var roles []models.PolicyRole

	if service == "" {
		roles, err = serviceRolesAPI.ListSystemDefinedRoles()
		if err != nil {
			log.Fatal(err)
		}

	} else {

		catalogClient, err := catalog.New(sess)

		if err != nil {
			log.Fatal(err)
		}
		resCatalogAPI := catalogClient.ResourceCatalog()

		service, err := resCatalogAPI.FindByName(service, true)
		if err != nil {
			log.Fatal(err)
		}
		roles, err = serviceRolesAPI.ListServiceRoles(service[0].Name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(roles)

	}

}
