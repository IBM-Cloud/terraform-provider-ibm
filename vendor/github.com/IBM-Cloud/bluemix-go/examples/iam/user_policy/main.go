package main

import (
	"flag"
	"log"
	"os"
	"strings"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
	"github.com/IBM-Cloud/bluemix-go/utils"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var userEmail string
	flag.StringVar(&userEmail, "userEmail", "", "Email of the user to be invited")

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

	trace.Logger = trace.NewLogger("true")
	c := new(bluemix.Config)
	flag.BoolVar(&c.Debug, "debug", false, "Show full trace if on")
	flag.Parse()

	if org == "" || userEmail == "" || roles == "" {
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

	accClient1, err := accountv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPIV1 := accClient1.Accounts()
	//Get list of users under account
	user, err := accountAPIV1.InviteAccountUser(myAccount.GUID, userEmail)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}

	userPolicyAPI := iamClient.UserPolicies()

	serviceRolesAPI := iamClient.ServiceRoles()

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

	userDetails, err := accountAPIV1.FindAccountUserByUserId(myAccount.GUID, userEmail)
	if err != nil {
		log.Fatal(err)
	}

	createdPolicy, err := userPolicyAPI.Create("a/"+myAccount.GUID, userDetails.IbmUniqueId, policy)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdPolicy)

	err = userPolicyAPI.Delete("a/"+myAccount.GUID, userDetails.IbmUniqueId, createdPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = accountAPIV1.DeleteAccountUser(myAccount.GUID, userDetails.Id)
	if err != nil {
		log.Fatal(err)
	}

}
