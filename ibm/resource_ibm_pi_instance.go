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
	PIInstanceName              = "servername"
	PIInstanceDate              = "creationdate"
	PIInstanceSSHKeyName        = "keypairname"
	PIInstanceImageName         = "imageid"
	PIInstanceProcessors        = "processors"
	PIInstanceProcType          = "proctype"
	PIInstanceMemory            = "memory"
	PIInstanceSystemType        = "systype"
	PIInstanceId                = "pvminstanceid"
	PIInstanceDiskSize          = "pvmdisksize"
	PIInstanceStatus            = "status"
	PIInstanceMinProc           = "minproc"
	PIInstanceVolumeIds         = "volumeids"
	PIInstanceNetworkIds        = "networkids"
	PIInstanceAddress           = "addresses"
	PIInstanceNetworkName       = "name"
	PIInstanceMigratable        = "migratable"
	PIInstanceAvailable         = "ACTIVE"
	PIInstanceHealthOk          = "OK"
	PIInstanceHealthWarning     = "WARNING"
	PIInstanceBuilding          = "BUILD"
	PIInstanceDeleting          = "DELETING"
	PIInstanceNetworkId         = "networkid"
	PIInstanceNetworkCidr       = "cidr"
	PIInstanceNotFound          = "Not Found"
	PIInstanceHealthStatus      = "healthstatus"
	PIInstanceReplicants        = "replicants"
	PIInstanceReplicationPolicy = "replicationpolicy"
	PIInstanceProgress          = "progress"
)

func resourceIBMPIInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPIInstanceCreate,
		Read:   resourceIBMPIInstanceRead,
		Update: resourceIBMPIInstanceUpdate,
		Delete: resourceIBMPIInstanceDelete,
		//Exists:   resourceIBMPIInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"powerinstanceid": {
				Type:     schema.TypeString,
				Required: true,
			},
			PIInstanceDiskSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			PIInstanceStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PIInstanceMigratable: {
				Type:     schema.TypeBool,
				Required: true,
			},
			PIInstanceMinProc: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			PIInstanceNetworkIds: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			PIInstanceVolumeIds: {
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

			PIInstanceHealthStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PIInstanceId: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			PIInstanceDate: {
				Type:     schema.TypeString,
				Computed: true,
			},
			PowerPVMImageName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PIInstanceProcessors: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			PIInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PIInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared"}),
			},
			PIInstanceSSHKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PIInstanceMemory: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			PIInstanceSystemType: {
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

func resourceIBMPIInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	name := d.Get(PIInstanceName).(string)
	sshkey := d.Get(PIInstanceSSHKeyName).(string)
	mem := d.Get(PIInstanceMemory).(float64)
	procs := d.Get(PIInstanceProcessors).(float64)
	migrateable := d.Get(PIInstanceMigratable).(bool)
	systype := d.Get(PIInstanceSystemType).(string)
	networks := expandStringList((d.Get(PIInstanceNetworkIds).(*schema.Set)).List())
	volids := expandStringList((d.Get(PIInstanceVolumeIds).(*schema.Set)).List())
	replicants := d.Get("replicants").(float64)
	if d.Get("replicants") == "" {
		replicants = 1
	}
	replicationpolicy := d.Get("replicationpolicy").(string)
	if d.Get("replicationpolicy") == "" {
		replicationpolicy = "none"
	}

	imageid := d.Get(PIInstanceImageName).(string)

	processortype := d.Get(PIInstanceProcType).(string)

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

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	pvm, _, _, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	}, powerinstanceid)
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

	_, err = isWaitForPIInstanceAvailable(client, *truepvmid, d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	return resourceIBMPIInstanceRead(d, meta)
}

func resourceIBMPIInstanceRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the PowerInstance Read code..")

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(d.Id(), powerinstanceid)

	if err != nil {
		return err
	}

	pvminstanceid := *powervmdata.PvmInstanceID

	log.Printf("The Power pvm instance id is %s", pvminstanceid)
	log.Printf("the power vm address data is %s", powervmdata.Addresses)
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

func resourceIBMPIInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()

	name := d.Get(PIInstanceName).(string)
	mem := d.Get(PIInstanceMemory).(float64)
	procs := d.Get(PIInstanceProcessors).(float64)
	migrateable := d.Get(PIInstanceMigratable).(bool)
	processortype := d.Get(PIInstanceProcType).(string)
	powerinstanceid := d.Get("powerinstanceid").(string)

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	body := &models.PVMInstanceUpdate{
		Memory:     mem,
		Migratable: &migrateable,
		ProcType:   processortype,
		Processors: procs,
		ServerName: name,
	}

	resp, err := client.Update(d.Id(), powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body})
	if err != nil {
		return err
	}

	log.Printf("Getting the response %s", resp.StatusURL)

	_, err = isWaitForPIInstanceAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	return resourceIBMPIInstanceRead(d, meta)

	return nil
}

func resourceIBMPIInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := data.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	err := client.Delete(data.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	_, err = isWaitForPIInstanceDeleted(client, data.Id(), data.Timeout(schema.TimeoutDelete), powerinstanceid)
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}

// Exists

func resourceIBMPIInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	instance, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return instance.PvmInstanceID == &id, nil
}

func isWaitForPIInstanceDeleted(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", PIInstanceDeleting},
		Target:     []string{PIInstanceNotFound},
		Refresh:    isPIInstanceDeleteRefreshFunc(client, id, powerinstanceid),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPIInstanceDeleteRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			log.Printf("The power vm does not exist")
			return pvm, PIInstanceNotFound, nil

		}
		return pvm, PIInstanceNotFound, nil

	}
}

func isWaitForPIInstanceAvailable(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be available and sleeping ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", PIInstanceHealthWarning},
		Target:     []string{"OK", PIInstanceHealthOk},
		Refresh:    isPIInstanceRefreshFunc(client, id, powerinstanceid),
		Delay:      3 * time.Minute,
		MinTimeout: 30 * time.Second,
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if pvm.Health.Status == PIInstanceHealthOk {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			return pvm, PIInstanceHealthOk, nil
			//}
		}

		return pvm, PIInstanceHealthWarning, nil
	}
}

func checkPIActive(vminstance *models.PVMInstance) bool {

	log.Printf("Calling the check vm status function and the health status is %s", vminstance.Health.Status)
	activeStatus := false

	if vminstance.Health.Status == "OK" {
		//if *vminstance.Status == "active" {
		log.Printf(" The status of the vm is now set to what we want it to be %s", vminstance.Health.Status)
		activeStatus = true

	}
	return activeStatus
}
