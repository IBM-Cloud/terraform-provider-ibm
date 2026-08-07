// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMISLBPools() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsLbPoolsRead,

		Schema: map[string]*schema.Schema{
			"lb": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The load balancer identifier.",
			},
			"pools": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of pools.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The load balancing algorithm.",
						},
						"client_authentication": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The client authentication used for this pool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_instance": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The certificate instance used for this pool.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN for this certificate instance.",
												},
											},
										},
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this pool was created.",
						},
						"failsafe_policy": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A load balancer failsafe policy action:- `forward`: Forwards requests to the `target` pool.- `fail`: Rejects requests with an HTTP `503` status code.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"healthy_member_threshold_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The healthy member count at which the failsafe policy action will be triggered. At present, this is always `0`, but may be modifiable in the future.",
									},
									"target": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If `action` is `forward`, the target pool to forward to.If `action` is `fail`, this property will be absent.The targets supported by this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deleted": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"more_info": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Link to documentation about deleted resources.",
															},
														},
													},
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this load balancer pool.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this load balancer pool.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this load balancer pool. The name is unique across all pools for the load balancer.",
												},
											},
										},
									},
								},
							},
						},
						"health_monitor": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The health monitor of this pool.If this pool has a member targeting a load balancer then:- If the targeted load balancer has multiple subnets, this health monitor is used to  direct traffic to the available subnets.- The health checks spawned by this health monitor is handled as any other traffic  (that is, subject to the configuration of listeners and pools on the target load  balancer).- This health monitor does not affect how pool member health is determined within the  target load balancer.For more information, see [Private Path network load balancer frequently askedquestions](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-faqs#ppnlb-faqs).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delay": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The seconds to wait between health checks.",
									},
									"max_retries": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The health check max retries.",
									},
									"port": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The health check port.If present, this overrides the pool member port values.",
									},
									"timeout": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The seconds to wait for a response to a health check.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol type used for health checks.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"request": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"body": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The HTTP request body used for health checks.If absent, the health checks will ignore the request body.",
												},
												"headers": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The HTTP request headers used for health checks.If absent, the health checks will ignore the request headers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"field": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The field of an HTTP request header used for health checks.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value of an HTTP request header used for health checks.",
															},
														},
													},
												},
												"method": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The HTTP request method used for health checks.",
												},
											},
										},
									},
									"response": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"body_regex": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The PCRE-flavor regular expression that HTTP response bodies must match for successful health checks.If absent, health checks will ignore any response body.",
												},
												"codes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The HTTP response codes expected for successful health checks.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"url_path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health check URL path, in the format of an [origin-form request target](https://tools.ietf.org/html/rfc7230#section-5.3.1).",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The pool's canonical URL.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this load balancer pool.",
						},
						"instance_group": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The instance group that is managing this pool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this instance group.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this instance group.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this instance group.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this instance group.",
									},
								},
							},
						},
						"members": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The backend server members of the pool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The member's canonical URL.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this load balancer pool member.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this load balancer pool.",
						},
						"protocol": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol used for this load balancer pool.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the pool on which the unexpected property value was encountered.",
						},
						"provisioning_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The provisioning status of this pool.",
						},
						"proxy_protocol": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The PROXY protocol setting for this pool:- `v1`: Enabled with version 1 (human-readable header format)- `v2`: Enabled with version 2 (binary header format)- `disabled`: DisabledSupported by load balancers in the `application` family (otherwise always `disabled`).",
						},
						"server_authentication": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The server authentication used for this pool. This property will be absent if the pool.protocol is not https.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_authority": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The certificate authority used for this pool.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN for this certificate instance.",
												},
											},
										},
									},
									"verify_certificate": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If set to true, the backend server certificate is verified.",
									},
								},
							},
						},
						"session_persistence": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The session persistence of this pool.The enumerated values for this property are expected to expand in the future. Whenprocessing this property, check for and log unknown values. Optionally haltprocessing and surface the error, or bypass the pool on which the unexpectedproperty value was encountered.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cookie_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The session persistence cookie name. Applicable only for type `app_cookie`. Names starting with `IBM` are not allowed.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The session persistence type. The `http_cookie` and `app_cookie` types are applicable only to the `http` and `https` protocols.",
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

func dataSourceIBMIsLbPoolsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_pools", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{}

	listLoadBalancerPoolsOptions.SetLoadBalancerID(d.Get("lb").(string))

	loadBalancerPoolCollection, _, err := sess.ListLoadBalancerPoolsWithContext(context, listLoadBalancerPoolsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListLoadBalancerPoolsWithContext failed: %s", err.Error()), "(Data) ibm_is_lb_pools", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("lb", d.Get("lb").(string)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lb: %s", err), "(Data) ibm_is_lb_pools", "read", "set-lb").GetDiag()
	}

	d.SetId(dataSourceIBMIsLbPoolsID(d))

	if loadBalancerPoolCollection.Pools != nil {
		err = d.Set("pools", dataSourceLoadBalancerPoolCollectionFlattenPools(loadBalancerPoolCollection.Pools))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting pools: %s", err), "(Data) ibm_is_lb_pools", "read", "set-pools").GetDiag()
		}
	}
	return nil
}

// dataSourceIBMIsLbPoolsID returns a reasonable ID for the list.
func dataSourceIBMIsLbPoolsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceLoadBalancerPoolCollectionFlattenPools(result []vpcv1.LoadBalancerPool) (pools []map[string]interface{}) {
	for _, poolsItem := range result {
		pools = append(pools, dataSourceLoadBalancerPoolCollectionPoolsToMap(poolsItem))
	}

	return pools
}

func dataSourceLoadBalancerPoolCollectionPoolsToMap(poolsItem vpcv1.LoadBalancerPool) (poolsMap map[string]interface{}) {
	poolsMap = map[string]interface{}{}

	if poolsItem.Algorithm != nil {
		poolsMap["algorithm"] = poolsItem.Algorithm
	}
	if poolsItem.ClientAuthentication != nil {
		clientAuthenticationList := []map[string]interface{}{}
		clientAuthenticationMap := dataSourceLoadBalancerPoolCollectionPoolsClientAuthenticationToMap(*poolsItem.ClientAuthentication)
		clientAuthenticationList = append(clientAuthenticationList, clientAuthenticationMap)
		poolsMap["client_authentication"] = clientAuthenticationList
	}
	if poolsItem.CreatedAt != nil {
		poolsMap["created_at"] = poolsItem.CreatedAt.String()
	}
	var healthMonitorMap map[string]interface{}
	if !core.IsNil(poolsItem.HealthMonitor) {
		healthMonitorMap, _ = DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorToMap(poolsItem.HealthMonitor)
	}
	if healthMonitorMap == nil {
		healthMonitorMap = map[string]interface{}{}
	}
	poolsMap["health_monitor"] = []map[string]interface{}{healthMonitorMap}
	if poolsItem.Href != nil {
		poolsMap["href"] = poolsItem.Href
	}
	if poolsItem.ID != nil {
		poolsMap["id"] = poolsItem.ID
	}
	if poolsItem.InstanceGroup != nil {
		instanceGroupList := []map[string]interface{}{}
		instanceGroupMap := dataSourceLoadBalancerPoolCollectionPoolsInstanceGroupToMap(*poolsItem.InstanceGroup)
		instanceGroupList = append(instanceGroupList, instanceGroupMap)
		poolsMap["instance_group"] = instanceGroupList
	}
	if poolsItem.Members != nil {
		membersList := []map[string]interface{}{}
		for _, membersItem := range poolsItem.Members {
			membersList = append(membersList, dataSourceLoadBalancerPoolCollectionPoolsMembersToMap(membersItem))
		}
		poolsMap["members"] = membersList
	}
	if poolsItem.Name != nil {
		poolsMap["name"] = poolsItem.Name
	}
	if poolsItem.Protocol != nil {
		poolsMap["protocol"] = poolsItem.Protocol
	}
	if poolsItem.ProvisioningStatus != nil {
		poolsMap["provisioning_status"] = poolsItem.ProvisioningStatus
	}
	if poolsItem.ProxyProtocol != nil {
		poolsMap["proxy_protocol"] = poolsItem.ProxyProtocol
	}
	if poolsItem.FailsafePolicy != nil {
		failsafePolicyMap, err := dataSourceIBMIsLbPoolsLoadBalancerPoolFailsafePolicyToMap(poolsItem.FailsafePolicy)
		if err != nil {
			return poolsMap
		}
		poolsMap["failsafe_policy"] = []map[string]interface{}{failsafePolicyMap}
	}
	if poolsItem.ServerAuthentication != nil {
		serverAuthenticationList := []map[string]interface{}{}
		serverAuthenticationMap := dataSourceLoadBalancerPoolCollectionPoolsServerAuthenticationToMap(*poolsItem.ServerAuthentication)
		serverAuthenticationList = append(serverAuthenticationList, serverAuthenticationMap)
		poolsMap["server_authentication"] = serverAuthenticationList
	}
	if poolsItem.SessionPersistence != nil {
		sessionPersistenceList := []map[string]interface{}{}
		sessionPersistenceMap := dataSourceLoadBalancerPoolCollectionPoolsSessionPersistenceToMap(*poolsItem.SessionPersistence)
		sessionPersistenceList = append(sessionPersistenceList, sessionPersistenceMap)
		poolsMap["session_persistence"] = sessionPersistenceList
	}

	return poolsMap
}

