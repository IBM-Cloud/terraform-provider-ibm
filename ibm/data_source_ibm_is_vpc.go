package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func dataSourceIBMISVPC() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISVPCRead,

		Schema: map[string]*schema.Schema{
			isVPCDefaultNetworkACL: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCClassicAccess: {
				Type:     schema.TypeBool,
				Computed: true,
			},

			isVPCName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPCResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCIDefaultSecurityGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group associated with VPC",
			},

			isVPCTags: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      resourceIBMVPCHash,
			},

			isVPCCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
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

			cseSourceAddresses: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud service endpoint IP Address",
						},

						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location info of CSE Address",
						},
					},
				},
			},

			subnetsList: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subent name",
						},

						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet ID",
						},

						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet status",
						},

						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "subnet location",
						},

						totalIPV4AddressCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total IPv4 address count in the subnet",
						},

						availableIPV4AddressCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Available IPv4 address count in the subnet",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := d.Get(isVPCName).(string)
	if userDetails.generation == 1 {
		err := classicVpcGetByName(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := vpcGetByName(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpcGetByName(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	listVpcsOptions := &vpcclassicv1.ListVpcsOptions{}
	vpcs, _, err := sess.ListVpcs(listVpcsOptions)
	if err != nil {
		return err
	}
	for _, vpc := range vpcs.Vpcs {
		if *vpc.Name == name {
			d.SetId(*vpc.ID)
			d.Set("id", *vpc.ID)
			d.Set(isVPCName, *vpc.Name)
			d.Set(isVPCClassicAccess, *vpc.ClassicAccess)
			d.Set(isVPCStatus, *vpc.Status)
			d.Set(isVPCResourceGroup, *vpc.ResourceGroup.ID)
			if vpc.DefaultNetworkAcl != nil {
				d.Set(isVPCDefaultNetworkACL, *vpc.DefaultNetworkAcl.ID)
			} else {
				d.Set(isVPCDefaultNetworkACL, nil)
			}
			if vpc.DefaultSecurityGroup != nil {
				d.Set(isVPCIDefaultSecurityGroup, *vpc.DefaultSecurityGroup.ID)
			} else {
				d.Set(isVPCIDefaultSecurityGroup, nil)
			}
			tags, err := GetTagsUsingCRN(meta, *vpc.Crn)
			if err != nil {
				log.Printf(
					"An error occured during reading of vpc (%s) tags : %s", d.Id(), err)
			}
			d.Set(isVPCTags, tags)
			d.Set(isVPCCRN, *vpc.Crn)

			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc/network/vpcs")
			d.Set(ResourceName, *vpc.Name)
			d.Set(ResourceCRN, *vpc.Crn)
			d.Set(ResourceStatus, *vpc.Status)
			if vpc.ResourceGroup != nil {
				d.Set(ResourceGroupName, *vpc.ResourceGroup.ID)
			}
			//set the cse ip addresses info
			if vpc.CseSourceIps != nil {
				cseSourceIpsList := make([]map[string]interface{}, 0)
				for _, sourceIP := range vpc.CseSourceIps {
					currentCseSourceIp := map[string]interface{}{}
					if sourceIP.Ip != nil {
						currentCseSourceIp["address"] = *sourceIP.Ip.Address
						currentCseSourceIp["zone_name"] = *sourceIP.Zone.Name
						cseSourceIpsList = append(cseSourceIpsList, currentCseSourceIp)
					}
				}
				d.Set(cseSourceAddresses, cseSourceIpsList)
			}
			options := &vpcclassicv1.ListSubnetsOptions{}
			s, response, err := sess.ListSubnets(options)
			if err != nil {
				log.Printf("Error Fetching subnets %s\n%s", err, response)
			} else {
				subnetsInfo := make([]map[string]interface{}, 0)
				for _, subnet := range s.Subnets {
					if *subnet.Vpc.ID == d.Id() {
						l := map[string]interface{}{
							"name":                    *subnet.Name,
							"id":                      *subnet.ID,
							"status":                  *subnet.Status,
							"zone":                    *subnet.Zone.Name,
							totalIPV4AddressCount:     *subnet.TotalIpv4AddressCount,
							availableIPV4AddressCount: *subnet.AvailableIpv4AddressCount,
						}
						subnetsInfo = append(subnetsInfo, l)
					}
				}
				d.Set(subnetsList, subnetsInfo)
			}
			return nil
		}
	}
	return fmt.Errorf("No VPC found with name %s", name)
}
func vpcGetByName(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	vpcs, _, err := sess.ListVpcs(listVpcsOptions)
	if err != nil {
		return err
	}
	for _, vpc := range vpcs.Vpcs {
		if *vpc.Name == name {
			d.SetId(*vpc.ID)
			d.Set("id", *vpc.ID)
			d.Set(isVPCName, *vpc.Name)
			d.Set(isVPCClassicAccess, *vpc.ClassicAccess)
			d.Set(isVPCStatus, *vpc.Status)
			d.Set(isVPCResourceGroup, *vpc.ResourceGroup.ID)
			if vpc.DefaultNetworkAcl != nil {
				d.Set(isVPCDefaultNetworkACL, *vpc.DefaultNetworkAcl.ID)
			} else {
				d.Set(isVPCDefaultNetworkACL, nil)
			}
			if vpc.DefaultSecurityGroup != nil {
				d.Set(isVPCIDefaultSecurityGroup, *vpc.DefaultSecurityGroup.ID)
			} else {
				d.Set(isVPCIDefaultSecurityGroup, nil)
			}
			tags, err := GetTagsUsingCRN(meta, *vpc.Crn)
			if err != nil {
				log.Printf(
					"An error occured during reading of vpc (%s) tags : %s", d.Id(), err)
			}
			d.Set(isVPCTags, tags)
			d.Set(isVPCCRN, *vpc.Crn)

			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc-ext/network/vpcs")
			d.Set(ResourceName, *vpc.Name)
			d.Set(ResourceCRN, *vpc.Crn)
			d.Set(ResourceStatus, *vpc.Status)
			if vpc.ResourceGroup != nil {
				d.Set(ResourceGroupName, *vpc.ResourceGroup.Name)
			}
			//set the cse ip addresses info
			if vpc.CseSourceIps != nil {
				cseSourceIpsList := make([]map[string]interface{}, 0)
				for _, sourceIP := range vpc.CseSourceIps {
					currentCseSourceIp := map[string]interface{}{}
					if sourceIP.Ip != nil {
						currentCseSourceIp["address"] = *sourceIP.Ip.Address
						currentCseSourceIp["zone_name"] = *sourceIP.Zone.Name
						cseSourceIpsList = append(cseSourceIpsList, currentCseSourceIp)
					}
				}
				d.Set(cseSourceAddresses, cseSourceIpsList)
			}
			options := &vpcv1.ListSubnetsOptions{}
			s, response, err := sess.ListSubnets(options)
			if err != nil {
				log.Printf("Error Fetching subnets %s\n%s", err, response)
			} else {
				subnetsInfo := make([]map[string]interface{}, 0)
				for _, subnet := range s.Subnets {
					if *subnet.Vpc.ID == d.Id() {
						l := map[string]interface{}{
							"name":                    *subnet.Name,
							"id":                      *subnet.ID,
							"status":                  *subnet.Status,
							"zone":                    *subnet.Zone.Name,
							totalIPV4AddressCount:     *subnet.TotalIpv4AddressCount,
							availableIPV4AddressCount: *subnet.AvailableIpv4AddressCount,
						}
						subnetsInfo = append(subnetsInfo, l)
					}
				}
				d.Set(subnetsList, subnetsInfo)
			}
			return nil
		}
	}
	return fmt.Errorf("No VPC found with name %s", name)
}
