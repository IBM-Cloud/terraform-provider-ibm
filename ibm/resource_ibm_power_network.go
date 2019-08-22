package ibm

import (
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"net"
	"strconv"

	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	PowerNetworkReady             = "ready"
	PowerNetworkID                = "networkid"
	PowerNetworkName              = "name"
	PowerNetworkCidr              = "cidr"
	PowerNetworkDNS               = "dns"
	PowerNetworkType              = "networktype"
	PowerNetworkGateway           = "gateway"
	PowerNetworkStartingIPAddress = "startip"
	PowerNetworkEndingIPAddress   = "endip"
	PowerNetworkIPAddressRange    = "ipaddressrange"
	PowerNetworkVlanId            = "vlanId"
	PowerNetworkProvisioning      = "build"
)

func resourceIBMPowerNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPowerNetworkCreate,
		Read:   resourceIBMPowerNetworkRead,
		Update: resourceIBMPowerNetworkUpdate,
		Delete: resourceIBMPowerNetworkDelete,
		//Exists:   resourceIBMPowerNetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PowerNetworkID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PowerNetworkType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"vlan", "public-vlan"}),
			},

			PowerNetworkName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PowerNetworkDNS: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			PowerNetworkCidr: {
				Type:     schema.TypeString,
				Required: true,
			},

			PowerNetworkIPAddressRange: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			PowerNetworkGateway: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIBMPowerNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	networkname := d.Get(PowerNetworkName).(string)
	networktype := d.Get(PowerNetworkType).(string)
	networkcidr := d.Get(PowerNetworkCidr).(string)
	networkdns := expandStringList((d.Get(PowerNetworkDNS).(*schema.Set)).List())

	log.Printf("Printing the data ")

	client := st.NewPowerNetworkClient(sess)
	networkgateway, firstip, lastip := generateData(networkcidr)
	networkResponse, _, err := client.Create(networkname, networktype, networkcidr, networkdns, networkgateway, firstip, lastip)

	if err != nil {
		return err
	}

	log.Printf("Printing the networkresponse %+v", &networkResponse)

	PowerNetworkID := *networkResponse.NetworkID

	d.SetId(PowerNetworkID)

	//log.Printf("The powernetwork id for the %s networname is %s",networkname,PowerNetworkID)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForPowerNetworkAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil

	return resourceIBMPowerNetworkRead(d, meta)
}

func resourceIBMPowerNetworkRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	networkC := st.NewPowerNetworkClient(sess)
	networkdata, err := networkC.Get(d.Get("name").(string))

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("networkid", networkdata.NetworkID)
	d.Set("cidr", networkdata.Cidr)
	d.Set("type", networkdata.Type)
	d.Set("gateway", networkdata.Gateway)
	d.Set("vlanid", networkdata.VlanID)

	return nil

}

func resourceIBMPowerNetworkUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPowerNetworkDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func expanddnsServerAttribute(strings []interface{}) []string {
	expandedStrings := make([]string, len(strings))
	for i, v := range strings {
		expandedStrings[i] = v.(string)
	}

	return expandedStrings
}

func isWaitForPowerNetworkAvailable(client *st.PowerNetworkClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Power Network (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", PowerNetworkProvisioning},
		Target:     []string{PowerNetworkReady},
		Refresh:    isPowerNetworkRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPowerNetworkRefreshFunc(client *st.PowerNetworkClient, id string) resource.StateRefreshFunc {

	log.Printf("Calling the IsPowerNetwork Refresh Function....")
	return func() (interface{}, string, error) {
		network, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if network.VlanID == nil {
			//if network.State == "available" {
			return network, PowerNetworkReady, nil
		}

		return network, PowerNetworkProvisioning, nil
	}
}

func generateData(cdir string) (gway, firstip, lastip string) {
	_, ipv4Net, err := net.ParseCIDR(cdir)

	if err != nil {
		log.Fatal(err)
	}

	var subnet_to_size = map[string]int{
		"21": 2048,
		"22": 1024,
		"23": 512,
		"24": 256,
		"25": 128,
		"26": 64,
		"27": 32,
		"28": 16,
		"29": 8,
		"30": 4,
		"31": 2,
	}

	subnetsize, _ := ipv4Net.Mask.Size()

	gateway, err := cidr.Host(ipv4Net, 1)
	ad := cidr.AddressCount(ipv4Net)

	convertedad := strconv.FormatUint(ad, 10)
	firstusable, err := cidr.Host(ipv4Net, 2)
	lastusable, err := cidr.Host(ipv4Net, subnet_to_size[convertedad]-2)
	log.Printf("The gateway  value is %s and  %s the count is %s and first ip is %s last one is  %s", gateway, subnetsize, convertedad, firstusable, lastusable)

	return gateway.String(), firstusable.String(), lastusable.String()

}
