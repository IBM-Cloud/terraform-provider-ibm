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

func DataSourceIBMIsDynamicRouteServerPeerGroupPeer() *schema.Resource {
	attributes := dataSourceIBMIsDynamicRouteServerPeerGroupPeerAttributes()
	delete(attributes, "id")
	delete(attributes, "name")
	attributes["dynamic_route_server_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The dynamic route server identifier.",
	}
	attributes["dynamic_route_server_peer_group_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The dynamic route server peer group identifier.",
	}
	attributes["dynamic_route_server_peer_group_peer_id"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"dynamic_route_server_peer_group_peer_id", "name"},
		Description:  "The dynamic route server peer group peer identifier.",
	}
	attributes["name"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"dynamic_route_server_peer_group_peer_id", "name"},
		Description:  "The name of the dynamic route server peer group peer.",
	}
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPeerRead,
		Schema:      attributes,
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPeerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	peerGroupID := d.Get("dynamic_route_server_peer_group_id").(string)

	var peerIntf vpcv1.DynamicRouteServerPeerGroupPeerIntf
	if peerID, ok := d.GetOk("dynamic_route_server_peer_group_peer_id"); ok {
		getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
		getOptions.SetID(peerID.(string))

		peerIntf, _, err = vpcClient.GetDynamicRouteServerPeerGroupPeerWithContext(context, getOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else {
		name := d.Get("name").(string)
		allItems, err := dataSourceIBMIsDynamicRouteServerPeerGroupPeersList(vpcClient, dynamicRouteServerID, peerGroupID)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupPeersPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		peerIntf, err = DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(allItems, name)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "find-by-name").GetDiag()
		}
	}

	peerMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerToMap(peerIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "peer-to-map").GetDiag()
	}
	peerID, _ := peerMap["id"].(string)
	if peerID == "" {
		return flex.DiscriminatedTerraformErrorf(fmt.Errorf("peer response has no id"), "peer response has no id", "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "peer-id").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", dynamicRouteServerID, peerGroupID, peerID))
	if err = d.Set("dynamic_route_server_peer_group_peer_id", peerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_peer_group_peer_id: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "set-dynamic_route_server_peer_group_peer_id").GetDiag()
	}
	for _, key := range []string{"asn", "bidirectional_forwarding_detection", "created_at", "creator", "endpoint", "href", "lifecycle_reasons", "lifecycle_state", "name", "priority", "resource_type", "sessions", "state", "status", "status_reasons"} {
		if value, ok := peerMap[key]; ok {
			if err = d.Set(key, value); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "(Data) ibm_is_dynamic_route_server_peer_group_peer", "read", "set-"+key).GetDiag()
			}
		}
	}

	return nil
}

// DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName returns the peer whose name matches, or an error when none or several match.
func DataSourceIBMIsDynamicRouteServerPeerGroupPeerFindByName(peers []vpcv1.DynamicRouteServerPeerGroupPeerIntf, name string) (vpcv1.DynamicRouteServerPeerGroupPeerIntf, error) {
	if name == "" {
		return nil, fmt.Errorf("name must not be empty")
	}
	var found vpcv1.DynamicRouteServerPeerGroupPeerIntf
	for _, peer := range peers {
		var peerName *string
		switch p := peer.(type) {
		case *vpcv1.DynamicRouteServerPeerGroupPeerAddress:
			peerName = p.Name
		case *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway:
			peerName = p.Name
		case *vpcv1.DynamicRouteServerPeerGroupPeer:
			peerName = p.Name
		default:
			return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerIntf subtype encountered")
		}
		if peerName != nil && *peerName == name {
			if found != nil {
				return nil, fmt.Errorf("more than one dynamic route server peer group peer named %q was found", name)
			}
			found = peer
		}
	}
	if found == nil {
		return nil, fmt.Errorf("no dynamic route server peer group peer named %q was found", name)
	}
	return found, nil
}
