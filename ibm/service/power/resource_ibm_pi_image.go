// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

// Attributes and Arguments defined in data_source_ibm_pi_image.go
func ResourceIBMPIImage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIImageCreate,
		ReadContext:   resourceIBMPIImageRead,
		DeleteContext: resourceIBMPIImageDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
				ForceNew:    true,
			},
			PIImageName: {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Image name",
				DiffSuppressFunc: flex.ApplyOnce,
				ForceNew:         true,
			},
			PIImageID: {
				Type:             schema.TypeString,
				Optional:         true,
				ExactlyOneOf:     []string{PIImageID, PIImageBucketName},
				Description:      "Instance image id",
				DiffSuppressFunc: flex.ApplyOnce,
				ConflictsWith:    []string{PIImageBucketName},
				ForceNew:         true,
			},

			// COS import variables
			PIImageBucketName: {
				Type:          schema.TypeString,
				Optional:      true,
				ExactlyOneOf:  []string{PIImageID, PIImageBucketName},
				Description:   "Cloud Object Storage bucket name; bucket-name[/optional/folder]",
				ConflictsWith: []string{PIImageID},
				RequiredWith:  []string{PIImageBucketRegion, PIImageBucketFile},
				ForceNew:      true,
			},
			PIImageBucketAccess: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Indicates if the bucket has public or private access",
				Default:       "public",
				ValidateFunc:  validate.ValidateAllowedStringValues([]string{"public", "private"}),
				ConflictsWith: []string{PIImageID},
				ForceNew:      true,
			},
			PIImageAccessKey: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Cloud Object Storage access key; required for buckets with private access",
				ForceNew:     true,
				Sensitive:    true,
				RequiredWith: []string{PIImageSecretKey},
			},
			PIImageSecretKey: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Cloud Object Storage secret key; required for buckets with private access",
				ForceNew:     true,
				Sensitive:    true,
				RequiredWith: []string{PIImageAccessKey},
			},
			PIImageBucketRegion: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Cloud Object Storage region",
				ConflictsWith: []string{PIImageID},
				RequiredWith:  []string{PIImageBucketName},
				ForceNew:      true,
			},
			PIImageBucketFile: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Cloud Object Storage image filename",
				ConflictsWith: []string{PIImageID},
				RequiredWith:  []string{PIImageBucketName},
				ForceNew:      true,
			},
			ImageStorageType: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of storage",
				ForceNew:    true,
			},
			ImageStoragePool: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Storage pool where the image will be loaded, if provided then pi_image_storage_type and pi_affinity_policy will be ignored",
				ForceNew:    true,
			},
			PIImageAffinityPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Affinity policy for image; ignored if pi_image_storage_pool provided; for policy affinity requires one of pi_affinity_instance or pi_affinity_volume to be specified; for policy anti-affinity requires one of pi_anti_affinity_instances or pi_anti_affinity_volumes to be specified",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"affinity", "anti-affinity"}),
				ForceNew:     true,
			},
			PIImageAffinityVolume: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Volume (ID or Name) to base storage affinity policy against; required if requesting affinity and pi_affinity_instance is not provided",
				ConflictsWith: []string{PIImageAffinityInstance},
				ForceNew:      true,
			},
			PIImageAffinityInstance: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "PVM Instance (ID or Name) to base storage affinity policy against; required if requesting storage affinity and pi_affinity_volume is not provided",
				ConflictsWith: []string{PIImageAffinityVolume},
				ForceNew:      true,
			},
			PIImageAntiAffinityVolumes: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of volumes to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_instances is not provided",
				ConflictsWith: []string{PIImageAntiAffinityInstances},
				ForceNew:      true,
			},
			PIImageAntiAffinityInstances: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of pvmInstances to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_volumes is not provided",
				ConflictsWith: []string{PIImageAntiAffinityVolumes},
				ForceNew:      true,
			},

			// Computed Attribute
			ImageID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image ID",
			},
		},
	}
}

func resourceIBMPIImageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		log.Printf("Failed to get the session")
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	imageName := d.Get(PIImageName).(string)

	client := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	// image copy
	if v, ok := d.GetOk(PIImageID); ok {
		imageid := v.(string)
		source := "root-project"
		var body = &models.CreateImage{
			ImageName: imageName,
			ImageID:   imageid,
			Source:    &source,
		}
		imageResponse, err := client.Create(body)
		if err != nil {
			return diag.FromErr(err)
		}

		IBMPIImageID := imageResponse.ImageID
		d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *IBMPIImageID))

		_, err = isWaitForIBMPIImageAvailable(ctx, client, *IBMPIImageID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			log.Printf("[DEBUG]  err %s", err)
			return diag.FromErr(err)
		}
	}

	// COS image import
	if v, ok := d.GetOk(PIImageBucketName); ok {
		bucketName := v.(string)
		bucketImageFileName := d.Get(PIImageBucketFile).(string)
		bucketRegion := d.Get(PIImageBucketRegion).(string)
		bucketAccess := d.Get(PIImageBucketAccess).(string)

		body := &models.CreateCosImageImportJob{
			ImageName:     &imageName,
			BucketName:    &bucketName,
			BucketAccess:  &bucketAccess,
			ImageFilename: &bucketImageFileName,
			Region:        &bucketRegion,
		}

		if v, ok := d.GetOk(PIImageAccessKey); ok {
			body.AccessKey = v.(string)
		}
		if v, ok := d.GetOk(PIImageSecretKey); ok {
			body.SecretKey = v.(string)
		}

		if v, ok := d.GetOk(ImageStorageType); ok {
			body.StorageType = v.(string)
		}
		if v, ok := d.GetOk(ImageStoragePool); ok {
			body.StoragePool = v.(string)
		}
		if ap, ok := d.GetOk(PIImageAffinityPolicy); ok {
			policy := ap.(string)
			affinity := &models.StorageAffinity{
				AffinityPolicy: &policy,
			}

			if policy == "affinity" {
				if av, ok := d.GetOk(PIImageAffinityVolume); ok {
					afvol := av.(string)
					affinity.AffinityVolume = &afvol
				}
				if ai, ok := d.GetOk(PIImageAffinityInstance); ok {
					afins := ai.(string)
					affinity.AffinityPVMInstance = &afins
				}
			} else {
				if avs, ok := d.GetOk(PIImageAntiAffinityVolumes); ok {
					afvols := flex.ExpandStringList(avs.([]interface{}))
					affinity.AntiAffinityVolumes = afvols
				}
				if ais, ok := d.GetOk(PIImageAntiAffinityInstances); ok {
					afinss := flex.ExpandStringList(ais.([]interface{}))
					affinity.AntiAffinityPVMInstances = afinss
				}
			}
			body.StorageAffinity = affinity
		}
		imageResponse, err := client.CreateCosImage(body)
		if err != nil {
			return diag.FromErr(err)
		}

		jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)
		_, err = waitForIBMPIJobCompleted(ctx, jobClient, *imageResponse.ID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}

		// Once the job is completed find by name
		image, err := client.Get(imageName)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *image.ImageID))
	}

	return resourceIBMPIImageRead(ctx, d, meta)
}

func resourceIBMPIImageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, imageID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	imageC := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	imagedata, err := imageC.Get(imageID)
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
	d.Set(ImageID, imageid)
	d.Set(PICloudInstanceID, cloudInstanceID)

	return nil
}

func resourceIBMPIImageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, imageID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	imageC := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)
	err = imageC.Delete(imageID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForIBMPIImageAvailable(ctx context.Context, client *st.IBMPIImageClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Power Image (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", "queued"},
		Target:     []string{"active"},
		Refresh:    isIBMPIImageRefreshFunc(ctx, client, id),
		Timeout:    timeout,
		Delay:      20 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIImageRefreshFunc(ctx context.Context, client *st.IBMPIImageClient, id string) resource.StateRefreshFunc {

	log.Printf("Calling the isIBMPIImageRefreshFunc Refresh Function....")
	return func() (interface{}, string, error) {
		image, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if image.State == "active" {
			return image, "active", nil
		}

		return image, "queued", nil
	}
}

func waitForIBMPIJobCompleted(ctx context.Context, client *st.IBMPIJobClient, jobID string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"queued", "readyForProcessing", "inProgress", "running", "waiting"},
		Target:  []string{"completed", "failed"},
		Refresh: func() (interface{}, string, error) {
			job, err := client.Get(jobID)
			if err != nil {
				log.Printf("[DEBUG] get job failed %v", err)
				return nil, "", fmt.Errorf(errors.GetJobOperationFailed, jobID, err)
			}
			if job == nil || job.Status == nil {
				log.Printf("[DEBUG] get job failed with empty response")
				return nil, "", fmt.Errorf("failed to get job status for job id %s", jobID)
			}
			if *job.Status.State == "failed" {
				log.Printf("[DEBUG] job status failed with message: %v", job.Status.Message)
				return nil, "failed", fmt.Errorf("job status failed for job id %s with message: %v", jobID, job.Status.Message)
			}
			return job, *job.Status.State, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}
