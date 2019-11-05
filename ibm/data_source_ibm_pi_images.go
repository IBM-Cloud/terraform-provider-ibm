package ibm

import (
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform/helper/validation"
)

/*
Datasource to get the list of images that are available when a power instance is created

*/
func dataSourceIBMPIImages() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIImagesAllRead,
		Schema: map[string]*schema.Schema{

			helpers.PIImageName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"image_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_href": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_state": {
							Type:     schema.TypeString,
							Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIImagesAllRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	imagename := d.Get(helpers.PIImageName).(string)

	imageC := instance.NewIBMPIImageClient(sess, powerinstanceid)

	imagedata, err := imageC.GetAll(powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	_ = d.Set("image_info", flattenStockImages(imagename, imagedata.Images))

	return nil

}

func flattenStockImages(imagename string, list []*models.ImageReference) []map[string]interface{} {
	log.Printf("Calling the flattenstockImages method and the size is %d", len(list))
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		if *i.Name == imagename {
			l := map[string]interface{}{
				"image_id":    *i.ImageID,
				"image_state": *i.State,
				"image_href":  *i.Href,
				"image_name":  *i.Name,
			}

			result = append(result, l)
		}
	}
	return result
}
