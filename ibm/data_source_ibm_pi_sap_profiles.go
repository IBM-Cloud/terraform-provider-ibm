package ibm

import (
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPISAPProfiles() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPISAPProfilesRead,
		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"sap_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cores": {
							Type:     schema.TypeInt,
							Computed: true,
						}, "certified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPISAPProfilesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	sapprofile := instance.NewIBMPIInstanceClient(sess, powerinstanceid)
	sapprofiledata, err := sapprofile.GetSAPProfiles(powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	d.Set("sap_profiles", flattensapprofile(sapprofiledata.Profiles))

	return nil

}

func flattensapprofile(list []*models.SAPProfile) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"profile_id":   *i.ProfileID,
			"memory":       *i.Memory,
			"certified":    *i.Certified,
			"profile_type": *i.Type,
			"cores":        *i.Cores,
		}
		result = append(result, l)
	}
	return result
}
