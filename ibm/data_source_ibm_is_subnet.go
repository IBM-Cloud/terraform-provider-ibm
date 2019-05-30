package ibm

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
)

func dataSourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSubnetRead,

		Schema: map[string]*schema.Schema{

			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},

			isSubnetIpv4CidrBlock: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetIpv6CidrBlock: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetAvailableIpv4AddressCount: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetTotalIpv4AddressCount: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isSubnetIPVersion: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isSubnetName: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetNetworkACL: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetPublicGateway: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetVPC: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetZone: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISSubnetRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	subnetC := network.NewSubnetClient(sess)

	subnet, err := subnetC.Get(d.Get("identifier").(string))
	if err != nil {
		return err
	}
	d.SetId(subnet.ID.String())
	d.Set("id", subnet.ID.String())
	d.Set(isSubnetName, subnet.Name)
	d.Set(isSubnetIPVersion, subnet.IPVersion)
	d.Set(isSubnetIpv4CidrBlock, subnet.IPV4CidrBlock)
	d.Set(isSubnetIpv6CidrBlock, subnet.IPV6CidrBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, subnet.AvailableIPV4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, subnet.TotalIPV4AddressCount)
	if subnet.NetworkACL != nil {
		d.Set(isSubnetNetworkACL, subnet.NetworkACL.ID.String())
	} else {
		d.Set(isSubnetNetworkACL, nil)
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, subnet.PublicGateway.ID.String())
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, subnet.Status)
	d.Set(isSubnetZone, subnet.Zone.Name)
	d.Set(isSubnetVPC, subnet.Vpc.ID.String())

	return nil
}
