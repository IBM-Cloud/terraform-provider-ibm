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

func dataSourceIBMAppDomainShared() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMAppDomainSharedRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description:  "The name of the shared domain",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateDomainName,
			},
		},
	}
}

func dataSourceIBMAppDomainSharedRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	domainName := d.Get("name").(string)
	shdomain, err := cfClient.SharedDomains().FindByName(domainName)
	if err != nil {
		return fmt.Errorf("Error retrieving shared domain: %s", err)
	}
	d.SetId(shdomain.GUID)
	return nil

}
