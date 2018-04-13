package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/utils"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var name string
	flag.StringVar(&name, "name", "", "Name of the service-instance")

	var servicename string
	flag.StringVar(&servicename, "service", "", "Name of the service offering")

	var serviceplan string
	flag.StringVar(&serviceplan, "plan", "", "Name of the service plan")

	var resourcegrp string
	flag.StringVar(&resourcegrp, "resource-group", "", "Name of the resource group")

	var location string
	flag.StringVar(&location, "location", "", "location or region to deploy")

	var alias string
	flag.StringVar(&alias, "alias", "", "Name of the service alias")

	var org string
	flag.StringVar(&org, "org", "", "Organization")

	var space string
	flag.StringVar(&space, "space", "", "Space")

	flag.Parse()

	fmt.Println(name, ":", servicename, ":", serviceplan, ":", location)

	if name == "" || servicename == "" || serviceplan == "" || location == "" || alias == "" || org == "" || space == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	catalogClient, err := catalog.New(sess)

	if err != nil {
		log.Fatal(err)
	}
	resCatalogAPI := catalogClient.ResourceCatalog()

	service, err := resCatalogAPI.FindByName(servicename, true)
	if err != nil {
		log.Fatal(err)
	}

	servicePlanID, err := resCatalogAPI.GetServicePlanID(service[0], serviceplan)
	if err != nil {
		log.Fatal(err)
	}

	if servicePlanID == "" {
		_, err := resCatalogAPI.GetServicePlan(serviceplan)
		if err != nil {
			log.Fatal(err)
		}
		servicePlanID = serviceplan
	}

	var crn crn.CRN
	deployments, err := resCatalogAPI.ListDeployments(servicePlanID)
	if err != nil {
		log.Fatal(err)
	}

	found := false
	var supportedLocations []string
	for _, deployment := range deployments {
		deploymentLocation := utils.GetLocationFromTargetCRN(deployment.Metadata.Deployment.TargetCrn.Resource)
		if deploymentLocation == location {
			crn = deployment.Metadata.Deployment.TargetCrn
			found = true
			break
		} else {
			supportedLocations = append(supportedLocations, deploymentLocation)
		}
	}

	if !found {
		log.Printf("The location %s you specified is not valid to the service plan %s. Valid location(s) are: %q", location, serviceplan, supportedLocations)
		os.Exit(1)

	}

	managementClient, err := management.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	var resourceGroupID string
	resGrpAPI := managementClient.ResourceGroup()

	if resourcegrp == "" {

		resourceGroupQuery := management.ResourceGroupQuery{
			Default: true,
		}

		grpList, err := resGrpAPI.List(&resourceGroupQuery)

		if err != nil {
			log.Fatal(err)
		}
		resourceGroupID = grpList[0].ID

	} else {
		grp, err := resGrpAPI.FindByName(nil, resourcegrp)
		if err != nil {
			log.Fatal(err)
		}
		resourceGroupID = grp[0].ID
	}

	controllerClient, err := controller.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	resServiceInstanceAPI := controllerClient.ResourceServiceInstance()

	var serviceInstancePayload = controller.CreateServiceInstanceRequest{
		Name:            name,
		ServicePlanID:   servicePlanID,
		ResourceGroupID: resourceGroupID,
		TargetCrn:       crn,
	}

	serviceInstance, err := resServiceInstanceAPI.CreateInstance(serviceInstancePayload)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resoure service Instance Details :", serviceInstance)

	serviceInstance, err = resServiceInstanceAPI.GetInstance(serviceInstance.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resoure service Instance Details :", serviceInstance)

	aliasrsp, err := resCatalogAPI.GetDeploymentAlias(serviceInstance.ServicePlanID, serviceInstance.RegionID, sess.Config.Region)

	if err != nil {
		log.Fatal(err)
	}
	if aliasrsp == nil {
		log.Fatalf("Cannot create alias for service instance: %s in region %s", serviceInstance.ID, sess.Config.Region)

	}

	sess, err = session.New()
	if err != nil {
		log.Fatal(err)
	}

	mccpClient, err := mccpv2.New(sess)

	orgAPI := mccpClient.Organizations()

	orgrsp, err := orgAPI.FindByName(org, "")
	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := mccpClient.Spaces()

	spacersp, err := spaceAPI.FindByNameInOrg(orgrsp.GUID, space, "")
	if err != nil {
		log.Fatal(err)
	}

	regionAPI := mccpClient.Regions()

	region, err := regionAPI.FindRegionByName(sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	crn = utils.GenerateSpaceCRN(*region, orgrsp.GUID, spacersp.GUID)

	serviceAliasAPI := controllerClient.ResourceServiceAlias()
	createServiceAlias := controller.CreateServiceAliasParams{
		Name:              alias,
		ServiceInstanceID: serviceInstance.ID,
		ScopeCRN:          crn,
	}
	aliasresp, err := serviceAliasAPI.CreateAlias(createServiceAlias)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aliasresp)

	err = serviceAliasAPI.DeleteAlias(aliasresp.ID)

	if err != nil {
		log.Fatal(err)
	}

	err = resServiceInstanceAPI.DeleteInstance(serviceInstance.ID, true)

	if err != nil {
		log.Fatal(err)
	}
}
