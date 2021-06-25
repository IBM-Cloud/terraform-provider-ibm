// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerBandwidth                    = "bandwidth"
	isBareMetalServerBootTarget                   = "boot_target"
	isBareMetalServerCPU                          = "cpu"
	isBareMetalServerCPUArchitecture              = "architecture"
	isBareMetalServerCPUCoreCount                 = "core_count"
	isBareMetalServerCpuSocketCount               = "socket_count"
	isBareMetalServerCpuThreadPerCore             = "threads_per_core"
	isBareMetalServerCRN                          = "crn"
	isBareMetalServerDisks                        = "disks"
	isBareMetalServerDiskID                       = "id"
	isBareMetalServerDiskSize                     = "size"
	isBareMetalServerDiskName                     = "name"
	isBareMetalServerDiskInterfaceType            = "interface_type"
	isBareMetalServerHref                         = "href"
	isBareMetalServerMemory                       = "memory"
	isBareMetalServerTags                         = "tags"
	isBareMetalServerName                         = "name"
	isBareMetalServerEnableSecureBoot             = "enable_secure_boot"
	isBareMetalServerNetworkInterfaces            = "network_interfaces"
	isBareMetalServerPrimaryNetworkInterface      = "primary_network_interface"
	isBareMetalServerProfile                      = "profile"
	isBareMetalServerResourceGroup                = "resource_group"
	isBareMetalServerResourceType                 = "resource_type"
	isBareMetalServerStatus                       = "status"
	isBareMetalServerStatusReasons                = "status_reasons"
	isBareMetalServerTrustedPlatformModule        = "trusted_platform_module"
	isBareMetalServerTrustedPlatformModuleEnabled = "enabled"
	isBareMetalServerTrustedPlatformModuleMode    = "mode"
	isBareMetalServerVPC                          = "vpc"
	isBareMetalServerZone                         = "zone"
	isBareMetalServerStatusReasonsCode            = "code"
	isBareMetalServerStatusReasonsMessage         = "message"
	isBareMetalServerImage                        = "image"
	isBareMetalServerKeys                         = "keys"
	isBareMetalServerUserData                     = "user_data"
	isBareMetalServerNicName                      = "name"
	isBareMetalServerNicPortSpeed                 = "port_speed"
	isBareMetalServerNicAllowIPSpoofing           = "allow_ip_spoofing"
	isBareMetalServerNicSecurityGroups            = "security_groups"
	isBareMetalServerNicSubnet                    = "subnet"
	isBareMetalServerUserAccounts                 = "user_accounts"
	isBareMetalServerActionDeleting               = "deleting"
	isBareMetalServerActionDeleted                = "deleted"
	isBareMetalServerActionStatusStopping         = "stopping"
	isBareMetalServerActionStatusStopped          = "stopped"
	isBareMetalServerActionStatusStarting         = "starting"
	isBareMetalServerStatusRunning                = "running"
	isBareMetalServerStatusPending                = "pending"
	isBareMetalServerStatusRestarting             = "restarting"
	isBareMetalServerStatusFailed                 = "failed"
)

