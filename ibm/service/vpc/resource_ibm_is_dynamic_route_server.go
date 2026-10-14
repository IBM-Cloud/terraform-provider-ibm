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
	isDynamicRouteServerStable   = "stable"
	isDynamicRouteServerPending  = "pending"
	isDynamicRouteServerUpdating = "updating"
	isDynamicRouteServerWaiting  = "waiting"
	isDynamicRouteServerDeleting = "deleting"
	isDynamicRouteServerFailed   = "failed"
	isDynamicRouteServerDeleted  = "done"
)

func ResourceIBMIsDynamicRouteServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerRead,
		UpdateContext: resourceIBMIsDynamicRouteServerUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"local_asn": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The local autonomous system number (ASN) used for this dynamic route server and its connections.",
			},
			"members": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				Description: "The initial members for the dynamic route server. Members cannot be changed in place; changing this argument replaces the dynamic route server. Use the `ibm_is_dynamic_route_server_member` resource to add members to an existing dynamic route server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The name for this dynamic route server member. The name is unique across all members in the dynamic route server.",
						},
						"virtual_network_interfaces": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							MinItems:    1,
							Description: "The virtual network interfaces for this dynamic route server member. Each entry identifies an existing virtual network interface by `id`, `href` or `crn`, or describes a new virtual network interface to create with `subnet`.",
							Elem:        ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceSchema(),
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dynamic route server member.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dynamic route server member.",
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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server", "name"),
				Description:  "The name for this dynamic route server. The name is unique across all dynamic route servers in the region.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The resource group for this dynamic route server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The unique identifier for this resource group.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource group.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dynamic route server was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this dynamic route server.",
			},
			"health_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `health_state` (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reason code for this health state.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A link to documentation about the reason for this health state.",
						},
					},
				},
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource: `ok`, `degraded`, `faulted` or `inapplicable`.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server.",
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
				Description: "The lifecycle state of the dynamic route server.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"vpc": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPC this dynamic route server resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this VPC.",
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
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this VPC.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPC.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this VPC. The name is unique across all VPCs in the region.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"etag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ETag of the dynamic route server, sent as If-Match on update and delete.",
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerOptions := &vpcv1.CreateDynamicRouteServerOptions{}

	createDynamicRouteServerOptions.SetLocalAsn(int64(d.Get("local_asn").(int)))
	var members []vpcv1.DynamicRouteServerMemberPrototype
	for _, v := range d.Get("members").([]interface{}) {
		value, _ := v.(map[string]interface{})
		membersItem, err := ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "create", "parse-members").GetDiag()
		}
		members = append(members, *membersItem)
	}
	createDynamicRouteServerOptions.SetMembers(members)
	if _, ok := d.GetOk("name"); ok {
		createDynamicRouteServerOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		resourceGroupModel, err := ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(d.Get("resource_group.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "create", "parse-resource_group").GetDiag()
		}
		createDynamicRouteServerOptions.SetResourceGroup(resourceGroupModel)
	}

	dynamicRouteServer, _, err := vpcClient.CreateDynamicRouteServerWithContext(context, createDynamicRouteServerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*dynamicRouteServer.ID)

	_, err = isWaitForDynamicRouteServerStable(vpcClient, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForDynamicRouteServerStable failed: %s", err.Error()), "ibm_is_dynamic_route_server", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIsDynamicRouteServerRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{}

	getDynamicRouteServerOptions.SetID(d.Id())

	dynamicRouteServer, response, err := vpcClient.GetDynamicRouteServerWithContext(context, getDynamicRouteServerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("local_asn", flex.IntValue(dynamicRouteServer.LocalAsn)); err != nil {
		err = fmt.Errorf("Error setting local_asn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-local_asn").GetDiag()
	}
	members := []map[string]interface{}{}
	for _, membersItem := range dynamicRouteServer.Members {
		membersItemMap, err := ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(&membersItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "members-to-map").GetDiag()
		}
		members = append(members, membersItemMap)
	}
	priorMembers, _ := d.Get("members").([]interface{})
	members = ResourceIBMIsDynamicRouteServerMergeMembers(priorMembers, members)
	if err = d.Set("members", members); err != nil {
		err = fmt.Errorf("Error setting members: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-members").GetDiag()
	}
	if !core.IsNil(dynamicRouteServer.Name) {
		if err = d.Set("name", dynamicRouteServer.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServer.ResourceGroup) {
		resourceGroupMap, err := ResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(dynamicRouteServer.ResourceGroup)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "resource_group-to-map").GetDiag()
		}
		if err = d.Set("resource_group", []map[string]interface{}{resourceGroupMap}); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-resource_group").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServer.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("crn", dynamicRouteServer.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-crn").GetDiag()
	}
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range dynamicRouteServer.HealthReasons {
		healthReasonsItemMap, err := ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(&healthReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "health_reasons-to-map").GetDiag()
		}
		healthReasons = append(healthReasons, healthReasonsItemMap)
	}
	if err = d.Set("health_reasons", healthReasons); err != nil {
		err = fmt.Errorf("Error setting health_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-health_reasons").GetDiag()
	}
	if err = d.Set("health_state", dynamicRouteServer.HealthState); err != nil {
		err = fmt.Errorf("Error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServer.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServer.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", dynamicRouteServer.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", dynamicRouteServer.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(dynamicRouteServer.VPC) {
		vpcMap, err := ResourceIBMIsDynamicRouteServerVPCReferenceToMap(dynamicRouteServer.VPC)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "vpc-to-map").GetDiag()
		}
		if err = d.Set("vpc", []map[string]interface{}{vpcMap}); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "read", "set-vpc").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerOptions := &vpcv1.UpdateDynamicRouteServerOptions{}

	updateDynamicRouteServerOptions.SetID(d.Id())

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerPatch{}
	if d.HasChange("local_asn") {
		newLocalAsn := int64(d.Get("local_asn").(int))
		patchVals.LocalAsn = &newLocalAsn
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		updateDynamicRouteServerOptions.SetIfMatch(d.Get("etag").(string))
		updateDynamicRouteServerOptions.DynamicRouteServerPatch = ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerWithContext(context, updateDynamicRouteServerOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForDynamicRouteServerStable(vpcClient, d.Id(), d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForDynamicRouteServerStable failed: %s", err.Error()), "ibm_is_dynamic_route_server", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Fetch a fresh ETag so the If-Match precondition does not fail because of a stale state value.
	getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{}
	getDynamicRouteServerOptions.SetID(d.Id())
	_, getResponse, err := vpcClient.GetDynamicRouteServerWithContext(context, getDynamicRouteServerOptions)
	if err != nil {
		if getResponse != nil && getResponse.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	etag := getResponse.Headers.Get("Etag")
	if etag == "" {
		etag = d.Get("etag").(string)
	}

	deleteDynamicRouteServerOptions := &vpcv1.DeleteDynamicRouteServerOptions{}

	deleteDynamicRouteServerOptions.SetID(d.Id())
	if etag != "" {
		deleteDynamicRouteServerOptions.SetIfMatch(etag)
	}

	_, response, err := vpcClient.DeleteDynamicRouteServerWithContext(context, deleteDynamicRouteServerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerDeleted(vpcClient, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForDynamicRouteServerDeleted failed: %s", err.Error()), "ibm_is_dynamic_route_server", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isWaitForDynamicRouteServerStable(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server (%s) to be stable.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerPending, isDynamicRouteServerUpdating, isDynamicRouteServerWaiting},
		Target:     []string{isDynamicRouteServerStable, isDynamicRouteServerFailed},
		Refresh:    isDynamicRouteServerRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{
			ID: &id,
		}
		dynamicRouteServer, response, err := sess.GetDynamicRouteServer(getDynamicRouteServerOptions)
		if err != nil {
			return nil, isDynamicRouteServerFailed, fmt.Errorf("[ERROR] Error getting dynamic route server: %s\n%s", err, response)
		}
		if dynamicRouteServer.LifecycleState == nil {
			return dynamicRouteServer, isDynamicRouteServerPending, nil
		}
		if *dynamicRouteServer.LifecycleState == isDynamicRouteServerFailed {
			return dynamicRouteServer, isDynamicRouteServerFailed, fmt.Errorf("[ERROR] Dynamic route server (%s) went into failed state during the operation. Running terraform apply again will remove the tainted dynamic route server and attempt to create it again", id)
		}
		if *dynamicRouteServer.LifecycleState == isDynamicRouteServerStable {
			return dynamicRouteServer, isDynamicRouteServerStable, nil
		}
		if *dynamicRouteServer.LifecycleState == isDynamicRouteServerUpdating || *dynamicRouteServer.LifecycleState == isDynamicRouteServerWaiting {
			return dynamicRouteServer, *dynamicRouteServer.LifecycleState, nil
		}
		return dynamicRouteServer, isDynamicRouteServerPending, nil
	}
}

func isWaitForDynamicRouteServerDeleted(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerDeleting},
		Target:     []string{isDynamicRouteServerDeleted, isDynamicRouteServerFailed},
		Refresh:    isDynamicRouteServerDeleteRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerDeleteRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getDynamicRouteServerOptions := &vpcv1.GetDynamicRouteServerOptions{
			ID: &id,
		}
		dynamicRouteServer, response, err := sess.GetDynamicRouteServer(getDynamicRouteServerOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return dynamicRouteServer, isDynamicRouteServerDeleted, nil
			}
			return nil, isDynamicRouteServerFailed, fmt.Errorf("[ERROR] The dynamic route server %s failed to delete: %s\n%s", id, err, response)
		}
		if dynamicRouteServer.LifecycleState != nil && *dynamicRouteServer.LifecycleState == isDynamicRouteServerFailed {
			return dynamicRouteServer, isDynamicRouteServerFailed, fmt.Errorf("[ERROR] Dynamic route server (%s) went into failed state while deleting", id)
		}
		return dynamicRouteServer, isDynamicRouteServerDeleting, nil
	}
}

// ResourceIBMIsDynamicRouteServerMergeMembers carries the prototype-only virtual network interface
// arguments from the prior state into the member list read from the API. Members are matched by
// name (unique within a dynamic route server), falling back to position.
func ResourceIBMIsDynamicRouteServerMergeMembers(prior []interface{}, current []map[string]interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(current))
	for i, member := range current {
		var priorVNIs []interface{}
		name, _ := member["name"].(string)
		matched := false
		for _, p := range prior {
			pm, ok := p.(map[string]interface{})
			if !ok {
				continue
			}
			if pname, _ := pm["name"].(string); name != "" && pname == name {
				priorVNIs, _ = pm["virtual_network_interfaces"].([]interface{})
				matched = true
				break
			}
		}
		if !matched && i < len(prior) {
			if pm, ok := prior[i].(map[string]interface{}); ok {
				priorVNIs, _ = pm["virtual_network_interfaces"].([]interface{})
			}
		}
		merged := make(map[string]interface{}, len(member))
		for k, v := range member {
			merged[k] = v
		}
		if vnis, ok := member["virtual_network_interfaces"].([]map[string]interface{}); ok {
			merged["virtual_network_interfaces"] = ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(priorVNIs, vnis)
		}
		result = append(result, merged)
	}
	return result
}

