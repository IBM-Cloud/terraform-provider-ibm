package ibm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/networking-go-sdk/directlinkapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
	dlTags                         = "tags"
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
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "BGP ASN",
			},
			dlBgpBaseCidr: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "BGP base CIDR",
			},
			dlCrossConnectRouter: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Cross connect router",
			},
			dlGlobal: {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    false,
				Description: "Gateways with global routing (true) can connect to networks outside their associated region",
			},
			dlLocationName: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Gateway location",
			},
			dlMetered: {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    false,
				Description: "Metered billing option",
			},
			dlName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "The unique user-defined name for this gateway",
				ValidateFunc: InvokeValidator("ibm_dl_gateway", dlName),
				// ValidateFunc: validateRegexpLen(1, 63, "^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$"),
			},
			dlCarrierName: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Carrier name",
				// ValidateFunc: validateRegexpLen(1, 128, "^[a-z][A-Z][0-9][ -_]$"),
			},
			dlCustomerName: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Customer name",
				// ValidateFunc: validateRegexpLen(1, 128, "^[a-z][A-Z][0-9][ -_]$"),
			},
			dlSpeedMbps: {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "Gateway speed in megabits per second",
			},
			dlType: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Gateway type",
				ValidateFunc: InvokeValidator("ibm_dl_gateway", dlType),
				// ValidateFunc: validateAllowedStringValue([]string{"dedicated", "connect"}),
			},
			dlBgpCerCidr: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "BGP customer edge router CIDR",
			},
			dlLoaRejectReason: {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    false,
				Description: "Loa reject reason",
			},
			dlBgpIbmCidr: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "BGP IBM CIDR",
			},
			dlResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Gateway resource group",
			},

			dlOperationalStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway operational status",
			},
			dlPort: {
				Type: schema.TypeString,

				Computed:    true,
				Description: "Gateway port",
			},
			dlProviderAPIManaged: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether gateway was created through a provider portal",
			},
			dlVlan: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "VLAN allocated for this gateway",
			},
			dlBgpIbmAsn: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IBM BGP ASN",
			},

			dlBgpStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway BGP status",
			},
			dlCompletionNoticeRejectReason: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reason for completion notice rejection",
			},
			dlCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time resource was created",
			},
			dlCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN (Cloud Resource Name) of this gateway",
			},
			dlLinkStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway link status",
			},
			dlLocationDisplayName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway location long name",
			},
			dlTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "Tags for the direct link gateway",
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

func resourceIBMDLGatewayValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 2)
	dlTypeAllowedValues := "dedicated, connect"

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 dlType,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              dlTypeAllowedValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 dlName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	ibmISDLGatewayResourceValidator := ResourceValidator{ResourceName: "ibm_dl_gateway", Schema: validateSchema}
	return &ibmISDLGatewayResourceValidator
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

	if dtype == "dedicated" {
		if _, ok := d.GetOk(dlCarrierName); ok {
			carrierName := d.Get(dlCarrierName).(string)
			gatewayTemplateModel.CarrierName = &carrierName
		} else {
			err = fmt.Errorf("Error creating gateway, %s is a required field", dlCarrierName)
			log.Printf("%s is a required field", dlCarrierName)
			return err
		}
		if _, ok := d.GetOk(dlCrossConnectRouter); ok {
			crossConnectRouter := d.Get(dlCrossConnectRouter).(string)
			gatewayTemplateModel.CrossConnectRouter = &crossConnectRouter
		} else {
			err = fmt.Errorf("Error creating gateway, %s is a required field", dlCrossConnectRouter)
			log.Printf("%s is a required field", dlCrossConnectRouter)
			return err
		}
		if _, ok := d.GetOk(dlLocationName); ok {
			locationName := d.Get(dlLocationName).(string)
			gatewayTemplateModel.LocationName = &locationName
		} else {
			err = fmt.Errorf("Error creating gateway, %s is a required field", dlLocationName)
			log.Printf("%s is a required field", dlLocationName)
			return err
		}
		if _, ok := d.GetOk(dlCustomerName); ok {
			customerName := d.Get(dlCustomerName).(string)
			gatewayTemplateModel.CustomerName = &customerName
		} else {
			err = fmt.Errorf("Error creating gateway, %s is a required field", dlCustomerName)
			log.Printf("%s is a required field", dlCustomerName)
			return err
		}
	}

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

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(dlTags); ok || v != "" {
		oldList, newList := d.GetChange(dlTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *gateway.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource direct link gateway dedicated (%s) tags: %s", d.Id(), err)
		}
	}

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
	tags, err := GetTagsUsingCRN(meta, *instance.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource direct link gateway (%s) tags: %s", d.Id(), err)
	}
	d.Set(dlTags, tags)
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
	instance, detail, err := directLink.GetGateway(getOptions)

	if err != nil {
		log.Printf("Error fetching Direct Link Gateway (Dedicated Template):%s", detail)
		return err
	}

	updateGatewayOptionsModel := &directlinkapisv1.UpdateGatewayOptions{}
	updateGatewayOptionsModel.ID = &ID

	if d.HasChange(dlTags) {
		oldList, newList := d.GetChange(dlTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *instance.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource direct link gateway dedicated (%s) tags: %s", *instance.ID, err)
		}
	}

	if d.HasChange(dlName) {
		name := d.Get(dlName).(string)
		updateGatewayOptionsModel.Name = &name
	}
	if d.HasChange(dlSpeedMbps) {
		speed := int64(d.Get(dlSpeedMbps).(int))
		updateGatewayOptionsModel.SpeedMbps = &speed
	}
	/*
		NOTE: Operational Status cannot be maintained in terraform. The status keeps changing automatically in server side.
		Hence, cannot be maintained in terraform.
		Operational Status and LoaRejectReason are linked.
		Hence, a user cannot update through terraform.

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
	*/
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
