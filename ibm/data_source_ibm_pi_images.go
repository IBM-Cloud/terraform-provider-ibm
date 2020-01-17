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
				Optional:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
				Deprecated:   "This field is deprectaed.",
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

	imageC := instance.NewIBMPIImageClient(sess, powerinstanceid)

	imagedata, err := imageC.GetAll(powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	_ = d.Set("image_info", flattenStockImages(imagedata.Images))

	return nil

}

func flattenStockImages(list []*models.ImageReference) []map[string]interface{} {
	log.Printf("Calling the flattenstockImages method and the size is %d", len(list))
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"image_id":    *i.ImageID,
			"image_state": *i.State,
			"image_href":  *i.Href,
			"image_name":  *i.Name,
		}

		result = append(result, l)
	}
	return result
}
