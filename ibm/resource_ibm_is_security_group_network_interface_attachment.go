package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
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
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgClient := network.NewSecurityGroupClient(sess)

	sgID := d.Get(isSGNICAGroupId).(string)
	nicID := d.Get(isSGNICANicId).(string)
	_, err = sgClient.AddNetworkInterface(sgID, nicID)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", sgID, nicID))
	return resourceIBMISSecurityGroupNetworkInterfaceAttachmentRead(d, meta)

}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgClient := network.NewSecurityGroupClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	sgID := parts[0]
	nicID := parts[1]
	instanceNic, err := sgClient.GetNetworkInterface(sgID, nicID)
	if err != nil {
		return err
	}

	d.Set(isSGNICAGroupId, sgID)
	d.Set(isSGNICANicId, nicID)
	d.Set(isSGNICAInstanceNwInterfaceID, instanceNic.ID.String())
	d.Set(isSGNICAName, instanceNic.Name)
	d.Set(isSGNICAPortSpeed, instanceNic.PortSpeed)
	d.Set(isSGNICAPrimaryIPV4Address, instanceNic.PrimaryIPV4Address)
	d.Set(isSGNICAStatus, instanceNic.Status)
	d.Set(isSGNICAType, instanceNic.Type)
	if instanceNic.Subnet != nil {
		d.Set(isSGNICASubnet, instanceNic.Subnet.ID.String())
	}
	sgs := make([]map[string]interface{}, len(instanceNic.SecurityGroups))
	for i, sgObj := range instanceNic.SecurityGroups {
		sg := make(map[string]interface{})
		sg[isSGNICASecurityGroupCRN] = sgObj.Crn
		sg[isSGNICASecurityGroupID] = sgObj.ID.String()
		sg[isSGNICASecurityGroupName] = sgObj.Name
		sgs[i] = sg
	}
	d.Set(isSGNICASecurityGroups, sgs)

	fps := make([]map[string]interface{}, len(instanceNic.FloatingIps))
	for i, fpObj := range instanceNic.FloatingIps {
		fp := make(map[string]interface{})
		fp[isSGNICAFloatingIpCRN] = fpObj.Crn
		fp[isSGNICAFloatingIpID] = fpObj.ID.String()
		fp[isSGNICAFloatingIpName] = fpObj.Name
		fps[i] = fp
	}
	d.Set(isSGNICAFloatingIps, fps)

	d.Set(isSGNICASecondaryAddresses, instanceNic.SecondaryAddresses)
	return nil
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgClient := network.NewSecurityGroupClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	sgID := parts[0]
	nicID := parts[1]
	err = sgClient.DeleteNetworkInterface(sgID, nicID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupNetworkInterfaceAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	sgClient := network.NewSecurityGroupClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	sgID := parts[0]
	nicID := parts[1]

	_, err = sgClient.GetNetworkInterface(sgID, nicID)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
