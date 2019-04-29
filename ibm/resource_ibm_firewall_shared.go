package ibm

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

const (
	FwHardwarePackageType = "ADDITIONAL_SERVICES_FIREWALL"
)

func resourceIBMFirewallShared() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMFirewallSharedCreate,
		Read:     resourceIBMFirewallSharedRead,
		Delete:   resourceIBMFirewallSharedDelete,
		Exists:   resourceIBMFirewallSharedExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"billing_item_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"firewall_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"10MBPS_HARDWARE_FIREWALL", "20MBPS_HARDWARE_FIREWALL", "100MBPS_HARDWARE_FIREWALL", "1000MBPS_HARDWARE_FIREWALL"}),
			},
			"guest_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_instance_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"hardware_instance_id"},
			},
			"hardware_instance_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"virtual_instance_id"},
			},
		},
	}
}

// keyName is in between:[10MBPS_HARDWARE_FIREWALL, 20MBPS_HARDWARE_FIREWALL,
//                         100MBPS_HARDWARE_FIREWALL, 1000MBPS_HARDWARE_FIREWALL]
func resourceIBMFirewallSharedCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	keyName := d.Get("firewall_type").(string)
	guestType := d.Get("guest_type").(string)

	var virtualId, hardwareId int
	if vID, ok := d.GetOk("virtual_instance_id"); ok {
		virtualId = vID.(int)
	}

	if hID, ok := d.GetOk("hardware_instance_id"); ok {
		hardwareId = hID.(int)
	}

	if virtualId == 0 && hardwareId == 0 {
		return fmt.Errorf("Provide either `virtual_instance_id` or `hardware_instance_id`")
	}

	//var productOrderContainer *string
	pkg, err := product.GetPackageByType(sess, FwHardwarePackageType)
	if err != nil {
		return err
	}

	// Get all prices for ADDITIONAL_SERVICES_FIREWALL with the given capacity
	productItems, err := product.GetPackageProducts(sess, *pkg.Id)
	if err != nil {
		return err
	}

	// Select only those product items with a matching keyname
	targetItems := []datatypes.Product_Item{}
	for _, item := range productItems {
		if *item.KeyName == keyName {
			targetItems = append(targetItems, item)
		}
	}

	if len(targetItems) == 0 {
		return fmt.Errorf("No product items matching %s could be found", keyName)
	}

	if guestType == "virtual machine" {
		productOrderContainer := datatypes.Container_Product_Order_Network_Protection_Firewall{
			Container_Product_Order: datatypes.Container_Product_Order{
				PackageId: pkg.Id,
				Prices: []datatypes.Product_Item_Price{
					{
						Id: targetItems[0].Prices[0].Id,
					},
				},
				Quantity: sl.Int(1),
				VirtualGuests: []datatypes.Virtual_Guest{{
					Id: sl.Int(virtualId),
				},
				},
			},
		}
		receipt, err := services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&productOrderContainer, sl.Bool(false))
		log.Print("receipt for order placed")
		log.Print(receipt)

		if err != nil {
			return fmt.Errorf("Error during creation of hardware firewall: %s", err)
		}

	}
	if guestType == "baremetal" {
		productOrderContainer := datatypes.Container_Product_Order_Network_Protection_Firewall{
			Container_Product_Order: datatypes.Container_Product_Order{
				PackageId: pkg.Id,
				Prices: []datatypes.Product_Item_Price{
					{
						Id: targetItems[0].Prices[0].Id,
					},
				},
				Quantity: sl.Int(1),
				Hardware: []datatypes.Hardware{{
					Id: sl.Int(hardwareId),
				},
				},
			},
		}
		receipt, err := services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&productOrderContainer, sl.Bool(false))
		log.Print("receipt for order placed")
		log.Print(receipt)

		if err != nil {
			return fmt.Errorf("Error during creation of hardware firewall: %s", err)
		}

	}
	log.Println("[INFO] Creating hardware firewall shared")

	log.Printf("[INFO] Wait one minute before fetching the firewall/device.")
	time.Sleep(time.Second * 30)

	return resourceIBMFirewallSharedRead(d, meta)
}

func resourceIBMFirewallSharedRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	firewall_type := (d.Get("firewall_type").(string))
	d.Set("firewall_type", firewall_type)

	var virtualId, hardwareId int
	if vID, ok := d.GetOk("virtual_instance_id"); ok {
		virtualId = vID.(int)
	}

	if hID, ok := d.GetOk("hardware_instance_id"); ok {
		hardwareId = hID.(int)
	}

	guestType := (d.Get("guest_type").(string))
	d.Set("guest_type", guestType)

	masked := "firewallServiceComponent.id"

	fservice := services.GetNetworkComponentFirewallService(sess)

	if guestType == "virtual machine" {
		service := services.GetVirtualGuestService(meta.(ClientSession).SoftLayerSession())
		result, err := service.Id(virtualId).Mask(masked).GetObject()

		if err != nil {
			return fmt.Errorf("Error retrieving firewall information: %s", err)
		}

		if result.FirewallServiceComponent == nil {
			return fmt.Errorf("Error retrieving firewall information.This resource has already been canceled.")
		}
		idd := *result.FirewallServiceComponent.Id

		d.SetId(fmt.Sprintf("%d", idd))
		// d.SetId(fmt.Sprintf("%d/%d", idd, macId))
		data, err := fservice.Id(idd).Mask("billingItem.id").GetObject()

		d.Set("billing_item_id", *data.BillingItem.Id)

	} else if guestType == "baremetal" {
		service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())
		resultNew, err := service.Id(hardwareId).Mask(masked).GetObject()

		if err != nil {
			return fmt.Errorf("Error retrieving firewall information: %s", err)
		}
		if resultNew.FirewallServiceComponent == nil {
			return fmt.Errorf("Error retrieving firewall information.This resource has already been canceled.")
		}
		idd2 := *resultNew.FirewallServiceComponent.Id
		// d.SetId(fmt.Sprintf("%d/%d", idd2, macId))
		d.SetId(fmt.Sprintf("%d", idd2))
		data2, err := fservice.Id(idd2).Mask("billingItem.id").GetObject()

		d.Set("billing_item_id", *data2.BillingItem.Id)

	}
	return nil
}

//detach hardware firewall from particular machine
func resourceIBMFirewallSharedDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	idd2 := (d.Get("billing_item_id")).(int)

	success, err := services.GetBillingItemService(sess).Id(idd2).CancelService()
	log.Print(success)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("SoftLayer reported an unsuccessful cancellation")
	}
	return nil
}

//exists method
func resourceIBMFirewallSharedExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	fservice := services.GetNetworkComponentFirewallService(sess)
	id, err := strconv.Atoi(d.Id())
	response, err := fservice.Id(id).GetObject()

	if err != nil {
		log.Printf("error fetching the firewall resource: %s", err)
		return false, err
	}
	log.Print(response)
	return true, nil
}
