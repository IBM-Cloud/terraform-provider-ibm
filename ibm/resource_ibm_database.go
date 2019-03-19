package ibm

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/flatmap"

	//	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	validation "github.com/hashicorp/terraform/helper/validation"
)

const (
	databaseInstanceSuccessStatus  = "active"
	databaseInstanceProgressStatus = "in progress"
	databaseInstanceInactiveStatus = "inactive"
	databaseInstanceFailStatus     = "failed"
	databaseInstanceRemovedStatus  = "removed"
)

const (
	databaseTaskSuccessStatus  = "completed"
	databaseTaskProgressStatus = "running"
	databaseTaskFailStatus     = "failed"
)

type CsEntry struct {
	Name       string
	Password   string
	String     string
	Composed   string
	CertName   string
	CertBase64 string
	Hosts      []struct {
		HostName string `json:"hostname"`
		Port     int    `json:"port"`
	}
	Scheme       string
	QueryOptions map[string]interface{}
	Path         string
	Database     string
}

func resourceIBMDatabaseInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMDatabaseInstanceCreate,
		Read:     resourceIBMDatabaseInstanceRead,
		Update:   resourceIBMDatabaseInstanceUpdate,
		Delete:   resourceIBMDatabaseInstanceDelete,
		Exists:   resourceIBMDatabaseInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Resource instance name for example, my Database instance",
				Type:        schema.TypeString,
				Required:    true,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the resource group in which the Database instance is present",
			},

			"location": {
				Description: "The location or the region in which Database instance exists",
				Type:        schema.TypeString,
				Required:    true,
			},

			"service": {
				Description:  "The name of the Cloud Internet database service",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"databases-for-etcd", "databases-for-postgresql", "databases-for-redis", "databases-for-elasticsearch", "databases-for-mongodb", "messages-for-rabbitmq", "databases-for-mysql"}),
			},
			"plan": {
				Description:  "The plan type of the Database instance",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"standard"}),
			},

			"status": {
				Description: "The resource instance status",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"adminuser": {
				Description: "The admin user id for the instance",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"adminpassword": {
				Description:  "The admin user password for the instance",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(10, 32),
				Sensitive:    true,
				// DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				//  return true
				// },
			},
			"version": {
				Description: "The database version to provision if specified",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"members_memory_allocation_mb": {
				Description:  "Memory allocation required for cluster",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(2048, 114688),
			},
			"members_disk_allocation_mb": {
				Description:  "Disk allocation required for cluster",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(2048, 1048576),
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"users": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description:  "User name",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(5, 32),
						},
						"password": {
							Description:  "User password",
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							ValidateFunc: validation.StringLenBetween(10, 32),
						},
					},
				},
			},
			"connectionstrings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "User name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"composed": {
							Description: "Connection string",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"scheme": {
							Description: "DB scheme",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"certname": {
							Description: "Certificate Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"certbase64": {
							Description: "Certificate in base64 encoding",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": {
							Description: "Password",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"queryoptions": {
							Description: "DB query options",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"database": {
							Description: "DB name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"path": {
							Description: "DB path",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"hosts": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostname": {
										Description: "DB host name",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"port": {
										Description: "DB port",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"whitelist": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description:  "Whitelist IP address in CIDR notation",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.CIDRNetwork(24, 32),
						},
						"description": {
							Description:  "Unique white list description",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(1, 32),
						},
					},
				},
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "Scaling group name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"count": {
							Description: "Count of scaling groups for the instance",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"memory": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"units": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The units memory is allocated in.",
									},
									"allocation_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The current memory allocation for a group instance",
									},
									"minimum_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum memory size for a group instance",
									},
									"step_size_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The step size memory increases or decreases in.",
									},
									"is_adjustable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the memory size adjustable.",
									},
									"can_scale_down": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Can memory scale down as well as up.",
									},
								},
							},
						},
						"cpu": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"units": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The .",
									},
									"allocation_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The current cpu allocation count",
									},
									"minimum_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum number of cpus allowed",
									},
									"step_size_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of CPUs allowed to step up or down by",
									},
									"is_adjustable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Are the number of CPUs adjustable",
									},
									"can_scale_down": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Can the number of CPUs be scaled down as well as up",
									},
								},
							},
						},
						"disk": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"units": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The units disk is allocated in",
									},
									"allocation_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The current disk allocation",
									},
									"minimum_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum disk size allowed",
									},
									"step_size_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The step size disk increases or decreases in",
									},
									"is_adjustable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the disk size adjustable",
									},
									"can_scale_down": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Can the disk size be scaled down as well as up",
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

