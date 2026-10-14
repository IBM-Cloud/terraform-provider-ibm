// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isDynamicRouteServerPeerGroupPending   = "pending"
	isDynamicRouteServerPeerGroupUpdating  = "updating"
	isDynamicRouteServerPeerGroupWaiting   = "waiting"
	isDynamicRouteServerPeerGroupDeleting  = "deleting"
	isDynamicRouteServerPeerGroupStable    = "stable"
	isDynamicRouteServerPeerGroupFailed    = "failed"
	isDynamicRouteServerPeerGroupSuspended = "suspended"
	isDynamicRouteServerPeerGroupDeleted   = "deleted"
)

func ResourceIBMIsDynamicRouteServerPeerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"health_monitor": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The health monitoring configuration used for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "health_monitor_mode"),
							Description:  "The mode used for this health monitor: `bfd`, `bgp` or `none`.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							MinItems:    1,
							Description: "The roles used for this dynamic route server peer group health monitor configuration. Allowed values are `next_hop` and `router`.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "name"),
				Description:  "The name for this dynamic route server peer group.",
			},
			"peers": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				Description: "The peers created with this peer group. Peers can be added later with the `ibm_is_dynamic_route_server_peer_group_peer` resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asn": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The autonomous system number (ASN) for this peer.",
						},
						"endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							MinItems:    1,
							MaxItems:    1,
							Description: "The endpoint for this peer. Specify either the reserved IP (`id` or `href`) or an IP address with an optional gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for the reserved IP used as the endpoint.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The URL for the reserved IP used as the endpoint.",
									},
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The IP address of the endpoint.",
									},
									"gateway": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										MaxItems:    1,
										Description: "The gateway used to reach the endpoint address.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The IP address of the gateway.",
												},
											},
										},
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name for this peer. The name must not be used by another peer in the peer group.",
						},
						"priority": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The priority of this peer. Lower values have higher priority.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The state of this peer: `enabled` or `disabled`.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this peer.",
						},
					},
				},
			},
			"prefix_watcher": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The prefix watcher configuration for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The prefixes excluded from monitoring.",
							Elem:        resourceIBMIsDynamicRouteServerPeerGroupPrefixSchema(),
						},
						"monitored_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The prefixes monitored by this peer group.",
							Elem:        resourceIBMIsDynamicRouteServerPeerGroupPrefixSchema(),
						},
					},
				},
			},
			"roles": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				Description: "The roles for this dynamic route server peer group. Allowed values are `next_hop` and `router`.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "state"),
				Description:  "The state of this dynamic route server peer group: `enabled` or `disabled`.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this dynamic route server peer group was created.",
			},
			"creator": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The types of resource that created peers in this peer group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group.",
			},
			"lifecycle_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `lifecycle_state` (if any).",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the dynamic route server peer group.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this dynamic route server peer group.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status reasons for the dynamic route server peer group.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server peer group.",
			},
			"etag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ETag identifier for the dynamic route server peer group.",
			},
		},
	}
}

