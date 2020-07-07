package ibm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isFloatingIPAddress       = "address"
	isFloatingIPName          = "name"
	isFloatingIPStatus        = "status"
	isFloatingIPZone          = "zone"
	isFloatingIPTarget        = "target"
	isFloatingIPResourceGroup = "resource_group"
	isFloatingIPTags          = "tags"

	isFloatingIPPending   = "pending"
	isFloatingIPAvailable = "available"
	isFloatingIPDeleting  = "deleting"
	isFloatingIPDeleted   = "done"
)

func resourceIBMISFloatingIP() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISFloatingIPCreate,
		Read:     resourceIBMISFloatingIPRead,
		Update:   resourceIBMISFloatingIPUpdate,
		Delete:   resourceIBMISFloatingIPDelete,
		Exists:   resourceIBMISFloatingIPExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isFloatingIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP address",
			},

			isFloatingIPName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "Name of the floating IP",
			},

			isFloatingIPStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP status",
			},

			isFloatingIPZone: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPTarget},
				Description:   "Zone name",
			},

			isFloatingIPTarget: {
				Type:          schema.TypeString,
				ForceNew:      false,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPZone},
				Description:   "Target info",
			},

			isFloatingIPResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group info",
			},

			isFloatingIPTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "Floating IP tags",
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISFloatingIPCreate(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := d.Get(isFloatingIPName).(string)
	if userDetails.generation == 1 {
		err := classicFipCreate(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := fipCreate(d, meta, name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISFloatingIPRead(d, meta)
}
func classicFipCreate(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	floatingIPPrototype := &vpcclassicv1.FloatingIPPrototype{
		Name: &name,
	}
	zone, target := "", ""

	if zn, ok := d.GetOk(isFloatingIPZone); ok {
		zone = zn.(string)
		floatingIPPrototype.Zone = &vpcclassicv1.ZoneIdentity{
			Name: &zone,
		}
	}

	if tgt, ok := d.GetOk(isFloatingIPTarget); ok {
		target = tgt.(string)
		floatingIPPrototype.Target = &vpcclassicv1.NetworkInterfaceIdentity{
			ID: &target,
		}
	}

	if zone == "" && target == "" {
		return fmt.Errorf("%s or %s need to be provided", isFloatingIPZone, isFloatingIPTarget)
	}

	options := &vpcclassicv1.ReserveFloatingIpOptions{
		FloatingIPPrototype: floatingIPPrototype,
	}
	floatingip, response, err := sess.ReserveFloatingIp(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Floating IP err %s\n%s", err, response)
	}
	d.SetId(*floatingip.ID)
	log.Printf("[INFO] Floating IP : %s[%s]", *floatingip.ID, *floatingip.Address)
	_, err = isWaitForClassicInstanceFloatingIP(sess, d.Id(), d)
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isFloatingIPTags); ok || v != "" {
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *floatingip.Crn)
		if err != nil {
			log.Printf(
				"Error on create of vpc Floating IP (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func fipCreate(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	floatingIPPrototype := &vpcv1.FloatingIPPrototype{
		Name: &name,
	}
	zone, target := "", ""
	if zn, ok := d.GetOk(isFloatingIPZone); ok {
		zone = zn.(string)
		floatingIPPrototype.Zone = &vpcv1.ZoneIdentity{
			Name: &zone,
		}
	}

	if tgt, ok := d.GetOk(isFloatingIPTarget); ok {
		target = tgt.(string)
		floatingIPPrototype.Target = &vpcv1.NetworkInterfaceIdentity{
			ID: &target,
		}
	}

	if zone == "" && target == "" {
		return fmt.Errorf("%s or %s need to be provided", isFloatingIPZone, isFloatingIPTarget)
	}

	if rgrp, ok := d.GetOk(isFloatingIPResourceGroup); ok {
		rg := rgrp.(string)
		floatingIPPrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	options := &vpcv1.ReserveFloatingIpOptions{
		FloatingIPPrototype: floatingIPPrototype,
	}

	floatingip, response, err := sess.ReserveFloatingIp(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Floating IP err %s\n%s", err, response)
	}
	d.SetId(*floatingip.ID)
	log.Printf("[INFO] Floating IP : %s[%s]", *floatingip.ID, *floatingip.Address)
	_, err = isWaitForInstanceFloatingIP(sess, d.Id(), d)
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isFloatingIPTags); ok || v != "" {
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *floatingip.Crn)
		if err != nil {
			log.Printf(
				"Error on create of vpc Floating IP (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISFloatingIPRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	if userDetails.generation == 1 {
		err := classicFipGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := fipGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicFipGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.GetFloatingIpOptions{
		ID: &id,
	}
	floatingip, response, err := sess.GetFloatingIp(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Floating IP (%s): %s\n%s", id, err, response)

	}
	d.Set(isFloatingIPName, *floatingip.Name)
	d.Set(isFloatingIPAddress, *floatingip.Address)
	d.Set(isFloatingIPStatus, *floatingip.Status)
	d.Set(isFloatingIPZone, *floatingip.Zone.Name)
	target, ok := floatingip.Target.(*vpcclassicv1.FloatingIPTarget)
	if ok {
		d.Set(isFloatingIPTarget, target.ID)
	}
	tags, err := GetTagsUsingCRN(meta, *floatingip.Crn)
	if err != nil {
		log.Printf(
			"Error on get of vpc Floating IP (%s) tags: %s", d.Id(), err)
	}
	d.Set(isFloatingIPTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/floatingIPs")
	d.Set(ResourceName, *floatingip.Name)
	d.Set(ResourceCRN, *floatingip.Crn)
	d.Set(ResourceStatus, *floatingip.Status)
	return nil
}

func fipGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetFloatingIpOptions{
		ID: &id,
	}
	floatingip, response, err := sess.GetFloatingIp(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Floating IP (%s): %s\n%s", id, err, response)

	}
	d.Set(isFloatingIPName, *floatingip.Name)
	d.Set(isFloatingIPAddress, *floatingip.Address)
	d.Set(isFloatingIPStatus, *floatingip.Status)
	d.Set(isFloatingIPZone, *floatingip.Zone.Name)
	target, ok := floatingip.Target.(*vpcv1.FloatingIPTarget)
	if ok {
		d.Set(isFloatingIPTarget, target.ID)
	}
	tags, err := GetTagsUsingCRN(meta, *floatingip.Crn)
	if err != nil {
		log.Printf(
			"Error on get of vpc Floating IP (%s) tags: %s", d.Id(), err)
	}
	d.Set(isFloatingIPTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/floatingIPs")
	d.Set(ResourceName, *floatingip.Name)
	d.Set(ResourceCRN, *floatingip.Crn)
	d.Set(ResourceStatus, *floatingip.Status)
	if floatingip.ResourceGroup != nil {
		d.Set(ResourceGroupName, *floatingip.ResourceGroup.Name)
		d.Set(isFloatingIPResourceGroup, *floatingip.ResourceGroup.ID)
	}
	return nil
}

func resourceIBMISFloatingIPUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	if userDetails.generation == 1 {
		err := classicFipUpdate(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := fipUpdate(d, meta, id)
		if err != nil {
			return err
		}
	}
	return resourceIBMISFloatingIPRead(d, meta)
}

func classicFipUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isFloatingIPTags) {
		options := &vpcclassicv1.GetFloatingIpOptions{
			ID: &id,
		}
		fip, response, err := sess.GetFloatingIp(options)
		if err != nil {
			return fmt.Errorf("Error getting Floating IP: %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *fip.Crn)
		if err != nil {
			log.Printf(
				"Error on update of vpc Floating IP (%s) tags: %s", id, err)
		}
	}
	hasChanged := false
	options := &vpcclassicv1.UpdateFloatingIpOptions{
		ID: &id,
	}
	if d.HasChange(isFloatingIPName) {
		name := d.Get(isFloatingIPName).(string)
		options.Name = &name
		hasChanged = true
	}

	if d.HasChange(isFloatingIPTarget) {
		target := d.Get(isFloatingIPTarget).(string)
		options.Target = &vpcclassicv1.NetworkInterfaceIdentity{
			ID: &target,
		}
		hasChanged = true
	}
	if hasChanged {
		_, response, err := sess.UpdateFloatingIp(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc Floating IP: %s\n%s", err, response)
		}
	}
	return nil
}

func fipUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isFloatingIPTags) {
		options := &vpcv1.GetFloatingIpOptions{
			ID: &id,
		}
		fip, response, err := sess.GetFloatingIp(options)
		if err != nil {
			return fmt.Errorf("Error getting Floating IP: %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *fip.Crn)
		if err != nil {
			log.Printf(
				"Error on update of vpc Floating IP (%s) tags: %s", id, err)
		}
	}
	hasChanged := false
	options := &vpcv1.UpdateFloatingIpOptions{
		ID: &id,
	}
	if d.HasChange(isFloatingIPName) {
		name := d.Get(isFloatingIPName).(string)
		options.Name = &name
		hasChanged = true
	}

	if d.HasChange(isFloatingIPTarget) {
		target := d.Get(isFloatingIPTarget).(string)
		options.Target = &vpcv1.NetworkInterfaceIdentity{
			ID: &target,
		}
		hasChanged = true
	}
	if hasChanged {
		_, response, err := sess.UpdateFloatingIp(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc Floating IP: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISFloatingIPDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicFipDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := fipDelete(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicFipDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getFloatingIpOptions := &vpcclassicv1.GetFloatingIpOptions{
		ID: &id,
	}
	_, response, err := sess.GetFloatingIp(getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}

		return fmt.Errorf("Error Getting Floating IP (%s): %s\n%s", id, err, response)

	}

	options := &vpcclassicv1.ReleaseFloatingIpOptions{
		ID: &id,
	}
	response, err = sess.ReleaseFloatingIp(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Floating IP : %s\n%s", err, response)
	}
	_, err = isWaitForClassicFloatingIPDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func fipDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getFloatingIpOptions := &vpcv1.GetFloatingIpOptions{
		ID: &id,
	}
	_, response, err := sess.GetFloatingIp(getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}

		return fmt.Errorf("Error Getting Floating IP (%s): %s\n%s", id, err, response)
	}

	options := &vpcv1.ReleaseFloatingIpOptions{
		ID: &id,
	}
	response, err = sess.ReleaseFloatingIp(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Floating IP : %s\n%s", err, response)
	}
	_, err = isWaitForFloatingIPDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISFloatingIPExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()

	if userDetails.generation == 1 {
		exists, err := classicFipExists(d, meta, id)
		return exists, err
	} else {
		exists, err := fipExists(d, meta, id)
		return exists, err
	}
}

func classicFipExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return false, err
	}
	getFloatingIpOptions := &vpcclassicv1.GetFloatingIpOptions{
		ID: &id,
	}
	_, response, err := sess.GetFloatingIp(getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting floating IP: %s\n%s", err, response)
	}
	return true, nil
}

func fipExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getFloatingIpOptions := &vpcv1.GetFloatingIpOptions{
		ID: &id,
	}
	_, response, err := sess.GetFloatingIp(getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting floating IP: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForClassicFloatingIPDeleted(fip *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for FloatingIP (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isFloatingIPPending, isFloatingIPDeleting},
		Target:     []string{"", isFloatingIPDeleted},
		Refresh:    isClassicFloatingIPDeleteRefreshFunc(fip, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicFloatingIPDeleteRefreshFunc(fip *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getfipoptions := &vpcclassicv1.GetFloatingIpOptions{
			ID: &id,
		}
		FloatingIP, response, err := fip.GetFloatingIp(getfipoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return FloatingIP, isFloatingIPDeleted, nil
			}
			return FloatingIP, "", fmt.Errorf("Error Getting Floating IP: %s\n%s", err, response)
		}
		return FloatingIP, isFloatingIPDeleting, err
	}
}

func isWaitForFloatingIPDeleted(fip *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for FloatingIP (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isFloatingIPPending, isFloatingIPDeleting},
		Target:     []string{"", isFloatingIPDeleted},
		Refresh:    isFloatingIPDeleteRefreshFunc(fip, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isFloatingIPDeleteRefreshFunc(fip *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		getfipoptions := &vpcv1.GetFloatingIpOptions{
			ID: &id,
		}
		FloatingIP, response, err := fip.GetFloatingIp(getfipoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return FloatingIP, isFloatingIPDeleted, nil
			}
			return FloatingIP, "", fmt.Errorf("Error Getting Floating IP: %s\n%s", err, response)
		}
		return FloatingIP, isFloatingIPDeleting, err
	}
}

func isWaitForClassicInstanceFloatingIP(floatingipC *vpcclassicv1.VpcClassicV1, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for floating IP (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isFloatingIPPending},
		Target:     []string{isFloatingIPAvailable, ""},
		Refresh:    isClassicInstanceFloatingIPRefreshFunc(floatingipC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isClassicInstanceFloatingIPRefreshFunc(floatingipC *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getfipoptions := &vpcclassicv1.GetFloatingIpOptions{
			ID: &id,
		}
		instance, response, err := floatingipC.GetFloatingIp(getfipoptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Floating IP for the instance: %s\n%s", err, response)
		}

		if *instance.Status == "available" {
			return instance, isFloatingIPAvailable, nil
		}

		return instance, isFloatingIPPending, nil
	}
}

func isWaitForInstanceFloatingIP(floatingipC *vpcv1.VpcV1, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for floating IP (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isFloatingIPPending},
		Target:     []string{isFloatingIPAvailable, ""},
		Refresh:    isInstanceFloatingIPRefreshFunc(floatingipC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceFloatingIPRefreshFunc(floatingipC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getfipoptions := &vpcv1.GetFloatingIpOptions{
			ID: &id,
		}
		instance, response, err := floatingipC.GetFloatingIp(getfipoptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Floating IP for the instance: %s\n%s", err, response)
		}

		if *instance.Status == "available" {
			return instance, isFloatingIPAvailable, nil
		}

		return instance, isFloatingIPPending, nil
	}
}
