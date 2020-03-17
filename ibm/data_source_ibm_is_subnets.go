package ibm

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
)

const (
	isSubnets = "subnets"
)

func dataSourceIBMISSubnets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSubnetsRead,

		Schema: map[string]*schema.Schema{

			isSubnets: {
				Type:        schema.TypeList,
				Description: "List of subnets",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4_cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISSubnetsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	subnetC := network.NewSubnetClient(sess)
	availableSubnets, _, err := subnetC.List("")
	if err != nil {
		return err
	}

	subnets := make([]map[string]string, len(availableSubnets))
	for i, subnet := range availableSubnets {

		subn := make(map[string]string)
		subn["name"] = subnet.Name
		subn["id"] = string(subnet.ID)
		subn["status"] = subnet.Status
		subn["ipv4_cidr_block"] = subnet.IPV4CidrBlock

		subnets[i] = subn
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isSubnets, subnets)
	return nil
}

// dataSourceIBMISSubnetsId returns a reasonable ID for a subnet list.
func dataSourceIBMISSubnetsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
