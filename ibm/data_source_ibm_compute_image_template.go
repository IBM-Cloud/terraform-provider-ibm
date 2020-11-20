package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"sort"
	"strings"
)

// ByCreateDate implements sort.Interface for []ImageTemplate based on
// the CreateField field in desc order.
type ByCreateDate []ImageTemplate

func (a ByCreateDate) Len() int      { return len(a) }
func (a ByCreateDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCreateDate) Less(i, j int) bool {
	return a[i].CreateDate.UnixNano() > a[j].CreateDate.UnixNano()
}

type ImageTemplate struct {
	ID         *int
	Name       *string
	CreateDate *datatypes.Time
}

func dataSourceIBMComputeImageTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMComputeImageTemplateRead,

		// TODO: based on need add properties for visibility, type of image,
		// notes, size, shared on accounts if needed
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The internal id of the image template",
				Type:        schema.TypeInt,
				Computed:    true,
			},

			"name": {
				Description: "The name of this image template - can match partially or fully",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMComputeImageTemplateRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	name := d.Get("name").(string)

	imageTemplates, err := service.
		Mask("id,name,createDate").
		GetBlockDeviceTemplateGroups()
	if err != nil {
		return fmt.Errorf("Error looking up image template [%s]: %s", name, err)
	}

	var matchingImageTemplates []ImageTemplate
	for _, imageTemplate := range imageTemplates {
		if imageTemplate.Name != nil && strings.Contains(*imageTemplate.Name, name) {
			matchingImageTemplates = append(matchingImageTemplates, ImageTemplate{
				ID:         imageTemplate.Id,
				Name:       imageTemplate.Name,
				CreateDate: imageTemplate.CreateDate,
			})
		}
	}

	if len(matchingImageTemplates) > 0 {
		// sort and pick newest image
		sort.Sort(ByCreateDate(matchingImageTemplates))
		d.SetId(fmt.Sprintf("%d", *matchingImageTemplates[0].ID))
		return nil
	}

	// Image not found among private nor shared images in the account.
	// Looking up in the public images
	templateService := services.GetVirtualGuestBlockDeviceTemplateGroupService(sess)
	pubImageTemplates, err := templateService.
		Mask("id,name").
		Filter(filter.Path("name").Contains(name).Build()).
		GetPublicImages()
	if err != nil {
		return fmt.Errorf("Error looking up image template [%s] among the public images: %s", name, err)
	}

	matchingImageTemplates = []ImageTemplate{}
	if len(pubImageTemplates) > 0 {
		for _, imageTemplate := range pubImageTemplates {
			matchingImageTemplates = append(matchingImageTemplates, ImageTemplate{
				ID:         imageTemplate.Id,
				Name:       imageTemplate.Name,
				CreateDate: imageTemplate.CreateDate,
			})
		}
		// sort and pick newest image
		sort.Sort(ByCreateDate(matchingImageTemplates))
		imageTemplate := matchingImageTemplates[0]
		d.SetId(fmt.Sprintf("%d", *imageTemplate.ID))
		return nil
	}

	return fmt.Errorf("Could not find image template with name [%s]", name)
}
