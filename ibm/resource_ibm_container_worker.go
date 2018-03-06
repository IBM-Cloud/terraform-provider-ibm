package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	v1 "github.com/IBM-Bluemix/bluemix-go/api/container/containerv1"
	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerWorker() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerWorkerCreate,
		Read:     resourceIBMContainerWorkerRead,
		Update:   resourceIBMContainerWorkerUpdate,
		Delete:   resourceIBMContainerWorkerDelete,
		Exists:   resourceIBMContainerWorkerExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster name",
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "add",
			},
			"machine_type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "u2c.2x4",
			},
			"isolation": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"worker_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"billing": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "hourly",
			},

			"public_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  nil,
			},

			"private_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  nil,
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"wait_time_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  90,
			},
		},
	}
}

func resourceIBMContainerWorkerCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	clusterName := d.Get("cluster_name").(string)
	machineType := d.Get("machine_type").(string)
	workerNum := d.Get("worker_num").(int)
	publicVlanID := d.Get("public_vlan_id").(string)
	privateVlanID := d.Get("private_vlan_id").(string)
	isolation := d.Get("isolation").(string)

	params := v1.WorkerParam{
		MachineType: machineType,
		PrivateVlan: privateVlanID,
		PublicVlan:  publicVlanID,
		Isolation:   isolation,
		WorkerNum:   workerNum,
	}

	targetEnv := getClusterTargetHeader(d)

	err = csClient.Workers().Add(clusterName, params, targetEnv)
	if err != nil {
		return err
	}

	//wait for worker  availability
	_, err = WaitForWorkerAvailable(d, meta, targetEnv)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
	}

	return resourceIBMContainerWorkerRead(d, meta)
}

func resourceIBMContainerWorkerRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	targetEnv := getClusterTargetHeader(d)

	workerID := d.Id()
	_, err = csClient.Workers().Get(workerID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving worker: %s", err)
	}
	return nil
}

func resourceIBMContainerWorkerUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	targetEnv := getClusterTargetHeader(d)

	wrkAPI := csClient.Workers()
	clusterName := d.Get("cluster_name").(string)
	workerID := d.Id()

	if d.Get("action").(string) == "update" {
		params := v1.WorkerUpdateParam{
			Action: "update",
		}
		err := wrkAPI.Update(clusterName, workerID, params, targetEnv)
		if err != nil {
			return err
		}
		_, err = WaitForWorkerVersionUpdate(d, meta, targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for worker (%s) version to be updated: %s", d.Id(), err)
		}
	}
	return resourceIBMContainerWorkerRead(d, meta)
}

func resourceIBMContainerWorkerDelete(d *schema.ResourceData, meta interface{}) error {
	targetEnv := getClusterTargetHeader(d)
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterName := d.Get("cluster_name").(string)
	workerID := d.Id()
	err = csClient.Workers().Delete(clusterName, workerID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error worker: %s", err)
	}
	return nil
}

// WaitForWorkerVersionUpdate Waits for worker creation
func WaitForWorkerVersionUpdate(d *schema.ResourceData, meta interface{}, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for worker of the cluster (%s) to be available.", d.Id())
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", workerProvisioning},
		Target:     []string{workerNormal},
		Refresh:    workerVersionRefreshFunc(csClient.Workers(), id, d, target),
		Timeout:    time.Duration(d.Get("wait_time_minutes").(int)) * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerVersionRefreshFunc(client v1.Workers, instanceID string, d *schema.ResourceData, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.List(instanceID, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		log.Println("Checking workers...")
		//Done worker has two fields State and Status , so check for those 2
		for _, e := range workerFields {
			if strings.Contains(e.KubeVersion, "pending") {
				return workerFields, versionUpdating, nil
			}
		}
		return workerFields, workerNormal, nil
	}
}

func resourceIBMContainerWorkerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	targetEnv := getClusterTargetHeader(d)
	if err != nil {
		return false, err
	}
	workerID := d.Id()
	worker, err := csClient.Workers().Get(workerID, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return worker.ID == workerID, nil
}
