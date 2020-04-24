package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isVPCRouteName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
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
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	routeName := d.Get(isVPCRouteName).(string)
	zoneName := d.Get(isVPCRouteLocation).(string)
	cidr := d.Get(isVPCRouteDestinationCIDR).(string)
	vpcID := d.Get(isVPCRouteVPCID).(string)
	nextHop := d.Get(isVPCRouteNextHop).(string)
	if userDetails.generation == 1 {
		err := classicVpcRouteCreate(d, meta, routeName, zoneName, cidr, vpcID, nextHop)
		if err != nil {
			return err
		}
	} else {
		err := vpcRouteCreate(d, meta, routeName, zoneName, cidr, vpcID, nextHop)
		if err != nil {
			return err
		}
	}
	return resourceIBMISVpcRouteRead(d, meta)
}

func classicVpcRouteCreate(d *schema.ResourceData, meta interface{}, routeName, zoneName, cidr, vpcID, nextHop string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	createRouteOptions := &vpcclassicv1.CreateVpcRouteOptions{
		VpcID:       &vpcID,
		Destination: &cidr,
		Name:        &routeName,
		NextHop: &vpcclassicv1.RouteNextHopPrototype{
			Address: &nextHop,
		},
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: &zoneName,
		},
	}
	route, response, err := sess.CreateVpcRoute(createRouteOptions)
	if err != nil {
		return fmt.Errorf("Error while creating VPC Route %s\n%s", err, response)
	}
	routeID := *route.ID

	d.SetId(fmt.Sprintf("%s/%s", vpcID, routeID))

	_, err = isWaitForClassicRouteStable(sess, d, vpcID, routeID)
	if err != nil {
		return err
	}
	return nil
}

func vpcRouteCreate(d *schema.ResourceData, meta interface{}, routeName, zoneName, cidr, vpcID, nextHop string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	createRouteOptions := &vpcv1.CreateVpcRouteOptions{
		VpcID:       &vpcID,
		Destination: &cidr,
		Name:        &routeName,
		NextHop: &vpcv1.RouteNextHopPrototype{
			Address: &nextHop,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zoneName,
		},
	}
	route, response, err := sess.CreateVpcRoute(createRouteOptions)
	if err != nil {
		return fmt.Errorf("Error while creating VPC Route err %s\n%s", err, response)
	}
	routeID := *route.ID

	d.SetId(fmt.Sprintf("%s/%s", vpcID, routeID))

	_, err = isWaitForRouteStable(sess, d, vpcID, routeID)
	if err != nil {
		return err
	}
	return nil
}

