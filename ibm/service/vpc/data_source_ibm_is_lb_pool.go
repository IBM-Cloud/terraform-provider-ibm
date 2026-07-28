// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMISLBPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsLbPoolRead,

		Schema: map[string]*schema.Schema{
			"lb": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The load balancer identifier.",
			},
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The pool identifier.",
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
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The user-defined name for this load balancer pool.",
			},
			"algorithm": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The load balancing algorithm.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this pool was created.",
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
			// http bundle
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
	}
}

func dataSourceIBMIsLbPoolRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_pool", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var loadBalancerPool *vpcv1.LoadBalancerPool

	if v, ok := d.GetOk("identifier"); ok {
		getLoadBalancerPoolOptions := &vpcv1.GetLoadBalancerPoolOptions{}

		getLoadBalancerPoolOptions.SetLoadBalancerID(d.Get("lb").(string))
		getLoadBalancerPoolOptions.SetID(v.(string))

		loadBalancerPoolInfo, _, err := sess.GetLoadBalancerPoolWithContext(context, getLoadBalancerPoolOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolWithContext failed: %s", err.Error()), "(Data) ibm_is_lb_pool", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		loadBalancerPool = loadBalancerPoolInfo

	} else if v, ok := d.GetOk("name"); ok {
		listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{}

		listLoadBalancerPoolsOptions.SetLoadBalancerID(d.Get("lb").(string))

		loadBalancerPoolCollection, _, err := sess.ListLoadBalancerPoolsWithContext(context, listLoadBalancerPoolsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListLoadBalancerPoolsWithContext failed: %s", err.Error()), "(Data) ibm_is_lb_pool", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		name := v.(string)
		for _, data := range loadBalancerPoolCollection.Pools {
			if *data.Name == name {
				loadBalancerPool = &data
				break
			}
		}
		if loadBalancerPool == nil {
			log.Printf("[DEBUG] No LoadBalancerPool found with name (%s)", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No LoadBalancerPool found with name: %s", name), "(Data) ibm_is_lb_pool", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}

	d.SetId(*loadBalancerPool.ID)
	if err = d.Set("algorithm", loadBalancerPool.Algorithm); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting algorithm: %s", err), "(Data) ibm_is_lb_pool", "read", "set-algorithm").GetDiag()
	}

	if loadBalancerPool.ClientAuthentication != nil {
		err = d.Set("client_authentication", dataSourceLoadBalancerPoolFlattenClientAuthentication(*loadBalancerPool.ClientAuthentication))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting client_authentication: %s", err), "(Data) ibm_is_lb_pool", "read", "set-client_authentication").GetDiag()
		}
	}

	failsafePolicy := []map[string]interface{}{}
	if loadBalancerPool.FailsafePolicy != nil {
		modelMap, err := dataSourceIBMIsLbPoolLoadBalancerPoolFailsafePolicyToMap(loadBalancerPool.FailsafePolicy)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_pool", "read", "failsafe_policy-to-map").GetDiag()
		}
		failsafePolicy = append(failsafePolicy, modelMap)
	}
	if err = d.Set("failsafe_policy", failsafePolicy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting failsafe_policy: %s", err), "(Data) ibm_is_lb_pool", "read", "set-failsafe_policy").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(loadBalancerPool.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_lb_pool", "read", "set-created_at").GetDiag()
	}

	// http bundle
	healthMonitor := []map[string]interface{}{}
	healthMonitorMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorToMap(loadBalancerPool.HealthMonitor)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_pool", "read", "health_monitor-to-map").GetDiag()
	}
	healthMonitor = append(healthMonitor, healthMonitorMap)
	if err = d.Set("health_monitor", healthMonitor); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_monitor: %s", err), "(Data) ibm_is_lb_pool", "read", "set-health_monitor").GetDiag()
	}
	if err = d.Set("href", loadBalancerPool.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_lb_pool", "read", "set-href").GetDiag()
	}

	if loadBalancerPool.InstanceGroup != nil {
		err = d.Set("instance_group", dataSourceLoadBalancerPoolFlattenInstanceGroup(*loadBalancerPool.InstanceGroup))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_group: %s", err), "(Data) ibm_is_lb_pool", "read", "set-instance_group").GetDiag()
		}
	}

	if loadBalancerPool.Members != nil {
		err = d.Set("members", dataSourceLoadBalancerPoolFlattenMembers(loadBalancerPool.Members))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting members: %s", err), "(Data) ibm_is_lb_pool", "read", "set-members").GetDiag()
		}
	}

	if err = d.Set("identifier", loadBalancerPool.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting identifier: %s", err), "(Data) ibm_is_lb_pool", "read", "set-identifier").GetDiag()
	}

	if err = d.Set("name", loadBalancerPool.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_lb_pool", "read", "set-name").GetDiag()
	}
	if err = d.Set("protocol", loadBalancerPool.Protocol); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_lb_pool", "read", "set-protocol").GetDiag()
	}
	if err = d.Set("provisioning_status", loadBalancerPool.ProvisioningStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provisioning_status: %s", err), "(Data) ibm_is_lb_pool", "read", "set-provisioning_status").GetDiag()
	}

	if err = d.Set("proxy_protocol", loadBalancerPool.ProxyProtocol); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting proxy_protocol: %s", err), "(Data) ibm_is_lb_pool", "read", "set-proxy_protocol").GetDiag()
	}

	if loadBalancerPool.ServerAuthentication != nil {
		err = d.Set("server_authentication", dataSourceLoadBalancerPoolFlattenServerAuthentication(*loadBalancerPool.ServerAuthentication))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting session_persistence: %s", err), "(Data) ibm_is_lb_pool", "read", "set-server_authentication").GetDiag()
		}
	}

	if loadBalancerPool.SessionPersistence != nil {
		err = d.Set("session_persistence", dataSourceLoadBalancerPoolFlattenSessionPersistence(*loadBalancerPool.SessionPersistence))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting session_persistence: %s", err), "(Data) ibm_is_lb_pool", "read", "set-session_persistence").GetDiag()
		}
	}

	return nil
}

