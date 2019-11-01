package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPITenant() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPITenantRead,
		Schema: map[string]*schema.Schema{
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"tenantid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creationdate": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"tenantname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloudinstances": {

				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudinstanceid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"href": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initialize": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						//"limits": {
						//	Type:     schema.TypeString,
						//	Computed: true,
						//},
					},
				},
			},
		},
	}
}

func dataSourceIBMPITenantRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	//tenantid := d.Get("tenantid").(string)

	tenantC := instance.NewIBMPITenantClient(sess, powerinstanceid)
	tenantData, err := tenantC.Get(powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	log.Printf("The creation date is %s", tenantData.CreationDate.String())
	d.Set("tenantid", tenantData.TenantID)
	d.Set("creationdate", tenantData.CreationDate)
	d.Set("enabled", tenantData.Enabled)

	if tenantData.CloudInstances != nil {

		d.Set("tenantname", tenantData.CloudInstances[0].Name)
	}

	if tenantData.CloudInstances != nil {
		tenants := make([]map[string]interface{}, len(tenantData.CloudInstances))
		for i, cloudinstance := range tenantData.CloudInstances {
			j := make(map[string]interface{})
			j["region"] = cloudinstance.Region
			j["cloudinstanceid"] = cloudinstance.CloudInstanceID
			tenants[i] = j
		}

		d.Set("cloudinstances", tenants)
	}

	return nil

}
