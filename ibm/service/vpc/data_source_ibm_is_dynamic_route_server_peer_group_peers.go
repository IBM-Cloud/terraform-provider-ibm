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

func DataSourceIBMIsDynamicRouteServerPeerGroupPeers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPeersRead,

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
			"peers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of dynamic route server peer group peers.",
				Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPeerSchema(),
			},
		},
	}
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPeerSchema is the computed schema of one peer, shared by the peer data sources.
func dataSourceIBMIsDynamicRouteServerPeerGroupPeerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: dataSourceIBMIsDynamicRouteServerPeerGroupPeerAttributes(),
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPeerAttributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier for this dynamic route server peer group peer.",
		},
		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The autonomous system number (ASN) for this dynamic route server peer group peer.",
		},
		"bidirectional_forwarding_detection": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The bidirectional forwarding detection (BFD) configuration for this dynamic route server peer group peer.",
			Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerBFDSchema(),
		},
		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date and time that this dynamic route server peer group peer was created.",
		},
		"creator": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The type of resource that created this peer: `dynamic_route_server` or `transit_gateway`.",
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The endpoint for this dynamic route server peer group peer.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"address": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The IP address of the endpoint.",
					},
					"crn": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The CRN for the transit gateway endpoint, if any.",
					},
					"deleted": &schema.Schema{
						Type:        schema.TypeList,
						Computed:    true,
						Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
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
					"gateway": &schema.Schema{
						Type:        schema.TypeList,
						Computed:    true,
						Description: "The gateway used to reach the endpoint address.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"address": &schema.Schema{
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The IP address of the gateway.",
								},
							},
						},
					},
					"href": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The URL for the reserved IP used as the endpoint.",
					},
					"id": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The unique identifier for the endpoint resource.",
					},
					"name": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The name for the reserved IP used as the endpoint.",
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
			Description: "The URL for this dynamic route server peer group peer.",
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
			Description: "The lifecycle state of the dynamic route server peer group peer.",
		},
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name for this dynamic route server peer group peer.",
		},
		"priority": &schema.Schema{
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The priority of this dynamic route server peer group peer.",
		},
		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The resource type.",
		},
		"sessions": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The BGP sessions for this dynamic route server peer group peer.",
			Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerSessionSchema(),
		},
		"state": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The state of this dynamic route server peer group peer: `enabled` or `disabled`.",
		},
		"status": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of this dynamic route server peer group peer connection: `down` or `up`.",
		},
		"status_reasons": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The status reasons for the dynamic route server peer group peer.",
			Elem:        resourceIBMIsDynamicRouteServerPeerGroupPeerReasonSchema(),
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPeersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := dataSourceIBMIsDynamicRouteServerPeerGroupPeersList(vpcClient, d.Get("dynamic_route_server_id").(string), d.Get("dynamic_route_server_peer_group_id").(string))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupPeersPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerPeerGroupPeersID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "peers-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("peers", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting peers %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "peers-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPeersList lists every peer of a peer group using the SDK pager.
func dataSourceIBMIsDynamicRouteServerPeerGroupPeersList(vpcClient *vpcv1.VpcV1, dynamicRouteServerID, peerGroupID string) ([]vpcv1.DynamicRouteServerPeerGroupPeerIntf, error) {
	listOptions := &vpcv1.ListDynamicRouteServerPeerGroupPeersOptions{}
	listOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	listOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)

	pager, err := vpcClient.NewDynamicRouteServerPeerGroupPeersPager(listOptions)
	if err != nil {
		return nil, err
	}
	return pager.GetAll()
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPeersID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerPeerGroupPeersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

// DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap flattens any peer subtype into the data source shape.
func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(model)
}
