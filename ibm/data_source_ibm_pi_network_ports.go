package ibm

import (
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

const dataSourcePrint = "Calling the datasource for Network Ports with the following data [ CloudInstance id - (%s) and the Network id - (%s) ]"

func dataSourceIBMPINetworkPorts() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPINetworkPortsRead,
		Schema: map[string]*schema.Schema{

			"pi_network_id": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Network Name",
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"network_ports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"macaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pvminstance": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pvminstanceid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"portid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"href": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPINetworkPortsRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	network_id := d.Get("pi_network_id").(string)
	log.Printf(dataSourcePrint, powerinstanceid, network_id)
	powerC := instance.NewIBMPINetworkClient(sess, powerinstanceid)
	networkportdata, err := powerC.GetAllPort(network_id, powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("network_ports", flattenNetworkPorts(networkportdata))
	return nil
}

func flattenNetworkPorts(list *models.NetworkPorts) []map[string]interface{} {

	networkports := make([]map[string]interface{}, len(list.Ports))
	for i, networkport := range list.Ports {
		p := map[string]interface{}{

			"description": networkport.Description,
			"status":      networkport.Status,
			"ipaddress":   networkport.IPAddress,
			"portid":      networkport.PortID,
			"href":        networkport.Href,
			"macaddress":  networkport.MacAddress,
			"pvminstance": flattenpvm(networkport.PvmInstance),
		}
		networkports[i] = p

	}
	return networkports
}

func flattenpvm(pvmInstance *models.NetworkPortPvmInstance) []map[string]interface{} {
	pvms := make([]map[string]interface{}, 0, 1)
	pvmdata := make(map[string]interface{})

	if pvmInstance != nil {
		pvmdata["pvminstanceid"] = pvmInstance.PvmInstanceID
		pvmdata["href"] = pvmInstance.Href
		pvms = append(pvms, pvmdata)

	} else {
		log.Printf("No instance has been allocated")
		pvmdata["pvminstanceid"] = ""
		pvmdata["href"] = ""
		pvms = append(pvms, pvmdata)

	}

	return pvms
}
