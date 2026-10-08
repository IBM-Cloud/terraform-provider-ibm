// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	isVPNGatewayName              = "name"
	isVPNGatewayResourceGroup     = "resource_group"
	isVPNGatewayMode              = "mode"
	isVPNGatewayLocalAsn          = "local_asn"
	isVPNGatewayAdvertisedCidrs   = "advertised_cidrs"
	isVPNGatewayCRN               = "crn"
	isVPNGatewayTags              = "tags"
	isVPNGatewaySubnet            = "subnet"
	isVPNGatewayStatus            = "status"
	isVPNGatewayDeleting          = "deleting"
	isVPNGatewayDeleted           = "done"
	isVPNGatewayProvisioning      = "provisioning"
	isVPNGatewayProvisioningDone  = "done"
	isVPNGatewayPublicIPAddress   = "public_ip_address"
	isVPNGatewayMembers           = "members"
	isVPNGatewayCreatedAt         = "created_at"
	isVPNGatewayPublicIPAddress2  = "public_ip_address2"
	isVPNGatewayPrivateIPAddress  = "private_ip_address"
	isVPNGatewayPrivateIPAddress2 = "private_ip_address2"
	isVPNGatewayAccessTags        = "access_tags"
	isVPNGatewayHealthState       = "health_state"
	isVPNGatewayHealthReasons     = "health_reasons"
	isVPNGatewayLifecycleState    = "lifecycle_state"
	isVPNGatewayLifecycleReasons  = "lifecycle_reasons"
	isVPNGatewayAvailabilityMode  = "availability_mode"
)

