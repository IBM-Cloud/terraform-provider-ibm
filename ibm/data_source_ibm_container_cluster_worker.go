package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMContainerClusterWorker() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerClusterWorkerRead,

		Schema: map[string]*schema.Schema{
			"worker_id": {
				Description: "ID of the worker",
				Type:        schema.TypeString,
				Required:    true,
			},
			"state": {
				Description: "State of the worker",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Status of the worker",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_vlan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_vlan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cluster region",
			},
		},
	}
}

func dataSourceIBMContainerClusterWorkerRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	wrkAPI := csClient.Workers()
	workerID := d.Get("worker_id").(string)
	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	workerFields, err := wrkAPI.Get(workerID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving worker: %s", err)
	}

	d.SetId(workerFields.ID)
	d.Set("state", workerFields.State)
	d.Set("status", workerFields.Status)
	d.Set("private_vlan", workerFields.PrivateVlan)
	d.Set("public_vlan", workerFields.PublicVlan)
	d.Set("private_ip", workerFields.PrivateIP)
	d.Set("public_ip", workerFields.PublicIP)

	return nil
}
