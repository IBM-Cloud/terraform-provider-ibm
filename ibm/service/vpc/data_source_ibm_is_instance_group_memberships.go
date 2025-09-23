// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isInstanceGroupMemberships = "memberships"
)

func DataSourceIBMISInstanceGroupMemberships() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupMembershipsRead,

		Schema: map[string]*schema.Schema{
			isInstanceGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance group identifier.",
			},

			isInstanceGroupMemberships: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of instance group memberships.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						isInstanceGroupMembershipName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this instance group membership. Names must be unique within the instance group.",
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
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupMembershipsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_memberships", "read", "initialize-client")
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
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupMembershipsWithContext failed %s", err), "(Data) ibm_is_instance_group_memberships", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		start = flex.GetNext(instanceGroupMembershipCollection.Next)
		allrecs = append(allrecs, instanceGroupMembershipCollection.Memberships...)

		if start == "" {
			break
		}

	}

	memberships := make([]map[string]interface{}, 0)
	for _, instanceGroupMembership := range allrecs {
		membership := map[string]interface{}{
			isInstanceGroupMemershipDeleteInstanceOnMembershipDelete: *instanceGroupMembership.DeleteInstanceOnMembershipDelete,
			isInstanceGroupMembership:                                *instanceGroupMembership.ID,
			isInstanceGroupMembershipName:                            *instanceGroupMembership.Name,
			isInstanceGroupMembershipStatus:                          *instanceGroupMembership.Status,
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
		membership[isInstanceGroupMemershipInstance] = instances

		instance_templates := make([]map[string]interface{}, 0)
		if instanceGroupMembership.InstanceTemplate != nil {
			instance_template := map[string]interface{}{
				isInstanceGroupMembershipCrn:                 *instanceGroupMembership.InstanceTemplate.CRN,
				isInstanceGroupMemershipInstanceTemplate:     *instanceGroupMembership.InstanceTemplate.ID,
				isInstanceGroupMemershipInstanceTemplateName: *instanceGroupMembership.InstanceTemplate.Name,
			}
			instance_templates = append(instance_templates, instance_template)
		}
		membership[isInstanceGroupMemershipInstanceTemplate] = instance_templates

		if instanceGroupMembership.PoolMember != nil && instanceGroupMembership.PoolMember.ID != nil {
			membership[isInstanceGroupMembershipLoadBalancerPoolMember] = *instanceGroupMembership.PoolMember.ID
		}

		memberships = append(memberships, membership)
	}
	if err = d.Set("memberships", memberships); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting memberships %s", err), "(Data) ibm_is_instance_group_memberships", "read", "memberships-set").GetDiag()
	}
	d.SetId(dataSourceIbmIsInstanceGroupMembershipsID(d))

	return nil
}

// dataSourceIbmIsInstanceGroupMembershipsID returns a reasonable ID for the list.
func dataSourceIbmIsInstanceGroupMembershipsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
