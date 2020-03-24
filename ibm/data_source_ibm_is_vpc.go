package ibm

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
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
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
		},
	}
}

func dataSourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)

	name := d.Get(isVPCName).(string)

	vpcs, _, err := vpcC.List("")
	if err != nil {
		return err
	}

	for _, vpc := range vpcs {
		if vpc.Name == name {
			d.SetId(vpc.ID.String())
			d.Set("id", vpc.ID.String())
			d.Set(isVPCName, vpc.Name)
			d.Set(isVPCClassicAccess, vpc.ClassicAccess)
			d.Set(isVPCStatus, vpc.Status)
			d.Set(isVPCResourceGroup, vpc.ResourceGroup.ID)
			if vpc.DefaultNetworkACL != nil {
				d.Set(isVPCDefaultNetworkACL, vpc.DefaultNetworkACL.ID)
			} else {
				d.Set(isVPCDefaultNetworkACL, nil)
			}
			tags, err := GetTagsUsingCRN(meta, vpc.Crn)
			if err != nil {
				return fmt.Errorf(
					"Error on get of resource vpc (%s) tags: %s", d.Id(), err)
			}
			d.Set(isVPCTags, tags)

			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			if sess.Generation == 1 {
				d.Set(ResourceControllerURL, controller+"/vpc/network/vpcs")
			} else {
				d.Set(ResourceControllerURL, controller+"/vpc-ext/network/vpcs")
			}
			d.Set(ResourceName, vpc.Name)
			d.Set(ResourceCRN, vpc.Crn)
			d.Set(ResourceStatus, vpc.Status)
			if vpc.ResourceGroup != nil {
				d.Set(ResourceGroupName, vpc.ResourceGroup.Name)
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
			return nil
		}
	}
	return fmt.Errorf("No VPC found with name %s", name)
}

func getIPZone(cseSourceIP interface{}) (string, string) {
	fmt.Println("The raw cse_source_ip is", cseSourceIP)
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
				fmt.Println("cse_source_ip map key is not string, ignore handling")
				continue
			}
			iterVal := iter.Value()
			if iterKey.Interface().(string) == "zone" {
				zoneValEle := iterVal.Elem()
				fmt.Println("zone element kind is:", zoneValEle.Kind())
				if zoneValEle.Kind() == reflect.Map {
					iterZone := zoneValEle.MapRange()
					for iterZone.Next() {
						zoneIterKey := iterZone.Key()
						if zoneIterKey.Kind() != reflect.String {
							fmt.Println("cse_source_ip zone map key is not string, ignore handling")
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
				fmt.Println("ip element type is:", ipValEle.Kind())
				if ipValEle.Kind() == reflect.Map {
					iterIP := ipValEle.MapRange()
					for iterIP.Next() {
						ipIterKey := iterIP.Key()
						if ipIterKey.Kind() != reflect.String {
							fmt.Println("cse_source_ip ip map key is not string, ignore handling")
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
		fmt.Println("unhandled type while parseing CSE source ip field")
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
