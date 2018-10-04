package ibm

import (
	"fmt"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMIAMAccessGroupMembers() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMAccessGroupMembersCreate,
		Read:     resourceIBMIAMAccessGroupMembersRead,
		Update:   resourceIBMIAMAccessGroupMembersUpdate,
		Delete:   resourceIBMIAMAccessGroupMembersDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"access_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the access group",
			},

			"ibm_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"iam_service_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iam_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMIAMAccessGroupMembersCreate(d *schema.ResourceData, meta interface{}) error {
	iamuumClient, err := meta.(ClientSession).IAMUUMAPI()
	if err != nil {
		return err
	}

	grpID := d.Get("access_group_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	var userids, serviceids []string

	users := expandStringList(d.Get("ibm_ids").(*schema.Set).List())
	services := expandStringList(d.Get("iam_service_ids").(*schema.Set).List())

	if len(users) == 0 && len(services) == 0 {
		return fmt.Errorf("Provide either `ibm_ids` or `iam_service_ids`")

	}

	userids, err = flattenUserIds(accountID, users, meta)
	if err != nil {
		return err
	}

	serviceids, err = flattenServiceIds(services, meta)
	if err != nil {
		return err
	}

	request := prepareMemberAddRequest(userids, serviceids)

	_, err = iamuumClient.AccessGroupMember().Add(grpID, request)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", grpID, time.Now().UTC().String()))

	return resourceIBMIAMAccessGroupMembersRead(d, meta)
}

func resourceIBMIAMAccessGroupMembersRead(d *schema.ResourceData, meta interface{}) error {
	iamuumClient, err := meta.(ClientSession).IAMUUMAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	grpID := parts[0]

	members, err := iamuumClient.AccessGroupMember().List(grpID)
	if err != nil {
		return fmt.Errorf("Error retrieving access group members: %s", err)
	}

	d.Set("access_group_id", grpID)

	mccpAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	accountv1Client, err := meta.(ClientSession).BluemixAcccountv1API()
	if err != nil {
		return err
	}

	users, err := accountv1Client.Accounts().GetAccountUsers(accountID)
	if err != nil {
		return err
	}

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	serviceIDs, err := iamClient.ServiceIds().List(GenerateBoundToCRN(*region, accountID).String())
	if err != nil {
		return err
	}

	d.Set("members", flattenAccessGroupMembers(members, users, serviceIDs))

	return nil
}

func resourceIBMIAMAccessGroupMembersUpdate(d *schema.ResourceData, meta interface{}) error {

	iamuumClient, err := meta.(ClientSession).IAMUUMAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	grpID := parts[0]

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	var removeUsers, addUsers, removeServiceids, addServiceids []string
	o, n := d.GetChange("ibm_ids")
	ou := o.(*schema.Set)
	nu := n.(*schema.Set)

	removeUsers = expandStringList(ou.Difference(nu).List())
	addUsers = expandStringList(nu.Difference(ou).List())

	os, ns := d.GetChange("iam_service_ids")
	osi := os.(*schema.Set)
	nsi := ns.(*schema.Set)

	removeServiceids = expandStringList(osi.Difference(nsi).List())
	addServiceids = expandStringList(nsi.Difference(osi).List())

	if len(addUsers) > 0 || len(addServiceids) > 0 && !d.IsNewResource() {
		var userids, serviceids []string
		userids, err = flattenUserIds(accountID, addUsers, meta)
		if err != nil {
			return err
		}

		serviceids, err = flattenServiceIds(addServiceids, meta)
		if err != nil {
			return err
		}
		request := prepareMemberAddRequest(userids, serviceids)

		_, err = iamuumClient.AccessGroupMember().Add(grpID, request)
		if err != nil {
			return err
		}

	}
	if len(removeUsers) > 0 || len(removeServiceids) > 0 && !d.IsNewResource() {
		iamClient, err := meta.(ClientSession).IAMAPI()
		if err != nil {
			return err
		}
		for _, u := range removeUsers {
			user, err := getAccountUser(accountID, u, meta)
			if err != nil {
				return err
			}
			err = iamuumClient.AccessGroupMember().Remove(grpID, user.IbmUniqueId)
			if err != nil {
				return err
			}

		}

		for _, s := range removeServiceids {
			serviceID, err := iamClient.ServiceIds().Get(s)
			if err != nil {
				return err
			}
			err = iamuumClient.AccessGroupMember().Remove(grpID, serviceID.IAMID)
			if err != nil {
				return err
			}

		}
	}

	return resourceIBMIAMAccessGroupMembersRead(d, meta)

}

func resourceIBMIAMAccessGroupMembersDelete(d *schema.ResourceData, meta interface{}) error {
	iamuumClient, err := meta.(ClientSession).IAMUUMAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	grpID := parts[0]

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	users := expandStringList(d.Get("ibm_ids").(*schema.Set).List())

	for _, name := range users {
		user, err := getAccountUser(userDetails.userAccount, name, meta)
		if err != nil {
			return err
		}
		err = iamuumClient.AccessGroupMember().Remove(grpID, user.IbmUniqueId)
		if err != nil {
			return err
		}

	}

	services := expandStringList(d.Get("iam_service_ids").(*schema.Set).List())

	for _, id := range services {
		serviceID, err := getServiceID(id, meta)
		if err != nil {
			return err
		}
		err = iamuumClient.AccessGroupMember().Remove(grpID, serviceID.IAMID)
		if err != nil {
			return err
		}
	}

	d.SetId("")

	return nil
}

func prepareMemberAddRequest(userIds, serviceIds []string) (req iamuumv1.AddGroupMemberRequest) {
	req.Members = make([]models.AccessGroupMember, len(userIds)+len(serviceIds))
	var i = 0
	for _, id := range userIds {
		req.Members[i] = models.AccessGroupMember{
			ID:   id,
			Type: iamuumv1.AccessGroupMemberUser,
		}
		i++
	}

	for _, id := range serviceIds {
		req.Members[i] = models.AccessGroupMember{
			ID:   id,
			Type: iamuumv1.AccessGroupMemberService,
		}
		i++
	}
	return
}

func getServiceID(id string, meta interface{}) (models.ServiceID, error) {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return models.ServiceID{}, err
	}
	serviceID, err := iamClient.ServiceIds().Get(id)
	if err != nil {
		return models.ServiceID{}, err
	}

	return serviceID, nil
}
