package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isSGNICAGroupId               = "security_group"
	isSGNICANicId                 = "network_interface"
	isSGNICAInstanceNwInterfaceID = "instance_network_interface"
	isSGNICAName                  = "name"
	isSGNICAPortSpeed             = "port_speed"
	isSGNICAPrimaryIPV4Address    = "primary_ipv4_address"
	isSGNICAPrimaryIPV6Address    = "primary_ipv6_address"
	isSGNICASecondaryAddresses    = "secondary_address"
	isSGNICASecurityGroups        = "security_groups"
	isSGNICASecurityGroupCRN      = "crn"
	isSGNICASecurityGroupID       = "id"
	isSGNICASecurityGroupName     = "name"
	isSGNICAStatus                = "status"
	isSGNICASubnet                = "subnet"
	isSGNICAType                  = "type"
	isSGNICAFloatingIps           = "floating_ips"
	isSGNICAFloatingIpID          = "id"
	isSGNICAFloatingIpAddress     = "address"
	isSGNICAFloatingIpName        = "name"
	isSGNICAFloatingIpCRN         = "crn"
)

func resourceIBMISSecurityGroupNetworkInterfaceAttachment() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSecurityGroupNetworkInterfaceAttachmentCreate,
		Read:     resourceIBMISSecurityGroupNetworkInterfaceAttachmentRead,
		Delete:   resourceIBMISSecurityGroupNetworkInterfaceAttachmentDelete,
		Exists:   resourceIBMISSecurityGroupNetworkInterfaceAttachmentExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isSGNICAGroupId: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "security group network interface attachment group ID",
			},
			isSGNICANicId: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "security group network interface attachment NIC ID",
			},
			isSGNICAName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment name",
			},
			isSGNICAInstanceNwInterfaceID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment network interface ID",
			},
			isSGNICAPortSpeed: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "security group network interface attachment port speed",
			},
			isSGNICAPrimaryIPV4Address: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment Primary IPV4 address",
			},
			isSGNICAPrimaryIPV6Address: {
				Type:     schema.TypeString,
				Computed: true,
				Removed:  "This field is removed",
			},
			isSGNICASecondaryAddresses: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "security group network interface attachment secondary address",
			},
			isSGNICAStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment status",
			},
			isSGNICASubnet: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment subnet",
			},
			isSGNICAType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group network interface attachment type",
			},
			isSGNICAFloatingIps: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSGNICAFloatingIpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment floating IP ID",
						},
						isSGNICAFloatingIpAddress: {
							Type:     schema.TypeString,
							Computed: true,
							Removed:  "This field is removed",
						},
						isSGNICAFloatingIpName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment floating IP name",
						},
						isSGNICAFloatingIpCRN: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment floating IP CRN",
						},
					},
				},
			},
			isSGNICASecurityGroups: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSGNICASecurityGroupID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment security group ID",
						},
						isSGNICASecurityGroupCRN: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment security group CRN",
						},
						isSGNICASecurityGroupName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group network interface attachment security group name",
						},
					},
				},
			},
		},
	}
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	sgID := d.Get(isSGNICAGroupId).(string)
	nicID := d.Get(isSGNICANicId).(string)

	if userDetails.generation == 1 {
		err := classicSgnicCreate(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	} else {
		err := sgnicCreate(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	}
	return resourceIBMISSecurityGroupNetworkInterfaceAttachmentRead(d, meta)

}

func classicSgnicCreate(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateSecurityGroupNetworkInterfaceBindingOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.CreateSecurityGroupNetworkInterfaceBinding(options)
	if err != nil {
		return fmt.Errorf("Error while creating SecurityGroup NetworkInterface Binding %s\n%s", err, response)
	}
	d.SetId(fmt.Sprintf("%s/%s", sgID, nicID))
	return nil
}

func sgnicCreate(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateSecurityGroupNetworkInterfaceBindingOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.CreateSecurityGroupNetworkInterfaceBinding(options)
	if err != nil {
		return fmt.Errorf("Error while creating SecurityGroup NetworkInterface Binding %s\n%s", err, response)
	}
	d.SetId(fmt.Sprintf("%s/%s", sgID, nicID))
	return nil
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	sgID := parts[0]
	nicID := parts[1]
	if userDetails.generation == 1 {
		err := classicSgnicGet(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	} else {
		err := sgnicGet(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicSgnicGet(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getSecurityGroupNetworkInterfaceOptions := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	instanceNic, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isSGNICAGroupId, sgID)
	d.Set(isSGNICANicId, nicID)
	d.Set(isSGNICAInstanceNwInterfaceID, *instanceNic.ID)
	d.Set(isSGNICAName, *instanceNic.Name)
	d.Set(isSGNICAPortSpeed, *instanceNic.PortSpeed)
	d.Set(isSGNICAPrimaryIPV4Address, *instanceNic.PrimaryIpv4Address)
	// d.Set(isSGNICAStatus, *instanceNic.Status)
	d.Set(isSGNICAType, *instanceNic.Type)
	if instanceNic.Subnet != nil {
		d.Set(isSGNICASubnet, *instanceNic.Subnet.ID)
	}
	sgs := make([]map[string]interface{}, len(instanceNic.SecurityGroups))
	for i, sgObj := range instanceNic.SecurityGroups {
		sg := make(map[string]interface{})
		sg[isSGNICASecurityGroupCRN] = *sgObj.Crn
		sg[isSGNICASecurityGroupID] = *sgObj.ID
		sg[isSGNICASecurityGroupName] = *sgObj.Name
		sgs[i] = sg
	}
	d.Set(isSGNICASecurityGroups, sgs)

	fps := make([]map[string]interface{}, len(instanceNic.FloatingIps))
	for i, fpObj := range instanceNic.FloatingIps {
		fp := make(map[string]interface{})
		fp[isSGNICAFloatingIpCRN] = fpObj.Crn
		fp[isSGNICAFloatingIpID] = *fpObj.ID
		fp[isSGNICAFloatingIpName] = *fpObj.Name
		fps[i] = fp
	}
	d.Set(isSGNICAFloatingIps, fps)

	// d.Set(isSGNICASecondaryAddresses, *instanceNic.SecondaryAddresses)
	return nil
}

func sgnicGet(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getSecurityGroupNetworkInterfaceOptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	instanceNic, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isSGNICAGroupId, sgID)
	d.Set(isSGNICANicId, nicID)
	d.Set(isSGNICAInstanceNwInterfaceID, *instanceNic.ID)
	d.Set(isSGNICAName, *instanceNic.Name)
	d.Set(isSGNICAPortSpeed, *instanceNic.PortSpeed)
	d.Set(isSGNICAPrimaryIPV4Address, *instanceNic.PrimaryIpv4Address)
	d.Set(isSGNICAStatus, *instanceNic.Status)
	d.Set(isSGNICAType, *instanceNic.Type)
	if instanceNic.Subnet != nil {
		d.Set(isSGNICASubnet, *instanceNic.Subnet.ID)
	}
	sgs := make([]map[string]interface{}, len(instanceNic.SecurityGroups))
	for i, sgObj := range instanceNic.SecurityGroups {
		sg := make(map[string]interface{})
		sg[isSGNICASecurityGroupCRN] = *sgObj.Crn
		sg[isSGNICASecurityGroupID] = *sgObj.ID
		sg[isSGNICASecurityGroupName] = *sgObj.Name
		sgs[i] = sg
	}
	d.Set(isSGNICASecurityGroups, sgs)

	fps := make([]map[string]interface{}, len(instanceNic.FloatingIps))
	for i, fpObj := range instanceNic.FloatingIps {
		fp := make(map[string]interface{})
		fp[isSGNICAFloatingIpCRN] = fpObj.Crn
		fp[isSGNICAFloatingIpID] = *fpObj.ID
		fp[isSGNICAFloatingIpName] = *fpObj.Name
		fps[i] = fp
	}
	d.Set(isSGNICAFloatingIps, fps)

	// d.Set(isSGNICASecondaryAddresses, *instanceNic.SecondaryAddresses)
	return nil
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	sgID := parts[0]
	nicID := parts[1]
	if userDetails.generation == 1 {
		err := classicSgnicDelete(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	} else {
		err := sgnicDelete(d, meta, sgID, nicID)
		if err != nil {
			return err
		}
	}
	d.SetId("")
	return nil
}

func classicSgnicDelete(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getSecurityGroupNetworkInterfaceOptions := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}

	deleteSecurityGroupNetworkInterfaceBindingOptions := &vpcclassicv1.DeleteSecurityGroupNetworkInterfaceBindingOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	response, err = sess.DeleteSecurityGroupNetworkInterfaceBinding(deleteSecurityGroupNetworkInterfaceBindingOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	d.SetId("")
	return nil
}

func sgnicDelete(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getSecurityGroupNetworkInterfaceOptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}

	deleteSecurityGroupNetworkInterfaceBindingOptions := &vpcv1.DeleteSecurityGroupNetworkInterfaceBindingOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	response, err = sess.DeleteSecurityGroupNetworkInterfaceBinding(deleteSecurityGroupNetworkInterfaceBindingOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	sgID := parts[0]
	nicID := parts[1]
	if userDetails.generation == 1 {
		err := classicSgnicExists(d, meta, sgID, nicID)
		if err != nil {
			return false, err
		}
	} else {
		err := sgnicExists(d, meta, sgID, nicID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicSgnicExists(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getSecurityGroupNetworkInterfaceOptions := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func sgnicExists(d *schema.ResourceData, meta interface{}, sgID, nicID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getSecurityGroupNetworkInterfaceOptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
		SecurityGroupID: &sgID,
		ID:              &nicID,
	}
	_, response, err := sess.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting NetworkInterface(%s) for the SecurityGroup (%s) : %s\n%s", nicID, sgID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}
