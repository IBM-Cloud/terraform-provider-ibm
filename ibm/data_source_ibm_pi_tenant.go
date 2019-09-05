package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"log"

	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
)

func dataSourceIBMPITenant() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPITenantRead,
		Schema: map[string]*schema.Schema{
			"powerinstanceid": {
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

	powerinstanceid := d.Get("powerinstanceid").(string)
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
	//d.Set("tenantname",tenantData.CloudInstances[0].Name)
	//d.Set("cloudinstances", flattenCloudInstances(tenantData.CloudInstances))
	log.Printf("Printing the tenant data %s", tenantData.CloudInstances)

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

func flattenCloudInstances(cloudinstances []*models.CloudInstanceReference) []map[string]interface{} {
	cloudInstances := make([]map[string]interface{}, len(cloudinstances))
	for _, i := range cloudinstances {
		l := map[string]interface{}{

			"region":          &i.Region,
			"href":            &i.Href,
			"cloudinstanceid": &i.CloudInstanceID,
			"initialized":     &i.Enabled,
			"name":            &i.Name,
		}
		cloudInstances = append(cloudInstances, l)
	}
	log.Printf("printing the cloudinstances %+v", cloudInstances)
	return cloudInstances
}

/*

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

*/