func resourceIBMisBareMetalServer() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISBareMetalServerCreate,
		Read:     resourceIBMISBareMetalServerRead,
		Update:   resourceIBMISBareMetalServerUpdate,
		Delete:   resourceIBMISBareMetalServerDelete,
		Exists:   resourceIBMISBareMetalServerExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			isBareMetalServerName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", isBareMetalServerName),
				Description:  "Bare metal server name",
			},

			isBareMetalServerBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total bandwidth (in megabits per second)",
			},
			isBareMetalServerBootTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this bare metal server disk",
			},

			isBareMetalServerCPU: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The bare metal server CPU configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerCPUArchitecture: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CPU architecture",
						},
						isBareMetalServerCPUCoreCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of cores",
						},
						isBareMetalServerCpuSocketCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of CPU sockets",
						},
						isBareMetalServerCpuThreadPerCore: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of hardware threads per core",
						},
					},
				},
			},
			isBareMetalServerCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this bare metal server",
			},
			isBareMetalServerDisks: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The disks for this bare metal server, including any disks that are associated with the boot_target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerDiskHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this bare metal server disk",
						},
						isBareMetalServerDiskID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this bare metal server disk",
						},
						isBareMetalServerDiskInterfaceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The disk interface used for attaching the disk. Supported values are [ nvme, sata ]",
						},
						isBareMetalServerDiskName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this disk",
						},
						isBareMetalServerDiskResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type",
						},
						isBareMetalServerDiskSize: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the disk in GB (gigabytes)",
						},
					},
				},
			},
			isBareMetalServerEnableSecureBoot: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether secure boot is enabled. If enabled, the image must support secure boot or the server will fail to boot.",
			},
			isBareMetalServerHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this bare metal server",
			},
			isBareMetalServerMemory: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of memory, truncated to whole gibibytes",
			},

			isBareMetalServerPrimaryNetworkInterface: {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Primary Network interface info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBareMetalServerNicHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network interface",
						},
						isBareMetalServerNicEnableInfraNAT: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.",
						},
						isBareMetalServerNicInterfaceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network interface type: [ pci, vlan ]",
						},
						isBareMetalServerNicReservedIps: {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The reserved IPs bound to this network interface.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
									isBareMetalServerNicIpAutoDelete: {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
									},
								},
							},
						},
						isBareMetalServerNicPrimaryIP: {
							Type:        schema.TypeList,
							Optional:    true,
							MinItems:    0,
							MaxItems:    1,
							Computed:    true,
							Description: "title: IPv4, The IP address. ",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
									isBareMetalServerNicIpAutoDelete: {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
									},
								},
							},
						},
						isBareMetalServerNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBareMetalServerNicName: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isBareMetalServerNicPortSpeed: {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
							Deprecated:       "This field is deprected",
						},

						isBareMetalServerNicSecurityGroups: {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBareMetalServerNicSubnet: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						isBareMetalServerNicAllowedVlans: {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Set:         schema.HashInt,
							Description: "Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.",
						},
					},
				},
			},

			isBareMetalServerNetworkInterfaces: {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBareMetalServerNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBareMetalServerNicName: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The user-defined name for this network interface. If unspecified, the name will be a hyphenated list of randomly-selected words",
						},
						isBareMetalServerNicHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network interface",
						},
						isBareMetalServerNicReservedIps: {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The reserved IPs bound to this network interface.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
									isBareMetalServerNicIpAutoDelete: {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
									},
								},
							},
						},
						isBareMetalServerNicEnableInfraNAT: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.",
						},
						isBareMetalServerNicInterfaceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network interface type: [ pci, vlan ]",
						},
						isBareMetalServerNicPrimaryIP: {
							Type:        schema.TypeList,
							MinItems:    0,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Description: "title: IPv4, The IP address. ",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
									isBareMetalServerNicIpAutoDelete: {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
									},
								},
							},
						},
						isBareMetalServerNicSecurityGroups: {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Description: "Collection of security group ids",
						},
						isBareMetalServerNicSubnet: {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         false,
							DiffSuppressFunc: applyOnce,
							Description:      "The associated subnet",
						},
						isBareMetalServerNicAllowedVlans: {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Set:         schema.HashInt,
							Description: "Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.",
						},

						isBareMetalServerNicAllowInterfaceToFloat: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
						},

						isBareMetalServerNicVlan: {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface",
						},
					},
				},
			},

			isBareMetalServerKeys: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "SSH key Ids for the bare metal server",
			},

			isBareMetalServerImage: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "image name",
			},
			isBareMetalServerProfile: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "profile name",
			},

			isBareMetalServerUserData: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "User data given for the bare metal server",
			},

			isBareMetalServerZone: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone name",
			},

			isBareMetalServerVPC: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The VPC the bare metal server is to be a part of",
			},

			isBareMetalServerResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource group name",
			},
			isBareMetalServerResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource type name",
			},

			isBareMetalServerStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bare metal server status",
			},

			isBareMetalServerStatusReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerStatusReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason",
						},

						isBareMetalServerStatusReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason",
						},
					},
				},
			},
			isBareMetalServerTrustedPlatformModule: {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerTrustedPlatformModuleEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Indicates whether the trusted platform module (TPM) is enabled. If enabled, mode will also be set.",
						},

						isBareMetalServerTrustedPlatformModuleMode: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered [ tpm_2, tpm_2_with_txt ] .",
						},
					},
				},
			},

			isBareMetalServerTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", "tag")},
				Set:         resourceIBMVPCHash,
				Description: "Tags for the Bare metal server",
			},
		},
	}
}

func resourceIBMISBareMetalServerValidator() *ResourceValidator {
	trustedPlatformModuleModes := "tpm_2, tpm_2_with_txt"
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBareMetalServerName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "tag",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBareMetalServerTrustedPlatformModuleMode,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              trustedPlatformModuleModes})
	ibmISBareMetalServerResourceValidator := ResourceValidator{ResourceName: "ibm_is_bare_metal_server", Schema: validateSchema}
	return &ibmISBareMetalServerResourceValidator
}