func resourceIBMIsDynamicRouteServerPeerGroupPrefixSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP prefix in CIDR format.",
			},
			"ge": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The minimum prefix length to match.",
			},
			"le": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The maximum prefix length to match.",
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "health_monitor_mode",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "bfd, bgp, none",
		},
		validate.ValidateSchema{
			Identifier:                 "state",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "disabled, enabled",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerPeerGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	createOptions := &vpcv1.CreateDynamicRouteServerPeerGroupOptions{}
	createOptions.SetDynamicRouteServerID(dynamicRouteServerID)

	healthMonitorModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(d.Get("health_monitor.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-health_monitor").GetDiag()
	}
	createOptions.SetHealthMonitor(healthMonitorModel)

	peers := []vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf{}
	for _, v := range d.Get("peers").([]interface{}) {
		peerMap, ok := v.(map[string]interface{})
		if !ok {
			return flex.DiscriminatedTerraformErrorf(fmt.Errorf("peers entry is not a block"), "peers entry is not a block", "ibm_is_dynamic_route_server_peer_group", "create", "parse-peers").GetDiag()
		}
		peerModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(peerMap)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-peers").GetDiag()
		}
		peers = append(peers, peerModel)
	}
	createOptions.SetPeers(peers)

	roles := []string{}
	for _, v := range d.Get("roles").([]interface{}) {
		roles = append(roles, v.(string))
	}
	createOptions.SetRoles(roles)

	if _, ok := d.GetOk("name"); ok {
		createOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("prefix_watcher"); ok {
		prefixWatcherModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(d.Get("prefix_watcher.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-prefix_watcher").GetDiag()
		}
		createOptions.SetPrefixWatcher(prefixWatcherModel)
	}
	if _, ok := d.GetOk("state"); ok {
		createOptions.SetState(d.Get("state").(string))
	}

	peerGroup, _, err := vpcClient.CreateDynamicRouteServerPeerGroupWithContext(context, createOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if peerGroup == nil || peerGroup.ID == nil {
		return flex.DiscriminatedTerraformErrorf(fmt.Errorf("peer group response has no id"), "peer group response has no id", "ibm_is_dynamic_route_server_peer_group", "create", "peer-group-id").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", dynamicRouteServerID, *peerGroup.ID))

	_, err = isWaitForDynamicRouteServerPeerGroupStable(vpcClient, dynamicRouteServerID, *peerGroup.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "wait-stable").GetDiag()
	}

	return resourceIBMIsDynamicRouteServerPeerGroupRead(context, d, meta)
}

// resourceIBMIsDynamicRouteServerPeerGroupParseID splits the resource ID "<dynamic_route_server_id>/<peer_group_id>".
func resourceIBMIsDynamicRouteServerPeerGroupParseID(id string) (string, string, error) {
	parts, err := flex.SepIdParts(id, "/")
	if err != nil {
		return "", "", err
	}
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid ID %q: expected <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>", id)
	}
	return parts[0], parts[1], nil
}

func resourceIBMIsDynamicRouteServerPeerGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, err := resourceIBMIsDynamicRouteServerPeerGroupParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "sep-id-parts").GetDiag()
	}

	getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetID(peerGroupID)

	peerGroup, response, err := vpcClient.GetDynamicRouteServerPeerGroupWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// The parent identifier is derived from the ID so that import works.
	if err = d.Set("dynamic_route_server_id", dynamicRouteServerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_id: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-dynamic_route_server_id").GetDiag()
	}
	if err = d.Set("dynamic_route_server_peer_group_id", peerGroupID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_id: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-dynamic_route_server_peer_group_id").GetDiag()
	}
	if peerGroup.HealthMonitor != nil {
		healthMonitorMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(peerGroup.HealthMonitor)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "health_monitor-to-map").GetDiag()
		}
		if err = d.Set("health_monitor", []map[string]interface{}{healthMonitorMap}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_monitor: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-health_monitor").GetDiag()
		}
	}
	if !core.IsNil(peerGroup.Name) {
		if err = d.Set("name", peerGroup.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-name").GetDiag()
		}
	}
	peers := []map[string]interface{}{}
	for _, peerItem := range peerGroup.Peers {
		peerMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(peerItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "peers-to-map").GetDiag()
		}
		peers = append(peers, peerMap)
	}
	if err = d.Set("peers", peers); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting peers: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-peers").GetDiag()
	}
	if peerGroup.PrefixWatcher != nil {
		prefixWatcherMap := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(peerGroup.PrefixWatcher)
		if err = d.Set("prefix_watcher", []map[string]interface{}{prefixWatcherMap}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting prefix_watcher: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-prefix_watcher").GetDiag()
		}
	}
	if err = d.Set("roles", peerGroup.Roles); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting roles: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-roles").GetDiag()
	}
	if !core.IsNil(peerGroup.State) {
		if err = d.Set("state", peerGroup.State); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting state: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-state").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(peerGroup.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("creator", peerGroup.Creator); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting creator: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-creator").GetDiag()
	}
	if err = d.Set("href", peerGroup.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for i := range peerGroup.LifecycleReasons {
		lifecycleReasons = append(lifecycleReasons, ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(&peerGroup.LifecycleReasons[i]))
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", peerGroup.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", peerGroup.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("status", peerGroup.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for i := range peerGroup.StatusReasons {
		statusReasons = append(statusReasons, ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(&peerGroup.StatusReasons[i]))
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-status_reasons").GetDiag()
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, err := resourceIBMIsDynamicRouteServerPeerGroupParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "sep-id-parts").GetDiag()
	}

	hasChange := false
	patchVals := &vpcv1.DynamicRouteServerPeerGroupPatch{}
	if d.HasChange("health_monitor.0.mode") {
		healthMonitor, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(d.Get("health_monitor.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "parse-health_monitor").GetDiag()
		}
		patchVals.HealthMonitor = healthMonitor
		hasChange = true
	}
	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("prefix_watcher") {
		prefixWatcherMap := map[string]interface{}{}
		if v, ok := d.Get("prefix_watcher.0").(map[string]interface{}); ok {
			prefixWatcherMap = v
		}
		prefixWatcher, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(prefixWatcherMap)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "parse-prefix_watcher").GetDiag()
		}
		patchVals.PrefixWatcher = prefixWatcher
		hasChange = true
	}
	if d.HasChange("state") {
		patchVals.State = core.StringPtr(d.Get("state").(string))
		hasChange = true
	}

	if hasChange {
		updateOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupOptions{}
		updateOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		updateOptions.SetID(peerGroupID)
		updateOptions.SetIfMatch(d.Get("etag").(string))
		updateOptions.DynamicRouteServerPeerGroupPatch = ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupWithContext(context, updateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForDynamicRouteServerPeerGroupStable(vpcClient, dynamicRouteServerID, peerGroupID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "wait-stable").GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, err := resourceIBMIsDynamicRouteServerPeerGroupParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "delete", "sep-id-parts").GetDiag()
	}

	deleteOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupOptions{}
	deleteOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	deleteOptions.SetID(peerGroupID)
	if etag, ok := d.GetOk("etag"); ok {
		deleteOptions.SetIfMatch(etag.(string))
	}

	_, response, err := vpcClient.DeleteDynamicRouteServerPeerGroupWithContext(context, deleteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerPeerGroupDeleted(vpcClient, dynamicRouteServerID, peerGroupID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "delete", "wait-deleted").GetDiag()
	}

	d.SetId("")

	return nil
}

