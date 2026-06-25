// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isLBPoolName                          = "name"
	isLBID                                = "lb"
	isLBPoolAlgorithm                     = "algorithm"
	isLBPoolProtocol                      = "protocol"
	isLBPoolHealthDelay                   = "health_delay"
	isLBPoolHealthRetries                 = "health_retries"
	isLBPoolHealthTimeout                 = "health_timeout"
	isLBPoolHealthType                    = "health_type"
	isLBPoolHealthMonitorURL              = "health_monitor_url"
	isLBPoolHealthMonitorPort             = "health_monitor_port"
	isLBPoolSessPersistenceType           = "session_persistence_type"
	isLBPoolSessPersistenceAppCookieName  = "session_persistence_app_cookie_name"
	isLBPoolSessPersistenceHttpCookieName = "session_persistence_http_cookie_name"
	isLBPoolProvisioningStatus            = "provisioning_status"
	isLBPoolProxyProtocol                 = "proxy_protocol"
	isLBPoolActive                        = "active"
	isLBPoolCreatePending                 = "create_pending"
	isLBPoolUpdatePending                 = "update_pending"
	isLBPoolDeletePending                 = "delete_pending"
	isLBPoolMaintainancePending           = "maintenance_pending"
	isLBPoolFailed                        = "failed"
	isLBPoolDeleteDone                    = "deleted"
	isLBPool                              = "pool_id"
)

func ResourceIBMISLBPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISLBPoolCreate,
		ReadContext:   resourceIBMISLBPoolRead,
		UpdateContext: resourceIBMISLBPoolUpdate,
		DeleteContext: resourceIBMISLBPoolDelete,
		Exists:        resourceIBMISLBPoolExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourceIBMISLBPoolCookieValidate(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isLBPoolName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolName),
				Description:  "Load Balancer Pool name",
			},

			isLBID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Load Balancer ID",
			},

			isLBPoolAlgorithm: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolAlgorithm),
				Description:  "Load Balancer Pool algorithm",
			},

			"failsafe_policy": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The failsafe policy to use for this pool.If unspecified, the default failsafe policy action from the profile will be used.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A load balancer failsafe policy action:- `forward`: Forwards requests to the `target` pool.- `fail`: Rejects requests with an HTTP `503` status code.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"healthy_member_threshold_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The healthy member count at which the failsafe policy action will be triggered. At present, this is always `0`, but may be modifiable in the future.",
						},
						"target": &schema.Schema{
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							DiffSuppressFunc: suppressNullTarget,
							Description:      "If `action` is `forward`, the target pool to forward to.If `action` is `fail`, this property will be absent.The targets supported by this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
										Optional:    true,
										Computed:    true,
										Description: "The URL for this load balancer pool.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
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
			isLBPoolProtocol: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolProtocol),
				Description:  "Load Balancer Protocol",
			},

			isLBPoolHealthDelay: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Blancer health delay time period",
			},

			isLBPoolHealthRetries: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Balancer health retry count",
			},

			isLBPoolHealthTimeout: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Balancer health timeout interval",
			},

			isLBPoolHealthType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolHealthType),
				Description:  "Load Balancer health type",
			},

			isLBPoolHealthMonitorURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Health monitor URL of LB Pool",
			},

			isLBPoolHealthMonitorPort: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Health monitor Port the LB Pool",
			},

			isLBPoolSessPersistenceType: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolSessPersistenceType),
				Description:  "Load Balancer Pool session persisence type.",
			},

			isLBPoolSessPersistenceAppCookieName: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolSessPersistenceAppCookieName),
				Description:  "Load Balancer Pool session persisence app cookie name.",
			},

			isLBPoolSessPersistenceHttpCookieName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer Pool session persisence http cookie name.",
			},

			isLBPoolProvisioningStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the LB Pool",
			},

			isLBPoolProxyProtocol: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool", isLBPoolProxyProtocol),
				Description:  "PROXY protocol setting for this pool",
			},

			"health_monitor": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The health monitor of this pool.If this pool has a member targeting a load balancer then:- If the targeted load balancer has multiple subnets, this health monitor is used to  direct traffic to the available subnets.- The health checks spawned by this health monitor is handled as any other traffic  (that is, subject to the configuration of listeners and pools on the target load  balancer).- This health monitor does not affect how pool member health is determined within the  target load balancer.For more information, see [Private Path network load balancer frequently askedquestions](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-faqs#ppnlb-faqs).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"body": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The HTTP request body used for health checks.If absent, the health checks will ignore the request body.",
									},
									"headers": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The HTTP request headers used for health checks.If absent, the health checks will ignore the request headers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"field": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The field of an HTTP request header used for health checks.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of an HTTP request header used for health checks.",
												},
											},
										},
									},
									"method": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The HTTP request method used for health checks.",
									},
								},
							},
						},
						"response": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"body_regex": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The PCRE-flavor regular expression that HTTP response bodies must match for successful health checks.If absent, health checks will ignore any response body.",
									},
									"codes": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The HTTP response codes expected for successful health checks.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},

			isLBPool: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The LB Pool id",
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the LB resource",
			},
		},
	}
}

func ResourceIBMISLBPoolValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	algorithm := "round_robin, weighted_round_robin, least_connections"
	protocol := "http, tcp, https, udp"
	persistanceType := "source_ip, app_cookie, http_cookie"
	proxyProtocol := "disabled, v1, v2"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolAlgorithm,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              algorithm})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolProtocol,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              protocol})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolHealthType,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              protocol})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolProxyProtocol,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              proxyProtocol})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolSessPersistenceAppCookieName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     "^[-A-Za-z0-9!#$%&'*+.^_`~|]+$",
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolSessPersistenceType,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              persistanceType})

	ibmISLBPoolResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_lb_pool", Schema: validateSchema}
	return &ibmISLBPoolResourceValidator
}

func resourceIBMISLBPoolCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] LB Pool create")
	name := d.Get(isLBPoolName).(string)
	lbID := d.Get(isLBID).(string)
	algorithm := d.Get(isLBPoolAlgorithm).(string)
	protocol := d.Get(isLBPoolProtocol).(string)
	healthDelay := int64(d.Get(isLBPoolHealthDelay).(int))
	maxRetries := int64(d.Get(isLBPoolHealthRetries).(int))
	healthTimeOut := int64(d.Get(isLBPoolHealthTimeout).(int))
	healthType := d.Get(isLBPoolHealthType).(string)

	var spType, cName, healthMonitorURL, pProtocol string
	var healthMonitorPort int64
	if pt, ok := d.GetOk(isLBPoolSessPersistenceType); ok {
		spType = pt.(string)
	}

	if cn, ok := d.GetOk(isLBPoolSessPersistenceAppCookieName); ok {
		cName = cn.(string)
	}
	if pp, ok := d.GetOk(isLBPoolProxyProtocol); ok {
		pProtocol = pp.(string)
	}

	if hmu, ok := d.GetOk(isLBPoolHealthMonitorURL); ok {
		healthMonitorURL = hmu.(string)
	}

	if hmp, ok := d.GetOk(isLBPoolHealthMonitorPort); ok {
		healthMonitorPort = int64(hmp.(int))
	}

	err := lbPoolCreate(context, d, meta, name, lbID, algorithm, protocol, healthType, spType, cName, healthMonitorURL, pProtocol, healthDelay, maxRetries, healthTimeOut, healthMonitorPort)
	if err != nil {
		return err
	}

	return resourceIBMISLBPoolRead(context, d, meta)
}

func lbPoolCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, lbID, algorithm, protocol, healthType, spType, cName, healthMonitorURL, pProtocol string, healthDelay, maxRetries, healthTimeOut, healthMonitorPort int64) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	healthMonitor := &vpcv1.LoadBalancerPoolHealthMonitorPrototype{
		Delay:      &healthDelay,
		MaxRetries: &maxRetries,
		Timeout:    &healthTimeOut,
		Type:       &healthType,
	}
	hmMap, _ := d.Get("health_monitor.0").(map[string]interface{})
	healthMonitorRequestModel, healthMonitorResponseModel, err := ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorPrototype(hmMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "create", "parse-health_monitor").GetDiag()
	}
	options := &vpcv1.CreateLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		Algorithm:      &algorithm,
		Protocol:       &protocol,
		Name:           &name,
	}
	if healthMonitorURL != "" {
		healthMonitor.URLPath = &healthMonitorURL
	}
	if healthMonitorRequestModel != nil {
		healthMonitor.Request = healthMonitorRequestModel
	}
	if healthMonitorResponseModel != nil {
		healthMonitor.Response = healthMonitorResponseModel
	}
	if healthMonitorPort > int64(0) {
		healthMonitor.Port = &healthMonitorPort
	}
	if spType != "" {
		options.SessionPersistence = &vpcv1.LoadBalancerPoolSessionPersistencePrototype{
			Type: &spType,
		}
		if cName != "" {
			options.SessionPersistence.CookieName = &cName
		}
	}
	if pProtocol != "" {
		options.ProxyProtocol = &pProtocol
	}
	if _, ok := d.GetOk("failsafe_policy"); ok {
		failsafePolicyModel, err := resourceIBMIsLbPoolMapToLoadBalancerPoolFailsafePolicyPrototype(d.Get("failsafe_policy.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "create", "parse-failsafe_policy").GetDiag()
		}
		options.SetFailsafePolicy(failsafePolicyModel)
	}
	options.HealthMonitor = healthMonitor
	lbPool, _, err := sess.CreateLoadBalancerPoolWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", lbID, *lbPool.ID))
	log.Printf("[INFO] lbpool : %s", *lbPool.ID)

	_, err = isWaitForLBPoolActive(sess, lbID, *lbPool.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func resourceIBMISLBPoolRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	diag := lbPoolGet(context, d, meta, lbID, lbPoolID)
	if diag != nil {
		return diag
	}

	return nil
}

func lbPoolGet(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getLoadBalancerPoolOptions := &vpcv1.GetLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		ID:             &lbPoolID,
	}

	loadBalancerPool, response, err := sess.GetLoadBalancerPoolWithContext(context, getLoadBalancerPoolOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("name", loadBalancerPool.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-name").GetDiag()
	}
	if err = d.Set(isLBPool, lbPoolID); err != nil {
		err = fmt.Errorf("Error setting pool_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-pool_id").GetDiag()
	}
	if err = d.Set(isLBID, lbID); err != nil {
		err = fmt.Errorf("Error setting lb: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-lb").GetDiag()
	}
	if err = d.Set("algorithm", loadBalancerPool.Algorithm); err != nil {
		err = fmt.Errorf("Error setting algorithm: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-algorithm").GetDiag()
	}
	if err = d.Set("protocol", loadBalancerPool.Protocol); err != nil {
		err = fmt.Errorf("Error setting protocol: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-protocol").GetDiag()
	}
	if loadBalancerPool.HealthMonitor != nil {
		poolHealthMonitor := loadBalancerPool.HealthMonitor.(*vpcv1.LoadBalancerPoolHealthMonitor)
		if err = d.Set(isLBPoolHealthDelay, *poolHealthMonitor.Delay); err != nil {
			err = fmt.Errorf("Error setting health_delay: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_delay").GetDiag()
		}
		if err = d.Set(isLBPoolHealthRetries, *poolHealthMonitor.MaxRetries); err != nil {
			err = fmt.Errorf("Error setting health_retries: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_retries").GetDiag()
		}
		if err = d.Set(isLBPoolHealthTimeout, *poolHealthMonitor.Timeout); err != nil {
			err = fmt.Errorf("Error setting health_timeout: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_timeout").GetDiag()
		}
		if poolHealthMonitor.Type != nil {
			if err = d.Set(isLBPoolHealthType, *poolHealthMonitor.Type); err != nil {
				err = fmt.Errorf("Error setting health_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_type").GetDiag()
			}
		}
		if poolHealthMonitor.URLPath != nil {
			if err = d.Set(isLBPoolHealthMonitorURL, *poolHealthMonitor.URLPath); err != nil {
				err = fmt.Errorf("Error setting health_monitor_url: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_monitor_url").GetDiag()
			}
		}
		if poolHealthMonitor.Port != nil {
			if err = d.Set(isLBPoolHealthMonitorPort, *poolHealthMonitor.Port); err != nil {
				err = fmt.Errorf("Error setting health_monitor_port: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_monitor_port").GetDiag()
			}
		}
	}
	var healthMonitorMap map[string]interface{}
	if !core.IsNil(loadBalancerPool.HealthMonitor) {
		healthMonitorMap, err = ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorToMap(loadBalancerPool.HealthMonitor)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "health_monitor-to-map").GetDiag()
		}
	} else {
		healthMonitorMap = map[string]interface{}{}
	}
	if err = d.Set("health_monitor", []map[string]interface{}{healthMonitorMap}); err != nil {
		err = fmt.Errorf("Error setting health_monitor: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-health_monitor").GetDiag()
	}
	if !core.IsNil(loadBalancerPool.FailsafePolicy) {
		failsafePolicyMap, err := resourceIBMIsLbPoolLoadBalancerPoolFailsafePolicyToMap(loadBalancerPool.FailsafePolicy)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "failsafe_policy-to-map").GetDiag()
		}
		if err = d.Set("failsafe_policy", []map[string]interface{}{failsafePolicyMap}); err != nil {
			err = fmt.Errorf("Error setting failsafe_policy: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-failsafe_policy").GetDiag()
		}
	}
	if loadBalancerPool.SessionPersistence != nil {
		if err = d.Set(isLBPoolSessPersistenceType, *loadBalancerPool.SessionPersistence.Type); err != nil {
			err = fmt.Errorf("Error setting session_persistence_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_type").GetDiag()
		}
		if loadBalancerPool.SessionPersistence.CookieName != nil {
			if *loadBalancerPool.SessionPersistence.Type == "app_cookie" {
				if err = d.Set(isLBPoolSessPersistenceAppCookieName, *loadBalancerPool.SessionPersistence.CookieName); err != nil {
					err = fmt.Errorf("Error setting session_persistence_type: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_app_cookie_name").GetDiag()
				}
				if err = d.Set(isLBPoolSessPersistenceHttpCookieName, ""); err != nil {
					err = fmt.Errorf("Error setting session_persistence_http_cookie_name: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_http_cookie_name").GetDiag()
				}
			} else if *loadBalancerPool.SessionPersistence.Type == "http_cookie" {
				if err = d.Set(isLBPoolSessPersistenceHttpCookieName, *loadBalancerPool.SessionPersistence.CookieName); err != nil {
					err = fmt.Errorf("Error setting session_persistence_http_cookie_name: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_http_cookie_name").GetDiag()
				}
				if err = d.Set(isLBPoolSessPersistenceAppCookieName, ""); err != nil {
					err = fmt.Errorf("Error setting session_persistence_app_cookie_name: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_app_cookie_name").GetDiag()
				}
			}

		}
	} else {
		if err = d.Set(isLBPoolSessPersistenceType, ""); err != nil {
			err = fmt.Errorf("Error setting session_persistence_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_type").GetDiag()
		}
		if err = d.Set(isLBPoolSessPersistenceHttpCookieName, ""); err != nil {
			err = fmt.Errorf("Error setting session_persistence_http_cookie_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_http_cookie_name").GetDiag()
		}
		if err = d.Set(isLBPoolSessPersistenceAppCookieName, ""); err != nil {
			err = fmt.Errorf("Error setting session_persistence_app_cookie_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-session_persistence_app_cookie_name").GetDiag()
		}
	}
	if err = d.Set("provisioning_status", loadBalancerPool.ProvisioningStatus); err != nil {
		err = fmt.Errorf("Error setting provisioning_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-provisioning_status").GetDiag()
	}
	if err = d.Set(isLBPoolProxyProtocol, *loadBalancerPool.ProxyProtocol); err != nil {
		err = fmt.Errorf("Error setting proxy_protocol: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-proxy_protocol").GetDiag()
	}

	getLoadBalancerOptions := &vpcv1.GetLoadBalancerOptions{
		ID: &lbID,
	}
	lb, response, err := sess.GetLoadBalancerWithContext(context, getLoadBalancerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.RelatedCRN, *lb.CRN); err != nil {
		err = fmt.Errorf("Error setting related_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "read", "set-related_crn").GetDiag()
	}
	return nil
}

func resourceIBMISLBPoolUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "update", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	diag := lbPoolUpdate(context, d, meta, lbID, lbPoolID)
	if diag != nil {
		return diag
	}

	return resourceIBMISLBPoolRead(context, d, meta)
}

func lbPoolUpdate(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	hasChanged := false

	updateLoadBalancerPoolOptions := &vpcv1.UpdateLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		ID:             &lbPoolID,
	}

	loadBalancerPoolPatchModel := &vpcv1.LoadBalancerPoolPatch{}

	lBPoolHealthMonitorPortRemoved := false
	requestWasRemoved := false
	headersWereCleared := false
	isFailSafePolicyTargetNull := false
	hasFailSafeChanged := false
	if d.HasChange("failsafe_policy") {
		failsafePolicy := &vpcv1.LoadBalancerPoolFailsafePolicyPatch{}
		hasFailSafeChanged = true
		if d.HasChange("failsafe_policy.0.action") {
			failsafepolicyAction := d.Get("failsafe_policy.0.action").(string)
			failsafePolicy.Action = &failsafepolicyAction
		}
		if d.HasChange("failsafe_policy.0.target") {
			failsafepolicyAction := d.Get("failsafe_policy.0.action").(string)
			failsafePolicy.Action = &failsafepolicyAction
			targetPatch := &vpcv1.LoadBalancerPoolFailsafePolicyTargetPatch{}
			failsafepolicyTargetId := d.Get("failsafe_policy.0.target.0.id").(string)
			if failsafepolicyTargetId == "null" {
				isFailSafePolicyTargetNull = true
				var nullStringPtr *string
				targetPatch.ID = nullStringPtr
				failsafePolicy.Target = targetPatch
			} else {
				targetPatch.ID = &failsafepolicyTargetId
				failsafePolicy.Target = targetPatch

			}
			failsafepolicyTargetHref := d.Get("failsafe_policy.0.target.0.href").(string)
			if failsafepolicyTargetHref == "null" {
				isFailSafePolicyTargetNull = true
				var nullStringPtr *string
				targetPatch.Href = nullStringPtr
				failsafePolicy.Target = targetPatch
			} else {
				targetPatch.Href = &failsafepolicyTargetHref
				failsafePolicy.Target = targetPatch
			}
		}
		loadBalancerPoolPatchModel.FailsafePolicy = failsafePolicy
	}

	if d.HasChange(isLBPoolHealthDelay) || d.HasChange(isLBPoolHealthRetries) ||
		d.HasChange(isLBPoolHealthTimeout) || d.HasChange(isLBPoolHealthType) || d.HasChange(isLBPoolHealthMonitorURL) || d.HasChange(isLBPoolHealthMonitorPort) || d.HasChange("failsafe_policy") || d.HasChange("health_monitor") {

		delay := int64(d.Get(isLBPoolHealthDelay).(int))
		maxretries := int64(d.Get(isLBPoolHealthRetries).(int))
		timeout := int64(d.Get(isLBPoolHealthTimeout).(int))
		healthtype := d.Get(isLBPoolHealthType).(string)
		urlpath := d.Get(isLBPoolHealthMonitorURL).(string)
		healthMonitorTemplate := &vpcv1.LoadBalancerPoolHealthMonitorPatch{
			Delay:      &delay,
			MaxRetries: &maxretries,
			Timeout:    &timeout,
			Type:       &healthtype,
			URLPath:    &urlpath,
		}
		port := int64(d.Get(isLBPoolHealthMonitorPort).(int))
		if port > int64(0) {
			healthMonitorTemplate.Port = &port
		} else {
			lBPoolHealthMonitorPortRemoved = true
		}

		hm, _ := d.Get("health_monitor.0").(map[string]interface{})
		if hm == nil {
			hm = map[string]interface{}{}
		}

		if d.HasChange("health_monitor.0.request") {
			reqList, _ := hm["request"].([]interface{})
			if len(reqList) > 0 {
				reqMap := reqList[0].(map[string]interface{})
				// Detect explicit header removal so we can send [] (not null) later.
				if d.HasChange("health_monitor.0.request.0.headers") {
					newHeaders, _ := reqMap["headers"].([]interface{})
					if len(newHeaders) == 0 {
						headersWereCleared = true
					}
				}
				requestModel, err :=
					ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestPatch(reqMap)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(
						err,
						err.Error(),
						"ibm_is_lb_pool",
						"update",
						"parse-health_monitor-request",
					).GetDiag()
				}
				healthMonitorTemplate.Request = requestModel
			} else {
				// request block was entirely removed; signal a null to the API.
				requestWasRemoved = true
			}
		}

		if d.HasChange("health_monitor.0.response") {
			if hm["response"] != nil && len(hm["response"].([]interface{})) > 0 {
				responseMap := hm["response"].([]interface{})[0].(map[string]interface{})
				hasBodyRegex := responseMap["body_regex"] != nil && responseMap["body_regex"].(string) != ""
				hasCodes := responseMap["codes"] != nil && len(responseMap["codes"].([]interface{})) > 0
				if hasBodyRegex || hasCodes {
					responseModel, err :=
						ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsResponsePatch(responseMap)
					if err != nil {
						return flex.DiscriminatedTerraformErrorf(
							err,
							err.Error(),
							"ibm_is_lb_pool",
							"update",
							"parse-health_monitor-response",
						).GetDiag()
					}

					healthMonitorTemplate.Response = responseModel
				}
			}
		}

		loadBalancerPoolPatchModel.HealthMonitor = healthMonitorTemplate
		hasChanged = true
	}

	sessionPersistenceRemoved := false
	if d.HasChange(isLBPoolSessPersistenceType) || d.HasChange(isLBPoolSessPersistenceAppCookieName) {
		sesspersistancetype := d.Get(isLBPoolSessPersistenceType).(string)
		sessPersistanceCookieName := d.Get(isLBPoolSessPersistenceAppCookieName).(string)
		sessionPersistence := &vpcv1.LoadBalancerPoolSessionPersistencePatch{}
		if sesspersistancetype != "" {
			sessionPersistence.Type = &sesspersistancetype
			if sessPersistanceCookieName != "" {
				sessionPersistence.CookieName = &sessPersistanceCookieName
			}
		} else {
			sessionPersistenceRemoved = true
		}

		loadBalancerPoolPatchModel.SessionPersistence = sessionPersistence

		hasChanged = true
	}

	if d.HasChange(isLBPoolProxyProtocol) {
		proxyProtocol := d.Get(isLBPoolProxyProtocol).(string)
		loadBalancerPoolPatchModel.ProxyProtocol = &proxyProtocol
		hasChanged = true
	}

	if d.HasChange(isLBPoolName) || d.HasChange(isLBPoolAlgorithm) || d.HasChange(isLBPoolProtocol) || hasChanged || hasFailSafeChanged {
		name := d.Get(isLBPoolName).(string)
		algorithm := d.Get(isLBPoolAlgorithm).(string)
		protocol := d.Get(isLBPoolProtocol).(string)

		loadBalancerPoolPatchModel.Algorithm = &algorithm
		loadBalancerPoolPatchModel.Name = &name
		loadBalancerPoolPatchModel.Protocol = &protocol

		isLBKey := "load_balancer_key_" + lbID
		conns.IbmMutexKV.Lock(isLBKey)
		defer conns.IbmMutexKV.Unlock(isLBKey)
		_, err := isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		LoadBalancerPoolPatch, err := loadBalancerPoolPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("loadBalancerPoolPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if sessionPersistenceRemoved {
			LoadBalancerPoolPatch["session_persistence"] = nil
		}
		if hmPatch, ok := LoadBalancerPoolPatch["health_monitor"].(map[string]interface{}); ok {
			if lBPoolHealthMonitorPortRemoved {
				hmPatch["port"] = nil
			}
			if requestWasRemoved {
				hmPatch["request"] = nil
			}
			// The IBM VPC API clears headers only when sent as [] (empty array), not null.
			// SDK asPatch() produces null for a non-nil empty slice, so we patch it here.
			if headersWereCleared {
				if reqPatch, ok := hmPatch["request"].(map[string]interface{}); ok {
					reqPatch["headers"] = []interface{}{}
				}
			}
		}
		if isFailSafePolicyTargetNull {
			if fpPatch, ok := LoadBalancerPoolPatch["failsafe_policy"].(map[string]interface{}); ok {
				fpPatch["target"] = nil
			}
		}

		updateLoadBalancerPoolOptions.LoadBalancerPoolPatch = LoadBalancerPoolPatch

		_, _, err = sess.UpdateLoadBalancerPoolWithContext(context, updateLoadBalancerPoolOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISLBPoolDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "delete", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	diag := lbPoolDelete(context, d, meta, lbID, lbPoolID)
	if diag != nil {
		return diag
	}

	return nil
}

func lbPoolDelete(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getLoadBalancerPoolOptions := &vpcv1.GetLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		ID:             &lbPoolID,
	}
	_, response, err := sess.GetLoadBalancerPoolWithContext(context, getLoadBalancerPoolOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteLoadBalancerPoolOptions := &vpcv1.DeleteLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		ID:             &lbPoolID,
	}
	response, err = sess.DeleteLoadBalancerPoolWithContext(context, deleteLoadBalancerPoolOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteLoadBalancerPoolWithContext failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForLBPoolDeleted(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolDeleted failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISLBPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "exists", "sep-id-parts")
	}
	if len(parts) != 2 {
		err = fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of lbID/lbPoolID", d.Id())
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "exists", "sep-id-parts")
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	exists, err := lbPoolExists(d, meta, lbID, lbPoolID)
	return exists, err

}

