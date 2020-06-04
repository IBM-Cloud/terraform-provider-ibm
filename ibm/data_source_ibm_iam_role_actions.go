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
