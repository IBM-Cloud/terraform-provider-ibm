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
	isDynamicRouteServerPeerGroupPolicyStable   = "stable"
	isDynamicRouteServerPeerGroupPolicyFailed   = "failed"
	isDynamicRouteServerPeerGroupPolicyDeleted  = "deleted"
	isDynamicRouteServerPeerGroupPolicyPending  = "pending"
	isDynamicRouteServerPeerGroupPolicyDeleting = "deleting"
	isDynamicRouteServerPeerGroupPolicyUpdating = "updating"
	isDynamicRouteServerPeerGroupPolicyWaiting  = "waiting"
)

func ResourceIBMIsDynamicRouteServerPeerGroupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupPolicyRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyDelete,
		Importer:      &schema.ResourceImporter{},
		CustomizeDiff: resourceIBMIsDynamicRouteServerPeerGroupPolicyCustomizeDiff,

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
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "dynamic_route_server_peer_group_id"),
				Description:  "The dynamic route server peer group identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "name"),
				Description:  "The name for this dynamic route server peer group policy. The name must not be used by another peer group policy. Names starting with ibm- are reserved. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enabled",
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "state"),
				Description:  "The state of this dynamic route server peer group policy. One of `enabled` or `disabled`.",
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "type"),
				Description:  "The type of dynamic route server peer group policy. One of `custom_routes`, `learned_routes`, `vpc_address_prefixes` or `vpc_routing_tables`.",
			},
			"custom_routes": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The custom routes to advertise. Only allowed when `type` is `custom_routes`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The destination CIDR of the route. The host identifier in the CIDR must be zero.",
						},
					},
				},
			},
			"excluded_prefixes": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The prefixes excluded from monitoring when learning routes from connected peer groups. Only allowed when `type` is `learned_routes` or `vpc_routing_tables`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ge": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The minimum prefix length to match. Zero disables `ge` filtering.",
						},
						"le": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The maximum prefix length to match. Zero disables `le` filtering.",
						},
						"prefix": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The prefix to be excluded.",
						},
					},
				},
			},
			"peer_groups": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The peer groups from which routes are learned and then advertised back to the peer group this policy is applied to. Only allowed when `type` is `learned_routes`. Each item is identified by `id` or `href`.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupIdentitySchema(),
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The next hop peer groups. Only allowed when `type` is `vpc_routing_tables`. Each item is identified by `id` or `href`.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupIdentitySchema(),
			},
			"route_delete_delay": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "route_delete_delay"),
				Description:  "The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`, or after it is no longer advertised. Only allowed when `type` is `vpc_routing_tables`.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The routing tables to update with the learned routes. Only allowed when `type` is `vpc_routing_tables`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
						},
						"vpc_routing_table": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The VPC routing table, identified by `id`, `href` or `crn`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this routing table.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The URL for this routing table.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The CRN for this routing table.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this routing table.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A link to documentation about deleted resources.",
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
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the peer group policy was created or last updated.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group policy.",
			},
			"lifecycle_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `lifecycle_state` (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reason code for this lifecycle state.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this lifecycle state.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the dynamic route server peer group policy.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"dynamic_route_server_peer_group_policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server peer group policy.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupIdentitySchema is the element schema shared by
// peer_groups and next_hops: the user supplies id or href, the service returns the full reference.
func resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupIdentitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server peer group.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this dynamic route server peer group.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"deleted": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to documentation about deleted resources.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidator() *validate.ResourceValidator {
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
			Identifier:                 "dynamic_route_server_peer_group_id",
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
			Identifier:                 "state",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "disabled, enabled",
		},
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "custom_routes, learned_routes, vpc_address_prefixes, vpc_routing_tables",
		},
		validate.ValidateSchema{
			Identifier:                 "route_delete_delay",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "60",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group_policy", Schema: validateSchema}
	return &resourceValidator
}

// policyFieldGetter is satisfied by both *schema.ResourceData and *schema.ResourceDiff.
type policyFieldGetter interface {
	Get(key string) interface{}
}

// ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields checks that only the fields allowed by the
// policy `type` are set, mirroring the oneOf prototypes of the API.
func ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(policyType string, d policyFieldGetter) error {
	listLen := func(key string) int {
		if v, ok := d.Get(key).([]interface{}); ok {
			return len(v)
		}
		return 0
	}
	allowed := map[string]map[string]bool{
		"custom_routes":        {"custom_routes": true},
		"learned_routes":       {"excluded_prefixes": true, "peer_groups": true},
		"vpc_address_prefixes": {},
		"vpc_routing_tables":   {"excluded_prefixes": true, "next_hops": true, "route_delete_delay": true, "routing_tables": true},
	}
	allowedForType, ok := allowed[policyType]
	if !ok {
		return fmt.Errorf("unsupported dynamic route server peer group policy type %q", policyType)
	}
	for _, key := range []string{"custom_routes", "excluded_prefixes", "peer_groups", "next_hops", "routing_tables"} {
		if listLen(key) > 0 && !allowedForType[key] {
			return fmt.Errorf("%q cannot be set when type is %q", key, policyType)
		}
	}
	if delay, ok := d.Get("route_delete_delay").(int); ok && delay != 0 && !allowedForType["route_delete_delay"] {
		return fmt.Errorf("%q cannot be set when type is %q", "route_delete_delay", policyType)
	}
	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	policyType, _ := d.Get("type").(string)
	if policyType == "" {
		return nil
	}
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(policyType, d)
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	policyType := d.Get("type").(string)
	if err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(policyType, d); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "validate-type").GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	bodyModelMap["type"] = policyType
	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if _, ok := d.GetOk("state"); ok {
		bodyModelMap["state"] = d.Get("state")
	}
	for _, key := range []string{"custom_routes", "excluded_prefixes", "peer_groups", "next_hops", "routing_tables"} {
		if _, ok := d.GetOk(key); ok {
			bodyModelMap[key] = d.Get(key)
		}
	}
	if _, ok := d.GetOk("route_delete_delay"); ok {
		bodyModelMap["route_delete_delay"] = d.Get("route_delete_delay")
	}

	convertedModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "parse-request-body").GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	peerGroupID := d.Get("dynamic_route_server_peer_group_id").(string)
	createOptions := &vpcv1.CreateDynamicRouteServerPeerGroupPolicyOptions{}
	createOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	createOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	createOptions.DynamicRouteServerPeerGroupPolicyPrototype = convertedModel

	policyIntf, _, err := vpcClient.CreateDynamicRouteServerPeerGroupPolicyWithContext(context, createOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	policy, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(policyIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "parse-response").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", dynamicRouteServerID, peerGroupID, *policy.ID))

	_, err = isWaitForDynamicRouteServerPeerGroupPolicyStable(vpcClient, dynamicRouteServerID, peerGroupID, *policy.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server peer group policy to become stable failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context, d, meta)
}