func ResourceIBMISVPNGateway() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPNGatewayCreate,
		ReadContext:   resourceIBMISVPNGatewayRead,
		UpdateContext: resourceIBMISVPNGatewayUpdate,
		DeleteContext: resourceIBMISVPNGatewayDelete,
		Exists:        resourceIBMISVPNGatewayExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			resourceIBMISVPNGatewayCustomizeDiff,
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{

			isVPNGatewayName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway", isVPNGatewayName),
				Description:  "VPN Gateway instance name",
			},

			isVPNGatewaySubnet: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The subnet for a zonal VPN gateway. Required when `availability_mode` is `zonal` or not set. Must not be set for a regional VPN gateway.",
			},

			isVPNGatewayResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group for this VPN gateway",
			},

			isVPNGatewayStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the VPN gateway",
			},

			isVPNGatewayHealthState: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.- `ok`: Healthy- `degraded`: Suffering from compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.",
			},
			isVPNGatewayHealthReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},

						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},

						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},

			isVPNGatewayPublicIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The public IP address assigned to the VPN gateway member.",
			},

			isVPNGatewayPublicIPAddress2: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The second public IP address assigned to the VPN gateway member.",
			},

			isVPNGatewayPrivateIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Private IP address assigned to the VPN gateway member.",
			},

			isVPNGatewayPrivateIPAddress2: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Second Private IP address assigned to the VPN gateway member.",
			},

			isVPNGatewayTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "VPN Gateway tags list",
			},

			isVPNGatewayAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
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

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			isVPNGatewayCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			isVPNGatewayCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created Time of the VPN Gateway",
			},
			isVPNGatewayLifecycleState: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the VPN route.",
			},
			isVPNGatewayLifecycleReasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current lifecycle_state (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this lifecycle state.",
						},

						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this lifecycle state.",
						},

						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},
			isVPNGatewayMode: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "route",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway", isVPNGatewayMode),
				Description:  "mode in VPN gateway(route/policy)",
			},

			isVPNGatewayLocalAsn: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The local autonomous system number (ASN) for this VPN gateway and its connections.",
			},

			isVPNGatewayAdvertisedCidrs: {
				Type: schema.TypeList,
				// Optional:    true,
				Computed:    true,
				Description: "The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			isVPNGatewayMembers: {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				MinItems:    2,
				MaxItems:    2,
				Description: "The members for the VPN gateway. Set this only for a regional VPN gateway, with one `private_ip.subnet` per member. The subnet is used at create time only: later member moves are done with `ibm_is_vpn_gateway_member_replace` and are not treated as drift.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public IP address assigned to the VPN gateway member",
						},

						"private_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private IP address assigned to the VPN gateway member",
						},
						"role": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The high availability role assigned to the VPN gateway member.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN gateway member",
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
										Description: "A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
							Description: "The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPN gateway member.",
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
										Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
							Description: "The lifecycle state of the VPN gateway member.",
						},
						"private_ip": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The reserved IP address assigned to the VPN gateway member. On create, set `subnet` to choose the member subnet.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
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
										Description: "The URL for this reserved IP.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this reserved IP.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
									"subnet": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
										MaxItems:    1,
										Description: "The subnet of the VPN gateway member. Specify one of `id`, `crn` or `href`.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"crn": &schema.Schema{
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: suppressVPNGatewayMemberSubnetDiff,
													Description:      "The CRN for this subnet.",
												},
												"deleted": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
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
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: suppressVPNGatewayMemberSubnetDiff,
													Description:      "The URL for this subnet.",
												},
												"id": &schema.Schema{
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: suppressVPNGatewayMemberSubnetDiff,
													Description:      "The unique identifier for this subnet.",
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
											},
										},
									},
								},
							},
						},
						"public_ip": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The public IP address assigned to the VPN gateway member.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
									},
								},
							},
						},
					},
				},
			},
			isVPNGatewayAvailabilityMode: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway", isVPNGatewayAvailabilityMode),
				Description:  "The availability mode of the VPN gateway. `zonal`: the gateway lives in the single zone of its `subnet`. `regional`: the gateway has two `members` that can be in the same or different zones of the region (route-based gateways only). This value cannot be updated in place; changing it recreates the gateway.",
			},
			"vpc": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VPC for the VPN Gateway",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this VPC.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this VPC.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPC.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this VPC.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISVPNGatewayValidator() *validate.ResourceValidator {

	modeCheckTypes := "route,policy"
	availabilityModeCheckTypes := "zonal,regional"
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayAvailabilityMode,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   false,
			AllowedValues:              availabilityModeCheckTypes})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayMode,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   false,
			AllowedValues:              modeCheckTypes})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISVPNGatewayResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpn_gateway", Schema: validateSchema}
	return &ibmISVPNGatewayResourceValidator
}

func resourceIBMISVPNGatewayCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] VPNGateway create")
	name := d.Get(isVPNGatewayName).(string)
	subnetID := d.Get(isVPNGatewaySubnet).(string)
	mode := d.Get(isVPNGatewayMode).(string)

	err := vpngwCreate(context, d, meta, name, subnetID, mode)
	if err != nil {
		return err
	}
	return resourceIBMISVPNGatewayRead(context, d, meta)
}

func vpngwCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, subnetID, mode string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpnGatewayPrototype := &vpcv1.VPNGatewayPrototype{
		Name: &name,
		Mode: &mode,
	}
	if subnetID != "" {
		vpnGatewayPrototype.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnetID,
		}
	}
	if availabilityMode, ok := d.GetOk(isVPNGatewayAvailabilityMode); ok {
		vpnGatewayPrototype.AvailabilityMode = core.StringPtr(availabilityMode.(string))
	}
	if membersIntf, ok := d.GetOk(isVPNGatewayMembers); ok {
		members, err := resourceIBMISVPNGatewayMapToMemberPrototypes(membersIntf.([]interface{}))
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "create", "parse-members")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		vpnGatewayPrototype.Members = members
	}
	if localAsnIntf, ok := d.GetOk(isVPNGatewayLocalAsn); ok {
		localAsn := int64(localAsnIntf.(int))
		vpnGatewayPrototype.LocalAsn = &localAsn
	}

	options := &vpcv1.CreateVPNGatewayOptions{
		VPNGatewayPrototype: vpnGatewayPrototype,
	}

	if rgrp, ok := d.GetOk(isVPNGatewayResourceGroup); ok {
		rg := rgrp.(string)
		vpnGatewayPrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	vpnGatewayIntf, _, err := sess.CreateVPNGatewayWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)

	d.SetId(*vpnGateway.ID)
	log.Printf("[INFO] VPNGateway : %s", *vpnGateway.ID)

	_, err = isWaitForVpnGatewayAvailable(sess, *vpnGateway.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpnGatewayAvailable failed: %s", err.Error()), "ibm_is_vpn_gateway", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVPNGatewayTags); ok || v != "" {
		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpnGateway.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isVPNGatewayAccessTags); ok {
		oldList, newList := d.GetChange(isVPNGatewayAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpnGateway.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource VPN Gateway (%s) access tags: %s", d.Id(), err)
		}
	}

	return nil
}