func dataSourceLoadBalancerPoolCollectionPoolsInstanceGroupToMap(instanceGroupItem vpcv1.InstanceGroupReference) (instanceGroupMap map[string]interface{}) {
	instanceGroupMap = map[string]interface{}{}

	if instanceGroupItem.CRN != nil {
		instanceGroupMap["crn"] = instanceGroupItem.CRN
	}
	if instanceGroupItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceLoadBalancerPoolCollectionInstanceGroupDeletedToMap(*instanceGroupItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		instanceGroupMap["deleted"] = deletedList
	}
	if instanceGroupItem.Href != nil {
		instanceGroupMap["href"] = instanceGroupItem.Href
	}
	if instanceGroupItem.ID != nil {
		instanceGroupMap["id"] = instanceGroupItem.ID
	}
	if instanceGroupItem.Name != nil {
		instanceGroupMap["name"] = instanceGroupItem.Name
	}

	return instanceGroupMap
}

func dataSourceLoadBalancerPoolCollectionInstanceGroupDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceLoadBalancerPoolCollectionPoolsMembersToMap(membersItem vpcv1.LoadBalancerPoolMemberReference) (membersMap map[string]interface{}) {
	membersMap = map[string]interface{}{}

	if membersItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceLoadBalancerPoolCollectionMembersDeletedToMap(*membersItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		membersMap["deleted"] = deletedList
	}
	if membersItem.Href != nil {
		membersMap["href"] = membersItem.Href
	}
	if membersItem.ID != nil {
		membersMap["id"] = membersItem.ID
	}

	return membersMap
}

func dataSourceLoadBalancerPoolCollectionMembersDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceLoadBalancerPoolCollectionPoolsSessionPersistenceToMap(sessionPersistenceItem vpcv1.LoadBalancerPoolSessionPersistence) (sessionPersistenceMap map[string]interface{}) {
	sessionPersistenceMap = map[string]interface{}{}

	if sessionPersistenceItem.CookieName != nil {
		sessionPersistenceMap["cookie_name"] = sessionPersistenceItem.CookieName
	}
	if sessionPersistenceItem.Type != nil {
		sessionPersistenceMap["type"] = sessionPersistenceItem.Type
	}

	return sessionPersistenceMap
}

func dataSourceIBMIsLbPoolsLoadBalancerPoolFailsafePolicyToMap(model *vpcv1.LoadBalancerPoolFailsafePolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["action"] = *model.Action
	modelMap["healthy_member_threshold_count"] = flex.IntValue(model.HealthyMemberThresholdCount)
	if model.Target != nil {
		targetMap, err := dataSourceIBMIsLbPoolsLoadBalancerPoolReferenceToMap(model.Target)
		if err != nil {
			return modelMap, err
		}
		modelMap["target"] = []map[string]interface{}{targetMap}
	}
	return modelMap, nil
}

