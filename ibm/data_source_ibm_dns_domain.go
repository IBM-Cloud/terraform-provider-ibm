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
