// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func ResourceIBMIsVPCDnsResolutionBinding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVPCDnsResolutionBindingCreate,
		ReadContext:   resourceIBMIsVPCDnsResolutionBindingRead,
		UpdateContext: resourceIBMIsVPCDnsResolutionBindingUpdate,
		DeleteContext: resourceIBMIsVPCDnsResolutionBindingDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"vpc_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPC identifier.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the DNS resolution binding was created.",
			},
			"health_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `health_state` (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.",
			},
			"endpoint_gateways": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The endpoint gateways in the bound to VPC that are allowed to participate in this DNS resolution binding.The endpoint gateways may be remote and therefore may not be directly retrievable.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this endpoint gateway.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this endpoint gateway.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this endpoint gateway.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this account.",
												},
												"resource_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource type.",
												},
											},
										},
									},
									"region": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this region.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The globally unique name for this region.",
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
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this DNS resolution binding.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the DNS resolution binding.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name for this DNS resolution binding. The name is unique across all DNS resolution bindings for the VPC.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"vpc": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Description: "The VPC bound to for DNS resolution.The VPC may be remote and therefore may not be directly retrievable.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Optional:     true,
							ExactlyOneOf: []string{"vpc.0.id", "vpc.0.href", "vpc.0.crn"},
							Type:         schema.TypeString,
							Computed:     true,
							Description:  "The CRN for this VPC.",
						},
						"href": &schema.Schema{
							Type:         schema.TypeString,
							ExactlyOneOf: []string{"vpc.0.id", "vpc.0.href", "vpc.0.crn"},
							Optional:     true,
							Computed:     true,
							Description:  "The URL for this VPC.",
						},
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ExactlyOneOf: []string{"vpc.0.id", "vpc.0.href", "vpc.0.crn"},
							Optional:     true,
							Computed:     true,
							Description:  "The unique identifier for this VPC.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this VPC. The name is unique across all VPCs in the region.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this account.",
												},
												"resource_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource type.",
												},
											},
										},
									},
									"region": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this region.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The globally unique name for this region.",
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
					},
				},
			},
		},
	}
}

