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
	isDynamicRouteServerMemberStable   = "stable"
	isDynamicRouteServerMemberPending  = "pending"
	isDynamicRouteServerMemberUpdating = "updating"
	isDynamicRouteServerMemberWaiting  = "waiting"
	isDynamicRouteServerMemberDeleting = "deleting"
	isDynamicRouteServerMemberFailed   = "failed"
	isDynamicRouteServerMemberDeleted  = "done"
)

func ResourceIBMIsDynamicRouteServerMember() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerMemberCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerMemberRead,
		DeleteContext: resourceIBMIsDynamicRouteServerMemberDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_member", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_member", "name"),
				Description:  "The name for this dynamic route server member. The name is unique across all members in the dynamic route server.",
			},
			"virtual_network_interfaces": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MinItems:    1,
				Description: "The virtual network interfaces for this dynamic route server member. Each entry identifies an existing virtual network interface by `id`, `href` or `crn`, or describes a new virtual network interface to create with `subnet`.",
				Elem:        ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceSchema(),
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
				Description: "The health of this dynamic route server member: `ok`, `faulted` or `inapplicable`.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server member.",
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
				Description: "The lifecycle state of the dynamic route server member.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"dynamic_route_server_member_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server member.",
			},
		},
	}
}

// ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceSchema returns the schema shared by the
// `virtual_network_interfaces` block of ibm_is_dynamic_route_server_member and the
// `members.virtual_network_interfaces` block of ibm_is_dynamic_route_server. Each entry is either an
// identity (exactly one of id, href, crn) or an inline virtual network interface prototype (subnet
// plus optional settings). All arguments force a new resource because members cannot be patched.
func ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The unique identifier for an existing virtual network interface. Conflicts with `href`, `crn` and the prototype arguments.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The URL for an existing virtual network interface. Conflicts with `id`, `crn` and the prototype arguments.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The CRN for an existing virtual network interface. Conflicts with `id`, `href` and the prototype arguments.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The name for the virtual network interface. The name is unique across all virtual network interfaces in the VPC. Only used when creating a new virtual network interface.",
			},
			"allow_ip_spoofing": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Indicates whether source IP spoofing is allowed on the virtual network interface. Only used when creating a new virtual network interface.",
			},
			"auto_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Indicates whether the virtual network interface will be automatically deleted when the member is deleted. Only used when creating a new virtual network interface.",
			},
			"enable_infrastructure_nat": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "If `true`, the VPC infrastructure performs any needed NAT operations. Only used when creating a new virtual network interface.",
			},
			"protocol_state_filtering_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The protocol state filtering mode for the virtual network interface: `auto`, `enabled` or `disabled`. Only used when creating a new virtual network interface.",
			},
			"ips": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Additional reserved IPs to bind to the virtual network interface. Each entry identifies an existing reserved IP by `id` or `href`, or describes a reserved IP to create. Only used when creating a new virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The unique identifier for an existing reserved IP.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The URL for an existing reserved IP.",
						},
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The IP address to reserve, which must not already be reserved on the subnet.",
						},
						"auto_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "Indicates whether the reserved IP will be automatically deleted when the virtual network interface is deleted.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The name for the reserved IP. The name is unique across all reserved IPs in a subnet.",
						},
					},
				},
			},
			"primary_ip": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The primary IP for the virtual network interface. Identify an existing reserved IP by `id` or `href`, or describe a reserved IP to create. Only used when creating a new virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The unique identifier for the reserved IP.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The URL for the reserved IP.",
						},
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The IP address. When creating a reserved IP, the address must not already be reserved on the subnet.",
						},
						"auto_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "Indicates whether the reserved IP will be automatically deleted when the virtual network interface is deleted.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The name for the reserved IP. The name is unique across all reserved IPs in a subnet.",
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
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The resource group for the virtual network interface. Only used when creating a new virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The unique identifier for this resource group.",
						},
					},
				},
			},
			"security_groups": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "The security groups for the virtual network interface, each identified by exactly one of `id`, `crn` or `href`. Only used when creating a new virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The unique identifier for this security group.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The CRN for this security group.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The URL for this security group.",
						},
					},
				},
			},
			"subnet": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The subnet for the virtual network interface, identified by exactly one of `id`, `crn` or `href`. Required when creating a new virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The unique identifier for this subnet.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The CRN for this subnet.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The URL for this subnet.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this subnet. The name is unique across all subnets in the VPC.",
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

