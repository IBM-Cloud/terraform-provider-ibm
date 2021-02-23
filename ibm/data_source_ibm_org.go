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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMOrg() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMOrgRead,

		Schema: map[string]*schema.Schema{
			"org": {
				Description:  "Org name, for example myorg@domain",
				Type:         schema.TypeString,
				Optional:     true,
				Deprecated:   "use name instead",
				ExactlyOneOf: []string{"org", "name"},
			},
			"name": {
				Description:  "Org name, for example myorg@domain",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"org", "name"},
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
	var org string
	if v, ok := d.GetOk("name"); ok {
		org = v.(string)
	}
	if v, ok := d.GetOk("org"); ok {
		org = v.(string)
	}

	orgFields, err := orgAPI.FindByName(org, BluemixRegion)
	if err != nil {
		return fmt.Errorf("Error retrieving organisation: %s", err)
	}
	d.SetId(orgFields.GUID)

	return nil
}
