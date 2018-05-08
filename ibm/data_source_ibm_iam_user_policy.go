package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// Data source to find all the policies for a user in a particular account
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
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	accountGUID := d.Get("account_guid").(string)
	ibmID := d.Get("ibm_id").(string)
	userID, err := getIBMID(accountGUID, ibmID, meta)
	if err != nil {
		return err
	}
	userPolicies, err := iamClient.IAMPolicy().List(accountGUID, userID)
	if err != nil {
		return fmt.Errorf("Error retrieving policies %s", err)
	}
	policies := userPolicies.Policies
	accountPolicyListMap := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {
		roles := flattenIAMPolicyRoles(policy.Roles)
		resources, err := flattenIAMPolicyResource(policy.Resources, iamClient)
		if err != nil {
			return err
		}
		l := map[string]interface{}{
			"id":        policy.ID,
			"roles":     roles,
			"resources": resources,
		}
		accountPolicyListMap = append(accountPolicyListMap, l)
	}
	//Id is composed of user in a particular account
	d.SetId(fmt.Sprintf("%s/%s", ibmID, accountGUID))
	d.Set("policies", accountPolicyListMap)
	return nil
}
