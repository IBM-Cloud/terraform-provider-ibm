// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"

	"errors"
	"fmt"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const captDestination string = "image-catalog"

func resourceIBMPICapture() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPICaptureCreate,
		ReadContext:   resourceIBMPICaptureRead,
		UpdateContext: resourceIBMPICaptureUpdate,
		DeleteContext: resourceIBMPICaptureDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			helpers.PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance Name of the Power VM",
			},

			helpers.PIInstanceCaptureName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the capture to create. Note : this must be unique",
			},

			helpers.PIInstanceCaptureDestination: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of destination to store the image capture to",
				ValidateFunc: validateAllowedStringValue([]string{"image-catalog", "cloud-storage", "both"}),
			},

			helpers.PIInstanceCaptureVolumeIds: {
				Type:             schema.TypeSet,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "List of PI volumes",
			},

			helpers.PIInstanceCaptureCloudStorageRegion: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "List of Regions to use",
				ValidateFunc: validateAllowedStringValue([]string{"us-south", "us-east", "us-de"}),
			},

			helpers.PIInstanceCaptureCloudStorageAccessKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of Cloud Storage Access Key",
			},
			helpers.PIInstanceCaptureCloudStorageSecretKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Cloud Storage Secret Key",
			},
			helpers.PIInstanceCaptureCloudStorageImagePath: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloud Object Storage bucket name; bucket-name[/optional/folder]",
			},
			// Computed Attribute
			"image_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Capture ID",
			},
		},
	}
}

func resourceIBMPICaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get(helpers.PIInstanceName).(string)
	capturename := d.Get(helpers.PIInstanceCaptureName).(string)
	capturedestination := d.Get(helpers.PIInstanceCaptureDestination).(string)
	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)

	captureBody := &models.PVMInstanceCapture{
		CaptureDestination: &capturedestination,
		CaptureName:        &capturename,
	}
	if v, ok := d.GetOk(helpers.PIInstanceCaptureCloudStorageRegion); ok {
		CloudStorageRegion := v.(string)
		captureBody.CloudStorageRegion = CloudStorageRegion
	}

	if v, ok := d.GetOk(helpers.PIInstanceCaptureVolumeIds); ok {
		volids := expandStringList((v.(*schema.Set)).List())
		if len(volids) > 0 {
			captureBody.CaptureVolumeIds = volids
		}
	}

	if v, ok := d.GetOk(helpers.PIInstanceCaptureCloudStorageAccessKey); ok {
		captureBody.CloudStorageAccessKey = v.(string)
	}
	if v, ok := d.GetOk(helpers.PIInstanceCaptureCloudStorageImagePath); ok {
		captureBody.CloudStorageImagePath = v.(string)
	}
	if v, ok := d.GetOk(helpers.PIInstanceCaptureCloudStorageSecretKey); ok {
		captureBody.CloudStorageSecretKey = v.(string)
	}

	captureResponse, err := client.CaptureInstanceToImageCatalogV2(name, captureBody)

	if err != nil {
		return diag.FromErr(err)
	}

	jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)
	_, err = waitForIBMPIJobCompleted(ctx, jobClient, *captureResponse.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	if capturedestination == captDestination {
		imageClient := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
		image, err := imageClient.Get(capturename)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *image.ImageID))
		return resourceIBMPICaptureRead(ctx, d, meta)
	}
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, capturename))
	return nil
}

func resourceIBMPICaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID, captureID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	capturedestination := d.Get(helpers.PIInstanceCaptureDestination).(string)
	if capturedestination == captDestination {
		imageClient := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
		imagedata, err := imageClient.Get(captureID)
		if err != nil {
			uErr := errors.Unwrap(err)
			switch uErr.(type) {
			case *p_cloud_images.PcloudCloudinstancesImagesGetNotFound:
				log.Printf("[DEBUG] image does not exist %v", err)
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] get image failed %v", err)
			return diag.FromErr(err)
		}
		imageid := *imagedata.ImageID
		d.Set("image_id", imageid)
		d.Set(helpers.PICloudInstanceId, cloudInstanceID)
	}
	return nil
}

func resourceIBMPICaptureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIBMPICaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, captureID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	capturedestination := d.Get(helpers.PIInstanceCaptureDestination).(string)
	if capturedestination == captDestination {
		imageClient := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
		err = imageClient.Delete(captureID)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return nil
}

func isWaitForImageCaptureAvailable(client *st.IBMPIInstanceClient, s string, s2 string, timeout time.Duration) (interface{}, error) {

	return nil, nil
}
