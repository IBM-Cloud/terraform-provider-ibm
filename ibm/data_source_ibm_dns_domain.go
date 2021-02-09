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
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
)

func dataSourceIBMDNSDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDNSDomainRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "A domain record's internal identifier",
				Type:        schema.TypeInt,
				Computed:    true,
			},

			"name": &schema.Schema{
				Description: "The name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMDNSDomainRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	name := d.Get("name").(string)

	names, err := service.
		Filter(filter.Build(filter.Path("domains.name").Eq(name))).
		Mask("id,name").
		GetDomains()

	if err != nil {
		return fmt.Errorf("Error retrieving domain: %s", err)
	}

	if len(names) == 0 {
		return fmt.Errorf("No domain found with name [%s]", name)
	}

	d.SetId(fmt.Sprintf("%d", *names[0].Id))
	return nil
}
