package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMIAMUserPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMUserPolicyRead,

		Schema: map[string]*schema.Schema{
			"account_guid": {
				Description: "The guid of the account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ibm_id": {
				Description: "The ibm id of user",
				Type:        schema.TypeString,
				Required:    true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString,

										Computed: true,
									},
								},
							},
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_instance": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"space_guid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"organization_guid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIAMUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	accClient, err := meta.(ClientSession).BluemixAcccountAPI()
	if err != nil {
		return err
	}

	scope := d.Get("account_guid").(string)
	ibmId := d.Get("ibm_id").(string)

	account, err := accClient.Accounts().Get(scope)
	if err != nil {
		return fmt.Errorf("Error retrieving account: %s", err)
	}
	userId, err := getIBMUniqueIdOfUser(scope, ibmId, meta)
	if userId == "" || err != nil {
		return fmt.Errorf("User %q does not exist in the account:%q", ibmId, account.Name)
	}

	accessPolicyList, err := iamClient.IAMPolicy().List(scope, userId)
	if err != nil {
		return fmt.Errorf("Error retrieving policies %s", err)
	}

	policies := accessPolicyList.Policies
	accountPolicyListMap := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {

		rolesMaps := make(map[string]string)
		rolesMaps[VIEWER_ID] = VIEWER
		rolesMaps[ADMINISTRATOR_ID] = ADMINISTRATOR
		rolesMaps[EDITOR_ID] = EDITOR
		rolesMaps[OPERATOR_ID] = OPERATOR
		roles := flattenIAMPolicyRoles(policy.Roles, rolesMaps)
		resources := flattenIAMPolicyResource(policy.Resources, iamClient)

		l := map[string]interface{}{
			"id":        policy.ID,
			"roles":     roles,
			"resources": resources,
		}
		accountPolicyListMap = append(accountPolicyListMap, l)
	}
	d.SetId(scope)
	d.Set("policies", accountPolicyListMap)
	return nil
}
