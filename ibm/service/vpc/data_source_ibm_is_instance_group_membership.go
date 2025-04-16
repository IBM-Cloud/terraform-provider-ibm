// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMISInstanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupMembershipRead,

		Schema: map[string]*schema.Schema{
			isInstanceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance group identifier.",
			},
			isInstanceGroupMembershipName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-defined name for this instance group membership. Names must be unique within the instance group.",
			},
			isInstanceGroupMemershipDeleteInstanceOnMembershipDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, when deleting the membership the instance will also be deleted.",
			},
			isInstanceGroupMembership: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this instance group membership.",
			},
			isInstanceGroupMemershipInstance: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGroupMembershipCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
						},
						isInstanceGroupMembershipVirtualServerInstance: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						isInstanceGroupMemershipInstanceName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
					},
				},
			},
			isInstanceGroupMemershipInstanceTemplate: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGroupMembershipCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this instance template.",
						},
						isInstanceGroupMemershipInstanceTemplate: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance template.",
						},
						isInstanceGroupMemershipInstanceTemplateName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this instance template.",
						},
					},
				},
			},
			isInstanceGroupMembershipLoadBalancerPoolMember: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this load balancer pool member.",
			},
			isInstanceGroupMembershipStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the instance group membership- `deleting`: Membership is deleting dependent resources- `failed`: Membership was unable to maintain dependent resources- `healthy`: Membership is active and serving in the group- `pending`: Membership is waiting for dependent resources- `unhealthy`: Membership has unhealthy dependent resources.",
			},
		},
	}
}

func dataSourceIBMISInstanceGroupMembershipRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_membership", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupID := d.Get(isInstanceGroup).(string)
	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroupMembership{}

	for {
		listInstanceGroupMembershipsOptions := vpcv1.ListInstanceGroupMembershipsOptions{
			InstanceGroupID: &instanceGroupID,
		}
		if start != "" {
			listInstanceGroupMembershipsOptions.Start = &start
		}
		instanceGroupMembershipCollection, _, err := sess.ListInstanceGroupMembershipsWithContext(context, &listInstanceGroupMembershipsOptions)
		if err != nil || instanceGroupMembershipCollection == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupMembershipsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_group_membership", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		start = flex.GetNext(instanceGroupMembershipCollection.Next)
		allrecs = append(allrecs, instanceGroupMembershipCollection.Memberships...)

		if start == "" {
			break
		}

	}

	instanceGroupMembershipName := d.Get(isInstanceGroupMembershipName).(string)
	for _, instanceGroupMembership := range allrecs {
		if instanceGroupMembershipName == *instanceGroupMembership.Name {
			d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupMembership.Instance.ID))

			if err = d.Set(isInstanceGroupMemershipDeleteInstanceOnMembershipDelete, *instanceGroupMembership.DeleteInstanceOnMembershipDelete); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting delete_instance_on_membership_delete: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-delete_instance_on_membership_delete").GetDiag()
			}

			if err = d.Set(isInstanceGroupMembership, *instanceGroupMembership.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_group_membership: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-instance_group_membership").GetDiag()
			}

			if err = d.Set(isInstanceGroupMembershipStatus, *instanceGroupMembership.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-status").GetDiag()
			}

			instances := make([]map[string]interface{}, 0)
			if instanceGroupMembership.Instance != nil {
				instance := map[string]interface{}{
					isInstanceGroupMembershipCrn:                   *instanceGroupMembership.Instance.CRN,
					isInstanceGroupMembershipVirtualServerInstance: *instanceGroupMembership.Instance.ID,
					isInstanceGroupMemershipInstanceName:           *instanceGroupMembership.Instance.Name,
				}
				instances = append(instances, instance)
			}

			if err = d.Set(isInstanceGroupMemershipInstance, instances); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-instance").GetDiag()
			}

			instance_templates := make([]map[string]interface{}, 0)
			if instanceGroupMembership.InstanceTemplate != nil {
				instance_template := map[string]interface{}{
					isInstanceGroupMembershipCrn:                 *instanceGroupMembership.InstanceTemplate.CRN,
					isInstanceGroupMemershipInstanceTemplate:     *instanceGroupMembership.InstanceTemplate.ID,
					isInstanceGroupMemershipInstanceTemplateName: *instanceGroupMembership.InstanceTemplate.Name,
				}
				instance_templates = append(instance_templates, instance_template)
			}
			if err = d.Set(isInstanceGroupMemershipInstanceTemplate, instance_templates); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_template: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-instance_template").GetDiag()
			}
			if instanceGroupMembership.PoolMember != nil && instanceGroupMembership.PoolMember.ID != nil {
				if err = d.Set(isInstanceGroupMembershipLoadBalancerPoolMember, *instanceGroupMembership.PoolMember.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting load_balancer_pool_member: %s", err), "(Data) ibm_is_instance_group_membership", "read", "set-load_balancer_pool_member").GetDiag()
				}
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance group membership %s not found", instanceGroupMembershipName), "(Data) ibm_is_instance_group_membership", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
