package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSubnetRead,

		Schema: map[string]*schema.Schema{

			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isSubnetName, "identifier"},
				ValidateFunc: InvokeDataSourceValidator("ibm_is_subnet", "identifier"),
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

			isSubnetName: {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{isSubnetName, "identifier"},
				ValidateFunc: InvokeDataSourceValidator("ibm_is_subnet", isSubnetName),
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

func dataSourceIBMISSubnetValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isSubnetName,
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})

	ibmISSubnetDataSourceValidator := ResourceValidator{ResourceName: "ibm_is_subnet", Schema: validateSchema}
	return &ibmISSubnetDataSourceValidator
}

func dataSourceIBMISSubnetRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	if userDetails.generation == 1 {
		err := classicSubnetGetByNameOrID(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := subnetGetByNameOrID(d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicSubnetGetByNameOrID(d *schema.ResourceData, meta interface{}) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	var subnet *vpcclassicv1.Subnet

	if v, ok := d.GetOk("identifier"); ok {
		id := v.(string)
		getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
			ID: &id,
		}
		subnetinfo, response, err := sess.GetSubnet(getSubnetOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
		}
		subnet = subnetinfo
	}
	if v, ok := d.GetOk(isSubnetName); ok {
		name := v.(string)
		getSubnetsListOptions := &vpcclassicv1.ListSubnetsOptions{}
		subnetsCollection, response, err := sess.ListSubnets(getSubnetsListOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Subnets List : %s\n%s", err, response)
		}
		for _, subnetInfo := range subnetsCollection.Subnets {
			if *subnetInfo.Name == name {
				subnet = &subnetInfo
				break
			}
		}
	}
	d.SetId(*subnet.ID)
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CIDRBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, *subnet.AvailableIpv4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, *subnet.TotalIpv4AddressCount)
	if subnet.NetworkACL != nil {
		d.Set(isSubnetNetworkACL, *subnet.NetworkACL.ID)
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID)
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, *subnet.Status)
	d.Set(isSubnetZone, *subnet.Zone.Name)
	d.Set(isSubnetVPC, *subnet.VPC.ID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/subnets")
	d.Set(ResourceName, *subnet.Name)
	d.Set(ResourceCRN, *subnet.CRN)
	d.Set(ResourceStatus, *subnet.Status)
	return nil
}

func subnetGetByNameOrID(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	var subnet *vpcv1.Subnet

	if v, ok := d.GetOk("identifier"); ok {
		id := v.(string)
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnetinfo, response, err := sess.GetSubnet(getSubnetOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
		}
		subnet = subnetinfo
	} else if v, ok := d.GetOk(isSubnetName); ok {
		name := v.(string)
		getSubnetsListOptions := &vpcv1.ListSubnetsOptions{}
		subnetsCollection, response, err := sess.ListSubnets(getSubnetsListOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Subnets List : %s\n%s", err, response)
		}
		for _, subnetInfo := range subnetsCollection.Subnets {
			if *subnetInfo.Name == name {
				subnet = &subnetInfo
				break
			}
		}
	}

	d.SetId(*subnet.ID)
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CIDRBlock)
	d.Set(isSubnetAvailableIpv4AddressCount, *subnet.AvailableIpv4AddressCount)
	d.Set(isSubnetTotalIpv4AddressCount, *subnet.TotalIpv4AddressCount)
	if subnet.NetworkACL != nil {
		d.Set(isSubnetNetworkACL, *subnet.NetworkACL.ID)
	}
	if subnet.PublicGateway != nil {
		d.Set(isSubnetPublicGateway, *subnet.PublicGateway.ID)
	} else {
		d.Set(isSubnetPublicGateway, nil)
	}
	d.Set(isSubnetStatus, *subnet.Status)
	d.Set(isSubnetZone, *subnet.Zone.Name)
	d.Set(isSubnetVPC, *subnet.VPC.ID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/subnets")
	d.Set(ResourceName, *subnet.Name)
	d.Set(ResourceCRN, *subnet.CRN)
	d.Set(ResourceStatus, *subnet.Status)
	if subnet.ResourceGroup != nil {
		d.Set(isSubnetResourceGroup, *subnet.ResourceGroup.ID)
		d.Set(ResourceGroupName, *subnet.ResourceGroup.Name)
	}
	return nil
}