func ResourceIBMIsDynamicRouteServerMapToDynamicRouteServerMemberPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerMemberPrototype, error) {
	model := &vpcv1.DynamicRouteServerMemberPrototype{}
	if modelMap == nil {
		return nil, fmt.Errorf("members entry must not be empty")
	}
	name, _ := modelMap["name"].(string)
	if name == "" {
		return nil, fmt.Errorf("members entry requires name")
	}
	model.Name = core.StringPtr(name)
	vniList, _ := modelMap["virtual_network_interfaces"].([]interface{})
	if len(vniList) == 0 {
		return nil, fmt.Errorf("members entry %q requires at least one virtual_network_interfaces entry", name)
	}
	virtualNetworkInterfaces := []vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf{}
	for _, virtualNetworkInterfacesItem := range vniList {
		itemMap, _ := virtualNetworkInterfacesItem.(map[string]interface{})
		virtualNetworkInterfacesItemModel, err := ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(itemMap)
		if err != nil {
			return nil, err
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemModel)
	}
	model.VirtualNetworkInterfaces = virtualNetworkInterfaces
	return model, nil
}

func ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentity(modelMap map[string]interface{}) (vpcv1.ResourceGroupIdentityIntf, error) {
	return ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(modelMap)
}

func ResourceIBMIsDynamicRouteServerMapToResourceGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.ResourceGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("resource_group requires id")
	}
	return &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerDynamicRouteServerMemberReferenceToMap(model *vpcv1.DynamicRouteServerMemberReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerDeletedToMap(model.Deleted)
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
	virtualNetworkInterfaces := []map[string]interface{}{}
	for _, virtualNetworkInterfacesItem := range model.VirtualNetworkInterfaces {
		virtualNetworkInterfacesItemMap, err := ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(&virtualNetworkInterfacesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemMap)
	}
	modelMap["virtual_network_interfaces"] = virtualNetworkInterfaces
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model != nil && model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerDeletedToMap(model.Deleted)
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
	if model.PrimaryIP != nil {
		primaryIPMap, err := ResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(model.PrimaryIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	if model.Subnet != nil {
		subnetMap, err := ResourceIBMIsDynamicRouteServerSubnetReferenceToMap(model.Subnet)
		if err != nil {
			return modelMap, err
		}
		modelMap["subnet"] = []map[string]interface{}{subnetMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerResourceGroupReferenceToMap(model *vpcv1.ResourceGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
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
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerDynamicRouteServerHealthReasonToMap(model *vpcv1.DynamicRouteServerHealthReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
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

func ResourceIBMIsDynamicRouteServerDynamicRouteServerLifecycleReasonToMap(model *vpcv1.DynamicRouteServerLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
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

func ResourceIBMIsDynamicRouteServerVPCReferenceToMap(model *vpcv1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerDeletedToMap(model.Deleted)
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

// ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch converts the patch into the merge-patch
// map that is sent to the service. Only the fields that changed are present in patchVals, so the map
// never touches fields the user did not change. A field that was removed from the configuration is
// sent as null only when the schema allows it to be unset (name is Optional+Computed, so it is kept).
func ResourceIBMIsDynamicRouteServerDynamicRouteServerPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPatch, d *schema.ResourceData) map[string]interface{} {
	patch := map[string]interface{}{}
	if patchVals == nil {
		return patch
	}
	if patchVals.LocalAsn != nil {
		patch["local_asn"] = *patchVals.LocalAsn
	}
	if patchVals.Name != nil && *patchVals.Name != "" {
		patch["name"] = *patchVals.Name
	}
	return patch
}
