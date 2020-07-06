package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isSubnetIpv4CidrBlock             = "ipv4_cidr_block"
	isSubnetIpv6CidrBlock             = "ipv6_cidr_block"
	isSubnetTotalIpv4AddressCount     = "total_ipv4_address_count"
	isSubnetIPVersion                 = "ip_version"
	isSubnetName                      = "name"
	isSubnetNetworkACL                = "network_acl"
	isSubnetPublicGateway             = "public_gateway"
	isSubnetStatus                    = "status"
	isSubnetVPC                       = "vpc"
	isSubnetZone                      = "zone"
	isSubnetAvailableIpv4AddressCount = "available_ipv4_address_count"
	isSubnetResourceGroup             = "resource_group"

	isSubnetProvisioning     = "provisioning"
	isSubnetProvisioningDone = "done"
	isSubnetDeleting         = "deleting"
	isSubnetDeleted          = "done"
)

func resourceIBMISSubnet() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSubnetCreate,
		Read:     resourceIBMISSubnetRead,
		Update:   resourceIBMISSubnetUpdate,
		Delete:   resourceIBMISSubnetDelete,
		Exists:   resourceIBMISSubnetExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isSubnetIpv4CidrBlock: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetTotalIpv4AddressCount},
				ValidateFunc:  validateCIDR,
				Description:   "IPV4 subnet - CIDR block",
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
				Type:          schema.TypeInt,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSubnetIpv4CidrBlock},
			},

			isSubnetIPVersion: {
				Type:         schema.TypeString,
				ForceNew:     true,
				Default:      "ipv4",
				Optional:     true,
				ValidateFunc: validateIPVersion,
				Description:  "Subnet IP version",
			},

			isSubnetName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "Subnet name",
			},

			isSubnetNetworkACL: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},

			isSubnetPublicGateway: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "Public Gateway of the subnet",
			},

			isSubnetStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isSubnetVPC: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "VPC instance ID",
			},

			isSubnetZone: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Subnet zone info",
			},

			isSubnetResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
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

func resourceIBMISSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := d.Get(isSubnetName).(string)
	vpc := d.Get(isSubnetVPC).(string)
	zone := d.Get(isSubnetZone).(string)

	ipv4cidr := ""
	if cidr, ok := d.GetOk(isSubnetIpv4CidrBlock); ok {
		ipv4cidr = cidr.(string)
	}
	ipv4addrcount64 := int64(0)
	ipv4addrcount := 0
	if ipv4addrct, ok := d.GetOk(isSubnetTotalIpv4AddressCount); ok {
		ipv4addrcount = ipv4addrct.(int)
		ipv4addrcount64 = int64(ipv4addrcount)
	}
	if ipv4cidr == "" && ipv4addrcount == 0 {
		return fmt.Errorf("%s or %s need to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
	}

	if ipv4cidr != "" && ipv4addrcount != 0 {
		return fmt.Errorf("only one of %s or %s needs to be provided", isSubnetIpv4CidrBlock, isSubnetTotalIpv4AddressCount)
	}
	isSubnetKey := "subnet_key_" + vpc + "_" + zone
	ibmMutexKV.Lock(isSubnetKey)
	defer ibmMutexKV.Unlock(isSubnetKey)

	acl := ""
	if nwacl, ok := d.GetOk(isSubnetNetworkACL); ok {
		acl = nwacl.(string)
	}

	gw := ""
	if pgw, ok := d.GetOk(isSubnetPublicGateway); ok {
		gw = pgw.(string)
	}

	if userDetails.generation == 1 {
		err := classicSubnetCreate(d, meta, name, vpc, zone, ipv4cidr, acl, gw, ipv4addrcount64)
		if err != nil {
			return err
		}
	} else {
		err := subnetCreate(d, meta, name, vpc, zone, ipv4cidr, acl, gw, ipv4addrcount64)
		if err != nil {
			return err
		}
	}

	return resourceIBMISSubnetRead(d, meta)
}