func resourceIBMISBareMetalServerCreate(d *schema.ResourceData, meta interface{}) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateBareMetalServerOptions{}

	if esb, ok := d.GetOk(isBareMetalServerEnableSecureBoot); ok {
		enableSecureBoot := esb.(bool)
		options.EnableSecureBoot = &enableSecureBoot
	}

	var imageStr string
	if image, ok := d.GetOk(isBareMetalServerImage); ok {
		imageStr = image.(string)
	}
	keySet := d.Get(isInstanceKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]vpcv1.KeyIdentityIntf, keySet.Len())
		for i, key := range keySet.List() {
			keystr := key.(string)
			keyobjs[i] = &vpcv1.KeyIdentity{
				ID: &keystr,
			}
		}
		options.Initialization = &vpcv1.BareMetalServerInitializationPrototype{
			Image: &vpcv1.ImageIdentity{
				ID: &imageStr,
			},
			Keys: keyobjs,
		}
		if userdata, ok := d.GetOk(isBareMetalServerUserData); ok {
			userdatastr := userdata.(string)
			options.Initialization.UserData = &userdatastr
		}
	}

	if name, ok := d.GetOk(isBareMetalServerName); ok {
		nameStr := name.(string)
		options.Name = &nameStr
	}

	if primnicintf, ok := d.GetOk(isBareMetalServerPrimaryNetworkInterface); ok && len(primnicintf.([]interface{})) > 0 {
		primnic := primnicintf.([]interface{})[0].(map[string]interface{})
		subnetintf, _ := primnic[isBareMetalServerNicSubnet]
		subnetintfstr := subnetintf.(string)
		var primnicobj = &vpcv1.BareMetalServerPrimaryNetworkInterfacePrototype{}
		primnicobj.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnetintfstr,
		}
		name, _ := primnic[isBareMetalServerNicName]
		namestr := name.(string)
		if namestr != "" {
			primnicobj.Name = &namestr
		}
		if ipsIntf, ok := primnic[isBareMetalServerNicReservedIps]; ok {
			ips := ipsIntf.([]interface{})
			var intfs []vpcv1.NetworkInterfaceIPPrototypeIntf
			for _, resource := range ips {
				ip := resource.(map[string]interface{})
				if reservedIpId, ok := ip[isBareMetalServerNicIpID]; ok {
					reservedIpIdStr := reservedIpId.(string)
					nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
						ID: &reservedIpIdStr,
					}
					intfs = append(intfs, nicReservedIP)
				} else {
					reservedIpAddress := ip[isBareMetalServerNicIpAddress].(string)
					reservedIpAutoDelete := ip[isBareMetalServerNicIpAutoDelete].(bool)
					reservedIpName := ip[isBareMetalServerNicIpName].(string)
					nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
						Address:    &reservedIpAddress,
						AutoDelete: &reservedIpAutoDelete,
						Name:       &reservedIpName,
					}
					intfs = append(intfs, nicReservedIP)
				}
			}
			primnicobj.Ips = intfs
		}

		if primaryIpIntf, ok := primnic[isBareMetalServerNicPrimaryIP]; ok && len(primaryIpIntf.([]interface{})) > 0 {
			primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})
			if reservedIpId, ok := primaryIp[isBareMetalServerNicIpID]; ok {
				reservedIpIdStr := reservedIpId.(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
					ID: &reservedIpIdStr,
				}
				primnicobj.PrimaryIP = nicReservedIP
			} else {
				reservedIpAddress := primaryIp[isBareMetalServerNicIpAddress].(string)
				reservedIpAutoDelete := primaryIp[isBareMetalServerNicIpAutoDelete].(bool)
				reservedIpName := primaryIp[isBareMetalServerNicIpName].(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
					Address:    &reservedIpAddress,
					AutoDelete: &reservedIpAutoDelete,
					Name:       &reservedIpName,
				}
				primnicobj.PrimaryIP = nicReservedIP
			}
		}

		allowIPSpoofing, ok := primnic[isBareMetalServerNicAllowIPSpoofing]
		allowIPSpoofingbool := allowIPSpoofing.(bool)
		if ok {
			primnicobj.AllowIPSpoofing = &allowIPSpoofingbool
		}

		enableInfraNAT, ok := primnic[isBareMetalServerNicEnableInfraNAT]
		enableInfraNATbool := enableInfraNAT.(bool)
		if ok {
			primnicobj.EnableInfrastructureNat = &enableInfraNATbool
		}

		secgrpintf, ok := primnic[isBareMetalServerNicSecurityGroups]
		if ok {
			secgrpSet := secgrpintf.(*schema.Set)
			if secgrpSet.Len() != 0 {
				var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
				for i, secgrpIntf := range secgrpSet.List() {
					secgrpIntfstr := secgrpIntf.(string)
					secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
						ID: &secgrpIntfstr,
					}
				}
				primnicobj.SecurityGroups = secgrpobjs
			}
		}

		if allowedVlansOk, ok := primnic[isBareMetalServerNicAllowedVlans]; ok {
			allowedVlansList := allowedVlansOk.(*schema.Set).List()

			allowedVlans := make([]int64, 0, len(allowedVlansList))
			for _, k := range allowedVlansList {
				allowedVlans = append(allowedVlans, int64(k.(int)))
			}
			primnicobj.AllowedVlans = allowedVlans
		}
		options.PrimaryNetworkInterface = primnicobj
	}

	if nicsintf, ok := d.GetOk(isBareMetalServerNetworkInterfaces); ok {
		nics := nicsintf.([]interface{})
		for _, resource := range nics {
			nic := resource.(map[string]interface{})
			interfaceType := ""
			if allowedVlansOk, ok := nic[isBareMetalServerNicAllowedVlans]; ok {
				interfaceType = "pci"
				var nicobj = &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{}
				nicobj.InterfaceType = &interfaceType

				allowedVlansList := allowedVlansOk.(*schema.Set).List()

				allowedVlans := make([]int64, 0, len(allowedVlansList))
				for _, k := range allowedVlansList {
					allowedVlans = append(allowedVlans, int64(k.(int)))
				}
				nicobj.AllowedVlans = allowedVlans

				subnetintf, _ := nic[isBareMetalServerNicSubnet]
				subnetintfstr := subnetintf.(string)
				nicobj.Subnet = &vpcv1.SubnetIdentity{
					ID: &subnetintfstr,
				}
				name, _ := nic[isBareMetalServerNicName]
				namestr := name.(string)
				if namestr != "" {
					nicobj.Name = &namestr
				}

				enableInfraNAT, ok := nic[isBareMetalServerNicEnableInfraNAT]
				enableInfraNATbool := enableInfraNAT.(bool)
				if ok {
					nicobj.EnableInfrastructureNat = &enableInfraNATbool
				}

				if ipsIntf, ok := nic[isBareMetalServerNicReservedIps]; ok && len(ipsIntf.([]interface{})) > 0 {
					ips := ipsIntf.([]interface{})
					var intfs []vpcv1.NetworkInterfaceIPPrototypeIntf
					for _, resource := range ips {
						ip := resource.(map[string]interface{})
						if reservedIpId, ok := ip[isBareMetalServerNicIpID]; ok {
							reservedIpIdStr := reservedIpId.(string)
							nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
								ID: &reservedIpIdStr,
							}
							intfs = append(intfs, nicReservedIP)
						} else {
							reservedIpAddress := ip[isBareMetalServerNicIpAddress].(string)
							reservedIpAutoDelete := ip[isBareMetalServerNicIpAutoDelete].(bool)
							reservedIpName := ip[isBareMetalServerNicIpName].(string)
							nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
								Address:    &reservedIpAddress,
								AutoDelete: &reservedIpAutoDelete,
								Name:       &reservedIpName,
							}
							intfs = append(intfs, nicReservedIP)
						}
					}
					nicobj.Ips = intfs
				}

				if primaryIpIntf, ok := nic[isBareMetalServerNicPrimaryIP]; ok && len(primaryIpIntf.([]interface{})) > 0 {
					primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})
					if reservedIpId, ok := primaryIp[isBareMetalServerNicIpID]; ok {
						reservedIpIdStr := reservedIpId.(string)
						nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
							ID: &reservedIpIdStr,
						}
						nicobj.PrimaryIP = nicReservedIP
					} else {
						reservedIpAddress := primaryIp[isBareMetalServerNicIpAddress].(string)
						reservedIpAutoDelete := primaryIp[isBareMetalServerNicIpAutoDelete].(bool)
						reservedIpName := primaryIp[isBareMetalServerNicIpName].(string)
						nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
							Address:    &reservedIpAddress,
							AutoDelete: &reservedIpAutoDelete,
							Name:       &reservedIpName,
						}
						nicobj.PrimaryIP = nicReservedIP
					}
				}

				allowIPSpoofing, ok := nic[isBareMetalServerNicAllowIPSpoofing]
				allowIPSpoofingbool := allowIPSpoofing.(bool)
				if ok {
					nicobj.AllowIPSpoofing = &allowIPSpoofingbool
				}
				secgrpintf, ok := nic[isBareMetalServerNicSecurityGroups]
				if ok {
					secgrpSet := secgrpintf.(*schema.Set)
					if secgrpSet.Len() != 0 {
						var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
						for i, secgrpIntf := range secgrpSet.List() {
							secgrpIntfstr := secgrpIntf.(string)
							secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
								ID: &secgrpIntfstr,
							}
						}
						nicobj.SecurityGroups = secgrpobjs
					}
				}
			} else {
				interfaceType = "vlan"
				var nicobj = &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByVlanPrototype{}
				nicobj.InterfaceType = &interfaceType

				if aitf, ok := nic[isBareMetalServerNicAllowInterfaceToFloat]; ok {
					allowInterfaceToFloat := aitf.(bool)
					nicobj.AllowInterfaceToFloat = &allowInterfaceToFloat
				}
				if vlan, ok := nic[isBareMetalServerNicVlan]; ok {
					vlanInt := int64(vlan.(int))
					nicobj.Vlan = &vlanInt
				}

				subnetintf, _ := nic[isBareMetalServerNicSubnet]
				subnetintfstr := subnetintf.(string)
				nicobj.Subnet = &vpcv1.SubnetIdentity{
					ID: &subnetintfstr,
				}
				name, _ := nic[isBareMetalServerNicName]
				namestr := name.(string)
				if namestr != "" {
					nicobj.Name = &namestr
				}

				enableInfraNAT, ok := nic[isBareMetalServerNicEnableInfraNAT]
				enableInfraNATbool := enableInfraNAT.(bool)
				if ok {
					nicobj.EnableInfrastructureNat = &enableInfraNATbool
				}

				if ipsIntf, ok := nic[isBareMetalServerNicReservedIps]; ok {
					ips := ipsIntf.([]interface{})
					var intfs []vpcv1.NetworkInterfaceIPPrototypeIntf
					for _, resource := range ips {
						ip := resource.(map[string]interface{})
						if reservedIpId, ok := ip[isBareMetalServerNicIpID]; ok {
							reservedIpIdStr := reservedIpId.(string)
							nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
								ID: &reservedIpIdStr,
							}
							intfs = append(intfs, nicReservedIP)
						} else {
							reservedIpAddress := ip[isBareMetalServerNicIpAddress].(string)
							reservedIpAutoDelete := ip[isBareMetalServerNicIpAutoDelete].(bool)
							reservedIpName := ip[isBareMetalServerNicIpName].(string)
							nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
								Address:    &reservedIpAddress,
								AutoDelete: &reservedIpAutoDelete,
								Name:       &reservedIpName,
							}
							intfs = append(intfs, nicReservedIP)
						}
					}
					nicobj.Ips = intfs
				}

				if primaryIpIntf, ok := nic[isBareMetalServerNicPrimaryIP]; ok && len(primaryIpIntf.([]interface{})) > 0 {
					primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})
					if reservedIpId, ok := primaryIp[isBareMetalServerNicIpID]; ok {
						reservedIpIdStr := reservedIpId.(string)
						nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
							ID: &reservedIpIdStr,
						}
						nicobj.PrimaryIP = nicReservedIP
					} else {
						reservedIpAddress := primaryIp[isBareMetalServerNicIpAddress].(string)
						reservedIpAutoDelete := primaryIp[isBareMetalServerNicIpAutoDelete].(bool)
						reservedIpName := primaryIp[isBareMetalServerNicIpName].(string)
						nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
							Address:    &reservedIpAddress,
							AutoDelete: &reservedIpAutoDelete,
							Name:       &reservedIpName,
						}
						nicobj.PrimaryIP = nicReservedIP
					}
				}

				allowIPSpoofing, ok := nic[isBareMetalServerNicAllowIPSpoofing]
				allowIPSpoofingbool := allowIPSpoofing.(bool)
				if ok {
					nicobj.AllowIPSpoofing = &allowIPSpoofingbool
				}
				secgrpintf, ok := nic[isBareMetalServerNicSecurityGroups]
				if ok {
					secgrpSet := secgrpintf.(*schema.Set)
					if secgrpSet.Len() != 0 {
						var secgrpobjs = make([]vpcv1.SecurityGroupIdentityIntf, secgrpSet.Len())
						for i, secgrpIntf := range secgrpSet.List() {
							secgrpIntfstr := secgrpIntf.(string)
							secgrpobjs[i] = &vpcv1.SecurityGroupIdentity{
								ID: &secgrpIntfstr,
							}
						}
						nicobj.SecurityGroups = secgrpobjs
					}
				}
			}
		}
	}

	if rgrp, ok := d.GetOk(isBareMetalServerResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	if p, ok := d.GetOk(isBareMetalServerProfile); ok {
		profile := p.(string)
		options.Profile = &vpcv1.BareMetalServerProfileIdentity{
			Name: &profile,
		}
	}

	if z, ok := d.GetOk(isBareMetalServerZone); ok {
		zone := z.(string)
		options.Zone = &vpcv1.ZoneIdentity{
			Name: &zone,
		}
	}

	if v, ok := d.GetOk(isBareMetalServerVPC); ok {
		vpc := v.(string)
		options.VPC = &vpcv1.VPCIdentity{
			ID: &vpc,
		}
	}

	if tpmIntf, ok := d.GetOk(isBareMetalServerTrustedPlatformModule); ok && len(tpmIntf.([]interface{})) > 0 {
		tpm := tpmIntf.([]interface{})[0].(map[string]interface{})
		enabledIntf, _ := tpm[isBareMetalServerTrustedPlatformModuleEnabled]
		enabled := enabledIntf.(bool)
		var tpmobj = &vpcv1.BareMetalServerTrustedPlatformModulePrototype{}
		tpmobj.Enabled = &enabled
		mode, _ := tpm[isBareMetalServerTrustedPlatformModuleMode]
		modestr := mode.(string)
		if modestr != "" {
			tpmobj.Mode = &modestr
		}
		options.TrustedPlatformModule = tpmobj
	}

	bms, response, err := sess.CreateBareMetalServer(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create bare metal server err %s\n%s", err, response)
	}
	d.SetId(*bms.ID)
	log.Printf("[INFO] Bare Metal Server : %s", *bms.ID)
	_, err = isWaitForBareMetalServerAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVolumeTags); ok || v != "" {
		oldList, newList := d.GetChange(isBareMetalServerTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *bms.CRN)
		if err != nil {
			log.Printf(
				"Error on create of resource bare metal server (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISBareMetalServerRead(d, meta)
}

func resourceIBMISBareMetalServerRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	err := bareMetalServerGet(d, meta, id)
	if err != nil {
		return err
	}

	return nil
}

func bareMetalServerGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	bms, response, err := sess.GetBareMetalServer(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Bare Metal Server (%s): %s\n%s", id, err, response)
	}
	d.SetId(*bms.ID)
	d.Set(isBareMetalServerBandwidth, *bms.Bandwidth)
	bmsBootTargetIntf := bms.BootTarget.(*vpcv1.BareMetalServerBootTarget)
	bmsBootTarget := bmsBootTargetIntf.ID
	d.Set(isBareMetalServerBootTarget, bmsBootTarget)
	cpuList := make([]map[string]interface{}, 0)
	if bms.Cpu != nil {
		currentCPU := map[string]interface{}{}
		currentCPU[isBareMetalServerCPUArchitecture] = *bms.Cpu.Architecture
		currentCPU[isBareMetalServerCPUCoreCount] = *bms.Cpu.CoreCount
		currentCPU[isBareMetalServerCpuSocketCount] = *bms.Cpu.SocketCount
		currentCPU[isBareMetalServerCpuThreadPerCore] = *bms.Cpu.ThreadsPerCore
		cpuList = append(cpuList, currentCPU)
	}
	d.Set(isBareMetalServerCPU, cpuList)
	d.Set(isBareMetalServerCRN, *bms.CRN)

	diskList := make([]map[string]interface{}, 0)
	if bms.Disks != nil {
		for _, disk := range bms.Disks {
			currentDisk := map[string]interface{}{
				isBareMetalServerDiskHref:          disk.Href,
				isBareMetalServerDiskID:            disk.ID,
				isBareMetalServerDiskInterfaceType: disk.InterfaceType,
				isBareMetalServerDiskName:          disk.Name,
				isBareMetalServerDiskResourceType:  disk.ResourceType,
				isBareMetalServerDiskSize:          disk.Size,
			}
			diskList = append(diskList, currentDisk)
		}
	}
	d.Set(isBareMetalServerDisks, diskList)
	if bms.EnableSecureBoot != nil {
		d.Set(isBareMetalServerEnableSecureBoot, *bms.EnableSecureBoot)
	}
	d.Set(isBareMetalServerHref, *bms.Href)
	d.Set(isBareMetalServerMemory, *bms.Memory)
	d.Set(isBareMetalServerName, *bms.Name)
	//pni

	if bms.PrimaryNetworkInterface != nil {
		primaryNicList := make([]map[string]interface{}, 0)
		currentPrimNic := map[string]interface{}{}
		currentPrimNic["id"] = *bms.PrimaryNetworkInterface.ID
		currentPrimNic[isBareMetalServerNicName] = *bms.PrimaryNetworkInterface.Name
		currentPrimNic[isBareMetalServerNicHref] = *bms.PrimaryNetworkInterface.Href
		currentPrimNic[isBareMetalServerNicSubnet] = *bms.PrimaryNetworkInterface.Subnet.ID
		// currentPrimNic[isBareMetalServerNicPrimaryIpv4Address] = *bms.PrimaryNetworkInterface.PrimaryIP.Address
		getnicoptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &id,
			ID:                bms.PrimaryNetworkInterface.ID,
		}
		bmsnic, response, err := sess.GetBareMetalServerNetworkInterface(getnicoptions)
		if err != nil {
			return fmt.Errorf("Error getting network interfaces attached to the bare metal server %s\n%s", err, response)
		}

		primaryIpList := make([]map[string]interface{}, 0)
		if bms.PrimaryNetworkInterface.PrimaryIP != nil {
			currentIP := map[string]interface{}{
				// isBareMetalServerNicIpHref:       *bms.PrimaryNetworkInterface.PrimaryIP.Href,
				isBareMetalServerNicIpID: *bms.PrimaryNetworkInterface.PrimaryIP.ID,
				// isBareMetalServerNicResourceType: *bms.PrimaryNetworkInterface.PrimaryIP.ResourceType,
				// isBareMetalServerNicIpName:       *bms.PrimaryNetworkInterface.PrimaryIP.Name,
				isBareMetalServerNicIpAddress: *bms.PrimaryNetworkInterface.PrimaryIP.Address,
			}
			primaryIpList = append(primaryIpList, currentIP)
		}
		currentPrimNic[isBareMetalServerNicPrimaryIP] = primaryIpList

		switch reflect.TypeOf(bmsnic).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				primNic := bmsnic.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
				currentPrimNic[isInstanceNicAllowIPSpoofing] = *primNic.AllowIPSpoofing
				if len(primNic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(primNic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string(*(primNic.SecurityGroups[i].ID)))
					}
					currentPrimNic[isInstanceNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}
			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				primNic := bmsnic.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				currentPrimNic[isInstanceNicAllowIPSpoofing] = *primNic.AllowIPSpoofing

				if len(primNic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(primNic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string(*(primNic.SecurityGroups[i].ID)))
					}
					currentPrimNic[isInstanceNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}
			}
		}

		primaryNicList = append(primaryNicList, currentPrimNic)
		d.Set(isBareMetalServerPrimaryNetworkInterface, primaryNicList)
	}

	//ni

	interfacesList := make([]map[string]interface{}, 0)
	for _, intfc := range bms.NetworkInterfaces {
		if *intfc.ID != *bms.PrimaryNetworkInterface.ID {
			currentNic := map[string]interface{}{}
			currentNic["id"] = *intfc.ID
			currentNic[isBareMetalServerNicName] = *intfc.Name
			// currentNic[isBareMetalServerNicPrimaryIpv4Address] = *intfc.PrimaryIP.Address
			getnicoptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: &id,
				ID:                intfc.ID,
			}
			bmsnicintf, response, err := sess.GetBareMetalServerNetworkInterface(getnicoptions)
			if err != nil {
				return fmt.Errorf("Error getting network interfaces attached to the bare metal server %s\n%s", err, response)
			}
			primaryIpList := make([]map[string]interface{}, 0)
			if intfc.PrimaryIP != nil {
				currentIP := map[string]interface{}{
					// isBareMetalServerNicIpHref:       *intfc.PrimaryIP.Href,
					isBareMetalServerNicIpID: *intfc.PrimaryIP.ID,
					// isBareMetalServerNicResourceType: *intfc.PrimaryIP.ResourceType,
					// isBareMetalServerNicIpName:       *intfc.PrimaryIP.Name,
					isBareMetalServerNicIpAddress: *intfc.PrimaryIP.Address,
				}
				primaryIpList = append(primaryIpList, currentIP)
			}
			currentNic[isBareMetalServerNicPrimaryIP] = primaryIpList
			switch reflect.TypeOf(bmsnicintf).String() {
			case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
				{
					bmsnic := bmsnicintf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
					currentNic[isBareMetalServerNicAllowIPSpoofing] = *bmsnic.AllowIPSpoofing
					currentNic[isBareMetalServerNicSubnet] = *bmsnic.Subnet.ID
					currentNic[isBareMetalServerNicInterfaceType] = "pci"
					if len(bmsnic.SecurityGroups) != 0 {
						secgrpList := []string{}
						for i := 0; i < len(bmsnic.SecurityGroups); i++ {
							secgrpList = append(secgrpList, string(*(bmsnic.SecurityGroups[i].ID)))
						}
						currentNic[isBareMetalServerNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
					}
				}
			case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
				{
					bmsnic := bmsnicintf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
					currentNic[isBareMetalServerNicAllowIPSpoofing] = *bmsnic.AllowIPSpoofing
					currentNic[isBareMetalServerNicSubnet] = *bmsnic.Subnet.ID
					currentNic[isBareMetalServerNicInterfaceType] = "vlan"

					if len(bmsnic.SecurityGroups) != 0 {
						secgrpList := []string{}
						for i := 0; i < len(bmsnic.SecurityGroups); i++ {
							secgrpList = append(secgrpList, string(*(bmsnic.SecurityGroups[i].ID)))
						}
						currentNic[isBareMetalServerNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
					}
				}
			}
			interfacesList = append(interfacesList, currentNic)
		}
	}
	d.Set(isBareMetalServerNetworkInterfaces, interfacesList)

	d.Set(isBareMetalServerProfile, *bms.Profile.Name)
	if bms.ResourceGroup != nil {
		d.Set(isBareMetalServerResourceGroup, *bms.ResourceGroup.ID)
	}
	d.Set(isBareMetalServerResourceType, *bms.ResourceType)
	d.Set(isBareMetalServerStatus, *bms.Status)
	statusReasonsList := make([]map[string]interface{}, 0)
	if bms.StatusReasons != nil {
		for _, sr := range bms.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isBareMetalServerStatusReasonsCode] = *sr.Code
				currentSR[isBareMetalServerStatusReasonsMessage] = *sr.Message
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
	}
	d.Set(isBareMetalServerStatusReasons, statusReasonsList)
	tpmList := make([]map[string]interface{}, 0)
	if bms.TrustedPlatformModule != nil {
		currentTPM := map[string]interface{}{
			isBareMetalServerTrustedPlatformModuleEnabled: *bms.TrustedPlatformModule.Enabled,
			isBareMetalServerTrustedPlatformModuleMode:    *bms.TrustedPlatformModule.Mode,
		}
		tpmList = append(tpmList, currentTPM)
	}
	d.Set(isBareMetalServerTrustedPlatformModule, tpmList)

	d.Set(isBareMetalServerVPC, *bms.VPC.ID)
	d.Set(isBareMetalServerZone, *bms.Zone.Name)

	tags, err := GetTagsUsingCRN(meta, *bms.CRN)
	if err != nil {
		log.Printf(
			"Error on get of resource bare metal server (%s) tags: %s", d.Id(), err)
	}
	d.Set(isBareMetalServerTags, tags)

	return nil
}

