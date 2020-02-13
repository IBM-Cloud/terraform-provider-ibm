package ibm

import (
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isImageHref            = "href"
	isImageName            = "name"
	isImageTags            = "tags"
	isImageOperatingSystem = "operating_system"
	isImageStatus          = "status"
	isImageVisibility      = "visibility"
	isImageFile            = "file"
	isImageFormat          = "format"
	isImageArchitecure     = "architecture"
	isImageResourceGroup   = "resource_group"

	isImageProvisioning     = "provisioning"
	isImageProvisioningDone = "done"
	isImageDeleting         = "deleting"
	isImageDeleted          = "done"
)

func resourceIBMISImage() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISImageCreate,
		Read:     resourceIBMISImageRead,
		Update:   resourceIBMISImageUpdate,
		Delete:   resourceIBMISImageDelete,
		Exists:   resourceIBMISImageExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isImageHref: {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: applyOnce,
			},

			isImageName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
			},

			isImageTags: {
				Type:             schema.TypeSet,
				Optional:         true,
				Computed:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              resourceIBMVPCHash,
				DiffSuppressFunc: applyOnce,
			},

			isImageOperatingSystem: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isImageStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isImageArchitecure: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isImageVisibility: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isImageFile: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isImageFormat: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isImageResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISImageCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Image create")
	href := d.Get(isImageHref).(string)
	name := d.Get(isImageName).(string)
	operatingSystem := d.Get(isImageOperatingSystem).(string)

	var rg string
	if grp, ok := d.GetOk(isImageResourceGroup); ok {
		rg = grp.(string)
	}
	imageC := compute.NewImageClient(sess)
	image, err := imageC.Create(href, name, operatingSystem, rg)
	if err != nil {
		log.Printf("[DEBUG] Image err %s", err)
		return err
	}

	d.SetId(image.ID.String())
	log.Printf("[INFO] Image : %s", image.ID.String())

	_, err = isWaitForImageAvailable(imageC, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isImageTags); ok || v != "" {
		oldList, newList := d.GetChange(isImageTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, image.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc image (%s) tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISImageRead(d, meta)
}

func isWaitForImageAvailable(imageC *compute.ImageClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for image (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isImageProvisioning},
		Target:     []string{isImageProvisioningDone},
		Refresh:    isImageRefreshFunc(imageC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isImageRefreshFunc(imageC *compute.ImageClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		image, err := imageC.Get(id)
		if err != nil {
			return nil, "", err
		}

		if image.Status == "available" || image.Status == "failed" {
			return image, isImageProvisioningDone, nil
		}

		return image, isImageProvisioning, nil
	}
}

func resourceIBMISImageUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	imageC := compute.NewImageClient(sess)

	image, err := imageC.Get(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange(isImageName) {
		name := d.Get(isImageName).(string)
		_, err := imageC.Update(d.Id(), name)
		if err != nil {
			return err
		}
	}
	if d.HasChange(isImageTags) {
		oldList, newList := d.GetChange(isImageTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, image.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Image (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISImageRead(d, meta)
}

func resourceIBMISImageRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	imageC := compute.NewImageClient(sess)

	image, err := imageC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set("id", image.ID.String())
	d.Set(isImageArchitecure, image.Architecture)
	d.Set(isImageName, image.Name)
	d.Set(isImageOperatingSystem, image.OperatingSystem)
	d.Set(isImageFormat, image.Format)
	d.Set(isImageFile, image.File)
	d.Set(isImageHref, image.Href)
	d.Set(isImageStatus, image.Status)
	d.Set(isImageVisibility, image.Visibility)
	tags, err := GetTagsUsingCRN(meta, image.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc Image (%s) tags: %s", d.Id(), err)
	}
	d.Set(isImageTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/compute/image")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/compute/image")
	}
	d.Set(ResourceName, image.Name)
	d.Set(ResourceStatus, image.Status)
	d.Set(ResourceCRN, image.Crn)
	if image.ResourceGroup != nil {
		d.Set(ResourceGroupName, image.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMISImageDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	imageC := compute.NewImageClient(sess)
	err = imageC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForImageDeleted(imageC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForImageDeleted(imageC *compute.ImageClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for image (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isImageDeleting},
		Target:     []string{},
		Refresh:    isImageDeleteRefreshFunc(imageC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isImageDeleteRefreshFunc(imageC *compute.ImageClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		image, err := imageC.Get(id)
		if err == nil {
			return image, isImageDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("[DEBUG] returning deleted")
				return nil, isImageDeleted, nil
			}
		}
		return nil, isImageDeleting, err
	}
}

func resourceIBMISImageExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	imageC := compute.NewImageClient(sess)

	_, err = imageC.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
