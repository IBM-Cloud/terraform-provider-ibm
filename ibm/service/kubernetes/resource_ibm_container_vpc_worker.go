// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

// Mutex to make resource creation sequential.
var resourceIBMContainerVpcWorkerCreateMutex sync.Mutex
var commonVarMutex sync.Mutex

// Status of worker replace
var workerReplaceStatus bool = false
var replaceInProgress bool = false

// Variable to identify the first run
var initRun int = 1

func ResourceIBMContainerVpcWorker() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerVpcWorkerCreate,
		Read:     resourceIBMContainerVpcWorkerRead,
		Delete:   resourceIBMContainerVpcWorkerDelete,
		Exists:   resourceIBMContainerVpcWorkerExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster name",
			},

			"replace_worker": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Worker name/id that needs to be replaced",
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "ID of the resource group.",
			},

			"kube_config_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Path of downloaded cluster config",
			},

			"check_ptx_status": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Check portworx status after worker replace",
			},

			"ptx_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "15m",
				Description: "Timeout for checking ptx pods/status",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					var err error
					_, err = time.ParseDuration(value)
					if err != nil {
						errors = append(errors, fmt.Errorf("[ERROR] Error parsing ptx_timeout: %s", err))
					}
					return
				},
			},

			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "IP of the replaced worker",
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource Controller URL",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMContainerVPCWorkerValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	containerVPCWorkerValidator := validate.ResourceValidator{ResourceName: "ibm_container_vpc_worker", Schema: validateSchema}
	return &containerVPCWorkerValidator
}

// Since Worker is being managed by Worker Pool, we can't create new workers
// but we can update/replace the existing workers to the new workers.
func resourceIBMContainerVpcWorkerCreate(d *schema.ResourceData, meta interface{}) error {

	//Current Resource status
	currentStatus := false

	workerID := d.Get("replace_worker").(string)

	defer func() {
		commonVarMutex.Lock()
		if currentStatus {
			workerReplaceStatus = true
		}
		replaceInProgress = false
		commonVarMutex.Unlock()
	}()

	//Continue only if the previous resource status is success
	err := waitForPreviousResource(workerID)
	if err != nil {
		return err
	}
	defer resourceIBMContainerVpcWorkerCreateMutex.Unlock()

	wkClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	clusterNameorID := d.Get("name").(string)
	cluster_config := d.Get("kube_config_path").(string)
	check_ptx_status := d.Get("check_ptx_status").(bool)

	worker, err := wkClient.Workers().Get(clusterNameorID, workerID, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting container vpc worker node: %s", err)
	}

	cls, err := wkClient.Clusters().GetCluster(clusterNameorID, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving conatiner vpc cluster: %s", err)
	}

	// Update the worker nodes after master node kube-version is updated.
	// workers will store the existing workers info to identify the replaced node
	workersInfo := make(map[string]int)

	workers, err := wkClient.Workers().ListWorkers(cls.ID, false, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving workers for cluster: %s", err)
	}

	for index, _worker := range workers {
		workersInfo[_worker.ID] = index
	}
	workersCount := len(workers)

	_, err = wkClient.Workers().ReplaceWokerNode(cls.ID, worker.ID, targetEnv)
	// As API returns http response 204 NO CONTENT, error raised will be exempted.
	if err != nil && !strings.Contains(err.Error(), "EmptyResponseBody") {
		return fmt.Errorf("[ERROR] Error replacing the worker node from the cluster: %s", err)
	}

	//1. wait for worker node to delete
	_, deleteError := waitForWorkerNodetoDelete(d, meta, targetEnv, worker.ID)
	if deleteError != nil {
		return fmt.Errorf("[ERROR] Worker node - %s is failed to replace", worker.ID)
	}

	//2. wait for new workerNode
	_, newWorkerError := waitForNewWorker(d, meta, targetEnv, workersCount)
	if newWorkerError != nil {
		return fmt.Errorf("[ERROR] Failed to spawn new worker node")
	}

	//3. Get new worker node ID and update the map
	newWorkerID, _, newNodeError := getNewWorkerID(d, meta, targetEnv, workersInfo)
	if newNodeError != nil {
		return fmt.Errorf("[ERROR] Unable to find the new worker node info")
	}

	d.SetId(newWorkerID)

	//4. wait for the worker's version update and normal state
	_, Err := WaitForVpcClusterWokersVersionUpdate(d, meta, targetEnv, cls.MasterKubeVersion, newWorkerID)
	if Err != nil {
		return fmt.Errorf(
			"[ERROR] Error waiting for cluster (%s) worker nodes kube version to be updated: %s", d.Id(), Err)
	}

	err = resourceIBMContainerVpcWorkerRead(d, meta)
	if err != nil {
		return err
	}

	if check_ptx_status {
		err = checkPortworxStatus(d, cluster_config)
		if err != nil {
			return err
		}
	}

	//Worker reloaded successfully
	currentStatus = true
	return nil
}

