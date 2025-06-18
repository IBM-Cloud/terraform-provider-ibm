// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerProfileName            = "name"
	isBareMetalServerProfileBandwidth       = "bandwidth"
	isBareMetalServerProfileType            = "type"
	isBareMetalServerProfileValue           = "value"
	isBareMetalServerProfileCPUArchitecture = "cpu_architecture"
	isBareMetalServerProfileCPUCoreCount    = "cpu_core_count"
	isBareMetalServerProfileCPUSocketCount  = "cpu_socket_count"
	isBareMetalServerProfileDisks           = "disks"
	isBareMetalServerProfileDiskQuantity    = "quantity"
	isBareMetalServerProfileDiskSize        = "size"
	isBareMetalServerProfileDiskSITs        = "supported_interface_types"
	isBareMetalServerProfileFamily          = "family"
	isBareMetalServerProfileHref            = "href"
	isBareMetalServerProfileMemory          = "memory"
	isBareMetalServerProfileOS              = "os_architecture"
	isBareMetalServerProfileValues          = "values"
	isBareMetalServerProfileDefault         = "default"
	isBareMetalServerProfileRT              = "resource_type"
	isBareMetalServerProfileSIFs            = "supported_image_flags"
	isBareMetalServerProfileSTPMMs          = "supported_trusted_platform_module_modes"
)

func DataSourceIBMIsBareMetalServerProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBMSProfileRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerProfileName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for this bare metal server profile",
			},

			// vni

			"virtual_network_interfaces_supported": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Indicates whether this profile supports virtual network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The value for this profile field.",
						},
					},
				},
			},
			"network_attachment_count": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum value for this profile field.",
						},
						"min": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The minimum value for this profile field.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
					},
				},
			},

			isBareMetalServerProfileFamily: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The product family this bare metal server profile belongs to",
			},
			isBareMetalServerProfileHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this bare metal server profile",
			},
			isBareMetalServerProfileBandwidth: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The total bandwidth (in megabits per second) shared across the network interfaces of a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValue: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"default": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default value for this profile field.",
						},
						"max": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum value for this profile field.",
						},
						"min": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The minimum value for this profile field.",
						},
						"step": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The increment step value for this profile field.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The permitted values for this profile field.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"network_interface_count": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum value for this profile field.",
						},
						"min": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The minimum value for this profile field.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
					},
				},
			},
			"console_types": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The console type configuration for a bare metal server with this profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The console types for a bare metal server with this profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			isBareMetalServerProfileRT: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type for this bare metal server profile",
			},

			isBareMetalServerProfileCPUArchitecture: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CPU architecture for a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValue: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBareMetalServerProfileCPUSocketCount: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The number of CPU sockets for a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValue: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBareMetalServerProfileCPUCoreCount: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CPU core count for a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValue: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},
			isBareMetalServerProfileMemory: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The memory (in gibibytes) for a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValue: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
					},
				},
			},

			isBareMetalServerProfileSTPMMs: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An array of supported trusted platform module (TPM) modes for this bare metal server profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValues: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "The supported trusted platform module (TPM) modes",
						},
					},
				},
			},
			isBareMetalServerProfileOS: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The supported OS architecture(s) for a bare metal server with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileDefault: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default for this profile field",
						},
						isBareMetalServerProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},

						isBareMetalServerProfileValues: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "The supported OS architecture(s) for a bare metal server with this profile",
						},
					},
				},
			},
			isBareMetalServerProfileDisks: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of the bare metal server profile's disks",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerProfileDiskQuantity: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The number of disks of this configuration for a bare metal server with this profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerProfileType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field",
									},
									isBareMetalServerProfileValue: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field",
									},
								},
							},
						},

						isBareMetalServerProfileDiskSize: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The size of the disk in GB (gigabytes)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerProfileType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field",
									},
									isBareMetalServerProfileValue: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field",
									},
								},
							},
						},
						isBareMetalServerProfileDiskSITs: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The disk interface used for attaching the disk.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerProfileDefault: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
									},
									isBareMetalServerProfileType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field",
									},
									isBareMetalServerProfileValues: {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The supported disk interfaces used for attaching the disk",
										Elem:        &schema.Schema{Type: schema.TypeString},
										Set:         flex.ResourceIBMVPCHash,
									},
								},
							},
						},
					},
				},
			},
			"reservation_terms": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The type for this profile field",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The supported committed use terms for a reservation using this profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBMSProfileRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_profile", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerProfileOptions{
		Name: &name,
	}
	bareMetalServerProfile, _, err := sess.GetBareMetalServerProfileWithContext(context, options)
	if err != nil || bareMetalServerProfile == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerProfileWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_profile", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*bareMetalServerProfile.Name)
	if err = d.Set(isBareMetalServerProfileName, *bareMetalServerProfile.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-name").GetDiag()
	}
	if err = d.Set("family", bareMetalServerProfile.Family); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting family: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-family").GetDiag()
	}
	if err = d.Set("href", bareMetalServerProfile.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-href").GetDiag()
	}
	if bareMetalServerProfile.Bandwidth != nil {
		bwList := make([]map[string]interface{}, 0)
		bw := bareMetalServerProfile.Bandwidth.(*vpcv1.BareMetalServerProfileBandwidth)
		bandwidth := map[string]interface{}{}
		if bw.Type != nil {
			bandwidth[isBareMetalServerProfileType] = *bw.Type
		}
		if bw.Value != nil {
			bandwidth[isBareMetalServerProfileValue] = *bw.Value
		}
		if bw.Values != nil && len(bw.Values) > 0 {
			bandwidth[isBareMetalServerProfileValues] = bw.Values
		}
		if bw.Default != nil {
			bandwidth["default"] = flex.IntValue(bw.Default)
		}
		if bw.Max != nil {
			bandwidth["max"] = flex.IntValue(bw.Max)
		}
		if bw.Min != nil {
			bandwidth["min"] = flex.IntValue(bw.Min)
		}
		if bw.Step != nil {
			bandwidth["step"] = flex.IntValue(bw.Step)
		}
		bwList = append(bwList, bandwidth)
		if err = d.Set(isBareMetalServerProfileBandwidth, bwList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bandwidth: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-bandwidth").GetDiag()
		}
	}
	consoleTypes := []map[string]interface{}{}
	if bareMetalServerProfile.ConsoleTypes != nil {
		modelMap, err := dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileConsoleTypesToMap(bareMetalServerProfile.ConsoleTypes)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_profile", "read", "console_types-to-map").GetDiag()
		}
		consoleTypes = append(consoleTypes, modelMap)
	}
	if err = d.Set("console_types", consoleTypes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting console_types: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-console_types").GetDiag()
	}

	networkInterfaceCount := []map[string]interface{}{}
	if bareMetalServerProfile.NetworkInterfaceCount != nil {
		modelMap, err := dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountToMap(bareMetalServerProfile.NetworkInterfaceCount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_profile", "read", "network_interface_count-to-map").GetDiag()
		}
		networkInterfaceCount = append(networkInterfaceCount, modelMap)
	}
	if err = d.Set("network_interface_count", networkInterfaceCount); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interface_count: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-network_interface_count").GetDiag()
	}

	if bareMetalServerProfile.CpuArchitecture != nil {
		caList := make([]map[string]interface{}, 0)
		ca := bareMetalServerProfile.CpuArchitecture
		architecture := map[string]interface{}{
			isBareMetalServerProfileType:  *ca.Type,
			isBareMetalServerProfileValue: *ca.Value,
		}
		caList = append(caList, architecture)

		if err = d.Set(isBareMetalServerProfileCPUArchitecture, caList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cpu_architecture: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-cpu_architecture").GetDiag()
		}
	}
	if bareMetalServerProfile.CpuCoreCount != nil {
		ccList := make([]map[string]interface{}, 0)
		cc := bareMetalServerProfile.CpuCoreCount.(*vpcv1.BareMetalServerProfileCpuCoreCount)
		coreCount := map[string]interface{}{
			isBareMetalServerProfileType:  *cc.Type,
			isBareMetalServerProfileValue: *cc.Value,
		}
		ccList = append(ccList, coreCount)

		if err = d.Set(isBareMetalServerProfileCPUCoreCount, ccList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cpu_core_count: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-cpu_core_count").GetDiag()
		}
	}
	if bareMetalServerProfile.CpuSocketCount != nil {
		scList := make([]map[string]interface{}, 0)
		sc := bareMetalServerProfile.CpuSocketCount.(*vpcv1.BareMetalServerProfileCpuSocketCount)
		socketCount := map[string]interface{}{
			isBareMetalServerProfileType:  *sc.Type,
			isBareMetalServerProfileValue: *sc.Value,
		}
		scList = append(scList, socketCount)
		if err = d.Set(isBareMetalServerProfileCPUSocketCount, scList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cpu_socket_count: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-cpu_socket_count").GetDiag()
		}
	}

	if bareMetalServerProfile.Memory != nil {
		memList := make([]map[string]interface{}, 0)
		mem := bareMetalServerProfile.Memory.(*vpcv1.BareMetalServerProfileMemory)
		m := map[string]interface{}{
			isBareMetalServerProfileType:  *mem.Type,
			isBareMetalServerProfileValue: *mem.Value,
		}
		memList = append(memList, m)
		if err = d.Set(isBareMetalServerProfileMemory, memList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting memory: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-memory").GetDiag()
		}
	}
	if err = d.Set(isBareMetalServerProfileRT, bareMetalServerProfile.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-resource_type").GetDiag()
	}
	if bareMetalServerProfile.SupportedTrustedPlatformModuleModes != nil {
		list := make([]map[string]interface{}, 0)
		var stpmmlist []string
		for _, item := range bareMetalServerProfile.SupportedTrustedPlatformModuleModes.Values {
			stpmmlist = append(stpmmlist, item)
		}
		m := map[string]interface{}{
			isBareMetalServerProfileType: *bareMetalServerProfile.SupportedTrustedPlatformModuleModes.Type,
		}
		m[isBareMetalServerProfileValues] = stpmmlist
		list = append(list, m)
		if err = d.Set(isBareMetalServerProfileSTPMMs, list); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting supported_trusted_platform_module_modes: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-supported_trusted_platform_module_modes").GetDiag()
		}
	}
	if bareMetalServerProfile.OsArchitecture != nil {
		list := make([]map[string]interface{}, 0)
		var valuelist []string
		for _, item := range bareMetalServerProfile.OsArchitecture.Values {
			valuelist = append(valuelist, item)
		}
		m := map[string]interface{}{
			isBareMetalServerProfileDefault: *bareMetalServerProfile.OsArchitecture.Default,
			isBareMetalServerProfileType:    *bareMetalServerProfile.OsArchitecture.Type,
		}
		m[isBareMetalServerProfileValues] = valuelist
		list = append(list, m)
		if err = d.Set(isBareMetalServerProfileOS, list); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting os_architecture: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-os_architecture").GetDiag()
		}
	}

	if bareMetalServerProfile.Disks != nil {
		list := make([]map[string]interface{}, 0)
		for _, disk := range bareMetalServerProfile.Disks {
			qlist := make([]map[string]interface{}, 0)
			slist := make([]map[string]interface{}, 0)
			sitlist := make([]map[string]interface{}, 0)
			quantity := disk.Quantity.(*vpcv1.BareMetalServerProfileDiskQuantity)
			q := make(map[string]interface{})
			q[isBareMetalServerProfileType] = *quantity.Type
			q[isBareMetalServerProfileValue] = *quantity.Value
			qlist = append(qlist, q)
			size := disk.Size.(*vpcv1.BareMetalServerProfileDiskSize)
			s := map[string]interface{}{
				isBareMetalServerProfileType:  *size.Type,
				isBareMetalServerProfileValue: *size.Value,
			}
			slist = append(slist, s)
			sit := map[string]interface{}{
				isBareMetalServerProfileDefault: *disk.SupportedInterfaceTypes.Default,
				isBareMetalServerProfileType:    *disk.SupportedInterfaceTypes.Type,
			}
			var valuelist []string
			for _, item := range disk.SupportedInterfaceTypes.Values {
				valuelist = append(valuelist, item)
			}
			sit[isBareMetalServerProfileValues] = valuelist
			sitlist = append(sitlist, sit)
			sz := map[string]interface{}{
				isBareMetalServerProfileDiskQuantity: qlist,
				isBareMetalServerProfileDiskSize:     slist,
				isBareMetalServerProfileDiskSITs:     sitlist,
			}
			list = append(list, sz)
		}
		if err = d.Set(isBareMetalServerProfileDisks, list); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting disks: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-disks").GetDiag()
		}
		// vni
		virtualNetworkInterfacesSupported := []map[string]interface{}{}
		if bareMetalServerProfile.VirtualNetworkInterfacesSupported != nil {
			modelMap, err := dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileVirtualNetworkInterfacesSupportedToMap(bareMetalServerProfile.VirtualNetworkInterfacesSupported)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_profile", "read", "virtual_network_interfaces_supported-to-map").GetDiag()
			}
			virtualNetworkInterfacesSupported = append(virtualNetworkInterfacesSupported, modelMap)
		}
		if err = d.Set("virtual_network_interfaces_supported", virtualNetworkInterfacesSupported); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting virtual_network_interfaces_supported: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-virtual_network_interfaces_supported").GetDiag()
		}
		networkAttachmentCount := []map[string]interface{}{}
		if bareMetalServerProfile.NetworkAttachmentCount != nil {
			modelMap, err := dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountToMap(bareMetalServerProfile.NetworkAttachmentCount)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_profile", "read", "network_attachment_count-to-map").GetDiag()
			}
			networkAttachmentCount = append(networkAttachmentCount, modelMap)
		}
		if err = d.Set("network_attachment_count", networkAttachmentCount); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachment_count: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-network_attachment_count").GetDiag()
		}
	}
	if bareMetalServerProfile.ReservationTerms != nil {
		err = d.Set("reservation_terms", dataSourceBaremetalServerProfileFlattenReservationTerms(*bareMetalServerProfile.ReservationTerms))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reservation_terms: %s", err), "(Data) ibm_is_bare_metal_server_profile", "read", "set-reservation_terms").GetDiag()
		}
	}

	return nil
}

