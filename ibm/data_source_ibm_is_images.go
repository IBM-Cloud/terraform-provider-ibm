package ibm

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
)

const (
	isImages = "images"
)

func dataSourceIBMISImages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISImagesRead,

		Schema: map[string]*schema.Schema{

			isImages: {
				Type:        schema.TypeList,
				Description: "List of images",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"visibility": {
							Type:     schema.TypeString,
							Computed: true,
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
		},
	}
}

func dataSourceIBMISImagesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	clientC := compute.NewImageClient(sess)
	availableImages, _, err := clientC.List("")
	if err != nil {
		return err
	}

	images := make([]map[string]string, len(availableImages))
	for i, image := range availableImages {

		img := make(map[string]string)
		img["name"] = image.Name
		img["id"] = string(image.ID)
		img["status"] = image.Status
		img["visibility"] = image.Visibility
		img["os"] = image.OperatingSystem.Name
		img["architecture"] = image.Architecture
		img["crn"] = image.Crn

		images[i] = img
	}
	d.SetId(dataSourceIBMISImagesID(d))
	d.Set(isImages, images)
	return nil
}

// dataSourceIBMISImagesId returns a reasonable ID for a image list.
func dataSourceIBMISImagesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
