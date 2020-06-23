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
	dlLoaRejectReason              = "loa_reject_reason"
	dlCustomerName                 = "customer_name"
	dlCarrierName                  = "carrier_name"
	dlResourceGroup                = "resource_group"
	dlBgpAsn                       = "bgp_asn"
	dlBgpBaseCidr                  = "bgp_base_cidr"
	dlBgpCerCidr                   = "bgp_cer_cidr"
	dlBgpIbmCidr                   = "bgp_ibm_cidr"
	dlCrossConnectRouter           = "cross_connect_router"
	dlGlobal                       = "global"
	dlLocationName                 = "location_name"
	dlName                         = "name"
	dlSpeedMbps                    = "speed_mbps"
	dlOperationalStatus            = "operational_status"
	dlBgpStatus                    = "bgp_status"
	dlLinkStatus                   = "link_status"
	dlType                         = "type"
	dlCrn                          = "crn"
	dlCreatedAt                    = "created_at"
	dlMetered                      = "metered"
	dlLocationDisplayName          = "location_display_name"
	dlBgpIbmAsn                    = "bgp_ibm_asn"
	dlCompletionNoticeRejectReason = "completion_notice_reject_reason"
	dlDedicatedHostingID           = "dedicated_hosting_id"
	dlPort                         = "port"
	dlProviderAPIManaged           = "provider_api_managed"
	dlVlan                         = "vlan"
)

func resourceIBMDLGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMdlGatewayCreate,
		Read:     resourceIBMdlGatewayRead,
		Delete:   resourceIBMdlGatewayDelete,
		Exists:   resourceIBMdlGatewayExists,
		Update:   resourceIBMdlGatewayUpdate,
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
			dlBgpAsn: {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			dlBgpBaseCidr: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlCrossConnectRouter: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlGlobal: {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
			},
			dlLocationName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlMetered: {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
			},
			dlName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			dlCarrierName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlCustomerName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlSpeedMbps: {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			dlType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dlBgpCerCidr: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			dlLoaRejectReason: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: false,
			},
			dlBgpIbmCidr: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			dlOperationalStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: false,
			},
			dlResourceGroup: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			dlPort: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlProviderAPIManaged: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			dlVlan: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			dlBgpIbmAsn: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			dlBgpStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlCompletionNoticeRejectReason: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlCreatedAt: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlCrn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlLinkStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			dlLocationDisplayName: {
				Type:     schema.TypeString,
				Computed: true,
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

func directlinkClient(meta interface{}) (*directlinkapisv1.DirectLinkApisV1, error) {
	sess, err := meta.(ClientSession).DirectlinkV1API()
	return sess, err
}

func resourceIBMdlGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	createGatewayOptionsModel := &directlinkapisv1.CreateGatewayOptions{}
	gatewayTemplateModel := &directlinkapisv1.GatewayTemplateGatewayTypeDedicatedTemplate{}

	name := d.Get(dlName).(string)
	gatewayTemplateModel.Name = &name
	dtype := d.Get(dlType).(string)
	gatewayTemplateModel.Type = &dtype
	speed := int64(d.Get(dlSpeedMbps).(int))
	gatewayTemplateModel.SpeedMbps = &speed
	global := d.Get(dlGlobal).(bool)
	gatewayTemplateModel.Global = &global
	bgpAsn := int64(d.Get(dlBgpAsn).(int))
	gatewayTemplateModel.BgpAsn = &bgpAsn
	bgpBaseCidr := d.Get(dlBgpBaseCidr).(string)
	gatewayTemplateModel.BgpBaseCidr = &bgpBaseCidr
	metered := d.Get(dlMetered).(bool)
	gatewayTemplateModel.Metered = &metered
	carrierName := d.Get(dlCarrierName).(string)
	gatewayTemplateModel.CarrierName = &carrierName
	crossConnectRouter := d.Get(dlCrossConnectRouter).(string)
	gatewayTemplateModel.CrossConnectRouter = &crossConnectRouter
	locationName := d.Get(dlLocationName).(string)
	gatewayTemplateModel.LocationName = &locationName
	customerName := d.Get(dlCustomerName).(string)
	gatewayTemplateModel.CustomerName = &customerName

	if _, ok := d.GetOk(dlBgpIbmCidr); ok {
		bgpIbmCidr := d.Get(dlBgpIbmCidr).(string)
		gatewayTemplateModel.BgpIbmCidr = &bgpIbmCidr
	}
	if _, ok := d.GetOk(dlBgpCerCidr); ok {
		bgpCerCidr := d.Get(dlBgpCerCidr).(string)
		gatewayTemplateModel.BgpCerCidr = &bgpCerCidr
	}
	if _, ok := d.GetOk(dlResourceGroup); ok {
		resourceGroup := d.Get(dlResourceGroup).(string)
		gatewayTemplateModel.ResourceGroup = &directlinkapisv1.ResourceGroupIdentity{ID: &resourceGroup}
	}

	createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel

	gateway, response, err := directLink.CreateGateway(createGatewayOptionsModel)
	if err != nil {
		log.Printf("[DEBUG] Create Direct Link Gateway (Dedicated) err %s\n%s", err, response)
		return err
	}
	d.SetId(*gateway.ID)

	log.Printf("[INFO] Created Direct Link Gateway (Dedicated Template) : %s", *gateway.ID)

	return resourceIBMdlGatewayRead(d, meta)
}