func resourceIBMContainerVpcWorkerRead(d *schema.ResourceData, meta interface{}) error {
	wkClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	workerID := d.Id()
	parts, err := flex.SepIdParts(workerID, "-")
	if err != nil {
		return err
	}
	if len(parts) < 2 {
		return fmt.Errorf("[ERROR] Incorrect ID %s: Id should be in kube-clusterID-* format", d.Id())
	}
	cluster := parts[1]

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	worker, err := wkClient.Workers().Get(cluster, workerID, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting container vpc worker node: %s", err)
	}
	for _, _network := range worker.NetworkInterfaces {
		if _network.Primary {
			d.Set("ip", _network.IpAddress)
			break
		}
	}
	d.Set(flex.ResourceStatus, worker.Health.State)

	cls, err := wkClient.Clusters().GetCluster(cluster, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving container vpc cluster: %s", err)
	}
	d.Set("name", cls.Name)
	d.Set("resource_group_id", cls.ResourceGroupID)

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, controller+"/kubernetes/clusters")
	d.Set(flex.ResourceGroupName, cls.ResourceGroupName)
	return nil
}

func resourceIBMContainerVpcWorkerDelete(d *schema.ResourceData, meta interface{}) error {
	_, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMContainerVpcWorkerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wkClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}

	workerID := d.Id()
	parts, err := flex.SepIdParts(workerID, "-")
	if err != nil {
		return false, err
	}
	if len(parts) < 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: Id should be in kube-clusterID-* format", d.Id())
	}
	cluster := parts[1]

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return false, err
	}

	worker, err := wkClient.Workers().Get(cluster, workerID, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 && strings.Contains(apiErr.Description(), "The specified worker node could not be found") {
				return false, nil
			}
		}
		return false, fmt.Errorf("[ERROR] Error getting container vpc worker node: %s", err)
	}
	if strings.Compare(worker.LifeCycle.ActualState, "deleted") == 0 {
		return false, nil
	}

	return worker.ID == workerID, nil
}

func waitForPreviousResource(worker_id string) error {
	time.Sleep(time.Second * 5)
	for {
		commonVarMutex.Lock()
		if !replaceInProgress {
			defer commonVarMutex.Unlock()
			if initRun == 1 || workerReplaceStatus {
				initRun = 0
				replaceInProgress = true
				log.Printf("Worker routine %s is taking mutex", worker_id)
				resourceIBMContainerVpcWorkerCreateMutex.Lock()
				return nil
			} else {
				return fmt.Errorf("[ERROR] Previous worker replace failed")
			}
		}
		commonVarMutex.Unlock()
		time.Sleep(time.Second * 10)
	}
}

func checkPortworxStatus(d *schema.ResourceData, cluster_config string) error {
	//1. Get worker ip
	worker_ip := d.Get("ip").(string)
	//1. Load the cluster config
	config, err := clientcmd.BuildConfigFromFlags("", cluster_config)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to set context: %s", err)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create clientset: %s", err)
	}

	//2. Retrieve portworx pod of current worker
	pod_name, err := WaitForPortworxPod(d, clientset, worker_ip)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to retrieve portworx pods: %s", err)
	}

	//3. Fetch portworx status json
	ptx_content, err := WaitForPortworxStatus(d, clientset, config, pod_name.(string))
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to fetch portworx status: %s", err)
	}

	//Execution should reach here only if the content is json & able to decode it without errors
	//4. Check portworx status on current worker
	node_data := (ptx_content.(map[string]interface{})["cluster"]).(map[string]interface{})["Nodes"].([]interface{})
	for _, n := range node_data {
		n_data := n.(map[string]interface{})
		if n_data["MgmtIp"] == worker_ip {
			ptx_status := n_data["NodeData"].(map[string]interface{})["STORAGE-INFO"].(map[string]interface{})["Status"].(string)
			if ptx_status == "Up" {
				//portworx is Up on this node
				log.Printf("Portworx Status is Up")
				return nil
			} else {
				return fmt.Errorf("[ERROR] Portworx Status is not Up on this node: %s", ptx_status)
			}
		}
	}
	return fmt.Errorf("[ERROR] No pods found with label name=portworx in kube-system namespace")
}

