package main

import (
	"flag"
	"log"
	"os"
	"strings"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
	"github.com/IBM-Cloud/bluemix-go/utils"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var accessGroup string
	flag.StringVar(&accessGroup, "accessgroup", "", "Bluemix access group name")

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

	var serviceType string
	flag.StringVar(&serviceType, "serviceType", "", "service type")

	trace.Logger = trace.NewLogger("true")
	c := new(bluemix.Config)
	flag.BoolVar(&c.Debug, "debug", false, "Show full trace if on")
	flag.Parse()

	if org == "" || roles == "" || accessGroup == "" {
		flag.Usage()
		os.Exit(1)
	}

	sess, err := session.New(c)
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

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	serviceRolesAPI := iamClient.ServiceRoles()

	var policy iampapv1.Policy

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

	policyResource := iampapv1.Resource{}

	if service != "" {
		policyResource.SetServiceName(service)
	}

	if serviceInstance != "" {
		policyResource.SetServiceInstance(serviceInstance)
	}

	if region != "" {
		policyResource.SetRegion(region)
	}

	if resourceType != "" {
		policyResource.SetResourceType(resourceType)
	}

	if resource != "" {
		policyResource.SetResource(resource)
	}

	if resourceGroupID != "" {
		policyResource.SetResourceGroupID(resourceGroupID)
	}

	switch serviceType {
	case "service":
		fallthrough
	case "platform_service":
		policyResource.SetServiceType(serviceType)
	}

	if len(policyResource.Attributes) == 0 {
		policyResource.SetServiceType("service")
	}

	policy = iampapv1.Policy{Roles: iampapv1.ConvertRoleModels(filterRoles), Resources: []iampapv1.Resource{policyResource}}

	policy.Resources[0].SetAccountID(myAccount.GUID)

	iamuumClient, err := iamuumv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accessGroupAPI := iamuumClient.AccessGroup()

	data := models.AccessGroup{
		Name: accessGroup,
	}
	agID, err := accessGroupAPI.Create(data, myAccount.GUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(agID)

	iampapClient, err := iampapv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	policy.Subjects = []iampapv1.Subject{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "access_group_id",
					Value: agID.ID,
				},
			},
		},
	}

	policy.Type = iampapv1.AccessPolicyType

	accessPolicy := iampapClient.V1Policy()

	createdPolicy, err := accessPolicy.Create(policy)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdPolicy)

	getPolicy, err := accessPolicy.Get(createdPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(getPolicy)

	err = accessPolicy.Delete(createdPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = accessGroupAPI.Delete(agID.ID, false)
	if err != nil {
		log.Fatal(err)
	}

}
