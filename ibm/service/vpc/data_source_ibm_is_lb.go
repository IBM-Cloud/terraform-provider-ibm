// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	name                      = "name"
	poolAlgorithm             = "algorithm"
	href                      = "href"
	family                    = "family"
	poolProtocol              = "protocol"
	poolCreatedAt             = "created_at"
	poolProvisioningStatus    = "provisioning_status"
	healthMonitor             = "health_monitor"
	instanceGroup             = "instance_group"
	members                   = "members"
	sessionPersistence        = "session_persistence"
	crnInstance               = "crn"
	sessionType               = "type"
	healthMonitorType         = "type"
	healthMonitorDelay        = "delay"
	healthMonitorMaxRetries   = "max_retries"
	healthMonitorPort         = "port"
	healthMonitorTimeout      = "timeout"
	healthMonitorURLPath      = "url_path"
	isLBPrivateIPDetail       = "private_ip"
	isLBPrivateIpAddress      = "address"
	isLBPrivateIpHref         = "href"
	isLBPrivateIpName         = "name"
	isLBPrivateIpId           = "reserved_ip"
	isLBPrivateIpResourceType = "resource_type"
)

func DataSourceIBMISLB() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISLBRead,

		Schema: map[string]*schema.Schema{
			isLBName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Load Balancer name",
			},
			isLBAccessMode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The access mode of this load balancer",
			},
			"dns": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The DNS configuration for this load balancer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this DNS instance",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the DNS zone.",
						},
					},
				},
			},
			isLBType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer type",
			},
			isLBAvailability: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The availability of this load balancer",
			},
			isLBInstanceGroupsSupported: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this load balancer supports instance groups.",
			},
			isLBSourceIPPersistenceSupported: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this load balancer supports source IP session persistence.",
			},
			isLBUdpSupported: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this load balancer supports UDP.",
			},

			isLBStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer status",
			},
			isLbProfile: {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The profile to use for this load balancer",
			},
			isLBRouteMode: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether route mode is enabled for this load balancer",
			},

			isLBCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this Load Balancer",
			},

			isLBOperatingStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer operating status",
			},

			isLBPublicIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Load Balancer Public IPs",
			},

			isLBPrivateIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Load Balancer private IPs",
			},

			isLBSubnets: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer subnets list",
			},

			isLBSecurityGroups: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer securitygroups list",
			},

			isLBSecurityGroupsSupported: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Security Group Supported for this Load Balancer",
			},

			isLBTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Tags associated to Load Balancer",
			},

			isLBAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},

			isLBResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer Resource group",
			},

			isLBHostName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer Host Name",
			},

			isLBLogging: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Logging of Load Balancer",
			},

			isLBListeners: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer Listeners list",
			},
			isLBPools: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Load Balancer Pools list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						poolAlgorithm: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The load balancing algorithm.",
						},
						healthMonitor: {
							Description: "The health monitor of this pool.",
							Computed:    true,
							Type:        schema.TypeMap,
						},

						instanceGroup: {
							Description: "The instance group that is managing this pool.",
							Computed:    true,
							Type:        schema.TypeMap,
						},

						members: {
							Description: "The backend server members of the pool.",
							Computed:    true,
							Type:        schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									href: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The member's canonical URL.",
									},
									ID: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this load balancer pool member.",
									},
								},
							},
						},
						sessionPersistence: {
							Description: "The session persistence of this pool.",
							Computed:    true,
							Type:        schema.TypeMap,
						},
						poolCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this pool was created.",
						},
						href: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The pool's canonical URL.",
						},
						ID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this load balancer pool",
						},
						name: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this load balancer pool",
						},
						poolProtocol: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol used for this load balancer pool.",
						},
						poolProvisioningStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The provisioning status of this pool.",
						},
					},
				},
			},
			isLBPrivateIPDetail: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The private IP addresses assigned to this load balancer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isLBPrivateIpAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address to reserve, which must not already be reserved on the subnet.",
						},
						isLBPrivateIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isLBPrivateIpName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
						},
						isLBPrivateIpId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a reserved IP by a unique property.",
						},
						isLBPrivateIpResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type",
						},
					},
				},
			},
			"failsafe_policy_actions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The supported `failsafe_policy.action` values for this load balancer's pools.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func dataSourceIBMISLBRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get(isLBName).(string)
	err := lbGetByName(context, d, meta, name)
	if err != nil {
		return err
	}
	return nil
}