func WaitForPortworxPod(d *schema.ResourceData, clientset *kubernetes.Clientset, worker_ip string) (interface{}, error) {
	log.Printf("Waiting for the portworx pod to be Available & Ready")

	ptx_timeout, err := time.ParseDuration(d.Get("ptx_timeout").(string))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error parsing ptx_timeout: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"ptx_not_ready"},
		Target:       []string{"ptx_ready"},
		Refresh:      ptxPodRefreshFunc(clientset, worker_ip),
		Timeout:      time.Duration(ptx_timeout),
		Delay:        30 * time.Second,
		PollInterval: 30 * time.Second,
	}

	return stateConf.WaitForState()
}

func ptxPodRefreshFunc(clientset *kubernetes.Clientset, worker_ip string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		//1. List pods from kube-system namespace
		podList, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{LabelSelector: "name=portworx"})
		if err != nil {
			log.Printf("Failed to list pods: %s", err)
			return nil, "ptx_not_ready", nil
		}

		//2. Get pod from the current worker
		for _, pod := range podList.Items {
			if pod.Status.HostIP == worker_ip {
				for _, container := range pod.Status.ContainerStatuses {
					if container.Name == "portworx" && container.Ready {
						log.Printf("Portworx pod %s is ready", pod.Name)
						return pod.Name, "ptx_ready", nil
					}
				}
				log.Printf("Portworx pod %s not ready yet", pod.Name)
				return pod.Name, "ptx_not_ready", nil
			}
		}
		return nil, "ptx_not_ready", nil
	}
}

func WaitForPortworxStatus(d *schema.ResourceData, clientset *kubernetes.Clientset, config *rest.Config, pod_name string) (interface{}, error) {
	log.Printf("Waiting to fetch portworx status json")

	ptx_timeout, err := time.ParseDuration(d.Get("ptx_timeout").(string))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error parsing ptx_timeout: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"pxctl_fail"},
		Target:       []string{"pxctl_success"},
		Refresh:      ptxStatusRefreshFunc(clientset, config, pod_name),
		Timeout:      time.Duration(ptx_timeout),
		Delay:        30 * time.Second,
		PollInterval: 30 * time.Second,
	}

	return stateConf.WaitForState()
}

func ptxStatusRefreshFunc(clientset *kubernetes.Clientset, config *rest.Config, pod_name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		//Byte buffers for the exec command output
		var stdout, stderr bytes.Buffer

		request := clientset.CoreV1().RESTClient().Post().
			Resource("pods").
			Name(pod_name).
			Namespace("kube-system").
			SubResource("exec")

		//portworx command to fetch the status as json file
		ptx_cmd := []string{
			"/opt/pwx/bin/pxctl",
			"status",
			"-j",
		}
		request.VersionedParams(&v1.PodExecOptions{
			Command: ptx_cmd,
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
		}, scheme.ParameterCodec)

		exec, err := remotecommand.NewSPDYExecutor(config, "POST", request.URL())
		if err != nil {
			log.Printf("[ERROR] Failed to Execute Remote Command: %s", err)
			return nil, "pxctl_fail", nil
		}
		err = exec.Stream(remotecommand.StreamOptions{
			Stdin:  os.Stdin,
			Stdout: &stdout,
			Stderr: &stderr,
			Tty:    false,
		})
		if err != nil {
			log.Printf("[ERROR] Failed to Read Remote Stream: %s", err)
			return nil, "pxctl_fail", nil
		}

		//If any error occurs in the exec command, log error & retry
		if len(stderr.String()) != 0 {
			log.Printf("[ERROR] Execute Remote Command Error: %s", stderr.String())
			return nil, "pxctl_fail", nil
		}

		//Parse json data to check the portworx status
		var parse_content map[string]interface{}
		err = json.Unmarshal(stdout.Bytes(), &parse_content)
		if err != nil {
			log.Printf("[ERROR] Failed to decode ptx status json: %s", err)
			return nil, "pxctl_fail", nil
		}

		log.Printf("Successfully fetched portworx status json")
		return parse_content, "pxctl_success", nil
	}
}
