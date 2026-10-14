// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerPeerGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"dynamic_route_server_peer_group_id", "name"},
				Description:  "The dynamic route server peer group identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"dynamic_route_server_peer_group_id", "name"},
				Description:  "The name of the dynamic route server peer group.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this dynamic route server peer group was created.",
			},
			"creator": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The types of resource that created peers in this peer group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"health_monitor": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The health monitoring configuration used for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mode used for this health monitor: `bfd`, `bgp` or `none`.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The roles used for this dynamic route server peer group health monitor configuration.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group.",
			},
			"lifecycle_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `lifecycle_state` (if any).",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the dynamic route server peer group.",
			},
			"peers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The peers in this dynamic route server peer group.",
				Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPeerSchema(),
			},
			"prefix_watcher": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The prefix watcher configuration for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The prefixes excluded from monitoring.",
							Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPrefixSchema(),
						},
						"monitored_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The prefixes monitored by this peer group.",
							Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPrefixSchema(),
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"roles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The roles for this dynamic route server peer group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of this dynamic route server peer group: `enabled` or `disabled`.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this dynamic route server peer group.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status reasons for the dynamic route server peer group.",
				Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPrefixSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP prefix in CIDR format.",
			},
			"ge": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum prefix length to match.",
			},
			"le": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum prefix length to match.",
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)

	var peerGroup *vpcv1.DynamicRouteServerPeerGroup
	if peerGroupID, ok := d.GetOk("dynamic_route_server_peer_group_id"); ok {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetID(peerGroupID.(string))

		peerGroup, _, err = vpcClient.GetDynamicRouteServerPeerGroupWithContext(context, getOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_peer_group", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else {
		name := d.Get("name").(string)
		listOptions := &vpcv1.ListDynamicRouteServerPeerGroupsOptions{}
		listOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		pager, err := vpcClient.NewDynamicRouteServerPeerGroupsPager(listOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		allItems, err := pager.GetAll()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupsPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		peerGroup, err = DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(allItems, name)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group", "read", "find-by-name").GetDiag()
		}
	}
	if peerGroup == nil || peerGroup.ID == nil {
		return flex.DiscriminatedTerraformErrorf(fmt.Errorf("peer group response has no id"), "peer group response has no id", "(Data) ibm_is_dynamic_route_server_peer_group", "read", "peer-group-id").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", dynamicRouteServerID, *peerGroup.ID))
	if err = d.Set("dynamic_route_server_peer_group_id", *peerGroup.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_id: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group", "read", "set-dynamic_route_server_peer_group_id").GetDiag()
	}

	peerGroupMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(peerGroup)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group", "read", "peer-group-to-map").GetDiag()
	}
	for _, key := range []string{"created_at", "creator", "health_monitor", "href", "lifecycle_reasons", "lifecycle_state", "name", "peers", "prefix_watcher", "resource_type", "roles", "state", "status", "status_reasons"} {
		if value, ok := peerGroupMap[key]; ok {
			if err = d.Set(key, value); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "(Data) ibm_is_dynamic_route_server_peer_group", "read", "set-"+key).GetDiag()
			}
		}
	}

	return nil
}

// DataSourceIBMIsDynamicRouteServerPeerGroupFindByName returns the peer group whose name matches, or an error when none or several match.
func DataSourceIBMIsDynamicRouteServerPeerGroupFindByName(peerGroups []vpcv1.DynamicRouteServerPeerGroup, name string) (*vpcv1.DynamicRouteServerPeerGroup, error) {
	if name == "" {
		return nil, fmt.Errorf("name must not be empty")
	}
	var found *vpcv1.DynamicRouteServerPeerGroup
	for i := range peerGroups {
		if peerGroups[i].Name != nil && *peerGroups[i].Name == name {
			if found != nil {
				return nil, fmt.Errorf("more than one dynamic route server peer group named %q was found", name)
			}
			found = &peerGroups[i]
		}
	}
	if found == nil {
		return nil, fmt.Errorf("no dynamic route server peer group named %q was found", name)
	}
	return found, nil
}

// DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap flattens a peer group into the data source shape.
func DataSourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupToMap(model *vpcv1.DynamicRouteServerPeerGroup) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("dynamic route server peer group is nil")
	}
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = flex.DateTimeToString(model.CreatedAt)
	}
	creator := model.Creator
	if creator == nil {
		creator = []string{}
	}
	modelMap["creator"] = creator
	if model.HealthMonitor != nil {
		healthMonitorMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model.HealthMonitor)
		if err != nil {
			return modelMap, err
		}
		modelMap["health_monitor"] = []map[string]interface{}{healthMonitorMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	lifecycleReasons := []map[string]interface{}{}
	for i := range model.LifecycleReasons {
		lifecycleReasons = append(lifecycleReasons, ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(&model.LifecycleReasons[i]))
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	peers := []map[string]interface{}{}
	for _, peer := range model.Peers {
		peerMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(peer)
		if err != nil {
			return modelMap, err
		}
		peers = append(peers, peerMap)
	}
	modelMap["peers"] = peers
	if model.PrefixWatcher != nil {
		modelMap["prefix_watcher"] = []map[string]interface{}{ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(model.PrefixWatcher)}
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	roles := model.Roles
	if roles == nil {
		roles = []string{}
	}
	modelMap["roles"] = roles
	if model.State != nil {
		modelMap["state"] = *model.State
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	statusReasons := []map[string]interface{}{}
	for i := range model.StatusReasons {
		statusReasons = append(statusReasons, ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(&model.StatusReasons[i]))
	}
	modelMap["status_reasons"] = statusReasons
	return modelMap, nil
}
