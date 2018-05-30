package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// Data source to find all the policies for a serviceID
func dataSourceIBMIAMServicePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMServicePolicyRead,

		Schema: map[string]*schema.Schema{
			"iam_service_id": {
				Description: "UUID of ServiceID",
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
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Role names of the policy definition",
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Service name of the policy definition",
									},
									"resource_instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "ID of resource instance of the policy definition",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region of the policy definition",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource type of the policy definition",
									},
									"resource": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource of the policy definition",
									},
									"resource_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the resource group.",
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

func dataSourceIBMIAMServicePolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}

	serviceIDUUID := d.Get("iam_service_id").(string)

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	mccpAPI, err := meta.(ClientSession).MccpAPI()
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

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return err
	}

	policies, err := iamClient.ServicePolicies().List(boundTo.ScopeSegment(), serviceID.IAMID)
	if err != nil {
		return err
	}

	servicePolicies := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {
		roles := make([]string, len(policy.Roles))
		for i, role := range policy.Roles {
			roles[i] = role.DisplayName
		}
		resources := flattenPolicyResource(policy.Resources)
		p := map[string]interface{}{
			"id":        fmt.Sprintf("%s/%s", serviceIDUUID, policy.ID),
			"roles":     roles,
			"resources": resources,
		}
		servicePolicies = append(servicePolicies, p)
	}
	d.SetId(serviceIDUUID)
	d.Set("policies", servicePolicies)
	return nil
}
