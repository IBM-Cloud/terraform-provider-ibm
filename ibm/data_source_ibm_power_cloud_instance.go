package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
)

func dataSourceIBMPowerCloudInstance() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPowerCloudInstanceRead,
		Schema: map[string]*schema.Schema{

			"cloudinstanceid": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Cloud Instance Id",
				ValidateFunc: validation.NoZeroValues,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenantid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"openstackid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"processors": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"initialized": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"procunits": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pvminstances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"proctype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
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
					},
				},
			},

			/*"limits": {
				Type: schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"processors": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"procunits": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"storage": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instances": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			}, */

		},
	}
}

func dataSourceIBMPowerCloudInstanceRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	cloudinstance := instance.NewPowerCloudInstanceClient(sess)
	clouddata, err := cloudinstance.Get(d.Get("cloudinstanceid").(string))

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	d.Set("tenantid", clouddata.TenantID)
	d.Set("processors", clouddata.Usage.Processors)
	d.Set("memory", clouddata.Usage.Memory)
	d.Set("instances", clouddata.Usage.Instances)
	d.Set("procunits", clouddata.Usage.ProcUnits)
	d.Set("region", clouddata.Region)

	if clouddata.PvmInstances != nil {
		pvms := make([]map[string]interface{}, len(clouddata.PvmInstances))
		for i, pvmdata := range clouddata.PvmInstances {

			p := make(map[string]interface{})
			p["servername"] = pvmdata.ServerName
			p["pvmstatus"] = pvmdata.Health.Status
			p["pvmid"] = pvmdata.PvmInstanceID
			pvms[i] = p
		}
		d.Set("pvminstances", pvms)
	}
	return nil
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
