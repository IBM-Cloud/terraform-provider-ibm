// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerProfiles = "profiles"
)

func dataSourceIBMISBareMetalServerProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBareMetalServerProfilesRead,

		Schema: map[string]*schema.Schema{

			isBareMetalServerProfiles: {
				Type:        schema.TypeList,
				Description: "List of BMS profile maps",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						isBareMetalServerProfileName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this bare metal server profile",
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
						isBareMetalServerProfileSIFs: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         resourceIBMVPCHash,
							Description: "An array of flags supported by this bare metal server profile",
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
										Set:         resourceIBMVPCHash,
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
										Set:         resourceIBMVPCHash,
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
													Set:         resourceIBMVPCHash,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBareMetalServerProfilesRead(d *schema.ResourceData, meta interface{}) error {

	err := bmsProfilesList(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func bmsProfilesList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	start := ""
	allrecs := []vpcv1.BareMetalServerProfile{}
	for {
		listBMSProfilesOptions := &vpcv1.ListBareMetalServerProfilesOptions{}
		if start != "" {
			listBMSProfilesOptions.Start = &start
		}
		availableProfiles, response, err := sess.ListBareMetalServerProfiles(listBMSProfilesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Bare Metal Server Profiles %s\n%s", err, response)
		}
		start = GetNext(availableProfiles.Next)
		allrecs = append(allrecs, availableProfiles.Profiles...)
		if start == "" {
			break
		}
	}

	profilesInfo := make([]map[string]interface{}, 0)
	for _, profile := range allrecs {

		l := map[string]interface{}{
			isBareMetalServerProfileName:   *profile.Name,
			isBareMetalServerProfileFamily: *profile.Family,
		}
		l[isBareMetalServerProfileHref] = *profile.Href
		if profile.Bandwidth != nil {
			bwList := make([]map[string]interface{}, 0)
			bw := profile.Bandwidth.(*vpcv1.BareMetalServerProfileBandwidth)
			bandwidth := map[string]interface{}{
				isBareMetalServerProfileType:  *bw.Type,
				isBareMetalServerProfileValue: *bw.Value,
			}
			bwList = append(bwList, bandwidth)
			l[isBareMetalServerProfileBandwidth] = bwList
		}
		if profile.CpuArchitecture != nil {
			caList := make([]map[string]interface{}, 0)
			ca := profile.CpuArchitecture
			architecture := map[string]interface{}{
				isBareMetalServerProfileType:  *ca.Type,
				isBareMetalServerProfileValue: *ca.Value,
			}
			caList = append(caList, architecture)
			l[isBareMetalServerProfileCPUArchitecture] = caList
		}
		if profile.CpuCoreCount != nil {
			ccList := make([]map[string]interface{}, 0)
			cc := profile.CpuCoreCount.(*vpcv1.BareMetalServerProfileCpuCoreCount)
			coreCount := map[string]interface{}{
				isBareMetalServerProfileType:  *cc.Type,
				isBareMetalServerProfileValue: *cc.Value,
			}
			ccList = append(ccList, coreCount)
			l[isBareMetalServerProfileCPUCoreCount] = ccList
		}
		if profile.CpuSocketCount != nil {
			scList := make([]map[string]interface{}, 0)
			sc := profile.CpuSocketCount.(*vpcv1.BareMetalServerProfileCpuSocketCount)
			socketCount := map[string]interface{}{
				isBareMetalServerProfileType:  *sc.Type,
				isBareMetalServerProfileValue: *sc.Value,
			}
			scList = append(scList, socketCount)
			l[isBareMetalServerProfileCPUSocketCount] = scList
		}

		if profile.Memory != nil {
			memList := make([]map[string]interface{}, 0)
			mem := profile.Memory.(*vpcv1.BareMetalServerProfileMemory)
			m := map[string]interface{}{
				isBareMetalServerProfileType:  *mem.Type,
				isBareMetalServerProfileValue: *mem.Value,
			}
			memList = append(memList, m)
			l[isBareMetalServerProfileCPUSocketCount] = memList
		}
		l[isBareMetalServerProfileRT] = *profile.ResourceType
		var siflist []string
		for _, item := range profile.SupportedImageFlags {
			siflist = append(siflist, item)
		}
		l[isBareMetalServerProfileSIFs] = siflist
		if profile.SupportedTrustedPlatformModuleModes != nil {
			list := make([]map[string]interface{}, 0)
			var stpmmlist []string
			for _, item := range profile.SupportedTrustedPlatformModuleModes.Values {
				stpmmlist = append(stpmmlist, item)
			}
			m := map[string]interface{}{
				isBareMetalServerProfileType: *profile.SupportedTrustedPlatformModuleModes.Type,
			}
			m[isBareMetalServerProfileValues] = stpmmlist
			list = append(list, m)
			l[isBareMetalServerProfileSTPMMs] = list
		}
		if profile.OsArchitecture != nil {
			list := make([]map[string]interface{}, 0)
			var valuelist []string
			for _, item := range profile.OsArchitecture.Values {
				valuelist = append(valuelist, item)
			}
			m := map[string]interface{}{
				isBareMetalServerProfileDefault: *profile.OsArchitecture.Default,
				isBareMetalServerProfileType:    *profile.OsArchitecture.Type,
			}
			m[isBareMetalServerProfileValues] = valuelist
			list = append(list, m)
			l[isBareMetalServerProfileOS] = list
		}

		if profile.Disks != nil {
			list := make([]map[string]interface{}, 0)
			for _, disk := range profile.Disks {
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
			l[isBareMetalServerProfileDisks] = list
		}

		profilesInfo = append(profilesInfo, l)
	}
	d.SetId(dataSourceIBMISBMSProfilesID(d))
	d.Set(isBareMetalServerProfiles, profilesInfo)
	return nil
}

// dataSourceIBMISBMSProfilesID returns a reasonable ID for a BMS Profile list.
func dataSourceIBMISBMSProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
