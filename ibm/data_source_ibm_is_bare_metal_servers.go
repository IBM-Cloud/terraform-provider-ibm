// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServers = "servers"
)

func dataSourceIBMISBareMetalServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBareMetalServersRead,

		Schema: map[string]*schema.Schema{

			isBareMetalServers: {
				Type:        schema.TypeList,
				Description: "List of Bare Metal Servers",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bare metal server id",
						},
						isBareMetalServerName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bare metal server name",
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
				},
			},
		},
	}
}

func dataSourceIBMISBareMetalServersRead(d *schema.ResourceData, meta interface{}) error {

	err := bmsList(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func bmsList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.BareMetalServer{}
	for {
		listBareMetalServersOptions := &vpcv1.ListBareMetalServersOptions{}
		if start != "" {
			listBareMetalServersOptions.Start = &start
		}
		availableServers, response, err := sess.ListBareMetalServers(listBareMetalServersOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Bare Metal Servers %s\n%s", err, response)
		}
		start = GetNext(availableServers.Next)
		allrecs = append(allrecs, availableServers.BareMetalServers...)
		if start == "" {
			break
		}
	}

	serversInfo := make([]map[string]interface{}, 0)
	for _, bms := range allrecs {

		l := map[string]interface{}{
			isBareMetalServerName: *bms.Name,
		}
		l["id"] = *bms.ID
		l[isBareMetalServerBandwidth] = *bms.Bandwidth
		bmsBootTargetIntf := bms.BootTarget.(*vpcv1.BareMetalServerBootTarget)
		bmsBootTarget := bmsBootTargetIntf.ID
		l[isBareMetalServerBootTarget] = bmsBootTarget
		cpuList := make([]map[string]interface{}, 0)
		if bms.Cpu != nil {
			currentCPU := map[string]interface{}{}
			currentCPU[isBareMetalServerCPUArchitecture] = *bms.Cpu.Architecture
			currentCPU[isBareMetalServerCPUCoreCount] = *bms.Cpu.CoreCount
			currentCPU[isBareMetalServerCpuSocketCount] = *bms.Cpu.SocketCount
			currentCPU[isBareMetalServerCpuThreadPerCore] = *bms.Cpu.ThreadsPerCore
			cpuList = append(cpuList, currentCPU)
		}
		l[isBareMetalServerCPU] = cpuList
		l[isBareMetalServerName] = *bms.Name
		l[isBareMetalServerCRN] = *bms.CRN

		// disks

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
		l[isBareMetalServerDisks] = diskList

		if bms.EnableSecureBoot != nil {
			l[isBareMetalServerEnableSecureBoot] = *bms.EnableSecureBoot
		}

		l[isBareMetalServerHref] = *bms.Href
		l[isBareMetalServerMemory] = *bms.Memory
		l[isBareMetalServerProfile] = *bms.Profile.Name
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
				BareMetalServerID: bms.ID,
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
			l[isBareMetalServerPrimaryNetworkInterface] = primaryNicList
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
					BareMetalServerID: bms.ID,
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
		l[isBareMetalServerNetworkInterfaces] = interfacesList

		//disks
		l[isBareMetalServerResourceType] = *bms.ResourceType
		l[isBareMetalServerStatus] = *bms.Status
		if bms.StatusReasons != nil {
			statusReasonsList := make([]map[string]interface{}, 0)
			for _, sr := range bms.StatusReasons {
				currentSR := map[string]interface{}{}
				if sr.Code != nil && sr.Message != nil {
					currentSR[isBareMetalServerStatusReasonsCode] = *sr.Code
					currentSR[isBareMetalServerStatusReasonsMessage] = *sr.Message
					statusReasonsList = append(statusReasonsList, currentSR)
				}
			}
			l[isBareMetalServerStatusReasons] = statusReasonsList
		}
		if bms.TrustedPlatformModule != nil {
			tpmList := make([]map[string]interface{}, 0)
			currentTPM := map[string]interface{}{
				isBareMetalServerTrustedPlatformModuleEnabled: *bms.TrustedPlatformModule.Enabled,
				isBareMetalServerTrustedPlatformModuleMode:    *bms.TrustedPlatformModule.Mode,
			}
			tpmList = append(tpmList, currentTPM)
			l[isBareMetalServerTrustedPlatformModule] = tpmList
		}
		l[isBareMetalServerVPC] = *bms.VPC.ID
		l[isBareMetalServerZone] = *bms.Zone.Name

		tags, err := GetTagsUsingCRN(meta, *bms.CRN)
		if err != nil {
			log.Printf(
				"Error on get of resource bare metal server (%s) tags: %s", d.Id(), err)
		}
		l[isBareMetalServerTags] = tags
		if bms.ResourceGroup != nil {
			l[isBareMetalServerResourceGroup] = *bms.ResourceGroup.ID
		}
		serversInfo = append(serversInfo, l)
	}
	d.SetId(dataSourceIBMISBareMetalServersID(d))
	d.Set(isBareMetalServers, serversInfo)
	return nil
}

// dataSourceIBMISBareMetalServersID returns a reasonable ID for a Bare Metal Servers list.
func dataSourceIBMISBareMetalServersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
