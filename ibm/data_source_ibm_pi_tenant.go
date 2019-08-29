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

			"cloudinstances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudinstancereferences": {
							Type:     schema.TypeList,
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
									"limits": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
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
	d.Set("creationdate", tenantData.CreationDate.String())
	d.Set("enabled", tenantData.Enabled)
	d.Set("cloudinstances", flattenCloudInstances(tenantData.CloudInstances))
	log.Printf("Printing the tenant data %s", tenantData.CloudInstances)

	return nil

}

func flattenCloudInstances(cloudinstances []*models.CloudInstanceReference) []map[string]interface{} {
	cloudInstances := make([]map[string]interface{}, len(cloudinstances))
	for i, cloudinstance := range cloudinstances {
		cloudinstancereferences := make([]map[string]interface{}, 1)
		cloudinstanceref := make(map[string]interface{})

		cloudinstanceref["cloudinstanceid"] = cloudinstance.CloudInstanceID
		cloudinstanceref["name"] = cloudinstance.Name
		cloudinstanceref["region"] = cloudinstance.Region
		cloudinstanceref["href"] = cloudinstance.Href

		cloudinstancereferences[0] = cloudinstanceref
		l := map[string]interface{}{

			//"cloudinstanceid": cloudinstance.CloudInstanceID,
			//"count":    group.Count,
			"cloudinstances": cloudinstances,
			//"cpu":      cpus,
			//"disk":     disks,
		}
		cloudInstances[i] = l
	}
	log.Printf("printing the cloudinstances %+v", cloudInstances)
	return cloudInstances
}
