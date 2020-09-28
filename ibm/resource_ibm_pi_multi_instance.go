package ibm

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func resourceIBMPIMultiInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIMultiInstanceCreate,
		Read:     resourceIBMPIMultiInstanceRead,
		Update:   resourceIBMPIMultiInstanceUpdate,
		Delete:   resourceIBMPIMultiInstanceDelete,
		Exists:   resourceIBMPIMultiInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(120 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Power Instance id that is assigned to the account",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI instance status",
			},
			"migratable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "set to true to enable migration of the PI instance",
			},
			"min_processors": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum number of the CPUs",
			},
			"min_memory": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum memory",
			},
			"max_processors": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum number of processors",
			},
			"max_memory": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum memory size",
			},
			helpers.PIInstanceNetworkIds: {
				Type:             schema.TypeList,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Description:      "Set of Networks that have been configured for the account",
				DiffSuppressFunc: applyOnce,
			},

			helpers.PIInstanceVolumeIds: {
				Type:             schema.TypeSet,
				Optional:         true,
				Computed:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "List of PI volumes",
			},

			helpers.PIInstanceUserData: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded data to be passed in for invoking a cloud init script",
			},

			"multivmdata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_virtual_cores": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"operating_system": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operating System",
						},
						"health_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PI Instance health status",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PI Instance health status",
						},
						"min_processors": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Minimum number of the CPUs",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance ID",
						},
						"addresses": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"macaddress": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						helpers.PIVirtualCoresAssigned: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Virtual Cores Assigned to the PVMInstance",
						},
						"os_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS Type",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS Type",
						},
						"max_processors": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Maximum number of processors",
						},
						"max_memory": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Maximum memory size",
						},
						"min_virtual_cores": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum Virtual Cores Assigned to the PVMInstance",
						},
					},
				},
			},

			"health_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI Instance health status",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance ID",
			},
			"pin_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PIN Policy of the Instance",
			},
			helpers.PIInstanceImageName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI instance image name",
			},
			helpers.PIInstanceProcessors: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Processors count",
			},
			helpers.PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI Instance name",
			},
			helpers.PIInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared", "capped"}),
				Description:  "Instance processor type",
			},
			helpers.PIInstanceSSHKeyName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SSH key name",
			},
			helpers.PIInstanceMemory: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Memory size",
			},
			helpers.PIInstanceSystemType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"s922", "e880", "e980"}),
				Description:  "PI Instance system type",
			},
			helpers.PIInstanceReplicants: {
				Type:        schema.TypeFloat,
				Optional:    true,
				Default:     "1",
				Description: "PI Instance replicas count",
			},
			helpers.PIInstanceReplicationPolicy: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"affinity", "anti-affinity", "none"}),
				Description:  "Replication policy for the PI Instance",
			},
			helpers.PIInstanceReplicationScheme: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"prefix", "suffix"}),
				Description:  "Replication scheme",
			},
			helpers.PIInstanceProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
			helpers.PIInstancePinPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Pin Policy of the instance",
				Default:      "none",
				ValidateFunc: validateAllowedStringValue([]string{"none", "soft", "hard"}),
			},

			"reboot_for_resource_change": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Flag to be passed for CPU/Memory changes that require a reboot to take effect",
			},
			"operating_system": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operating System",
			},
			"os_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS Type",
			},
			helpers.PIInstanceHealthStatus: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"OK", "WARNING"}),
				Default:      "OK",
				Description:  "Allow the user to set the status of the lpar so that they can connect to it faster",
			},
			helpers.PIVirtualCoresAssigned: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Virtual Cores Assigned to the PVMInstance",
			},
			"max_virtual_cores": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum Virtual Cores Assigned to the PVMInstance",
			},
			"min_virtual_cores": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Virtual Cores Assigned to the PVMInstance",
			},
		},
	}
}

func resourceIBMPIMultiInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Now in the Multi PVM Create")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIInstanceName).(string)
	sshkey := d.Get(helpers.PIInstanceSSHKeyName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	systype := d.Get(helpers.PIInstanceSystemType).(string)
	networks := expandStringList(d.Get(helpers.PIInstanceNetworkIds).([]interface{}))
	volids := expandStringList((d.Get(helpers.PIInstanceVolumeIds).(*schema.Set)).List())
	replicants := d.Get(helpers.PIInstanceReplicants).(float64)
	replicationpolicy := d.Get(helpers.PIInstanceReplicationPolicy).(string)
	replicationNamingScheme := d.Get(helpers.PIInstanceReplicationScheme).(string)
	imageid := d.Get(helpers.PIInstanceImageName).(string)
	processortype := d.Get(helpers.PIInstanceProcType).(string)
	pinpolicy := d.Get(helpers.PIInstancePinPolicy).(string)

	if d.Get(helpers.PIInstancePinPolicy) == "" {
		pinpolicy = "none"
	}

	instance_ready_status := d.Get(helpers.PIInstanceHealthStatus).(string)
	if d.Get(helpers.PIInstanceHealthStatus) == "" {
		log.Printf("Instance Ready Status is not provided. Setting the default to OK")
		instance_ready_status = "OK"
	}

	log.Printf("The accepted instance status for the LPAR [%s] can be [%s] ", name, instance_ready_status)

	//var userdata = ""
	user_data := d.Get(helpers.PIInstanceUserData).(string)
	if d.Get(helpers.PIInstanceUserData) == "" {
		user_data = ""
	}
	err = checkBase64(user_data)
	if err != nil {
		log.Printf("Data is not base64 encoded")
		return err
	}

	sort.Strings(networks)

	//publicinterface := d.Get(helpers.PIInstancePublicNetwork).(bool)
	body := &models.PVMInstanceCreate{
		//NetworkIds: networks,
		Processors:              &procs,
		Memory:                  &mem,
		ServerName:              ptrToString(name),
		SysType:                 systype,
		KeyPairName:             sshkey,
		ImageID:                 ptrToString(imageid),
		ProcType:                ptrToString(processortype),
		Replicants:              replicants,
		UserData:                user_data,
		ReplicantNamingScheme:   ptrToString(replicationNamingScheme),
		ReplicantAffinityPolicy: ptrToString(replicationpolicy),
		Networks:                buildPVMNetworks(networks),
	}
	if len(volids) > 0 {
		body.VolumeIds = volids
	}
	if d.Get(helpers.PIInstancePinPolicy) == "soft" || d.Get(helpers.PIInstancePinPolicy) == "hard" {
		body.PinPolicy = models.PinPolicy(pinpolicy)
	}

	if d.Get(helpers.PIVirtualCoresAssigned) != "" && d.Get(helpers.PIVirtualCoresAssigned) != 0 {
		assigned_virtual_cores := int64(d.Get(helpers.PIVirtualCoresAssigned).(int))
		body.VirtualCores = &models.VirtualCores{Assigned: &assigned_virtual_cores}
	} else {
		log.Printf("Virtual cores is not provided")
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	pvm, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	}, powerinstanceid, createTimeOut)

	if err != nil {
		return fmt.Errorf("failed to provision %v", err)
	} else {
		log.Printf("Printing the instance info %+v", &pvm)
	}

	var pvminstanceids []string
	if replicants > 1 {
		log.Printf("We are in a multi create mode")
		for i := 0; i < int(replicants); i++ {
			truepvmid := (*pvm)[i].PvmInstanceID
			log.Printf("Printing the instance id %s", *truepvmid)
			pvminstanceids = append(pvminstanceids, fmt.Sprintf("%s", *truepvmid))
			log.Printf("Printing each of the pvminstance ids %s", pvminstanceids)
		}

		d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, pvminstanceids))

	} else {
		log.Printf("Single Create Mode ")
		truepvmid := (*pvm)[0].PvmInstanceID
		d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, *truepvmid))

		pvminstanceids = append(pvminstanceids, *truepvmid)
		log.Printf("Printing the instance id .. after the create ... %s", *truepvmid)
	}

	for inst_count := 0; inst_count < len(pvminstanceids); inst_count++ {
		log.Printf("The pvm instance id is [%s] .Checking for status", pvminstanceids[inst_count])

		_, err = isWaitForPIMultiInstanceAvailable(client, pvminstanceids[inst_count], d.Timeout(schema.TimeoutCreate), powerinstanceid, instance_ready_status)
		if err != nil {
			log.Printf("exceeded the timeout value to get the instance")
			return err
		}
	}

	return resourceIBMPIMultiInstanceRead(d, meta)

}

func resourceIBMPIMultiInstanceRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the Multi PowerInstance Read code..")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	if d.Get(helpers.PIInstanceReplicants) != "" && d.Get(helpers.PIInstanceReplicants).(float64) > 1 {
		myvmdata, err := getmultivmdata(d.Id(), sess)
		if err != nil {
			return err
		}
		d.Set("multivmdata", myvmdata)

	} else {
		log.Printf("Single Instance Get Mode ")
		powerinstanceid := parts[0]
		powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
		powervmdata, err := powerC.Get(parts[1], powerinstanceid, getTimeOut)

		if err != nil {
			return fmt.Errorf("failed to get the instance %v", err)
		}

		d.Set(helpers.PIInstanceMemory, powervmdata.Memory)
		d.Set(helpers.PIInstanceProcessors, powervmdata.Processors)
		d.Set("status", powervmdata.Status)
		d.Set(helpers.PIInstanceProcType, powervmdata.ProcType)
		d.Set("migratable", powervmdata.Migratable)
		d.Set("min_processors", powervmdata.Minproc)
		d.Set(helpers.PIInstanceProgress, powervmdata.Progress)
		d.Set(helpers.PICloudInstanceId, powerinstanceid)
		d.Set("instance_id", powervmdata.PvmInstanceID)
		d.Set(helpers.PIInstanceName, powervmdata.ServerName)
		d.Set(helpers.PIInstanceImageName, powervmdata.ImageID)
		var networks []string
		networks = make([]string, 0)
		if powervmdata.Networks != nil {
			for _, n := range powervmdata.Networks {
				if n != nil {
					networks = append(networks, n.NetworkID)
				}

			}
		}
		d.Set(helpers.PIInstanceNetworkIds, networks)
		d.Set(helpers.PIInstanceVolumeIds, powervmdata.VolumeIds)
		d.Set(helpers.PIInstanceSystemType, powervmdata.SysType)
		d.Set("min_memory", powervmdata.Minmem)
		d.Set("max_processors", powervmdata.Maxproc)
		d.Set("max_memory", powervmdata.Maxmem)
		d.Set("pin_policy", powervmdata.PinPolicy)
		d.Set("operating_system", powervmdata.OperatingSystem)
		d.Set("os_type", powervmdata.OsType)

		if powervmdata.Addresses != nil {
			pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
			for i, pvmip := range powervmdata.Addresses {
				log.Printf("Now entering the powervm address space....")

				p := make(map[string]interface{})
				p["ip"] = pvmip.IP
				p["network_name"] = pvmip.NetworkName
				p["network_id"] = pvmip.NetworkID
				p["macaddress"] = pvmip.MacAddress
				p["type"] = pvmip.Type
				p["external_ip"] = pvmip.ExternalIP
				pvmaddress[i] = p
			}
			d.Set("addresses", pvmaddress)

		}

		if powervmdata.Health != nil {
			d.Set("health_status", powervmdata.Health.Status)
		}

		if powervmdata.VirtualCores.Assigned != nil {
			d.Set(helpers.PIVirtualCoresAssigned, powervmdata.VirtualCores.Assigned)
			d.Set("max_virtual_cores", powervmdata.VirtualCores.Max)
			d.Set("min_virtual_cores", powervmdata.VirtualCores.Min)
		}
	}

	return nil

}

func resourceIBMPIMultiInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Calling the Multi Instance Update method")
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return fmt.Errorf("failed to get the session from the IBM Cloud Service")
	}
	powerinstanceid := parts[0]
	instlist, _ := pvmparts(parts[1])
	log.Printf("the size of inslist is %d", len(instlist))

	if len(instlist) > 1 {
		log.Printf("Performing the update operation on a multi-lpar setup")
		err := executeMultivmChange(d, sess)
		if err != nil {
			return fmt.Errorf("failed to execute the change ")
		}
	}

	name := d.Get(helpers.PIInstanceName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	processortype := d.Get(helpers.PIInstanceProcType).(string)
	assigned_virtual_cores := int64(d.Get(helpers.PIVirtualCoresAssigned).(int))

	if d.Get("health_status") == "WARNING" {

		return fmt.Errorf("the operation cannot be performed when the lpar health in the WARNING State")
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	//if d.HasChange(helpers.PIInstanceName) || d.HasChange(helpers.PIInstanceProcessors) || d.HasChange(helpers.PIInstanceProcType) || d.HasChange(helpers.PIInstancePinPolicy){
	if d.HasChange(helpers.PIInstanceProcType) {

		// Stop the lpar
		processortype := d.Get(helpers.PIInstanceProcType).(string)
		if d.Get("status") == "SHUTOFF" {
			log.Printf("the lpar is in the shutoff state. Nothing to do . Moving on ")
		} else {

			body := &models.PVMInstanceAction{
				Action: ptrToString("immediate-shutdown"),
			}
			resp, err := client.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: body}, parts[1], powerinstanceid, postTimeOut)
			if err != nil {
				log.Printf("Stop Action failed on [%s]", name)

				return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)

			}
			log.Printf("Getting the response from the shutdown ... %v", resp)

			_, err = isWaitForPIInstanceStopped(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid)
			if err != nil {
				return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)
			}
		}

		// Modify

		log.Printf("At this point the lpar should be off. Executing the Processor Update Change")
		updatebody := &models.PVMInstanceUpdate{ProcType: processortype}
		updateresp, err := client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: updatebody}, updateTimeOut)
		if err != nil {
			return fmt.Errorf("failed to perform the modify operation on the pvm instance %v", err)
		} else {
			log.Printf("Getting the response from the change %s", updateresp.StatusURL)
		}
		// To check if the verify resize operation is complete.. and then it will go to SHUTOFF

		_, err = isWaitForPIInstanceStopped(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid)
		if err != nil {
			return err
		}

		// Start

		startbody := &models.PVMInstanceAction{
			Action: ptrToString("start"),
		}
		startresp, err := client.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: startbody}, parts[1], powerinstanceid, postTimeOut)
		if err != nil {
			return fmt.Errorf("failed to perform the start action on the pvm instance %v", err)
		} else {
			log.Printf("Performing the start operation on the pvminstance")
		}

		log.Printf("Getting the response from the start %s", startresp)

		_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
		if err != nil {
			return err
		}

	}

	// Start of the change for Memory and Processors
	if d.HasChange(helpers.PIVirtualCoresAssigned) {
		log.Printf("Calling the change for the Virtual Cores")
		max_vc := d.Get("max_virtual_cores").(int)
		log.Printf("the max virtual cores is set to %d", max_vc)
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]

		client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

		body := &models.PVMInstanceUpdate{
			VirtualCores: &models.VirtualCores{Assigned: &assigned_virtual_cores},
		}
		resp, err := client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
		if err != nil {
			return fmt.Errorf("failed to update the lpar with the change for virtual cores")
		}
		log.Printf("Getting the response from the bigger change block %s", resp.StatusURL)

		_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
		if err != nil {
			return err
		}
	}

	if d.HasChange(helpers.PIInstanceMemory) || d.HasChange(helpers.PIInstanceProcessors) {
		log.Printf("Checking for cpu / memory change..and also virtual cores")

		max_mem_lpar := d.Get("max_memory").(float64)
		max_cpu_lpar := d.Get("max_processors").(float64)
		//log.Printf("the required memory is set to [%d] and current max memory is set to  [%d] ", int(mem), int(max_mem_lpar))

		if mem > max_mem_lpar || procs > max_cpu_lpar {
			log.Printf("Will require a shutdown to perform the change")

		} else {
			log.Printf("max_mem_lpar is set to %f", max_mem_lpar)
			log.Printf("max_cpu_lpar is set to %f", max_cpu_lpar)
		}

		//if d.GetOkExists("reboot_for_resource_change")

		if mem > max_mem_lpar || procs > max_cpu_lpar {

			_, err = performChangeAndReboot(client, parts[1], powerinstanceid, mem, procs)
			//_, err = stopLparForResourceChange(client, parts[1], powerinstanceid)
			if err != nil {
				return fmt.Errorf("failed to perform the operation for the change")
			}

		} else {
			log.Printf("Memory change is within limits")
			parts, err := idParts(d.Id())
			if err != nil {
				return err
			}
			powerinstanceid := parts[0]

			client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

			body := &models.PVMInstanceUpdate{
				Memory:     mem,
				ProcType:   processortype,
				Processors: procs,
				ServerName: name,
			}
			body.VirtualCores = &models.VirtualCores{Assigned: &assigned_virtual_cores}

			resp, err := client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
			if err != nil {
				return fmt.Errorf("failed to update the lpar with the change")
			}
			log.Printf("Getting the response from the bigger change block %s", resp.StatusURL)

			_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
			if err != nil {
				return err
			}

		}

	}

	return resourceIBMPIInstanceRead(d, meta)

}

func resourceIBMPIMultiInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the Multi Instance Delete method")
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]

	instlist, _ := pvmparts(parts[1])
	log.Printf("the size of inslist is %d", len(instlist))
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	if len(instlist) > 1 {
		log.Printf("we are in a multi-vm delete mode")
		for inst := range instlist {
			log.Printf("**** IN THE READINSTANCE CODE ***")
			log.Printf("Printing the list of instances %s", instlist[inst])
			r := strings.NewReplacer("[", "", "]", "")
			myinstance := r.Replace(instlist[inst])
			log.Printf("The instance to be deleted has the id .. [%s]", myinstance)
			err = client.Delete(myinstance, powerinstanceid, deleteTimeOut)
			if err != nil {
				return fmt.Errorf("failed to get the instance %v", err)
			}
			_, err = isWaitForPIInstanceDeleted(client, myinstance, d.Timeout(schema.TimeoutDelete), powerinstanceid)
			if err != nil {
				return err
			}

		}
		d.SetId("")
		return nil
	}
	log.Printf("Deleting the instance with name/id %s and cloud_instance_id %s", parts[1], powerinstanceid)
	err = client.Delete(parts[1], powerinstanceid, deleteTimeOut)
	if err != nil {

		return fmt.Errorf("failed to perform the delete action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceDeleted(client, parts[1], d.Timeout(schema.TimeoutDelete), powerinstanceid)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// Exists

func resourceIBMPIMultiInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	log.Printf("Calling the Multi PowerInstance Exists method")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	powerinstanceid := parts[0]
	instlist, _ := pvmparts(parts[1])
	log.Printf("Instance Count  is %d", len(instlist))
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	if len(instlist) > 1 {
		log.Printf("we are in a multi-vm create mode")
		for inst := range instlist {
			log.Printf("**** IN THE READINSTANCE CODE ***")
			log.Printf("Printing the list of instances %s", instlist[inst])
			r := strings.NewReplacer("[", "", "]", "")
			myinstance := r.Replace(instlist[inst])
			//log.Printf("The instance to be checked  is ... %s ******", myinstance)
			instance, err := client.Get(myinstance, powerinstanceid, deleteTimeOut)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok {
					if apiErr.StatusCode() == 404 {
						return false, nil
					}
				}
				return false, fmt.Errorf("error communicating with the API: %s", err)
			}
			truepvmid := *instance.PvmInstanceID
			return truepvmid == myinstance, nil
		}
	}
	instance, err := client.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("error communicating with the API: %s", err)
	}

	truepvmid := *instance.PvmInstanceID
	return truepvmid == parts[1], nil
}

