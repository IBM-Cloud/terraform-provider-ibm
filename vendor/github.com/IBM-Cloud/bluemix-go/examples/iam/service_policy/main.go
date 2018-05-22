package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/utils"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var serviceID string
	flag.StringVar(&serviceID, "serviceID", "", "Bluemix service id name")

	var service string
	flag.StringVar(&service, "service", "", "Bluemix service name")

	var roles string
	flag.StringVar(&roles, "roles", "", "Comma seperated list of roles")

	var serviceInstance string
	flag.StringVar(&serviceInstance, "serviceInstance", "", "Bluemix service instance name")

	var region string
	flag.StringVar(&region, "region", "", "Bluemix region")

	var resourceType string
	flag.StringVar(&resourceType, "resourceType", "", "Bluemix resource type")

	var resource string
	flag.StringVar(&resource, "resource", "", "Bluemix resource")

	var resourceGroupID string
	flag.StringVar(&resourceGroupID, "resourceGroupID", "", "Bluemix resource group ")

	flag.Parse()
	if org == "" || serviceID == "" || roles == "" {
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
	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, sess.Config.Region)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	regionAPI := client.Regions()
	regionList, err := regionAPI.FindRegionByName(sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	serviceIDAPI := iamClient.ServiceIds()

	serviceRolesAPI := iamClient.ServiceRoles()

	boundTo := utils.GenerateBoundToCRN(*regionList, myAccount.GUID).String()

	data := models.ServiceID{
		Name:    serviceID,
		BoundTo: boundTo,
	}
	sID, err := serviceIDAPI.Create(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sID)

	sID, err = serviceIDAPI.Get(sID.UUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sID)

	policyBoundTo := utils.GenerateBoundToCRN(*regionList, myAccount.GUID)

	var policy models.Policy

	var definedRoles []models.PolicyRole

	if service == "" {
		definedRoles, err = serviceRolesAPI.ListSystemDefinedRoles()
	} else {
		definedRoles, err = serviceRolesAPI.ListServiceRoles(service)
	}

	if err != nil {
		log.Fatal(err)
	}

	filterRoles, err := utils.GetRolesFromRoleNames(strings.Split(roles, ","), definedRoles)

	if err != nil {
		log.Fatal(err)
	}

	policyResource := models.PolicyResource{}

	if service != "" {
		policyResource.ServiceName = service
	}

	if serviceInstance != "" {
		policyResource.ServiceInstance = serviceInstance
	}

	if region != "" {
		policyResource.Region = region
	}

	if resourceType != "" {
		policyResource.ResourceType = resourceType
	}

	if resource != "" {
		policyResource.Resource = resource
	}

	if resourceGroupID != "" {
		policyResource.ResourceGroupID = resourceGroupID
	}

	policy = models.Policy{Roles: filterRoles, Resources: []models.PolicyResource{policyResource}}

	policy.Resources[0].AccountID = myAccount.GUID

	servicePolicyAPI := iamClient.ServicePolicies()

	createdPolicy, err := servicePolicyAPI.Create(policyBoundTo.ScopeSegment(), sID.IAMID, policy)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdPolicy)

	createdPolicy, err = servicePolicyAPI.Get(policyBoundTo.ScopeSegment(), sID.IAMID, createdPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdPolicy)

	err = servicePolicyAPI.Delete(iamv1.ServicePolicyIdentifier{
		Scope:    policyBoundTo.ScopeSegment(),
		IAMID:    sID.IAMID,
		PolicyID: createdPolicy.ID,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = serviceIDAPI.Delete(sID.UUID)
	if err != nil {
		log.Fatal(err)
	}

}
