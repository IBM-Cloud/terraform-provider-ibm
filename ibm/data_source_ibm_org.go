/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

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
