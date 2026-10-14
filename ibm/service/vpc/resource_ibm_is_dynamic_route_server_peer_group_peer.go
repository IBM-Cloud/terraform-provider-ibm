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
	"github.com/go-openapi/strfmt"
)

const (
	isDynamicRouteServerPeerGroupPeerPending   = "pending"
	isDynamicRouteServerPeerGroupPeerUpdating  = "updating"
	isDynamicRouteServerPeerGroupPeerWaiting   = "waiting"
	isDynamicRouteServerPeerGroupPeerDeleting  = "deleting"
	isDynamicRouteServerPeerGroupPeerStable    = "stable"
	isDynamicRouteServerPeerGroupPeerFailed    = "failed"
	isDynamicRouteServerPeerGroupPeerSuspended = "suspended"
	isDynamicRouteServerPeerGroupPeerDeleted   = "deleted"
)

func ResourceIBMIsDynamicRouteServerPeerGroupPeer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupPeerCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupPeerRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupPeerUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupPeerDelete,
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
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "dynamic_route_server_peer_group_id"),
				Description:  "The dynamic route server peer group identifier.",
			},
			"asn": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The autonomous system number (ASN) for this dynamic route server peer group peer.",
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "The endpoint for this dynamic route server peer group peer. Specify either the reserved IP (`id` or `href`) or an IP address with an optional gateway.",
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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "name"),
				Description:  "The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.",
			},
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "priority"),
				Description:  "The priority of this dynamic route server peer group peer. Lower values have higher priority.",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "state"),
				Description:  "The state of this dynamic route server peer group peer: `enabled` or `disabled`.",
			},
			"bidirectional_forwarding_detection": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The bidirectional forwarding detection (BFD) configuration for this dynamic route server peer group peer.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerBFDSchema(),
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this dynamic route server peer group peer was created.",
			},
			"creator": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of resource that created this peer: `dynamic_route_server` or `transit_gateway`.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group peer.",
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
				Description: "The lifecycle state of the dynamic route server peer group peer.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"sessions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The BGP sessions for this dynamic route server peer group peer.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerSessionSchema(),
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this dynamic route server peer group peer connection: `down` or `up`.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status reasons for the dynamic route server peer group peer.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
			},
			"dynamic_route_server_peer_group_peer_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server peer group peer.",
			},
			"etag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ETag identifier for the dynamic route server peer group peer.",
			},
		},
	}
}

// resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema is the schema shared by lifecycle and status reasons.
func resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A reason code for this state.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An explanation of the reason for this state.",
			},
			"more_info": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A link to documentation about the reason for this state.",
			},
		},
	}
}

// resourceIBMIsDynamicRouteServerPeerGroupPeerReferenceSchema is the schema for a resource reference (member, peer, VNI, subnet, reserved IP).
func resourceIBMIsDynamicRouteServerPeerGroupPeerReferenceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this resource, if any.",
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
							Description: "Link to documentation about deleted resources.",
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this resource.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this resource.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this resource.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}

// resourceIBMIsDynamicRouteServerPeerGroupPeerMemberSchema is the schema for a dynamic route server member reference.
func resourceIBMIsDynamicRouteServerPeerGroupPeerMemberSchema() *schema.Resource {
	s := resourceIBMIsDynamicRouteServerPeerGroupPeerReferenceSchema()
	delete(s.Schema, "crn")
	s.Schema["virtual_network_interfaces"] = &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "The virtual network interfaces for this dynamic route server member.",
		Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReferenceSchema(),
	}
	return s
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerBFDSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"detect_multiplier": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The detection multiplier used for BFD.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether BFD is enabled for this peer.",
			},
			"mode": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The BFD mode.",
			},
			"receive_interval": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum receive interval in milliseconds.",
			},
			"role": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The BFD role of this peer.",
			},
			"sessions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The BFD sessions for this peer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The dynamic route server member for the local side of this session.",
							Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerMemberSchema(),
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The peer for the remote side of this session.",
							Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReferenceSchema(),
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of this BFD session.",
						},
					},
				},
			},
			"transmit_interval": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum transmit interval in milliseconds.",
			},
		},
	}
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerSessionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"established_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this BGP session was established.",
			},
			"local": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The dynamic route server member for the local side of this session.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerMemberSchema(),
			},
			"protocol_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The BGP protocol state of this session.",
			},
			"remote": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The remote side of this session.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the remote peer.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerValidator() *validate.ResourceValidator {
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
			Identifier:                 "priority",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "4",
		},
		validate.ValidateSchema{
			Identifier:                 "state",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "disabled, enabled",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group_peer", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{
		"asn":      d.Get("asn"),
		"endpoint": d.Get("endpoint"),
	}
	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if resourceIBMIsDynamicRouteServerPeerGroupPeerConfigHas(d, "priority") {
		bodyModelMap["priority"] = d.Get("priority")
	}
	if _, ok := d.GetOk("state"); ok {
		bodyModelMap["state"] = d.Get("state")
	}
	convertedModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "parse-request-body").GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	peerGroupID := d.Get("dynamic_route_server_peer_group_id").(string)
	createOptions := &vpcv1.CreateDynamicRouteServerPeerGroupPeerOptions{}
	createOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	createOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	createOptions.DynamicRouteServerPeerGroupPeerPrototype = convertedModel

	peerIntf, _, err := vpcClient.CreateDynamicRouteServerPeerGroupPeerWithContext(context, createOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	peerID, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(peerIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "peer-id").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", dynamicRouteServerID, peerGroupID, peerID))

	_, err = isWaitForDynamicRouteServerPeerGroupPeerStable(vpcClient, dynamicRouteServerID, peerGroupID, peerID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "wait-stable").GetDiag()
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context, d, meta)
}

