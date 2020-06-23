package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
)
const (
	dlGatewaysVirtualConnections = "gateway_vcs"
	dlVCNetworkAccount           = "network_account"
	dlVCNetworkId                = "network_id"
	dlVCName                     = "name"
	dlVCType                     = "type"
	dlVCCreatedAt                = "created_at"
	dlVCStatus                   = "status"
	dlGatewayId                  = "gateway"
	ID                           = "id"
)

func resourceIBMDLGatewayVC() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMdlGatewayVCCreate,
		Read:     resourceIBMdlGatewayVCRead,
		Delete:   resourceIBMdlGatewayVCDelete,
		Exists:   resourceIBMdlGatewayVCExists,
		Update:   resourceIBMdlGatewayVCUpdate,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlVCType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlVCName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			dlVCNetworkId: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			dlVCCreatedAt: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlVCStatus: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: false,
			},

			dlVCNetworkAccount: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMdlGatewayVCCreate(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	createGatewayVCOptions := &directlinkapisv1.CreateGatewayVirtualConnectionOptions{}

	gatewayId := d.Get(dlGatewayId).(string)
	createGatewayVCOptions.SetGatewayID(gatewayId)
	vcName := d.Get(dlVCName).(string)
	createGatewayVCOptions.SetName(vcName)
	vcType := d.Get(dlVCType).(string)
	createGatewayVCOptions.SetType(vcType)

	if _, ok := d.GetOk(dlVCNetworkId); ok {
		vcNetworkId := d.Get(dlVCNetworkId).(string)
		createGatewayVCOptions.SetNetworkID(vcNetworkId)
	}

	gatewayVC, response, err := directLink.CreateGatewayVirtualConnection(createGatewayVCOptions)
	if err != nil {
		log.Printf("[DEBUG] Create Direct Link Gateway (Dedicated) Virtual connection err %s\n%s", err, response)
		return err
	}
	d.SetId(*gatewayVC.ID)

	log.Printf("[INFO] Created Direct Link Gateway (Dedicated Template) Virtual connection : %s", *gatewayVC.ID)

	if err != nil {
		log.Printf("Error creating Direct Link Gateway (Dedicated Template) Virtual Connection :%s", response)
		return err
	}

	return resourceIBMdlGatewayVCRead(d, meta)
}

func resourceIBMdlGatewayVCRead(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()

	getGatewayVirtualConnectionOptions := &directlinkapisv1.GetGatewayVirtualConnectionOptions{}
	gatewayId := d.Get(dlGatewayId).(string)
	gatewayVCId := ID

	getGatewayVirtualConnectionOptions.SetGatewayID(gatewayId)
	getGatewayVirtualConnectionOptions.SetID(gatewayVCId)
	instance, _, err := directLink.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptions)
	if err != nil {
		return err
	}
	if instance.ID != nil {
		d.SetId(*instance.ID)
	}
	if instance.Name != nil {
		d.Set(dlVCName, *instance.Name)
	}
	if instance.Type != nil {
		d.Set(dlVCType, *instance.Type)
	}
	if instance.NetworkAccount != nil {
		d.Set(dlVCNetworkAccount, *instance.NetworkAccount)
	}
	if instance.NetworkID != nil {
		d.Set(dlVCNetworkId, *instance.NetworkID)
	}
	if instance.CreatedAt != nil {
		d.Set(dlVCCreatedAt, instance.CreatedAt.String())
	}
	if instance.Status != nil {
		d.Set(dlVCStatus, *instance.ID)
	}
	return nil
}

func resourceIBMdlGatewayVCUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	gatewayId := d.Get(dlGatewayId).(string)

	getVCOptions := &directlinkapisv1.GetGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	getVCOptions.SetGatewayID(gatewayId)
	_, detail, err := directLink.GetGatewayVirtualConnection(getVCOptions)

	if err != nil {
		log.Printf("Error fetching Direct Link Gateway (Dedicated Template) Virtual Connection:%s", detail)
		return err
	}

	updateGatewayVCOptions := &directlinkapisv1.UpdateGatewayVirtualConnectionOptions{}
	updateGatewayVCOptions.ID = &ID
	updateGatewayVCOptions.SetGatewayID(gatewayId)
	if d.HasChange(dlName) {
		if d.Get(dlName) != nil {
			name := d.Get(dlName).(string)
			updateGatewayVCOptions.Name = &name
		}
	}
	if d.HasChange(dlVCStatus) {
		if d.Get(dlVCStatus) != nil {
			status := d.Get(dlVCStatus).(string)
			updateGatewayVCOptions.SetStatus(status)
		}
	}

	_, response, err := directLink.UpdateGatewayVirtualConnection(updateGatewayVCOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway (Dedicated) Virtual Connection err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlGatewayVCRead(d, meta)
}

func resourceIBMdlGatewayVCDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	gatewayId := d.Get(dlGatewayId).(string)

	delVCOptions := &directlinkapisv1.DeleteGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	delVCOptions.SetGatewayID(gatewayId)
	response, err := directLink.DeleteGatewayVirtualConnection(delVCOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Direct Link Gateway (Dedicated Template) Virtual Connection: %s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMdlGatewayVCExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return false, err
	}

	ID := d.Id()
	gatewayId := d.Get(dlGatewayId).(string)

	getVCOptions := &directlinkapisv1.GetGatewayVirtualConnectionOptions{
		ID: &ID,
	}
	getVCOptions.SetGatewayID(gatewayId)
	_, response, err := directLink.GetGatewayVirtualConnection(getVCOptions)
	if err != nil {
		return false, fmt.Errorf("Error Getting Direct Link Gateway (Dedicated Template) Virtual Connection: %s\n%s", err, response)
	}

	if response.StatusCode == 404 {
		d.SetId("")
		return false, nil
	}
	return true, nil
}