func resourceIBMIsVPCDnsResolutionBindingCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	spokeVPCID := d.Get("vpc_id").(string)
	createVPCDnsResolutionBindingOptions := &vpcv1.CreateVPCDnsResolutionBindingOptions{}
	vpchref := d.Get("vpc.0.href").(string)
	vpccrn := d.Get("vpc.0.crn").(string)
	vpcid := d.Get("vpc.0.id").(string)

	createVPCDnsResolutionBindingOptions.SetVPCID(spokeVPCID)
	if d.Get("name").(string) != "" {
		createVPCDnsResolutionBindingOptions.SetName(d.Get("name").(string))
	}
	if vpchref != "" {
		vPCIdentityIntf := &vpcv1.VPCIdentityByHref{
			Href: &vpchref,
		}
		createVPCDnsResolutionBindingOptions.SetVPC(vPCIdentityIntf)
	} else if vpcid != "" {
		vPCIdentityIntf := &vpcv1.VPCIdentityByID{
			ID: &vpcid,
		}
		createVPCDnsResolutionBindingOptions.SetVPC(vPCIdentityIntf)
	} else {
		vPCIdentityIntf := &vpcv1.VPCIdentityByCRN{
			CRN: &vpccrn,
		}
		createVPCDnsResolutionBindingOptions.SetVPC(vPCIdentityIntf)
	}
	vpcdnsResolutionBinding, _, err := sess.CreateVPCDnsResolutionBindingWithContext(context, createVPCDnsResolutionBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPCDnsResolutionBindingWithContext failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(MakeTerraformVPCDNSID(spokeVPCID, *vpcdnsResolutionBinding.ID))
	intf, err := isWaitForVpcDnsCreated(sess, spokeVPCID, *vpcdnsResolutionBinding.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	if vpcdnsResolutionBinding, ok := intf.(*vpcv1.VpcdnsResolutionBinding); ok {
		diagErr := resourceIBMIsVPCDnsResolutionBindingGet(vpcdnsResolutionBinding, d)
		if diagErr != nil {
			return diagErr
		}
	} else {
		return resourceIBMIsVPCDnsResolutionBindingRead(context, d, meta)
	}

	return nil
}
func resourceIBMIsVPCDnsResolutionBindingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpcId, id, err := ParseVPCDNSTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "sep-id-parts").GetDiag()
	}
	getVPCDnsResolutionBindingOptions := &vpcv1.GetVPCDnsResolutionBindingOptions{}

	getVPCDnsResolutionBindingOptions.SetVPCID(vpcId)
	getVPCDnsResolutionBindingOptions.SetID(id)

	vpcdnsResolutionBinding, response, err := sess.GetVPCDnsResolutionBindingWithContext(context, getVPCDnsResolutionBindingOptions)
	if err != nil {
		log.Printf("[DEBUG] GetVPCDnsResolutionBindingWithContext failed %s\n%s", err, response)
		if response.StatusCode != 404 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCDnsResolutionBindingWithContext failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		} else {
			d.SetId("")
			return nil
		}
	}
	diagErr := resourceIBMIsVPCDnsResolutionBindingGet(vpcdnsResolutionBinding, d)
	if diagErr != nil {
		return diagErr
	}
	return nil
}
func resourceIBMIsVPCDnsResolutionBindingGet(vpcdnsResolutionBinding *vpcv1.VpcdnsResolutionBinding, d *schema.ResourceData) diag.Diagnostics {
	var err error
	if err = d.Set("created_at", flex.DateTimeToString(vpcdnsResolutionBinding.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-created_at").GetDiag()
	}

	endpointGateways := []map[string]interface{}{}
	if vpcdnsResolutionBinding.EndpointGateways != nil {
		for _, modelItem := range vpcdnsResolutionBinding.EndpointGateways {
			modelMap, err := dataSourceIBMIsVPCDnsResolutionBindingEndpointGatewayReferenceRemoteToMap(&modelItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "endpoint_gateways-to-map").GetDiag()
			}
			endpointGateways = append(endpointGateways, modelMap)
		}
	}
	if err := d.Set("endpoint_gateways", endpointGateways); err != nil {
		err = fmt.Errorf("Error setting endpoint_gateways: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-endpoint_gateways").GetDiag()
	}

	if err := d.Set("href", vpcdnsResolutionBinding.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-href").GetDiag()
	}

	if err := d.Set("lifecycle_state", vpcdnsResolutionBinding.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-lifecycle_state").GetDiag()
	}
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range vpcdnsResolutionBinding.HealthReasons {
		healthReasonsItemMap, err := resourceIBMIsVPCDnsResolutionBindingVpcdnsResolutionBindingHealthReasonToMap(&healthReasonsItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "health_reasons-to-map").GetDiag()
		}
		healthReasons = append(healthReasons, healthReasonsItemMap)
	}
	if err := d.Set("health_reasons", healthReasons); err != nil {
		err = fmt.Errorf("Error setting health_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-health_reasons").GetDiag()
	}
	if err := d.Set("health_state", vpcdnsResolutionBinding.HealthState); err != nil {
		err = fmt.Errorf("Error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-health_state").GetDiag()
	}
	if err := d.Set("name", vpcdnsResolutionBinding.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-name").GetDiag()
	}

	if err := d.Set("resource_type", vpcdnsResolutionBinding.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-resource_type").GetDiag()
	}

	vpc := []map[string]interface{}{}
	if vpcdnsResolutionBinding.VPC != nil {
		modelMap, err := dataSourceIBMIsVPCDnsResolutionBindingVPCReferenceRemoteToMap(vpcdnsResolutionBinding.VPC)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "vpc-to-map").GetDiag()
		}
		vpc = append(vpc, modelMap)
	}
	if err := d.Set("vpc", vpc); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "read", "set-vpc").GetDiag()
	}
	return nil
}
func resourceIBMIsVPCDnsResolutionBindingUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpcId, id, err := ParseVPCDNSTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "update", "sep-id-parts").GetDiag()
	}
	if d.HasChange(isVPCDnsResolutionBindingName) {
		nameChange := d.Get(isVPCDnsResolutionBindingName).(string)
		vpcdnsResolutionBindingPatch := &vpcv1.VpcdnsResolutionBindingPatch{
			Name: &nameChange,
		}
		vpcdnsResolutionBindingPatchAsPatch, err := vpcdnsResolutionBindingPatch.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf(" vpcdnsResolutionBindingPatch.AsPatch() failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateVPCDnsResolutionBindingOptions := &vpcv1.UpdateVPCDnsResolutionBindingOptions{}

		updateVPCDnsResolutionBindingOptions.SetVPCID(vpcId)
		updateVPCDnsResolutionBindingOptions.SetID(id)
		updateVPCDnsResolutionBindingOptions.SetVpcdnsResolutionBindingPatch(vpcdnsResolutionBindingPatchAsPatch)

		vpcdnsResolutionBinding, _, err := sess.UpdateVPCDnsResolutionBindingWithContext(context, updateVPCDnsResolutionBindingOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPCDnsResolutionBindingWithContext failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		diagErr := resourceIBMIsVPCDnsResolutionBindingGet(vpcdnsResolutionBinding, d)
		if diagErr != nil {
			return diagErr
		}
	}

	return nil
}
func resourceIBMIsVPCDnsResolutionBindingDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpcId, id, err := ParseVPCDNSTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_dns_resolution_binding", "delete", "sep-id-parts").GetDiag()
	}
	deleteVPCDnsResolutionBindingOptions := &vpcv1.DeleteVPCDnsResolutionBindingOptions{}

	deleteVPCDnsResolutionBindingOptions.SetVPCID(vpcId)
	deleteVPCDnsResolutionBindingOptions.SetID(id)

	dns, _, err := sess.DeleteVPCDnsResolutionBindingWithContext(context, deleteVPCDnsResolutionBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVPCDnsResolutionBindingWithContext failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForVpcDnsDeleted(sess, vpcId, id, d.Timeout(schema.TimeoutDelete), dns)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpcDnsDeleted failed: %s", err.Error()), "ibm_is_vpc_dns_resolution_binding", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}
func MakeTerraformVPCDNSID(id1, id2 string) string {
	// Include both  vpc id and binding id to create a unique Terraform id.  As a bonus,
	// we can extract the bindings as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}

func ParseVPCDNSTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

func isWaitForVpcDnsDeleted(sess *vpcv1.VpcV1, vpcid, id string, timeout time.Duration, dns *vpcv1.VpcdnsResolutionBinding) (interface{}, error) {
	log.Printf("Waiting for vpc dns (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "pending", "updating", "waiting"},
		Target:     []string{"stable", "failed", "suspended", ""},
		Refresh:    isVpcDnsDeleteRefreshFunc(sess, vpcid, id, dns),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpcDnsDeleteRefreshFunc(sess *vpcv1.VpcV1, vpcid, id string, dns *vpcv1.VpcdnsResolutionBinding) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVPCDnsResolutionBindingOptions := &vpcv1.GetVPCDnsResolutionBindingOptions{}

		getVPCDnsResolutionBindingOptions.SetVPCID(vpcid)
		getVPCDnsResolutionBindingOptions.SetID(id)

		vpcdnsResolutionBinding, response, err := sess.GetVPCDnsResolutionBinding(getVPCDnsResolutionBindingOptions)
		if vpcdnsResolutionBinding == nil {
			vpcdnsResolutionBinding = dns
		}
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vpcdnsResolutionBinding, "", nil
			}
			return vpcdnsResolutionBinding, "", fmt.Errorf("[ERROR] Error getting vpcdnsResolutionBinding: %s\n%s", err, response)
		}
		return vpcdnsResolutionBinding, *vpcdnsResolutionBinding.LifecycleState, err
	}
}