func lbPoolExists(d *schema.ResourceData, meta interface{}, lbID, lbPoolID string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	getLoadBalancerPoolOptions := &vpcv1.GetLoadBalancerPoolOptions{
		LoadBalancerID: &lbID,
		ID:             &lbPoolID,
	}
	_, response, err := sess.GetLoadBalancerPool(getLoadBalancerPoolOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPool failed: %s", err.Error()), "ibm_is_lb_pool", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func isWaitForLBPoolActive(sess *vpcv1.VpcV1, lbId, lbPoolId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool (%s) to be available.", lbPoolId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolCreatePending, isLBPoolUpdatePending, isLBPoolMaintainancePending},
		Target:     []string{isLBPoolActive, ""},
		Refresh:    isLBPoolRefreshFunc(sess, lbId, lbPoolId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolRefreshFunc(sess *vpcv1.VpcV1, lbId, lbPoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlbpOptions := &vpcv1.GetLoadBalancerPoolOptions{
			LoadBalancerID: &lbId,
			ID:             &lbPoolId,
		}
		lbPool, response, err := sess.GetLoadBalancerPool(getlbpOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer Pool: %s\n%s", err, response)
		}

		if *lbPool.ProvisioningStatus == isLBPoolActive || *lbPool.ProvisioningStatus == isLBPoolFailed {
			return lbPool, isLBPoolActive, nil
		}

		return lbPool, *lbPool.ProvisioningStatus, nil
	}
}

func isWaitForLBPoolDeleted(lbc *vpcv1.VpcV1, lbId, lbPoolId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", lbPoolId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolUpdatePending, isLBPoolMaintainancePending, isLBPoolDeletePending},
		Target:     []string{isLBPoolDeleteDone, ""},
		Refresh:    isLBPoolDeleteRefreshFunc(lbc, lbId, lbPoolId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolDeleteRefreshFunc(lbc *vpcv1.VpcV1, lbId, lbPoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] is lb pool delete function here")
		getlbpOptions := &vpcv1.GetLoadBalancerPoolOptions{
			LoadBalancerID: &lbId,
			ID:             &lbPoolId,
		}
		lbPool, response, err := lbc.GetLoadBalancerPool(getlbpOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return lbPool, isLBPoolDeleteDone, nil
			}
			return nil, "", fmt.Errorf("[ERROR] The vpc load balancer pool %s failed to delete: %s\n%s", lbPoolId, err, response)
		}
		return lbPool, isLBPoolDeletePending, nil
	}
}

