package ibm

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
)

const (
	dlSpeeds       = "offering_speeds"
	dlLinkSpeed    = "link_speed"
	dlOfferingType = "offering_type"
)

func dataSourceIBMDLSpeedOptions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDLSpeedOptionsRead,
		Schema: map[string]*schema.Schema{
			dlOfferingType: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The Direct Link offering type",
				ValidateFunc: InvokeValidator("ibm_dl_speed_options", dlOfferingType),
			},
			dlSpeeds: {
				Type:        schema.TypeList,
				Description: "Collection of direct link speeds",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlLinkSpeed: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Direct Link offering speed for the specified offering type",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDLSpeedOptionsRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}
	dlType := d.Get(dlOfferingType).(string)
	listSpeedsOptionsModel := &directlinkapisv1.ListOfferingTypeSpeedsOptions{}
	listSpeedsOptionsModel.OfferingType = &dlType
	listSpeeds, _, err := directLink.ListOfferingTypeSpeeds(listSpeedsOptionsModel)

	if err != nil {
		return err
	}
	speeds := make([]map[string]interface{}, 0)
	for _, instance := range listSpeeds.Speeds {
		speed := map[string]interface{}{}
		if instance.LinkSpeed != nil {
			speed[dlLinkSpeed] = *instance.LinkSpeed
			log.Println("Link Speed ", *instance.LinkSpeed)
		}
		log.Println("Speed ", speed)
		speeds = append(speeds, speed)
	}
	d.SetId(dataSourceIBMDLSpeedOptionsID(d))
	d.Set(dlSpeeds, speeds)
	return nil
}

// dataSourceIBMDLSpeedOptionsID returns a reasonable ID for a direct link speeds list.
func dataSourceIBMDLSpeedOptionsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func datasourceIBMDLSpeedOptionsValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 2)
	dlTypeAllowedValues := "dedicated, connect"

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 dlOfferingType,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              dlTypeAllowedValues})

	ibmDLSpeedOptionsDatasourceValidator := ResourceValidator{ResourceName: "ibm_dl_speed_options", Schema: validateSchema}
	return &ibmDLSpeedOptionsDatasourceValidator
}
