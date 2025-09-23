// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIBLBPoolMember() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsLbPoolMemberRead,

		Schema: map[string]*schema.Schema{
			"lb": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The load balancer identifier.",
			},
			"pool": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The pool identifier.",
			},
			"member": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The member identifier.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this member was created.",
			},
			"health": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Health of the server member in the pool.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The member's canonical URL.",
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The port number of the application running in the server member.",
			},
			"provisioning_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The provisioning status of this member.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The pool member target. Load balancers in the `network` family support virtual serverinstances. Load balancers in the `application` family support IP addresses.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
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
							Description: "The URL for this virtual server instance.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
						},
					},
				},
			},
			"weight": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Weight of the server member. Applicable only if the pool algorithm is`weighted_round_robin`.",
			},
		},
	}
}

func dataSourceIBMIsLbPoolMemberRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_lb_pool_member", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getLoadBalancerPoolMemberOptions := &vpcv1.GetLoadBalancerPoolMemberOptions{}

	getLoadBalancerPoolMemberOptions.SetLoadBalancerID(d.Get("lb").(string))
	getLoadBalancerPoolMemberOptions.SetPoolID(d.Get("pool").(string))
	getLoadBalancerPoolMemberOptions.SetID(d.Get("member").(string))

	loadBalancerPoolMember, _, err := sess.GetLoadBalancerPoolMemberWithContext(context, getLoadBalancerPoolMemberOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "(Data) ibm_is_lb_pool_member", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*loadBalancerPoolMember.ID)
	if err = d.Set("created_at", flex.DateTimeToString(loadBalancerPoolMember.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("health", loadBalancerPoolMember.Health); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-health").GetDiag()
	}
	if err = d.Set("href", loadBalancerPoolMember.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-href").GetDiag()
	}
	if err = d.Set("port", flex.IntValue(loadBalancerPoolMember.Port)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting port: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-port").GetDiag()
	}
	if err = d.Set("provisioning_status", loadBalancerPoolMember.ProvisioningStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provisioning_status: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-provisioning_status").GetDiag()
	}

	if loadBalancerPoolMember.Target != nil {
		target := loadBalancerPoolMember.Target.(*vpcv1.LoadBalancerPoolMemberTarget)
		err = d.Set("target", dataSourceLoadBalancerPoolMemberFlattenTarget(*target))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-target").GetDiag()
		}
	}
	if !core.IsNil(loadBalancerPoolMember.Weight) {
		if err = d.Set("weight", flex.IntValue(loadBalancerPoolMember.Weight)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting weight: %s", err), "(Data) ibm_is_lb_pool_member", "read", "set-weight").GetDiag()
		}
	}

	return nil
}

func dataSourceLoadBalancerPoolMemberFlattenTarget(result vpcv1.LoadBalancerPoolMemberTarget) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceLoadBalancerPoolMemberTargetToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceLoadBalancerPoolMemberTargetToMap(targetItem vpcv1.LoadBalancerPoolMemberTarget) (targetMap map[string]interface{}) {
	targetMap = map[string]interface{}{}

	if targetItem.CRN != nil {
		targetMap["crn"] = targetItem.CRN
	}
	if targetItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceLoadBalancerPoolMemberTargetDeletedToMap(*targetItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		targetMap["deleted"] = deletedList
	}
	if targetItem.Href != nil {
		targetMap["href"] = targetItem.Href
	}
	if targetItem.ID != nil {
		targetMap["id"] = targetItem.ID
	}
	if targetItem.Name != nil {
		targetMap["name"] = targetItem.Name
	}
	if targetItem.Address != nil {
		targetMap["address"] = targetItem.Address
	}

	return targetMap
}

func dataSourceLoadBalancerPoolMemberTargetDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}
