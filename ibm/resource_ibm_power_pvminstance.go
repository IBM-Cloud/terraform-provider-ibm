package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.ibm.com/Bluemix/power-go-client/power/models"

	"log"
	"time"
)

const (
	PowerPVMInstanceName          = "server_name"
	PowerPVMInstanceDate          = "creation_date"
	PowerPVMInstanceSSHKeyName    = "key_pair_name"
	PowerPVMImageName             = "image_id"
	PowerPVMInstanceProcessors    = "processors"
	PowerPVMInstanceProcType      = "proc_type"
	PowerPVMInstanceMemory        = "memory"
	PowerPVMInstanceSystemType    = "sys_type"
	PowerPVMInstanceId            = "pvm_instance_id"
	PowerPVMInstanceDiskSize      = "pvm_disk_size"
	PowerPVMInstanceStatus        = "status"
	PowerPVMInstanceMinProc       = "min_proc"
	PowerPVMInstanceVolumeIds     = "volume_ids"
	PowerPVMInstanceNetworkIds    = "network_ids"
	PowerPVMInstanceAddress       = "addresses"
	PowerPVMInstanceNetworkName   = "name"
	PowerPVMInstanceMigratable    = "migratable"
	PowerPVMInstanceAvailable     = "ACTIVE"
	PowerPVMInstanceHealthOk      = "OK"
	PowerPVMInstanceHealthWarning = "WARNING"
	PowerPVMInstanceBuilding      = "BUILD"
	PowerPVMInstanceDeleting      = "DELETING"
	PowerPVMInstanceNetworkId     = "networkid"
	PowerPVMInstanceNetworkCidr   = "cidr"
	PowerPVMInstanceNotFound      = "Not Found"
	PowerPVMInstanceHealthStatus  = "health_status"
	PowerPVMReplicants            = "replicants"
	PowerPVMReplicationPolicy     = "replication_policy"
	PowerPVMProgress              = "progress"
)

func resourceIBMPowerPVMInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPowerPVMInstanceCreate,
		Read:   resourceIBMPowerPVMInstanceRead,
		Update: resourceIBMPowerPVMInstanceUpdate,
		Delete: resourceIBMPowerPVMInstanceDelete,
		//Exists:   resourceIBMPowerPVMInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			PowerPVMInstanceDiskSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			PowerPVMInstanceStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PowerPVMInstanceMigratable: {
				Type:     schema.TypeBool,
				Required: true,
			},
			PowerPVMInstanceMinProc: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			PowerPVMInstanceNetworkIds: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			PowerPVMInstanceVolumeIds: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
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
						"networkid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networkname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*"version": {
							Type:     schema.TypeFloat,
							Computed: true,
						},*/
					},
				},
			},

			PowerPVMInstanceHealthStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PowerPVMInstanceId: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			PowerPVMInstanceDate: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PowerPVMImageName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PowerPVMInstanceProcessors: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			PowerPVMInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PowerPVMInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared"}),
			},
			PowerPVMInstanceSSHKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PowerPVMInstanceMemory: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			PowerPVMInstanceSystemType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"any", "s922", "e880"}),
			},
			PowerPVMReplicants: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			PowerPVMReplicationPolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},
			PowerPVMProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
		},
	}
}

func resourceIBMPowerPVMInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	name := d.Get(PowerPVMInstanceName).(string)
	sshkey := d.Get(PowerPVMInstanceSSHKeyName).(string)
	mem := d.Get(PowerPVMInstanceMemory).(float64)
	procs := d.Get(PowerPVMInstanceProcessors).(float64)
	migrateable := d.Get(PowerPVMInstanceMigratable).(bool)
	systype := d.Get(PowerPVMInstanceSystemType).(string)
	networks := expandStringList((d.Get(PowerPVMInstanceNetworkIds).(*schema.Set)).List())
	volids := expandStringList((d.Get(PowerPVMInstanceVolumeIds).(*schema.Set)).List())
	replicants := d.Get("replicants").(float64)
	if d.Get("replicants") == "" {
		replicants = 1
	}
	replicationpolicy := d.Get("replicationpolicy").(string)
	if d.Get("replicationpolicy") == "" {
		replicationpolicy = "none"
	}

	imageid := d.Get(PowerPVMImageName).(string)

	processortype := d.Get(PowerPVMInstanceProcType).(string)

	body := &models.PVMInstanceCreate{

		VolumeIds: volids, NetworkIds: networks, Processors: &procs, Memory: &mem, ServerName: ptrToString(name),
		Migratable:              &migrateable,
		SysType:                 ptrToString(systype),
		KeyPairName:             sshkey,
		ImageID:                 ptrToString(imageid),
		ProcType:                ptrToString(processortype),
		Replicants:              replicants,
		ReplicantAffinityPolicy: ptrToString(replicationpolicy),
	}

	client := st.NewPowerPvmClient(sess)

	pvm, _, _, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	})
	log.Printf("the number of instances is %d", len(*pvm))

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("Printing the instance info %+v", &pvm)

	truepvmid := (*pvm)[0].PvmInstanceID
	d.SetId(*truepvmid)
	//d.Set("addresses",(*pvm)[0].Addresses)

	log.Printf("Printing the instance id .. after the create ... %s", *truepvmid)

	_, err = isWaitForPowerPVMInstanceAvailable(client, *truepvmid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMPowerPVMInstanceRead(d, meta)
}

func resourceIBMPowerPVMInstanceRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the PowerInstance Read code..")

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	powerC := st.NewPowerPvmClient(sess)
	powervmdata, err := powerC.Get(d.Id())

	if err != nil {
		return err
	}

	pvminstanceid := *powervmdata.PvmInstanceID

	log.Printf("The Power pvm instance id is %s", pvminstanceid)

	d.SetId(pvminstanceid)
	d.Set("memory", powervmdata.Memory)
	d.Set("processors", powervmdata.Processors)
	d.Set("status", powervmdata.Status)
	d.Set("proctype", powervmdata.ProcType)
	d.Set("migratable", powervmdata.Migratable)
	d.Set("minproc", powervmdata.Minproc)
	d.Set("progress", powervmdata.Progress)

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {
			log.Printf("Now entering the powervm address space....")

			p := make(map[string]interface{})
			p["ip"] = pvmip.IP
			p["networkname"] = pvmip.NetworkName
			p["networkid"] = pvmip.NetworkID
			p["macaddress"] = pvmip.MacAddress
			p["type"] = pvmip.Type
			pvmaddress[i] = p
		}
		d.Set("addresses", pvmaddress)

		//log.Printf("Printing the value after the read - this should set it.... %+v", pvmaddress)

	}

	if powervmdata.Health != nil {
		d.Set("healthstatus", powervmdata.Health.Status)

	}

	return nil

}

func resourceIBMPowerPVMInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerPvmClient(sess)

	name := d.Get(PowerPVMInstanceName).(string)
	mem := d.Get(PowerPVMInstanceMemory).(float64)
	procs := d.Get(PowerPVMInstanceProcessors).(float64)
	migrateable := d.Get(PowerPVMInstanceMigratable).(bool)
	processortype := d.Get(PowerPVMInstanceProcType).(string)

	body := &models.PVMInstanceUpdate{
		Memory:     mem,
		Migratable: &migrateable,
		ProcType:   processortype,
		Processors: procs,
		ServerName: name,
	}

	resp, err := client.Update(d.Id(), &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body})
	if err != nil {
		return err
	}

	log.Printf("Getting the response %s", resp.StatusURL)

	_, err = isWaitForPowerPVMInstanceAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMPowerPVMInstanceRead(d, meta)

	return nil
}

func resourceIBMPowerPVMInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).PowerSession()
	client := st.NewPowerPvmClient(sess)
	err := client.Delete(data.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForPowerPvmInstanceDeleted(client, data.Id(), data.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}

func isWaitForPowerPvmInstanceDeleted(client *st.PowerPvmClient, id string, timeout time.Duration) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", PowerPVMInstanceDeleting},
		Target:     []string{PowerPVMInstanceNotFound},
		Refresh:    isPowerPvmInstanceDeleteRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPowerPvmInstanceDeleteRefreshFunc(client *st.PowerPvmClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pvm, err := client.Get(id)
		if err != nil {
			log.Printf("The power vm does not exist")
			return pvm, PowerPVMInstanceNotFound, nil

		}
		return pvm, PowerPVMInstanceNotFound, nil

	}
}

func isWaitForPowerPVMInstanceAvailable(client *st.PowerPvmClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for PowerPVMInstance (%s) to be available and sleeping ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", PowerPVMInstanceHealthWarning},
		Target:     []string{"OK", PowerPVMInstanceHealthOk},
		Refresh:    isPowerPVMInstanceRefreshFunc(client, id),
		Delay:      3 * time.Minute,
		MinTimeout: 30 * time.Second,
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPowerPVMInstanceRefreshFunc(client *st.PowerPvmClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if pvm.Health.Status == PowerPVMInstanceHealthOk {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			return pvm, PowerPVMInstanceHealthOk, nil
			//}
		}

		return pvm, PowerPVMInstanceHealthWarning, nil
	}
}

func checkActive(vminstance *models.PVMInstance) bool {

	log.Printf("Calling the check vm status function and the health status is %s", vminstance.Health.Status)
	activeStatus := false

	if vminstance.Health.Status == "OK" {
		//if *vminstance.Status == "active" {
		log.Printf(" The status of the vm is now set to what we want it to be %s", vminstance.Health.Status)
		activeStatus = true

	}
	return activeStatus
}

func checkifIP(vmip []*models.PVMInstanceAddress) bool {
	log.Printf("calling the check ip function ")
	ipexists := false
	if len(vmip) > 0 {
		ipexists = true
	}
	return ipexists
}
