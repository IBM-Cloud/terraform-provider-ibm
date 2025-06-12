// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isNetworkInterfacePending   = "pending"
	isNetworkInterfaceAvailable = "available"
	isNetworkInterfaceFailed    = "failed"
	isNetworkInterfaceDeleting  = "deleting"
	isNetworkInterfaceDeleted   = "deleted"
)

func ResourceIBMIsInstanceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsInstanceNetworkInterfaceCreate,
		ReadContext:   resourceIBMIsInstanceNetworkInterfaceRead,
		UpdateContext: resourceIBMIsInstanceNetworkInterfaceUpdate,
		DeleteContext: resourceIBMIsInstanceNetworkInterfaceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier of the instance.",
			},
			isInstanceNicSubnet: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier of the subnet.",
			},
			isInstanceNicAllowIPSpoofing: &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.",
			},
			isInstanceNicName: &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_network_interface", isInstanceNicName),
				Description:  "The user-defined name for this network interface. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			isInstanceNicPrimaryIpv4Address: &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"primary_ip.0.address"},
				Deprecated:    "primary_ipv4_address is deprecated and support will be removed. Use primary_ip instead",
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_network_interface", isInstanceNicPrimaryIpv4Address),
				Description:   "The primary IPv4 address. If specified, it must be an available address on the network interface's subnet. If unspecified, an available address on the subnet will be automatically selected.",
			},
			isInstanceNicPrimaryIP: {
				Type:        schema.TypeList,
				MinItems:    0,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceNicReservedIpAddress: {
							Type:          schema.TypeString,
							Computed:      true,
							ForceNew:      true,
							Optional:      true,
							ConflictsWith: []string{"primary_ipv4_address"},
							Description:   "The IP address to reserve, which must not already be reserved on the subnet.",
						},
						isInstanceNicReservedIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isInstanceNicReservedIpAutoDelete: {
							Type:             schema.TypeBool,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: flex.ApplyOnce,
							Description:      "Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.",
						},
						isInstanceNicReservedIpName: {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: flex.ApplyOnce,
							Description:      "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
						},
						isInstanceNicReservedIpId: {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"primary_ipv4_address"},
							Computed:      true,
							Description:   "Identifies a reserved IP by a unique property.",
						},
						isInstanceNicReservedIpResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type",
						},
					},
				},
			},
			"network_interface": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The globally unique ID of this network interface",
			},
			isInstanceNicSecurityGroups: {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the network interface was created.",
			},
			isInstanceNicFloatingIP: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the floating IP to attach to this network interface",
			},
			isInstanceNicFloatingIPs: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The floating IPs associated with this network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique IP address.",
						},
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this floating IP.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this floating IP.",
						},
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique identifier for this floating IP.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this floating IP.",
						},
					},
				},
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this network interface.",
			},
			"port_speed": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The network interface port speed in Mbps.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the network interface.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this network interface as it relates to an instance.",
			},
		},
	}
}