func isWaitForClassicRouteStable(sess *vpcclassicv1.VpcClassicV1, d *schema.ResourceData, vpcID, routeID string) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending: []string{isRouteStatusPending, isRouteStatusUpdating},
		Target:  []string{isRouteStatusStable, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			getVpcRouteOptions := &vpcclassicv1.GetVpcRouteOptions{
				VpcID: &vpcID,
				ID:    &routeID,
			}
			route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
			if err != nil {
				return route, "", fmt.Errorf("Error Getting VPC Route: %s\n%s", err, response)
			}

			if *route.LifecycleState == "stable" || *route.LifecycleState == "failed" {
				return route, *route.LifecycleState, nil
			}
			return route, *route.LifecycleState, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isWaitForRouteStable(sess *vpcv1.VpcV1, d *schema.ResourceData, vpcID, routeID string) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending: []string{isRouteStatusPending, isRouteStatusUpdating},
		Target:  []string{isRouteStatusStable, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			getVpcRouteOptions := &vpcv1.GetVpcRouteOptions{
				VpcID: &vpcID,
				ID:    &routeID,
			}
			route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
			if err != nil {
				return route, "", fmt.Errorf("Error Getting VPC Route: %s\n%s", err, response)
			}

			if *route.LifecycleState == "stable" || *route.LifecycleState == "failed" {
				return route, *route.LifecycleState, nil
			}
			return route, *route.LifecycleState, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMISVpcRouteRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	routeID := parts[1]
	if userDetails.generation == 1 {
		err := classicVpcRouteGet(d, meta, vpcID, routeID)
		if err != nil {
			return err
		}
	} else {
		err := vpcRouteGet(d, meta, vpcID, routeID)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVpcRouteGet(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getVpcRouteOptions := &vpcclassicv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Route (%s): %s\n%s", routeID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set("id", *route.ID)
	d.Set(isVPCRouteName, route.Name)
	if route.Zone != nil {
		d.Set(isVPCRouteLocation, *route.Zone.Name)
	}
	d.Set(isVPCRouteDestinationCIDR, *route.Destination)
	nexthop := route.NextHop.(*vpcclassicv1.RouteNextHop)
	d.Set(isVPCRouteNextHop, *nexthop.Address)
	d.Set(isVPCRouteState, *route.LifecycleState)
	return nil
}

func vpcRouteGet(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getVpcRouteOptions := &vpcv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Route (%s): %s\n%s", routeID, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set("id", *route.ID)
	d.Set(isVPCRouteName, route.Name)
	if route.Zone != nil {
		d.Set(isVPCRouteLocation, *route.Zone.Name)
	}
	d.Set(isVPCRouteDestinationCIDR, *route.Destination)
	nexthop := route.NextHop.(*vpcv1.RouteNextHop)
	d.Set(isVPCRouteNextHop, *nexthop.Address)
	d.Set(isVPCRouteState, *route.LifecycleState)

	return nil
}

func resourceIBMISVpcRouteUpdate(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := ""
	hasChanged := false

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	vpcID := parts[0]
	routeID := parts[1]
	if d.HasChange(isVPCRouteName) {
		name = d.Get(isVPCRouteName).(string)
		hasChanged = true
	}

	if userDetails.generation == 1 {
		err := classicVpcRouteUpdate(d, meta, vpcID, routeID, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := vpcRouteUpdate(d, meta, vpcID, routeID, name, hasChanged)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVpcRouteRead(d, meta)
}

func classicVpcRouteUpdate(d *schema.ResourceData, meta interface{}, vpcID, routeID, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if hasChanged {
		updateVpcRouteOptions := &vpcclassicv1.UpdateVpcRouteOptions{
			VpcID: &vpcID,
			ID:    &routeID,
			Name:  &name,
		}
		_, response, err := sess.UpdateVpcRoute(updateVpcRouteOptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC Route: %s\n%s", err, response)
		}
	}
	return nil
}

func vpcRouteUpdate(d *schema.ResourceData, meta interface{}, vpcID, routeID, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if hasChanged {
		updateVpcRouteOptions := &vpcv1.UpdateVpcRouteOptions{
			VpcID: &vpcID,
			ID:    &routeID,
			Name:  &name,
		}
		_, response, err := sess.UpdateVpcRoute(updateVpcRouteOptions)
		if err != nil {
			return fmt.Errorf("Error Updating VPC Route: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISVpcRouteDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	vpcID := parts[0]
	routeID := parts[1]
	if userDetails.generation == 1 {
		err := classicVpcRouteDelete(d, meta, vpcID, routeID)
		if err != nil {
			return err
		}
	} else {
		err := vpcRouteDelete(d, meta, vpcID, routeID)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil
}

func classicVpcRouteDelete(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getVpcRouteOptions := &vpcclassicv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	_, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Route (%s): %s\n%s", routeID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	deleteRouteOptions := &vpcclassicv1.DeleteVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	response, err = sess.DeleteVpcRoute(deleteRouteOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting VPC Route: %s\n%s", err, response)
	}
	_, err = isWaitForClassicVPCRouteDeleted(sess, vpcID, routeID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func vpcRouteDelete(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getVpcRouteOptions := &vpcv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	_, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting VPC Route (%s): %s\n%s", routeID, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	deleteRouteOptions := &vpcv1.DeleteVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	response, err = sess.DeleteVpcRoute(deleteRouteOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting VPC Route: %s\n%s", err, response)
	}
	_, err = isWaitForVPCRouteDeleted(sess, vpcID, routeID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForClassicVPCRouteDeleted(sess *vpcclassicv1.VpcClassicV1, vpcID, routeID string, timeout time.Duration) (interface{}, error) {

	log.Printf("Waiting for VPC Route (%s) to be deleted.", routeID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", isRouteStatusDeleting},
		Target:  []string{isRouteStatusDeleted, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			getVpcRouteOptions := &vpcclassicv1.GetVpcRouteOptions{
				VpcID: &vpcID,
				ID:    &routeID,
			}
			route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
			if err != nil && response.StatusCode != 404 {
				return route, isRouteStatusDeleting, fmt.Errorf("The VPC route %s failed to delete: %s\n%s", routeID, err, response)
			}
			if response.StatusCode == 404 {
				return route, isRouteStatusDeleted, nil
			}
			return route, isRouteStatusDeleting, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isWaitForVPCRouteDeleted(sess *vpcv1.VpcV1, vpcID, routeID string, timeout time.Duration) (interface{}, error) {

	log.Printf("Waiting for VPC Route (%s) to be deleted.", routeID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", isRouteStatusDeleting},
		Target:  []string{isRouteStatusDeleted, isRouteStatusFailed},
		Refresh: func() (interface{}, string, error) {
			getVpcRouteOptions := &vpcv1.GetVpcRouteOptions{
				VpcID: &vpcID,
				ID:    &routeID,
			}
			route, response, err := sess.GetVpcRoute(getVpcRouteOptions)
			if err != nil && response.StatusCode != 404 {
				return route, isRouteStatusDeleting, fmt.Errorf("The VPC route %s failed to delete: %s\n%s", routeID, err, response)
			}
			if response.StatusCode == 404 {
				return route, isRouteStatusDeleted, nil
			}
			return route, isRouteStatusDeleting, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMISVpcRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	vpcID := parts[0]
	routeID := parts[1]
	if userDetails.generation == 1 {
		err := classicVpcRouteExists(d, meta, vpcID, routeID)
		if err != nil {
			return false, err
		}
	} else {
		err := vpcRouteExists(d, meta, vpcID, routeID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicVpcRouteExists(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getVpcRouteOptions := &vpcclassicv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	_, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting VPC Route: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func vpcRouteExists(d *schema.ResourceData, meta interface{}, vpcID, routeID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getVpcRouteOptions := &vpcv1.GetVpcRouteOptions{
		VpcID: &vpcID,
		ID:    &routeID,
	}
	_, response, err := sess.GetVpcRoute(getVpcRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting VPC Route: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}
