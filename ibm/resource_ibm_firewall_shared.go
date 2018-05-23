package ibm

import (
	"fmt"

	"log"

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
		Create: resourceIBMFirewallSharedCreate,
		Read:   resourceIBMFirewallSharedRead,
		// Update:   resourceIBMFirewallSharedUpdate,
		Delete: resourceIBMFirewallSharedDelete,
		// Exists:   resourceIBMFirewallSharedExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"firewall_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"guest_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"guest_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

// keyName is in between:[10MBPS_HARDWARE_FIREWALL, 20MBPS_HARDWARE_FIREWALL,
//                         100MBPS_HARDWARE_FIREWALL, 1024MBPS_HARDWARE_FIREWALL]
func resourceIBMFirewallSharedCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	keyName := d.Get("firewall_type").(string)
	guestType := d.Get("guest_type").(string)
	machineId := d.Get("guest_id").(int)

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
					Id: sl.Int(machineId),
				},
				},
			},
		}
		receipt, err := services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&productOrderContainer, sl.Bool(false))
		log.Print(*receipt.OrderId)
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
					Id: sl.Int(machineId),
				},
				},
			},
		}
		receipt, err := services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&productOrderContainer, sl.Bool(false))
		log.Print(*receipt.OrderId)
		if err != nil {
			return fmt.Errorf("Error during creation of hardware firewall: %s", err)
		}
	}
	log.Println("[INFO] Creating hardware firewall shared")

	d.Set("firewall_type", keyName)
	d.Set("guest_id", machineId)
	d.Set("guest_type", guestType)

	return resourceIBMFirewallSharedRead(d, meta)
}

func resourceIBMFirewallSharedRead(d *schema.ResourceData, meta interface{}) error {

	macId := (d.Get("guest_id").(int))

	guestType := (d.Get("guest_type").(string))

	masked := "firewallServiceComponent"

	if guestType == "virtual machine" {
		service := services.GetVirtualGuestService(meta.(ClientSession).SoftLayerSession())

		result, err := service.Id(macId).Mask(masked).GetObject()

		if err != nil {
			return fmt.Errorf("Error retrieving firewall information: %s", err)
		}

		log.Print(*result.FirewallServiceComponent.GuestNetworkComponent.Id)
		d.SetId(fmt.Sprintf("%d", *result.FirewallServiceComponent.GuestNetworkComponent.Id))

		return nil
	} else {
		if guestType == "baremetal" {
			service := services.GetHardwareService(meta.(ClientSession).SoftLayerSession())

			resultNew, err := service.Id(macId).Mask(masked).GetObject()

			if err != nil {
				return fmt.Errorf("Error retrieving firewall information: %s", err)
			}

			log.Print(*resultNew.FirewallServiceComponent.NetworkComponent.Id)

			d.SetId(fmt.Sprintf("%d", (*resultNew.FirewallServiceComponent.NetworkComponent.Id)))

			return nil
		}
	}
	return nil
}

func resourceIBMFirewallSharedDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