func resourceIBMdlGatewayRead(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()

	getOptions := &directlinkapisv1.GetGatewayOptions{
		ID: &ID,
	}
	instance, response, err := directLink.GetGateway(getOptions)
	if err != nil {
		return fmt.Errorf("Error Getting Direct Link Gateway (Dedicated Template): %s\n%s", err, response)
	}
	if instance.ID != nil {
		d.Set("id", *instance.ID)
	}
	if instance.Name != nil {
		d.Set(dlName, *instance.Name)
	}
	if instance.Crn != nil {
		d.Set(dlCrn, *instance.Crn)
	}
	if instance.BgpAsn != nil {
		d.Set(dlBgpAsn, *instance.BgpAsn)
	}
	if instance.BgpIbmCidr != nil {
		d.Set(dlBgpIbmCidr, *instance.BgpIbmCidr)
	}
	if instance.BgpIbmAsn != nil {
		d.Set(dlBgpIbmAsn, *instance.BgpIbmAsn)
	}
	if instance.Metered != nil {
		d.Set(dlMetered, *instance.Metered)
	}
	if instance.CrossConnectRouter != nil {
		d.Set(dlCrossConnectRouter, *instance.CrossConnectRouter)
	}
	if instance.BgpBaseCidr != nil {
		d.Set(dlBgpBaseCidr, *instance.BgpBaseCidr)
	}
	if instance.BgpCerCidr != nil {
		d.Set(dlBgpCerCidr, *instance.BgpCerCidr)
	}
	if instance.ProviderApiManaged != nil {
		d.Set(dlProviderAPIManaged, *instance.ProviderApiManaged)
	}
	if instance.Type != nil {
		d.Set(dlType, *instance.Type)
	}
	if instance.SpeedMbps != nil {
		d.Set(dlSpeedMbps, *instance.SpeedMbps)
	}
	if instance.OperationalStatus != nil {
		d.Set(dlOperationalStatus, *instance.OperationalStatus)
	}
	if instance.BgpStatus != nil {
		d.Set(dlBgpStatus, *instance.BgpStatus)
	}
	if instance.CompletionNoticeRejectReason != nil {
		d.Set(dlCompletionNoticeRejectReason, *instance.CompletionNoticeRejectReason)
	}
	if instance.LocationName != nil {
		d.Set(dlLocationName, *instance.LocationName)
	}
	if instance.LocationDisplayName != nil {
		d.Set(dlLocationDisplayName, *instance.LocationDisplayName)
	}
	if instance.Vlan != nil {
		d.Set(dlVlan, *instance.Vlan)
	}
	if instance.Global != nil {
		d.Set(dlGlobal, *instance.Global)
	}
	if instance.Port != nil {
		d.Set(dlPort, *instance.Port.ID)
	}
	if instance.LinkStatus != nil {
		d.Set(dlLinkStatus, *instance.LinkStatus)
	}
	if instance.CreatedAt != nil {
		d.Set(dlCreatedAt, instance.CreatedAt.String())
	}
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/interconnectivity/direct-link")
	d.Set(ResourceName, *instance.Name)
	d.Set(ResourceCRN, *instance.Crn)
	d.Set(ResourceStatus, *instance.LinkStatus)
	if instance.ResourceGroup != nil {
		rg := instance.ResourceGroup
		d.Set(dlResourceGroup, *rg.ID)
		d.Set(ResourceGroupName, *rg.ID)
	}

	return nil
}

func resourceIBMdlGatewayUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	getOptions := &directlinkapisv1.GetGatewayOptions{
		ID: &ID,
	}
	_, detail, err := directLink.GetGateway(getOptions)

	if err != nil {
		log.Printf("Error fetching Direct Link Gateway (Dedicated Template):%s", detail)
		return err
	}

	updateGatewayOptionsModel := &directlinkapisv1.UpdateGatewayOptions{}
	updateGatewayOptionsModel.ID = &ID
	if d.HasChange(dlName) {
		name := d.Get(dlName).(string)
		updateGatewayOptionsModel.Name = &name
	}
	if d.HasChange(dlSpeedMbps) {
		speed := int64(d.Get(dlSpeedMbps).(int))
		updateGatewayOptionsModel.SpeedMbps = &speed
	}
	if d.HasChange(dlOperationalStatus) {
		if _, ok := d.GetOk(dlOperationalStatus); ok {
			operStatus := d.Get(dlOperationalStatus).(string)
			updateGatewayOptionsModel.OperationalStatus = &operStatus
		}
		if _, ok := d.GetOk(dlLoaRejectReason); ok {
			loaRejectReason := d.Get(dlLoaRejectReason).(string)
			updateGatewayOptionsModel.LoaRejectReason = &loaRejectReason
		}
	}
	if d.HasChange(dlGlobal) {
		global := d.Get(dlGlobal).(bool)
		updateGatewayOptionsModel.Global = &global
	}
	if d.HasChange(dlMetered) {
		metered := d.Get(dlMetered).(bool)
		updateGatewayOptionsModel.Metered = &metered
	}
	_, response, err := directLink.UpdateGateway(updateGatewayOptionsModel)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Gateway (Dedicated) err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlGatewayRead(d, meta)
}

func resourceIBMdlGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	delOptions := &directlinkapisv1.DeleteGatewayOptions{
		ID: &ID,
	}
	response, err := directLink.DeleteGateway(delOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Direct Link Gateway (Dedicated Template): %s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMdlGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return false, err
	}

	ID := d.Id()

	getOptions := &directlinkapisv1.GetGatewayOptions{
		ID: &ID,
	}
	_, response, err := directLink.GetGateway(getOptions)
	if err != nil {
		return false, fmt.Errorf("Error Getting Direct Link Gateway (Dedicated Template): %s\n%s", err, response)
	}

	if response.StatusCode == 404 {
		d.SetId("")
		return false, nil
	}
	return true, nil
}