func lbGetByName(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	start := ""
	allrecs := []vpcv1.LoadBalancer{}
	for {
		listLoadBalancersOptions := &vpcv1.ListLoadBalancersOptions{}
		if start != "" {
			listLoadBalancersOptions.Start = &start
		}
		lbs, _, err := sess.ListLoadBalancersWithContext(context, listLoadBalancersOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListLoadBalancersWithContext failed: %s", err.Error()), "(Data) ibm_is_lb", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(lbs.Next)
		allrecs = append(allrecs, lbs.LoadBalancers...)
		if start == "" {
			break
		}
	}

	for _, loadBalancer := range allrecs {
		if *loadBalancer.Name == name {
			d.SetId(*loadBalancer.ID)
			if err = d.Set("availability", loadBalancer.Availability); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability: %s", err), "(Data) ibm_is_lb", "read", "set-availability").GetDiag()
			}
			if err = d.Set("access_mode", loadBalancer.AccessMode); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_mode: %s", err), "(Data) ibm_is_lb", "read", "set-access_mode").GetDiag()
			}
			if err = d.Set("instance_groups_supported", loadBalancer.InstanceGroupsSupported); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_groups_supported: %s", err), "(Data) ibm_is_lb", "read", "set-instance_groups_supported").GetDiag()
			}
			if err = d.Set("source_ip_session_persistence_supported", loadBalancer.SourceIPSessionPersistenceSupported); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_ip_session_persistence_supported: %s", err), "(Data) ibm_is_lb", "read", "set-source_ip_session_persistence_supported").GetDiag()
			}
			dnsList := make([]map[string]interface{}, 0)
			if loadBalancer.Dns != nil {
				dns := map[string]interface{}{}
				dns["instance_crn"] = loadBalancer.Dns.Instance.CRN
				dns["zone_id"] = loadBalancer.Dns.Zone.ID
				dnsList = append(dnsList, dns)
				if err = d.Set("dns", dnsList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dns: %s", err), "(Data) ibm_is_lb", "read", "set-dns").GetDiag()
				}
			}
			if err = d.Set("name", loadBalancer.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_lb", "read", "set-name").GetDiag()
			}
			if loadBalancer.Logging != nil && loadBalancer.Logging.Datapath != nil {
				if err = d.Set(isLBLogging, *loadBalancer.Logging.Datapath.Active); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logging: %s", err), "(Data) ibm_is_lb", "read", "set-logging").GetDiag()
				}
			}
			if loadBalancer.IsPublic != nil && *loadBalancer.IsPublic {
				if err = d.Set(isLBType, "public"); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_lb", "read", "set-type").GetDiag()
				}
			} else if loadBalancer.IsPrivatePath != nil && *loadBalancer.IsPrivatePath {
				if err = d.Set(isLBType, "private_path"); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_lb", "read", "set-type").GetDiag()
				}
			} else {
				if err = d.Set(isLBType, "private"); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_lb", "read", "set-type").GetDiag()
				}
			}
			lbProfile := make(map[string]interface{})
			if loadBalancer.Profile != nil {
				lbProfile[isLBName] = *loadBalancer.Profile.Name
				lbProfile[href] = *loadBalancer.Profile.Href
				lbProfile[family] = *loadBalancer.Profile.Family
			}
			if err = d.Set("profile", lbProfile); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_lb", "read", "set-profile").GetDiag()
			}
			if err = d.Set("status", loadBalancer.ProvisioningStatus); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_lb", "read", "set-status").GetDiag()
			}
			if loadBalancer.RouteMode != nil {
				if err = d.Set(isLBRouteMode, *loadBalancer.RouteMode); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode: %s", err), "(Data) ibm_is_lb", "read", "set-route_mode").GetDiag()
				}
			}
			if err = d.Set("udp_supported", loadBalancer.UDPSupported); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported: %s", err), "(Data) ibm_is_lb", "read", "set-udp_supported").GetDiag()
			}
			if err = d.Set("crn", loadBalancer.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_lb", "read", "set-crn").GetDiag()
			}
			if err = d.Set("operating_status", loadBalancer.OperatingStatus); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_status: %s", err), "(Data) ibm_is_lb", "read", "set-operating_status").GetDiag()
			}
			publicIpList := make([]string, 0)
			if loadBalancer.PublicIps != nil {
				for _, ip := range loadBalancer.PublicIps {
					if ip.Address != nil {
						pubip := *ip.Address
						publicIpList = append(publicIpList, pubip)
					}
				}
			}
			if err = d.Set(isLBPublicIPs, publicIpList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting public_ips: %s", err), "(Data) ibm_is_lb", "read", "set-public_ips").GetDiag()
			}
			privateIpList := make([]string, 0)
			privateIpDetailList := make([]map[string]interface{}, 0)
			if loadBalancer.PrivateIps != nil {
				for _, ip := range loadBalancer.PrivateIps {
					if ip.Address != nil {
						prip := *ip.Address
						privateIpList = append(privateIpList, prip)
					}
					currentPriIp := map[string]interface{}{}

					if ip.Address != nil {
						currentPriIp[isLBPrivateIpAddress] = ip.Address
					}
					if ip.Href != nil {
						currentPriIp[isLBPrivateIpHref] = ip.Href
					}
					if ip.Name != nil {
						currentPriIp[isLBPrivateIpName] = ip.Name
					}
					if ip.ID != nil {
						currentPriIp[isLBPrivateIpId] = ip.ID
					}
					if ip.ResourceType != nil {
						currentPriIp[isLBPrivateIpResourceType] = ip.ResourceType
					}
					privateIpDetailList = append(privateIpDetailList, currentPriIp)

				}
			}
			if err = d.Set(isLBPrivateIPDetail, privateIpDetailList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting private_ip: %s", err), "(Data) ibm_is_lb", "read", "set-private_ip").GetDiag()
			}
			if err = d.Set(isLBPrivateIPs, privateIpList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting private_ips: %s", err), "(Data) ibm_is_lb", "read", "set-private_ips").GetDiag()
			}
			if loadBalancer.Subnets != nil {
				subnetList := make([]string, 0)
				for _, subnet := range loadBalancer.Subnets {
					if subnet.ID != nil {
						sub := *subnet.ID
						subnetList = append(subnetList, sub)
					}
				}
				if err = d.Set(isLBSubnets, subnetList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnets: %s", err), "(Data) ibm_is_lb", "read", "set-subnets").GetDiag()
				}
			}

			if err = d.Set(isLBSecurityGroupsSupported, false); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_group_supported: %s", err), "(Data) ibm_is_lb", "read", "set-security_group_supported").GetDiag()
			}
			if loadBalancer.SecurityGroups != nil {
				securitygroupList := make([]string, 0)
				for _, securityGroup := range loadBalancer.SecurityGroups {
					if securityGroup.ID != nil {
						securityGroupID := *securityGroup.ID
						securitygroupList = append(securitygroupList, securityGroupID)
					}
				}
				if err = d.Set(isLBSecurityGroups, securitygroupList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_groups: %s", err), "(Data) ibm_is_lb", "read", "set-security_groups").GetDiag()
				}
				if err = d.Set(isLBSecurityGroupsSupported, true); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_group_supported: %s", err), "(Data) ibm_is_lb", "read", "set-security_group_supported").GetDiag()
				}
			}

			if loadBalancer.Listeners != nil {
				listenerList := make([]string, 0)
				for _, listener := range loadBalancer.Listeners {
					if listener.ID != nil {
						lis := *listener.ID
						listenerList = append(listenerList, lis)
					}
				}

				if err = d.Set(isLBListeners, listenerList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting listeners: %s", err), "(Data) ibm_is_lb", "read", "set-listeners").GetDiag()
				}
			}
			if err = d.Set("failsafe_policy_actions", loadBalancer.FailsafePolicyActions); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting failsafe_policy_actions: %s", err), "(Data) ibm_is_lb", "read", "set-failsafe_policy_actions").GetDiag()
			}
			listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{}
			listLoadBalancerPoolsOptions.SetLoadBalancerID(*loadBalancer.ID)
			poolsResult, _, _ := sess.ListLoadBalancerPools(listLoadBalancerPoolsOptions)
			if poolsResult != nil {
				poolsInfo := make([]map[string]interface{}, 0)

				for _, p := range poolsResult.Pools {
					//	log.Printf("******* p ******** : (%+v)", p)
					pool := make(map[string]interface{})
					pool[poolAlgorithm] = *p.Algorithm
					pool[ID] = *p.ID
					pool[href] = *p.Href
					pool[poolProtocol] = *p.Protocol
					pool[poolCreatedAt] = p.CreatedAt.String()
					pool[poolProvisioningStatus] = *p.ProvisioningStatus
					pool["name"] = *p.Name
					if p.HealthMonitor != nil {
						poolHealthMonitor := p.HealthMonitor.(*vpcv1.LoadBalancerPoolHealthMonitor)
						healthMonitorInfo := make(map[string]interface{})
						delayfinal := strconv.FormatInt(*(poolHealthMonitor.Delay), 10)
						healthMonitorInfo[healthMonitorDelay] = delayfinal
						maxRetriesfinal := strconv.FormatInt(*(poolHealthMonitor.MaxRetries), 10)
						timeoutfinal := strconv.FormatInt(*(poolHealthMonitor.Timeout), 10)
						healthMonitorInfo[healthMonitorMaxRetries] = maxRetriesfinal
						healthMonitorInfo[healthMonitorTimeout] = timeoutfinal
						if poolHealthMonitor.URLPath != nil {
							healthMonitorInfo[healthMonitorURLPath] = *(poolHealthMonitor.URLPath)
						}
						healthMonitorInfo[healthMonitorType] = *(poolHealthMonitor.Type)
						pool[healthMonitor] = healthMonitorInfo
					}

					if p.SessionPersistence != nil {
						sessionPersistenceInfo := make(map[string]interface{})
						sessionPersistenceInfo[sessionType] = *p.SessionPersistence.Type
						pool[sessionPersistence] = sessionPersistenceInfo
					}
					if p.Members != nil {
						memberList := make([]map[string]interface{}, len(p.Members))
						for j, m := range p.Members {
							member := make(map[string]interface{})
							member[ID] = *m.ID
							member[href] = *m.Href
							memberList[j] = member
						}
						pool[members] = memberList
					}

					if p.InstanceGroup != nil {
						instanceGroupInfo := make(map[string]interface{})
						instanceGroupInfo[ID] = *(p.InstanceGroup.ID)
						instanceGroupInfo[crnInstance] = *(p.InstanceGroup.CRN)
						instanceGroupInfo[href] = *(p.InstanceGroup.Href)
						instanceGroupInfo[name] = *(p.InstanceGroup.Name)
						pool[instanceGroup] = instanceGroupInfo
					}
					poolsInfo = append(poolsInfo, pool)
				} //for

				if err = d.Set(isLBPools, poolsInfo); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting pools: %s", err), "(Data) ibm_is_lb", "read", "set-pools").GetDiag()
				}
			}

			if err = d.Set(isLBResourceGroup, *loadBalancer.ResourceGroup.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_lb", "read", "set-resource_group").GetDiag()
			}
			if err = d.Set("hostname", loadBalancer.Hostname); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting hostname: %s", err), "(Data) ibm_is_lb", "read", "set-hostname").GetDiag()
			}
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *loadBalancer.CRN, "", isUserTagType)
			if err != nil {
				log.Printf(
					"Error on get of resource vpc Load Balancer (%s) tags: %s", d.Id(), err)
			}
			if err = d.Set(isLBTags, tags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_lb", "read", "set-tags").GetDiag()
			}
			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *loadBalancer.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error on get of resource Load Balancer (%s) access tags: %s", d.Id(), err)
			}
			if err = d.Set(isLBAccessTags, accesstags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_lb", "read", "set-access_tags").GetDiag()
			}
			controller, err := flex.GetBaseController(meta)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_lb", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/loadBalancers"); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_lb", "read", "set-resource_controller_url").GetDiag()
			}
			if err = d.Set(flex.ResourceName, *loadBalancer.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_lb", "read", "set-resource_name").GetDiag()
			}
			if loadBalancer.ResourceGroup != nil {
				if err = d.Set(flex.ResourceGroupName, *loadBalancer.ResourceGroup.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_lb", "read", "set-resource_group_name").GetDiag()
				}
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No Load balancer found with name: %s", name), "(Data) ibm_is_lb", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
