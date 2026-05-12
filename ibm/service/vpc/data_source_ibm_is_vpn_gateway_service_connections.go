// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isVPNGatewayServiceConnectionId               = "id"
	isVPNGatewayServiceConnectionCreatedAt        = "created_at"
	isVPNGatewayServiceConnectionResourceGroup    = "resource_group"
	isVPNGatewayServiceConnections                = "service_connections"
	isVPNGatewayServiceConnectionCreator          = "creator"
	isVPNGatewayServiceConnectionStatus           = "status"
	isVPNGatewayServiceConnectionStatusReasons    = "status_reasons"
	isVPNGatewayServiceConnectionLifecycleState   = "lifecycle_state"
	isVPNGatewayServiceConnectionLifecycleReasons = "lifecycle_reasons"
)

func DataSourceIBMIsVPNGatewayServiceConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayServiceConnectionsRead,

		Schema: map[string]*schema.Schema{
			"vpn_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPN gateway identifier.",
			},
			isVPNGatewayServiceConnections: {
				Type:        schema.TypeList,
				Description: "Collection of VPN Gateway service connections",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this VPN service connection was created.",
						},
						"creator": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for transit gateway resource.",
									},
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for transit gateway resource.",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPN gateway service connection",
						},
						"lifecycle_reasons": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current `lifecycle_state` (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this lifecycle state.",
									},
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about the reason for this lifecycle state.",
									},
								},
							},
						},
						"lifecycle_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the VPN service connection.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of this service connection:- `up`: operating normally- `degraded`: operating with compromised performance- `down`: not operational.",
						},
						"status_reasons": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current VPN service connection status (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A snake case string succinctly identifying the status reason. The enumerated values for this property may https://cloud.ibm.com/apidocs/vpc#property-value-expansion in the future.",
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this VPN service connection's status.",
									},
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about this status reason.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsVPNGatewayServiceConnectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway_service_connections", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	vpnGateway := ""
	if vpnGatewayIntf, ok := d.GetOk("vpn_gateway"); ok {
		vpnGateway = vpnGatewayIntf.(string)
	}

	start := ""
	allrecs := []vpcv1.VPNGatewayServiceConnection{}
	for {
		listvpnGWServiceConnectionsOptions := sess.NewListVPNGatewayServiceConnectionsOptions(vpnGateway)
		listvpnGWServiceConnectionsOptions.VPNGatewayID = &vpnGateway
		if start != "" {
			listvpnGWServiceConnectionsOptions.Start = &start
		}
		availableVPNGatewayServiceConnections, detail, err := sess.ListVPNGatewayServiceConnections(listvpnGWServiceConnectionsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error reading list of VPN Gateway service connections:%s\n%s", err, detail), "(Data) ibm_is_vpn_gateway_service_connections", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return diag.FromErr(tfErr)
		}
		start = flex.GetNext(availableVPNGatewayServiceConnections.Next)
		allrecs = append(allrecs, availableVPNGatewayServiceConnections.ServiceConnections...)
		if start == "" {
			break
		}
	}

	vpngatewayServiceConnections := make([]map[string]interface{}, 0)
	for _, serviceConnection := range allrecs {
		connection := map[string]interface{}{}
		connection[isVPNGatewayServiceConnectionCreatedAt] = serviceConnection.CreatedAt.String()
		connection[isVPNGatewayServiceConnectionId] = *serviceConnection.ID
		connection[isVPNGatewayServiceConnectionCreator] = resourceVPNGatewayServiceConnectionFlattenCreator(serviceConnection.Creator)
		connection[isVPNGatewayServiceConnectionLifecycleReasons] = resourceVPNGatewayServiceConnectionFlattenLifecycleReasons(serviceConnection.LifecycleReasons)
		connection[isVPNGatewayServiceConnectionLifecycleState] = *serviceConnection.LifecycleState
		connection[isVPNGatewayServiceConnectionStatus] = *serviceConnection.Status
		connection[isVPNGatewayServiceConnectionStatusReasons] = resourceVPNGatewayServiceConnectionFlattenStateReasons(serviceConnection.StatusReasons)

		vpngatewayServiceConnections = append(vpngatewayServiceConnections, connection)
	}

	d.SetId(dataSourceIBMVPNGatewayServiceConnectionsID(d))
	if err = d.Set(isVPNGatewayServiceConnections, vpngatewayServiceConnections); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_connections %s", err), "(Data) ibm_is_vpn_gateway_service_connections", "read", "vpn_gateway-service-connections-set").GetDiag()
	}
	return nil
}

// dataSourceIBMVPNGatewayServiceConnectionsID returns a reasonable ID  list.
func dataSourceIBMVPNGatewayServiceConnectionsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceVPNGatewayServiceConnectionFlattenLifecycleReasons(lifecycleReasons []vpcv1.VPNGatewayServiceConnectionLifecycleReason) (lifecycleReasonsList []map[string]interface{}) {
	lifecycleReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range lifecycleReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			lifecycleReasonsList = append(lifecycleReasonsList, currentLR)
		}
	}
	return lifecycleReasonsList
}

func resourceVPNGatewayServiceConnectionFlattenStateReasons(healthReasons []vpcv1.VPNGatewayServiceConnectionStatusReason) (statusReasonsList []map[string]interface{}) {
	statusReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range healthReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			statusReasonsList = append(statusReasonsList, currentLR)
		}
	}
	return statusReasonsList
}

func resourceVPNGatewayServiceConnectionFlattenCreator(model vpcv1.VPNGatewayServiceConnectionCreatorIntf) []map[string]interface{} {
	modelMap := make(map[string]interface{})

	connectionCreatorItem, ok := model.(*vpcv1.VPNGatewayServiceConnectionCreator)
	if !ok || connectionCreatorItem == nil {
		return nil
	}
	if connectionCreatorItem.CRN != nil {
		modelMap["crn"] = *connectionCreatorItem.CRN
	}
	if connectionCreatorItem.ID != nil {
		modelMap["id"] = *connectionCreatorItem.ID
	}
	if connectionCreatorItem.ResourceType != nil {
		modelMap["resource_type"] = *connectionCreatorItem.ResourceType
	}
	return []map[string]interface{}{modelMap}
}
