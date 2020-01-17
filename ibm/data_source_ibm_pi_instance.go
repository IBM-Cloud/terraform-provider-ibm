package ibm

import (
	_ "fmt"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceIBMPIInstance() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIInstancesRead,
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

			// Computed Attributes
			"volumeid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicationpolicy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicants": {
				Type:     schema.TypeInt,
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
						"externalip": {
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

func dataSourceIBMPIInstancesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

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

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {

			p := make(map[string]interface{})
			p["ip"] = pvmip.IP
			p["networkname"] = pvmip.NetworkName
			p["networkid"] = pvmip.NetworkID
			p["macaddress"] = pvmip.MacAddress
			p["type"] = pvmip.Type
			p["externalip"] = pvmip.ExternalIP
			pvmaddress[i] = p
		}
		d.Set("addresses", pvmaddress)

	}

	if powervmdata.Health != nil {

		d.Set("healthstatus", powervmdata.Health.Status)

	}

	return nil

}
