// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"reflect"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBareMetalServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBareMetalServerRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{isBareMetalServerName},
				ValidateFunc:  InvokeDataSourceValidator("ibm_is_bare_metal_server", "identifier"),
			},
			isBareMetalServerName: {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"identifier"},
				ValidateFunc:  InvokeValidator("ibm_is_bare_metal_server", isBareMetalServerName),
				Description:   "Bare metal server name",
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
				Type: schema.TypeList,
				// MinItems:    1,
				// MaxItems:    1,
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
				Computed:    true,
				Description: "Primary Network interface info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBareMetalServerNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBareMetalServerNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBareMetalServerNicPortSpeed: {
							Type:       schema.TypeInt,
							Computed:   true,
							Deprecated: "This field is deprected",
						},
						isBareMetalServerNicHref: {
							Type:       schema.TypeString,
							Computed:   true,
							Deprecated: "This URL of the interface",
						},

						isBareMetalServerNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBareMetalServerNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isBareMetalServerNetworkInterfaces: {
				Type:     schema.TypeList,
				Computed: true,
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
						isBareMetalServerNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether IP spoofing is allowed on this interface.",
						},
						isBareMetalServerNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isBareMetalServerNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isBareMetalServerNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isBareMetalServerKeys: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "SSH key Ids for the bare metal server",
			},

			isBareMetalServerImage: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "image name",
			},
			isBareMetalServerProfile: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "profil name",
			},

			isBareMetalServerUserData: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User data given for the bare metal server",
			},

			isBareMetalServerZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			isBareMetalServerVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The VPC the bare metal server is to be a part of",
			},

			isBareMetalServerResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerTrustedPlatformModuleEnabled: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the trusted platform module (TPM) is enabled. If enabled, mode will also be set.",
						},

						isBareMetalServerTrustedPlatformModuleMode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered [ tpm_2, tpm_2_with_txt ] .",
						},
					},
				},
			},

			isBareMetalServerTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: InvokeValidator("ibm_is_bare_metal_server", "tag")},
				Set:         resourceIBMVPCHash,
				Description: "Tags for the Bare metal server",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBareMetalServerName,
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString})

	ibmISBMSDataSourceValidator := ResourceValidator{ResourceName: "ibm_is_bare_metal_server", Schema: validateSchema}
	return &ibmISBMSDataSourceValidator
}

func dataSourceIBMISBareMetalServerRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Get("identifier").(string)
	name := d.Get(isBareMetalServerName).(string)

	err := bmsGetById(d, meta, id, name)
	if err != nil {
		return err
	}
	return nil
}

func bmsGetById(d *schema.ResourceData, meta interface{}, id, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerOptions{}
	if id != "" {
		options.ID = &id
	}
	// else if name != "" {
	// 	options.Name = &name
	// }
	bms, response, err := sess.GetBareMetalServer(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Bare Metal Server (%s): %s\n%s", id, err, response)
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

			switch reflect.TypeOf(bmsnicintf).String() {
			case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
				{
					bmsnic := bmsnicintf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
					currentNic[isBareMetalServerNicAllowIPSpoofing] = *bmsnic.AllowIPSpoofing
					currentNic[isBareMetalServerNicSubnet] = *bmsnic.Subnet.ID
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