// resourceIBMIsDynamicRouteServerPeerGroupPolicyParseID splits "<drs>/<peer_group>/<policy>".
func resourceIBMIsDynamicRouteServerPeerGroupPolicyParseID(id string) (string, string, string, error) {
	parts, err := flex.SepIdParts(id, "/")
	if err != nil {
		return "", "", "", err
	}
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("invalid ID %q, expected <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<policy_id>", id)
	}
	return parts[0], parts[1], parts[2], nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, policyID, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "sep-id-parts").GetDiag()
	}

	getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	getOptions.SetID(policyID)

	policyIntf, response, err := vpcClient.GetDynamicRouteServerPeerGroupPolicyWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	policy, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(policyIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "parse-response").GetDiag()
	}

	// On import only the ID is known; derive the parent identifiers from it.
	if err = d.Set("dynamic_route_server_id", dynamicRouteServerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_id: %s", err), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-dynamic_route_server_id").GetDiag()
	}
	if err = d.Set("dynamic_route_server_peer_group_id", peerGroupID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_id: %s", err), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-dynamic_route_server_peer_group_id").GetDiag()
	}

	modelMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(policy)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "policy-to-map").GetDiag()
	}
	for key, value := range modelMap {
		if key == "id" {
			key = "dynamic_route_server_peer_group_policy_id"
		}
		if err = d.Set(key, value); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-"+key).GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, policyID, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "sep-id-parts").GetDiag()
	}

	policyType := d.Get("type").(string)
	if err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(policyType, d); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "validate-type").GetDiag()
	}

	patchVals := &vpcv1.DynamicRouteServerPeerGroupPolicyPatch{}
	patch := map[string]interface{}{}
	hasChange := false

	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("state") {
		patchVals.State = core.StringPtr(d.Get("state").(string))
		hasChange = true
	}
	if d.HasChange("route_delete_delay") {
		patchVals.RouteDeleteDelay = core.Int64Ptr(int64(d.Get("route_delete_delay").(int)))
		hasChange = true
	}
	if d.HasChange("custom_routes") {
		customRoutes := []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
		for _, v := range d.Get("custom_routes").([]interface{}) {
			item, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(v.(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-custom_routes").GetDiag()
			}
			customRoutes = append(customRoutes, *item)
		}
		if len(customRoutes) > 0 {
			patchVals.CustomRoutes = customRoutes
		}
		patch["custom_routes"] = []map[string]interface{}{}
		hasChange = true
	}
	if d.HasChange("excluded_prefixes") {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, v := range d.Get("excluded_prefixes").([]interface{}) {
			item, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(v.(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-excluded_prefixes").GetDiag()
			}
			excludedPrefixes = append(excludedPrefixes, *item)
		}
		if len(excludedPrefixes) > 0 {
			patchVals.ExcludedPrefixes = excludedPrefixes
		}
		patch["excluded_prefixes"] = []map[string]interface{}{}
		hasChange = true
	}
	if d.HasChange("next_hops") {
		nextHops := []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{}
		for _, v := range d.Get("next_hops").([]interface{}) {
			item, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(v.(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-next_hops").GetDiag()
			}
			nextHops = append(nextHops, item)
		}
		if len(nextHops) > 0 {
			patchVals.NextHops = nextHops
		}
		patch["next_hops"] = []map[string]interface{}{}
		hasChange = true
	}
	if d.HasChange("peer_groups") {
		peerGroups := []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{}
		for _, v := range d.Get("peer_groups").([]interface{}) {
			item, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(v.(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-peer_groups").GetDiag()
			}
			peerGroups = append(peerGroups, item)
		}
		if len(peerGroups) > 0 {
			patchVals.PeerGroups = peerGroups
		}
		patch["peer_groups"] = []map[string]interface{}{}
		hasChange = true
	}
	if d.HasChange("routing_tables") {
		routingTables := []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch{}
		for _, v := range d.Get("routing_tables").([]interface{}) {
			item, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(v.(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-routing_tables").GetDiag()
			}
			routingTables = append(routingTables, *item)
		}
		if len(routingTables) > 0 {
			patchVals.RoutingTables = routingTables
		}
		patch["routing_tables"] = []map[string]interface{}{}
		hasChange = true
	}

	if hasChange {
		// A list cleared by the user is sent as an empty array (pre-seeded above); AsPatch only
		// emits the populated lists, which overwrite the empty placeholders.
		asPatch, err := patchVals.AsPatch()
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "as-patch").GetDiag()
		}
		for key, value := range asPatch {
			patch[key] = value
		}

		updateOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupPolicyOptions{}
		updateOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		updateOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		updateOptions.SetID(policyID)
		updateOptions.SetIfMatch(d.Get("etag").(string))
		updateOptions.DynamicRouteServerPeerGroupPolicyPatch = patch

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupPolicyWithContext(context, updateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForDynamicRouteServerPeerGroupPolicyStable(vpcClient, dynamicRouteServerID, peerGroupID, policyID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server peer group policy to become stable failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, policyID, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "delete", "sep-id-parts").GetDiag()
	}

	deleteOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupPolicyOptions{}
	deleteOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	deleteOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	deleteOptions.SetID(policyID)
	if etag, ok := d.GetOk("etag"); ok && etag.(string) != "" {
		deleteOptions.SetIfMatch(etag.(string))
	}

	_, response, err := vpcClient.DeleteDynamicRouteServerPeerGroupPolicyWithContext(context, deleteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerPeerGroupPolicyDeleted(vpcClient, dynamicRouteServerID, peerGroupID, policyID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server peer group policy to be deleted failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isDynamicRouteServerPeerGroupPolicyRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, policyID string, forDelete bool) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		getOptions.SetID(policyID)
		policyIntf, response, err := sess.GetDynamicRouteServerPeerGroupPolicy(getOptions)
		if err != nil {
			if forDelete && response != nil && response.StatusCode == 404 {
				return struct{}{}, isDynamicRouteServerPeerGroupPolicyDeleted, nil
			}
			return nil, "", fmt.Errorf("Error getting dynamic route server peer group policy: %s\n%s", err, response)
		}
		policy, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(policyIntf)
		if err != nil {
			return nil, "", err
		}
		state := ""
		if policy.LifecycleState != nil {
			state = *policy.LifecycleState
		}
		if state == isDynamicRouteServerPeerGroupPolicyFailed {
			return policy, state, fmt.Errorf("dynamic route server peer group policy (%s) entered the failed lifecycle state", policyID)
		}
		return policy, state, nil
	}
}

func isWaitForDynamicRouteServerPeerGroupPolicyStable(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, policyID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group policy (%s) to be stable.", policyID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupPolicyPending, isDynamicRouteServerPeerGroupPolicyUpdating, isDynamicRouteServerPeerGroupPolicyWaiting, ""},
		Target:     []string{isDynamicRouteServerPeerGroupPolicyStable},
		Refresh:    isDynamicRouteServerPeerGroupPolicyRefreshFunc(sess, dynamicRouteServerID, peerGroupID, policyID, false),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isWaitForDynamicRouteServerPeerGroupPolicyDeleted(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, policyID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group policy (%s) to be deleted.", policyID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupPolicyDeleting, isDynamicRouteServerPeerGroupPolicyStable, isDynamicRouteServerPeerGroupPolicyPending, isDynamicRouteServerPeerGroupPolicyUpdating, isDynamicRouteServerPeerGroupPolicyWaiting, ""},
		Target:     []string{isDynamicRouteServerPeerGroupPolicyDeleted},
		Refresh:    isDynamicRouteServerPeerGroupPolicyRefreshFunc(sess, dynamicRouteServerID, peerGroupID, policyID, true),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

// ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase normalizes any DynamicRouteServerPeerGroupPolicyIntf
// subtype (custom routes, learned routes, VPC address prefixes, VPC routing tables) into the base struct.
func ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model vpcv1.DynamicRouteServerPeerGroupPolicyIntf) (*vpcv1.DynamicRouteServerPeerGroupPolicy, error) {
	switch policy := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPolicy:
		if policy == nil {
			return nil, fmt.Errorf("dynamic route server peer group policy is nil")
		}
		return policy, nil
	case *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes:
		if policy == nil {
			return nil, fmt.Errorf("dynamic route server peer group policy is nil")
		}
		return &vpcv1.DynamicRouteServerPeerGroupPolicy{
			CreatedAt:        policy.CreatedAt,
			Href:             policy.Href,
			ID:               policy.ID,
			LifecycleReasons: policy.LifecycleReasons,
			LifecycleState:   policy.LifecycleState,
			Name:             policy.Name,
			ResourceType:     policy.ResourceType,
			State:            policy.State,
			Type:             policy.Type,
			CustomRoutes:     policy.CustomRoutes,
		}, nil
	case *vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes:
		if policy == nil {
			return nil, fmt.Errorf("dynamic route server peer group policy is nil")
		}
		return &vpcv1.DynamicRouteServerPeerGroupPolicy{
			CreatedAt:        policy.CreatedAt,
			Href:             policy.Href,
			ID:               policy.ID,
			LifecycleReasons: policy.LifecycleReasons,
			LifecycleState:   policy.LifecycleState,
			Name:             policy.Name,
			ResourceType:     policy.ResourceType,
			State:            policy.State,
			Type:             policy.Type,
			ExcludedPrefixes: policy.ExcludedPrefixes,
			PeerGroups:       policy.PeerGroups,
		}, nil
	case *vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes:
		if policy == nil {
			return nil, fmt.Errorf("dynamic route server peer group policy is nil")
		}
		return &vpcv1.DynamicRouteServerPeerGroupPolicy{
			CreatedAt:        policy.CreatedAt,
			Href:             policy.Href,
			ID:               policy.ID,
			LifecycleReasons: policy.LifecycleReasons,
			LifecycleState:   policy.LifecycleState,
			Name:             policy.Name,
			ResourceType:     policy.ResourceType,
			State:            policy.State,
			Type:             policy.Type,
		}, nil
	case *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables:
		if policy == nil {
			return nil, fmt.Errorf("dynamic route server peer group policy is nil")
		}
		return &vpcv1.DynamicRouteServerPeerGroupPolicy{
			CreatedAt:        policy.CreatedAt,
			Href:             policy.Href,
			ID:               policy.ID,
			LifecycleReasons: policy.LifecycleReasons,
			LifecycleState:   policy.LifecycleState,
			Name:             policy.Name,
			ResourceType:     policy.ResourceType,
			State:            policy.State,
			Type:             policy.Type,
			ExcludedPrefixes: policy.ExcludedPrefixes,
			NextHops:         policy.NextHops,
			RouteDeleteDelay: policy.RouteDeleteDelay,
			RoutingTables:    policy.RoutingTables,
		}, nil
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPolicyIntf subtype encountered")
	}
}

// ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap flattens a policy into the
// resource attribute map (keys match the schema; "id" is the policy identifier).
func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("dynamic route server peer group policy is nil")
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = flex.DateTimeToString(model.CreatedAt)
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, item := range model.LifecycleReasons {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyLifecycleReasonToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, itemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	if model.State != nil {
		modelMap["state"] = *model.State
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	customRoutes := []map[string]interface{}{}
	for _, item := range model.CustomRoutes {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		customRoutes = append(customRoutes, itemMap)
	}
	modelMap["custom_routes"] = customRoutes
	excludedPrefixes := []map[string]interface{}{}
	for _, item := range model.ExcludedPrefixes {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		excludedPrefixes = append(excludedPrefixes, itemMap)
	}
	modelMap["excluded_prefixes"] = excludedPrefixes
	peerGroups := []map[string]interface{}{}
	for _, item := range model.PeerGroups {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		peerGroups = append(peerGroups, itemMap)
	}
	modelMap["peer_groups"] = peerGroups
	nextHops := []map[string]interface{}{}
	for _, item := range model.NextHops {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(item)
		if err != nil {
			return modelMap, err
		}
		nextHops = append(nextHops, itemMap)
	}
	modelMap["next_hops"] = nextHops
	modelMap["route_delete_delay"] = 0
	if model.RouteDeleteDelay != nil {
		modelMap["route_delete_delay"] = flex.IntValue(model.RouteDeleteDelay)
	}
	routingTables := []map[string]interface{}{}
	for _, item := range model.RoutingTables {
		itemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		routingTables = append(routingTables, itemMap)
	}
	modelMap["routing_tables"] = routingTables
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyLifecycleReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("lifecycle reason is nil")
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
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
	destination, _ := modelMap["destination"].(string)
	if destination == "" {
		return model, fmt.Errorf("custom_routes item requires destination")
	}
	model.Destination = core.StringPtr(destination)
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
	if ge, ok := modelMap["ge"].(int); ok {
		model.Ge = core.Int64Ptr(int64(ge))
	}
	if le, ok := modelMap["le"].(int); ok {
		model.Le = core.Int64Ptr(int64(le))
	}
	prefix, _ := modelMap["prefix"].(string)
	if prefix == "" {
		return model, fmt.Errorf("excluded_prefixes item requires prefix")
	}
	model.Prefix = core.StringPtr(prefix)
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupIdentityIntf, error) {
	id, _ := modelMap["id"].(string)
	href, _ := modelMap["href"].(string)
	if id != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(modelMap)
	}
	return nil, fmt.Errorf("peer_groups item requires id or href")
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("peer group identity requires id")
	}
	return &vpcv1.DynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("peer group identity requires href")
	}
	return &vpcv1.DynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap)
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf, error) {
	id, _ := modelMap["id"].(string)
	href, _ := modelMap["href"].(string)
	if id != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap)
	}
	return nil, fmt.Errorf("next_hops item requires id or href")
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("next hop identity requires id")
	}
	return &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("next hop identity requires href")
	}
	return &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr(href)}, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyVPCRoutingTableIdentityFromMap(modelMap map[string]interface{}) (vpcv1.RoutingTableIdentityIntf, error) {
	list, _ := modelMap["vpc_routing_table"].([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, fmt.Errorf("routing_tables item requires vpc_routing_table")
	}
	identityMap, ok := list[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("routing_tables item requires vpc_routing_table")
	}
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(identityMap)
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{}
	advertise, _ := modelMap["advertise"].(bool)
	model.Advertise = core.BoolPtr(advertise)
	identity, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyVPCRoutingTableIdentityFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = identity
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch{}
	advertise, _ := modelMap["advertise"].(bool)
	model.Advertise = core.BoolPtr(advertise)
	identity, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyVPCRoutingTableIdentityFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = identity
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(modelMap map[string]interface{}) (vpcv1.RoutingTableIdentityIntf, error) {
	id, _ := modelMap["id"].(string)
	crn, _ := modelMap["crn"].(string)
	href, _ := modelMap["href"].(string)
	if id != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(modelMap)
	}
	if crn != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(modelMap)
	}
	return nil, fmt.Errorf("vpc_routing_table requires id, crn or href")
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("routing table identity requires id")
	}
	return &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByCRN, error) {
	crn, _ := modelMap["crn"].(string)
	if crn == "" {
		return nil, fmt.Errorf("routing table identity requires crn")
	}
	return &vpcv1.RoutingTableIdentityByCRN{CRN: core.StringPtr(crn)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("routing table identity requires href")
	}
	return &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr(href)}, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap map[string]interface{}, key string) []interface{} {
	list, _ := modelMap[key].([]interface{})
	return list
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyCustomRoutesFromMap(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype, error) {
	result := []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
	for _, item := range resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap, "custom_routes") {
		itemMap, _ := item.(map[string]interface{})
		model, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(itemMap)
		if err != nil {
			return nil, err
		}
		result = append(result, *model)
	}
	return result, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyExcludedPrefixesFromMap(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, error) {
	result := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
	for _, item := range resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap, "excluded_prefixes") {
		itemMap, _ := item.(map[string]interface{})
		model, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(itemMap)
		if err != nil {
			return nil, err
		}
		result = append(result, *model)
	}
	return result, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupsFromMap(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupIdentityIntf, error) {
	result := []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{}
	for _, item := range resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap, "peer_groups") {
		itemMap, _ := item.(map[string]interface{})
		model, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(itemMap)
		if err != nil {
			return nil, err
		}
		result = append(result, model)
	}
	return result, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyNextHopsFromMap(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf, error) {
	result := []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{}
	for _, item := range resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap, "next_hops") {
		itemMap, _ := item.(map[string]interface{})
		model, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(itemMap)
		if err != nil {
			return nil, err
		}
		result = append(result, model)
	}
	return result, nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTablesFromMap(modelMap map[string]interface{}) ([]vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype, error) {
	result := []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{}
	for _, item := range resourceIBMIsDynamicRouteServerPeerGroupPolicyListFromMap(modelMap, "routing_tables") {
		itemMap, _ := item.(map[string]interface{})
		model, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(itemMap)
		if err != nil {
			return nil, err
		}
		result = append(result, *model)
	}
	return result, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype selects the oneOf
// prototype subtype from `type` and validates that only the fields allowed by that subtype are present.
func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeIntf, error) {
	policyType, _ := modelMap["type"].(string)
	if policyType == "" {
		return nil, fmt.Errorf("type is required")
	}
	if err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidateTypeFields(policyType, policyMapGetter(modelMap)); err != nil {
		return nil, err
	}
	switch policyType {
	case "custom_routes":
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(modelMap)
	case "learned_routes":
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(modelMap)
	case "vpc_address_prefixes":
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(modelMap)
	case "vpc_routing_tables":
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(modelMap)
	}
	return nil, fmt.Errorf("unsupported dynamic route server peer group policy type %q", policyType)
}

// policyMapGetter adapts a plain map to the policyFieldGetter interface.
type policyMapGetter map[string]interface{}

func (m policyMapGetter) Get(key string) interface{} {
	return m[key]
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyNameState(modelMap map[string]interface{}) (*string, *string) {
	var name, state *string
	if v, _ := modelMap["name"].(string); v != "" {
		name = core.StringPtr(v)
	}
	if v, _ := modelMap["state"].(string); v != "" {
		state = core.StringPtr(v)
	}
	return name, state
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype{}
	model.Name, model.State = resourceIBMIsDynamicRouteServerPeerGroupPolicyNameState(modelMap)
	customRoutes, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyCustomRoutesFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.CustomRoutes = customRoutes
	model.Type = core.StringPtr("custom_routes")
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype{}
	model.Name, model.State = resourceIBMIsDynamicRouteServerPeerGroupPolicyNameState(modelMap)
	excludedPrefixes, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyExcludedPrefixesFromMap(modelMap)
	if err != nil {
		return model, err
	}
	if len(excludedPrefixes) > 0 {
		model.ExcludedPrefixes = excludedPrefixes
	}
	peerGroups, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyPeerGroupsFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.PeerGroups = peerGroups
	model.Type = core.StringPtr("learned_routes")
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype{}
	model.Name, model.State = resourceIBMIsDynamicRouteServerPeerGroupPolicyNameState(modelMap)
	model.Type = core.StringPtr("vpc_address_prefixes")
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype{}
	model.Name, model.State = resourceIBMIsDynamicRouteServerPeerGroupPolicyNameState(modelMap)
	excludedPrefixes, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyExcludedPrefixesFromMap(modelMap)
	if err != nil {
		return model, err
	}
	if len(excludedPrefixes) > 0 {
		model.ExcludedPrefixes = excludedPrefixes
	}
	nextHops, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyNextHopsFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.NextHops = nextHops
	if delay, ok := modelMap["route_delete_delay"].(int); ok {
		model.RouteDeleteDelay = core.Int64Ptr(int64(delay))
	}
	routingTables, err := resourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTablesFromMap(modelMap)
	if err != nil {
		return model, err
	}
	model.RoutingTables = routingTables
	model.Type = core.StringPtr("vpc_routing_tables")
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("custom route is nil")
	}
	if model.Destination != nil {
		modelMap["destination"] = *model.Destination
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("excluded prefix is nil")
	}
	modelMap["ge"] = 0
	if model.Ge != nil {
		modelMap["ge"] = flex.IntValue(model.Ge)
	}
	modelMap["le"] = 0
	if model.Le != nil {
		modelMap["le"] = flex.IntValue(model.Le)
	}
	if model.Prefix != nil {
		modelMap["prefix"] = *model.Prefix
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("peer group reference is nil")
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("deleted is nil")
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	switch nextHop := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference:
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(nextHop)
	case *vpcv1.DynamicRouteServerPeerGroupPolicyNextHop:
		if nextHop == nil {
			return nil, fmt.Errorf("next hop is nil")
		}
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference{
			Deleted:      nextHop.Deleted,
			Href:         nextHop.Href,
			ID:           nextHop.ID,
			Name:         nextHop.Name,
			ResourceType: nextHop.ResourceType,
		})
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf subtype encountered")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("next hop is nil")
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("routing table is nil")
	}
	modelMap["advertise"] = false
	if model.Advertise != nil {
		modelMap["advertise"] = *model.Advertise
	}
	if model.VPCRoutingTable != nil {
		vpcRoutingTableMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model.VPCRoutingTable)
		if err != nil {
			return modelMap, err
		}
		modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("routing table reference is nil")
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}
