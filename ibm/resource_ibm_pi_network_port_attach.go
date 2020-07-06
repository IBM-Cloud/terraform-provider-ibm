package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMPINetworkPortAttach() *schema.Resource {
	return &schema.Resource{

		Create: resourceIBMPINetworkPortAttachCreate,
		Read:   resourceIBMPINetworkPortAttachRead,
		Update: resourceIBMPINetworkPortAttachUpdate,
		Delete: resourceIBMPINetworkPortAttachDelete,
		//Exists:   resourceIBMPINetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{

			"port_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PIInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PINetworkName: {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

}

func resourceIBMPINetworkPortAttachCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	networkname := d.Get(helpers.PINetworkName).(string)
	portid := d.Get("port_id").(string)
	instancename := d.Get(helpers.PIInstanceName).(string)
	description := d.Get("description").(string)
	client := st.NewIBMPINetworkClient(sess, powerinstanceid)

	log.Printf("Printing the input to the resource powerinstance [%s] and network name [%s] and the portid [%s]", powerinstanceid, networkname, portid)
	networkPortResponse, err := client.AttachPort(powerinstanceid, networkname, portid, description, instancename, postTimeOut)

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

	return resourceIBMPINetworkPortAttachRead(d, meta)
}

func resourceIBMPINetworkPortAttachRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Calling ther NetworkAttach Read code")

	return nil
}

func resourceIBMPINetworkPortAttachUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Calling the attach update ")
	return nil
}

func resourceIBMPINetworkPortAttachDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Detaching the network port from the Instance ")

	/*	log.Printf("Calling the network delete functions. ")
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
		portid := parts[1]
		client := st.NewIBMPINetworkClient(sess, powerinstanceid)

		network, err := client.DetachPort(powerinstanceid, powernetworkname, portid, deleteTimeOut)*/

	return nil
}

func isWaitForIBMPINetworkPortAvailable(client *st.IBMPINetworkClient, id string, timeout time.Duration, powerinstanceid, networkname string) (interface{}, error) {
	log.Printf("Waiting for Power Network (%s) that was created for Network Zone (%s) to be available.", id, networkname)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PINetworkProvisioning},
		Target:     []string{"DOWN"},
		Refresh:    isIBMPINetworkPortRefreshFunc(client, id, powerinstanceid, networkname),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isIBMPINetworkPortRefreshFunc(client *st.IBMPINetworkClient, id, powerinstanceid, networkname string) resource.StateRefreshFunc {

	log.Printf("Calling the IsIBMPINetwork Refresh Function....with the following id (%s) for network port and following id (%s) for network name and waiting for network to be READY", id, networkname)
	return func() (interface{}, string, error) {
		network, err := client.GetPort(networkname, powerinstanceid, id, getTimeOut)
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
