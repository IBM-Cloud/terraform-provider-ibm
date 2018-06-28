package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	clusterNormal     = "normal"
	workerNormal      = "normal"
	subnetNormal      = "normal"
	workerReadyState  = "Ready"
	workerDeleteState = "deleted"

	versionUpdating     = "updating"
	clusterProvisioning = "provisioning"
	workerProvisioning  = "provisioning"
	subnetProvisioning  = "provisioning"

	hardwareShared    = "shared"
	hardwareDedicated = "dedicated"
	isolationPublic   = "public"
	isolationPrivate  = "private"

	defaultWorkerPool = "default"
)

const PUBLIC_SUBNET_TYPE = "public"

func resourceIBMContainerCluster() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerClusterCreate,
		Read:     resourceIBMContainerClusterRead,
		Update:   resourceIBMContainerClusterUpdate,
		Delete:   resourceIBMContainerClusterDelete,
		Exists:   resourceIBMContainerClusterExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster name",
			},
			"datacenter": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The datacenter where this cluster will be deployed",
			},
			"workers": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"worker_num"},
				Deprecated:    "Use worker_num instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "add",
							ValidateFunc: validateAllowedStringValue([]string{"add", "reboot", "reload"}),
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"worker_num": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				Description:   "Number of worker nodes",
				ConflictsWith: []string{"workers"},
				ValidateFunc:  validateWorkerNum,
			},

			"workers_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IDs of the worker node",
			},

			"disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},

			"kube_version": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"machine_type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"isolation": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{"hardware"},
				Deprecated:    "Use hardware instead",
			},
			"hardware": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{"isolation"},
				Default:       hardwareShared,
				ValidateFunc:  validateAllowedStringValue([]string{hardwareShared, hardwareDedicated}),
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
			"ingress_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"no_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"is_trusted": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"server_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"subnet_id": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"webhook": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"slack"}),
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"worker_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_per_zone": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hardware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kube_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeMap,
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
					},
				},
			},
		},
	}
}

func resourceIBMContainerClusterCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	datacenter := d.Get("datacenter").(string)
	billing := d.Get("billing").(string)
	machineType := d.Get("machine_type").(string)
	publicVlanID := d.Get("public_vlan_id").(string)
	privateVlanID := d.Get("private_vlan_id").(string)
	webhooks := d.Get("webhook").([]interface{})
	noSubnet := d.Get("no_subnet").(bool)
	enableTrusted := d.Get("is_trusted").(bool)
	diskEncryption := d.Get("disk_encryption").(bool)
	var workers []interface{}
	var workerNum int
	if v, ok := d.GetOk("workers"); ok {
		workers = v.([]interface{})
		workerNum = len(workers)
	}

	if v, ok := d.GetOk("worker_num"); ok {
		workerNum = v.(int)
	}

	if workerNum == 0 {
		return fmt.Errorf(
			"Please set either the wokers with valid array or worker_num with value grater than 0")
	}

	//Read the hardware and convert it to appropriate
	var isolation string

	hardware := d.Get("hardware").(string)
	switch strings.ToLower(hardware) {
	case "": // do nothing
	case hardwareDedicated:
		isolation = isolationPrivate
	case hardwareShared:
		isolation = isolationPublic
	}

	if v, ok := d.GetOk("isolation"); ok {
		isolation = v.(string)
	}

	params := v1.ClusterCreateRequest{
		Name:           name,
		Datacenter:     datacenter,
		WorkerNum:      workerNum,
		Billing:        billing,
		MachineType:    machineType,
		PublicVlan:     publicVlanID,
		PrivateVlan:    privateVlanID,
		NoSubnet:       noSubnet,
		Isolation:      isolation,
		DiskEncryption: diskEncryption,
		EnableTrusted:  enableTrusted,
	}

	if v, ok := d.GetOk("kube_version"); ok {
		params.MasterVersion = v.(string)
	}

	targetEnv := getClusterTargetHeader(d)

	cls, err := csClient.Clusters().Create(params, targetEnv)
	if err != nil {
		return err
	}
	d.SetId(cls.ID)
	//wait for cluster availability
	_, err = WaitForClusterAvailable(d, meta, targetEnv)
	//wait for worker  availability
	_, err = WaitForWorkerAvailable(d, meta, targetEnv)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
	}

	subnetAPI := csClient.Subnets()
	subnetIDs := d.Get("subnet_id").(*schema.Set)
	var publicSubnetAdded bool
	if noSubnet == false {
		publicSubnetAdded = true
	}
	var subnets []v1.Subnet
	if len(subnetIDs.List()) > 0 {
		subnets, err = subnetAPI.List(targetEnv)
		if err != nil {
			return err
		}
	}
	for _, subnetID := range subnetIDs.List() {
		if subnetID != "" {
			err = subnetAPI.AddSubnet(cls.ID, subnetID.(string), targetEnv)
			if err != nil {
				return err
			}
			subnet := getSubnet(subnets, subnetID.(string))
			if subnet.Type == PUBLIC_SUBNET_TYPE {
				publicSubnetAdded = true
			}
		} else {
			return fmt.Errorf(
				"subnet_id can not contain empty value")
		}
	}

	if publicSubnetAdded {
		_, err = WaitForSubnetAvailable(d, meta, targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for initializing ingress hostname and secret: %s", err)
		}
	}
	whkAPI := csClient.WebHooks()

	for _, e := range webhooks {
		pack := e.(map[string]interface{})
		webhook := v1.WebHook{
			Level: pack["level"].(string),
			Type:  pack["type"].(string),
			URL:   pack["url"].(string),
		}

		whkAPI.Add(cls.ID, webhook, targetEnv)

	}

	workersInfo := []map[string]string{}
	wrkAPI := csClient.Workers()
	workerFields, err := wrkAPI.List(cls.ID, targetEnv)
	if err != nil {
		return err
	}
	//Create a map with worker name and id
	for i, e := range workers {
		pack := e.(map[string]interface{})
		var worker = map[string]string{
			"name":    pack["name"].(string),
			"id":      workerFields[i].ID,
			"action":  pack["action"].(string),
			"version": strings.Split(workerFields[i].KubeVersion, "_")[0],
		}
		workersInfo = append(workersInfo, worker)
	}
	d.Set("workers", workersInfo)

	if err != nil {
		return fmt.Errorf(
			"Error waiting for cluster (%s) to become ready: %s", d.Id(), err)
	}

	return resourceIBMContainerClusterRead(d, meta)
}

func resourceIBMContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	wrkAPI := csClient.Workers()
	workerPoolsAPI := csClient.WorkerPools()

	targetEnv := getClusterTargetHeader(d)

	clusterID := d.Id()
	cls, err := csClient.Clusters().Find(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving armada cluster: %s", err)
	}

	workerFields, err := wrkAPI.List(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving workers for cluster: %s", err)
	}
	workers := make([]string, len(workerFields))
	for i, worker := range workerFields {
		workers[i] = worker.ID
	}

	workersByPool, err := wrkAPI.ListByWorkerPool(clusterID, defaultWorkerPool, false)
	if err != nil {
		return fmt.Errorf("Error retrieving workers for cluster: %s", err)
	}

	hardware := workersByPool[0].Isolation
	switch strings.ToLower(hardware) {
	case "":
		hardware = hardwareShared
	case isolationPrivate:
		hardware = hardwareDedicated
	case isolationPublic:
		hardware = hardwareShared
	}

	workerPools, err := workerPoolsAPI.ListWorkerPools(clusterID)
	if err != nil {
		return err
	}

	defaultWorkerPool, err := workerPoolsAPI.GetWorkerPool(clusterID, "default")
	if err != nil {
		return err
	}
	zones := defaultWorkerPool.Zones
	for _, zone := range zones {
		if zone.ID == cls.DataCenter {
			d.Set("worker_num", zone.WorkerCount)
			break
		}
	}

	d.Set("name", cls.Name)
	d.Set("server_url", cls.ServerURL)
	d.Set("ingress_hostname", cls.IngressHostname)
	d.Set("ingress_secret", cls.IngressSecretName)

	d.Set("subnet_id", d.Get("subnet_id").(*schema.Set))
	d.Set("workers_info", workers)
	d.Set("kube_version", strings.Split(cls.MasterKubeVersion, "_")[0])
	d.Set("is_trusted", cls.IsTrusted)
	d.Set("worker_pools", flattenWorkerPools(workerPools))
	d.Set("hardware", hardware)

	return nil
}

func resourceIBMContainerClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	targetEnv := getClusterTargetHeader(d)

	subnetAPI := csClient.Subnets()
	whkAPI := csClient.WebHooks()
	wrkAPI := csClient.Workers()
	clusterAPI := csClient.Clusters()

	clusterID := d.Id()

	if d.HasChange("kube_version") {
		var masterVersion string
		if v, ok := d.GetOk("kube_version"); ok {
			masterVersion = v.(string)
		}
		params := v1.ClusterUpdateParam{
			Action:  "update",
			Force:   true,
			Version: masterVersion,
		}
		err := clusterAPI.Update(clusterID, params, targetEnv)
		if err != nil {
			return err
		}
		_, err = WaitForClusterVersionUpdate(d, meta, targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for cluster (%s) version to be updated: %s", d.Id(), err)
		}
	}

	workersInfo := []map[string]string{}
	if d.HasChange("worker_num") {
		workerPoolsAPI := csClient.WorkerPools()

		worker_num := d.Get("worker_num").(int)
		err = workerPoolsAPI.ResizeWorkerPool(clusterID, "default", worker_num)
		if err != nil {
			return fmt.Errorf(
				"Error updating the worker_num %d: %s", worker_num, err)
		}

		_, err = WaitForWorkerAvailable(d, meta, targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
		}
	}

	if d.HasChange("workers") {
		oldWorkers, newWorkers := d.GetChange("workers")
		oldWorker := oldWorkers.([]interface{})
		newWorker := newWorkers.([]interface{})
		for _, nW := range newWorker {
			newPack := nW.(map[string]interface{})
			exists := false
			for _, oW := range oldWorker {
				oldPack := oW.(map[string]interface{})
				if strings.Compare(newPack["name"].(string), oldPack["name"].(string)) == 0 {
					exists = true
					if strings.Compare(newPack["action"].(string), oldPack["action"].(string)) != 0 {
						params := v1.WorkerUpdateParam{
							Action: newPack["action"].(string),
						}
						err := wrkAPI.Update(clusterID, oldPack["id"].(string), params, targetEnv)
						if err != nil {
							return fmt.Errorf("Error updating worker %s: %s", oldPack["id"].(string), err)
						}
						_, err = WaitForWorkerAvailable(d, meta, targetEnv)
						if err != nil {
							return fmt.Errorf(
								"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
						}
						workerObj, err := wrkAPI.Get(oldPack["id"].(string), targetEnv)
						if err != nil {
							return fmt.Errorf("Error getting worker %s: %s", oldPack["id"].(string), err)
						}
						var worker = map[string]string{
							"name":    newPack["name"].(string),
							"id":      newPack["id"].(string),
							"action":  newPack["action"].(string),
							"version": strings.Split(workerObj.KubeVersion, "_")[0],
						}
						workersInfo = append(workersInfo, worker)
					} else if strings.Compare(newPack["version"].(string), oldPack["version"].(string)) != 0 {
						cluster, err := clusterAPI.Find(clusterID, targetEnv)
						if err != nil {
							return fmt.Errorf("Error retrieving cluster %s: %s", clusterID, err)
						}
						if newPack["version"].(string) != strings.Split(cluster.MasterKubeVersion, "_")[0] {
							return fmt.Errorf("Worker version %s should match the master kube version %s", newPack["version"].(string), strings.Split(cluster.MasterKubeVersion, "_")[0])
						}
						params := v1.WorkerUpdateParam{
							Action: "update",
						}
						err = wrkAPI.Update(clusterID, oldPack["id"].(string), params, targetEnv)
						if err != nil {
							return fmt.Errorf("Error updating worker %s: %s", oldPack["id"].(string), err)
						}
						_, err = WaitForWorkerAvailable(d, meta, targetEnv)
						if err != nil {
							return fmt.Errorf(
								"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
						}
						workerObj, err := wrkAPI.Get(oldPack["id"].(string), targetEnv)
						if err != nil {
							return fmt.Errorf("Error getting worker %s: %s", oldPack["id"].(string), err)
						}
						var worker = map[string]string{
							"name":    newPack["name"].(string),
							"id":      newPack["id"].(string),
							"action":  newPack["action"].(string),
							"version": strings.Split(workerObj.KubeVersion, "_")[0],
						}
						workersInfo = append(workersInfo, worker)

					} else {
						workerObj, err := wrkAPI.Get(oldPack["id"].(string), targetEnv)
						if err != nil {
							return fmt.Errorf("Error getting worker %s: %s", oldPack["id"].(string), err)
						}
						var worker = map[string]string{
							"name":    oldPack["name"].(string),
							"id":      oldPack["id"].(string),
							"action":  oldPack["action"].(string),
							"version": strings.Split(workerObj.KubeVersion, "_")[0],
						}
						workersInfo = append(workersInfo, worker)
					}
				}
			}
			if !exists {
				params := v1.WorkerParam{
					Action: "add",
					Count:  1,
				}
				err := wrkAPI.Add(clusterID, params, targetEnv)
				if err != nil {
					return fmt.Errorf("Error adding worker to cluster")
				}
				id, err := getID(d, meta, clusterID, oldWorker, workersInfo)
				if err != nil {
					return fmt.Errorf("Error getting id of worker")
				}
				_, err = WaitForWorkerAvailable(d, meta, targetEnv)
				if err != nil {
					return fmt.Errorf(
						"Error waiting for workers of cluster (%s) to become ready: %s", d.Id(), err)
				}
				workerObj, err := wrkAPI.Get(id, targetEnv)
				if err != nil {
					return fmt.Errorf("Error getting worker %s: %s", id, err)
				}
				var worker = map[string]string{
					"name":    newPack["name"].(string),
					"id":      id,
					"action":  newPack["action"].(string),
					"version": strings.Split(workerObj.KubeVersion, "_")[0],
				}
				workersInfo = append(workersInfo, worker)
			}
		}
		for _, oW := range oldWorker {
			oldPack := oW.(map[string]interface{})
			exists := false
			for _, nW := range newWorker {
				newPack := nW.(map[string]interface{})
				exists = exists || (strings.Compare(oldPack["name"].(string), newPack["name"].(string)) == 0)
			}
			if !exists {
				wrkAPI.Delete(clusterID, oldPack["id"].(string), targetEnv)
			}

		}
		//wait for new workers to available
		//Done - Can we not put WaitForWorkerAvailable after all client.DeleteWorker
		d.Set("workers", workersInfo)
	}

	//TODO put webhooks can't deleted in the error message if such case is observed in the chnages
	if d.HasChange("webhook") {
		oldHooks, newHooks := d.GetChange("webhook")
		oldHook := oldHooks.([]interface{})
		newHook := newHooks.([]interface{})
		for _, nH := range newHook {
			newPack := nH.(map[string]interface{})
			exists := false
			for _, oH := range oldHook {
				oldPack := oH.(map[string]interface{})
				if (strings.Compare(newPack["level"].(string), oldPack["level"].(string)) == 0) && (strings.Compare(newPack["type"].(string), oldPack["type"].(string)) == 0) && (strings.Compare(newPack["url"].(string), oldPack["url"].(string)) == 0) {
					exists = true
				}
			}
			if !exists {
				webhook := v1.WebHook{
					Level: newPack["level"].(string),
					Type:  newPack["type"].(string),
					URL:   newPack["url"].(string),
				}

				whkAPI.Add(clusterID, webhook, targetEnv)
			}
		}
	}
	//TODO put subnet can't deleted in the error message if such case is observed in the chnages
	if d.HasChange("subnet_id") {
		oldSubnets, newSubnets := d.GetChange("subnet_id")
		oldSubnet := oldSubnets.(*schema.Set)
		newSubnet := newSubnets.(*schema.Set)
		rem := oldSubnet.Difference(newSubnet).List()
		if len(rem) > 0 {
			return fmt.Errorf("Subnet(s) %v cannot be deleted", rem)
		}
		var publicSubnetAdded bool
		subnets, err := subnetAPI.List(targetEnv)
		if err != nil {
			return err
		}
		for _, nS := range newSubnet.List() {
			exists := false
			for _, oS := range oldSubnet.List() {
				if strings.Compare(nS.(string), oS.(string)) == 0 {
					exists = true
				}
			}
			if !exists {
				err := subnetAPI.AddSubnet(clusterID, nS.(string), targetEnv)
				if err != nil {
					return err
				}
				subnet := getSubnet(subnets, nS.(string))
				if subnet.Type == PUBLIC_SUBNET_TYPE {
					publicSubnetAdded = true
				}
			}
		}
		if publicSubnetAdded {
			_, err = WaitForSubnetAvailable(d, meta, targetEnv)
			if err != nil {
				return fmt.Errorf(
					"Error waiting for initializing ingress hostname and secret: %s", err)
			}
		}
	}
	return resourceIBMContainerClusterRead(d, meta)
}

func getID(d *schema.ResourceData, meta interface{}, clusterID string, oldWorkers []interface{}, workerInfo []map[string]string) (string, error) {
	targetEnv := getClusterTargetHeader(d)
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return "", err
	}
	workerFields, err := csClient.Workers().List(clusterID, targetEnv)
	if err != nil {
		return "", err
	}
	for _, wF := range workerFields {
		exists := false
		for _, oW := range oldWorkers {
			oldPack := oW.(map[string]interface{})
			if strings.Compare(wF.ID, oldPack["id"].(string)) == 0 || strings.Compare(wF.State, "deleted") == 0 {
				exists = true
			}
		}
		if !exists {
			for i := 0; i < len(workerInfo); i++ {
				pack := workerInfo[i]
				exists = exists || (strings.Compare(wF.ID, pack["id"]) == 0)
			}
			if !exists {
				return wF.ID, nil
			}
		}
	}

	return "", fmt.Errorf("Unable to get ID of worker")
}

