package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/helpers"
)

func dataSourceIBMPINetwork() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPINetworksRead,
		Schema: map[string]*schema.Schema{

			helpers.PINetworkName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Network Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"networkid": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"vlanid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPINetworksRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	networkC := instance.NewIBMPINetworkClient(sess, powerinstanceid)
	networkdata, err := networkC.Get(d.Get(helpers.PINetworkName).(string), powerinstanceid)

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
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