func isWaitForDynamicRouteServerPeerGroupStable(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group (%s) to be stable.", peerGroupID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupPending, isDynamicRouteServerPeerGroupUpdating, isDynamicRouteServerPeerGroupWaiting},
		Target:     []string{isDynamicRouteServerPeerGroupStable, isDynamicRouteServerPeerGroupFailed, isDynamicRouteServerPeerGroupSuspended},
		Refresh:    isDynamicRouteServerPeerGroupRefreshFunc(sess, dynamicRouteServerID, peerGroupID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerPeerGroupRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetID(peerGroupID)
		peerGroup, response, err := sess.GetDynamicRouteServerPeerGroup(getOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting dynamic route server peer group: %s\n%s", err, response)
		}
		state := ""
		if peerGroup.LifecycleState != nil {
			state = *peerGroup.LifecycleState
		}
		if state == isDynamicRouteServerPeerGroupFailed || state == isDynamicRouteServerPeerGroupSuspended {
			return peerGroup, state, fmt.Errorf("[ERROR] The dynamic route server peer group (%s) is in %s state", peerGroupID, state)
		}
		if state == "" {
			return peerGroup, isDynamicRouteServerPeerGroupStable, nil
		}
		return peerGroup, state, nil
	}
}

func isWaitForDynamicRouteServerPeerGroupDeleted(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group (%s) to be deleted.", peerGroupID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupDeleting, isDynamicRouteServerPeerGroupStable, isDynamicRouteServerPeerGroupPending, isDynamicRouteServerPeerGroupUpdating, isDynamicRouteServerPeerGroupWaiting},
		Target:     []string{isDynamicRouteServerPeerGroupDeleted, isDynamicRouteServerPeerGroupFailed},
		Refresh:    isDynamicRouteServerPeerGroupDeleteRefreshFunc(sess, dynamicRouteServerID, peerGroupID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerPeerGroupDeleteRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetID(peerGroupID)
		peerGroup, response, err := sess.GetDynamicRouteServerPeerGroup(getOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return peerGroup, isDynamicRouteServerPeerGroupDeleted, nil
			}
			return nil, "", fmt.Errorf("[ERROR] Error getting dynamic route server peer group: %s\n%s", err, response)
		}
		state := ""
		if peerGroup.LifecycleState != nil {
			state = *peerGroup.LifecycleState
		}
		if state == isDynamicRouteServerPeerGroupFailed {
			return peerGroup, state, fmt.Errorf("[ERROR] The dynamic route server peer group (%s) failed to delete", peerGroupID)
		}
		if state == "" {
			return peerGroup, isDynamicRouteServerPeerGroupDeleting, nil
		}
		return peerGroup, state, nil
	}
}

// ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype builds the health monitor prototype.
// The next hop and router prototypes share the same shape, so the base prototype is used for both.
func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototype{}
	if mode, ok := modelMap["mode"].(string); ok && mode != "" {
		model.Mode = core.StringPtr(mode)
	}
	rolesList, _ := modelMap["roles"].([]interface{})
	if len(rolesList) == 0 {
		return nil, fmt.Errorf("health_monitor roles is required")
	}
	roles := []string{}
	for _, role := range rolesList {
		roleString, ok := role.(string)
		if !ok || roleString == "" {
			return nil, fmt.Errorf("health_monitor roles must be non-empty strings")
		}
		roles = append(roles, roleString)
	}
	model.Roles = roles
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPatch{}
	if mode, ok := modelMap["mode"].(string); ok && mode != "" {
		model.Mode = core.StringPtr(mode)
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPrototype{}
	excluded, monitored, err := resourceIBMIsDynamicRouteServerPeerGroupMapToPrefixLists(modelMap)
	if err != nil {
		return nil, err
	}
	model.ExcludedPrefixes = excluded
	model.MonitoredPrefixes = monitored
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPatch{}
	excluded, monitored, err := resourceIBMIsDynamicRouteServerPeerGroupMapToPrefixLists(modelMap)
	if err != nil {
		return nil, err
	}
	model.ExcludedPrefixes = excluded
	model.MonitoredPrefixes = monitored
	return model, nil
}

// resourceIBMIsDynamicRouteServerPeerGroupMapToPrefixLists always returns non-nil slices so that a
// patch removing every prefix sends an empty list instead of omitting the property.
func resourceIBMIsDynamicRouteServerPeerGroupMapToPrefixLists(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype, error) {
	excluded := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
	monitored := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{}
	if list, ok := modelMap["excluded_prefixes"].([]interface{}); ok {
		for _, item := range list {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("excluded_prefixes entry is not a block")
			}
			prefix, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(itemMap)
			if err != nil {
				return nil, nil, err
			}
			excluded = append(excluded, *prefix)
		}
	}
	if list, ok := modelMap["monitored_prefixes"].([]interface{}); ok {
		for _, item := range list {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("monitored_prefixes entry is not a block")
			}
			prefix, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(itemMap)
			if err != nil {
				return nil, nil, err
			}
			monitored = append(monitored, *prefix)
		}
	}
	return excluded, monitored, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, error) {
	prefix, _ := modelMap["prefix"].(string)
	if prefix == "" {
		return nil, fmt.Errorf("prefix is required for an excluded prefix")
	}
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{Prefix: core.StringPtr(prefix)}
	if ge, ok := modelMap["ge"].(int); ok && ge != 0 {
		model.Ge = core.Int64Ptr(int64(ge))
	}
	if le, ok := modelMap["le"].(int); ok && le != 0 {
		model.Le = core.Int64Ptr(int64(le))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype, error) {
	prefix, _ := modelMap["prefix"].(string)
	if prefix == "" {
		return nil, fmt.Errorf("prefix is required for a monitored prefix")
	}
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{Prefix: core.StringPtr(prefix)}
	if ge, ok := modelMap["ge"].(int); ok && ge != 0 {
		model.Ge = core.Int64Ptr(int64(ge))
	}
	if le, ok := modelMap["le"].(int); ok && le != 0 {
		model.Le = core.Int64Ptr(int64(le))
	}
	return model, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap flattens the health monitor oneOf (next hop, router or base).
func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf) (map[string]interface{}, error) {
	var mode *string
	var roles []string
	switch monitor := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop:
		if monitor == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf subtype encountered")
		}
		mode, roles = monitor.Mode, monitor.Roles
	case *vpcv1.DynamicRouteServerPeerGroupHealthMonitorRouter:
		if monitor == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf subtype encountered")
		}
		mode, roles = monitor.Mode, monitor.Roles
	case *vpcv1.DynamicRouteServerPeerGroupHealthMonitor:
		if monitor == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf subtype encountered")
		}
		mode, roles = monitor.Mode, monitor.Roles
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf subtype encountered")
	}
	modelMap := make(map[string]interface{})
	if mode != nil {
		modelMap["mode"] = *mode
	}
	if roles == nil {
		roles = []string{}
	}
	modelMap["roles"] = roles
	return modelMap, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap flattens any peer subtype into the prototype-shaped `peers` block of the resource.
