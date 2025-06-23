// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISLbProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISLbProfileRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for this load balancer profile",
			},
			isLBAccessModes: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The access mode for a load balancer with this profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for access mode",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access modes for this profile",
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Access modes for this profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"targetable_load_balancer_profiles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The load balancer profiles that load balancers with this profile can target",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The product family this load balancer profile belongs to",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this load balancer profile",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this load balancer profile",
						},
					},
				},
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this load balancer profile",
			},
			"family": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The product family this load balancer profile belongs to",
			},
			"route_mode_supported": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The route mode support for a load balancer with this profile depends on its configuration",
			},
			"route_mode_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The route mode type for this load balancer profile, one of [fixed, dependent]",
			},
			"udp_supported": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The UDP support for a load balancer with this profile",
			},
			"udp_supported_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UDP support type for a load balancer with this profile",
			},
			"failsafe_policy_actions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default failsafe policy action for this profile.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The supported failsafe policy actions.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISLbProfileRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_profile", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	lbprofilename := d.Get(isLbsProfileName).(string)
	getLoadBalancerProfileOptions := &vpcv1.GetLoadBalancerProfileOptions{
		Name: &lbprofilename,
	}
	loadBalancerProfile, _, err := sess.GetLoadBalancerProfileWithContext(context, getLoadBalancerProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerProfileWithContext failed: %s", err.Error()), "(Data) ibm_is_lb_profile", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", loadBalancerProfile.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_lb_profile", "read", "set-name").GetDiag()
	}

	if err = d.Set("href", loadBalancerProfile.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_lb_profile", "read", "set-href").GetDiag()
	}
	if err = d.Set("family", loadBalancerProfile.Family); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting family: %s", err), "(Data) ibm_is_lb_profile", "read", "set-family").GetDiag()
	}
	if loadBalancerProfile.AccessModes != nil {
		accessModes := loadBalancerProfile.AccessModes
		AccessModesMap := map[string]interface{}{}
		AccessModesList := []map[string]interface{}{}
		if accessModes.Type != nil {
			AccessModesMap["type"] = *accessModes.Type
		}
		if len(accessModes.Values) > 0 {
			AccessModesMap["values"] = accessModes.Values
		}
		AccessModesList = append(AccessModesList, AccessModesMap)
		if err = d.Set("access_modes", AccessModesList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_modes: %s", err), "(Data) ibm_is_lb_profile", "read", "set-access_modes").GetDiag()
		}
	}

	if loadBalancerProfile.TargetableLoadBalancerProfiles != nil {
		d.Set("targetable_load_balancer_profiles", dataSourceLbProfileFlattenTargetableLoadBalancerProfiles(loadBalancerProfile.TargetableLoadBalancerProfiles))
	}
	log.Printf("[INFO] loadBalancerProfile udp %v", loadBalancerProfile.UDPSupported)
	if loadBalancerProfile.UDPSupported != nil {
		udpSupport := loadBalancerProfile.UDPSupported
		log.Printf("[INFO] loadBalancerProfile udp %s", reflect.TypeOf(udpSupport).String())

		switch reflect.TypeOf(udpSupport).String() {
		case "*vpcv1.LoadBalancerProfileUDPSupportedFixed":
			{
				udp := udpSupport.(*vpcv1.LoadBalancerProfileUDPSupportedFixed)
				if err = d.Set("udp_supported", udp.Value); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported: %s", err), "(Data) ibm_is_lb_profile", "read", "set-udp_supported").GetDiag()
				}
				if err = d.Set("udp_supported_type", udp.Type); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-udp_supported_type").GetDiag()
				}
			}
		case "*vpcv1.LoadBalancerProfileUDPSupportedDependent":
			{
				udp := udpSupport.(*vpcv1.LoadBalancerProfileUDPSupportedDependent)
				if udp.Type != nil {
					if err = d.Set("udp_supported_type", *udp.Type); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-udp_supported_type").GetDiag()
					}
				}
			}
		case "*vpcv1.LoadBalancerProfileUDPSupported":
			{
				udp := udpSupport.(*vpcv1.LoadBalancerProfileUDPSupported)
				if udp.Type != nil {
					if err = d.Set("udp_supported_type", *udp.Type); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-udp_supported_type").GetDiag()
					}
				}
				if udp.Value != nil {
					if err = d.Set("udp_supported", *udp.Value); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp_supported: %s", err), "(Data) ibm_is_lb_profile", "read", "set-udp_supported").GetDiag()
					}
				}
			}
		}
	}

	failsafePolicyActions := []map[string]interface{}{}
	if loadBalancerProfile.FailsafePolicyActions != nil {
		modelMap, err := dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsToMap(loadBalancerProfile.FailsafePolicyActions)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_profile", "read", "failsafe_policy_actions-to-map").GetDiag()
		}
		failsafePolicyActions = append(failsafePolicyActions, modelMap)
	}
	if err = d.Set("failsafe_policy_actions", failsafePolicyActions); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting failsafe_policy_actions: %s", err), "(Data) ibm_is_lb_profile", "read", "set-failsafe_policy_actions").GetDiag()
	}
	if loadBalancerProfile.RouteModeSupported != nil {
		routeMode := loadBalancerProfile.RouteModeSupported
		switch reflect.TypeOf(routeMode).String() {
		case "*vpcv1.LoadBalancerProfileRouteModeSupportedFixed":
			{
				rms := routeMode.(*vpcv1.LoadBalancerProfileRouteModeSupportedFixed)
				if err = d.Set("route_mode_supported", rms.Value); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode_supported: %s", err), "(Data) ibm_is_lb_profile", "read", "set-route_mode_supported").GetDiag()
				}
				if err = d.Set("route_mode_type", rms.Type); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-route_mode_type").GetDiag()
				}
			}
		case "*vpcv1.LoadBalancerProfileRouteModeSupportedDependent":
			{
				rms := routeMode.(*vpcv1.LoadBalancerProfileRouteModeSupportedDependent)
				if rms.Type != nil {
					if err = d.Set("route_mode_type", *rms.Type); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-route_mode_type").GetDiag()
					}
				}
			}
		case "*vpcv1.LoadBalancerProfileRouteModeSupported":
			{
				rms := routeMode.(*vpcv1.LoadBalancerProfileRouteModeSupported)
				if rms.Type != nil {
					if err = d.Set("route_mode_type", *rms.Type); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode_type: %s", err), "(Data) ibm_is_lb_profile", "read", "set-route_mode_type").GetDiag()
					}
				}
				if rms.Value != nil {
					if err = d.Set("route_mode_supported", *rms.Value); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_mode_supported: %s", err), "(Data) ibm_is_lb_profile", "read", "set-route_mode_supported").GetDiag()
					}
				}
			}
		}
	}
	d.SetId(*loadBalancerProfile.Name)
	return nil
}

func dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsToMap(model vpcv1.LoadBalancerProfileFailsafePolicyActionsIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.LoadBalancerProfileFailsafePolicyActionsEnum); ok {
		return dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsEnumToMap(model.(*vpcv1.LoadBalancerProfileFailsafePolicyActionsEnum))
	} else if _, ok := model.(*vpcv1.LoadBalancerProfileFailsafePolicyActionsDependent); ok {
		return dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsDependentToMap(model.(*vpcv1.LoadBalancerProfileFailsafePolicyActionsDependent))
	} else if _, ok := model.(*vpcv1.LoadBalancerProfileFailsafePolicyActions); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.LoadBalancerProfileFailsafePolicyActions)
		if model.Default != nil {
			modelMap["default"] = *model.Default
		}
		if model.Type != nil {
			modelMap["type"] = *model.Type
		}
		if model.Values != nil {
			modelMap["values"] = model.Values
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.LoadBalancerProfileFailsafePolicyActionsIntf subtype encountered")
	}
}

func dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsEnumToMap(model *vpcv1.LoadBalancerProfileFailsafePolicyActionsEnum) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["default"] = *model.Default
	modelMap["type"] = *model.Type
	modelMap["values"] = model.Values
	return modelMap, nil
}

func dataSourceIBMIsLbProfileLoadBalancerProfileFailsafePolicyActionsDependentToMap(model *vpcv1.LoadBalancerProfileFailsafePolicyActionsDependent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func dataSourceLbProfileFlattenTargetableLoadBalancerProfiles(result []vpcv1.LoadBalancerProfileReference) (targetableLoadBalancerProfiles []map[string]interface{}) {
	for _, targetableLoadBalancerProfileItem := range result {
		targetableLoadBalancerProfiles = append(targetableLoadBalancerProfiles, dataSourceLbProfileTargetableLoadBalancerProfilesToMap(targetableLoadBalancerProfileItem))
	}

	return targetableLoadBalancerProfiles
}

func dataSourceLbProfileTargetableLoadBalancerProfilesToMap(targetableLoadBalancerProfileItem vpcv1.LoadBalancerProfileReference) (targetableLoadBalancerProfileMap map[string]interface{}) {
	targetableLoadBalancerProfileMap = map[string]interface{}{}

	if targetableLoadBalancerProfileItem.Family != nil {
		targetableLoadBalancerProfileMap["family"] = targetableLoadBalancerProfileItem.Family
	}

	if targetableLoadBalancerProfileItem.Href != nil {
		targetableLoadBalancerProfileMap["href"] = targetableLoadBalancerProfileItem.Href
	}

	if targetableLoadBalancerProfileItem.Name != nil {
		targetableLoadBalancerProfileMap["name"] = targetableLoadBalancerProfileItem.Name
	}

	return targetableLoadBalancerProfileMap
}
