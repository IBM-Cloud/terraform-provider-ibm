package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

const (
	isVPCRouteName            = "name"
	isVPCRouteState           = "status"
	isVPCRouteNextHop         = "next_hop"
	isVPCRouteDestinationCIDR = "destination"
	isVPCRouteLocation        = "zone"
	isVPCRouteVPCID           = "vpc"

	isRouteStatusPending  = "pending"
	isRouteStatusUpdating = "updating"
	isRouteStatusStable   = "stable"
	isRouteStatusFailed   = "failed"

	isRouteStatusDeleting = "deleting"
	isRouteStatusDeleted  = "deleted"
)

func resourceIBMISVpcRoute() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVpcRouteCreate,
		Read:     resourceIBMISVpcRouteRead,
		Update:   resourceIBMISVpcRouteUpdate,
		Delete:   resourceIBMISVpcRouteDelete,
		Exists:   resourceIBMISVpcRouteExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isVPCRouteName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			isVPCRouteLocation: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPCRouteDestinationCIDR: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPCRouteState: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCRouteVPCID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPCRouteNextHop: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIBMISVpcRouteCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)

	routeName := d.Get(isVPCRouteName).(string)
	zoneName := d.Get(isVPCRouteLocation).(string)
	cidr := d.Get(isVPCRouteDestinationCIDR).(string)
	vpcID := d.Get(isVPCRouteVPCID).(string)
	nextHop := d.Get(isVPCRouteNextHop).(string)

	routeTemp := &models.RouteTemplate{
		Destination: models.CIDR(cidr),
		Name:        routeName,
		NextHop: models.RouteNextHopIP{
			IP: models.IP{
				Address: nextHop,
			},
		},
		Zone: &models.ZoneIdentityByName{
			Name: &zoneName,
		},
	}

	route, err := vpcClient.CreateRoute(routeTemp, vpcID)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", vpcID, *route.ID))

	_, err = isWaitForRouteStable(d, meta)
	if err != nil {
		return err
	}

	return resourceIBMISVpcRouteRead(d, meta)

}

func resourceIBMISVpcRouteRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	routeID := parts[1]
	route, err := vpcClient.GetRoute(vpcID, routeID)
	if err != nil {
		return err
	}

	d.Set(isVPCRouteName, route.Name)
	if route.Zone != nil {
		d.Set(isVPCRouteLocation, route.Zone.Name)
	}

	d.Set(isVPCRouteDestinationCIDR, route.Destination)
	d.Set(isVPCAddressPrefixVPCID, vpcID)
	d.Set(isVPCRouteNextHop, route.NextHop)
	d.Set(isVPCRouteState, route.LifecycleState)

	return nil
}

func resourceIBMISVpcRouteUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)
	hasChanged := false

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	routeID := parts[1]
	route, err := vpcClient.GetRoute(vpcID, routeID)
	if err != nil {
		return err
	}

	if d.HasChange(isVPCRouteName) {
		*route.Name = (d.Get(isVPCRouteName).(string))
		hasChanged = true
	}
	routePatch := &models.RoutePatch{
		Name: *route.Name,
	}

	if hasChanged {
		_, err = vpcClient.UpdateRoute(routePatch, vpcID, routeID)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVpcRouteRead(d, meta)
}

func resourceIBMISVpcRouteDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	routeID := parts[1]

	err = vpcClient.DeleteRoute(vpcID, routeID)
	if err != nil {
		return err
	}
	_, err = isWaitForVPCRouteDeleted(vpcClient, vpcID, routeID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISVpcRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpcClient := network.NewVPCClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	vpcID := parts[0]
	routeID := parts[1]
	_, err = vpcClient.GetRoute(vpcID, routeID)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func isWaitForRouteStable(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpcClient := network.NewVPCClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return nil, err
	}
	vpcID := parts[0]
	routeID := parts[1]
	stateConf := &resource.StateChangeConf{
		Pending: []string{isRouteStatusPending, isRouteStatusUpdating},
		Target:  []string{isRouteStatusStable, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			route, err := vpcClient.GetRoute(vpcID, routeID)
			if err != nil {
				return nil, "", err
			}
			return route, string(route.LifecycleState), nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isWaitForVPCRouteDeleted(vpcClient *network.VPCClient, vpcID, routeID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC Route (%s) to be deleted.", routeID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", isRouteStatusDeleting},
		Target:  []string{isRouteStatusDeleted, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			route, err := vpcClient.GetRoute(vpcID, routeID) //Only in case there's a rias error with code "not found", resource is deleted, all other cases we keep attempting to delete
			if err == nil {
				if route.LifecycleState == isRouteStatusFailed {
					return route, isRouteStatusFailed, fmt.Errorf("The VPC route %s failed to delete: %v", routeID, err)
				}
				return route, isRouteStatusDeleting, nil
			}

			iserror, ok := err.(iserrors.RiaasError)
			if ok {
				log.Printf("[DEBUG] %s", iserror.Error())
				if len(iserror.Payload.Errors) == 1 &&
					iserror.Payload.Errors[0].Code == "not_found" {
					return route, isRouteStatusDeleted, nil
				}
			}
			return nil, isRouteStatusDeleting, err
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
