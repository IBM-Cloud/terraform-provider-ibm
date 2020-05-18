package ibm

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isInstanceTemplates                     = "templates"
	isInstanceTemplatesFirst                = "first"
	isInstanceTemplatesHref                 = "href"
	isInstanceTemplatesCrn                  = "crn"
	isInstanceTemplatesLimit                = "limit"
	isInstanceTemplatesNext                 = "next"
	isInstanceTemplatesTotalCount           = "total_count"
	isInstanceTemplatesName                 = "name"
	isInstanceTemplatesPortSpeed            = "port_speed"
	isInstanceTemplatesPortType             = "type"
	isInstanceTemplatesPortValue            = "value"
	isInstanceTemplatesDeleteVol            = "delete_volume_on_instance_delete"
	isInstanceTemplatesVol                  = "volume"
	isInstanceTemplatesMemory               = "memory"
	isInstanceTemplatesMemoryValue          = "value"
	isInstanceTemplatesMemoryType           = "type"
	isInstanceTemplatesMemoryValues         = "values"
	isInstanceTemplatesMemoryDefault        = "default"
	isInstanceTemplatesMemoryMin            = "min"
	isInstanceTemplatesMemoryMax            = "max"
	isInstanceTemplatesMemoryStep           = "step"
	isInstanceTemplatesSocketCount          = "socket_count"
	isInstanceTemplatesSocketValue          = "value"
	isInstanceTemplatesSocketType           = "type"
	isInstanceTemplatesSocketValues         = "values"
	isInstanceTemplatesSocketDefault        = "default"
	isInstanceTemplatesSocketMin            = "min"
	isInstanceTemplatesSocketMax            = "max"
	isInstanceTemplatesSocketStep           = "step"
	isInstanceTemplatesVcpuArch             = "vcpu_architecture"
	isInstanceTemplatesVcpuArchType         = "type"
	isInstanceTemplatesVcpuArchValue        = "value"
	isInstanceTemplatesVcpuCount            = "vcpu_count"
	isInstanceTemplatesVcpuCountValue       = "value"
	isInstanceTemplatesVcpuCountType        = "type"
	isInstanceTemplatesVcpuCountValues      = "values"
	isInstanceTemplatesVcpuCountDefault     = "default"
	isInstanceTemplatesVcpuCountMin         = "min"
	isInstanceTemplatesVcpuCountMax         = "max"
	isInstanceTemplatesVcpuCountStep        = "step"
	isInstanceTemplatesStart                = "start"
	isInstanceTemplatesVersion              = "version"
	isInstanceTemplatesGeneration           = "generation"
	isInstanceTemplatesNicSecondaryAddr     = "secondary_addresses"
	isInstanceTemplatesBootVolumeAttachment = "boot_volume_attachment"
)

func dataSourceIBMISInstanceTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceTemplatesRead,
		Schema: map[string]*schema.Schema{
			isInstanceTemplatesFirst: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplatesHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			isInstanceTemplatesLimit: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			isInstanceTemplatesNext: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplatesHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			isInstanceTemplatesTotalCount: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			isInstanceTemplates: {
				Type:        schema.TypeList,
				Description: "Collection of instance templates",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplatesCrn: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateVPC: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateZone: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplatesName: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateProfile: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplatesName: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateKeys: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						isInstanceTemplateVolumeAttachments: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplatesDeleteVol: {
										Type:     schema.TypeBool,
										Computed: true,
									},
									isInstanceTemplatesName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceTemplatesVol: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplatePrimaryNetworkInterface: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									/*
										isInstanceTemplateNicAllowIpSpoofing: {
											Type:     schema.TypeBool,
											Optional: true,
											Default:  false,
										},
									*/
									isInstanceTemplateNicName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceTemplateNicPrimaryIpv4Address: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceNicSecurityGroups: {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									isInstanceNicSubnet: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateNetworkInterfaces: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									/*
										isInstanceTemplateNicAllowIpSpoofing: {
											Type:     schema.TypeBool,
											Optional: true,
											Default:  false,
										},
									*/
									isInstanceTemplateNicName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceTemplateNicPrimaryIpv4Address: {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									isInstanceNicSecurityGroups: {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									isInstanceNicSubnet: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateGeneration: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateUserData: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateImage: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplatesBootVolumeAttachment: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplatesDeleteVol: {
										Type:     schema.TypeBool,
										Computed: true,
									},
									isInstanceTemplatesName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceTemplatesVol: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						isInstanceTemplateResourceGroup: {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									isInstanceTemplatesName: {
										Type:     schema.TypeString,
										Computed: true,
									},
									isInstanceTemplatesHref: {
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceIBMISInstanceTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	instanceC, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listInstanceTemplatesOptions := &vpcv1.ListInstanceTemplatesOptions{}
	availableTemplates, _, err := instanceC.ListInstanceTemplates(listInstanceTemplatesOptions)
	if err != nil {
		return err
	}
	firstnode := map[string]interface{}{}
	firstnode["href"] = *availableTemplates.First.Href
	log.Println("Href of instance templates:", *availableTemplates.First.Href)
	log.Println("firstnode of instance templates:", firstnode)
	d.Set(isInstanceTemplatesFirst, firstnode)
	d.Set(isInstanceTemplatesLimit, availableTemplates.Limit)
	log.Println("limit of instance templates:", availableTemplates.Limit)
	d.Set(isInstanceTemplatesTotalCount, availableTemplates.TotalCount)
	log.Println("total count of instance templates:", availableTemplates.TotalCount)
	log.Println("count of instance templates:", len(availableTemplates.Templates))
	// len_available_profiles := len(availableProfiles.Profiles)
	templates := make([]map[string]interface{}, 0)
	for _, instTempl := range availableTemplates.Templates {
		template := map[string]interface{}{}
		instance := instTempl.(*vpcv1.InstanceTemplate)
		template["id"] = instance.ID
		template[isInstanceTemplatesHref] = instance.Href
		template[isInstanceTemplatesCrn] = instance.Crn
		template[isInstanceTemplateName] = instance.Name
		template[isInstanceTemplateUserData] = instance.UserData

		if instance.Profile != nil {
			identityNode := map[string]interface{}{}
			instanceProfileIntf := instance.Profile
			identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
			identityNode[isInstanceTemplateName] = identity.Name
			template[isInstanceTemplateProfile] = identityNode
		}
		if instance.PrimaryNetworkInterface != nil {
			currentPrimNic := map[string]interface{}{}
			currentPrimNic[isInstanceTemplateNicName] = *instance.PrimaryNetworkInterface.Name
			log.Println("Primary Network Interface name:", *instance.PrimaryNetworkInterface.Name)
			if instance.PrimaryNetworkInterface.PrimaryIpv4Address != nil {
				currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *instance.PrimaryNetworkInterface.PrimaryIpv4Address
				log.Println("Primary Network Interface ipv4:", *instance.PrimaryNetworkInterface.PrimaryIpv4Address)
			}
			subInf := instance.PrimaryNetworkInterface.Subnet
			subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
			log.Println("Primary Network Interface subnet id:", *subnetIdentity.ID)
			currentPrimNic[isInstanceNicSubnet] = *subnetIdentity.ID

			//currentPrimNic[isInstanceTemplateNicAllowIpSpoofing] = instance.PrimaryNetworkInterface.AllowIpSpoofing
			/*
				if len(instance.PrimaryNetworkInterface.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(instance.PrimaryNetworkInterface.SecurityGroups); i++ {
						secGrpInf := instance.PrimaryNetworkInterface.SecurityGroups[i]
						secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
						secGrp := map[string]interface{}{}
						log.Println("Primary Network Interface secgrp id:", *secGrpIdentity.ID)
						secGrp["id"] = *secGrpIdentity.ID
						secGrp[isInstanceTemplatesCrn] = *secGrpIdentity.Crn
						log.Println("Primary Network Interface secgrp Crn:", *secGrpIdentity.Crn)
						secGrp[isInstanceTemplatesHref] = *secGrpIdentity.Href
						log.Println("Primary Network Interface secgrp Href:", *secGrpIdentity.Href)
						secgrpList = append(secgrpList, *secGrpIdentity.ID)
					}
					currentPrimNic[isInstanceTemplateNicSecurityGroups] = secgrpList
					log.Println("Primary Network Interface ", currentPrimNic)
				}
			*/

			if len(instance.PrimaryNetworkInterface.SecurityGroups) > 0 {
				secgrpList := []string{}
				for i := 0; i < len(instance.PrimaryNetworkInterface.SecurityGroups); i++ {
					secGrpInf := instance.PrimaryNetworkInterface.SecurityGroups[i]
					secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
					secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
					log.Println("Primary Network Interface secgrp id:", string(*secGrpIdentity.ID))
				}
				//currentPrimNic[isInstanceNicSecurityGroups] = strings.Join(secgrpList, ",")
				currentPrimNic[isInstanceNicSecurityGroups] = secgrpList
				log.Println("Primary Network Interface secgrp id:", secgrpList)
			}

			/*
				secgrpList := []string{}
				for i := 0; i < len(instance.PrimaryNetworkInterface.SecurityGroups); i++ {
					secGrpInf := instance.PrimaryNetworkInterface.SecurityGroups[i]
					secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
					secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
					log.Println("Primary Network Interface secgrp id:", string(*secGrpIdentity.ID))
				}
				currentPrimNic[isInstanceNicSecurityGroups] = strings.Join(secgrpList, ",")

			*/
			/*
				if len(instance.PrimaryNetworkInterface.SecurityGroups) > 0 {
					secGrpsData, err := json.Marshal(instance.PrimaryNetworkInterface.SecurityGroups)
					if err != nil {
						log.Printf("Error reading security groups of Primary Network Interface :%s", err)
						return err
					}
					jsonStr := string(secGrpsData)
					currentPrimNic[isInstanceNicSecurityGroups] = jsonStr
					log.Println("Primary Network Interface secgrp id:", jsonStr)
				}
			*/
			template[isInstanceTemplatePrimaryNetworkInterface] = currentPrimNic
			log.Println("Primary Network Interface template", template)
		}
		/*
			if instance.NetworkInterfaces != nil {
				interfacesList := make([]map[string]interface{}, 0)
				log.Println("count of Network Interfaces ", len(instance.NetworkInterfaces))
				for _, intfc := range instance.NetworkInterfaces {
					currentNic := map[string]interface{}{}
					currentNic[isInstanceTemplateNicName] = *intfc.Name
					log.Println("Network Interface name:", *intfc.Name)
					if intfc.PrimaryIpv4Address != nil {
						currentNic[isInstanceTemplateNicPrimaryIpv4Address] = *intfc.PrimaryIpv4Address
						log.Println("Network Interface ipv4:", *instance.PrimaryNetworkInterface.PrimaryIpv4Address)
					}
					//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
					subInf := intfc.Subnet
					subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
					currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
					log.Println("Network Interface subnet id:", *subnetIdentity.ID)
					/*
						if len(intfc.SecurityGroups) != 0 {
							secgrpList := make([]map[string]interface{}, 0)
							for i := 0; i < len(intfc.SecurityGroups); i++ {
								secGrpInf := intfc.SecurityGroups[i]
								secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
								secGrp := map[string]interface{}{}
								secGrp["id"] = *secGrpIdentity.ID
								secGrp[isInstanceTemplatesCrn] = *secGrpIdentity.Crn
								secGrp[isInstanceTemplatesHref] = *secGrpIdentity.Href
								secgrpList = append(secgrpList, secGrp)
							}
							currentNic[isInstanceTemplateNicSecurityGroups] = secgrpList
						}
		*/
		/*
					if len(intfc.SecurityGroups) > 0 {
						secGrpsData, err := json.Marshal(intfc.SecurityGroups)
						if err != nil {
							log.Printf("Error reading security groups of Network Interface :%s", err)
							return err
						}
						jsonStr := string(secGrpsData)
						currentNic[isInstanceNicSecurityGroups] = jsonStr
						log.Println("Network Interface secgrp id:", jsonStr)
					}
					interfacesList = append(interfacesList, currentNic)
				}
				template[isInstanceTemplateNetworkInterfaces] = interfacesList
			}
		*/
		if instance.Image != nil {
			imageNode := map[string]interface{}{}
			imageInf := instance.Image
			imageIdentity := imageInf.(*vpcv1.ImageIdentity)
			imageNode["id"] = imageIdentity.ID
			template[isInstanceTemplateImage] = imageNode
		}

		if instance.Vpc != nil {
			vpcNode := map[string]interface{}{}
			vpcInf := instance.Vpc
			vpcRef := vpcInf.(*vpcv1.VPCIdentity)
			vpcNode["id"] = vpcRef.ID
			template[isInstanceTemplateVPC] = vpcNode
		}

		if instance.Zone != nil {
			zoneNode := map[string]interface{}{}
			zoneInf := instance.Zone
			zone := zoneInf.(*vpcv1.ZoneIdentity)
			zoneNode[isInstanceTemplateName] = zone.Name
			template[isInstanceTemplateZone] = zoneNode
		}
		/*
			interfacesList := make([]map[string]interface{}, 0)
			if instance.VolumeAttachments != nil {
				for _, volume := range instance.VolumeAttachments {
					volumeAttach := map[string]interface{}{}
					volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
					volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
					/*
						volumeId := map[string]interface{}{}
						volumeIntf := volume.Volume
						volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeInstanceContextVolume)
						volumeId[isInstanceTemplateVolAttVolName] = volumeInst.Name
						volumeId[isInstanceTemplateVolAttVolIops] = volumeInst.Iops

					volumeIntf := volume.Volume
					volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeInstanceContextVolume)
					volData, err := json.Marshal(volumeInst)
					if err != nil {
						log.Printf("Error reading volume of volume attachments :%s", err)
						return err
					}
					jsonStr := string(volData)
					volumeAttach[isInstanceTemplateVolAttVolume] = jsonStr
					interfacesList = append(interfacesList, volumeAttach)
				}
				template[isInstanceTemplateVolumeAttachments] = interfacesList
			}
		*/
		/*
			if instance.BootVolumeAttachment != nil {
				bootVol := map[string]interface{}{}
				bootVol[isInstanceTemplateBootName] = *instance.BootVolumeAttachment.Name
				log.Println("Type of delete ", reflect.TypeOf(instance.BootVolumeAttachment.DeleteVolumeOnInstanceDelete).String())
				// bootVol[isInstanceTemplateDeleteVolume] = *instance.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
				log.Println("Boot Volume name :", *instance.BootVolumeAttachment.Name)
				log.Println("Boot Volume delete :", *instance.BootVolumeAttachment.DeleteVolumeOnInstanceDelete)

				volumeIntf := instance.BootVolumeAttachment.Volume
				volData, err := json.Marshal(volumeIntf)
				if err != nil {
					log.Printf("Error reading volume of boot volume attachment :%s", err)
					return err
				}
				jsonStr := string(volData)
				bootVol[isInstanceTemplateVolAttVolume] = jsonStr
				log.Println("Boot Volume json :", jsonStr)
				log.Println("Boot Volume  :", bootVol)

				// getvolattoptions := &vpcclassicv1.GetVolumeAttachmentOptions{
				// 	InstanceID: &ID,
				// 	ID:         instance.BootVolumeAttachment.Volume.ID,
				// }
				// vol, _, err := instanceC.GetVolumeAttachment(getvolattoptions)
				// if err != nil {
				// 	return fmt.Errorf("Error while retrieving boot volume %s for instance %s: %v", getvolattoptions.ID, d.Id(), err)
				// }

				// bootVol[isInstanceBootSize] = instance.BootVolumeAttachment.Capacity
				// bootVol[isInstanceBootIOPS] = instance.BootVolumeAttachment.Iops
				// bootVol[isInstanceBootProfile] = instance.BootVolumeAttachment.Name

				template[isInstanceTemplatesBootVolumeAttachment] = bootVol
			}
		*/
		if instance.ResourceGroup != nil {
			rgNode := map[string]interface{}{}
			rg := instance.ResourceGroup
			rgNode["id"] = rg.ID
			rgNode[isInstanceTemplateName] = rg.Name
			rgNode[isInstanceTemplatesHref] = rg.Href
			template[isInstanceTemplateResourceGroup] = rgNode
		}

		templates = append(templates, template)
	}
	d.SetId(dataSourceIBMISInstanceTemplatesID(d))
	d.Set(isInstanceTemplates, templates)
	return nil
}

// dataSourceIBMISInstanceTemplatesID returns a reasonable ID for a instance templates list.
func dataSourceIBMISInstanceTemplatesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