func executeMultivmChange(d *schema.ResourceData, pvmsession *ibmpisession.IBMPISession) error {

	log.Printf("Now in the executeMultiVM Change.. ")
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	instlist, _ := pvmparts(parts[1])
	if err != nil {
		fmt.Errorf("failed to get the parts %v", err)
	}
	log.Printf("We are now in the multi vm update mode.. ")
	log.Printf("the size of inslist is %d and the powerinstanceid is [%s]", len(instlist), powerinstanceid)
	powervcclient := st.NewIBMPIInstanceClient(pvmsession, powerinstanceid)
	log.Printf("getting the processortype [%s]", d.Get(helpers.PIInstanceProcType).(string))
	if d.HasChange(helpers.PIInstanceProcType) {
		for inst := range instlist {
			r := strings.NewReplacer("[", "", "]", "")
			myinstance := r.Replace(instlist[inst])
			// Stop the lpar
			log.Printf("are we now in the haschange state ?? ")
			processortype := d.Get(helpers.PIInstanceProcType).(string)
			if d.Get("status") == "SHUTOFF" {
				log.Printf("the lpar is in the shutoff state. Nothing to do . Moving on ")
			} else {
				body := &models.PVMInstanceAction{
					Action: ptrToString("immediate-shutdown"),
				}
				resp, err := powervcclient.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: body}, myinstance, powerinstanceid, postTimeOut)
				if err != nil {
					log.Printf("Stop Action failed on [%s]", myinstance)

					return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)

				}
				log.Printf("Getting the response from the shutdown ... %v", resp)

				_, err = isWaitForPIInstanceStopped(powervcclient, myinstance, d.Timeout(schema.TimeoutUpdate), powerinstanceid)
				if err != nil {
					return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)
				}
				log.Printf("executing stop on the instance for modify operation with id [%s] ", myinstance)
			}
			// Modify

			log.Printf("At this point the lpar should be off. Executing the Processor Update Change for lpar with id [%s] ", myinstance)
			updatebody := &models.PVMInstanceUpdate{ProcType: processortype}
			updateresp, err := powervcclient.Update(myinstance, powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: updatebody}, updateTimeOut)
			if err != nil {
				return fmt.Errorf("failed to perform the modify operation on the pvm instance %v", err)
			} else {
				log.Printf("Getting the response from the change %s", updateresp.StatusURL)
			}
			// To check if the verify resize operation is complete.. and then it will go to SHUTOFF

			_, err = isWaitForPIInstanceStopped(powervcclient, myinstance, d.Timeout(schema.TimeoutUpdate), powerinstanceid)
			if err != nil {
				return err
			}

			// Start

			startbody := &models.PVMInstanceAction{
				Action: ptrToString("start"),
			}
			startresp, err := powervcclient.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: startbody}, myinstance, powerinstanceid, postTimeOut)
			if err != nil {
				return fmt.Errorf("failed to perform the start action on the pvm instance %v", err)
			} else {
				log.Printf("Performing the start operation on the pvminstance")
			}

			log.Printf("Getting the response from the start %s", startresp)

			_, err = isWaitForPIInstanceAvailable(powervcclient, myinstance, d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
			if err != nil {
				return err
			}

		}

	}

	return err
}

