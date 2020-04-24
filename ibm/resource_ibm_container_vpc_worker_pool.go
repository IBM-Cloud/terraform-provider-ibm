package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	workerDesired = "deployed"
)

func resourceIBMContainerVpcWorkerPool() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerVpcWorkerPoolCreate,
		Update:   resourceIBMContainerVpcWorkerPoolUpdate,
		Read:     resourceIBMContainerVpcWorkerPoolRead,
		Delete:   resourceIBMContainerVpcWorkerPoolDelete,
		Exists:   resourceIBMContainerVpcWorkerPoolExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(90 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"flavor": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"worker_pool_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"zones": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: applyOnce,
				Elem:             schema.TypeString,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group.",
				ForceNew:    true,
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The vpc id where the cluster is",
				ForceNew:    true,
			},
			"worker_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of workers",
				ForceNew:    true,
			},
		},
	}
}

func resourceIBMContainerVpcWorkerPoolCreate(d *schema.ResourceData, meta interface{}) error {

	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	clusterNameorID := d.Get("cluster").(string)
	var zonei []interface{}

	zone := []v2.Zone{}

	if res, ok := d.GetOk("zones"); ok {
		zonei = res.([]interface{})
		for _, e := range zonei {
			r, _ := e.(map[string]interface{})
			zoneParam := v2.Zone{
				ID:       r["name"].(string),
				SubnetID: r["subnet_id"].(string),
			}
			zone = append(zone, zoneParam)
		}

	}

	// for _, e := range d.Get("zones").(*schema.Set).List() {
	// 	value := e.(map[string]interface{})
	// 	id := value["id"].(string)
	// 	subnetid := value["subnet_id"].(string)

	// }

	workerPoolConfig := v2.WorkerPoolConfig{
		Name:        d.Get("worker_pool_name").(string),
		VpcID:       d.Get("vpc_id").(string),
		Flavor:      d.Get("flavor").(string),
		WorkerCount: d.Get("worker_count").(int),
		Zones:       zone,
	}

	if l, ok := d.GetOk("labels"); ok {
		labels := make(map[string]string)
		for k, v := range l.(map[string]interface{}) {
			labels[k] = v.(string)
		}
		workerPoolConfig.Labels = labels
	}
	params := v2.WorkerPoolRequest{
		WorkerPoolConfig: workerPoolConfig,
		Cluster:          clusterNameorID,
	}

	workerPoolsAPI := wpClient.WorkerPools()
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	res, err := workerPoolsAPI.CreateWorkerPool(params, targetEnv)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", clusterNameorID, res.ID))

	//wait for workerpool availability
	_, err = WaitForWorkerPoolAvailable(d, meta, clusterNameorID, res.ID, d.Timeout(schema.TimeoutCreate), targetEnv)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for workerpool (%s) to become ready: %s", d.Id(), err)
	}

	return resourceIBMContainerVpcWorkerPoolUpdate(d, meta)
}

func resourceIBMContainerVpcWorkerPoolUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("worker_count") {
		clusterNameOrID := d.Get("cluster").(string)
		workerPoolName := d.Get("worker_pool_name").(string)
		count := d.Get("worker_count").(int)
		targetEnv, err := getVpcClusterTargetHeader(d, meta)
		if err != nil {
			return err
		}
		ClusterClient, err := meta.(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		Env := v1.ClusterTargetHeader{ResourceGroup: targetEnv.ResourceGroup}

		err = ClusterClient.WorkerPools().ResizeWorkerPool(clusterNameOrID, workerPoolName, count, Env)
		if err != nil {
			return fmt.Errorf(
				"Error updating the worker_count %d: %s", count, err)
		}
	}
	return resourceIBMContainerVpcWorkerPoolRead(d, meta)
}

