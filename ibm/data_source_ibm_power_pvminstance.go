package ibm

import (
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
	_ "github.ibm.com/Bluemix/power-go-client/power/client"
)

func dataSourceIBMPowerPVMInstance() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPowerPVMInstancesRead,
		Schema: map[string]*schema.Schema{

			"servername": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Server Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			"volumeid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"processors": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"shareable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthstatus": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lastupdate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"macaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networkid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networkname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*"version": {
							Type:     schema.TypeFloat,
							Computed: true,
						},*/
					},
				},
			},
			"proctype": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPowerPVMInstancesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	powerC := instance.NewPowerPvmClient(sess)
	powervmdata, err := powerC.Get(d.Get("servername").(string))

	if err != nil {
		return err
	}

	pvminstanceid := *powervmdata.PvmInstanceID
	d.SetId(pvminstanceid)
	d.Set("memory", powervmdata.Memory)
	d.Set("processors", powervmdata.Processors)
	d.Set("status", powervmdata.Status)
	d.Set("proctype", powervmdata.ProcType)

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {

			p := make(map[string]interface{})
			p["ip"] = pvmip.IP
			p["networkname"] = pvmip.NetworkName
			p["networkid"] = pvmip.NetworkID
			p["macaddress"] = pvmip.MacAddress
			p["type"] = pvmip.Type
			pvmaddress[i] = p
		}
		d.Set("addresses", pvmaddress)

	}

	if powervmdata.Health != nil {

		d.Set("healthstatus", powervmdata.Health.Status)

	}

	return nil

}
