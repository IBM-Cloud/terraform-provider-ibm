package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/helpers/location"
	"github.com/softlayer/softlayer-go/services"
	slsession "github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMMultiVlanFirewall() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkMultiVlanCreate,
		Read:     resourceIBMFirewallRead,
		Delete:   resourceIBMMultiVlanFirewallDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"pod": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.TrimSpace(old) == strings.TrimSpace(new)
				},
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"public_vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"firewall_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"addon_configuration": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

const (
	productpackagefilter      = `{"keyName":{"operation":"FIREWALL_APPLIANCE"}}`
	complextype               = "SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated"
	productpackageservicemask = "description,prices.locationGroupId,prices.id"
	mandatoryfirewalltype     = "FortiGate Security Appliance"
)

func resourceIBMNetworkMultiVlanCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	name := d.Get("name").(string)
	firewalltype := d.Get("firewall_type").(string)
	datacenter := d.Get("datacenter").(string)
	pod := d.Get("pod").(string)
	nameofthepod := datacenter + "." + pod
	podservice := services.GetNetworkPodService(sess)
	podfilter := strings.Replace(`{"datacenterName":{"operation":"datacentername"}}`, "datacentername", datacenter, -1)
	podmask := `frontendRouterId,name`
	publicVlanID := d.Get("public_vlan_id").(int)

	// 1.Getting the router ID
	routerids, err := podservice.Filter(podfilter).Mask(podmask).GetAllObjects()
	if err != nil {
		return fmt.Errorf("Datacenter doesnt support multi-vlan-firewall,Please enter a different datacenter")
	}
	var routerid int
	for _, iterate := range routerids {
		if *iterate.Name == nameofthepod {
			routerid = *iterate.FrontendRouterId
		}
	}
	datacentername, ok := d.GetOk("datacenter")
	if !ok {
		return fmt.Errorf("The attribute datacenter is not defined")
	}

	//2.Get the datacenter id
	dc, err := location.GetDatacenterByName(sess, datacentername.(string), "id")
	if err != nil {
		return fmt.Errorf("Datacenter not found")
	}
	locationservice := services.GetLocationService(sess)

	//3. get the pricegroups that the datacenter belongs to
	priceidds, _ := locationservice.Id(*dc.Id).GetPriceGroups()
	var listofpriceids []int
	//store all the pricegroups a datacenter belongs to
	for _, priceidd := range priceidds {
		listofpriceids = append(listofpriceids, *priceidd.Id)
	}

	//4.get the addons that are specified
	addonconfigurations, ok := d.Get("addon_configuration").([]interface{})
	if !ok {
		return fmt.Errorf("Addons is an array of strings")
	}
	var actualaddons []string
	for _, addons := range addonconfigurations {
		actualaddons = append(actualaddons, addons.(string))
	}
	//appending the 20000GB Bandwidth item as it is mandatory
	actualaddons = append(actualaddons, firewalltype, "20000 GB Bandwidth")
	//appending the Fortigate Security Appliance as it is mandatory parameter for placing an order
	if firewalltype != mandatoryfirewalltype {
		actualaddons = append(actualaddons, mandatoryfirewalltype)
	}

	//5. Getting the priceids of items which have to be ordered
	priceItems := []datatypes.Product_Item_Price{}
	for _, addon := range actualaddons {
		actualpriceid, err := returnpriceidaccordingtopackageid(addon, listofpriceids, sess)
		if err != nil || actualpriceid == 0 {
			return fmt.Errorf("The addon or the firewall is not available for the datacenter you have selected. Please enter a different datacenter")
		}
		priceItem := datatypes.Product_Item_Price{
			Id: &actualpriceid,
		}
		priceItems = append(priceItems, priceItem)
	}

	//6.Get the package ID
	productpackageservice, _ := services.GetProductPackageService(sess).Filter(productpackagefilter).Mask(`id`).GetAllObjects()
	var productid int
	for _, packageid := range productpackageservice {
		productid = *packageid.Id
	}

	//7. Populate the container which needs to be sent for verify order and place order
	productOrderContainer := datatypes.Container_Product_Order_Network_Protection_Firewall_Dedicated{
		Container_Product_Order: datatypes.Container_Product_Order{
			PackageId:   &productid,
			Prices:      priceItems,
			Quantity:    sl.Int(1),
			Location:    &datacenter,
			ComplexType: sl.String(complextype),
		},
		Name:     sl.String(name),
		RouterId: &routerid,
		VlanId:   &publicVlanID,
	}

	//8.Calling verify order
	_, err = services.GetProductOrderService(sess.SetRetries(0)).
		VerifyOrder(&productOrderContainer)
	if err != nil {
		return fmt.Errorf("Error during Verify order for Creating: %s", err)
	}

	//9.Calling place order
	receipt, err := services.GetProductOrderService(sess.SetRetries(0)).
		PlaceOrder(&productOrderContainer, sl.Bool(false))
	if err != nil {
		return fmt.Errorf("Error during Place order for Creating: %s", err)
	}
	vlan, err := findDedicatedFirewallByOrderID(sess, *receipt.OrderId)
	if err != nil {
		return fmt.Errorf("Error during creation of dedicated hardware firewall: %s", err)
	}
	id := *vlan.NetworkVlanFirewall.Id
	d.SetId(fmt.Sprintf("%d", id))
	log.Printf("[INFO] Firewall ID: %s", d.Id())
	return err
}

func returnpriceidaccordingtopackageid(addon string, listofpriceids []int, sess *slsession.Session) (int, error) {
	productpackageservice := services.GetProductPackageService(sess)
	productpackageservicefilter := strings.Replace(`{"items":{"description":{"operation":"appliance"}}}`, "appliance", addon, -1)
	resp, err := productpackageservice.Mask(productpackageservicemask).Filter(productpackageservicefilter).Id(863).GetItems()
	if err != nil {
		return 0, err
	}
	m := make(map[int]int)
	for _, item := range listofpriceids {
		for _, items := range resp {
			for _, temp := range items.Prices {
				if temp.LocationGroupId == nil {
					m[item] = *temp.Id
				} else if item == *temp.LocationGroupId {
					m[*temp.LocationGroupId] = *temp.Id
				}
			}
		}
		if val, ok := m[item]; ok {
			return val, nil
		}
	}
	return 0, nil
}

func resourceIBMMultiVlanFirewallDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	fwService := services.GetNetworkVlanFirewallService(sess)

	fwID, _ := strconv.Atoi(d.Id())

	// Get billing item associated with the firewall
	billingItem, err := fwService.Id(fwID).GetBillingItem()

	if err != nil {
		return fmt.Errorf("Error while looking up billing item associated with the firewall: %s", err)
	}

	if billingItem.Id == nil {
		return fmt.Errorf("Error while looking up billing item associated with the firewall: No billing item for ID:%d", fwID)
	}

	success, err := services.GetBillingItemService(sess).Id(*billingItem.Id).CancelService()
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("SoftLayer reported an unsuccessful cancellation")
	}

	return resourceIBMMultiVlanFirewallRead(d, meta)
}
func resourceIBMMultiVlanFirewallRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()

	fwID, _ := strconv.Atoi(d.Id())

	_, err := services.GetNetworkVlanFirewallService(sess).
		Id(fwID).
		Mask(fwMask).
		GetObject()

	if err != nil {
		return fmt.Errorf("Error retrieving firewall information: %s", err)
	}
	return nil
}