func isWaitForVpnGatewayAvailable(vpnGateway *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for vpn gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayProvisioning},
		Target:     []string{isVPNGatewayProvisioningDone, ""},
		Refresh:    isVpnGatewayRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpnGatewayRefreshFunc(vpnGateway *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
			ID: &id,
		}
		vpnGatewayIntf, response, err := vpnGateway.GetVPNGateway(getVpnGatewayOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Vpn Gateway: %s\n%s", err, response)
		}
		vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)

		if *vpnGateway.LifecycleState == "stable" || *vpnGateway.LifecycleState == "failed" || *vpnGateway.LifecycleState == "suspended" {
			return vpnGateway, isVPNGatewayProvisioningDone, nil
		}

		return vpnGateway, isVPNGatewayProvisioning, nil
	}
}

func resourceIBMISVPNGatewayRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := vpngwGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func vpngwGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
		ID: &id,
	}
	vpnGatewayIntf, response, err := sess.GetVPNGatewayWithContext(context, getVpnGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)

	if !core.IsNil(vpnGateway.Name) {
		if err = d.Set("name", vpnGateway.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(vpnGateway.AvailabilityMode) {
		if err = d.Set(isVPNGatewayAvailabilityMode, vpnGateway.AvailabilityMode); err != nil {
			err = fmt.Errorf("Error setting availability_mode: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-availability_mode").GetDiag()
		}
	}
	if !core.IsNil(vpnGateway.Subnet) {
		if err = d.Set(isVPNGatewaySubnet, *vpnGateway.Subnet.ID); err != nil {
			err = fmt.Errorf("Error setting subnet: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-subnet").GetDiag()
		}
	}

	if err = d.Set("health_state", vpnGateway.HealthState); err != nil {
		err = fmt.Errorf("Error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("health_reasons", resourceVPNGatewayRouteFlattenHealthReasons(vpnGateway.HealthReasons)); err != nil {
		err = fmt.Errorf("Error setting health_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-health_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", vpnGateway.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("lifecycle_reasons", resourceVPNGatewayFlattenLifecycleReasons(vpnGateway.LifecycleReasons)); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-lifecycle_reasons").GetDiag()
	}
	members := []vpcv1.VPNGatewayMember{}
	for _, member := range vpnGateway.Members {
		members = append(members, member)
	}
	if len(members) > 0 && members[0].PublicIP != nil && members[0].PublicIP.Address != nil {
		if err = d.Set(isVPNGatewayPublicIPAddress, *members[0].PublicIP.Address); err != nil {
			err = fmt.Errorf("Error setting public_ip_address: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-public_ip_address").GetDiag()
		}
	}
	if len(members) > 0 {
		if members[0].PrivateIP != nil && members[0].PrivateIP.Address != nil {
			if err = d.Set(isVPNGatewayPrivateIPAddress, *members[0].PrivateIP.Address); err != nil {
				err = fmt.Errorf("Error setting private_ip_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-private_ip_address").GetDiag()
			}
		}
	}
	if len(members) > 1 {
		if members[1].PublicIP != nil && members[1].PublicIP.Address != nil {
			if err = d.Set(isVPNGatewayPublicIPAddress2, *members[1].PublicIP.Address); err != nil {
				err = fmt.Errorf("Error setting public_ip_address2: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-public_ip_address2").GetDiag()
			}
		}
		if members[1].PrivateIP != nil && members[1].PrivateIP.Address != nil {
			if err = d.Set(isVPNGatewayPrivateIPAddress2, *members[1].PrivateIP.Address); err != nil {
				err = fmt.Errorf("Error setting private_ip_address2: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-private_ip_address2").GetDiag()
			}
		}

	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *vpnGateway.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isVPNGatewayTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *vpnGateway.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource VPC VPN Gateway (%s) access tags: %s", d.Id(), err)
	}

	if err = d.Set(isVPNGatewayAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-access_tags").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_vpn_gateway", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc/network/vpngateways"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *vpnGateway.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *vpnGateway.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("crn", vpnGateway.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-crn").GetDiag()
	}
	if vpnGateway.ResourceGroup != nil {
		if err = d.Set(flex.ResourceGroupName, vpnGateway.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-resource_group_name").GetDiag()
		}
		if err = d.Set(isVPNGatewayResourceGroup, vpnGateway.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-resource_group").GetDiag()
		}
	}
	if err = d.Set(isVPNGatewayAdvertisedCidrs, vpnGateway.AdvertisedCIDRs); err != nil {
		err = fmt.Errorf("Error setting advertised_cidrs: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-advertised_cidrs").GetDiag()
	}
	if vpnGateway.LocalAsn != nil {
		if err = d.Set(isVPNGatewayLocalAsn, *vpnGateway.LocalAsn); err != nil {
			err = fmt.Errorf("Error setting local_asn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-local_asn").GetDiag()
		}
	}
	if err = d.Set(isVPNGatewayMode, *vpnGateway.Mode); err != nil {
		err = fmt.Errorf("Error setting mode: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-mode").GetDiag()
	}
	if vpnGateway.Members != nil {
		memberSchema := ResourceIBMISVPNGateway().Schema[isVPNGatewayMembers].Elem.(*schema.Resource).Schema
		vpcMembersIpsList := make([]map[string]interface{}, 0, len(vpnGateway.Members))
		for _, membersItem := range vpnGateway.Members {
			membersItemMap, err := vpnGatewayMemberMapForSchema(&membersItem, memberSchema) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "members-to-map").GetDiag()
			}
			vpcMembersIpsList = append(vpcMembersIpsList, membersItemMap)
		}
		if err = d.Set(isVPNGatewayMembers, vpcMembersIpsList); err != nil {
			err = fmt.Errorf("Error setting members: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-members").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(vpnGateway.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-created_at").GetDiag()
	}
	if vpnGateway.VPC != nil {
		vpcList := []map[string]interface{}{}
		vpcList = append(vpcList, dataSourceVPNServerCollectionVPNGatewayVpcReferenceToMap(vpnGateway.VPC))
		if err = d.Set("vpc", vpcList); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "read", "set-vpc").GetDiag()
		}
	}
	return nil
}

func resourceIBMISVPNGatewayUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	hasChanged := false

	err := vpngwUpdate(context, d, meta, id, hasChanged)
	if err != nil {
		return err
	}
	return resourceIBMISVPNGatewayRead(context, d, meta)
}

func vpngwUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id string, hasChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isVPNGatewayTags) {
		getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
			ID: &id,
		}
		vpnGatewayIntf, _, err := sess.GetVPNGatewayWithContext(context, getVpnGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)

		oldList, newList := d.GetChange(isVPNGatewayTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpnGateway.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Vpn Gateway (%s) tags: %s", id, err)
		}
	}
	if d.HasChange(isVPNGatewayAccessTags) {
		getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
			ID: &id,
		}
		vpnGatewayIntf, _, err := sess.GetVPNGatewayWithContext(context, getVpnGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)

		oldList, newList := d.GetChange(isVPNGatewayAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpnGateway.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource VPC VPN Gateway  (%s) access tags: %s", d.Id(), err)
		}
	}

	options := &vpcv1.UpdateVPNGatewayOptions{
		ID: &id,
	}
	vpnGatewayPatchModel := &vpcv1.VPNGatewayPatch{}
	if d.HasChange(isVPNGatewayName) {
		name := d.Get(isVPNGatewayName).(string)
		vpnGatewayPatchModel.Name = &name
		hasChanged = true
	}

	if d.HasChange(isVPNGatewayLocalAsn) {
		if localAsnIntf, ok := d.GetOk(isVPNGatewayLocalAsn); ok {
			localAsn := core.Int64Ptr(int64(localAsnIntf.(int)))
			vpnGatewayPatchModel.LocalAsn = localAsn
			hasChanged = true
		}
	}

	if hasChanged {
		vpnGatewayPatch, err := vpnGatewayPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpnGatewayPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_vpn_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.VPNGatewayPatch = vpnGatewayPatch
		_, _, err = sess.UpdateVPNGatewayWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISVPNGatewayDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()

	err := vpngwDelete(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func vpngwDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVPNGatewayWithContext(context, getVpnGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.DeleteVPNGatewayOptions{
		ID: &id,
	}
	response, err = sess.DeleteVPNGatewayWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForVpnGatewayDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpnGatewayDeleted failed: %s", err.Error()), "ibm_is_vpn_gateway", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForVpnGatewayDeleted(vpnGateway *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayDeleting},
		Target:     []string{isVPNGatewayDeleted, ""},
		Refresh:    isVpnGatewayDeleteRefreshFunc(vpnGateway, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpnGatewayDeleteRefreshFunc(vpnGateway *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
			ID: &id,
		}
		vpngw, response, err := vpnGateway.GetVPNGateway(getVpnGatewayOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return "", isVPNGatewayDeleted, nil
			}
			return "", "", fmt.Errorf("[ERROR] Error Getting Vpn Gateway: %s\n%s", err, response)
		}
		return vpngw, isVPNGatewayDeleting, err
	}
}

func resourceIBMISVPNGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()

	exists, err := vpngwExists(d, meta, id)
	return exists, err
}

func vpngwExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getVpnGatewayOptions := &vpcv1.GetVPNGatewayOptions{
		ID: &id,
	}
	_, response, err := sess.GetVPNGateway(getVpnGatewayOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGateway failed: %s", err.Error()), "ibm_is_vpn_gateway", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, fmt.Errorf("[ERROR] Error getting Vpn Gatewa: %s\n%s", err, response)
	}
	return true, nil
}

func resourceVPNGatewayRouteFlattenHealthReasons(healthReasons []vpcv1.VPNGatewayHealthReason) (healthReasonsList []map[string]interface{}) {
	healthReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range healthReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			healthReasonsList = append(healthReasonsList, currentLR)
		}
	}
	return healthReasonsList
}

func resourceVPNGatewayFlattenLifecycleReasons(lifecycleReasons []vpcv1.VPNGatewayLifecycleReason) (lifecycleReasonsList []map[string]interface{}) {
	lifecycleReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range lifecycleReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			lifecycleReasonsList = append(lifecycleReasonsList, currentLR)
		}
	}
	return lifecycleReasonsList
}

// resourceIBMISVPNGatewayCustomizeDiff validates the zonal and regional rules
// at plan time for new gateways, so a bad combination fails before any API call.
func resourceIBMISVPNGatewayCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
	if diff.Id() != "" {
		return nil
	}
	rawConfig := diff.GetRawConfig()
	if rawConfig.IsNull() || !rawConfig.IsKnown() {
		return nil
	}
	availabilityMode := ""
	if v := rawConfig.GetAttr(isVPNGatewayAvailabilityMode); v.IsKnown() && !v.IsNull() {
		availabilityMode = v.AsString()
	}
	mode := "route"
	if v := rawConfig.GetAttr(isVPNGatewayMode); v.IsKnown() && !v.IsNull() {
		mode = v.AsString()
	}
	// An unknown value (for example a subnet created in the same apply) still counts as set.
	subnetSet := !rawConfig.GetAttr(isVPNGatewaySubnet).IsNull()
	memberCount := 0
	if v := rawConfig.GetAttr(isVPNGatewayMembers); v.IsKnown() && !v.IsNull() {
		memberCount = v.LengthInt()
	} else if !v.IsKnown() {
		return nil
	}
	return ResourceIBMISVPNGatewayValidateAvailability(availabilityMode, mode, subnetSet, memberCount)
}