func resourceIBMISBareMetalServerUpdate(d *schema.ResourceData, meta interface{}) error {

	id := d.Id()

	err := bareMetalServerUpdate(d, meta, id)
	if err != nil {
		return err
	}

	return resourceIBMISVolumeRead(d, meta)
}

func bareMetalServerUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.UpdateBareMetalServerOptions{
		ID: &id,
	}
	bmsPatchModel := &vpcv1.BareMetalServerPatch{}
	flag := false
	if d.HasChange(isBareMetalServerName) {
		flag = true
		nameStr := ""
		if name, ok := d.GetOk(isBareMetalServerName); ok {
			nameStr = name.(string)
		}
		bmsPatchModel.Name = &nameStr
	}

	if d.HasChange(isBareMetalServerEnableSecureBoot) {
		flag = true
		enableSecureBoot := false
		if esb, ok := d.GetOk(isBareMetalServerEnableSecureBoot); ok {
			enableSecureBoot = esb.(bool)
		}
		bmsPatchModel.EnableSecureBoot = &enableSecureBoot
	}
	mode := fmt.Sprintf("%s.0.%s", isBareMetalServerTrustedPlatformModule, isBareMetalServerTrustedPlatformModuleMode)
	enabled := fmt.Sprintf("%s.0.%s", isBareMetalServerTrustedPlatformModule, isBareMetalServerTrustedPlatformModuleEnabled)
	if d.HasChange(mode) && !d.IsNewResource() {
		flag = true
		modeStr := ""
		if modeOk, ok := d.GetOk(mode); ok {
			modeStr = modeOk.(string)
		}
		bmsPatchModel.TrustedPlatformModule.Mode = &modeStr
	}
	if d.HasChange(enabled) && !d.IsNewResource() {
		flag = true
		enabledBool := false
		if enabledOk, ok := d.GetOk(enabled); ok {
			enabledBool = enabledOk.(bool)
		}
		bmsPatchModel.TrustedPlatformModule.Enabled = &enabledBool
	}
	if flag {
		bmsPatch, err := bmsPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("Error calling asPatch for BareMetalServerPatch: %s", err)
		}
		options.BareMetalServerPatch = bmsPatch
		_, response, err := sess.UpdateBareMetalServer(options)
		if err != nil {
			return fmt.Errorf("Error updating Bare Metal Server: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISBareMetalServerDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()

	err := bareMetalServerDelete(d, meta, id)
	if err != nil {
		return err
	}

	return nil
}

func bareMetalServerDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getBmsOptions := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	bms, response, err := sess.GetBareMetalServer(getBmsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error Getting Bare Metal Server (%s): %s\n%s", id, err, response)
	}
	if *bms.Status == "running" {
		options := &vpcv1.CreateBareMetalServerStopOptions{
			ID: bms.ID,
		}
		response, err = sess.CreateBareMetalServerStop(options)
	}

	options := &vpcv1.DeleteBareMetalServerOptions{
		ID: &id,
	}
	response, err = sess.DeleteBareMetalServer(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Bare Metal Server : %s\n%s", err, response)
	}
	_, err = isWaitForBareMetalServerDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForBareMetalServerDeleted(bmsC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isBareMetalServerActionDeleting},
		Target:     []string{"done", "", isBareMetalServerActionDeleted},
		Refresh:    isBareMetalServerDeleteRefreshFunc(bmsC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isBareMetalServerDeleteRefreshFunc(bmsC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := bmsC.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return bms, isBareMetalServerActionDeleted, nil
			}
			return bms, "", fmt.Errorf("Error Getting Bare Metal Server: %s\n%s", err, response)
		}
		return bms, isBareMetalServerActionDeleting, err
	}
}

func resourceIBMISBareMetalServerExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()

	exists, err := bareMetalServerExists(d, meta, id)
	return exists, err

}

func bareMetalServerExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	options := &vpcv1.GetBareMetalServerOptions{
		ID: &id,
	}
	_, response, err := sess.GetBareMetalServer(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Bare Metal Server: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForBareMetalServerAvailable(client *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) to be available.", id)
	communicator := make(chan interface{})
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isBareMetalServerStatusPending, isBareMetalServerActionStatusStarting},
		Target:     []string{isBareMetalServerStatusRunning, isBareMetalServerStatusFailed},
		Refresh:    isBareMetalServerRefreshFunc(client, id, d, communicator),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerRefreshFunc(client *vpcv1.VpcV1, id string, d *schema.ResourceData, communicator chan interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := client.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error getting Bare Metal Server: %s\n%s", err, response)
		}
		d.Set(isBareMetalServerStatus, *bms.Status)

		select {
		case data := <-communicator:
			return nil, "", data.(error)
		default:
			fmt.Println("no message sent")
		}

		if *bms.Status == "running" || *bms.Status == "failed" {
			// let know the isRestartStartAction() to stop
			close(communicator)
			return bms, *bms.Status, nil

		}
		return bms, isBareMetalServerStatusPending, nil
	}
}

