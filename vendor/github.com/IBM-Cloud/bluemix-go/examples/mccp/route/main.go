package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	var host string
	flag.StringVar(&host, "host", "myexample", "Bluemix Host")

	var path string
	flag.StringVar(&path, "path", "/mypath", "Bluemix Path")

	var newHost string
	flag.StringVar(&newHost, "new_host", "myexample1", "Bluemix Path")

	flag.Parse()

	if org == "" || space == "" {
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

	region := sess.Config.Region
	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, region)
	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, region)
	if err != nil {
		log.Fatal(err)
	}

	sharedDomainAPI := client.SharedDomains()

	sd, err := sharedDomainAPI.FindByName("mybluemix.net")
	if err != nil {
		log.Fatal(err)
	}

	routesAPI := client.Routes()

	payload := mccpv2.RouteRequest{
		Host:       host,
		SpaceGUID:  myspace.GUID,
		DomainGUID: sd.GUID,
		Path:       path,
	}
	r, err := routesAPI.Create(payload)
	if err != nil {
		log.Fatal(err)
	}

	updatePayload := mccpv2.RouteUpdateRequest{
		Host: helpers.String(newHost),
		Path: helpers.String(""),
	}

	updatedRoute, err := routesAPI.Update(r.Metadata.GUID, updatePayload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*updatedRoute)

	routeFilter := mccpv2.RouteFilter{
		DomainGUID: sd.GUID,
		Host:       helpers.String(newHost),
	}

	routes, err := spaceAPI.ListRoutes(myspace.GUID, routeFilter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(routes)

	err = routesAPI.Delete(r.Metadata.GUID, true)
	if err != nil {
		log.Fatal(err)
	}

}
