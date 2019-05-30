package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
)

func dataSourceIBMISVPC() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISVPCRead,

		Schema: map[string]*schema.Schema{
			isVPCDefaultNetworkACL: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCClassicAccess: {
				Type:     schema.TypeBool,
				Computed: true,
			},

			isVPCName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPCResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCTags: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)

	name := d.Get(isVPCName).(string)

	vpcs, _, err := vpcC.List("")
	if err != nil {
		return err
	}

	for _, vpc := range vpcs {
		if vpc.Name == name {
			d.SetId(vpc.ID.String())
			d.Set("id", vpc.ID.String())
			d.Set(isVPCName, vpc.Name)
			d.Set(isVPCClassicAccess, vpc.ClassicAccess)
			d.Set(isVPCStatus, vpc.Status)
			if vpc.DefaultNetworkACL != nil {
				d.Set(isVPCDefaultNetworkACL, vpc.DefaultNetworkACL.ID)
			} else {
				d.Set(isVPCDefaultNetworkACL, nil)
			}
			tags, err := GetTagsUsingCRN(meta, vpc.Crn)
			if err != nil {
				return fmt.Errorf(
					"Error on get of resource vpc (%s) tags: %s", d.Id(), err)
			}
			d.Set(isVPCTags, tags)
			return nil
		}
	}
	return fmt.Errorf("No VPC found with name %s", name)
}