func isWaitForBareMetalServerActionStop(bmsC *vpcv1.VpcV1, timeout time.Duration, id string, d *schema.ResourceData) (interface{}, error) {
	communicator := make(chan interface{})
	stateConf := &resource.StateChangeConf{
		Pending: []string{isBareMetalServerStatusRunning, isBareMetalServerStatusPending, isBareMetalServerActionStatusStopping},
		Target:  []string{isBareMetalServerActionStatusStopped, isBareMetalServerStatusFailed, ""},
		Refresh: func() (interface{}, string, error) {
			getbmsoptions := &vpcv1.GetBareMetalServerOptions{
				ID: &id,
			}
			bms, response, err := bmsC.GetBareMetalServer(getbmsoptions)
			if err != nil {
				return nil, "", fmt.Errorf("Error Getting Bare Metal Server: %s\n%s", err, response)
			}
			select {
			case data := <-communicator:
				return nil, "", data.(error)
			default:
				fmt.Println("no message sent")
			}
			if *bms.Status == isBareMetalServerStatusFailed {
				// let know the isRestartStopAction() to stop
				close(communicator)
				return bms, *bms.Status, fmt.Errorf("The  Bare Metal Server %s failed to stop: %v", id, err)
			}
			return bms, *bms.Status, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	if v, ok := d.GetOk("force_recovery_time"); ok {
		forceTimeout := v.(int)
		go isRestartStopAction(bmsC, id, d, forceTimeout, communicator)
	}

	return stateConf.WaitForState()
}

func isBareMetalServerRestartStopAction(bmsC *vpcv1.VpcV1, id string, d *schema.ResourceData, forceTimeout int, communicator chan interface{}) {
	subticker := time.NewTicker(time.Duration(forceTimeout) * time.Minute)
	//subticker := time.NewTicker(time.Duration(forceTimeout) * time.Second)
	for {
		select {

		case <-subticker.C:
			log.Println("Bare Metal Server is still in stopping state, retrying to stop with -force")
			actiontype := "hard"
			createbmssactoptions := &vpcv1.CreateBareMetalServerStopOptions{
				ID:   &id,
				Type: &actiontype,
			}
			response, err := bmsC.CreateBareMetalServerStop(createbmssactoptions)
			if err != nil {
				communicator <- fmt.Errorf("Error retrying Bare Metal Server action stop: %s\n%s", err, response)
				return
			}
		case <-communicator:
			// indicates refresh func is reached target and not proceed with the thread)
			subticker.Stop()
			return

		}
	}
}
