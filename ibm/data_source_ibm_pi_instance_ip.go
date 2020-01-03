package ibm

import (
	"log"
	"net"
	"strconv"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPIInstanceIP() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIInstancesIPRead,
		Schema: map[string]*schema.Schema{

			helpers.PIInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Server Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PINetworkName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"requestip": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipoctet": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIInstancesIPRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	checkValidSubnet(d, meta)

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	powerinstancesubnet := d.Get(helpers.PINetworkName).(string)
	powerC := instance.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(d.Get(helpers.PIInstanceName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	pvminstanceid := *powervmdata.PvmInstanceID
	d.SetId(pvminstanceid)
	d.Set("memory", powervmdata.Memory)
	d.Set("processors", powervmdata.Processors)
	d.Set("status", powervmdata.Status)
	d.Set("proctype", powervmdata.ProcType)
	d.Set("volumeid", powervmdata.VolumeIds)

	for i, _ := range powervmdata.Addresses {
		if powervmdata.Addresses[i].NetworkName == powerinstancesubnet {
			log.Printf("Found the  zone")
			log.Printf("Printing the ip %s", powervmdata.Addresses[i].IP)
			d.Set("requestip", powervmdata.Addresses[i].IP)

			IPObject := net.ParseIP(powervmdata.Addresses[i].IP).To4()

			//log.Printf("Printing the value %s", IPObject.String())
			//log.Printf("Printing %s", strconv.Itoa(int(IPObject[3])))
			d.Set("ipoctet", strconv.Itoa(int(IPObject[3])))

			//log.Printf("Printing the value %s",ipoctet)
		}

		log.Printf("Got the  subnet.. %s", powerinstancesubnet)

	}

	return nil

}

func checkValidSubnet(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	powerinstancesubnet := d.Get(helpers.PINetworkName).(string)

	log.Printf("The subnet name is %s", powerinstancesubnet)

	networkC := instance.NewIBMPINetworkClient(sess, powerinstanceid)
	networkdata, err := networkC.Get(powerinstancesubnet, powerinstanceid)

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
	d.Set("pi_network_id", networkdata.NetworkID)
	d.Set("pi_vlan_id", networkdata.VlanID)

	return nil
}