// resourceIBMIsDynamicRouteServerPeerGroupPeerConfigHas reports whether the attribute is set in the
// configuration, so that a zero value such as priority 0 is sent to the API.
func resourceIBMIsDynamicRouteServerPeerGroupPeerConfigHas(d *schema.ResourceData, key string) bool {
	raw := d.GetRawConfig()
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute(key) {
		return false
	}
	value := raw.GetAttr(key)
	return value.IsKnown() && !value.IsNull()
}

// resourceIBMIsDynamicRouteServerPeerGroupPeerParseID splits the resource ID "<dynamic_route_server_id>/<peer_group_id>/<peer_id>".
func resourceIBMIsDynamicRouteServerPeerGroupPeerParseID(id string) (string, string, string, error) {
	parts, err := flex.SepIdParts(id, "/")
	if err != nil {
		return "", "", "", err
	}
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("invalid ID %q: expected <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<peer_id>", id)
	}
	return parts[0], parts[1], parts[2], nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, peerID, err := resourceIBMIsDynamicRouteServerPeerGroupPeerParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "sep-id-parts").GetDiag()
	}

	getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	getOptions.SetID(peerID)

	peerIntf, response, err := vpcClient.GetDynamicRouteServerPeerGroupPeerWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	peerMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(peerIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "peer-to-map").GetDiag()
	}

	// The parent identifiers are derived from the ID so that import works.
	if err = d.Set("dynamic_route_server_id", dynamicRouteServerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_id: %s", err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-dynamic_route_server_id").GetDiag()
	}
	if err = d.Set("dynamic_route_server_peer_group_id", peerGroupID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_id: %s", err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-dynamic_route_server_peer_group_id").GetDiag()
	}
	if err = d.Set("dynamic_route_server_peer_group_peer_id", peerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_peer_id: %s", err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-dynamic_route_server_peer_group_peer_id").GetDiag()
	}
	for _, key := range []string{"asn", "endpoint", "name", "priority", "state", "bidirectional_forwarding_detection", "created_at", "creator", "href", "lifecycle_reasons", "lifecycle_state", "resource_type", "sessions", "status", "status_reasons"} {
		if value, ok := peerMap[key]; ok {
			if err = d.Set(key, value); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-"+key).GetDiag()
			}
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, peerID, err := resourceIBMIsDynamicRouteServerPeerGroupPeerParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "update", "sep-id-parts").GetDiag()
	}

	hasChange := false
	patchVals := &vpcv1.DynamicRouteServerPeerGroupPeerPatch{}
	if d.HasChange("asn") {
		patchVals.Asn = core.Int64Ptr(int64(d.Get("asn").(int)))
		hasChange = true
	}
	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("priority") {
		patchVals.Priority = core.Int64Ptr(int64(d.Get("priority").(int)))
		hasChange = true
	}
	if d.HasChange("state") {
		patchVals.State = core.StringPtr(d.Get("state").(string))
		hasChange = true
	}

	if hasChange {
		updateOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupPeerOptions{}
		updateOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		updateOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		updateOptions.SetID(peerID)
		updateOptions.SetIfMatch(d.Get("etag").(string))
		updateOptions.DynamicRouteServerPeerGroupPeerPatch = ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupPeerWithContext(context, updateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForDynamicRouteServerPeerGroupPeerStable(vpcClient, dynamicRouteServerID, peerGroupID, peerID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "update", "wait-stable").GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, peerGroupID, peerID, err := resourceIBMIsDynamicRouteServerPeerGroupPeerParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "delete", "sep-id-parts").GetDiag()
	}

	deleteOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupPeerOptions{}
	deleteOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	deleteOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	deleteOptions.SetID(peerID)
	if etag, ok := d.GetOk("etag"); ok {
		deleteOptions.SetIfMatch(etag.(string))
	}

	_, response, err := vpcClient.DeleteDynamicRouteServerPeerGroupPeerWithContext(context, deleteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerPeerGroupPeerDeleted(vpcClient, dynamicRouteServerID, peerGroupID, peerID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "delete", "wait-deleted").GetDiag()
	}

	d.SetId("")

	return nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID returns the ID carried by any peer subtype.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfID(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (string, error) {
	var id *string
	switch peer := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddress:
		id = peer.ID
	case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway:
		id = peer.ID
	case *vpcv1.DynamicRouteServerPeerGroupPeer:
		id = peer.ID
	default:
		return "", fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerIntf subtype encountered")
	}
	if id == nil {
		return "", fmt.Errorf("dynamic route server peer group peer response has no id")
	}
	return *id, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState returns the lifecycle state carried by any peer subtype.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) string {
	var state *string
	switch peer := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddress:
		state = peer.LifecycleState
	case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway:
		state = peer.LifecycleState
	case *vpcv1.DynamicRouteServerPeerGroupPeer:
		state = peer.LifecycleState
	}
	if state == nil {
		return ""
	}
	return *state
}

