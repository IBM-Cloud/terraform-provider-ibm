package ibm

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	if userDetails.generation == 1 {
		err := classicImageList(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := imageList(d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicImageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcclassicv1.Image{}
	for {
		listImagesOptions := &vpcclassicv1.ListImagesOptions{
			Start: &start,
		}
		availableImages, _, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return err
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.Crn,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

func imageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.Image{}
	for {
		listImagesOptions := &vpcv1.ListImagesOptions{
			Start: &start,
		}
		availableImages, _, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return err
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.Crn,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

// dataSourceIBMISImagesId returns a reasonable ID for a image list.
func dataSourceIBMISImagesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
