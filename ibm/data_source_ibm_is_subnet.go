package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func dataSourceIBMISSubnetRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Get("identifier").(string)
	if userDetails.generation == 1 {
		err := classicSubnetGet(d, meta, id)
		d.SetId(id)
		if err != nil {
			return err
		}
	} else {
		err := subnetGet(d, meta, id)
		d.SetId(id)
		if err != nil {
			return err
		}
	}
	return nil
}
