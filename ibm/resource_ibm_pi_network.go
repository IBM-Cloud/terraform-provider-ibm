package ibm

import (
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	IBMPINetworkReady             = "ready"
	IBMPINetworkID                = "networkid"
	IBMPINetworkName              = "name"
	IBMPINetworkCidr              = "cidr"
	IBMPINetworkDNS               = "dns"
	IBMPINetworkType              = "networktype"
	IBMPINetworkGateway           = "gateway"
	IBMPINetworkStartingIPAddress = "startip"
	IBMPINetworkEndingIPAddress   = "endip"
	IBMPINetworkIPAddressRange    = "ipaddressrange"
	IBMPINetworkVlanId            = "vlanId"
	IBMPINetworkProvisioning      = "build"
)

func resourceIBMPINetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPINetworkCreate,
		Read:   resourceIBMPINetworkRead,
		Update: resourceIBMPINetworkUpdate,
		Delete: resourceIBMPINetworkDelete,
		//Exists:   resourceIBMPINetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			IBMPINetworkType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"vlan", "public-vlan"}),
			},

			IBMPINetworkName: {
				Type:     schema.TypeString,
				Required: true,
			},
			IBMPINetworkDNS: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			IBMPINetworkCidr: {
				Type:     schema.TypeString,
				Required: true,
			},

			IBMPINetworkIPAddressRange: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			IBMPINetworkGateway: {
				Type:     schema.TypeString,
				Optional: true,
			},

			"powerinstanceid": {
				Type:     schema.TypeString,
				Required: true,
			},

			//Computed Attributes

			IBMPINetworkID: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMPINetworkCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	networkname := d.Get(IBMPINetworkName).(string)
	networktype := d.Get(IBMPINetworkType).(string)
	networkcidr := d.Get(IBMPINetworkCidr).(string)
	networkdns := expandStringList((d.Get(IBMPINetworkDNS).(*schema.Set)).List())

	log.Printf("Printing the data ")

	client := st.NewIBMPINetworkClient(sess, powerinstanceid)
	networkgateway, firstip, lastip := generateData(networkcidr)
	networkResponse, _, err := client.Create(networkname, networktype, networkcidr, networkdns, networkgateway, firstip, lastip, powerinstanceid)

	if err != nil {
		return err
	}

	log.Printf("Printing the networkresponse %+v", &networkResponse)

	IBMPINetworkID := *networkResponse.NetworkID

	d.SetId(IBMPINetworkID)
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForIBMPINetworkAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	return resourceIBMPINetworkRead(d, meta)
}

func resourceIBMPINetworkRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	networkC := st.NewIBMPINetworkClient(sess, powerinstanceid)
	networkdata, err := networkC.Get(d.Get("name").(string), powerinstanceid)

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

func resourceIBMPINetworkUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPINetworkDelete(data *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the network delete functions. ")
	return nil
}

func resourceIBMPINetworkExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPINetworkClient(sess, powerinstanceid)

	network, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return network.NetworkID == &id, nil
}

func isWaitForIBMPINetworkAvailable(client *st.IBMPINetworkClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for Power Network (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", IBMPINetworkProvisioning},
		Target:     []string{IBMPINetworkReady},
		Refresh:    isIBMPINetworkRefreshFunc(client, id, powerinstanceid),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isIBMPINetworkRefreshFunc(client *st.IBMPINetworkClient, id, powerinstanceid string) resource.StateRefreshFunc {

	log.Printf("Calling the IsIBMPINetwork Refresh Function....")
	return func() (interface{}, string, error) {
		network, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if network.VlanID == nil {
			//if network.State == "available" {
			return network, IBMPINetworkReady, nil
		}

		return network, IBMPINetworkProvisioning, nil
	}
}

func generateIPData(cdir string) (gway, firstip, lastip string) {
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
	if err != nil {
		log.Printf("Failed to get the gateway for this cdir passed in %s", cdir)
		log.Fatal(err)
	}
	ad := cidr.AddressCount(ipv4Net)

	convertedad := strconv.FormatUint(ad, 10)
	firstusable, err := cidr.Host(ipv4Net, 2)
	lastusable, err := cidr.Host(ipv4Net, subnet_to_size[convertedad]-2)
	log.Printf("The gateway  value is %s and  %s the count is %s and first ip is %s last one is  %s", gateway, subnetsize, convertedad, firstusable, lastusable)

	return gateway.String(), firstusable.String(), lastusable.String()

}
