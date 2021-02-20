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
