package ibm

import (
	"fmt"
	"log"
	"reflect"

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
		err := classicVpcList(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := vpcList(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpcList(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	listVpcsOptions := &vpcclassicv1.ListVpcsOptions{}
	vpcs, response, err := sess.ListVpcs(listVpcsOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Listing VPCs : %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
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
			// set the cse ip addresses info
			if vpc.CseSourceIps != nil {
				displaySourceIps := []VPCCSESourceIP{}
				sourceIPs := vpc.CseSourceIps

				for _, sourceIP := range sourceIPs {
					// work around to parse the cse_source_ip data structure from map[string]interface{} type as we define it as any type in swagger.yaml file
					ip, zone := safeGetIPZone(sourceIP)
					if ip == "" {
						continue
					}
					displaySourceIps = append(displaySourceIps, VPCCSESourceIP{
						Address:  ip,
						ZoneName: zone,
					})
				}
				info := flattenCseIPs(displaySourceIps)
				d.Set(cseSourceAddresses, info)
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
func vpcList(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	vpcs, response, err := sess.ListVpcs(listVpcsOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Listing VPCs : %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
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
			// set the cse ip addresses info
			if vpc.CseSourceIps != nil {
				displaySourceIps := []VPCCSESourceIP{}
				sourceIPs := vpc.CseSourceIps

				for _, sourceIP := range sourceIPs {
					// work around to parse the cse_source_ip data structure from map[string]interface{} type as we define it as any type in swagger.yaml file
					ip, zone := safeGetIPZone(sourceIP)
					if ip == "" {
						continue
					}
					displaySourceIps = append(displaySourceIps, VPCCSESourceIP{
						Address:  ip,
						ZoneName: zone,
					})
				}
				info := flattenCseIPs(displaySourceIps)
				d.Set(cseSourceAddresses, info)
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

func getIPZone(cseSourceIP interface{}) (string, string) {
	log.Println("The raw cse_source_ip is", cseSourceIP)
	cseSourceIPType := reflect.TypeOf(cseSourceIP)
	ip := ""
	zone := ""
	switch v := reflect.ValueOf(cseSourceIP); cseSourceIPType.Kind() {
	case reflect.String:
		ip = v.Interface().(string)
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			iterKey := iter.Key()
			if iterKey.Kind() != reflect.String {
				log.Println("cse_source_ip map key is not string, ignore handling")
				continue
			}
			iterVal := iter.Value()
			if iterKey.Interface().(string) == "zone" {
				zoneValEle := iterVal.Elem()
				log.Println("zone element kind is:", zoneValEle.Kind())
				if zoneValEle.Kind() == reflect.Map {
					iterZone := zoneValEle.MapRange()
					for iterZone.Next() {
						zoneIterKey := iterZone.Key()
						if zoneIterKey.Kind() != reflect.String {
							log.Println("cse_source_ip zone map key is not string, ignore handling")
							continue
						}
						zoneIterVal := iterZone.Value()
						if zoneIterKey.Interface().(string) == "name" && zoneIterVal.Elem().Kind() == reflect.String {
							zone = zoneIterVal.Elem().Interface().(string)
						}
					}
				}
			}
			if iterKey.Interface().(string) == "ip" {
				ipValEle := iterVal.Elem()
				log.Println("ip element type is:", ipValEle.Kind())
				if ipValEle.Kind() == reflect.Map {
					iterIP := ipValEle.MapRange()
					for iterIP.Next() {
						ipIterKey := iterIP.Key()
						if ipIterKey.Kind() != reflect.String {
							log.Println("cse_source_ip ip map key is not string, ignore handling")
							continue
						}
						ipIterVal := iterIP.Value()
						if ipIterKey.Interface().(string) == "address" && ipIterVal.Elem().Kind() == reflect.String {
							ip = ipIterVal.Elem().Interface().(string)
						}
					}
				}
				if ipValEle.Kind() == reflect.String {
					ip = ipValEle.Interface().(string)
				}

			}
		}
	default:
		log.Println("unhandled type while parseing CSE source ip field")
	}
	return ip, zone
}

func safeGetIPZone(cseSourceIP interface{}) (string, string) {
	// If there is any unexpected error while parsing the cse_source_ip, we just swallow the panic
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered while getting cse source ip: ", r)
		}
	}()
	ip, zone := getIPZone(cseSourceIP)
	return ip, zone
}

// VPCCSESourceIP ...
type VPCCSESourceIP struct {
	Address  string
	ZoneName string
}

type sourceIPSortedByZoneName []VPCCSESourceIP

func (s sourceIPSortedByZoneName) Len() int {
	return len(s)
}
func (s sourceIPSortedByZoneName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sourceIPSortedByZoneName) Less(i, j int) bool {
	return s[i].ZoneName < s[j].ZoneName
}
