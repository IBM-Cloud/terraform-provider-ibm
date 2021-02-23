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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceIBMIAMRoleAction() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIBMIAMRoleActionRead,

		Schema: map[string]*schema.Schema{
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Service Name",
				ForceNew:    true,
			},
			"reader": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Reader action ids",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"manager": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "manager action ids",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"reader_plus": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "readerplus action ids",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"writer": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "writer action ids",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

}

func datasourceIBMIAMRoleActionRead(d *schema.ResourceData, meta interface{}) error {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}

	serviceName := d.Get("service").(string)
	d.SetId(serviceName)
	serviceRoles, err := iampapv2Client.IAMRoles().ListServiceRoles(serviceName)
	if err != nil {
		return err
	}

	d.Set("reader", flattenActionbyDisplayName("Reader", serviceRoles))
	d.Set("manager", flattenActionbyDisplayName("Manager", serviceRoles))
	d.Set("reader_plus", flattenActionbyDisplayName("ReaderPlus", serviceRoles))
	d.Set("writer", flattenActionbyDisplayName("Writer", serviceRoles))

	return nil
}