func resourceIBMIsLbPoolMapToLoadBalancerPoolFailsafePolicyPrototype(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolFailsafePolicyPrototype, error) {
	model := &vpcv1.LoadBalancerPoolFailsafePolicyPrototype{}
	if modelMap["action"] != nil && modelMap["action"].(string) != "" {
		model.Action = core.StringPtr(modelMap["action"].(string))
	}
	if modelMap["target"] != nil && len(modelMap["target"].([]interface{})) > 0 {
		TargetModel, err := resourceIBMIsLbPoolMapToLoadBalancerPoolIdentity(modelMap["target"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Target = TargetModel
	}
	return model, nil
}

func resourceIBMIsLbPoolMapToLoadBalancerPoolIdentity(modelMap map[string]interface{}) (vpcv1.LoadBalancerPoolIdentityIntf, error) {
	model := &vpcv1.LoadBalancerPoolIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func resourceIBMIsLbPoolLoadBalancerPoolFailsafePolicyToMap(model *vpcv1.LoadBalancerPoolFailsafePolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["action"] = *model.Action
	modelMap["healthy_member_threshold_count"] = flex.IntValue(model.HealthyMemberThresholdCount)
	if model.Target != nil {
		targetMap, err := resourceIBMIsLbPoolLoadBalancerPoolReferenceToMap(model.Target)
		if err != nil {
			return modelMap, err
		}
		modelMap["target"] = []map[string]interface{}{targetMap}
	}
	return modelMap, nil
}

func resourceIBMIsLbPoolLoadBalancerPoolReferenceToMap(model *vpcv1.LoadBalancerPoolReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := resourceIBMIsLbPoolDeletedToMap(model.Deleted)
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

func resourceIBMIsLbPoolDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func suppressNullTarget(k, old, new string, d *schema.ResourceData) bool {
	// If resource already exists (has an ID) and new value is "null" while old is empty
	// then suppress the diff
	if new != old && new == "null" && old == "" && d.Id() != "" {
		return true
	}

	oldId, newId := d.GetChange("failsafe_policy.0.target.0.id")
	oldHref, newHref := d.GetChange("failsafe_policy.0.target.0.href")

	// Check id field
	if newId.(string) == "null" && oldId.(string) == "" && d.Id() != "" {
		return true
	}

	// Check href field
	if newHref.(string) == "null" && oldHref.(string) == "" && d.Id() != "" {
		return true
	}

	return false
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorPrototype(
	modelMap map[string]interface{},
) (
	*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype,
	*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype,
	error,
) {
	var (
		requestModel  *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype
		responseModel *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype
		err           error
	)

	if len(modelMap) == 0 {
		return nil, nil, nil
	}

	if modelMap["request"] != nil && len(modelMap["request"].([]interface{})) > 0 {
		requestModel, err =
			ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype(
				modelMap["request"].([]interface{})[0].(map[string]interface{}),
			)
		if err != nil {
			return requestModel, responseModel, err
		}
	}

	if modelMap["response"] != nil && len(modelMap["response"].([]interface{})) > 0 {
		responseModel, err =
			ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype(
				modelMap["response"].([]interface{})[0].(map[string]interface{}),
			)
		if err != nil {
			return requestModel, responseModel, err
		}
	}

	return requestModel, responseModel, nil
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype, error) {
	model := &vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPrototype{}
	if modelMap["headers"] != nil {
		headers := []vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype{}
		for _, headersItem := range modelMap["headers"].([]interface{}) {
			headersItemModel, err := ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype(headersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			headers = append(headers, *headersItemModel)
		}
		if len(headers) > 0 {
			model.HeadersVar = headers
		}
	}
	if modelMap["method"] != nil && modelMap["method"].(string) != "" {
		model.Method = core.StringPtr(modelMap["method"].(string))
	}
	if modelMap["body"] != nil && modelMap["body"].(string) != "" {
		model.Body = core.StringPtr(modelMap["body"].(string))
	}
	return model, nil
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype, error) {
	model := &vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype{}
	if modelMap["field"] != nil && modelMap["field"].(string) != "" {
		model.Field = core.StringPtr(modelMap["field"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype, error) {
	model := &vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePrototype{}
	if modelMap["body_regex"] != nil && modelMap["body_regex"].(string) != "" {
		model.BodyRegex = core.StringPtr(modelMap["body_regex"].(string))
	}
	if modelMap["codes"] != nil {
		codes := []string{}
		for _, codesItem := range modelMap["codes"].([]interface{}) {
			codes = append(codes, codesItem.(string))
		}
		if len(codes) > 0 {
			model.Codes = codes
		}
	}
	return model, nil
}

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorToMap(model vpcv1.LoadBalancerPoolHealthMonitorIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps); ok {
		return ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps))
	} else if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP); ok {
		return ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeTCPToMap(model.(*vpcv1.LoadBalancerPoolHealthMonitorTypeTCP))
	} else if _, ok := model.(*vpcv1.LoadBalancerPoolHealthMonitor); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.LoadBalancerPoolHealthMonitor)

		if model.Request != nil {
			requestMap, err := ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
			if err != nil {
				return modelMap, err
			}
			modelMap["request"] = []map[string]interface{}{requestMap}
		}
		if model.Response != nil {
			responseMap, err := ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
			if err != nil {
				return modelMap, err
			}
			modelMap["response"] = []map[string]interface{}{responseMap}
		}

		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.LoadBalancerPoolHealthMonitorIntf subtype encountered")
	}
}

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequest) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Body != nil {
		modelMap["body"] = *model.Body
	}
	if model.HeadersVar != nil {
		headers := []map[string]interface{}{}
		for _, headersItem := range model.HeadersVar {
			headersItemMap, err := ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(&headersItem) // #nosec G601
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

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeader) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Field != nil {
		modelMap["field"] = *model.Field
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BodyRegex != nil {
		modelMap["body_regex"] = *model.BodyRegex
	}
	if model.Codes != nil {
		modelMap["codes"] = model.Codes
	}
	return modelMap, nil
}

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttps) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Request != nil {
		requestMap, err := ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsRequestToMap(model.Request)
		if err != nil {
			return modelMap, err
		}
		modelMap["request"] = []map[string]interface{}{requestMap}
	}
	if model.Response != nil {
		responseMap, err := ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeHttphttpsResponseToMap(model.Response)
		if err != nil {
			return modelMap, err
		}
		modelMap["response"] = []map[string]interface{}{responseMap}
	}
	return modelMap, nil
}

func ResourceIBMIsLbPoolLoadBalancerPoolHealthMonitorTypeTCPToMap(model *vpcv1.LoadBalancerPoolHealthMonitorTypeTCP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	return modelMap, nil
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestPatch(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPatch, error) {
	model := &vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestPatch{}
	if modelMap["body"] != nil && modelMap["body"].(string) != "" {
		model.Body = core.StringPtr(modelMap["body"].(string))
	}
	// Always initialize headers as a non-nil slice so that removing all headers
	// sends "headers": null in the PATCH body, telling the API to clear them.
	headers := []vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype{}
	if modelMap["headers"] != nil {
		for _, headersItem := range modelMap["headers"].([]interface{}) {
			headersItemModel, err := ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsRequestHeaderPrototype(headersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			headers = append(headers, *headersItemModel)
		}
	}
	model.HeadersVar = headers
	model.Method = core.StringPtr(modelMap["method"].(string))
	return model, nil
}

func ResourceIBMIsLbPoolMapToLoadBalancerPoolHealthMonitorTypeHttphttpsResponsePatch(modelMap map[string]interface{}) (*vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePatch, error) {
	model := &vpcv1.LoadBalancerPoolHealthMonitorTypeHttphttpsResponsePatch{}
	if modelMap["body_regex"] != nil && modelMap["body_regex"].(string) != "" {
		model.BodyRegex = core.StringPtr(modelMap["body_regex"].(string))
	}
	if modelMap["codes"] != nil {
		codes := []string{}
		for _, codesItem := range modelMap["codes"].([]interface{}) {
			codes = append(codes, codesItem.(string))
		}
		if len(codes) > 0 {
			model.Codes = codes
		}
	}
	return model, nil
}
