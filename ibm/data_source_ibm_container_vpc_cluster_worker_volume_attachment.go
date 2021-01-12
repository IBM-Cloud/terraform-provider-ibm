package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMContainerVpcWorkerVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerVpcWorkerVolumeAttachmentRead,

		Schema: map[string]*schema.Schema{
			"volume_attachment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vpc volume attachment ID",
			},

			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster name or ID",
			},

			"worker": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "worker node ID",
			},

			"volume": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC Volume ID",
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
			"volume_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of volume",
			},
		},
	}
}

func dataSourceIBMContainerVpcWorkerVolumeAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	workersAPI := wpClient.Workers()
	target, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	clusterNameorID := d.Get("cluster").(string)
	volumeAttachmentID := d.Get("volume_attachment_id").(string)
	workerID := d.Get("worker").(string)

	volume, err := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, target)
	if err != nil {
		return err
	}
	d.Set("volume_attachment_name", volume.Name)
	d.Set("status", volume.Status)
	d.Set("volume_type", volume.Type)
	d.SetId(fmt.Sprintf("%s/%s/%s", clusterNameorID, workerID, volumeAttachmentID))
	return nil
}