func dataSourceBaremetalServerProfileFlattenReservationTerms(result vpcv1.BareMetalServerProfileReservationTerms) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceBaremetalServerProfileReservationTermsToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceBaremetalServerProfileReservationTermsToMap(resTermItem vpcv1.BareMetalServerProfileReservationTerms) map[string]interface{} {
	resTermMap := map[string]interface{}{}

	if resTermItem.Type != nil {
		resTermMap["type"] = resTermItem.Type
	}
	if resTermItem.Values != nil {
		resTermMap["values"] = resTermItem.Values
	}

	return resTermMap
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileConsoleTypesToMap(model *vpcv1.BareMetalServerProfileConsoleTypes) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["values"] = model.Values
	return modelMap, nil
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountToMap(model vpcv1.BareMetalServerProfileNetworkInterfaceCountIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCountRange); ok {
		return dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountRangeToMap(model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCountRange))
	} else if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCountDependent); ok {
		return dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountDependentToMap(model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCountDependent))
	} else if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCount); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.BareMetalServerProfileNetworkInterfaceCount)
		if model.Max != nil {
			modelMap["max"] = flex.IntValue(model.Max)
		}
		if model.Min != nil {
			modelMap["min"] = flex.IntValue(model.Min)
		}
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.BareMetalServerProfileNetworkInterfaceCountIntf subtype encountered")
	}
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountRangeToMap(model *vpcv1.BareMetalServerProfileNetworkInterfaceCountRange) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Max != nil {
		modelMap["max"] = flex.IntValue(model.Max)
	}
	if model.Min != nil {
		modelMap["min"] = flex.IntValue(model.Min)
	}
	modelMap["type"] = model.Type
	return modelMap, nil
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkInterfaceCountDependentToMap(model *vpcv1.BareMetalServerProfileNetworkInterfaceCountDependent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	return modelMap, nil
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileVirtualNetworkInterfacesSupportedToMap(model *vpcv1.BareMetalServerProfileVirtualNetworkInterfacesSupported) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountToMap(model vpcv1.BareMetalServerProfileNetworkAttachmentCountIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCountRange); ok {
		return dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountRangeToMap(model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCountRange))
	} else if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCountDependent); ok {
		return dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountDependentToMap(model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCountDependent))
	} else if _, ok := model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCount); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.BareMetalServerProfileNetworkAttachmentCount)
		if model.Max != nil {
			modelMap["max"] = flex.IntValue(model.Max)
		}
		if model.Min != nil {
			modelMap["min"] = flex.IntValue(model.Min)
		}
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.BareMetalServerProfileNetworkAttachmentCountIntf subtype encountered")
	}
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountRangeToMap(model *vpcv1.BareMetalServerProfileNetworkAttachmentCountRange) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Max != nil {
		modelMap["max"] = flex.IntValue(model.Max)
	}
	if model.Min != nil {
		modelMap["min"] = flex.IntValue(model.Min)
	}
	modelMap["type"] = model.Type
	return modelMap, nil
}

func dataSourceIBMIsBareMetalServerProfileBareMetalServerProfileNetworkAttachmentCountDependentToMap(model *vpcv1.BareMetalServerProfileNetworkAttachmentCountDependent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	return modelMap, nil
}
