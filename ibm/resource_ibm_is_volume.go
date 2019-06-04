package ibm

import (
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	st "github.ibm.com/riaas/rias-api/riaas/client/storage"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isVolumeName             = "name"
	isVolumeProfileName      = "profile"
	isVolumeZone             = "zone"
	isVolumeEncryptionKey    = "encryption_key"
	isVolumeCapacity         = "capacity"
	isVolumeIops             = "iops"
	isVolumeCrn              = "crn"
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

		Schema: map[string]*schema.Schema{

			isVolumeName: {
				Type:     schema.TypeString,
				Required: true,
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
	volCapacity := d.Get(isVolumeCapacity).(int)
	client := storage.NewStorageClient(sess)

	volZone := &models.PostVolumesParamsBodyZone{
		Name: &zone,
	}
	volProfile := &models.PostVolumesParamsBodyProfile{
		Name: profile,
	}

	body := &models.PostVolumesParamsBody{
		Name:     volName,
		Zone:     volZone,
		Profile:  volProfile,
		Capacity: int64(volCapacity),
	}

	var encryptionKey string
	if key, ok := d.GetOk(isVolumeEncryptionKey); ok {
		encryptionKey = key.(string)
		volEncryptionKey := models.PostVolumesParamsBodyEncryptionKey{
			Crn: encryptionKey,
		}
		body.EncryptionKey = &volEncryptionKey
	}

	if rg, ok := d.GetOk(isVolumeResourceGroup); ok {
		rgref := models.PostVolumesParamsBodyResourceGroup{
			ID: strfmt.UUID(rg.(string)),
		}
		body.ResourceGroup = &rgref
	}

	if iops, ok := d.GetOk(isVolumeIops); ok {
		body.Iops = iops.(int64)
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
	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
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

	return nil
}

func resourceIBMISVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := storage.NewStorageClient(sess)

	// Generating parameters for

	if d.HasChange(isVolumeName) {
		body := &models.PatchVolumesIDParamsBody{
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

	return resourceIBMISVolumeRead(d, meta)
}

func resourceIBMISVolumeDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := storage.NewStorageClient(sess)
	err := client.Delete(d.Id())
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
				iserror.Payload.Errors[0].Code == "volume_not_found" {
				return nil, isVolumeDeleted, nil
			}
		}
		return nil, isVolumeDeleting, err
	}
}

func resourceIBMISVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).ISSession()
	client := storage.NewStorageClient(sess)

	_, err := client.Get(d.Id())
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
