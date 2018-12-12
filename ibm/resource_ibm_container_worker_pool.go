package ibm

import (
	"fmt"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerWorkerPool() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerWorkerPoolCreate,
		Read:     resourceIBMContainerWorkerPoolRead,
		Update:   resourceIBMContainerWorkerPoolUpdate,
		Delete:   resourceIBMContainerWorkerPoolDelete,
		Exists:   resourceIBMContainerWorkerPoolExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(90 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"machine_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			"worker_pool_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"size_per_zone": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateSizePerZone,
			},

			"hardware": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      hardwareShared,
				ValidateFunc: validateAllowedStringValue([]string{hardwareShared, hardwareDedicated}),
			},

			"disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"private_vlan": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"public_vlan": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"worker_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
				Elem:             schema.TypeString,
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The worker pool region",
			},

			"resource_group_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "ID of the resource group.",
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},
		},
	}
}

func resourceIBMContainerWorkerPoolCreate(d *schema.ResourceData, meta interface{}) error {

	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	clusterNameorID := d.Get("cluster").(string)

	workerPoolConfig := v1.WorkerPoolConfig{
		Name:        d.Get("worker_pool_name").(string),
		Size:        d.Get("size_per_zone").(int),
		MachineType: d.Get("machine_type").(string),
	}
	if v, ok := d.GetOk("hardware"); ok {
		hardware := v.(string)
		switch strings.ToLower(hardware) {
		case "": // do nothing
		case hardwareDedicated:
			hardware = isolationPrivate
		case hardwareShared:
			hardware = isolationPublic
		}
		workerPoolConfig.Isolation = hardware
	}
	if l, ok := d.GetOk("labels"); ok {
		labels := make(map[string]string)
		for k, v := range l.(map[string]interface{}) {
			labels[k] = v.(string)
		}
		workerPoolConfig.Labels = labels
	}
	params := v1.WorkerPoolRequest{
		WorkerPoolConfig: workerPoolConfig,
		DiskEncryption:   d.Get("disk_encryption").(bool),
	}

	workerPoolsAPI := csClient.WorkerPools()
	targetEnv, err := getWorkerPoolTargetHeader(d, meta)
	if err != nil {
		return err
	}

	res, err := workerPoolsAPI.CreateWorkerPool(clusterNameorID, params, targetEnv)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", clusterNameorID, res.ID))

	return resourceIBMContainerWorkerPoolRead(d, meta)
}

func resourceIBMContainerWorkerPoolRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := csClient.WorkerPools()
	targetEnv, err := getWorkerPoolTargetHeader(d, meta)
	if err != nil {
		return err
	}

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID, targetEnv)
	if err != nil {
		return err
	}

	machineType := workerPool.MachineType
	d.Set("worker_pool_name", workerPool.Name)
	d.Set("machine_type", machineType)
	d.Set("size_per_zone", workerPool.Size)
	hardware := workerPool.Isolation
	switch strings.ToLower(hardware) {
	case "":
		hardware = hardwareShared
	case isolationPrivate:
		hardware = hardwareDedicated
	case isolationPublic:
		hardware = hardwareShared
	}
	d.Set("hardware", hardware)
	d.Set("state", workerPool.State)
	d.Set("labels", workerPool.Labels)
	d.Set("zones", flattenZones(workerPool.Zones))
	d.Set("cluster", cluster)
	d.Set("region", workerPool.Region)
	if strings.Contains(machineType, "encrypted") {
		d.Set("disk_encryption", true)
	} else {
		d.Set("disk_encryption", false)
	}
	return nil
}

func resourceIBMContainerWorkerPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameorID := parts[0]
	workerPoolNameorID := parts[1]
	workerPoolsAPI := csClient.WorkerPools()
	targetEnv, err := getWorkerPoolTargetHeader(d, meta)
	if err != nil {
		return err
	}

	if d.HasChange("size_per_zone") {
		err = workerPoolsAPI.ResizeWorkerPool(clusterNameorID, workerPoolNameorID, d.Get("size_per_zone").(int), targetEnv)
		if err != nil {
			return err
		}

		_, err = WaitForWorkerNormal(clusterNameorID, workerPoolNameorID, meta, d.Timeout(schema.TimeoutUpdate), targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for workers of worker pool (%s) of cluster (%s) to become ready: %s", workerPoolNameorID, clusterNameorID, err)
		}
	}

	return resourceIBMContainerWorkerPoolRead(d, meta)
}

func resourceIBMContainerWorkerPoolDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameorID := parts[0]
	workerPoolNameorID := parts[1]

	workerPoolsAPI := csClient.WorkerPools()
	targetEnv, err := getWorkerPoolTargetHeader(d, meta)
	if err != nil {
		return err
	}

	err = workerPoolsAPI.DeleteWorkerPool(clusterNameorID, workerPoolNameorID, targetEnv)
	if err != nil {
		return err
	}
	_, err = WaitForWorkerDelete(clusterNameorID, workerPoolNameorID, meta, d.Timeout(schema.TimeoutUpdate), targetEnv)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for removing workers of worker pool (%s) of cluster (%s): %s", workerPoolNameorID, clusterNameorID, err)
	}
	return nil
}

func resourceIBMContainerWorkerPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := csClient.WorkerPools()
	targetEnv, err := getWorkerPoolTargetHeader(d, meta)
	if err != nil {
		return false, err
	}

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return workerPool.ID == workerPoolID, nil
}

func WaitForWorkerNormal(clusterNameOrID, workerPoolNameOrID string, meta interface{}, timeout time.Duration, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", workerProvisioning},
		Target:     []string{workerNormal},
		Refresh:    workerPoolStateRefreshFunc(csClient.Workers(), clusterNameOrID, workerPoolNameOrID, target),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerPoolStateRefreshFunc(client v1.Workers, instanceID, workerPoolNameOrID string, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, workerPoolNameOrID, false, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		//Done worker has two fields State and Status , so check for those 2
		for _, e := range workerFields {
			if strings.Contains(e.KubeVersion, "pending") || strings.Compare(e.State, workerNormal) != 0 || strings.Compare(e.Status, workerReadyState) != 0 {
				if strings.Compare(e.State, "deleted") != 0 {
					return workerFields, workerProvisioning, nil
				}
			}
		}
		return workerFields, workerNormal, nil
	}
}

func WaitForWorkerDelete(clusterNameOrID, workerPoolNameOrID string, meta interface{}, timeout time.Duration, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{workerDeleteState},
		Refresh:    workerPoolDeleteStateRefreshFunc(csClient.Workers(), clusterNameOrID, workerPoolNameOrID, target),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerPoolDeleteStateRefreshFunc(client v1.Workers, instanceID, workerPoolNameOrID string, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, "", true, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		//Done worker has two fields State and Status , so check for those 2
		for _, e := range workerFields {
			if e.PoolName == workerPoolNameOrID || e.PoolID == workerPoolNameOrID {
				if strings.Compare(e.State, "deleted") != 0 {
					return workerFields, "deleting", nil
				}
			}
		}
		return workerFields, workerDeleteState, nil
	}
}

func getWorkerPoolTargetHeader(d *schema.ResourceData, meta interface{}) (v1.ClusterTargetHeader, error) {
	region := d.Get("region").(string)
	resourceGroup := d.Get("resource_group_id").(string)

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return v1.ClusterTargetHeader{}, err
	}

	if region == "" {
		region = sess.Config.Region
	}
	if resourceGroup == "" {
		resourceGroup = sess.Config.ResourceGroup

		if resourceGroup == "" {
			rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
			if err != nil {
				return v1.ClusterTargetHeader{}, err
			}
			resourceGroupQuery := management.ResourceGroupQuery{
				Default: true,
			}
			grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
			if err != nil {
				return v1.ClusterTargetHeader{}, err
			}
			resourceGroup = grpList[0].ID
		}
	}

	targetEnv := v1.ClusterTargetHeader{
		Region:        region,
		ResourceGroup: resourceGroup,
	}
	return targetEnv, nil
}
