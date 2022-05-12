// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const cloudStorageDestination string = "cloud-storage"
const imageCatalogDestination string = "image-catalog"

func ResourceIBMPICapture() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPICaptureCreate,
		ReadContext:   resourceIBMPICaptureRead,
		DeleteContext: resourceIBMPICaptureDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(75 * time.Minute),
			Delete: schema.DefaultTimeout(50 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance Name of the Power VM",
			},

			PICaptureName: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the capture to create. Note : this must be unique",
			},

			PICaptureDestination: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Destination for the deployable image",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"image-catalog", "cloud-storage", "both"}),
			},

			PICaptureVolumeIDs: {
				Type:             schema.TypeSet,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				ForceNew:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "List of Data volume IDs",
			},

			PICaptureCloudStorageRegion: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "List of Regions to use",
			},

			PICaptureCloudStorageAccessKey: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Name of Cloud Storage Access Key",
			},
			PICaptureCloudStorageSecretKey: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Name of the Cloud Storage Secret Key",
			},
			PICaptureStorageImagePath: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Cloud Storage Image Path (bucket-name [/folder/../..])",
			},
			// Computed Attribute
			CaptureImageID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image ID of Capture Instance",
			},
		},
	}
}

func resourceIBMPICaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get(PIInstanceName).(string)
	capturename := d.Get(PICaptureName).(string)
	capturedestination := d.Get(PICaptureDestination).(string)
	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)

	captureBody := &models.PVMInstanceCapture{
		CaptureDestination: &capturedestination,
		CaptureName:        &capturename,
	}
	if capturedestination != imageCatalogDestination {
		if v, ok := d.GetOk(PICaptureCloudStorageRegion); ok {
			captureBody.CloudStorageRegion = v.(string)
		} else {
			return diag.Errorf("%s is required when capture destination is %s", PICaptureCloudStorageRegion, capturedestination)
		}
		if v, ok := d.GetOk(PICaptureCloudStorageAccessKey); ok {
			captureBody.CloudStorageAccessKey = v.(string)
		} else {
			return diag.Errorf("%s is required when capture destination is %s ", PICaptureCloudStorageAccessKey, capturedestination)
		}
		if v, ok := d.GetOk(PICaptureStorageImagePath); ok {
			captureBody.CloudStorageImagePath = v.(string)
		} else {
			return diag.Errorf("%s is required when capture destination is %s ", PICaptureStorageImagePath, capturedestination)
		}
		if v, ok := d.GetOk(PICaptureCloudStorageSecretKey); ok {
			captureBody.CloudStorageSecretKey = v.(string)
		} else {
			return diag.Errorf("%s is required when capture destination is %s ", PICaptureCloudStorageSecretKey, capturedestination)
		}
	}

	if v, ok := d.GetOk(PICaptureVolumeIDs); ok {
		volids := flex.ExpandStringList((v.(*schema.Set)).List())
		if len(volids) > 0 {
			captureBody.CaptureVolumeIDs = volids
		}
	}

	captureResponse, err := client.CaptureInstanceToImageCatalogV2(name, captureBody)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, capturename, capturedestination))
	jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)
	_, err = waitForIBMPIJobCompleted(ctx, jobClient, *captureResponse.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIBMPICaptureRead(ctx, d, meta)
}

func resourceIBMPICaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := parts[0]
	captureID := parts[1]
	capturedestination := parts[2]
	if capturedestination != cloudStorageDestination {
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
		d.Set(CaptureImageID, imageid)
	}
	d.Set(PICloudInstanceID, cloudInstanceID)
	return nil
}

func resourceIBMPICaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := parts[0]
	captureID := parts[1]
	capturedestination := parts[2]
	if capturedestination != cloudStorageDestination {
		imageClient := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
		err = imageClient.Delete(captureID)
		if err != nil {
			uErr := errors.Unwrap(err)
			switch uErr.(type) {
			case *p_cloud_images.PcloudCloudinstancesImagesGetNotFound:
				log.Printf("[DEBUG] image does not exist while deleting %v", err)
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] delete image failed %v", err)
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return nil
}
