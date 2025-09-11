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

func ResourceIBMPINetworkPeerRouteFilter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPINetworkPeerRouteFilterCreate,
		ReadContext:   resourceIBMPINetworkPeerRouteFilterRead,
		DeleteContext: resourceIBMPINetworkPeerRouteFilterDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_Action: {
				Default:      "allow",
				Description:  "Action of the filter.",
				ForceNew:     true,
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{Allow, Deny}),
			},
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Direction: {
				Description:  "Direction of the filter.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{Export, Import}),
			},
			Arg_GE: {
				Description: "The minimum matching length of the prefix-set(1 ≤ value ≤ 32 & value ≤ LE).",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			Arg_Index: {
				Description: "Priority or order of the filter.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeInt,
			},
			Arg_LE: {
				Description: "The maximum matching length of the prefix-set( 1 ≤ value ≤ 32 & value >= GE).",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			Arg_NetworkPeerID: {
				Description: "Network peer ID.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_Prefix: {
				Description: "IP prefix representing an address and mask length of the prefix-set.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			// Attribute
			Attr_CreationDate: {
				Computed:    true,
				Description: "Time stamp for create route filter.",
				Type:        schema.TypeString,
			},
			Attr_Error: {
				Computed:    true,
				Description: "Error description.",
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
	}
}

func resourceIBMPINetworkPeerRouteFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	action := d.Get(Arg_Action).(string)
	direction := d.Get(Arg_Direction).(string)
	index := int64(d.Get(Arg_Index).(int))
	networkPeerID := d.Get(Arg_NetworkPeerID).(string)
	prefix := d.Get(Arg_Prefix).(string)
	body := &models.RouteFilterCreate{
		Action:    &action,
		Direction: &direction,
		Index:     &index,
		Prefix:    &prefix,
	}
	if ge, ok := d.Get(Arg_GE).(int); ok {
		body.GE = int64(ge)
	}
	if le, ok := d.Get(Arg_LE).(int); ok {
		body.LE = int64(le)
	}
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, cloudInstanceID)
	routeFilter, err := networkC.CreateNetworkPeersRouteFilters(networkPeerID, body)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating the network peer route: %s", err))
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, networkPeerID, *routeFilter.RouteFilterID))
	_, err = isWaitForIBMPINetworkPeerRouteCreated(ctx, networkC, networkPeerID, *routeFilter.RouteFilterID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIBMPINetworkPeerRouteFilterRead(ctx, d, meta)
}

func resourceIBMPINetworkPeerRouteFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, parts[0])
	routeFilter, err := networkC.GetNetworkPeersRouteFilter(parts[1], parts[2])
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.Set(Arg_Action, routeFilter.Action)
	d.Set(Arg_Direction, routeFilter.Direction)
	d.Set(Arg_GE, routeFilter.GE)
	d.Set(Arg_Index, routeFilter.Index)
	d.Set(Arg_LE, routeFilter.LE)
	d.Set(Arg_Prefix, routeFilter.Prefix)
	d.Set(Attr_CreationDate, routeFilter.CreationDate)
	d.Set(Attr_Error, routeFilter.Error)
	d.Set(Attr_RouteFilterID, routeFilter.RouteFilterID)
	d.Set(Attr_State, routeFilter.State)

	return nil
}

func resourceIBMPINetworkPeerRouteFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, parts[0])
	err = networkC.DeleteNetworkPeersRouteFilter(parts[1], parts[2])
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = isWaitForIBMPINetworkPeerRouteFilterDeleted(ctx, networkC, parts[1], parts[2], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}
func isWaitForIBMPINetworkPeerRouteCreated(ctx context.Context, client *instance.IBMPINetworkPeerClient, networkPeerID, routeFilterID string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Configuring, State_Removing, State_Updating},
		Target:     []string{State_Active, State_Error},
		Refresh:    isIBMPINetworkPeerRouteRefreshCreateFunc(client, networkPeerID, routeFilterID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}
func isIBMPINetworkPeerRouteRefreshCreateFunc(client *instance.IBMPINetworkPeerClient, networkPeerID, routeFilterID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		networkPeerRouteFilter, err := client.GetNetworkPeersRouteFilter(networkPeerID, routeFilterID)
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
			return networkPeerRouteFilter, State_Error, nil
		}
		return networkPeerRouteFilter, State_Active, nil
	}
}
func isWaitForIBMPINetworkPeerRouteFilterDeleted(ctx context.Context, client *instance.IBMPINetworkPeerClient, networkPeerID, routeFilterID string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Found},
		Target:     []string{State_NotFound},
		Refresh:    isIBMPINetworkPeerRouteRefreshDeleteFunc(client, networkPeerID, routeFilterID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkPeerRouteRefreshDeleteFunc(client *instance.IBMPINetworkPeerClient, networkPeerID, routeFilterID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		networkPeerRouteFilter, err := client.GetNetworkPeersRouteFilter(networkPeerID, routeFilterID)
		if err != nil {
			return networkPeerRouteFilter, State_NotFound, nil
		}
		return networkPeerRouteFilter, State_Found, nil
	}
}