func getmultivmdata(pvmid string, pvmsession *ibmpisession.IBMPISession) ([]map[string]interface{}, error) {

	log.Printf("Now in the getmultivm data code.. ")
	parts, err := idParts(pvmid)
	if err != nil {
		return nil, err
	}

	powerinstanceid := parts[0]
	instlist, _ := pvmparts(parts[1])
	if err != nil {
		fmt.Errorf("failed to get the parts %v", err)
	}
	log.Printf("We are now in the multi vm create mode.. ")
	powervcclient := st.NewIBMPIInstanceClient(pvmsession, powerinstanceid)
	vmdata := make([]map[string]interface{}, len(instlist))

	for inst := range instlist {
		log.Printf("**** IN THE READINSTANCE CODE ***")
		log.Printf("Printing the list of instances %s", instlist[inst])
		r := strings.NewReplacer("[", "", "]", "")
		myinstance := r.Replace(instlist[inst])
		log.Printf("The instance is ... %s ******", myinstance)
		powervmdata, err := powervcclient.Get(myinstance, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, fmt.Errorf("failed to get the instance %v", err)
		}

		p := make(map[string]interface{})
		if powervmdata.Addresses != nil {
			pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
			for i, pvmip := range powervmdata.Addresses {
				log.Printf("Now entering the powervm address space....")

				q := make(map[string]interface{})
				q["ip"] = pvmip.IP
				q["network_name"] = pvmip.NetworkName
				q["network_id"] = pvmip.NetworkID
				q["macaddress"] = pvmip.MacAddress
				q["type"] = pvmip.Type
				q["external_ip"] = pvmip.ExternalIP
				pvmaddress[i] = q
			}
			p["addresses"] = pvmaddress
		}

		p["max_virtual_cores"] = powervmdata.VirtualCores.Max
		p[helpers.PIVirtualCoresAssigned] = powervmdata.VirtualCores.Assigned
		p["os_type"] = powervmdata.OsType
		p["max_processors"] = powervmdata.Maxproc
		p["max_memory"] = powervmdata.Maxmem
		p["instance_name"] = powervmdata.ServerName
		p["health_status"] = powervmdata.Health.Status
		p["min_processors"] = powervmdata.Minproc
		p["instance_id"] = powervmdata.PvmInstanceID
		p["status"] = powervmdata.Status
		p["operating_system"] = powervmdata.OperatingSystem
		p["max_virtual_cores"] = powervmdata.VirtualCores.Max
		p["min_virtual_cores"] = powervmdata.VirtualCores.Max
		vmdata[inst] = p

	}
	return vmdata, err
}

func isWaitForPIMultiInstanceAvailable(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string, instance_ready_status string) (interface{}, error) {
	log.Printf("Waiting for Multi PIInstance (%s) to be available and active ", id)
	var queryTimeOut time.Duration
	if instance_ready_status == "WARNING" {
		queryTimeOut = warningTimeOut
	} else {
		queryTimeOut = activeTimeOut
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", "BUILD", helpers.PIInstanceHealthWarning},
		Target:     []string{"OK", "ACTIVE", helpers.PIInstanceHealthOk, "ERROR"},
		Refresh:    isPIMultiInstanceRefreshFunc(client, id, powerinstanceid, instance_ready_status),
		Delay:      10 * time.Second,
		MinTimeout: queryTimeOut,
		Timeout:    timeout,
	}

	return stateConf.WaitForState()
}

func isPIMultiInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid, instance_ready_status string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, "", err
		}
		allowableStatus := instance_ready_status
		log.Printf("*** InstanceRefreshFunc Multi-VM Model- the allowable instance status is [%s]", allowableStatus)
		if *pvm.Status == helpers.PIInstanceAvailable && (pvm.Health.Status == allowableStatus) {
			////if *pvm.Status == helpers.PIInstanceAvailable  {
			log.Printf("The health status is now %s", allowableStatus)

			return pvm, helpers.PIInstanceAvailable, nil

		}

		return pvm, helpers.PIInstanceBuilding, nil
	}
}
