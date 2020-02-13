package ibm

import (
	"log"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	st "github.ibm.com/Bluemix/riaas-go-client/riaas/client/storage"
)

const (
	isVolumeName             = "name"
	isVolumeProfileName      = "profile"
	isVolumeZone             = "zone"
	isVolumeEncryptionKey    = "encryption_key"
	isVolumeCapacity         = "capacity"
	isVolumeIops             = "iops"
	isVolumeCrn              = "crn"
	isVolumeTags             = "tags"
	isVolumeStatus           = "status"
	isVolumeDeleting         = "deleting"
	isVolumeDeleted          = "done"
	isVolumeProvisioning     = "provisioning"
	isVolumeProvisioningDone = "done"
	isVolumeResourceGroup    = "resource_group"
)

func resourceIBMISVolume() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVolumeCreate,
		Read:     resourceIBMISVolumeRead,
		Update:   resourceIBMISVolumeUpdate,
		Delete:   resourceIBMISVolumeDelete,
		Exists:   resourceIBMISVolumeExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			isVolumeName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateISName,
			},

			isVolumeProfileName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVolumeZone: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVolumeEncryptionKey: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			isVolumeCapacity: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
				ForceNew: true,
			},
			isVolumeResourceGroup: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			isVolumeIops: {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			isVolumeCrn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isVolumeStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVolumeTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      resourceIBMVPCHash,
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

func resourceIBMISVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	volName := d.Get(isVolumeName).(string)
	profile := d.Get(isVolumeProfileName).(string)
	zone := d.Get(isVolumeZone).(string)
	volCapacity := int64(d.Get(isVolumeCapacity).(int))
	client := storage.NewStorageClient(sess)

	volZone := &st.PostVolumesParamsBodyZone{
		Name: &zone,
	}
	volProfile := &st.PostVolumesParamsBodyProfile{
		Name: profile,
	}

	body := st.PostVolumesBody{
		Name:     volName,
		Zone:     volZone,
		Profile:  volProfile,
		Capacity: &volCapacity,
	}

	var encryptionKey string
	if key, ok := d.GetOk(isVolumeEncryptionKey); ok {
		encryptionKey = key.(string)
		volEncryptionKey := st.PostVolumesParamsBodyEncryptionKey{
			Crn: encryptionKey,
		}
		body.EncryptionKey = &volEncryptionKey
	}

	if rg, ok := d.GetOk(isVolumeResourceGroup); ok {
		rgref := st.PostVolumesParamsBodyResourceGroup{
			ID: strfmt.UUID(rg.(string)),
		}
		body.ResourceGroup = &rgref
	}

	if iops, ok := d.GetOk(isVolumeIops); ok {
		i := int64(iops.(int))
		body.Iops = &i
	}

	vol, err := client.Create(&st.PostVolumesParams{
		Body: body,
	})
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	d.SetId(vol.ID.String())
	log.Printf("[INFO]  : %s", vol.ID.String())

	_, err = isWaitForVolumeAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVolumeTags); ok || v != "" {
		oldList, newList := d.GetChange(isVolumeTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, vol.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc volume (%s) tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := storage.NewStorageClient(sess)

	vol, err := client.Get(d.Id())
	if err != nil {
		return err
	}

	d.SetId(vol.ID.String())
	d.Set(isVolumeName, vol.Name)
	d.Set(isVolumeProfileName, vol.Profile.Name)
	d.Set(isVolumeZone, vol.Zone.Name)
	if vol.EncryptionKey != nil {
		d.Set(isVolumeEncryptionKey, vol.EncryptionKey.Crn)
	}
	d.Set(isVolumeIops, vol.Iops)
	d.Set(isVolumeCapacity, vol.Capacity)
	d.Set(isVolumeCrn, vol.Crn)
	d.Set(isVolumeResourceGroup, vol.ResourceGroup.ID)
	d.Set(isVolumeStatus, vol.Status)
	tags, err := GetTagsUsingCRN(meta, vol.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc volume (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVolumeTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/storage/storageVolumes")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/storage/storageVolumes")
	}
	d.Set(ResourceName, vol.Name)
	d.Set(ResourceCRN, vol.Crn)
	d.Set(ResourceStatus, vol.Status)
	if vol.ResourceGroup != nil {
		d.Set(ResourceGroupName, vol.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMISVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := storage.NewStorageClient(sess)

	vol, err := client.Get(d.Id())
	if err != nil {
		return err
	}

	// Generating parameters for

	if d.HasChange(isVolumeName) {
		body := st.PatchVolumesIDBody{
			Name: d.Get(isVolumeName).(string),
		}

		patchVolParms := &st.PatchVolumesIDParams{
			Body: body,
		}
		_, err := client.Update(d.Id(), patchVolParms)
		if err != nil {
			return err
		}
	}
	if d.HasChange(isVolumeTags) {
		oldList, newList := d.GetChange(isVolumeTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, vol.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc volume (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := storage.NewStorageClient(sess)
	err = client.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForVolumeDeleted(client, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForVolumeDeleted(vol *storage.StorageClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeDeleting},
		Target:     []string{},
		Refresh:    isVolumeDeleteRefreshFunc(vol, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVolumeDeleteRefreshFunc(vol *storage.StorageClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := vol.Get(id)
		if err == nil {
			return vol, isVolumeDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "volume_id_not_found" {
				return nil, isVolumeDeleted, nil
			}
		}
		return nil, isVolumeDeleting, err
	}
}

func resourceIBMISVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	client := storage.NewStorageClient(sess)

	_, err = client.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "volume_id_not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func isWaitForVolumeAvailable(client *storage.StorageClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeProvisioning},
		Target:     []string{isVolumeProvisioningDone},
		Refresh:    isVolumeRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVolumeRefreshFunc(client *storage.StorageClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.Status == "available" {
			return vol, isVolumeProvisioningDone, nil
		}

		return vol, isVolumeProvisioning, nil
	}
}
