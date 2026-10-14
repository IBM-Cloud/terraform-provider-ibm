// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicies() *schema.Resource {
	policySchema := dataSourceIBMIsDynamicRouteServerPeerGroupPolicyItemSchema()
	policySchema["id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The unique identifier for this dynamic route server peer group policy.",
	}
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server peer group identifier.",
			},
			"sort": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "-created_at",
				Description: "Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. Allowable values are `created_at` and `name`.",
			},
			"policies": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The policies of the dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: policySchema,
				},
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listOptions := &vpcv1.ListDynamicRouteServerPeerGroupPoliciesOptions{}
	listOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	listOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))
	if sort, ok := d.GetOk("sort"); ok && sort.(string) != "" {
		listOptions.SetSort(sort.(string))
	}

	pager, err := vpcClient.NewDynamicRouteServerPeerGroupPoliciesPager(listOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAllWithContext(context)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupPoliciesPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "policies-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("policies", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policies %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "policies-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

// DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap flattens any policy
// subtype (or the base struct) into the list item map.
func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyIntf) (map[string]interface{}, error) {
	policy, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
	if err != nil {
		return nil, err
	}
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(policy)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group policy is nil")
	}
	return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group policy is nil")
	}
	return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group policy is nil")
	}
	return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group policy is nil")
	}
	return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model)
}