func ResourceIBMIsInstanceNetworkInterfaceValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceNicName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 isInstanceNicPrimaryIpv4Address,
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_network_interface", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsInstanceNetworkInterfaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instance_id := d.Get("instance").(string)

	createInstanceNetworkInterfaceOptions := &vpcv1.CreateInstanceNetworkInterfaceOptions{}

	createInstanceNetworkInterfaceOptions.SetInstanceID(instance_id)

	subnetId := d.Get(isInstanceNicSubnet).(string)
	subnetIdentity := vpcv1.SubnetIdentity{
		ID: &subnetId,
	}
	createInstanceNetworkInterfaceOptions.SetSubnet(&subnetIdentity)
	if allow_ip_spoofing, ok := d.GetOk(isInstanceNicAllowIPSpoofing); ok {
		createInstanceNetworkInterfaceOptions.SetAllowIPSpoofing(allow_ip_spoofing.(bool))
	}
	if name, ok := d.GetOk(isInstanceNicName); ok {
		createInstanceNetworkInterfaceOptions.SetName(name.(string))
	}

	var primary_ipv4, reservedIp, reservedipv4, reservedipname string
	var autodelete, okAuto bool
	if primary_ipv4Ok, ok := d.GetOk(isInstanceNicPrimaryIpv4Address); ok {
		primary_ipv4 = primary_ipv4Ok.(string)
	}

	//reserved ip changes
	if primaryIpOk, ok := d.GetOk(isInstanceNicPrimaryIP); ok {
		primip := primaryIpOk.([]interface{})[0].(map[string]interface{})

		reservedipok, _ := primip[isInstanceNicReservedIpId]
		reservedIp = reservedipok.(string)
		reservedipv4Ok, _ := primip[isInstanceNicReservedIpAddress]
		reservedipv4 = reservedipv4Ok.(string)

		reservedipnameOk, _ := primip[isInstanceNicReservedIpName]
		reservedipname = reservedipnameOk.(string)

		reservedipautodeleteok, okAuto := primip[isInstanceNicReservedIpAutoDelete]
		if okAuto {
			autodelete = reservedipautodeleteok.(bool)
		}
	}

	if primary_ipv4 != "" && reservedipv4 != "" && primary_ipv4 != reservedipv4 {
		err = fmt.Errorf("Error creating instance, network_interfaces error, use either primary_ipv4_address(%s) or primary_ip.0.address(%s)", primary_ipv4, reservedipv4)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceNetworkInterface failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if reservedIp != "" && (primary_ipv4 != "" || reservedipv4 != "" || reservedipname != "") {
		err = fmt.Errorf("Error creating instance, network_interfaces error, reserved_ip(%s) is mutually exclusive with other primary_ip attributes", reservedIp)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceNetworkInterface failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if reservedIp != "" {
		createInstanceNetworkInterfaceOptions.PrimaryIP = &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
			ID: &reservedIp,
		}
	} else {
		if primary_ipv4 != "" || reservedipv4 != "" || reservedipname != "" || okAuto {
			primaryipobj := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{}
			if primary_ipv4 != "" {
				primaryipobj.Address = &primary_ipv4
			}
			if reservedipv4 != "" {
				primaryipobj.Address = &reservedipv4
			}
			if reservedipname != "" {
				primaryipobj.Name = &reservedipname
			}
			if okAuto {
				primaryipobj.AutoDelete = &autodelete
			}
			createInstanceNetworkInterfaceOptions.PrimaryIP = primaryipobj
		}
	}

	if secgrpintf, ok := d.GetOk(isInstanceNicSecurityGroups); ok {
		secgrpSet := secgrpintf.(*schema.Set)
		if secgrpSet.Len() != 0 {
			var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
			for i, secgrpIntf := range secgrpSet.List() {
				secgrpIntfstr := secgrpIntf.(string)
				secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
					ID: &secgrpIntfstr,
				}
			}
			createInstanceNetworkInterfaceOptions.SecurityGroups = secgrpobjs
		}
	}

	isNICKey := "instance_key_" + instance_id
	conns.IbmMutexKV.Lock(isNICKey)
	defer conns.IbmMutexKV.Unlock(isNICKey)

	networkInterface, _, err := vpcClient.CreateInstanceNetworkInterfaceWithContext(context, createInstanceNetworkInterfaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createInstanceNetworkInterfaceOptions.InstanceID, *networkInterface.ID))
	d.Set("network_interface", *networkInterface.ID)

	_, err = isWaitForNetworkInterfaceAvailable(vpcClient, d.Id(), d.Timeout(schema.TimeoutUpdate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if floating_ip_Intf, ok := d.GetOk(isInstanceNicFloatingIP); ok && floating_ip_Intf.(string) != "" {
		floating_ip_id := floating_ip_Intf.(string)

		addInstanceNetworkInterfaceFloatingIPOptions := &vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions{
			InstanceID:         createInstanceNetworkInterfaceOptions.InstanceID,
			NetworkInterfaceID: networkInterface.ID,
			ID:                 &floating_ip_id,
		}

		_, _, err := vpcClient.AddInstanceNetworkInterfaceFloatingIP(addInstanceNetworkInterfaceFloatingIPOptions)

		if err != nil {
			d.Set(isInstanceNicFloatingIP, "")
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddInstanceNetworkInterfaceFloatingIP failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForNetworkInterfaceAvailable(vpcClient, d.Id(), d.Timeout(schema.TimeoutUpdate), d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}

	_, err = isWaitForInstanceAvailable(vpcClient, instance_id, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return resourceIBMIsInstanceNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsInstanceNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "sep-id-parts").GetDiag()
	}

	getInstanceNetworkInterfaceOptions.SetInstanceID(parts[0])
	getInstanceNetworkInterfaceOptions.SetID(parts[1])

	networkInterface, response, err := vpcClient.GetInstanceNetworkInterfaceWithContext(context, getInstanceNetworkInterfaceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(fmt.Sprintf("%s/%s", parts[0], *networkInterface.ID))
	if err = d.Set("network_interface", *networkInterface.ID); err != nil {
		err = fmt.Errorf("Error setting network_interface: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-network_interface").GetDiag()
	}

	if err = d.Set("instance", parts[0]); err != nil {
		err = fmt.Errorf("Error setting instance: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-instance").GetDiag()
	}
	if err = d.Set(isInstanceNicSubnet, *networkInterface.Subnet.ID); err != nil {
		err = fmt.Errorf("Error setting subnet: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-subnet").GetDiag()
	}
	if err = d.Set(isInstanceNicAllowIPSpoofing, *networkInterface.AllowIPSpoofing); err != nil {
		err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
	}
	if err = d.Set(isInstanceNicName, *networkInterface.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-name").GetDiag()
	}
	if networkInterface.PrimaryIP != nil {
		if err = d.Set(isInstanceNicPrimaryIpv4Address, *networkInterface.PrimaryIP.Address); err != nil {
			err = fmt.Errorf("Error setting primary_ipv4_address: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-primary_ipv4_address").GetDiag()
		}
		// reserved ip changes
		primaryIpList := make([]map[string]interface{}, 0)
		currentPrimIp := map[string]interface{}{}

		if networkInterface.PrimaryIP.Address != nil {
			currentPrimIp[isInstanceNicReservedIpAddress] = *networkInterface.PrimaryIP.Address
		}
		if networkInterface.PrimaryIP.Href != nil {
			currentPrimIp[isInstanceNicReservedIpHref] = *networkInterface.PrimaryIP.Href
		}
		if networkInterface.PrimaryIP.Name != nil {
			currentPrimIp[isInstanceNicReservedIpName] = *networkInterface.PrimaryIP.Name
		}
		if networkInterface.PrimaryIP.ID != nil {
			currentPrimIp[isInstanceNicReservedIpId] = *networkInterface.PrimaryIP.ID
		}
		if networkInterface.PrimaryIP.ResourceType != nil {
			currentPrimIp[isInstanceNicReservedIpResourceType] = *networkInterface.PrimaryIP.ResourceType
		}
		getripoptions := &vpcv1.GetSubnetReservedIPOptions{
			SubnetID: networkInterface.Subnet.ID,
			ID:       networkInterface.PrimaryIP.ID,
		}
		insRip, _, err := vpcClient.GetSubnetReservedIPWithContext(context, getripoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		currentPrimIp[isInstanceNicReservedIpAutoDelete] = insRip.AutoDelete

		primaryIpList = append(primaryIpList, currentPrimIp)
		if err = d.Set("primary_ip", primaryIpList); err != nil {
			err = fmt.Errorf("Error setting primary_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-primary_ip").GetDiag()
		}
	}
	if networkInterface.SecurityGroups != nil && len(networkInterface.SecurityGroups) != 0 {
		secgrpList := []string{}
		for _, secGrp := range networkInterface.SecurityGroups {
			secgrpList = append(secgrpList, string(*(secGrp.ID)))
		}
		if err = d.Set("security_groups", flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
			err = fmt.Errorf("Error setting security_groups: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-security_groups").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(networkInterface.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-created_at").GetDiag()
	}
	floatingIps := []map[string]interface{}{}
	if networkInterface.FloatingIps != nil {

		for _, floatingIpsItem := range networkInterface.FloatingIps {
			floatingIpsItemMap := resourceIBMIsInstanceNetworkInterfaceFloatingIPReferenceToMap(floatingIpsItem)
			floatingIps = append(floatingIps, floatingIpsItemMap)
		}
	}
	if err = d.Set(isInstanceNicFloatingIPs, floatingIps); err != nil {
		err = fmt.Errorf("Error setting floating_ips: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-floating_ips").GetDiag()
	}
	if err = d.Set("href", networkInterface.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-href").GetDiag()
	}
	if err = d.Set("port_speed", flex.IntValue(networkInterface.PortSpeed)); err != nil {
		err = fmt.Errorf("Error setting port_speed: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-port_speed").GetDiag()
	}
	if err = d.Set("resource_type", networkInterface.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("status", networkInterface.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-status").GetDiag()
	}
	if err = d.Set("type", networkInterface.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "read", "set-type").GetDiag()
	}

	return nil
}

func resourceIBMIsInstanceNetworkInterfaceFloatingIPReferenceToMap(floatingIPReference vpcv1.FloatingIPReference) map[string]interface{} {
	floatingIPReferenceMap := map[string]interface{}{}

	floatingIPReferenceMap["address"] = floatingIPReference.Address
	floatingIPReferenceMap["crn"] = floatingIPReference.CRN
	if floatingIPReference.Deleted != nil {
		DeletedMap := resourceIBMIsInstanceNetworkInterfaceFloatingIPReferenceDeletedToMap(*floatingIPReference.Deleted)
		floatingIPReferenceMap["deleted"] = []map[string]interface{}{DeletedMap}
	}
	floatingIPReferenceMap["href"] = floatingIPReference.Href
	floatingIPReferenceMap["id"] = floatingIPReference.ID
	floatingIPReferenceMap["name"] = floatingIPReference.Name

	return floatingIPReferenceMap
}

func resourceIBMIsInstanceNetworkInterfaceFloatingIPReferenceDeletedToMap(floatingIPReferenceDeleted vpcv1.Deleted) map[string]interface{} {
	floatingIPReferenceDeletedMap := map[string]interface{}{}

	floatingIPReferenceDeletedMap["more_info"] = floatingIPReferenceDeleted.MoreInfo

	return floatingIPReferenceDeletedMap
}

func resourceIBMIsInstanceNetworkInterfaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceNetworkInterfaceOptions := &vpcv1.UpdateInstanceNetworkInterfaceOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "update", "sep-id-parts").GetDiag()
	}
	instance_id := parts[0]
	network_interface_id := parts[1]
	updateInstanceNetworkInterfaceOptions.SetInstanceID(instance_id)
	updateInstanceNetworkInterfaceOptions.SetID(network_interface_id)

	hasChange := false

	patchVals := &vpcv1.NetworkInterfacePatch{}

	if d.HasChange(isInstanceNicAllowIPSpoofing) {
		patchVals.AllowIPSpoofing = core.BoolPtr(d.Get(isInstanceNicAllowIPSpoofing).(bool))
		hasChange = true
	}
	if d.HasChange(isInstanceNicName) {
		patchVals.Name = core.StringPtr(d.Get(isInstanceNicName).(string))
		hasChange = true
	}
	if !d.IsNewResource() && (d.HasChange("primary_network_interface.0.primary_ip.0.name") || d.HasChange("primary_network_interface.0.primary_ip.0.auto_delete")) {
		subnetId := d.Get(isBareMetalServerNicSubnet).(string)
		ripId := d.Get("primary_network_interface.0.primary_ip.0.reserved_ip").(string)
		updateripoptions := &vpcv1.UpdateSubnetReservedIPOptions{
			SubnetID: &subnetId,
			ID:       &ripId,
		}
		reservedIpPath := &vpcv1.ReservedIPPatch{}
		if d.HasChange("primary_network_interface.0.primary_ip.0.name") {
			name := d.Get("primary_network_interface.0.primary_ip.0.name").(string)
			reservedIpPath.Name = &name
		}
		if d.HasChange("primary_network_interface.0.primary_ip.0.auto_delete") {
			auto := d.Get("primary_network_interface.0.primary_ip.0.auto_delete").(bool)
			reservedIpPath.AutoDelete = &auto
		}
		reservedIpPathAsPatch, err := reservedIpPath.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIpPath.AsPatch() failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateripoptions.ReservedIPPatch = reservedIpPathAsPatch
		_, _, err = vpcClient.UpdateSubnetReservedIPWithContext(context, updateripoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if d.HasChange(isInstanceNicSecurityGroups) && !d.IsNewResource() {

		ovs, nvs := d.GetChange(isInstanceNicSecurityGroups)
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)
		remove := flex.ExpandStringList(ov.Difference(nv).List())
		add := flex.ExpandStringList(nv.Difference(ov).List())
		if len(add) > 0 {
			for i := range add {
				createsgnicoptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{
					SecurityGroupID: &add[i],
					ID:              &network_interface_id,
				}
				_, _, err := vpcClient.CreateSecurityGroupTargetBindingWithContext(context, createsgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForInstanceAvailable(vpcClient, instance_id, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				deletesgnicoptions := &vpcv1.DeleteSecurityGroupTargetBindingOptions{
					SecurityGroupID: &remove[i],
					ID:              &network_interface_id,
				}
				_, err := vpcClient.DeleteSecurityGroupTargetBindingWithContext(context, deletesgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForInstanceAvailable(vpcClient, instance_id, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
		hasChange = true
	}
	if hasChange {
		isNICKey := "instance_key_" + instance_id
		conns.IbmMutexKV.Lock(isNICKey)
		defer conns.IbmMutexKV.Unlock(isNICKey)
		updateInstanceNetworkInterfaceOptions.NetworkInterfacePatch, _ = patchVals.AsPatch()
		_, _, err := vpcClient.UpdateInstanceNetworkInterfaceWithContext(context, updateInstanceNetworkInterfaceOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if d.HasChange(isInstanceNicFloatingIP) {
		floating_ip_id_old, floating_ip_id_new := d.GetChange(isInstanceNicFloatingIP)
		instance_id := parts[0]
		network_interface_id := parts[1]
		if floating_ip_id_new == nil || floating_ip_id_new.(string) == "" {
			removeInstanceNetworkInterfaceFloatingIPOptions := vpcClient.NewRemoveInstanceNetworkInterfaceFloatingIPOptions(instance_id, network_interface_id, floating_ip_id_old.(string))
			response, err := vpcClient.RemoveInstanceNetworkInterfaceFloatingIPWithContext(context, removeInstanceNetworkInterfaceFloatingIPOptions)
			if err != nil {
				if response.StatusCode == 404 {
					log.Println("[DEBUG] The specified floating IP address is not associated with the network interface with the specified identifier. ", err.Error())
				} else {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveInstanceNetworkInterfaceFloatingIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		} else {
			floating_ip_id := floating_ip_id_new.(string)
			getFloatingIPOptions := &vpcv1.GetFloatingIPOptions{
				ID: &floating_ip_id,
			}
			floatingip, response, err := vpcClient.GetFloatingIPWithContext(context, getFloatingIPOptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					d.SetId("")
					return nil
				}
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			if floatingip != nil && floatingip.Target != nil {
				floatingIpTarget := floatingip.Target.(*vpcv1.FloatingIPTarget)
				if *floatingIpTarget.ID != network_interface_id {
					d.Set(isInstanceNicFloatingIP, "")
					err = fmt.Errorf("Provided floating ip is already bound to another resource")
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

			addInstanceNetworkInterfaceFloatingIPOptions := &vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions{
				InstanceID:         &instance_id,
				NetworkInterfaceID: &network_interface_id,
				ID:                 &floating_ip_id,
			}

			_, response, err = vpcClient.AddInstanceNetworkInterfaceFloatingIPWithContext(context, addInstanceNetworkInterfaceFloatingIPOptions)

			if err != nil {
				d.Set(isInstanceNicFloatingIP, "")
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddInstanceNetworkInterfaceFloatingIPWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}

	}

	_, err = isWaitForNetworkInterfaceAvailable(vpcClient, d.Id(), d.Timeout(schema.TimeoutUpdate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForInstanceAvailable(vpcClient, instance_id, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return resourceIBMIsInstanceNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsInstanceNetworkInterfaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceNetworkInterfaceOptions := &vpcv1.DeleteInstanceNetworkInterfaceOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_network_interface", "delete", "sep-id-parts").GetDiag()
	}
	instance_id := parts[0]
	network_intf_id := parts[1]
	isNICKey := "instance_key_" + instance_id
	conns.IbmMutexKV.Lock(isNICKey)
	defer conns.IbmMutexKV.Unlock(isNICKey)

	deleteInstanceNetworkInterfaceOptions.SetInstanceID(instance_id)
	deleteInstanceNetworkInterfaceOptions.SetID(network_intf_id)

	_, err = vpcClient.DeleteInstanceNetworkInterfaceWithContext(context, deleteInstanceNetworkInterfaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_instance_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForNetworkInterfaceDelete(vpcClient, d.Id(), d.Timeout(schema.TimeoutUpdate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForNetworkInterfaceDelete failed: %s", err.Error()), "ibm_is_instance_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForInstanceAvailable(vpcClient, instance_id, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isWaitForNetworkInterfaceAvailable(vpcClient *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for dedicated host (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isNetworkInterfacePending},
		Target:     []string{isNetworkInterfaceAvailable, isNetworkInterfaceFailed},
		Refresh:    isNetworkInterfaceRefreshFunc(vpcClient, id, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isNetworkInterfaceRefreshFunc(vpcClient *vpcv1.VpcV1, id string, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{}
		parts, err := flex.SepIdParts(d.Id(), "/")
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error splitting ID in parts %s", err)
		}

		getInstanceNetworkInterfaceOptions.SetInstanceID(parts[0])
		getInstanceNetworkInterfaceOptions.SetID(parts[1])

		networkInterface, response, err := vpcClient.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)
		if err != nil {
			return nil, "", fmt.Errorf("GetInstanceNetworkInterface failed %s\n%s", err, response)
		}
		d.Set("status", *networkInterface.Status)

		if *networkInterface.Status == isNetworkInterfaceFailed {
			return networkInterface, *networkInterface.Status, fmt.Errorf("Network Interface creationg failed with status %s ", *networkInterface.Status)
		}
		return networkInterface, *networkInterface.Status, nil
	}
}

func isWaitForNetworkInterfaceDelete(vpcClient *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for dedicated host (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isNetworkInterfacePending, isNetworkInterfaceDeleting, isNetworkInterfaceAvailable},
		Target:     []string{isNetworkInterfaceDeleted},
		Refresh:    isNetworkInterfaceRefreshDeleteFunc(vpcClient, id, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isNetworkInterfaceRefreshDeleteFunc(vpcClient *vpcv1.VpcV1, id string, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{}
		parts, err := flex.SepIdParts(d.Id(), "/")
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error splitting ID in parts %s", err)
		}

		getInstanceNetworkInterfaceOptions.SetInstanceID(parts[0])
		getInstanceNetworkInterfaceOptions.SetID(parts[1])

		networkInterface, response, err := vpcClient.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return networkInterface, isNetworkInterfaceDeleted, nil
			}
			return nil, "", fmt.Errorf("GetInstanceNetworkInterface failed %s\n%s", err, response)
		}
		d.Set("status", *networkInterface.Status)

		if *networkInterface.Status == isNetworkInterfaceFailed {
			return networkInterface, *networkInterface.Status, fmt.Errorf("Network Interface creationg failed with status %s ", *networkInterface.Status)
		}
		return networkInterface, *networkInterface.Status, nil
	}
}