// suppressVPNGatewayMemberSubnetDiff ignores member subnet changes on an
// existing gateway. The member subnet is a create-time input; after create it
// can be moved with ibm_is_vpn_gateway_member_replace, and that move must not
// show up as drift that would destroy and recreate the gateway.
func suppressVPNGatewayMemberSubnetDiff(k, old, new string, d *schema.ResourceData) bool {
	return d.Id() != ""
}

func resourceIBMISVPNGatewayMapToMemberPrototypes(membersList []interface{}) ([]vpcv1.VPNGatewayMemberPrototype, error) {
	members := make([]vpcv1.VPNGatewayMemberPrototype, 0, len(membersList))
	for i, memberIntf := range membersList {
		member, _ := memberIntf.(map[string]interface{})
		privateIPList, _ := member["private_ip"].([]interface{})
		if len(privateIPList) == 0 || privateIPList[0] == nil {
			return nil, fmt.Errorf("members.%d.private_ip.subnet is required", i)
		}
		privateIP := privateIPList[0].(map[string]interface{})
		subnetList, _ := privateIP["subnet"].([]interface{})
		if len(subnetList) == 0 || subnetList[0] == nil {
			return nil, fmt.Errorf("members.%d.private_ip.subnet is required", i)
		}
		subnetIdentity, err := vpnGatewaySubnetIdentityFromMap(subnetList[0].(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("members.%d.private_ip.%s", i, err)
		}
		members = append(members, vpcv1.VPNGatewayMemberPrototype{
			PrivateIP: &vpcv1.VPNGatewayMemberPrivateIPPrototype{
				Subnet: subnetIdentity,
			},
		})
	}
	return members, nil
}

// ResourceIBMIsVPNGatewayVPNGatewayMemberToMap flattens a VPN gateway member.
// Every pointer is checked because a member that is still pending can come back
// without private_ip or public_ip. The legacy keys (address, private_address,
// public_ip_address, private_ip_address) are kept so existing configurations
// that read them keep working.
func ResourceIBMIsVPNGatewayVPNGatewayMemberToMap(model *vpcv1.VPNGatewayMember) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range model.HealthReasons {
		healthReasonsItemMap, err := ResourceIBMIsVPNGatewayVPNGatewayMemberHealthReasonToMap(&healthReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		healthReasons = append(healthReasons, healthReasonsItemMap)
	}
	modelMap["health_reasons"] = healthReasons
	if model.HealthState != nil {
		modelMap["health_state"] = *model.HealthState
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsVPNGatewayVPNGatewayMemberLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	if model.PrivateIP != nil {
		privateIPMap, err := ResourceIBMIsVPNGatewayReservedIPReferenceVPNGatewayContextToMap(model.PrivateIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["private_ip"] = []map[string]interface{}{privateIPMap}
		if model.PrivateIP.Address != nil {
			modelMap["private_address"] = *model.PrivateIP.Address
			modelMap["private_ip_address"] = *model.PrivateIP.Address
		}
	}
	if model.PublicIP != nil {
		publicIPMap, err := ResourceIBMIsVPNGatewayIPToMap(model.PublicIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["public_ip"] = []map[string]interface{}{publicIPMap}
		if model.PublicIP.Address != nil {
			modelMap["address"] = *model.PublicIP.Address
			modelMap["public_ip_address"] = *model.PublicIP.Address
		}
	}
	if model.Role != nil {
		modelMap["role"] = *model.Role
	}
	return modelMap, nil
}

// vpnGatewayMemberMapForSchema keeps only the keys that exist in the given
// member schema, so one flattener can serve the resource and data sources.
func vpnGatewayMemberMapForSchema(model *vpcv1.VPNGatewayMember, keys map[string]*schema.Schema) (map[string]interface{}, error) {
	full, err := ResourceIBMIsVPNGatewayVPNGatewayMemberToMap(model)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{}, len(keys))
	for k, v := range full {
		if _, ok := keys[k]; ok {
			out[k] = v
		}
	}
	return out, nil
}

func ResourceIBMIsVPNGatewayVPNGatewayMemberHealthReasonToMap(model *vpcv1.VPNGatewayMemberHealthReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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

func ResourceIBMIsVPNGatewayVPNGatewayMemberLifecycleReasonToMap(model *vpcv1.VPNGatewayMemberLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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

func ResourceIBMIsVPNGatewayReservedIPReferenceVPNGatewayContextToMap(model *vpcv1.ReservedIPReferenceVPNGatewayMemberContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsVPNGatewayDeletedToMap(model.Deleted)
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
	if model.Subnet != nil {
		subnetMap, err := ResourceIBMIsVPNGatewaySubnetReferenceToMap(model.Subnet)
		if err != nil {
			return modelMap, err
		}
		modelMap["subnet"] = []map[string]interface{}{subnetMap}
	}
	return modelMap, nil
}

func ResourceIBMIsVPNGatewayDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model != nil && model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsVPNGatewaySubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsVPNGatewayDeletedToMap(model.Deleted)
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

func ResourceIBMIsVPNGatewayIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model != nil && model.Address != nil {
		modelMap["address"] = *model.Address
	}
	return modelMap, nil
}

// vpnGatewaySubnetIdentityFromMap builds a subnet identity from a nested
// subnet block. Exactly one of id, crn or href is expected.
func vpnGatewaySubnetIdentityFromMap(subnet map[string]interface{}) (vpcv1.SubnetIdentityIntf, error) {
	if id, ok := subnet["id"].(string); ok && id != "" {
		return &vpcv1.SubnetIdentityByID{ID: &id}, nil
	}
	if crn, ok := subnet["crn"].(string); ok && crn != "" {
		return &vpcv1.SubnetIdentityByCRN{CRN: &crn}, nil
	}
	if href, ok := subnet["href"].(string); ok && href != "" {
		return &vpcv1.SubnetIdentityByHref{Href: &href}, nil
	}
	return nil, fmt.Errorf("subnet must specify one of id, crn or href")
}

// ResourceIBMISVPNGatewayValidateAvailability checks the create-time rules for
// availability_mode, mode, subnet and members:
//   - regional gateways must be route based and need exactly two members and no subnet
//   - zonal gateways (the default) need subnet and must not set members
func ResourceIBMISVPNGatewayValidateAvailability(availabilityMode, mode string, subnetSet bool, memberCount int) error {
	if availabilityMode == "regional" {
		if mode != "route" {
			return fmt.Errorf("availability_mode \"regional\" is supported only for route-based VPN gateways (mode = \"route\")")
		}
		if subnetSet {
			return fmt.Errorf("subnet must not be set when availability_mode is \"regional\"; use members instead")
		}
		if memberCount != 2 {
			return fmt.Errorf("availability_mode \"regional\" requires exactly 2 members, got %d", memberCount)
		}
		return nil
	}
	if memberCount > 0 {
		return fmt.Errorf("members can be set only when availability_mode is \"regional\"")
	}
	if !subnetSet {
		return fmt.Errorf("subnet is required when availability_mode is \"zonal\" or not set")
	}
	return nil
}