func resourceIBMContainerClusterDelete(d *schema.ResourceData, meta interface{}) error {
	targetEnv := getClusterTargetHeader(d)
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterID := d.Id()
	err = csClient.Clusters().Delete(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error deleting cluster: %s", err)
	}
	return nil
}

// WaitForClusterAvailable Waits for cluster creation
func WaitForClusterAvailable(d *schema.ResourceData, meta interface{}, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for cluster (%s) to be available.", d.Id())
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", clusterProvisioning},
		Target:     []string{clusterNormal},
		Refresh:    clusterStateRefreshFunc(csClient.Clusters(), id, d, target),
		Timeout:    time.Duration(d.Get("wait_time_minutes").(int)) * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func clusterStateRefreshFunc(client v1.Clusters, instanceID string, d *schema.ResourceData, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		clusterFields, err := client.Find(instanceID, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving cluster: %s", err)
		}
		// Check active transactions
		log.Println("Checking cluster")
		//Check for cluster state to be normal
		log.Println("Checking cluster state", strings.Compare(clusterFields.State, clusterNormal))
		if strings.Compare(clusterFields.State, clusterNormal) != 0 {
			return clusterFields, clusterProvisioning, nil
		}
		return clusterFields, clusterNormal, nil
	}
}

// WaitForWorkerAvailable Waits for worker creation
func WaitForWorkerAvailable(d *schema.ResourceData, meta interface{}, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for worker of the cluster (%s) to be available.", d.Id())
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", workerProvisioning},
		Target:     []string{workerNormal},
		Refresh:    workerStateRefreshFunc(csClient.Workers(), id, d, target),
		Timeout:    time.Duration(d.Get("wait_time_minutes").(int)) * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerStateRefreshFunc(client v1.Workers, instanceID string, d *schema.ResourceData, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.List(instanceID, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		log.Println("Checking workers...")
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

func WaitForSubnetAvailable(d *schema.ResourceData, meta interface{}, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for Ingress Subdomain and secret being assigned.")
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", workerProvisioning},
		Target:     []string{workerNormal},
		Refresh:    subnetStateRefreshFunc(csClient.Clusters(), id, d, target),
		Timeout:    time.Duration(d.Get("wait_time_minutes").(int)) * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func subnetStateRefreshFunc(client v1.Clusters, instanceID string, d *schema.ResourceData, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cluster, err := client.Find(instanceID, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving cluster: %s", err)
		}
		if cluster.IngressHostname == "" || cluster.IngressSecretName == "" {
			return cluster, subnetProvisioning, nil
		}
		return cluster, subnetNormal, nil
	}
}

// WaitForClusterVersionUpdate Waits for cluster creation
func WaitForClusterVersionUpdate(d *schema.ResourceData, meta interface{}, target v1.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for cluster (%s) version to be updated.", d.Id())
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", versionUpdating},
		Target:     []string{clusterNormal},
		Refresh:    clusterVersionRefreshFunc(csClient.Clusters(), id, d, target),
		Timeout:    time.Duration(d.Get("wait_time_minutes").(int)) * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func clusterVersionRefreshFunc(client v1.Clusters, instanceID string, d *schema.ResourceData, target v1.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		clusterFields, err := client.Find(instanceID, target)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving cluster: %s", err)
		}
		// Check active transactions
		log.Println("Checking cluster version", clusterFields.MasterKubeVersion, d.Get("kube_version").(string))
		if strings.Contains(clusterFields.MasterKubeVersion, "pending") {
			return clusterFields, versionUpdating, nil
		}
		return clusterFields, clusterNormal, nil
	}
}

func resourceIBMContainerClusterExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	targetEnv := getClusterTargetHeader(d)
	if err != nil {
		return false, err
	}
	clusterID := d.Id()
	cls, err := csClient.Clusters().Find(clusterID, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return cls.ID == clusterID, nil
}

func getSubnet(subnets []v1.Subnet, subnetId string) v1.Subnet {
	for _, subnet := range subnets {
		if subnet.ID == subnetId {
			return subnet
		}
	}
	return v1.Subnet{}
}