func classicSubnetCreate(d *schema.ResourceData, meta interface{}, name, vpc, zone, ipv4cidr, acl, gw string, ipv4addrcount64 int64) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	subnetTemplate := &vpcclassicv1.SubnetPrototype{
		Name: &name,
		Vpc: &vpcclassicv1.VPCIdentity{
			ID: &vpc,
		},
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: &zone,
		},
	}
	if ipv4cidr != "" {
		subnetTemplate.Ipv4CidrBlock = &ipv4cidr
	}
	if ipv4addrcount64 != int64(0) {
		subnetTemplate.TotalIpv4AddressCount = &ipv4addrcount64
	}
	if gw != "" {
		subnetTemplate.PublicGateway = &vpcclassicv1.PublicGatewayIdentity{
			ID: &gw,
		}
	}

	if acl != "" {
		subnetTemplate.NetworkAcl = &vpcclassicv1.NetworkACLIdentity{
			ID: &acl,
		}
	}
	//create a subnet
	createSubnetOptions := &vpcclassicv1.CreateSubnetOptions{
		SubnetPrototype: subnetTemplate,
	}
	subnet, response, err := sess.CreateSubnet(createSubnetOptions)
	if err != nil {
		log.Printf("[DEBUG] Subnet err %s\n%s", err, response)
		return fmt.Errorf("Error while creating Subnet %s\n%s", err, response)
	}
	d.SetId(*subnet.ID)
	log.Printf("[INFO] Subnet : %s", *subnet.ID)
	_, err = isWaitForClassicSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func subnetCreate(d *schema.ResourceData, meta interface{}, name, vpc, zone, ipv4cidr, acl, gw string, ipv4addrcount64 int64) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	subnetTemplate := &vpcv1.SubnetPrototype{
		Name: &name,
		Vpc: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	}
	if ipv4cidr != "" {
		subnetTemplate.Ipv4CidrBlock = &ipv4cidr
	}
	if ipv4addrcount64 != int64(0) {
		subnetTemplate.TotalIpv4AddressCount = &ipv4addrcount64
	}
	if gw != "" {
		subnetTemplate.PublicGateway = &vpcv1.PublicGatewayIdentity{
			ID: &gw,
		}
	}

	if acl != "" {
		subnetTemplate.NetworkAcl = &vpcv1.NetworkACLIdentity{
			ID: &acl,
		}
	}
	rg := ""
	if grp, ok := d.GetOk(isSubnetResourceGroup); ok {
		rg = grp.(string)
		subnetTemplate.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	//create a subnet
	createSubnetOptions := &vpcv1.CreateSubnetOptions{
		SubnetPrototype: subnetTemplate,
	}
	subnet, response, err := sess.CreateSubnet(createSubnetOptions)
	if err != nil {
		log.Printf("[DEBUG] Subnet err %s\n%s", err, response)
		return fmt.Errorf("Error while creating Subnet %s\n%s", err, response)
	}
	d.SetId(*subnet.ID)
	log.Printf("[INFO] Subnet : %s", *subnet.ID)
	_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func isWaitForClassicSubnetAvailable(subnetC *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetProvisioning},
		Target:     []string{isSubnetProvisioningDone, ""},
		Refresh:    isClassicSubnetRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicSubnetRefreshFunc(subnetC *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error getting Subnet : %s\n%s", err, response)
		}

		if *subnet.Status == "available" || *subnet.Status == "failed" {
			return subnet, isSubnetProvisioningDone, nil
		}

		return subnet, isSubnetProvisioning, nil
	}
}

func isWaitForSubnetAvailable(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetProvisioning},
		Target:     []string{isSubnetProvisioningDone, ""},
		Refresh:    isSubnetRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error getting Subnet : %s\n%s", err, response)
		}

		if *subnet.Status == "available" || *subnet.Status == "failed" {
			return subnet, isSubnetProvisioningDone, nil
		}

		return subnet, isSubnetProvisioning, nil
	}
}

func resourceIBMISSubnetRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicSubnetGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := subnetGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}
func classicSubnetGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	// d.Set(isSubnetIPVersion, *subnet..IPVersion)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CidrBlock)
	// d.Set(isSubnetIpv6CidrBlock, *subnet.IPV6CidrBlock)
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

func subnetGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	d.Set("id", *subnet.ID)
	d.Set(isSubnetName, *subnet.Name)
	// d.Set(isSubnetIPVersion, *subnet..IPVersion)
	d.Set(isSubnetIpv4CidrBlock, *subnet.Ipv4CidrBlock)
	// d.Set(isSubnetIpv6CidrBlock, *subnet.IPV6CidrBlock)
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

func resourceIBMISSubnetUpdate(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()

	if userDetails.generation == 1 {
		err := classicSubnetUpdate(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := subnetUpdate(d, meta, id)
		if err != nil {
			return err
		}
	}
	return resourceIBMISSubnetRead(d, meta)
}

func classicSubnetUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	hasChanged := false
	name := ""
	acl := ""
	updateSubnetOptions := &vpcclassicv1.UpdateSubnetOptions{}
	if d.HasChange(isSubnetName) {
		name = d.Get(isSubnetName).(string)
		updateSubnetOptions.Name = &name
		hasChanged = true
	}
	if d.HasChange(isSubnetNetworkACL) {
		acl = d.Get(isSubnetNetworkACL).(string)
		updateSubnetOptions.NetworkAcl = &vpcclassicv1.NetworkACLIdentity{
			ID: &acl,
		}
		hasChanged = true
	}
	if d.HasChange(isSubnetPublicGateway) {
		gw := d.Get(isSubnetPublicGateway).(string)
		if gw == "" {
			deleteSubnetPublicGatewayBindingOptions := &vpcclassicv1.DeleteSubnetPublicGatewayBindingOptions{
				ID: &id,
			}
			response, err := sess.DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptions)
			if err != nil {
				return fmt.Errorf("Error Detaching the public gateway attached to the subnet : %s\n%s", err, response)
			}
			_, err = isWaitForClassicSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		} else {
			setSubnetPublicGatewayBindingOptions := &vpcclassicv1.SetSubnetPublicGatewayBindingOptions{
				ID: &id,
				PublicGatewayIdentity: &vpcclassicv1.PublicGatewayIdentity{
					ID: &gw,
				},
			}
			_, response, err := sess.SetSubnetPublicGatewayBinding(setSubnetPublicGatewayBindingOptions)
			if err != nil {
				return fmt.Errorf("Error Attaching public gateway to the subnet : %s\n%s", err, response)
			}
			_, err = isWaitForClassicSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		}
	}
	if hasChanged {
		updateSubnetOptions.ID = &id
		_, response, err := sess.UpdateSubnet(updateSubnetOptions)
		if err != nil {
			return fmt.Errorf("Error Updating Subnet : %s\n%s", err, response)
		}
	}
	return nil
}

func subnetUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	hasChanged := false
	name := ""
	acl := ""
	updateSubnetOptions := &vpcv1.UpdateSubnetOptions{}
	if d.HasChange(isSubnetName) {
		name = d.Get(isSubnetName).(string)
		updateSubnetOptions.Name = &name
		hasChanged = true
	}
	if d.HasChange(isSubnetNetworkACL) {
		acl = d.Get(isSubnetNetworkACL).(string)
		updateSubnetOptions.NetworkAcl = &vpcv1.NetworkACLIdentity{
			ID: &acl,
		}
		hasChanged = true
	}
	if d.HasChange(isSubnetPublicGateway) {
		gw := d.Get(isSubnetPublicGateway).(string)
		if gw == "" {
			deleteSubnetPublicGatewayBindingOptions := &vpcv1.DeleteSubnetPublicGatewayBindingOptions{
				ID: &id,
			}
			response, err := sess.DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptions)
			if err != nil {
				return fmt.Errorf("Error Detaching the public gateway attached to the subnet : %s\n%s", err, response)
			}
			_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		} else {
			setSubnetPublicGatewayBindingOptions := &vpcv1.SetSubnetPublicGatewayBindingOptions{
				ID: &id,
				PublicGatewayIdentity: &vpcv1.PublicGatewayIdentity{
					ID: &gw,
				},
			}
			_, response, err := sess.SetSubnetPublicGatewayBinding(setSubnetPublicGatewayBindingOptions)
			if err != nil {
				return fmt.Errorf("Error Attaching public gateway to the subnet : %s\n%s", err, response)
			}
			_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		}
	}
	if hasChanged {
		updateSubnetOptions.ID = &id
		_, response, err := sess.UpdateSubnet(updateSubnetOptions)
		if err != nil {
			return fmt.Errorf("Error Updating Subnet : %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicSubnetDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := subnetDelete(d, meta, id)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil
}

func classicSubnetDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	if subnet.PublicGateway != nil {
		deleteSubnetPublicGatewayBindingOptions := &vpcclassicv1.DeleteSubnetPublicGatewayBindingOptions{
			ID: &id,
		}
		_, err = sess.DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptions)
		if err != nil {
			return err
		}
		_, err = isWaitForClassicSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return err
		}
	}
	deleteSubnetOptions := &vpcclassicv1.DeleteSubnetOptions{
		ID: &id,
	}
	response, err = sess.DeleteSubnet(deleteSubnetOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting Subnet : %s\n%s", err, response)
	}
	_, err = isWaitForClassicSubnetDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func subnetDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnet(getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error Getting Subnet (%s): %s\n%s", id, err, response)
	}
	if subnet.PublicGateway != nil {
		deleteSubnetPublicGatewayBindingOptions := &vpcv1.DeleteSubnetPublicGatewayBindingOptions{
			ID: &id,
		}
		_, err = sess.DeleteSubnetPublicGatewayBinding(deleteSubnetPublicGatewayBindingOptions)
		if err != nil {
			return err
		}
		_, err = isWaitForSubnetAvailable(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return err
		}
	}
	deleteSubnetOptions := &vpcv1.DeleteSubnetOptions{
		ID: &id,
	}
	response, err = sess.DeleteSubnet(deleteSubnetOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting Subnet : %s\n%s", err, response)
	}
	_, err = isWaitForSubnetDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForClassicSubnetDeleted(subnetC *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetDeleting},
		Target:     []string{isSubnetDeleted, ""},
		Refresh:    isClassicSubnetDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicSubnetDeleteRefreshFunc(subnetC *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getSubnetOptions := &vpcclassicv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return subnet, isSubnetDeleted, nil
			}
			return subnet, "", fmt.Errorf("The Subnet %s failed to delete: %s\n%s", id, err, response)
		}
		return subnet, isSubnetDeleting, err
	}
}

func isWaitForSubnetDeleted(subnetC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for subnet (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isSubnetDeleting},
		Target:     []string{isSubnetDeleted, ""},
		Refresh:    isSubnetDeleteRefreshFunc(subnetC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSubnetDeleteRefreshFunc(subnetC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getSubnetOptions := &vpcv1.GetSubnetOptions{
			ID: &id,
		}
		subnet, response, err := subnetC.GetSubnet(getSubnetOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return subnet, isSubnetDeleted, nil
			}
			return subnet, "", fmt.Errorf("The Subnet %s failed to delete: %s\n%s", id, err, response)
		}
		return subnet, isSubnetDeleting, err
	}
}

func resourceIBMISSubnetExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		exists, err := classicSubnetExists(d, meta, id)
		return exists, err
	} else {
		exists, err := subnetExists(d, meta, id)
		return exists, err
	}
}

func classicSubnetExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return false, err
	}
	getsubnetOptions := &vpcclassicv1.GetSubnetOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnet(getsubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Subnet: %s\n%s", err, response)
	}
	return true, nil
}

func subnetExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getsubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnet(getsubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Subnet: %s\n%s", err, response)
	}
	return true, nil
}
