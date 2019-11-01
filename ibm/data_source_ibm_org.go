package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMOrg() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMOrgRead,

		Schema: map[string]*schema.Schema{
			"org": {
				Description: "Org name, for example myorg@domain",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMOrgRead(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	orgAPI := cfAPI.Organizations()
	org := d.Get("org").(string)
	orgFields, err := orgAPI.FindByName(org, BluemixRegion)
	if err != nil {
		return fmt.Errorf("Error retrieving organisation: %s", err)
	}

	d.SetId(orgFields.GUID)

	return nil
}
