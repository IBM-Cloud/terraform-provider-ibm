package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMIAMServiceID() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMServiceIDRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the serviceID",
				Type:        schema.TypeString,
				Required:    true,
			},

			"service_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"bound_to": {
							Description: "bound to of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"crn": {
							Description: "CRN of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": {
							Description: "description of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": {
							Description: "Version of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},

						"locked": {
							Description: "lock state of the serviceID",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIAMServiceIDRead(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
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

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount).String()

	serviceIDS, err := iamClient.ServiceIds().FindByName(boundTo, name)
	if err != nil {
		return err
	}

	if len(serviceIDS) == 0 {
		return fmt.Errorf("No serviceID found with name [%s]", name)

	}

	serviceIDListMap := make([]map[string]interface{}, 0, len(serviceIDS))
	for _, serviceID := range serviceIDS {
		l := map[string]interface{}{
			"id":          serviceID.UUID,
			"bound_to":    serviceID.BoundTo,
			"version":     serviceID.Version,
			"description": serviceID.Description,
			"crn":         serviceID.CRN,
			"locked":      serviceID.Locked,
		}
		serviceIDListMap = append(serviceIDListMap, l)
	}
	d.SetId(name)
	d.Set("service_ids", serviceIDListMap)
	return nil
}
