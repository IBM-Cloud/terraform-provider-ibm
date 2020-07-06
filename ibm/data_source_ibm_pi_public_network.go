package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPIPublicNetwork() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIPublicNetworksRead,
		Schema: map[string]*schema.Schema{

			helpers.PINetworkName: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Network Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
				Deprecated:   "This field is deprectaed.",
			},

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"network_id": {
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

			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIPublicNetworksRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	networkC := instance.NewIBMPINetworkClient(sess, powerinstanceid)
	networkdata, err := networkC.GetPublic(powerinstanceid, getTimeOut)

	if err != nil {
		return err
	}

	d.SetId(*networkdata.Networks[0].NetworkID)
	d.Set("type", networkdata.Networks[0].Type)
	d.Set("name", networkdata.Networks[0].Name)
	d.Set("vlan_id", networkdata.Networks[0].VlanID)

	return nil

}