type Params struct {
	Version       string `json:"version,omitempty"`
	KeyProtectKey string `json:"key_protect_key,omitempty"`
	Memory        int    `json:"members_memory_allocation_mb,omitempty"`
	Disk          int    `json:"members_disk_allocation_mb,omitempty"`
}

// Replace with func wrapper for resourceIBMResourceInstanceCreate specifying serviceName := "database......."
func resourceIBMDatabaseInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := controller.CreateServiceInstanceRequest{
		Name: name,
	}

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return fmt.Errorf("Error retrieving database service offering: %s", err)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	rsInst.ServicePlanID = servicePlan

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return fmt.Errorf("Error retrieving deployment for plan %s : %s", plan, err)
	}
	if len(deployments) == 0 {
		return fmt.Errorf("No deployment found for service plan : %s", plan)
	}
	deployments, supportedLocations := filterDatabaseDeployments(deployments, location)

	if len(deployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		return fmt.Errorf("No deployment found for service plan %s at location %s.\nValid location(s) are: %q.", plan, location, locationList)
	}

	rsInst.TargetCrn = deployments[0].CatalogCRN

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rsInst.ResourceGroupID = rsGrpID.(string)
	} else {
		rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
		if err != nil {
			return err
		}
		resourceGroupQuery := management.ResourceGroupQuery{
			Default: true,
		}
		grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
		if err != nil {
			return err
		}
		rsInst.ResourceGroupID = grpList[0].ID
	}

	params := Params{}
	if memory, ok := d.GetOk("members_memory_allocation_mb"); ok {
		params.Memory = memory.(int)
	}
	if disk, ok := d.GetOk("members_disk_allocation_mb"); ok {
		params.Disk = disk.(int)
	}
	if version, ok := d.GetOk("version"); ok {
		params.Version = version.(string)
	}
	if keyProtect, ok := d.GetOk("key_protect_key"); ok {
		params.KeyProtectKey = keyProtect.(string)
	}
	parameters, _ := json.Marshal(params)
	var raw map[string]interface{}
	json.Unmarshal(parameters, &raw)
	//paramString := string(parameters[:])
	rsInst.Parameters = raw

	if _, ok := d.GetOk("tags"); ok {
		rsInst.Tags = getServiceTags(d)
	}

	instance, err := rsConClient.ResourceServiceInstance().CreateInstance(rsInst)
	if err != nil {
		return fmt.Errorf("Error creating database instance: %s", err)
	}

	// Moved d.SetId(instance.ID) to after waiting for resource to finish creation. Otherwise Terraform initates depedent tasks too early.
	// Original flow had SetId here as its required as input to waitForDatabaseInstanceCreate

	_, err = waitForDatabaseInstanceCreate(d, meta, instance.ID)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for create database instance (%s) to complete: %s", d.Id(), err)
	}

	icdId := EscapeUrlParm(instance.ID)
	icdClient, err := meta.(ClientSession).ICDAPI()
	if err != nil {
		return fmt.Errorf("Error getting database client settings: %s", err)
	}

	if pw, ok := d.GetOk("adminpassword"); ok {
		adminPassword := pw.(string)
		cdb, err := icdClient.Cdbs().GetCdb(icdId)
		if err != nil {
			if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
				return fmt.Errorf("The database instance was not found in the region set for the Provider, or the default of us-south. Specify the correct region in the provider definition, or create a provider alias for the correct region. %v", err)
			}
			return fmt.Errorf("Error getting database config for: %s with error %s\n", icdId, err)
		}

		userParams := icdv4.UserReq{
			User: icdv4.User{
				Password: adminPassword,
			},
		}
		task, err := icdClient.Users().UpdateUser(icdId, cdb.AdminUser, userParams)
		if err != nil {
			return fmt.Errorf("Error updating database admin password: %s", err)
		}
		_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for update of database (%s) admin password task to complete: %s", icdId, err)
		}
	}

	if wl, ok := d.GetOk("whitelist"); ok {
		whitelist := expandWhitelist(wl.(*schema.Set))
		for _, wlEntry := range whitelist {
			whitelistReq := icdv4.WhitelistReq{
				WhitelistEntry: icdv4.WhitelistEntry{
					Address:     wlEntry.Address,
					Description: wlEntry.Description,
				},
			}
			task, err := icdClient.Whitelists().CreateWhitelist(icdId, whitelistReq)
			if err != nil {
				return fmt.Errorf("Error updating database whitelist entry: %s", err)
			}
			_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
			if err != nil {
				return fmt.Errorf(
					"Error waiting for update of database (%s) whitelist task to complete: %s", icdId, err)
			}
		}
	}

	if userlist, ok := d.GetOk("users"); ok {
		users := expandUsers(userlist.(*schema.Set))
		for _, user := range users {
			userReq := icdv4.UserReq{
				User: icdv4.User{
					UserName: user.UserName,
					Password: user.Password,
				},
			}
			task, err := icdClient.Users().CreateUser(icdId, userReq)
			if err != nil {
				return fmt.Errorf("Error updating database user (%s) entry: %s", user.UserName, err)
			}
			_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
			if err != nil {
				return fmt.Errorf(
					"Error waiting for update of database (%s) user (%s) create task to complete: %s", icdId, user.UserName, err)
			}
		}
	}

	d.SetId(instance.ID)

	return resourceIBMDatabaseInstanceRead(d, meta)
}

