package ibm

import (
	"strconv"
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
						"crn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4_cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_ipv4_address_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_acl": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_ipv4_address_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone": {
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

		var aac string = strconv.FormatInt(subnet.AvailableIPV4AddressCount, 10)
		var tac string = strconv.FormatInt(subnet.TotalIPV4AddressCount, 10)

		subn := make(map[string]string)
		subn["name"] = subnet.Name
		subn["id"] = string(subnet.ID)
		subn["status"] = subnet.Status
		subn["crn"] = subnet.Crn
		subn["ipv4_cidr_block"] = subnet.IPV4CidrBlock
		subn["ipv6_cidr_block"] = subnet.IPV6CidrBlock
		subn["available_ipv4_address_count"] = aac
		subn["network_acl"] = subnet.NetworkACL.Name
		if subnet.PublicGateway != nil {
			subn["public_gateway"] = string(subnet.PublicGateway.ID)
		}
		if subnet.ResourceGroup != nil {
			subn["resource_group"] = string(subnet.ResourceGroup.ID)
		}
		subn["total_ipv4_address_count"] = tac
		subn["vpc"] = string(subnet.Vpc.ID)
		subn["zone"] = subnet.Zone.Name

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