func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (map[string]interface{}, error) {
	fullMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(model)
	if err != nil {
		return nil, err
	}
	modelMap := make(map[string]interface{})
	for _, key := range []string{"asn", "name", "priority", "state", "id"} {
		if value, ok := fullMap[key]; ok {
			modelMap[key] = value
		}
	}
	if endpoints, ok := fullMap["endpoint"].([]map[string]interface{}); ok && len(endpoints) > 0 {
		endpointMap := make(map[string]interface{})
		for _, key := range []string{"id", "href", "address", "gateway"} {
			if value, ok := endpoints[0][key]; ok {
				endpointMap[key] = value
			}
		}
		modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcher) map[string]interface{} {
	modelMap := make(map[string]interface{})
	excluded := []map[string]interface{}{}
	monitored := []map[string]interface{}{}
	if model != nil {
		for i := range model.ExcludedPrefixes {
			excluded = append(excluded, ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(&model.ExcludedPrefixes[i]))
		}
		for i := range model.MonitoredPrefixes {
			monitored = append(monitored, ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(&model.MonitoredPrefixes[i]))
		}
	}
	modelMap["excluded_prefixes"] = excluded
	modelMap["monitored_prefixes"] = monitored
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefix) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.Ge != nil {
		modelMap["ge"] = flex.IntValue(model.Ge)
	}
	if model.Le != nil {
		modelMap["le"] = flex.IntValue(model.Le)
	}
	if model.Prefix != nil {
		modelMap["prefix"] = *model.Prefix
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefix) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.Ge != nil {
		modelMap["ge"] = flex.IntValue(model.Ge)
	}
	if model.Le != nil {
		modelMap["le"] = flex.IntValue(model.Le)
	}
	if model.Prefix != nil {
		modelMap["prefix"] = *model.Prefix
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupStatusReason) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap
}

// ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch converts the patch struct to a map containing only the changed fields.
func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPeerGroupPatch, d *schema.ResourceData) map[string]interface{} {
	patch := map[string]interface{}{}
	if patchVals == nil {
		return patch
	}
	if patchVals.HealthMonitor != nil {
		healthMonitor := map[string]interface{}{}
		if patchVals.HealthMonitor.Mode != nil {
			healthMonitor["mode"] = *patchVals.HealthMonitor.Mode
		}
		patch["health_monitor"] = healthMonitor
	}
	if patchVals.Name != nil {
		patch["name"] = *patchVals.Name
	}
	if patchVals.PrefixWatcher != nil {
		excluded := []map[string]interface{}{}
		for i := range patchVals.PrefixWatcher.ExcludedPrefixes {
			excluded = append(excluded, resourceIBMIsDynamicRouteServerPeerGroupPrefixPatchEntry(patchVals.PrefixWatcher.ExcludedPrefixes[i].Prefix, patchVals.PrefixWatcher.ExcludedPrefixes[i].Ge, patchVals.PrefixWatcher.ExcludedPrefixes[i].Le))
		}
		monitored := []map[string]interface{}{}
		for i := range patchVals.PrefixWatcher.MonitoredPrefixes {
			monitored = append(monitored, resourceIBMIsDynamicRouteServerPeerGroupPrefixPatchEntry(patchVals.PrefixWatcher.MonitoredPrefixes[i].Prefix, patchVals.PrefixWatcher.MonitoredPrefixes[i].Ge, patchVals.PrefixWatcher.MonitoredPrefixes[i].Le))
		}
		patch["prefix_watcher"] = map[string]interface{}{
			"excluded_prefixes":  excluded,
			"monitored_prefixes": monitored,
		}
	}
	if patchVals.State != nil {
		patch["state"] = *patchVals.State
	}
	return patch
}

func resourceIBMIsDynamicRouteServerPeerGroupPrefixPatchEntry(prefix *string, ge, le *int64) map[string]interface{} {
	entry := map[string]interface{}{}
	if prefix != nil {
		entry["prefix"] = *prefix
	}
	if ge != nil {
		entry["ge"] = *ge
	}
	if le != nil {
		entry["le"] = *le
	}
	return entry
}
