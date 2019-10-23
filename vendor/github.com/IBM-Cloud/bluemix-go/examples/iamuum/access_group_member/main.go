package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/utils"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var accessGroup string
	flag.StringVar(&accessGroup, "accessGroup", "", "Bluemix access group name")

	var user string
	flag.StringVar(&user, "user", "", "IBM-id or email id of the user to be added")

	var serviceID string
	flag.StringVar(&serviceID, "serviceID", "", "Bluemix service id name")

	flag.Parse()
	if org == "" || accessGroup == "" || user == "" || serviceID == "" {
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

	regionAPI := client.Regions()
	region, err := regionAPI.FindRegionByName(sess.Config.Region)
	if err != nil {
		log.Fatal(err)
	}

	iamClient, err := iamv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	serviceIdAPI := iamClient.ServiceIds()

	boundTo := utils.GenerateBoundToCRN(*region, myAccount.GUID).String()

	serviceData := models.ServiceID{
		Name:    serviceID,
		BoundTo: boundTo,
	}
	sID, err := serviceIdAPI.Create(serviceData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sID)

	accClient1, err := accountv1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPIV1 := accClient1.Accounts()
	//Get list of users under account
	userres, err := accountAPIV1.InviteAccountUser(myAccount.GUID, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userres)

	userDetails, err := accountAPIV1.FindAccountUserByUserId(myAccount.GUID, user)
	if err != nil {
		log.Fatal(err)
	}

	accessGroupMemAPI := iamuumClient.AccessGroupMember()

	var members []models.AccessGroupMember

	grpmem1 := models.AccessGroupMember{
		ID:   userDetails.IbmUniqueId,
		Type: iamuumv1.AccessGroupMemberUser,
	}

	members = append(members, grpmem1)

	grpmem2 := models.AccessGroupMember{
		ID:   sID.IAMID,
		Type: iamuumv1.AccessGroupMemberService,
	}

	members = append(members, grpmem2)

	addRequest := iamuumv1.AddGroupMemberRequest{
		Members: members,
	}

	resp, err := accessGroupMemAPI.Add(agID.ID, addRequest)
	if err != nil {
		log.Fatal(err)
	}

	err = accessGroupMemAPI.Remove(agID.ID, resp.Members[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	err = accessGroupMemAPI.Remove(agID.ID, resp.Members[1].ID)
	if err != nil {
		log.Fatal(err)
	}

	err = accessGroupAPI.Delete(agID.ID, false)
	if err != nil {
		log.Fatal(err)
	}

	err = accountAPIV1.DeleteAccountUser(myAccount.GUID, userDetails.Id)
	if err != nil {
		log.Fatal(err)
	}

}
