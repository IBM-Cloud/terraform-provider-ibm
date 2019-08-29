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

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"icn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"tenantinstances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationdate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"tenantid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloudinstances": {
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
						"os": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"architecture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"crn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"sshkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operatingsystem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
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
	d.Set("icn", tenantData.TenantID)
	d.Set("creationdate", tenantData.CreationDate)
	d.Set("sshkey", tenantData.SSHKeys[0].Name)
	d.Set("enabled", tenantData.Enabled)
	d.Set("tenantinstances", flattenCloudInstances(tenantData.CloudInstances))
	log.Printf("Printing the tenant data %s", tenantData.CloudInstances)
	/*if tenantData.CloudInstances != nil {
		tenantInstances := make([]map[string]interface{}, len(tenantData.CloudInstances))
		for i, tenantinfo := range tenantData.CloudInstances {

			p := make(map[string]interface{})
			p["name"] = tenantinfo.Name
			p["region"] = tenantinfo.Region
			p["memory"] = tenantinfo.Limits.Memory

			tenantInstances[i] = p
		}
		d.Set("tenantinstances", tenantInstances)

	}*/

	return nil

}

/*
groups := make([]map[string]interface{}, len(grouplist.Groups))
	for i, group := range grouplist.Groups {
		memorys := make([]map[string]interface{}, 1)
		memory := make(map[string]interface{})
		memory["units"] = group.Memory.Units
		memory["allocation_mb"] = group.Memory.AllocationMb
		memory["minimum_mb"] = group.Memory.MinimumMb
		memory["step_size_mb"] = group.Memory.StepSizeMb
		memory["is_adjustable"] = group.Memory.IsAdjustable
		memory["can_scale_down"] = group.Memory.CanScaleDown
		memorys[0] = memory



		l := map[string]interface{}{
			"group_id": group.Id,
			"count":    group.Count,
			"memory":   memorys,
			"cpu":      cpus,
			"disk":     disks,
		}
		groups[i] = l
	}
	return groups
*/

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
