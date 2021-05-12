// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func dataSourceIBMISInstanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceGroupMembershipRead,

		Schema: map[string]*schema.Schema{
			"instance_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance group identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-defined name for this instance group membership. Names must be unique within the instance group.",
			},
			"delete_instance_on_membership_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, when deleting the membership the instance will also be deleted.",
			},
			"instance_group_membership": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this instance group membership.",
			},
			"instance": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
						},
						"virtual_server_instance": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
					},
				},
			},
			"instance_template": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this instance template.",
						},
						"instance_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance template.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this instance template.",
						},
					},
				},
			},
			"load_balancer_pool_member": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this load balancer pool member.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the instance group membership- `deleting`: Membership is deleting dependent resources- `failed`: Membership was unable to maintain dependent resources- `healthy`: Membership is active and serving in the group- `pending`: Membership is waiting for dependent resources- `unhealthy`: Membership has unhealthy dependent resources.",
			},
		},
	}
}

func dataSourceIBMISInstanceGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	instanceGroupID := d.Get("instance_group").(string)
	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroupMembership{}

	for {
		listInstanceGroupMembershipsOptions := vpcv1.ListInstanceGroupMembershipsOptions{
			InstanceGroupID: &instanceGroupID,
		}
		instanceGroupMembershipCollection, response, err := sess.ListInstanceGroupMemberships(&listInstanceGroupMembershipsOptions)
		if err != nil || instanceGroupMembershipCollection == nil {
			return fmt.Errorf("Error Getting InstanceGroup Membership Collection %s\n%s", err, response)
		}

		start = GetNext(instanceGroupMembershipCollection.Next)
		allrecs = append(allrecs, instanceGroupMembershipCollection.Memberships...)

		if start == "" {
			break
		}

	}

	instanceGroupMembershipName := d.Get("name").(string)
	for _, instanceGroupMembership := range allrecs {
		if instanceGroupMembershipName == *instanceGroupMembership.Name {
			d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupMembership.Instance.ID))
			d.Set("delete_instance_on_membership_delete", *instanceGroupMembership.DeleteInstanceOnMembershipDelete)
			d.Set("instance_group_membership", *instanceGroupMembership.ID)
			d.Set("status", *instanceGroupMembership.Status)

			instances := make([]map[string]interface{}, 0)
			if instanceGroupMembership.Instance != nil {
				instance := map[string]interface{}{
					"crn":                     *instanceGroupMembership.Instance.CRN,
					"virtual_server_instance": *instanceGroupMembership.Instance.ID,
					"name":                    *instanceGroupMembership.Instance.Name,
				}
				instances = append(instances, instance)
			}
			d.Set("instance", instances)

			instance_templates := make([]map[string]interface{}, 0)
			if instanceGroupMembership.InstanceTemplate != nil {
				instance_template := map[string]interface{}{
					"crn":               *instanceGroupMembership.InstanceTemplate.CRN,
					"instance_template": *instanceGroupMembership.InstanceTemplate.ID,
					"name":              *instanceGroupMembership.InstanceTemplate.Name,
				}
				instance_templates = append(instance_templates, instance_template)
			}
			d.Set("instance_template", instance_templates)

			if instanceGroupMembership.PoolMember != nil && instanceGroupMembership.PoolMember.ID != nil {
				d.Set("load_balancer_pool_member", *instanceGroupMembership.PoolMember.ID)
			}
			return nil
		}
	}
	return fmt.Errorf("Instance group membership %s not found", instanceGroupMembershipName)
}
