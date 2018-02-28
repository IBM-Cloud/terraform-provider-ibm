package ibm

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	slsession "github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMMultiVlanFirewall() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkMultiVlanCreate,
		Read:     resourceIBMComputeSSHKeyRead,
		Update:   resourceIBMComputeSSHKeyUpdate,
		Delete:   resourceIBMComputeSSHKeyDelete,
		Exists:   resourceIBMComputeSSHKeyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
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
			},

			"firewall_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"addon_configuration": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}
func resourceIBMNetworkMultiVlanCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	name := d.Get("name").(string)

	datacenter := d.Get("datacenter").(string)
	datacenterfilter := "{\"name\":{\"operation\":\"datacentername\"}}"
	datacenterfilter = strings.Replace(datacenterfilter, "datacentername", datacenter, -1)
	datacentermask := "priceGroups.id"
	pod := d.Get("pod").(string)
	id := datacenter + "." + pod
	log.Print("================ the id is ------------", id)
	service := services.GetNetworkPodService(sess)
	resp, err := service.StringId(id).GetCapabilities()
	if stringinSlice("SUPPORTS_MULTI_VLAN_DEDICATED_FIREWALL", resp) {
		return fmt.Errorf("Datacenter and Pod Combination do not support Multi Vlan Dedicated Firewall")
	}
	router, err := service.StringId(id).Mask("frontendRouterId").GetObject()
	routerid := router.FrontendRouterId
	datacenterservice := services.GetLocationDatacenterService(sess)
	resp1, err2 := datacenterservice.Filter(datacenterfilter).Mask(datacentermask).GetDatacenters()
	if err2 != nil {
		return fmt.Errorf("Cannot find the datacenter specified")
	}
	var s []*int
	for _, item := range resp1 {
		for _, temp := range item.PriceGroups {
			s = append(s, temp.Id)
		}
	}
	var pricegroupid int
	for _, ids := range s {
		if *ids == 503 || *ids == 505 || *ids == 507 || *ids == 509 || *ids == 545 {
			pricegroupid = *ids
		}
	}
	var priceids []int
	addonconfigurations, ok := d.Get("addon_configuration").([]interface{})
	var actualaddons []string
	for _, addons := range addonconfigurations {
		actualaddons = append(actualaddons, addons.(string))
	}
	if !ok {
		return fmt.Errorf("Couldnt convert addons")
	}
	for _, addons := range actualaddons {
		log.Println("The addon is ", addons)
		priceid, err := returnpriceidaccordingtopackageid(addons, pricegroupid, sess)
		if err != nil {
			return fmt.Errorf("Erorr in returnpriceidaccordingtopackageid")
		}
		if priceid != 0 {
			priceids = append(priceids, priceid)
		} else {
			return fmt.Errorf("Please enter a valid configuration addon")
		}
	}
	priceItems := []datatypes.Product_Item_Price{}
	for _, priceid := range priceids {
		priceItem := datatypes.Product_Item_Price{
			Id: &priceid,
		}
		priceItems = append(priceItems, priceItem)
	}
	packageid := 863
	Complextype := "SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated"
	productOrderContainer := datatypes.Container_Product_Order_Network_Protection_Firewall_Dedicated{
		Container_Product_Order: datatypes.Container_Product_Order{
			PackageId:   &packageid,
			Prices:      priceItems,
			Quantity:    sl.Int(1),
			Location:    &datacenter,
			ComplexType: &Complextype,
		},
		Name:     sl.String(name),
		RouterId: routerid,
	}
	_, err = services.GetProductOrderService(sess.SetRetries(0)).
		VerifyOrder(&productOrderContainer)
	if err != nil {
		return fmt.Errorf("Error during Verify order for Creating: %s", err)
	}
	//place order
	_, err = services.GetProductOrderService(sess.SetRetries(0)).
		PlaceOrder(&productOrderContainer, sl.Bool(false))
	if err != nil {
		return fmt.Errorf("Error during Verify order for Creating: %s", err)
	}

	return err
}

func stringinSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func returnpriceidaccordingtopackageid(appliance string, pricegroupid int, sess *slsession.Session) (int, error) {
	productpackageservice := services.GetProductPackageService(sess)
	productpackageservicefilter := strings.Replace("{\"items\":{\"description\":{\"operation\":\"appliance\"}}}", "appliance", appliance, -1)
	resp, err := productpackageservice.Mask("description,prices.locationGroupId,prices.id").Filter(productpackageservicefilter).Id(863).GetItems()
	if err != nil {
		return 0, fmt.Errorf("Eroor in returnpriceidaccordingtopackageid")
	}
	m := make(map[*int]*int)
	for _, items := range resp {
		for _, temp := range items.Prices {
			m[temp.LocationGroupId] = temp.Id
		}
	}
	//var priceid []int
	if val, ok := m[&pricegroupid]; ok {
		return *val, nil
	}
	return 0, nil

}
