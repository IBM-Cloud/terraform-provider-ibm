// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsFlowLog() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsFlowLogRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identifier", "name"},
				Description:  "The unique user-defined name for this flow log collector.",
			},

			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identifier", "name"},
				Description:  "The flow log collector identifier.",
			},
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this collector is active.",
			},
			"auto_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to `true`, this flow log collector will be automatically deleted when the target is deleted.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the flow log collector was created.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this flow log collector.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this flow log collector.",
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the flow log collector.",
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this flow log collector.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			"storage_bucket": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Cloud Object Storage bucket where the collected flows are logged.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name of this COS bucket.",
						},
					},
				},
			},
			"target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target this collector is collecting flow logs for. If the target is an instance,subnet, or VPC, flow logs will not be collected for any network interfaces within thetarget that are themselves the target of a more specific flow log collector.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "The URL for this network interface.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this network interface.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this network interface.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
						},
					},
				},
			},
			"vpc": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPC this flow log collector is associated with.",
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
			isFlowLogAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func dataSourceIBMIsFlowLogRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_ibm_is_flow_log", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	name := d.Get("name").(string)
	identifier := d.Get("identifier").(string)
	var flowLogCollector *vpcv1.FlowLogCollector

	if name != "" {

		listOptions := &vpcv1.ListFlowLogCollectorsOptions{
			Name: &name,
		}

		flowlogCollectors, _, err := vpcClient.ListFlowLogCollectorsWithContext(context, listOptions) // Use WithContext
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListFlowLogCollectorsWithContext failed: %s", err.Error()), "(Data) ibm_ibm_is_flow_log", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		allrecs := flowlogCollectors.FlowLogCollectors

		if len(allrecs) == 0 {
			err = fmt.Errorf("No flow log collector found with name (%s)", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Not found: %s", err.Error()), "(Data) ibm_ibm_is_flow_log", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		flc := allrecs[0]
		flowLogCollector = &flc

	} else if identifier != "" {
		getFlowLogCollectorOptions := &vpcv1.GetFlowLogCollectorOptions{}

		getFlowLogCollectorOptions.SetID(d.Get("identifier").(string))

		flowlogCollector, _, err := vpcClient.GetFlowLogCollectorWithContext(context, getFlowLogCollectorOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFlowLogCollectorWithContext failed: %s", err.Error()), "(Data) ibm_ibm_is_flow_log", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		flowLogCollector = flowlogCollector
	}

	d.SetId(*flowLogCollector.ID)
	if err = d.Set("active", flowLogCollector.Active); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting active: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-active").GetDiag()
	}
	if err = d.Set("auto_delete", flowLogCollector.AutoDelete); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-auto_delete").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(flowLogCollector.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("crn", flowLogCollector.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-crn").GetDiag()
	}
	if err = d.Set("href", flowLogCollector.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-href").GetDiag()
	}
	if err = d.Set("lifecycle_state", flowLogCollector.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("name", flowLogCollector.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-name").GetDiag()
	}
	if err = d.Set("identifier", *flowLogCollector.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting identifier: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-identifier").GetDiag()
	}

	if flowLogCollector.ResourceGroup != nil {
		if err = d.Set("resource_group", dataSourceFlowLogCollectorFlattenResourceGroup(*flowLogCollector.ResourceGroup)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-resource_group").GetDiag()
		}
	}

	if flowLogCollector.StorageBucket != nil {
		if err = d.Set("storage_bucket", dataSourceFlowLogCollectorFlattenStorageBucket(*flowLogCollector.StorageBucket)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_bucket: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-storage_bucket").GetDiag()
		}
	}

	if flowLogCollector.Target != nil {
		targetIntf := flowLogCollector.Target
		target := targetIntf.(*vpcv1.FlowLogCollectorTarget) // type assertion
		if err = d.Set("target", dataSourceFlowLogCollectorFlattenTarget(*target)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-target").GetDiag()
		}
	}

	if flowLogCollector.VPC != nil {
		if err = d.Set("vpc", dataSourceFlowLogCollectorFlattenVPC(*flowLogCollector.VPC)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-vpc").GetDiag()
		}
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *flowLogCollector.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource VPC Flow Log (%s) access tags: %s", d.Id(), err)
	}

	if err = d.Set(isFlowLogAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_ibm_is_flow_log", "read", "set-access_tags").GetDiag()
	}
	return nil
}

func dataSourceFlowLogCollectorFlattenResourceGroup(result vpcv1.ResourceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceFlowLogCollectorResourceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceFlowLogCollectorResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func dataSourceFlowLogCollectorFlattenStorageBucket(result vpcv1.LegacyCloudObjectStorageBucketReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceFlowLogCollectorStorageBucketToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceFlowLogCollectorStorageBucketToMap(storageBucketItem vpcv1.LegacyCloudObjectStorageBucketReference) (storageBucketMap map[string]interface{}) {
	storageBucketMap = map[string]interface{}{}

	if storageBucketItem.Name != nil {
		storageBucketMap["name"] = storageBucketItem.Name
	}

	return storageBucketMap
}

func dataSourceFlowLogCollectorFlattenTarget(result vpcv1.FlowLogCollectorTarget) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceFlowLogCollectorTargetToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceFlowLogCollectorTargetToMap(targetItem vpcv1.FlowLogCollectorTarget) (targetMap map[string]interface{}) {
	targetMap = map[string]interface{}{}

	if targetItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceFlowLogCollectorTargetDeletedToMap(*targetItem.Deleted)
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
	if targetItem.ResourceType != nil {
		targetMap["resource_type"] = targetItem.ResourceType
	}
	if targetItem.CRN != nil {
		targetMap["crn"] = targetItem.CRN
	}

	return targetMap
}

func dataSourceFlowLogCollectorTargetDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceFlowLogCollectorFlattenVPC(result vpcv1.VPCReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceFlowLogCollectorVPCToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceFlowLogCollectorVPCToMap(vpcItem vpcv1.VPCReference) (vpcMap map[string]interface{}) {
	vpcMap = map[string]interface{}{}

	if vpcItem.CRN != nil {
		vpcMap["crn"] = vpcItem.CRN
	}
	if vpcItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceFlowLogCollectorVPCDeletedToMap(*vpcItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		vpcMap["deleted"] = deletedList
	}
	if vpcItem.Href != nil {
		vpcMap["href"] = vpcItem.Href
	}
	if vpcItem.ID != nil {
		vpcMap["id"] = vpcItem.ID
	}
	if vpcItem.Name != nil {
		vpcMap["name"] = vpcItem.Name
	}

	return vpcMap
}

func dataSourceFlowLogCollectorVPCDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}
