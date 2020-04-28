package ibm

import (
	"fmt"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMPINetworkPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPINetworkPortCreate,
		Read:   resourceIBMPINetworkPortRead,
		Update: resourceIBMPINetworkPortUpdate,
		Delete: resourceIBMPINetworkPortDelete,
		//Exists:   resourceIBMPINetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PINetworkName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},

			/*
			 "description": "",
			    "ipAddress": "172.31.96.93",
			    "macAddress": "fa:16:3e:30:b9:64",
			    "portID": "6c9d0e42-73f3-492f-840d-3d4a7a573014",
			*/

			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"macaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			//Computed Attributes

		},
	}
}

func resourceIBMPINetworkPortCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	networkname := d.Get(helpers.PINetworkName).(string)

	client := st.NewIBMPINetworkClient(sess, powerinstanceid)

	networkPortResponse, err := client.CreatePort(networkname, powerinstanceid)

	if err != nil {
		return err
	}

	log.Printf("Printing the networkresponse %+v", &networkPortResponse)

	IBMPINetworkPortID := *networkPortResponse.PortID

	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, IBMPINetworkPortID))
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForIBMPINetworkPortAvailable(client, IBMPINetworkPortID, d.Timeout(schema.TimeoutCreate), powerinstanceid, networkname)
	if err != nil {
		return err
	}

	return resourceIBMPINetworkPortRead(d, meta)
}

func resourceIBMPINetworkPortRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	powerinstanceid := parts[0]
	powernetworkname := d.Get(helpers.PINetworkName).(string)
	networkC := st.NewIBMPINetworkClient(sess, powerinstanceid)
	networkdata, err := networkC.GetPort(powernetworkname, powerinstanceid, parts[1])

	if err != nil {
		return err
	}

	d.Set("ipaddress", networkdata.IPAddress)
	d.Set("macaddress", networkdata.MacAddress)
	d.Set("status", networkdata.Status)
	d.Set("portid", networkdata.PortID)

	return nil

}

func resourceIBMPINetworkPortUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPINetworkPortDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the network delete functions. ")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	powernetworkname := d.Get(helpers.PINetworkName).(string)
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPINetworkClient(sess, powerinstanceid)
	log.Printf("Calling the client %v", client)

	log.Printf("Calling the delete with the following params delete with cloudinstance -> (%s) and networkid -->  (%s) and portid --> (%s) ", powerinstanceid, powernetworkname, parts[1])
	networkdata, err := client.DeletePort(powernetworkname, powerinstanceid, parts[1])

	log.Printf("Response from the deleteport call %v", networkdata)

	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMPINetworPortExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPINetworkClient(sess, powerinstanceid)

	network, err := client.Get(parts[0], powerinstanceid)
	if err != nil {

		return false, err
	}
	return *network.NetworkID == parts[1], nil
}

func isWaitForIBMPINetworkPortAvailable(client *st.IBMPINetworkClient, id string, timeout time.Duration, powerinstanceid, networkname string) (interface{}, error) {
	log.Printf("Waiting for Power Network (%s) that was created for Network Zone (%s) to be available.", id, networkname)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PINetworkProvisioning},
		Target:     []string{"DOWN"},
		Refresh:    isIBMPINetworkPortRefreshFunc(client, id, powerinstanceid, networkname),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isIBMPINetworkPortRefreshFunc(client *st.IBMPINetworkClient, id, powerinstanceid, networkname string) resource.StateRefreshFunc {

	log.Printf("Calling the IsIBMPINetwork Refresh Function....with the following id (%s) for network port and following id (%s) for network name and waiting for network to be READY", id, networkname)
	return func() (interface{}, string, error) {
		network, err := client.GetPort(networkname, powerinstanceid, id)
		if err != nil {
			return nil, "", err
		}

		if &network.PortID != nil {
			//if network.State == "available" {
			log.Printf(" The port has been created with the following ip address")
			return network, "DOWN", nil
		}

		return network, helpers.PINetworkProvisioning, nil
	}
}