func dataSourceLoadBalancerPoolFlattenInstanceGroup(result vpcv1.InstanceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceLoadBalancerPoolInstanceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceLoadBalancerPoolInstanceGroupToMap(instanceGroupItem vpcv1.InstanceGroupReference) (instanceGroupMap map[string]interface{}) {
	instanceGroupMap = map[string]interface{}{}

	if instanceGroupItem.CRN != nil {
		instanceGroupMap["crn"] = instanceGroupItem.CRN
	}
	if instanceGroupItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceLoadBalancerPoolInstanceGroupDeletedToMap(*instanceGroupItem.Deleted)
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

func dataSourceLoadBalancerPoolInstanceGroupDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceLoadBalancerPoolFlattenMembers(result []vpcv1.LoadBalancerPoolMemberReference) (members []map[string]interface{}) {
	for _, membersItem := range result {
		members = append(members, dataSourceLoadBalancerPoolMembersToMap(membersItem))
	}

	return members
}

func dataSourceLoadBalancerPoolMembersToMap(membersItem vpcv1.LoadBalancerPoolMemberReference) (membersMap map[string]interface{}) {
	membersMap = map[string]interface{}{}

	if membersItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceLoadBalancerPoolMembersDeletedToMap(*membersItem.Deleted)
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

func dataSourceLoadBalancerPoolMembersDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceLoadBalancerPoolFlattenSessionPersistence(result vpcv1.LoadBalancerPoolSessionPersistence) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceLoadBalancerPoolSessionPersistenceToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceLoadBalancerPoolSessionPersistenceToMap(sessionPersistenceItem vpcv1.LoadBalancerPoolSessionPersistence) (sessionPersistenceMap map[string]interface{}) {
	sessionPersistenceMap = map[string]interface{}{}

	if sessionPersistenceItem.CookieName != nil {
		sessionPersistenceMap["cookie_name"] = sessionPersistenceItem.CookieName
	}
	if sessionPersistenceItem.Type != nil {
		sessionPersistenceMap["type"] = sessionPersistenceItem.Type
	}

	return sessionPersistenceMap
}

func dataSourceIBMIsLbPoolLoadBalancerPoolFailsafePolicyToMap(model *vpcv1.LoadBalancerPoolFailsafePolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["action"] = *model.Action
	modelMap["healthy_member_threshold_count"] = flex.IntValue(model.HealthyMemberThresholdCount)
	if model.Target != nil {
		targetMap, err := dataSourceIBMIsLbPoolLoadBalancerPoolReferenceToMap(model.Target)
		if err != nil {
			return modelMap, err
		}
		modelMap["target"] = []map[string]interface{}{targetMap}
	}
	return modelMap, nil
}

func dataSourceIBMIsLbPoolLoadBalancerPoolReferenceToMap(model *vpcv1.LoadBalancerPoolReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsLbPoolDeletedToMap(model.Deleted)
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

func dataSourceIBMIsLbPoolDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorToMap(model vpcv1.LoadBalancerPoolHealthMonitorIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps); ok {
		return DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps))
	} else if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP); ok {
		return DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeTCPToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP))
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
			requestMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
			if err != nil {
				return modelMap, err
			}
			modelMap["request"] = []map[string]interface{}{requestMap}
		}
		if model.Response != nil {
			responseMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
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

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequest) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Body != nil {
		modelMap["body"] = *model.Body
	}
	if model.HeadersVar != nil {
		headers := []map[string]interface{}{}
		for _, headersItem := range model.HeadersVar {
			headersItemMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(&headersItem) // #nosec G601
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

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeader) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Field != nil {
		modelMap["field"] = *model.Field
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BodyRegex != nil {
		modelMap["body_regex"] = *model.BodyRegex
	}
	if model.Codes != nil {
		modelMap["codes"] = model.Codes
	}
	return modelMap, nil
}

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["delay"] = flex.IntValue(model.Delay)
	modelMap["max_retries"] = flex.IntValue(model.MaxRetries)
	if model.Port != nil {
		modelMap["port"] = flex.IntValue(model.Port)
	}
	modelMap["timeout"] = flex.IntValue(model.Timeout)
	if model.Request != nil {
		requestMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
		if err != nil {
			return modelMap, err
		}
		modelMap["request"] = []map[string]interface{}{requestMap}
	}
	if model.Response != nil {
		responseMap, err := DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
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

func DataSourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeTCPToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeTCP) (map[string]interface{}, error) {
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

func dataSourceLoadBalancerPoolFlattenClientAuthentication(result vpcv1.LoadBalancerPoolClientAuthentication) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceLoadBalancerPoolClientAuthenticationToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceLoadBalancerPoolClientAuthenticationToMap(clientAuthItem vpcv1.LoadBalancerPoolClientAuthentication) (clientAuthMap map[string]interface{}) {
	clientAuthMap = map[string]interface{}{}

	if clientAuthItem.CertificateInstance != nil {
		certificateInstanceList := []map[string]interface{}{}
		certificateInstanceMap := dataSourceLoadBalancerPoolCertificateInstanceToMap(*clientAuthItem.CertificateInstance)
		certificateInstanceList = append(certificateInstanceList, certificateInstanceMap)
		clientAuthMap["certificate_instance"] = certificateInstanceList
	}

	return clientAuthMap
}

func dataSourceLoadBalancerPoolCertificateInstanceToMap(certificateInstanceItem vpcv1.CertificateInstanceReference) (certificateInstanceMap map[string]interface{}) {
	certificateInstanceMap = map[string]interface{}{}

	if certificateInstanceItem.CRN != nil {
		certificateInstanceMap["crn"] = certificateInstanceItem.CRN
	}

	return certificateInstanceMap
}

func dataSourceLoadBalancerPoolFlattenServerAuthentication(result vpcv1.LoadBalancerPoolServerAuthentication) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceLoadBalancerPoolServerAuthenticationToMap(result)
	finalList = append(finalList, finalMap)
	return finalList
}

func dataSourceLoadBalancerPoolServerAuthenticationToMap(serverAuthItem vpcv1.LoadBalancerPoolServerAuthentication) (serverAuthMap map[string]interface{}) {
	serverAuthMap = map[string]interface{}{}

	if serverAuthItem.CertificateAuthority != nil {
		certificateAuthorityList := []map[string]interface{}{}
		certificateAuthorityMap := dataSourceLoadBalancerPoolCertificateInstanceToMap(*serverAuthItem.CertificateAuthority)
		certificateAuthorityList = append(certificateAuthorityList, certificateAuthorityMap)
		serverAuthMap["certificate_authority"] = certificateAuthorityList
	}
	if serverAuthItem.VerifyCertificate != nil {
		serverAuthMap["verify_certificate"] = serverAuthItem.VerifyCertificate
	}
	return serverAuthMap
}