func isWaitForVpcDnsCreated(sess *vpcv1.VpcV1, vpcid, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for vpc dns (%s) to be created.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "pending", "updating", "waiting"},
		Target:     []string{"stable", "failed", "suspended", ""},
		Refresh:    isVpcDnsCreateRefreshFunc(sess, vpcid, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVpcDnsCreateRefreshFunc(sess *vpcv1.VpcV1, vpcid, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVPCDnsResolutionBindingOptions := &vpcv1.GetVPCDnsResolutionBindingOptions{}

		getVPCDnsResolutionBindingOptions.SetVPCID(vpcid)
		getVPCDnsResolutionBindingOptions.SetID(id)

		vpcdnsResolutionBinding, response, err := sess.GetVPCDnsResolutionBinding(getVPCDnsResolutionBindingOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vpcdnsResolutionBinding, "", nil
			}
			return vpcdnsResolutionBinding, "", fmt.Errorf("[ERROR] Error getting vpcdnsResolutionBinding: %s\n%s", err, response)
		}
		if *vpcdnsResolutionBinding.LifecycleState == "failed" || *vpcdnsResolutionBinding.LifecycleState == "suspended" {
			return vpcdnsResolutionBinding, "", fmt.Errorf("[ERROR] DnsResolutionBinding in %s state", *vpcdnsResolutionBinding.LifecycleState)
		}
		return vpcdnsResolutionBinding, *vpcdnsResolutionBinding.LifecycleState, err
	}
}

func resourceIBMIsVPCDnsResolutionBindingVpcdnsResolutionBindingHealthReasonToMap(model *vpcv1.VpcdnsResolutionBindingHealthReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = model.Code
	modelMap["message"] = model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = model.MoreInfo
	}
	return modelMap, nil
}
