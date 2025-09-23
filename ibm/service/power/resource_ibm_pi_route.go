// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIRouteCreate,
		ReadContext:   resourceIBMPIRouteRead,
		UpdateContext: resourceIBMPIRouteUpdate,
		DeleteContext: resourceIBMPIRouteDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourcePowerUserTagsCustomizeDiff(diff)
			},
		),
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_Action: {
				Default:      Deliver,
				Description:  "Action for route. Valid values are \"deliver\".",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{Deliver}, false),
			},
			Arg_Advertise: {
				Default:      Enable,
				Description:  "Indicates if the route is advertised. Valid values are \"enable\" and \"disable\".",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{Enable, Disable}, false),
			},
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Destination: {
				Description: "Destination of route.",
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_DestinationType: {
				Default:      IPV4_Address,
				Description:  "The destination type. Valid values are \"ipv4-address\".",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{IPV4_Address}, false),
			},
			Arg_Enabled: {
				Default:     false,
				Description: "Indicates if the route should be enabled in the fabric.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			Arg_Name: {
				Description: "Name of the route.",
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_NextHop: {
				Description: "The next hop.",
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_NextHopType: {
				Default:      IPV4_Address,
				Description:  "The next hop type. Valid values are \"ipv4-address\".",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{IPV4_Address}, false),
			},
			Arg_UserTags: {
				Computed:    true,
				Description: "The user tags attached to this resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},

			// Attributes
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of this resource.",
				Type:        schema.TypeString,
			},
			Attr_RouteID: {
				Computed:    true,
				Description: "The unique route ID.",
				Type:        schema.TypeString,
			},
			Attr_State: {
				Computed:    true,
				Description: "The state of the route.",
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPIRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	action := d.Get(Arg_Action).(string)
	advertise := d.Get(Arg_Advertise).(string)
	destination := d.Get(Arg_Destination).(string)
	destinationType := d.Get(Arg_DestinationType).(string)
	enabled := d.Get(Arg_Enabled).(bool)
	name := d.Get(Arg_Name).(string)
	nextHop := d.Get(Arg_NextHop).(string)
	nextHopType := d.Get(Arg_NextHopType).(string)
	routeClient := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)

	body := &models.RouteCreate{
		Action:          &action,
		Advertise:       &advertise,
		Destination:     &destination,
		DestinationType: &destinationType,
		Enabled:         &enabled,
		Name:            &name,
		NextHop:         &nextHop,
		NextHopType:     &nextHopType,
	}

	if v, ok := d.GetOk(Arg_UserTags); ok {
		userTags := flex.FlattenSet(v.(*schema.Set))
		body.UserTags = userTags
	}

	route, err := routeClient.Create(body)
	if err != nil {
		return diag.FromErr(err)
	}
	if _, ok := d.GetOk(Arg_UserTags); ok {
		if route.Crn != nil {
			oldList, newList := d.GetChange(Arg_UserTags)
			err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, string(*route.Crn), "", UserTagType)
			if err != nil {
				log.Printf("Error on update of pi network security group (%s) pi_user_tags during creation: %s", *route.ID, err)
			}
		}
	}

	routeID := *route.ID
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, routeID))

	return resourceIBMPIRouteRead(ctx, d, meta)
}

func resourceIBMPIRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, routeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	routeClient := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)
	route, err := routeClient.Get(routeID)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.Set(Arg_CloudInstanceID, cloudInstanceID)
	d.Set(Arg_Action, route.Action)
	d.Set(Arg_Advertise, route.Advertise)
	if route.Crn != nil {
		d.Set(Attr_CRN, route.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(*route.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of pi route (%s) pi_user_tags: %s", *route.ID, err)
		}
		d.Set(Arg_UserTags, tags)
	}
	d.Set(Arg_Destination, route.Destination)
	d.Set(Arg_DestinationType, route.DestinationType)
	d.Set(Arg_Enabled, route.Enabled)
	d.Set(Attr_RouteID, route.ID)
	d.Set(Arg_Name, route.Name)
	d.Set(Arg_NextHop, route.NextHop)
	d.Set(Arg_NextHopType, route.NextHopType)
	d.Set(Attr_State, route.State)

	return nil
}

func resourceIBMPIRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID, routeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	routeClient := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)
	body := &models.RouteUpdate{}

	if d.HasChange(Arg_Action) {
		action := d.Get(Arg_Action).(string)
		body.Action = action
	}

	if d.HasChange(Arg_Advertise) {
		advertise := d.Get(Arg_Advertise).(string)
		body.Advertise = advertise
	}

	if d.HasChange(Arg_Destination) {
		destination := d.Get(Arg_Destination).(string)
		body.Destination = destination
	}

	if d.HasChange(Arg_DestinationType) {
		destinationType := d.Get(Arg_DestinationType).(string)
		body.DestinationType = destinationType
	}

	if d.HasChange(Arg_Enabled) {
		enabled := d.Get(Arg_Enabled).(bool)
		body.Enabled = &enabled
	}

	if d.HasChange(Arg_Name) {
		name := d.Get(Arg_Name).(string)
		body.Name = name
	}

	if d.HasChange(Arg_NextHop) {
		nextHop := d.Get(Arg_NextHop).(string)
		body.NextHopType = nextHop
	}

	if d.HasChange(Arg_NextHopType) {
		nextHopType := d.Get(Arg_NextHopType).(string)
		body.NextHopType = nextHopType
	}

	_, err = routeClient.Update(routeID, body)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange(Arg_UserTags) {
		if crn, ok := d.GetOk(Attr_CRN); ok {
			oldList, newList := d.GetChange(Arg_UserTags)
			err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, crn.(string), "", UserTagType)
			if err != nil {
				log.Printf("Error on update of pi network security group (%s) pi_user_tags: %s", routeID, err)
			}
		}
	}
	return resourceIBMPIRouteRead(ctx, d, meta)
}

func resourceIBMPIRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID, routeID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	routeClient := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)
	err = routeClient.Delete(routeID)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = isWaitForIBMPIRouteDeleted(ctx, routeClient, routeID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForIBMPIRouteDeleted(ctx context.Context, client *instance.IBMPIRouteClient, id string, timeout time.Duration) (interface{}, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Found},
		Target:     []string{State_NotFound},
		Refresh:    isIBMPIRouteRefreshDeleteFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIRouteRefreshDeleteFunc(client *instance.IBMPIRouteClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		route, err := client.Get(id)
		if err != nil && strings.Contains(err.Error(), NotFound) {
			return route, State_NotFound, nil
		}
		return route, State_Found, nil
	}
}
