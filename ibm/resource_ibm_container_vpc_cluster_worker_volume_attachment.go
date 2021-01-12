package ibm

import (
	"fmt"
	"strings"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	volumeAttaching = "attaching"
	volumeAttached  = "attached"
)

func resourceIBMContainerVpcWorkerVolumeAttachment() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerVpcWorkerVolumeAttachmentCreate,
		Read:     resourceIBMContainerVpcWorkerVolumeAttachmentRead,
		Delete:   resourceIBMContainerVpcWorkerVolumeAttachmentDelete,
		Exists:   resourceIBMContainerVpcWorkerVolumeAttachmentExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"volume": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC Volume ID",
			},

			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster name or ID",
			},

			"worker": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "worker node ID",
			},

			"volume_attachment_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume attachment name",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume attachment status",
			},

			"volume_attachment_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume attachment ID",
			},

			"volume_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of volume",
			},
		},
	}
}

func resourceIBMContainerVpcWorkerVolumeAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	workersAPI := wpClient.Workers()

	volumeID := d.Get("volume").(string)
	clusterNameorID := d.Get("cluster").(string)
	workerID := d.Get("worker").(string)

	target, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}
	attachVolumeRequest := v2.VolumeRequest{
		Cluster:  clusterNameorID,
		VolumeID: volumeID,
		Worker:   workerID,
	}

	volumeattached, err := workersAPI.CreateStorageAttachment(attachVolumeRequest, target)
	if err != nil {
		fmt.Println(err)
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", clusterNameorID, workerID, volumeattached.Id))
	_, attachErr := waitforVolumetoAttach(d, meta)
	if attachErr != nil {
		return attachErr
	}
	return resourceIBMContainerVpcWorkerVolumeAttachmentRead(d, meta)
}

func resourceIBMContainerVpcWorkerVolumeAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	workersAPI := wpClient.Workers()
	target, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameorID := parts[0]
	workerID := parts[1]
	volumeAttachmentID := parts[2]

	volume, err := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, target)
	if err != nil {
		return err
	}
	d.Set("cluster", clusterNameorID)
	d.Set("worker", workerID)
	d.Set("volume_attachment_id", volumeAttachmentID)
	d.Set("volume_attachment_name", volume.Name)
	d.Set("status", volume.Status)
	d.Set("volume_type", volume.Type)
	return nil
}

func resourceIBMContainerVpcWorkerVolumeAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	workersAPI := wpClient.Workers()
	target, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameorID := parts[0]
	workerID := parts[1]
	volumeAttachmentID := parts[2]

	detachVolumeRequest := v2.VolumeRequest{
		Cluster:            clusterNameorID,
		VolumeAttachmentID: volumeAttachmentID,
		Worker:             workerID,
	}

	response, deleteErr := workersAPI.DeleteStorageAttachment(detachVolumeRequest, target)
	if deleteErr != nil && !strings.Contains(deleteErr.Error(), "EmptyResponseBody") {
		if response != "Ok" && strings.Contains(response, "Not found") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Failed to delete the volume attachment: %s", deleteErr)
	}
	return nil
}

func resourceIBMContainerVpcWorkerVolumeAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}

	workersAPI := wpClient.Workers()
	target, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return false, err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	clusterNameorID := parts[0]
	workerID := parts[1]
	volumeAttachmentID := parts[2]

	_, getErr := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, target)
	if getErr != nil {
		if apiErr, ok := getErr.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", getErr)
	}
	return true, nil
}

func waitforVolumetoAttach(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}

	workersAPI := wpClient.Workers()

	target, trgetErr := getVpcClusterTargetHeader(d, meta)
	if trgetErr != nil {
		return nil, trgetErr
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return nil, err
	}
	clusterNameorID := parts[0]
	workerID := parts[1]
	volumeAttachmentID := parts[2]

	createStateConf := &resource.StateChangeConf{
		Pending: []string{volumeAttaching},
		Target:  []string{volumeAttached},
		Refresh: func() (interface{}, string, error) {
			volume, err := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, target)

			if err != nil {
				return volume, volumeAttaching, err
			}

			if volume.Status == volumeAttached {
				return volume, volumeAttached, nil
			}
			return volume, volumeAttaching, nil

		},
		Timeout:                   d.Timeout(schema.TimeoutCreate),
		Delay:                     10 * time.Second,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
	}
	return createStateConf.WaitForState()
}