func ResourceIBMIsDynamicRouteServerMemberValidator() *validate.ResourceValidator {
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
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_member", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerMemberCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerMemberOptions := &vpcv1.CreateDynamicRouteServerMemberOptions{}

	createDynamicRouteServerMemberOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	createDynamicRouteServerMemberOptions.SetName(d.Get("name").(string))
	var virtualNetworkInterfaces []vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf
	for _, v := range d.Get("virtual_network_interfaces").([]interface{}) {
		value := v.(map[string]interface{})
		virtualNetworkInterfacesItem, err := ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "create", "parse-virtual_network_interfaces").GetDiag()
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItem)
	}
	createDynamicRouteServerMemberOptions.SetVirtualNetworkInterfaces(virtualNetworkInterfaces)

	dynamicRouteServerMember, _, err := vpcClient.CreateDynamicRouteServerMemberWithContext(context, createDynamicRouteServerMemberOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerMemberWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createDynamicRouteServerMemberOptions.DynamicRouteServerID, *dynamicRouteServerMember.ID))

	_, err = isWaitForDynamicRouteServerMemberStable(vpcClient, *createDynamicRouteServerMemberOptions.DynamicRouteServerID, *dynamicRouteServerMember.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForDynamicRouteServerMemberStable failed: %s", err.Error()), "ibm_is_dynamic_route_server_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIsDynamicRouteServerMemberRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerMemberRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerMemberOptions := &vpcv1.GetDynamicRouteServerMemberOptions{}

	// The ID is "<dynamic_route_server_id>/<member_id>" so that import works without any arguments.
	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "sep-id-parts").GetDiag()
	}
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		err = fmt.Errorf("invalid ID %q: expected <dynamic_route_server_id>/<member_id>", d.Id())
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerMemberOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerMemberOptions.SetID(parts[1])

	dynamicRouteServerMember, response, err := vpcClient.GetDynamicRouteServerMemberWithContext(context, getDynamicRouteServerMemberOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerMemberWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_member", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("dynamic_route_server_id", parts[0]); err != nil {
		err = fmt.Errorf("Error setting dynamic_route_server_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-dynamic_route_server_id").GetDiag()
	}
	if err = d.Set("name", dynamicRouteServerMember.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-name").GetDiag()
	}
	virtualNetworkInterfaces := []map[string]interface{}{}
	for _, virtualNetworkInterfacesItem := range dynamicRouteServerMember.VirtualNetworkInterfaces {
		virtualNetworkInterfacesItemMap, err := ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(&virtualNetworkInterfacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "virtual_network_interfaces-to-map").GetDiag()
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemMap)
	}
	prior, _ := d.Get("virtual_network_interfaces").([]interface{})
	virtualNetworkInterfaces = ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(prior, virtualNetworkInterfaces)
	if err = d.Set("virtual_network_interfaces", virtualNetworkInterfaces); err != nil {
		err = fmt.Errorf("Error setting virtual_network_interfaces: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-virtual_network_interfaces").GetDiag()
	}
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range dynamicRouteServerMember.HealthReasons {
		healthReasonsItemMap, err := ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(&healthReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "health_reasons-to-map").GetDiag()
		}
		healthReasons = append(healthReasons, healthReasonsItemMap)
	}
	if err = d.Set("health_reasons", healthReasons); err != nil {
		err = fmt.Errorf("Error setting health_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-health_reasons").GetDiag()
	}
	if err = d.Set("health_state", dynamicRouteServerMember.HealthState); err != nil {
		err = fmt.Errorf("Error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerMember.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServerMember.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", dynamicRouteServerMember.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", dynamicRouteServerMember.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("dynamic_route_server_member_id", dynamicRouteServerMember.ID); err != nil {
		err = fmt.Errorf("Error setting dynamic_route_server_member_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "read", "set-dynamic_route_server_member_id").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerMemberDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerMemberOptions := &vpcv1.DeleteDynamicRouteServerMemberOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "delete", "sep-id-parts").GetDiag()
	}
	if len(parts) != 2 {
		err = fmt.Errorf("invalid ID %q: expected <dynamic_route_server_id>/<member_id>", d.Id())
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_member", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerMemberOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerMemberOptions.SetID(parts[1])

	_, response, err := vpcClient.DeleteDynamicRouteServerMemberWithContext(context, deleteDynamicRouteServerMemberOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerMemberWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerMemberDeleted(vpcClient, parts[0], parts[1], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForDynamicRouteServerMemberDeleted failed: %s", err.Error()), "ibm_is_dynamic_route_server_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isWaitForDynamicRouteServerMemberStable(sess *vpcv1.VpcV1, dynamicRouteServerID string, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server member (%s/%s) to be stable.", dynamicRouteServerID, id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerMemberPending, isDynamicRouteServerMemberUpdating, isDynamicRouteServerMemberWaiting},
		Target:     []string{isDynamicRouteServerMemberStable, isDynamicRouteServerMemberFailed},
		Refresh:    isDynamicRouteServerMemberRefreshFunc(sess, dynamicRouteServerID, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerMemberRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID string, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getDynamicRouteServerMemberOptions := &vpcv1.GetDynamicRouteServerMemberOptions{
			DynamicRouteServerID: &dynamicRouteServerID,
			ID:                   &id,
		}
		member, response, err := sess.GetDynamicRouteServerMember(getDynamicRouteServerMemberOptions)
		if err != nil {
			return nil, isDynamicRouteServerMemberFailed, fmt.Errorf("[ERROR] Error getting dynamic route server member: %s\n%s", err, response)
		}
		if member.LifecycleState == nil {
			return member, isDynamicRouteServerMemberPending, nil
		}
		if *member.LifecycleState == isDynamicRouteServerMemberFailed {
			return member, isDynamicRouteServerMemberFailed, fmt.Errorf("[ERROR] Dynamic route server member (%s) went into failed state during the operation. Running terraform apply again will remove the tainted member and attempt to create it again", id)
		}
		if *member.LifecycleState == isDynamicRouteServerMemberStable {
			return member, isDynamicRouteServerMemberStable, nil
		}
		if *member.LifecycleState == isDynamicRouteServerMemberUpdating || *member.LifecycleState == isDynamicRouteServerMemberWaiting {
			return member, *member.LifecycleState, nil
		}
		return member, isDynamicRouteServerMemberPending, nil
	}
}

func isWaitForDynamicRouteServerMemberDeleted(sess *vpcv1.VpcV1, dynamicRouteServerID string, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server member (%s/%s) to be deleted.", dynamicRouteServerID, id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerMemberDeleting},
		Target:     []string{isDynamicRouteServerMemberDeleted, isDynamicRouteServerMemberFailed},
		Refresh:    isDynamicRouteServerMemberDeleteRefreshFunc(sess, dynamicRouteServerID, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDynamicRouteServerMemberDeleteRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID string, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getDynamicRouteServerMemberOptions := &vpcv1.GetDynamicRouteServerMemberOptions{
			DynamicRouteServerID: &dynamicRouteServerID,
			ID:                   &id,
		}
		member, response, err := sess.GetDynamicRouteServerMember(getDynamicRouteServerMemberOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return member, isDynamicRouteServerMemberDeleted, nil
			}
			return nil, isDynamicRouteServerMemberFailed, fmt.Errorf("[ERROR] The dynamic route server member %s failed to delete: %s\n%s", id, err, response)
		}
		if member.LifecycleState != nil && *member.LifecycleState == isDynamicRouteServerMemberFailed {
			return member, isDynamicRouteServerMemberFailed, fmt.Errorf("[ERROR] Dynamic route server member (%s) went into failed state while deleting", id)
		}
		return member, isDynamicRouteServerMemberDeleting, nil
	}
}

// ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces carries prototype-only arguments
// (settings that are sent on create but are not part of the virtual network interface reference the
// API returns) from the prior state into the freshly read list, so that a refresh does not produce a
// spurious diff. Entries are matched by id, then by name, then by position. When nothing is known
// (for example on import) the schema defaults are used.
func ResourceIBMIsDynamicRouteServerMemberMergeVirtualNetworkInterfaces(prior []interface{}, current []map[string]interface{}) []map[string]interface{} {
	findPrior := func(index int, item map[string]interface{}) map[string]interface{} {
		id, _ := item["id"].(string)
		name, _ := item["name"].(string)
		for _, p := range prior {
			pm, ok := p.(map[string]interface{})
			if !ok {
				continue
			}
			if pid, _ := pm["id"].(string); id != "" && pid == id {
				return pm
			}
		}
		for _, p := range prior {
			pm, ok := p.(map[string]interface{})
			if !ok {
				continue
			}
			pid, _ := pm["id"].(string)
			if pname, _ := pm["name"].(string); name != "" && pname == name && pid == "" {
				return pm
			}
		}
		if index < len(prior) {
			if pm, ok := prior[index].(map[string]interface{}); ok {
				pid, _ := pm["id"].(string)
				pname, _ := pm["name"].(string)
				if pid == "" && pname == "" {
					return pm
				}
			}
		}
		return nil
	}

	result := make([]map[string]interface{}, 0, len(current))
	for i, item := range current {
		merged := make(map[string]interface{}, len(item)+7)
		for k, v := range item {
			merged[k] = v
		}
		pm := findPrior(i, item)
		boolFrom := func(key string, def bool) bool {
			if pm != nil {
				if v, ok := pm[key].(bool); ok {
					return v
				}
			}
			return def
		}
		merged["allow_ip_spoofing"] = boolFrom("allow_ip_spoofing", false)
		merged["auto_delete"] = boolFrom("auto_delete", true)
		merged["enable_infrastructure_nat"] = boolFrom("enable_infrastructure_nat", true)
		if pm != nil {
			if v, ok := pm["protocol_state_filtering_mode"].(string); ok && v != "" {
				merged["protocol_state_filtering_mode"] = v
			}
			if v, ok := pm["ips"].([]interface{}); ok && len(v) > 0 {
				merged["ips"] = v
			}
			if v, ok := pm["security_groups"].([]interface{}); ok && len(v) > 0 {
				merged["security_groups"] = v
			}
			if v, ok := pm["resource_group"].([]interface{}); ok && len(v) > 0 {
				merged["resource_group"] = v
			}
		}
		if primaryIPs, ok := merged["primary_ip"].([]map[string]interface{}); ok && len(primaryIPs) > 0 {
			primaryIP := make(map[string]interface{}, len(primaryIPs[0])+1)
			for k, v := range primaryIPs[0] {
				primaryIP[k] = v
			}
			autoDelete := true
			if pm != nil {
				if pips, ok := pm["primary_ip"].([]interface{}); ok && len(pips) > 0 {
					if pip, ok := pips[0].(map[string]interface{}); ok {
						if v, ok := pip["auto_delete"].(bool); ok {
							autoDelete = v
						}
					}
				}
			}
			primaryIP["auto_delete"] = autoDelete
			merged["primary_ip"] = []map[string]interface{}{primaryIP}
		}
		result = append(result, merged)
	}
	return result
}

func dynamicRouteServerMemberIdentityKeys(modelMap map[string]interface{}, keys ...string) []string {
	set := []string{}
	for _, key := range keys {
		if v, ok := modelMap[key].(string); ok && v != "" {
			set = append(set, key)
		}
	}
	return set
}

// ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface builds
// the oneOf prototype: an identity when exactly one of id, href or crn is set, otherwise an inline
// virtual network interface prototype (which requires subnet).
func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterface(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceIntf, error) {
	if modelMap == nil {
		return nil, fmt.Errorf("virtual_network_interfaces entry must not be empty")
	}
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href", "crn")
	switch len(identity) {
	case 0:
		return ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(modelMap)
	case 1:
		return ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(modelMap)
	default:
		return nil, fmt.Errorf("virtual_network_interfaces entry must set only one of id, href or crn, got %v", identity)
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href", "crn")
	if len(identity) != 1 {
		return nil, fmt.Errorf("virtual network interface identity requires exactly one of id, href or crn")
	}
	switch identity[0] {
	case "id":
		return ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(modelMap)
	case "href":
		return ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(modelMap)
	default:
		return ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(modelMap)
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("virtual network interface identity by id requires id")
	}
	return &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("virtual network interface identity by href requires href")
	}
	return &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN, error) {
	crn, _ := modelMap["crn"].(string)
	if crn == "" {
		return nil, fmt.Errorf("virtual network interface identity by crn requires crn")
	}
	return &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN{CRN: core.StringPtr(crn)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToDynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext, error) {
	model := &vpcv1.DynamicRouteServerMemberPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeDynamicRouteServerContext{}
	subnets, _ := modelMap["subnet"].([]interface{})
	if len(subnets) == 0 || subnets[0] == nil {
		return nil, fmt.Errorf("subnet is required when a virtual network interface is not identified by id, href or crn")
	}
	subnetModel, err := ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(subnets[0].(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	model.Subnet = subnetModel
	if v, ok := modelMap["allow_ip_spoofing"].(bool); ok {
		model.AllowIPSpoofing = core.BoolPtr(v)
	}
	if v, ok := modelMap["auto_delete"].(bool); ok {
		model.AutoDelete = core.BoolPtr(v)
	}
	if v, ok := modelMap["enable_infrastructure_nat"].(bool); ok {
		model.EnableInfrastructureNat = core.BoolPtr(v)
	}
	if ipsList, ok := modelMap["ips"].([]interface{}); ok && len(ipsList) > 0 {
		ips := []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf{}
		for _, ipsItem := range ipsList {
			ipsItemModel, err := ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(ipsItem.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			ips = append(ips, ipsItemModel)
		}
		model.Ips = ips
	}
	if v, ok := modelMap["name"].(string); ok && v != "" {
		model.Name = core.StringPtr(v)
	}
	if primaryIPs, ok := modelMap["primary_ip"].([]interface{}); ok && len(primaryIPs) > 0 && primaryIPs[0] != nil {
		primaryIPModel, err := ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(primaryIPs[0].(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		model.PrimaryIP = primaryIPModel
	}
	if v, ok := modelMap["protocol_state_filtering_mode"].(string); ok && v != "" {
		model.ProtocolStateFilteringMode = core.StringPtr(v)
	}
	if resourceGroups, ok := modelMap["resource_group"].([]interface{}); ok && len(resourceGroups) > 0 && resourceGroups[0] != nil {
		resourceGroupModel, err := ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(resourceGroups[0].(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		model.ResourceGroup = resourceGroupModel
	}
	if securityGroupsList, ok := modelMap["security_groups"].([]interface{}); ok && len(securityGroupsList) > 0 {
		securityGroups := []vpcv1.SecurityGroupIdentityIntf{}
		for _, securityGroupsItem := range securityGroupsList {
			securityGroupsItemModel, err := ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(securityGroupsItem.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			securityGroups = append(securityGroups, securityGroupsItemModel)
		}
		model.SecurityGroups = securityGroups
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototype(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfaceIPPrototypeIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href")
	switch len(identity) {
	case 0:
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(modelMap)
	case 1:
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(modelMap)
	default:
		return nil, fmt.Errorf("ips entry must set only one of id or href")
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href")
	if len(identity) != 1 {
		return nil, fmt.Errorf("reserved IP identity requires exactly one of id or href")
	}
	if identity[0] == "id" {
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(modelMap)
	}
	return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(modelMap)
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("reserved IP identity by id requires id")
	}
	return &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("reserved IP identity by href requires href")
	}
	return &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext, error) {
	model := &vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext{}
	if v, ok := modelMap["address"].(string); ok && v != "" {
		model.Address = core.StringPtr(v)
	}
	if v, ok := modelMap["auto_delete"].(bool); ok {
		model.AutoDelete = core.BoolPtr(v)
	}
	if v, ok := modelMap["name"].(string); ok && v != "" {
		model.Name = core.StringPtr(v)
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototype(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href")
	switch len(identity) {
	case 0:
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(modelMap)
	case 1:
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(modelMap)
	default:
		return nil, fmt.Errorf("primary_ip must set only one of id or href")
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "href")
	if len(identity) != 1 {
		return nil, fmt.Errorf("primary IP identity requires exactly one of id or href")
	}
	if identity[0] == "id" {
		return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(modelMap)
	}
	return ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(modelMap)
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("primary IP identity by id requires id")
	}
	return &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("primary IP identity by href requires href")
	}
	return &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext(modelMap map[string]interface{}) (*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext, error) {
	model := &vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext{}
	if v, ok := modelMap["address"].(string); ok && v != "" {
		model.Address = core.StringPtr(v)
	}
	if v, ok := modelMap["auto_delete"].(bool); ok {
		model.AutoDelete = core.BoolPtr(v)
	}
	if v, ok := modelMap["name"].(string); ok && v != "" {
		model.Name = core.StringPtr(v)
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentity(modelMap map[string]interface{}) (vpcv1.ResourceGroupIdentityIntf, error) {
	return ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(modelMap)
}

func ResourceIBMIsDynamicRouteServerMemberMapToResourceGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.ResourceGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("resource_group requires id")
	}
	return &vpcv1.ResourceGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentity(modelMap map[string]interface{}) (vpcv1.SecurityGroupIdentityIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "crn", "href")
	if len(identity) != 1 {
		return nil, fmt.Errorf("security_groups entry requires exactly one of id, crn or href")
	}
	switch identity[0] {
	case "id":
		return ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(modelMap)
	case "crn":
		return ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(modelMap)
	default:
		return ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(modelMap)
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.SecurityGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("security group identity by id requires id")
	}
	return &vpcv1.SecurityGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.SecurityGroupIdentityByCRN, error) {
	crn, _ := modelMap["crn"].(string)
	if crn == "" {
		return nil, fmt.Errorf("security group identity by crn requires crn")
	}
	return &vpcv1.SecurityGroupIdentityByCRN{CRN: core.StringPtr(crn)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSecurityGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.SecurityGroupIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("security group identity by href requires href")
	}
	return &vpcv1.SecurityGroupIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentity(modelMap map[string]interface{}) (vpcv1.SubnetIdentityIntf, error) {
	identity := dynamicRouteServerMemberIdentityKeys(modelMap, "id", "crn", "href")
	if len(identity) != 1 {
		return nil, fmt.Errorf("subnet requires exactly one of id, crn or href")
	}
	switch identity[0] {
	case "id":
		return ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(modelMap)
	case "crn":
		return ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(modelMap)
	default:
		return ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(modelMap)
	}
}

func ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByID(modelMap map[string]interface{}) (*vpcv1.SubnetIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("subnet identity by id requires id")
	}
	return &vpcv1.SubnetIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.SubnetIdentityByCRN, error) {
	crn, _ := modelMap["crn"].(string)
	if crn == "" {
		return nil, fmt.Errorf("subnet identity by crn requires crn")
	}
	return &vpcv1.SubnetIdentityByCRN{CRN: core.StringPtr(crn)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberMapToSubnetIdentityByHref(modelMap map[string]interface{}) (*vpcv1.SubnetIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("subnet identity by href requires href")
	}
	return &vpcv1.SubnetIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerMemberVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model.Deleted)
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
		primaryIPMap, err := ResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(model.PrimaryIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	if model.Subnet != nil {
		subnetMap, err := ResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(model.Subnet)
		if err != nil {
			return modelMap, err
		}
		modelMap["subnet"] = []map[string]interface{}{subnetMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model != nil && model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerMemberReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerMemberSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerMemberDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerHealthReasonToMap(model *vpcv1.DynamicRouteServerHealthReason) (map[string]interface{}, error) {
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

func ResourceIBMIsDynamicRouteServerMemberDynamicRouteServerLifecycleReasonToMap(model *vpcv1.DynamicRouteServerLifecycleReason) (map[string]interface{}, error) {
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
