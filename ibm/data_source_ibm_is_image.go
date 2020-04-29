package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func dataSourceIBMISImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISImageRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"public", "private"}),
			},

			"status": {
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
	}
}

func dataSourceIBMISImageRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}
	if userDetails.generation == 1 {
		err := classicImageGet(d, meta, name, visibility)
		if err != nil {
			return err
		}
	} else {
		err := imageGet(d, meta, name, visibility)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicImageGet(d *schema.ResourceData, meta interface{}, name, visibility string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	listImagesOptions := &vpcclassicv1.ListImagesOptions{}
	if visibility != "" {
		listImagesOptions.Visibility = &visibility
	}
	images, _, err := sess.ListImages(listImagesOptions)
	if err != nil {
		return err
	}
	for _, image := range images.Images {
		if *image.Name == name {
			d.SetId(*image.ID)
			d.Set("status", *image.Status)
			d.Set("name", *image.Name)
			d.Set("visibility", *image.Visibility)
			d.Set("os", *image.OperatingSystem.Name)
			d.Set("architecture", *image.OperatingSystem.Architecture)
			d.Set("crn", *image.Crn)
			return nil
		}
	}
	return fmt.Errorf("No Image found with name %s", name)
}

func imageGet(d *schema.ResourceData, meta interface{}, name, visibility string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listImagesOptions := &vpcv1.ListImagesOptions{}
	if visibility != "" {
		listImagesOptions.Visibility = &visibility
	}
	images, _, err := sess.ListImages(listImagesOptions)
	if err != nil {
		return err
	}
	for _, image := range images.Images {
		if *image.Name == name {
			d.SetId(*image.ID)
			d.Set("status", *image.Status)
			d.Set("name", *image.Name)
			d.Set("visibility", *image.Visibility)
			d.Set("os", *image.OperatingSystem.Name)
			d.Set("architecture", *image.OperatingSystem.Architecture)
			d.Set("crn", *image.Crn)
			return nil
		}
	}
	return nil
}
