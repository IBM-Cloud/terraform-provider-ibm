package main

import (
	"flag"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	iam "github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
	"github.com/IBM-Cloud/bluemix-go/utils"
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

	var key string
	flag.StringVar(&key, "key", "", "Name of the service key")

	var role string
	flag.StringVar(&role, "role", "", "Name of the role")

	flag.Parse()

	if name == "" || servicename == "" || serviceplan == "" || location == "" {
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

	deployments, err := resCatalogAPI.ListDeployments(servicePlanID)
	if err != nil {
		log.Fatal(err)
	}

	if len(deployments) == 0 {

		log.Printf("No deployment found for service plan : %s", serviceplan)
		os.Exit(1)
	}

	supportedDeployments := []models.ServiceDeployment{}
	supportedLocations := make(map[string]bool)
	for _, d := range deployments {
		if d.Metadata.RCCompatible {
			deploymentLocation := d.Metadata.Deployment.Location
			supportedLocations[deploymentLocation] = true
			if deploymentLocation == location {
				supportedDeployments = append(supportedDeployments, d)
			}
		}
	}

	if len(supportedDeployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		log.Printf("No deployment found for service plan %s at location %s.\nValid location(s) are: %q.\nUse service instance example if the service is a Cloud Foundry service.", serviceplan, location, locationList)
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
		TargetCrn:       supportedDeployments[0].CatalogCRN,
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

	instance, err := resServiceInstanceAPI.GetInstance(serviceInstance.ID)
	if err != nil {
		log.Fatal(err)
	}

	serviceresp, err := resCatalogAPI.Get(instance.ServiceID, true)
	if err != nil {
		log.Fatal(err)
	}

	iamClient, err := iam.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	serviceRolesAPI := iamClient.ServiceRoles()

	var roles []models.PolicyRole

	if serviceresp.Name == "" {
		roles, err = serviceRolesAPI.ListSystemDefinedRoles()

	} else {
		roles, err = serviceRolesAPI.ListServiceRoles(serviceresp.Name)
	}

	if err != nil {
		log.Fatal(err)
	}

	roleMatch, err := utils.FindRoleByName(roles, role)
	if err != nil {
		log.Fatal(err)
	}

	parameters := make(map[string]interface{})

	parameters["role_crn"] = roleMatch.ID

	resServiceKeyAPI := controllerClient.ResourceServiceKey()

	params := controller.CreateServiceKeyRequest{
		Name:       key,
		SourceCRN:  instance.Crn,
		Parameters: parameters,
	}

	keyresp, err := resServiceKeyAPI.CreateKey(params)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resoure service key Details :", keyresp)

	keyresp, err = resServiceKeyAPI.GetKey(keyresp.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resoure service key Details :", keyresp)

	err = resServiceKeyAPI.DeleteKey(keyresp.ID)

	if err != nil {
		log.Fatal(err)
	}

	err = resServiceInstanceAPI.DeleteInstance(serviceInstance.ID, true)

	if err != nil {
		log.Fatal(err)
	}

}