func isWaitForDynamicRouteServerPeerGroupPeerStable(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, peerID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group peer (%s) to be stable.", peerID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupPeerPending, isDynamicRouteServerPeerGroupPeerUpdating, isDynamicRouteServerPeerGroupPeerWaiting},
		Target:     []string{isDynamicRouteServerPeerGroupPeerStable, isDynamicRouteServerPeerGroupPeerFailed, isDynamicRouteServerPeerGroupPeerSuspended},
		Refresh:    isDynamicRouteServerPeerGroupPeerRefreshFunc(sess, dynamicRouteServerID, peerGroupID, peerID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerPeerGroupPeerRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, peerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		getOptions.SetID(peerID)
		peer, response, err := sess.GetDynamicRouteServerPeerGroupPeer(getOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting dynamic route server peer group peer: %s\n%s", err, response)
		}
		state := ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(peer)
		if state == isDynamicRouteServerPeerGroupPeerFailed || state == isDynamicRouteServerPeerGroupPeerSuspended {
			return peer, state, fmt.Errorf("[ERROR] The dynamic route server peer group peer (%s) is in %s state", peerID, state)
		}
		if state == "" {
			// A peer without a lifecycle state is not transitioning.
			return peer, isDynamicRouteServerPeerGroupPeerStable, nil
		}
		return peer, state, nil
	}
}