func resourceIBMDatabaseInstanceRead(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Id()
	instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
	if err != nil {
		if strings.Contains(err.Error(), "Object not found") ||
			strings.Contains(err.Error(), "status code: 404") {
			log.Printf("[WARN] Removing record from state because it's not found via the API")
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving resource instance: %s", err)
	}
	if strings.Contains(instance.State, "removed") {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return nil
	}

	err = GetTags(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error on get of resource instance (%s) tags: %s", d.Id(), err)
	}
	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("parameters", flatmap.Flatten(instance.Parameters))
	d.Set("location", instance.RegionID)

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.GetServiceName(instance.ServiceID)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}

	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(instance.ServicePlanID)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	d.Set("plan", servicePlan)

	icdClient, err := meta.(ClientSession).ICDAPI()
	if err != nil {
		return fmt.Errorf("Error getting database client settings: %s", err)
	}

	icdId := EscapeUrlParm(instanceID)
	cdb, err := icdClient.Cdbs().GetCdb(icdId)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
			return fmt.Errorf("The database instance was not found in the region set for the Provider. Specify the correct region in the provider definition. %v", err)
		}
		return fmt.Errorf("Error getting database config for: %s with error %s\n", icdId, err)
	}
	d.Set("adminuser", cdb.AdminUser)
	d.Set("version", cdb.Version)

	groupList, err := icdClient.Groups().GetGroups(icdId)
	if err != nil {
		return fmt.Errorf("Error getting database groups: %s", err)
	}
	d.Set("groups", flattenIcdGroups(groupList))
	d.Set("members_memory_allocation_mb", groupList.Groups[0].Memory.AllocationMb)
	d.Set("members_disk_allocation_mb", groupList.Groups[0].Disk.AllocationMb)

	whitelist, err := icdClient.Whitelists().GetWhitelist(icdId)
	if err != nil {
		return fmt.Errorf("Error getting database whitelist: %s", err)
	}
	d.Set("whitelist", flattenWhitelist(whitelist))

	var connectionStrings []CsEntry
	//ICD does not implement a GetUsers API. Users populated from tf configuration.
	tfusers := d.Get("users").(*schema.Set)
	users := expandUsers(tfusers)
	user := icdv4.User{
		UserName: cdb.AdminUser,
	}
	users = append(users, user)
	for _, user := range users {
		userName := user.UserName
		csEntry, err := getConnectionString(d, userName, meta)
		if err != nil {
			return fmt.Errorf("Error getting user connection string for user (%s): %s", userName, err)
		}
		connectionStrings = append(connectionStrings, csEntry)
	}
	d.Set("connectionstrings", flattenConnectionStrings(connectionStrings))

	return nil
}

func resourceIBMDatabaseInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Id()
	updateReq := controller.UpdateServiceInstanceRequest{}
	if d.HasChange("name") {
		updateReq.Name = d.Get("name").(string)
	}

	_, err = rsConClient.ResourceServiceInstance().UpdateInstance(instanceID, updateReq)
	if err != nil {
		return fmt.Errorf("Error updating resource instance: %s", err)
	}

	_, err = waitForDatabaseInstanceUpdate(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for update of resource instance (%s) to complete: %s", d.Id(), err)
	}

	if d.HasChange("tags") {
		err = UpdateTags(d, meta)
		if err != nil {
			return fmt.Errorf(
				"Error on update of resource instance (%s) tags: %s", d.Id(), err)
		}
	}

	icdClient, err := meta.(ClientSession).ICDAPI()
	if err != nil {
		return fmt.Errorf("Error getting database client settings: %s", err)
	}
	icdId := EscapeUrlParm(instanceID)

	if d.HasChange("members_memory_allocation_mb") || d.HasChange("members_disk_allocation_mb") {
		params := icdv4.GroupReq{}
		if d.HasChange("members_memory_allocation_mb") {
			memory := d.Get("members_memory_allocation_mb").(int)
			memoryReq := icdv4.MemoryReq{AllocationMb: memory}
			params.GroupBdy.Memory = &memoryReq
		}
		if d.HasChange("members_disk_allocation_mb") {
			disk := d.Get("members_disk_allocation_mb").(int)
			diskReq := icdv4.DiskReq{AllocationMb: disk}
			params.GroupBdy.Disk = &diskReq
		}
		task, err := icdClient.Groups().UpdateGroup(icdId, "member", params)
		if err != nil {
			return fmt.Errorf("Error updating database scaling group: %s", err)
		}
		_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for database (%s) scaling group update task to complete: %s", icdId, err)
		}

	}

	if d.HasChange("adminpassword") {
		adminUser := d.Get("adminuser").(string)
		password := d.Get("adminpassword").(string)
		userParams := icdv4.UserReq{
			User: icdv4.User{
				Password: password,
			},
		}
		task, err := icdClient.Users().UpdateUser(icdId, adminUser, userParams)
		if err != nil {
			return fmt.Errorf("Error updating database admin password: %s", err)
		}
		_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for database (%s) admin password update task to complete: %s", icdId, err)
		}
	}

	if d.HasChange("whitelist") {
		oldList, newList := d.GetChange("whitelist")
		if oldList == nil {
			oldList = new(schema.Set)
		}
		if newList == nil {
			newList = new(schema.Set)
		}
		os := oldList.(*schema.Set)
		ns := newList.(*schema.Set)
		remove := os.Difference(ns).List()
		add := ns.Difference(os).List()

		if len(add) > 0 {
			for _, entry := range add {
				newEntry := entry.(map[string]interface{})
				wlEntry := icdv4.WhitelistEntry{
					Address:     newEntry["address"].(string),
					Description: newEntry["description"].(string),
				}
				whitelistReq := icdv4.WhitelistReq{
					WhitelistEntry: wlEntry,
				}
				task, err := icdClient.Whitelists().CreateWhitelist(icdId, whitelistReq)
				if err != nil {
					return fmt.Errorf("Error updating database whitelist entry %v : %s", wlEntry.Address, err)
				}
				_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
				if err != nil {
					return fmt.Errorf(
						"Error waiting for database (%s) whitelist create task to complete for entry %s : %s", icdId, wlEntry.Address, err)
				}

			}

		}

		if len(remove) > 0 {
			for _, entry := range remove {
				newEntry := entry.(map[string]interface{})
				wlEntry := icdv4.WhitelistEntry{
					Address:     newEntry["address"].(string),
					Description: newEntry["description"].(string),
				}
				ipAddress := wlEntry.Address
				task, err := icdClient.Whitelists().DeleteWhitelist(icdId, ipAddress)
				if err != nil {
					return fmt.Errorf("Error deleting database whitelist entry: %s", err)
				}
				_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
				if err != nil {
					return fmt.Errorf(
						"Error waiting for database (%s) whitelist delete task to complete for ipAddress %s : %s", icdId, ipAddress, err)
				}

			}
		}
	}

	if d.HasChange("users") {
		oldList, newList := d.GetChange("users")
		if oldList == nil {
			oldList = new(schema.Set)
		}
		if newList == nil {
			newList = new(schema.Set)
		}
		os := oldList.(*schema.Set)
		ns := newList.(*schema.Set)
		remove := os.Difference(ns).List()
		add := ns.Difference(os).List()

		if len(add) > 0 {
			for _, entry := range add {
				newEntry := entry.(map[string]interface{})
				userEntry := icdv4.User{
					UserName: newEntry["name"].(string),
					Password: newEntry["password"].(string),
				}
				userReq := icdv4.UserReq{
					User: userEntry,
				}
				task, err := icdClient.Users().CreateUser(icdId, userReq)
				if err != nil {
					// ICD does not report if error was due to user already being defined. Check if can
					// successfully update password by itself.
					userParams := icdv4.UserReq{
						User: icdv4.User{
							Password: newEntry["password"].(string),
						},
					}
					task, err := icdClient.Users().UpdateUser(icdId, newEntry["name"].(string), userParams)
					if err != nil {
						return fmt.Errorf("Error updating database user (%s) password: %s", newEntry["name"].(string), err)
					}
					_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
					if err != nil {
						return fmt.Errorf(
							"Error waiting for database (%s) user (%s) password update task to complete: %s", icdId, newEntry["name"].(string), err)
					}
				} else {
					_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
					if err != nil {
						return fmt.Errorf(
							"Error waiting for database (%s) user (%s) create task to complete: %s", icdId, newEntry["name"].(string), err)
					}
				}
			}

		}

		if len(remove) > 0 {
			for _, entry := range remove {
				newEntry := entry.(map[string]interface{})
				userEntry := icdv4.User{
					UserName: newEntry["name"].(string),
					Password: newEntry["password"].(string),
				}
				user := userEntry.UserName
				task, err := icdClient.Users().DeleteUser(icdId, user)
				if err != nil {
					return fmt.Errorf("Error deleting database user (%s) entry: %s", user, err)
				}
				_, err = waitForDatabaseTaskComplete(task.Id, d, meta)
				if err != nil {
					return fmt.Errorf(
						"Error waiting for database (%s) user (%s) delete task to complete: %s", icdId, user, err)
				}
			}
		}
	}

	return resourceIBMDatabaseInstanceRead(d, meta)
}