func dataSourceIBMIsLbPoolsLoadBalancerPoolReferenceToMap(model *vpcv1.LoadBalancerPoolReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsLbPoolsDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func dataSourceIBMIsLbPoolsDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorToMap(model vpcv1.LoadBalancerPoolHealthMonitorIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps); ok {
		return DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps))
	} else if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP); ok {
		return DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeTCPToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP))
	} else if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitor); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.LoadBalancerPoolHealthMonitor)
		modelMap["delay"] = flex.IntValue(model.Delay)
		modelMap["max_retries"] = flex.IntValue(model.MaxRetries)
		if model.Port != nil {
			modelMap["port"] = flex.IntValue(model.Port)
		}
		modelMap["timeout"] = flex.IntValue(model.Timeout)
		modelMap["type"] = *model.Type
		if model.Request != nil {
			requestMap, err := DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
			if err != nil {
				return modelMap, err
			}
			modelMap["request"] = []map[string]interface{}{requestMap}
		}
		if model.Response != nil {
			responseMap, err := DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
			if err != nil {
				return modelMap, err
			}
			modelMap["response"] = []map[string]interface{}{responseMap}
		}
		if model.URLPath != nil {
			modelMap["url_path"] = *model.URLPath
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.LoadBalancerPoolHealthMonitorIntf subtype encountered")
	}
}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequest) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Body != nil {
		modelMap["body"] = *model.Body
	}
	if model.HeadersVar != nil {
		headers := []map[string]interface{}{}
		for _, headersItem := range model.HeadersVar {
			headersItemMap, err := DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(&headersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			headers = append(headers, headersItemMap)
		}
		modelMap["headers"] = headers
	}
	modelMap["method"] = *model.Method
	return modelMap, nil
}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeader) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Field != nil {
		modelMap["field"] = *model.Field
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponse) (map[string]interface{}, error) {

	modelMap := make(map[string]interface{})
	if model.BodyRegex != nil {
		modelMap["body_regex"] = *model.BodyRegex
	}
	if model.Codes != nil {
		modelMap["codes"] = model.Codes
	}
	return modelMap, nil

}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["delay"] = flex.IntValue(model.Delay)
	modelMap["max_retries"] = flex.IntValue(model.MaxRetries)
	if model.Port != nil {
		modelMap["port"] = flex.IntValue(model.Port)
	}
	modelMap["timeout"] = flex.IntValue(model.Timeout)
	if model.Request != nil {
		requestMap, err := DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
		if err != nil {
			return modelMap, err
		}
		modelMap["request"] = []map[string]interface{}{requestMap}
	}
	if model.Response != nil {
		responseMap, err := DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
		if err != nil {
			return modelMap, err
		}
		modelMap["response"] = []map[string]interface{}{responseMap}
	}
	modelMap["type"] = *model.Type
	if model.URLPath != nil {
		modelMap["url_path"] = *model.URLPath
	}
	return modelMap, nil
}

func DataSourceIBMIsLbPoolsLoadBalancerPoolHealthMonitorTypeTCPToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeTCP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["delay"] = flex.IntValue(model.Delay)
	modelMap["max_retries"] = flex.IntValue(model.MaxRetries)
	if model.Port != nil {
		modelMap["port"] = flex.IntValue(model.Port)
	}
	modelMap["timeout"] = flex.IntValue(model.Timeout)
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func dataSourceLoadBalancerPoolCollectionPoolsClientAuthenticationToMap(clientAuthItem vpcv1.LoadBalancerPoolClientAuthentication) (clientAuthMap map[string]interface{}) {
	clientAuthMap = map[string]interface{}{}

	if clientAuthItem.CertificateInstance != nil {
		certificateInstanceList := []map[string]interface{}{}
		certificateInstanceMap := dataSourceLoadBalancerPoolCollectionPoolsCertificateInstanceToMap(*clientAuthItem.CertificateInstance)
		certificateInstanceList = append(certificateInstanceList, certificateInstanceMap)
		clientAuthMap["certificate_instance"] = certificateInstanceList
	}

	return clientAuthMap
}

func dataSourceLoadBalancerPoolCollectionPoolsServerAuthenticationToMap(serverAuthItem vpcv1.LoadBalancerPoolServerAuthentication) (serverAuthMap map[string]interface{}) {
	serverAuthMap = map[string]interface{}{}

	if serverAuthItem.CertificateAuthority != nil {
		certificateAuthorityList := []map[string]interface{}{}
		certificateAuthorityMap := dataSourceLoadBalancerPoolCollectionPoolsCertificateInstanceToMap(*serverAuthItem.CertificateAuthority)
		certificateAuthorityList = append(certificateAuthorityList, certificateAuthorityMap)
		serverAuthMap["certificate_authority"] = certificateAuthorityList
	}
	if serverAuthItem.VerifyCertificate != nil {
		serverAuthMap["verify_certificate"] = serverAuthItem.VerifyCertificate
	}

	return serverAuthMap
}

func dataSourceLoadBalancerPoolCollectionPoolsCertificateInstanceToMap(certificateInstanceItem vpcv1.CertificateInstanceReference) (certificateInstanceMap map[string]interface{}) {
	certificateInstanceMap = map[string]interface{}{}

	if certificateInstanceItem.CRN != nil {
		certificateInstanceMap["crn"] = certificateInstanceItem.CRN
	}

	return certificateInstanceMap
}