func isWaitForDynamicRouteServerPeerGroupPeerDeleted(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, peerID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server peer group peer (%s) to be deleted.", peerID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPeerGroupPeerDeleting, isDynamicRouteServerPeerGroupPeerStable, isDynamicRouteServerPeerGroupPeerPending, isDynamicRouteServerPeerGroupPeerUpdating, isDynamicRouteServerPeerGroupPeerWaiting},
		Target:     []string{isDynamicRouteServerPeerGroupPeerDeleted, isDynamicRouteServerPeerGroupPeerFailed},
		Refresh:    isDynamicRouteServerPeerGroupPeerDeleteRefreshFunc(sess, dynamicRouteServerID, peerGroupID, peerID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerPeerGroupPeerDeleteRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID, peerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		getOptions.SetID(peerID)
		peer, response, err := sess.GetDynamicRouteServerPeerGroupPeer(getOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return peer, isDynamicRouteServerPeerGroupPeerDeleted, nil
			}
			return nil, "", fmt.Errorf("[ERROR] Error getting dynamic route server peer group peer: %s\n%s", err, response)
		}
		state := ResourceIBMIsDynamicRouteServerPeerGroupPeerIntfLifecycleState(peer)
		if state == isDynamicRouteServerPeerGroupPeerFailed {
			return peer, state, fmt.Errorf("[ERROR] The dynamic route server peer group peer (%s) failed to delete", peerID)
		}
		if state == "" {
			return peer, isDynamicRouteServerPeerGroupPeerDeleting, nil
		}
		return peer, state, nil
	}
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype builds the oneOf peer prototype.
// Only the address peer prototype can be created through the API.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(modelMap)
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype{}
	asn, ok := modelMap["asn"].(int)
	if !ok {
		return nil, fmt.Errorf("asn is required for a dynamic route server peer group peer")
	}
	model.Asn = core.Int64Ptr(int64(asn))
	endpointList, _ := modelMap["endpoint"].([]interface{})
	if len(endpointList) == 0 || endpointList[0] == nil {
		return nil, fmt.Errorf("endpoint is required for a dynamic route server peer group peer")
	}
	endpointMap, ok := endpointList[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("endpoint is required for a dynamic route server peer group peer")
	}
	endpointModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(endpointMap)
	if err != nil {
		return nil, err
	}
	model.Endpoint = endpointModel
	if name, ok := modelMap["name"].(string); ok && name != "" {
		model.Name = core.StringPtr(name)
	}
	if priority, ok := modelMap["priority"].(int); ok {
		model.Priority = core.Int64Ptr(int64(priority))
	}
	if state, ok := modelMap["state"].(string); ok && state != "" {
		model.State = core.StringPtr(state)
	}
	return model, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype builds the oneOf endpoint prototype:
// a reserved IP identity (by id or by href) or an IP address with an optional gateway.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeIntf, error) {
	id, _ := modelMap["id"].(string)
	href, _ := modelMap["href"].(string)
	address, _ := modelMap["address"].(string)
	gatewayList, _ := modelMap["gateway"].([]interface{})
	if (id != "" || href != "") && (address != "" || len(gatewayList) > 0) {
		return nil, fmt.Errorf("endpoint must specify either a reserved IP (id or href) or an address with an optional gateway, not both")
	}
	if id != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(modelMap)
	}
	if address != "" {
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(modelMap)
	}
	return nil, fmt.Errorf("endpoint must specify a reserved IP (id or href) or an address")
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("id is required for a reserved IP identity by id")
	}
	return &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("href is required for a reserved IP identity by href")
	}
	return &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype, error) {
	address, _ := modelMap["address"].(string)
	if address == "" {
		return nil, fmt.Errorf("address is required for an IP gateway endpoint")
	}
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype{Address: core.StringPtr(address)}
	if gatewayList, ok := modelMap["gateway"].([]interface{}); ok && len(gatewayList) > 0 && gatewayList[0] != nil {
		gatewayModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(gatewayList[0].(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		model.Gateway = gatewayModel
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway, error) {
	address, _ := modelMap["address"].(string)
	if address == "" {
		return nil, fmt.Errorf("address is required for an endpoint gateway")
	}
	return &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{Address: core.StringPtr(address)}, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap flattens any peer subtype (address, transit gateway or base) into the full response shape.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (map[string]interface{}, error) {
	switch peer := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddress:
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressToMap(peer)
	case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway:
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayToMap(peer)
	case *vpcv1.DynamicRouteServerPeerGroupPeer:
		if peer == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerIntf subtype encountered")
		}
		modelMap := resourceIBMIsDynamicRouteServerPeerGroupPeerCommonToMap(peer.CreatedAt, peer.Creator, peer.Href, peer.ID, peer.LifecycleReasons, peer.LifecycleState, peer.Name, peer.ResourceType, peer.Status, peer.StatusReasons, peer.State)
		if peer.Asn != nil {
			modelMap["asn"] = flex.IntValue(peer.Asn)
		}
		if peer.Priority != nil {
			modelMap["priority"] = flex.IntValue(peer.Priority)
		}
		if peer.BidirectionalForwardingDetection != nil {
			bfdMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(peer.BidirectionalForwardingDetection)
			if err != nil {
				return modelMap, err
			}
			modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bfdMap}
		}
		if peer.Endpoint != nil {
			endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(peer.Endpoint)
			if err != nil {
				return modelMap, err
			}
			modelMap["endpoint"] = []map[string]interface{}{endpointMap}
		}
		sessions := []map[string]interface{}{}
		for i := range peer.Sessions {
			sessionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(&peer.Sessions[i])
			if err != nil {
				return modelMap, err
			}
			sessions = append(sessions, sessionMap)
		}
		modelMap["sessions"] = sessions
		return modelMap, nil
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerIntf subtype encountered")
	}
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerCommonToMap(createdAt *strfmt.DateTime, creator, href, id *string, lifecycleReasons []vpcv1.LifecycleReason, lifecycleState, name, resourceType, status *string, statusReasons []vpcv1.DynamicRouteServerPeerGroupPeerStatusReason, state *string) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if createdAt != nil {
		modelMap["created_at"] = flex.DateTimeToString(createdAt)
	}
	if creator != nil {
		modelMap["creator"] = *creator
	}
	if href != nil {
		modelMap["href"] = *href
	}
	if id != nil {
		modelMap["id"] = *id
	}
	reasons := []map[string]interface{}{}
	for i := range lifecycleReasons {
		reasons = append(reasons, ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(&lifecycleReasons[i]))
	}
	modelMap["lifecycle_reasons"] = reasons
	if lifecycleState != nil {
		modelMap["lifecycle_state"] = *lifecycleState
	}
	if name != nil {
		modelMap["name"] = *name
	}
	if resourceType != nil {
		modelMap["resource_type"] = *resourceType
	}
	if status != nil {
		modelMap["status"] = *status
	}
	statusReasonMaps := []map[string]interface{}{}
	for i := range statusReasons {
		statusReasonMaps = append(statusReasonMaps, ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasons[i]))
	}
	modelMap["status_reasons"] = statusReasonMaps
	if state != nil {
		modelMap["state"] = *state
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddress) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group address peer is nil")
	}
	modelMap := resourceIBMIsDynamicRouteServerPeerGroupPeerCommonToMap(model.CreatedAt, model.Creator, model.Href, model.ID, model.LifecycleReasons, model.LifecycleState, model.Name, model.ResourceType, model.Status, model.StatusReasons, model.State)
	if model.Asn != nil {
		modelMap["asn"] = flex.IntValue(model.Asn)
	}
	if model.Priority != nil {
		modelMap["priority"] = flex.IntValue(model.Priority)
	}
	if model.BidirectionalForwardingDetection != nil {
		bfdMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model.BidirectionalForwardingDetection)
		if err != nil {
			return modelMap, err
		}
		modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bfdMap}
	}
	if model.Endpoint != nil {
		endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model.Endpoint)
		if err != nil {
			return modelMap, err
		}
		modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	}
	sessions := []map[string]interface{}{}
	for i := range model.Sessions {
		sessionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(&model.Sessions[i])
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionMap)
	}
	modelMap["sessions"] = sessions
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group transit gateway peer is nil")
	}
	modelMap := resourceIBMIsDynamicRouteServerPeerGroupPeerCommonToMap(model.CreatedAt, model.Creator, model.Href, model.ID, model.LifecycleReasons, model.LifecycleState, model.Name, model.ResourceType, model.Status, model.StatusReasons, model.State)
	if model.Endpoint != nil {
		endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model.Endpoint)
		if err != nil {
			return modelMap, err
		}
		modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	}
	sessions := []map[string]interface{}{}
	for i := range model.Sessions {
		sessionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(&model.Sessions[i])
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionMap)
	}
	modelMap["sessions"] = sessions
	return modelMap, nil
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap flattens the address endpoint oneOf (reserved IP, IP gateway or base).
func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	switch endpoint := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway:
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(endpoint)
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP:
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(endpoint)
	case *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint:
		if endpoint == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf subtype encountered")
		}
		if endpoint.Address != nil {
			modelMap["address"] = *endpoint.Address
		}
		if endpoint.Gateway != nil {
			modelMap["gateway"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(endpoint.Gateway)}
		}
		if endpoint.Deleted != nil {
			modelMap["deleted"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(endpoint.Deleted)}
		}
		if endpoint.Href != nil {
			modelMap["href"] = *endpoint.Href
		}
		if endpoint.ID != nil {
			modelMap["id"] = *endpoint.ID
		}
		if endpoint.Name != nil {
			modelMap["name"] = *endpoint.Name
		}
		if endpoint.ResourceType != nil {
			modelMap["resource_type"] = *endpoint.ResourceType
		}
		return modelMap, nil
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf subtype encountered")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("IP gateway endpoint is nil")
	}
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Gateway != nil {
		modelMap["gateway"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("reserved IP endpoint is nil")
	}
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Deleted != nil {
		modelMap["deleted"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)}
	}
	if model.Gateway != nil {
		modelMap["gateway"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model.Gateway)}
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model != nil && model.Address != nil {
		modelMap["address"] = *model.Address
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model *vpcv1.IP) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model != nil && model.Address != nil {
		modelMap["address"] = *model.Address
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model *vpcv1.Deleted) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model != nil && model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap flattens the transit gateway endpoint oneOf.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	var crn, id, resourceType *string
	switch endpoint := model.(type) {
	case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference:
		if endpoint == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf subtype encountered")
		}
		crn, id, resourceType = endpoint.CRN, endpoint.ID, endpoint.ResourceType
	case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint:
		if endpoint == nil {
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf subtype encountered")
		}
		crn, id, resourceType = endpoint.CRN, endpoint.ID, endpoint.ResourceType
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf subtype encountered")
	}
	if crn != nil {
		modelMap["crn"] = *crn
	}
	if id != nil {
		modelMap["id"] = *id
	}
	if resourceType != nil {
		modelMap["resource_type"] = *resourceType
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(model *vpcv1.LifecycleReason) map[string]interface{} {
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerStatusReason) map[string]interface{} {
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("bidirectional forwarding detection is nil")
	}
	modelMap := make(map[string]interface{})
	if model.DetectMultiplier != nil {
		modelMap["detect_multiplier"] = flex.IntValue(model.DetectMultiplier)
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Mode != nil {
		modelMap["mode"] = *model.Mode
	}
	if model.ReceiveInterval != nil {
		modelMap["receive_interval"] = flex.IntValue(model.ReceiveInterval)
	}
	if model.Role != nil {
		modelMap["role"] = *model.Role
	}
	sessions := []map[string]interface{}{}
	for i := range model.Sessions {
		sessionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&model.Sessions[i])
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionMap)
	}
	modelMap["sessions"] = sessions
	if model.TransmitInterval != nil {
		modelMap["transmit_interval"] = flex.IntValue(model.TransmitInterval)
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("bidirectional forwarding detection session is nil")
	}
	modelMap := make(map[string]interface{})
	if model.Local != nil {
		modelMap["local"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model.Local)}
	}
	if model.Remote != nil {
		modelMap["remote"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model.Remote)}
	}
	if model.State != nil {
		modelMap["state"] = *model.State
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model *vpcv1.DynamicRouteServerMemberReference) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.Deleted != nil {
		modelMap["deleted"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)}
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
	vnis := []map[string]interface{}{}
	for i := range model.VirtualNetworkInterfaces {
		vnis = append(vnis, ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(&model.VirtualNetworkInterfaces[i]))
	}
	modelMap["virtual_network_interfaces"] = vnis
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		modelMap["deleted"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)}
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
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerReference) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap
	}
	if model.Deleted != nil {
		modelMap["deleted"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)}
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
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressSession) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("peer session is nil")
	}
	modelMap := make(map[string]interface{})
	if model.EstablishedAt != nil {
		modelMap["established_at"] = flex.DateTimeToString(model.EstablishedAt)
	}
	if model.Local != nil {
		modelMap["local"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model.Local)}
	}
	if model.ProtocolState != nil {
		modelMap["protocol_state"] = *model.ProtocolState
	}
	if model.Remote != nil {
		modelMap["remote"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model.Remote)}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerName) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model != nil && model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("transit gateway peer session is nil")
	}
	modelMap := make(map[string]interface{})
	if model.EstablishedAt != nil {
		modelMap["established_at"] = flex.DateTimeToString(model.EstablishedAt)
	}
	if model.Local != nil {
		modelMap["local"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model.Local)}
	}
	if model.ProtocolState != nil {
		modelMap["protocol_state"] = *model.ProtocolState
	}
	if model.Remote != nil {
		modelMap["remote"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model.Remote)}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model != nil && model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap
}

// ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch converts the patch struct to a map containing only the changed fields.
func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPeerGroupPeerPatch, d *schema.ResourceData) map[string]interface{} {
	patch := map[string]interface{}{}
	if patchVals == nil {
		return patch
	}
	if patchVals.Asn != nil {
		patch["asn"] = *patchVals.Asn
	}
	if patchVals.Name != nil {
		patch["name"] = *patchVals.Name
	}
	if patchVals.Priority != nil {
		patch["priority"] = *patchVals.Priority
	}
	if patchVals.State != nil {
		patch["state"] = *patchVals.State
	}
	return patch
}
