package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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

			isSubnetResourceGroup: {
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
		err := classicSubnetGetByID(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := subnetGetByID(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicSubnetGetByID(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	d.SetId(*subnet.ID)
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CidrBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, *subnet.AvailableIpv4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, *subnet.TotalIpv4AddressCount)
	if subnet.NetworkAcl != nil {
		d.Set(isSubnetNetworkACL, *subnet.NetworkAcl.ID)
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID)
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, *subnet.Status)
	d.Set(isSubnetZone, *subnet.Zone.Name)
	d.Set(isSubnetVPC, *subnet.Vpc.ID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/subnets")
	d.Set(ResourceName, *subnet.Name)
	d.Set(ResourceCRN, *subnet.Crn)
	d.Set(ResourceStatus, *subnet.Status)
	return nil
}

func subnetGetByID(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	d.SetId(*subnet.ID)
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CidrBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, *subnet.AvailableIpv4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, *subnet.TotalIpv4AddressCount)
	if subnet.NetworkAcl != nil {
		d.Set(isSubnetNetworkACL, *subnet.NetworkAcl.ID)
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID)
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, *subnet.Status)
	d.Set(isSubnetZone, *subnet.Zone.Name)
	d.Set(isSubnetVPC, *subnet.Vpc.ID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/subnets")
	d.Set(ResourceName, *subnet.Name)
	d.Set(ResourceCRN, *subnet.Crn)
	d.Set(ResourceStatus, *subnet.Status)
	if subnet.ResourceGroup != nil {
		d.Set(isSubnetResourceGroup, *subnet.ResourceGroup.ID)
		d.Set(ResourceGroupName, *subnet.ResourceGroup.Name)
	}
	return nil
}
