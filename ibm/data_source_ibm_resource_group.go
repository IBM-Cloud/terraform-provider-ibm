package ibm

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMResourceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMResourceGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description:   "Resource group name",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"is_default"},
			},
			"is_default": {
				Description:   "Default Resource group",
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"name"},
			},
		},
	}
}

func dataSourceIBMResourceGroupRead(d *schema.ResourceData, meta interface{}) error {
	rsManagementAPI, err := meta.(ClientSession).ResourceManagementAPI()
	if err != nil {
		return err
	}
	rsGroup := rsManagementAPI.ResourceGroup()
	defaultGrp := d.Get("is_default").(bool)
	var name string
	if n, ok := d.GetOk("name"); ok {
		name = n.(string)
	}
	var grp []models.ResourceGroup
	if defaultGrp {
		resourceGroupQuery := management.ResourceGroupQuery{
			Default: true,
		}

		grp, err = rsGroup.List(&resourceGroupQuery)

		if err != nil {
			return fmt.Errorf("Error retrieving default resource group: %s", err)
		}
		d.SetId(grp[0].ID)

	} else if name != "" {
		grp, err := rsGroup.FindByName(nil, name)
		if err != nil {
			return fmt.Errorf("Error retrieving resource group %s: %s", name, err)
		}
		d.SetId(grp[0].ID)

	} else {
		return fmt.Errorf("Missing required properties. Need a resource group name, or the is_default true")
	}

	return nil
}