func getConnectionString(d *schema.ResourceData, userName string, meta interface{}) (CsEntry, error) {
	csEntry := CsEntry{}
	icdClient, err := meta.(ClientSession).ICDAPI()
	if err != nil {
		return csEntry, fmt.Errorf("Error getting database client settings: %s", err)
	}

	icdId := d.Id()
	connection, err := icdClient.Connections().GetConnection(icdId, userName)
	if err != nil {
		return csEntry, fmt.Errorf("Error getting database user connection string via ICD API: %s", err)
	}

	service := d.Get("service")
	dbConnection := icdv4.Uri{}
	switch service {
	case "databases-for-postgresql":
		dbConnection = connection.Postgres
	case "databases-for-redis":
		dbConnection = connection.Rediss
	case "databases-for-mongodb":
		dbConnection = connection.Mongo
	// case "databases-for-mysql":
	// 	dbConnection = connection.Mysql
	case "databases-for-elasticsearch":
		dbConnection = connection.Https
	case "databases-for-etcd":
		dbConnection = connection.Grpc
	case "messages-for-rabbitmq":
		dbConnection = connection.Amqps
	default:
		return csEntry, fmt.Errorf("Unrecognised database type during connection string lookup: %s", service)
	}

	csEntry = CsEntry{
		Name:     userName,
		Password: "",
		// Populate only first 'composed' connection string as an example
		Composed:     dbConnection.Composed[0],
		CertName:     dbConnection.Certificate.Name,
		CertBase64:   dbConnection.Certificate.CertificateBase64,
		Hosts:        dbConnection.Hosts,
		Scheme:       dbConnection.Scheme,
		Path:         dbConnection.Path,
		QueryOptions: dbConnection.QueryOptions.(map[string]interface{}),
	}
	// Postgres DB name is of type string, Redis is json.Number, others are nil
	if dbConnection.Database != nil {
		switch v := dbConnection.Database.(type) {
		default:
			return csEntry, fmt.Errorf("Unexpected data type: %T", v)
		case json.Number:
			csEntry.Database = dbConnection.Database.(json.Number).String()
		case string:
			csEntry.Database = dbConnection.Database.(string)
		}
	} else {
		csEntry.Database = ""
	}
	return csEntry, nil
}

func resourceIBMDatabaseInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	id := d.Id()

	err = rsConClient.ResourceServiceInstance().DeleteInstance(id, true)
	if err != nil {
		// If prior delete occurs, instance is not immediately deleted, but remains in "removed" state"
		// RC 410 with "Gone" returned as error
		if strings.Contains(err.Error(), "Gone") ||
			strings.Contains(err.Error(), "status code: 410") {
			log.Printf("[WARN] Resource instance already deleted %s\n ", err)
			err = nil
		} else {
			return fmt.Errorf("Error deleting resource instance: %s", err)
		}
	}

	_, err = waitForDatabaseInstanceDelete(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for resource instance (%s) to be deleted: %s", d.Id(), err)
	}

	d.SetId("")

	return nil
}
func resourceIBMDatabaseInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()
	instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	if strings.Contains(instance.State, "removed") {
		log.Printf("[WARN] Removing instance from state because it's in removed state")
		d.SetId("")
		return false, nil
	}

	return instance.ID == instanceID, nil
}

func waitForDatabaseInstanceCreate(d *schema.ResourceData, meta interface{}, instanceID string) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus},
		Target:  []string{databaseInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if instance.State == databaseInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func waitForDatabaseInstanceUpdate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus},
		Target:  []string{databaseInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if instance.State == databaseInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func waitForDatabaseTaskComplete(taskId string, d *schema.ResourceData, meta interface{}) (bool, error) {
	icdClient, err := meta.(ClientSession).ICDAPI()
	if err != nil {
		return false, fmt.Errorf("Error getting database client settings: %s", err)
	}
	delayDuration := 5 * time.Second

	timeout := time.After(15 * time.Minute)
	delay := time.Tick(delayDuration)
	innerTask := icdv4.Task{}

	for {
		select {
		case <-timeout:
			return false, fmt.Errorf("[Error] Time out waiting for database task to complete")
		case <-delay:
			innerTask, err = icdClient.Tasks().GetTask(EscapeUrlParm(taskId))
			if err != nil {
				return false, fmt.Errorf("The ICD Get task on database update errored: %v", err)
			}
			if innerTask.Status == "failed" {
				return false, fmt.Errorf("[Error] Database task failed")
			}
			// Completed status could be returned as "" due to interaction between bluemix-go and icd task response
			// Otherwise Running an queued
			if innerTask.Status == "completed" || innerTask.Status == "" {
				return true, nil
			}

		}
	}
}

func waitForDatabaseInstanceDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()
	stateConf := &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus, databaseInstanceSuccessStatus},
		Target:  []string{databaseInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, databaseInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if instance.State == databaseInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed to delete: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func filterDatabaseDeployments(deployments []models.ServiceDeployment, location string) ([]models.ServiceDeployment, map[string]bool) {
	supportedDeployments := []models.ServiceDeployment{}
	supportedLocations := make(map[string]bool)
	for _, d := range deployments {
		if d.Metadata.RCCompatible {
			deploymentLocation := d.Metadata.Deployment.Location
			supportedLocations[deploymentLocation] = true
			if deploymentLocation == location {
				supportedDeployments = append(supportedDeployments, d)
			}
		}
	}
	return supportedDeployments, supportedLocations
}
