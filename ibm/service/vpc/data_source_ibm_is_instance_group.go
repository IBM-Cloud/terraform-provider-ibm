// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISInstanceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-defined name for this instance group",
			},

			"instance_template": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "instance template ID",
			},

			"membership_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of instances in the instance group",
			},

			"resource_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group ID",
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this instance group",
			},

			"subnets": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of subnet IDs",
			},

			"application_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Used by the instance group when scaling up instances to supply the port for the load balancer pool member.",
			},

			"load_balancer_pool": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "load balancer pool ID",
			},

			"managers": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of Managers associated with instancegroup",
			},

			"vpc": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "vpc instance",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance group status - deleting, healthy, scaling, unhealthy",
			},

			isInstanceGroupAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},
		},
	}
}

func dataSourceIBMISInstanceGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	name := d.Get("name")

	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroup{}
	for {
		listInstanceGroupOptions := vpcv1.ListInstanceGroupsOptions{}
		if start != "" {
			listInstanceGroupOptions.Start = &start
		}
		instanceGroupsCollection, _, err := sess.ListInstanceGroupsWithContext(context, &listInstanceGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_group", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(instanceGroupsCollection.Next)
		allrecs = append(allrecs, instanceGroupsCollection.InstanceGroups...)

		if start == "" {
			break
		}

	}

	for _, instanceGroup := range allrecs {
		if *instanceGroup.Name == name {
			if err = d.Set("name", *instanceGroup.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_group", "read", "set-name").GetDiag()
			}
			if !core.IsNil(instanceGroup.InstanceTemplate) {
				if err = d.Set("instance_template", *instanceGroup.InstanceTemplate.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_template: %s", err), "(Data) ibm_is_instance_group", "read", "set-instance_template").GetDiag()
				}
			}
			if err = d.Set("membership_count", flex.IntValue(instanceGroup.MembershipCount)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting membership_count: %s", err), "(Data) ibm_is_instance_group", "read", "set-membership_count").GetDiag()
			}
			if !core.IsNil(instanceGroup.ResourceGroup) {
				if err = d.Set("resource_group", *instanceGroup.ResourceGroup.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance_group", "read", "set-resource_group").GetDiag()
				}
			}
			if err = d.Set("crn", instanceGroup.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance_group", "read", "set-crn").GetDiag()
			}
			d.SetId(*instanceGroup.ID)
			if !core.IsNil(instanceGroup.ApplicationPort) {
				if err = d.Set("application_port", flex.IntValue(instanceGroup.ApplicationPort)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting application_port: %s", err), "(Data) ibm_is_instance_group", "read", "set-application_port").GetDiag()
				}
			}
			subnets := make([]string, 0)
			for i := 0; i < len(instanceGroup.Subnets); i++ {
				subnets = append(subnets, string(*(instanceGroup.Subnets[i].ID)))
			}
			if instanceGroup.LoadBalancerPool != nil {
				if err = d.Set("load_balancer_pool", *instanceGroup.LoadBalancerPool.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting load_balancer_pool: %s", err), "(Data) ibm_is_instance_group", "read", "set-load_balancer_pool").GetDiag()
				}
			}
			if err = d.Set("subnets", subnets); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnets: %s", err), "(Data) ibm_is_instance_group", "read", "set-subnets").GetDiag()
			}

			managers := make([]string, 0)
			for i := 0; i < len(instanceGroup.Managers); i++ {
				managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
			}
			if err = d.Set("managers", managers); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managers: %s", err), "(Data) ibm_is_instance_group", "read", "set-managers").GetDiag()
			}
			if !core.IsNil(instanceGroup.VPC) {
				if err = d.Set("vpc", *instanceGroup.VPC.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance_group", "read", "set-vpc").GetDiag()
				}
			}
			if err = d.Set("status", instanceGroup.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_instance_group", "read", "set-status").GetDiag()
			}
			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *instanceGroup.CRN, "", isInstanceGroupAccessTagType)
			if err != nil {
				log.Printf(
					"[ERROR] Error occured during reading of instance group (%s) access tags: %s", d.Id(), err)
			}
			if err = d.Set(isInstanceGroupAccessTags, accesstags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_instance_group", "read", "set-access_tags").GetDiag()
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance group %s not found", name), "(Data) ibm_is_instance_group", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