func resourceIBMContainerVpcWorkerPoolRead(d *schema.ResourceData, meta interface{}) error {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := wpClient.WorkerPools()
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID, targetEnv)
	if err != nil {
		return err
	}

	var zones = make([]map[string]interface{}, 0)
	for _, zone := range workerPool.Zones {
		for _, subnet := range zone.Subnets {
			zoneInfo := map[string]interface{}{
				"name":      zone.ID,
				"subnet_id": subnet.ID,
			}
			zones = append(zones, zoneInfo)
		}
	}

	d.Set("worker_pool_name", workerPool.PoolName)
	d.Set("flavor", workerPool.Flavor)
	d.Set("worker_count", workerPool.WorkerCount)
	d.Set("provider", workerPool.Provider)
	d.Set("labels", workerPool.Labels)
	d.Set("zones", zones)
	d.Set("cluster", cluster)
	d.Set("vpc_id", workerPool.VpcID)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/kubernetes/clusters")
	return nil
}

func resourceIBMContainerVpcWorkerPoolDelete(d *schema.ResourceData, meta interface{}) error {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameorID := parts[0]
	workerPoolNameorID := parts[1]

	workerPoolsAPI := wpClient.WorkerPools()
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	err = workerPoolsAPI.DeleteWorkerPool(clusterNameorID, workerPoolNameorID, targetEnv)
	if err != nil {
		return err
	}
	_, err = WaitForVpcWorkerDelete(clusterNameorID, workerPoolNameorID, meta, d.Timeout(schema.TimeoutDelete), targetEnv)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for removing workers of worker pool (%s) of cluster (%s): %s", workerPoolNameorID, clusterNameorID, err)
	}
	d.SetId("")
	return nil
}

func resourceIBMContainerVpcWorkerPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := wpClient.WorkerPools()
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
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

// WaitForWorkerPoolAvailable Waits for worker creation
func WaitForWorkerPoolAvailable(d *schema.ResourceData, meta interface{}, clusterNameOrID, workerPoolNameOrID string, timeout time.Duration, target v2.ClusterTargetHeader) (interface{}, error) {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for workerpool (%s) to be available.", d.Id())
	// id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"provision_pending"},
		Target:     []string{workerDesired},
		Refresh:    vpcWorkerPoolStateRefreshFunc(wpClient.Workers(), clusterNameOrID, workerPoolNameOrID, target),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func vpcWorkerPoolStateRefreshFunc(client v2.Workers, instanceID string, workerPoolNameOrID string, target v2.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, "", false, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		// Check active transactions
		//Check for worker state to be deployed
		//Done worker has two fields desiredState and actualState , so check for those 2
		for _, e := range workerFields {
			if e.PoolName == workerPoolNameOrID || e.PoolID == workerPoolNameOrID {
				if strings.Compare(e.LifeCycle.ActualState, "deployed") != 0 {
					log.Printf("worker: %s state: %s", e.ID, e.LifeCycle.ActualState)
					return workerFields, "provision_pending", nil
				}
			}
		}
		return workerFields, workerDesired, nil
	}
}

func WaitForVpcWorkerDelete(clusterNameOrID, workerPoolNameOrID string, meta interface{}, timeout time.Duration, target v2.ClusterTargetHeader) (interface{}, error) {
	wpClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{workerDeleteState},
		Refresh:    vpcworkerPoolDeleteStateRefreshFunc(wpClient.Workers(), clusterNameOrID, workerPoolNameOrID, target),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func vpcworkerPoolDeleteStateRefreshFunc(client v2.Workers, instanceID, workerPoolNameOrID string, target v2.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, "", true, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		//Done worker has two fields desiredState and actualState , so check for those 2
		for _, e := range workerFields {
			if e.PoolName == workerPoolNameOrID || e.PoolID == workerPoolNameOrID {
				if strings.Compare(e.LifeCycle.ActualState, "deleted") != 0 {
					log.Printf("Deleting worker %s", e.ID)
					return workerFields, "deleting", nil
				}
			}
		}
		return workerFields, workerDeleteState, nil
	}
}
