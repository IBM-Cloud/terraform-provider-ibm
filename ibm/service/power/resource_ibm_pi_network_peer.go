// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

func ResourceIBMPINetworkPeer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPINetworkPeerCreate,
		ReadContext:   resourceIBMPINetworkPeerRead,
		UpdateContext: resourceIBMPINetworkPeerUpdate,
		DeleteContext: resourceIBMPINetworkPeerDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_CustomerASN: {
				Description:  "ASN number at customer network side.",
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_CustomerCIDR: {
				Description:  "IP address used for configuring customer network interface with network subnet mask. customerCidr and ibmCidr must have matching network and subnet mask values.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_DefaultExportRouteFilter: {
				Default:      "allow",
				Description:  "Default action for export route filter. Allowed values: allow, deny.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{Allow, Deny}),
			},
			Arg_DefaultImportRouteFilter: {
				Default:      "allow",
				Description:  "Default action for import route filter. Allowed values: allow, deny.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{Allow, Deny}),
			},
			Arg_IBMASN: {
				Description:  "ASN number at IBM PowerVS side.",
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_IBMCIDR: {
				Description:  "IP address used for configuring IBM network interface with network subnet mask. customerCidr and ibmCidr must have matching network and subnet mask values.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Name: {
				Description:  "User defined name.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_PeerInterfaceID: {
				Description:  "Peer interface id. Use datasource 'ibmi_pi_network_peer_interfaces' to get a list of valid peer interface id.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Type: {
				Default:      "dcnetwork_bgp",
				Description:  "Type of the peer network * dcnetwork_bgp: broader gateway protocol is used to share routes between two autonomous network.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{DCNetworkBGP}),
			},
			Arg_VLAN: {
				Description:  "A vlan configured at the customer network.",
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_CreationDate: {
				Computed:    true,
				Description: "Time stamp for create network peer.",
				Type:        schema.TypeString,
			},
			Attr_Error: {
				Computed:    true,
				Description: "Error description.",
				Type:        schema.TypeString,
			},
			Attr_ExportRouteFilters: {
				Computed:    true,
				Description: "List of export route filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Action: {
							Computed:    true,
							Description: "Action of the filter.",
							Type:        schema.TypeString,
						},
						Attr_CreationDate: {
							Computed:    true,
							Description: "Time stamp for create route filter.",
							Type:        schema.TypeString,
						},
						Attr_Direction: {
							Computed:    true,
							Description: "Direction of the filter.",
							Type:        schema.TypeString,
						},
						Attr_Error: {
							Computed:    true,
							Description: "Error description.",
							Type:        schema.TypeString,
						},
						Attr_GE: {
							Computed:    true,
							Description: "The minimum matching length of the prefix-set.",
							Type:        schema.TypeInt,
						},
						Attr_Index: {
							Computed:    true,
							Description: "Priority or order of the filter.",
							Type:        schema.TypeInt,
						},
						Attr_LE: {
							Computed:    true,
							Description: "The maximum matching length of the prefix-set.",
							Type:        schema.TypeInt,
						},
						Attr_Prefix: {
							Computed:    true,
							Description: "IP prefix representing an address and mask length of the prefix-set.",
							Type:        schema.TypeString,
						},
						Attr_RouteFilterID: {
							Computed:    true,
							Description: "Route filter ID.",
							Type:        schema.TypeString,
						},
						Attr_State: {
							Computed:    true,
							Description: "Status of the route filter.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_ImportRouteFilters: {
				Computed:    true,
				Description: "List of import route filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Action: {
							Computed:    true,
							Description: "Action of the filter.",
							Type:        schema.TypeString,
						},
						Attr_CreationDate: {
							Computed:    true,
							Description: "Time stamp for create route filter.",
							Type:        schema.TypeString,
						},
						Attr_Direction: {
							Computed:    true,
							Description: "Direction of the filter.",
							Type:        schema.TypeString,
						},
						Attr_Error: {
							Computed:    true,
							Description: "Error description.",
							Type:        schema.TypeString,
						},
						Attr_GE: {
							Computed:    true,
							Description: "The minimum matching length of the prefix-set.",
							Type:        schema.TypeInt,
						},
						Attr_Index: {
							Computed:    true,
							Description: "Priority or order of the filter.",
							Type:        schema.TypeInt,
						},
						Attr_LE: {
							Computed:    true,
							Description: "The maximum matching length of the prefix-set.",
							Type:        schema.TypeInt,
						},
						Attr_Prefix: {
							Computed:    true,
							Description: "IP prefix representing an address and mask length of the prefix-set.",
							Type:        schema.TypeString,
						},
						Attr_RouteFilterID: {
							Computed:    true,
							Description: "Route filter ID.",
							Type:        schema.TypeString,
						},
						Attr_State: {
							Computed:    true,
							Description: "Status of the route filter.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_PeerID: {
				Computed:    true,
				Description: "Network peer id.",
				Type:        schema.TypeString,
			},
			Attr_State: {
				Computed:    true,
				Description: "Status of the network peer.",
				Type:        schema.TypeString,
			},
			Attr_UpdatedDate: {
				Computed:    true,
				Description: "Time stamp for update network peer.",
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPINetworkPeerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	customerASN := (int64)(d.Get(Arg_CustomerASN).(int))
	customerCIDR := d.Get(Arg_CustomerCIDR).(string)
	ibmASN := (int64)(d.Get(Arg_IBMASN).(int))
	ibmCIDR := d.Get(Arg_IBMCIDR).(string)
	name := d.Get(Arg_Name).(string)
	peerInterfaceID := d.Get(Arg_PeerInterfaceID).(string)
	vlan := (int64)(d.Get(Arg_VLAN).(int))

	body := &models.NetworkPeerCreate{
		CustomerASN:     &customerASN,
		CustomerCidr:    &customerCIDR,
		IbmASN:          &ibmASN,
		IbmCidr:         &ibmCIDR,
		Name:            &name,
		PeerInterfaceID: &peerInterfaceID,
		Vlan:            &vlan,
	}
	if derf, ok := d.GetOk(Arg_DefaultExportRouteFilter); ok {
		body.DefaultExportRouteFilter = flex.PtrToString(derf.(string))
	}
	if dirf, ok := d.GetOk(Arg_DefaultImportRouteFilter); ok {
		body.DefaultImportRouteFilter = flex.PtrToString(dirf.(string))
	}
	if networkType, ok := d.GetOk(Arg_Type); ok {
		body.Type = flex.PtrToString(networkType.(string))
	}

	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, cloudInstanceID)
	networkPeer, err := networkC.CreateNetworkPeer(body)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating network peer: %s", err))
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *networkPeer.ID))
	_, err = isWaitForIBMPINetworkPeerCreated(ctx, networkC, *networkPeer.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIBMPINetworkPeerRead(ctx, d, meta)
}

func resourceIBMPINetworkPeerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, parts[0])
	networkPeer, err := networkC.GetNetworkPeer(parts[1])
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set(Arg_CustomerASN, networkPeer.CustomerASN)
	d.Set(Arg_CustomerCIDR, networkPeer.CustomerCidr)
	d.Set(Arg_DefaultExportRouteFilter, networkPeer.DefaultExportRouteFilter)
	d.Set(Arg_DefaultImportRouteFilter, networkPeer.DefaultImportRouteFilter)
	d.Set(Arg_IBMASN, networkPeer.IbmASN)
	d.Set(Arg_IBMCIDR, networkPeer.IbmCidr)
	d.Set(Arg_Name, networkPeer.Name)
	d.Set(Arg_PeerInterfaceID, networkPeer.PeerInterfaceID)
	d.Set(Arg_Type, networkPeer.Type)
	d.Set(Arg_VLAN, networkPeer.Vlan)
	d.Set(Attr_CreationDate, networkPeer.CreationDate)
	d.Set(Attr_Error, networkPeer.Error)
	exportRouteFilters := []map[string]interface{}{}
	if networkPeer.ExportRouteFilters != nil {
		for _, erp := range networkPeer.ExportRouteFilters {
			exportRouteFilter := dataSourceIBMPINetworkPeerRouteFilterToMap(erp)
			exportRouteFilters = append(exportRouteFilters, exportRouteFilter)
		}
	}
	d.Set(Attr_ExportRouteFilters, exportRouteFilters)

	importRouteFilters := []map[string]interface{}{}
	if networkPeer.ImportRouteFilters != nil {
		for _, irp := range networkPeer.ImportRouteFilters {
			importRouteFilter := dataSourceIBMPINetworkPeerRouteFilterToMap(irp)
			importRouteFilters = append(importRouteFilters, importRouteFilter)
		}
	}
	d.Set(Attr_ImportRouteFilters, importRouteFilters)
	d.Set(Attr_PeerID, networkPeer.ID)
	d.Set(Attr_State, networkPeer.State)
	d.Set(Attr_UpdatedDate, networkPeer.UpdatedDate)

	return nil
}

func resourceIBMPINetworkPeerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	updateBody := &models.NetworkPeerUpdate{}
	hasChange := false

	if d.HasChange(Arg_CustomerASN) {
		asn := (int64)(d.Get(Arg_CustomerASN).(int))
		updateBody.CustomerASN = &asn
		hasChange = true
	}
	if d.HasChange(Arg_CustomerCIDR) {
		cidr := d.Get(Arg_CustomerCIDR).(string)
		updateBody.CustomerCidr = &cidr
		hasChange = true
	}
	if d.HasChange(Arg_DefaultExportRouteFilter) {
		updateBody.DefaultExportRouteFilter = flex.PtrToString(d.Get(Arg_DefaultExportRouteFilter).(string))
		hasChange = true
	}
	if d.HasChange(Arg_DefaultImportRouteFilter) {
		updateBody.DefaultImportRouteFilter = flex.PtrToString(d.Get(Arg_DefaultImportRouteFilter).(string))
		hasChange = true
	}
	if d.HasChange(Arg_IBMASN) {
		asn := (int64)(d.Get(Arg_IBMASN).(int))
		updateBody.IbmASN = &asn
		hasChange = true
	}
	if d.HasChange(Arg_IBMCIDR) {
		cidr := d.Get(Arg_IBMCIDR).(string)
		updateBody.IbmCidr = &cidr
		hasChange = true
	}
	if d.HasChange(Arg_Name) {
		name := d.Get(Arg_Name).(string)
		updateBody.Name = &name
		hasChange = true
	}
	if d.HasChange(Arg_PeerInterfaceID) {
		id := d.Get(Arg_PeerInterfaceID).(string)
		updateBody.PeerInterfaceID = &id
		hasChange = true
	}
	if d.HasChange(Arg_Type) {
		updateBody.Type = flex.PtrToString(d.Get(Arg_Type).(string))
		hasChange = true
	}
	if d.HasChange(Arg_VLAN) {
		vlan := (int64)(d.Get(Arg_VLAN).(int))
		updateBody.Vlan = &vlan
		hasChange = true
	}

	if hasChange {
		parts, err := flex.SepIdParts(d.Id(), "/")
		if err != nil {
			return diag.FromErr(err)
		}
		networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, parts[0])
		_, err = networkC.UpdateNetworkPeer(parts[1], updateBody)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = isWaitForIBMPINetworkPeerUpdated(ctx, networkC, parts[1], d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIBMPINetworkPeerRead(ctx, d, meta)
}

func resourceIBMPINetworkPeerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, parts[0])
	err = networkC.DeleteNetworkPeer(parts[1])
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = isWaitForIBMPINetworkPeerDeleted(ctx, networkC, parts[1], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}
func isWaitForIBMPINetworkPeerCreated(ctx context.Context, client *instance.IBMPINetworkPeerClient, networkPeerID string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Configuring, State_Removing, State_Updating},
		Target:     []string{State_Active, State_Error},
		Refresh:    isIBMPINetworkPeerRefreshFunc(client, networkPeerID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}
func isIBMPINetworkPeerRefreshFunc(client *instance.IBMPINetworkPeerClient, networkPeerID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		networkPeerRouteFilter, err := client.GetNetworkPeer(networkPeerID)
		if err != nil {
			return nil, "", err
		}
		if strings.ToLower(*networkPeerRouteFilter.State) == State_Configuring {
			return networkPeerRouteFilter, State_Configuring, nil
		}
		if strings.ToLower(*networkPeerRouteFilter.State) == State_Removing {
			return networkPeerRouteFilter, State_Removing, nil
		}
		if strings.ToLower(*networkPeerRouteFilter.State) == State_Updating {
			return networkPeerRouteFilter, State_Updating, nil
		}
		if strings.ToLower(*networkPeerRouteFilter.State) == State_Error {
			return networkPeerRouteFilter, State_Error, fmt.Errorf("[ERROR] the network peer %s failed with %s", networkPeerID, *networkPeerRouteFilter.Error)
		}
		return networkPeerRouteFilter, State_Active, nil
	}
}
func isWaitForIBMPINetworkPeerDeleted(ctx context.Context, client *instance.IBMPINetworkPeerClient, networkPeerID string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Found},
		Target:     []string{State_NotFound},
		Refresh:    isIBMPINetworkPeerRefreshDeleteFunc(client, networkPeerID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkPeerRefreshDeleteFunc(client *instance.IBMPINetworkPeerClient, networkPeerID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		networkPeer, err := client.GetNetworkPeer(networkPeerID)
		if err != nil {
			return networkPeer, State_NotFound, nil
		}
		return networkPeer, State_Found, nil
	}
}

func isWaitForIBMPINetworkPeerUpdated(ctx context.Context, client *instance.IBMPINetworkPeerClient, networkPeerID string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Configuring, State_Removing, State_Updating},
		Target:     []string{State_Active, State_Error},
		Refresh:    isIBMPINetworkPeerRefreshFunc(client, networkPeerID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}
