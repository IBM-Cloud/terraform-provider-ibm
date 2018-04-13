package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	v1 "github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
)

func main() {
	c := new(bluemix.Config)
	flag.BoolVar(&c.Debug, "debug", false, "Show full trace if on")
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	var region string
	flag.StringVar(&region, "region", "us-south", "Bluemix region")

	var user_id string
	flag.StringVar(&user_id, "user_id", "", "Bluemix user id")

	var service_name string
	flag.StringVar(&service_name, "service_name", "", "Bluemix service name")

	var role string
	flag.StringVar(&role, "role", "", "Access Policy role. Ex: crn:v1:bluemix:public:iam::::role:Viewer")

	flag.Parse()
	if org == "" || space == "" || user_id == "" {
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
	myorg, err := orgAPI.FindByName(org, region)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, region)
	if err != nil {
		log.Fatal(err)
	}
	var roles = []v1.Roles{
		v1.Roles{
			ID: role,
		},
	}
	var resources = []v1.Resources{
		v1.Resources{
			ServiceName: service_name,
		},
	}
	var accessPolicyRequest = v1.AccessPolicyRequest{
		Roles:     roles,
		Resources: resources,
	}

	iampapClient, err := v1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	iampapAPI := iampapClient.IAMPolicy()
	accessPolicyCreateResp, _, err := iampapAPI.Create(myAccount.GUID, user_id, accessPolicyRequest)
	if err != nil {
		log.Fatal(err)
	}

	accessPolicyListResp, err := iampapAPI.List(myAccount.GUID, user_id)
	if err != nil {
		log.Fatal(err)
	}

	policies := accessPolicyListResp.Policies

	for i, policy := range policies {
		fmt.Println("Policy Id ", i, policy.ID)
	}

	err = iampapAPI.Delete(myAccount.GUID, user_id, accessPolicyCreateResp.ID)
	if err != nil {
		log.Fatal(err)
	}
}
