/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMIAMAuthorizationPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMAuthorizationPolicyCreate,
		Read:     resourceIBMIAMAuthorizationPolicyRead,
		Update:   resourceIBMIAMAuthorizationPolicyUpdate,
		Delete:   resourceIBMIAMAuthorizationPolicyDelete,
		Exists:   resourceIBMIAMAuthorizationPolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"source_service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source service name",
				ForceNew:    true,
			},

			"target_service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target service name",
			},

			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Role names of the policy definition",
			},

			"source_resource_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The source resource instance Id",
			},

			"target_resource_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The target resource instance Id",
			},

			"source_resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The source resource group Id",
			},

			"target_resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The target resource group Id",
			},

			"source_resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Resource type of source service",
			},

			"target_resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Resource type of target service",
			},

			"source_service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Account GUID of source service",
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMIAMAuthorizationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	sourceServiceName := d.Get("source_service_name").(string)
	targetServiceName := d.Get("target_service_name").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	policy := iampapv1.Policy{
		Type: iampapv1.AuthorizationPolicyType,
	}

	sourceServiceAccount := userDetails.userAccount

	if account, ok := d.GetOk("source_service_account"); ok {
		sourceServiceAccount = account.(string)

	}

	policy.Subjects = []iampapv1.Subject{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "accountId",
					Value: sourceServiceAccount,
				},
				{
					Name:  "serviceName",
					Value: sourceServiceName,
				},
			},
		},
	}

	policy.Resources = []iampapv1.Resource{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "accountId",
					Value: userDetails.userAccount,
				},
				{
					Name:  "serviceName",
					Value: targetServiceName,
				},
			},
		},
	}
	if sID, ok := d.GetOk("source_resource_instance_id"); ok {
		policy.Subjects[0].SetServiceInstance(sID.(string))
	}

	if tID, ok := d.GetOk("target_resource_instance_id"); ok {
		policy.Resources[0].SetServiceInstance(tID.(string))
	}

	if sType, ok := d.GetOk("source_resource_type"); ok {
		policy.Subjects[0].SetResourceType(sType.(string))
	}

	if tType, ok := d.GetOk("target_resource_type"); ok {
		policy.Resources[0].SetResourceType(tType.(string))
	}

	if sResGrpID, ok := d.GetOk("source_resource_group_id"); ok {
		policy.Subjects[0].SetResourceGroupID(sResGrpID.(string))
	}

	if tResGrpID, ok := d.GetOk("target_resource_group_id"); ok {
		policy.Resources[0].SetResourceGroupID(tResGrpID.(string))
	}

	roles, err := getAuthorizationRolesByName(expandStringList(d.Get("roles").([]interface{})), sourceServiceName, targetServiceName, meta)
	if err != nil {
		return err
	}

	policy.Roles = iampapv1.ConvertRoleModels(roles)

	authPolicy, err := iampapClient.V1Policy().Create(policy)

	if err != nil {
		return fmt.Errorf("Error creating authorization policy: %s", err)
	}

	d.SetId(authPolicy.ID)

	return resourceIBMIAMAuthorizationPolicyRead(d, meta)
}

func resourceIBMIAMAuthorizationPolicyRead(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	authorizationPolicy, err := iampapClient.V1Policy().Get(d.Id())
	if err != nil {
		return fmt.Errorf("Error retrieving authorizationPolicy: %s", err)
	}
	roles := make([]string, len(authorizationPolicy.Roles))
	for i, role := range authorizationPolicy.Roles {
		roles[i] = role.Name
	}
	d.Set("roles", roles)
	d.Set("version", authorizationPolicy.Version)
	source := authorizationPolicy.Subjects[0]
	target := authorizationPolicy.Resources[0]
	d.Set("source_service_name", source.ServiceName())
	d.Set("target_service_name", target.ServiceName())
	d.Set("source_resource_instance_id", source.ServiceInstance())
	d.Set("target_resource_instance_id", target.ServiceInstance())
	d.Set("source_resource_type", source.ResourceType())
	d.Set("target_resource_type", target.ResourceType())
	d.Set("source_service_account", source.AccountID())
	d.Set("source_resource_group_id", source.ResourceGroupID())
	d.Set("target_resource_group_id", target.ResourceGroupID())
	return nil
}

func resourceIBMIAMAuthorizationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMIAMAuthorizationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	authorizationPolicyID := d.Id()

	err = iampapClient.V1Policy().Delete(authorizationPolicyID)
	if err != nil {
		log.Printf(
			"Error deleting authorization policy: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMAuthorizationPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return false, err
	}

	authorizationPolicy, err := iampapClient.V1Policy().Get(d.Id())
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return authorizationPolicy.ID == d.Id(), nil
}

func getAuthorizationRolesByName(roleNames []string, sourceServiceName string, targetServiceName string, meta interface{}) ([]models.PolicyRole, error) {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return []models.PolicyRole{}, err
	}

	iamRepo := iamClient.ServiceRoles()
	roles, err := iamRepo.ListAuthorizationRoles(sourceServiceName, targetServiceName)
	if err != nil {
		return []models.PolicyRole{}, err
	}

	filteredRoles := []models.PolicyRole{}
	filteredRoles, err = getRolesFromRoleNames(roleNames, roles)
	if err != nil {
		return []models.PolicyRole{}, err
	}
	return filteredRoles, nil
}
