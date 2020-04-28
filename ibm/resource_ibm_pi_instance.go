package ibm

import (
	"encoding/base64"
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
	"time"
)

func resourceIBMPIInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIInstanceCreate,
		Read:     resourceIBMPIInstanceRead,
		Update:   resourceIBMPIInstanceUpdate,
		Delete:   resourceIBMPIInstanceDelete,
		Exists:   resourceIBMPIInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"migratable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"min_processors": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"min_memory": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_processors": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_memory": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			helpers.PIInstanceNetworkIds: {
				Type:             schema.TypeSet,
				Required:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
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
			},

			helpers.PIInstanceUserData: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded data to be passed in for invoking a cloud init script",
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
						/*"version": {
							Type:     schema.TypeFloat,
							Computed: true,
						},*/
					},
				},
			},

			"health_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			helpers.PIInstanceImageName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceProcessors: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			helpers.PIInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared", "capped"}),
			},
			helpers.PIInstanceSSHKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceMemory: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			helpers.PIInstanceSystemType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"any", "s922", "e880", "e980"}),
			},
			helpers.PIInstanceReplicants: {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "1",
			},
			helpers.PIInstanceReplicationPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"affinity", "anti-affinity", "none"}),
				Default:      "none",
			},
			helpers.PIInstanceReplicationScheme: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"prefix", "suffix"}),
				Default:      "suffix",
			},
			helpers.PIInstanceProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
			helpers.PIInstancePinPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"none", "soft", "hard"}),
				Description:  "Pin Policy that is attached to a PVMInstance",
				Default:      "none",
			},
			"pin_policy": {
				Type:     schema.TypeString,
				Computed: true,
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
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	name := d.Get(helpers.PIInstanceName).(string)
	sshkey := d.Get(helpers.PIInstanceSSHKeyName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	systype := d.Get(helpers.PIInstanceSystemType).(string)
	networks := expandStringList((d.Get(helpers.PIInstanceNetworkIds).(*schema.Set)).List())
	volids := expandStringList((d.Get(helpers.PIInstanceVolumeIds).(*schema.Set)).List())
	replicants := d.Get(helpers.PIInstanceReplicants).(float64)
	replicationpolicy := d.Get(helpers.PIInstanceReplicationPolicy).(string)
	replicationNamingScheme := d.Get(helpers.PIInstanceReplicationScheme).(string)
	imageid := d.Get(helpers.PIInstanceImageName).(string)
	processortype := d.Get(helpers.PIInstanceProcType).(string)
	pinpolicy := string(helpers.PIInstancePinPolicy)
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

	//publicinterface := d.Get(helpers.PIInstancePublicNetwork).(bool)
	if d.Get(helpers.PIInstancePinPolicy) == "" {
		pinpolicy = "none"
	}

	pinpolicy = d.Get(helpers.PIInstancePinPolicy).(string)

	body := &models.PVMInstanceCreate{
		NetworkIds: networks, Processors: &procs, Memory: &mem, ServerName: ptrToString(name),
		SysType:                 systype,
		KeyPairName:             sshkey,
		ImageID:                 ptrToString(imageid),
		ProcType:                ptrToString(processortype),
		Replicants:              replicants,
		UserData:                user_data,
		ReplicantNamingScheme:   ptrToString(replicationNamingScheme),
		ReplicantAffinityPolicy: ptrToString(replicationpolicy),
	}
	if len(volids) > 0 {
		body.VolumeIds = volids
	}
	if d.Get(helpers.PIInstancePinPolicy) == "soft" || d.Get(helpers.PIInstancePinPolicy) == "hard" {
		body.PinPolicy = models.PinPolicy(pinpolicy)
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	pvm, _, _, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	}, powerinstanceid)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("Printing the instance info %+v", &pvm)

	truepvmid := (*pvm)[0].PvmInstanceID
	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, *truepvmid))

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
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(parts[1], powerinstanceid)

	if err != nil {
		return err
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
	d.Set(helpers.PIInstanceNetworkIds, newStringSet(schema.HashString, networks))
	d.Set(helpers.PIInstanceVolumeIds, powervmdata.VolumeIds)
	d.Set(helpers.PIInstanceSystemType, powervmdata.SysType)
	d.Set("min_memory", powervmdata.Minmem)
	d.Set("max_processors", powervmdata.Maxproc)
	d.Set("max_memory", powervmdata.Maxmem)
	d.Set("pin_policy", powervmdata.PinPolicy)

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

		//log.Printf("Printing the value after the read - this should set it.... %+v", pvmaddress)

	}

	if powervmdata.Health != nil {
		d.Set("health_status", powervmdata.Health.Status)

	}

	return nil

}

func resourceIBMPIInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	if d.HasChange(helpers.PIInstanceName) || d.HasChange(helpers.PIInstanceMemory) || d.HasChange(helpers.PIInstanceProcessors) || d.HasChange(helpers.PIInstanceProcType) || d.HasChange(
		helpers.PIInstancePinPolicy) {
		name := d.Get(helpers.PIInstanceName).(string)
		mem := d.Get(helpers.PIInstanceMemory).(float64)
		procs := d.Get(helpers.PIInstanceProcessors).(float64)
		processortype := d.Get(helpers.PIInstanceProcType).(string)
		pinpolicy := d.Get(helpers.PIInstancePinPolicy).(string)

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
			PinPolicy:  models.PinPolicy(pinpolicy),
		}

		resp, err := client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body})
		if err != nil {
			return err
		}

		log.Printf("Getting the response %s", resp.StatusURL)

		_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid)
		if err != nil {
			return err
		}

	}

	return resourceIBMPIInstanceRead(d, meta)

}

func resourceIBMPIInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the Instance Delete method")
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	log.Printf("Deleting the instance with name/id %s and cloud_instance_id %s", parts[1], powerinstanceid)
	err = client.Delete(parts[1], powerinstanceid)
	if err != nil {
		return err
	}

	_, err = isWaitForPIInstanceDeleted(client, parts[1], d.Timeout(schema.TimeoutDelete), powerinstanceid)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// Exists

func resourceIBMPIInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	instance, err := client.Get(parts[1], powerinstanceid)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	truepvmid := *instance.PvmInstanceID
	return truepvmid == parts[1], nil
}

func isWaitForPIInstanceDeleted(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIInstanceDeleting},
		Target:     []string{helpers.PIInstanceNotFound},
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
			return pvm, helpers.PIInstanceNotFound, nil

		}
		return pvm, helpers.PIInstanceNotFound, nil

	}
}

func isWaitForPIInstanceAvailable(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be available and sleeping ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", "BUILD", helpers.PIInstanceHealthWarning},
		Target:     []string{"OK", "ACTIVE", helpers.PIInstanceHealthOk},
		Refresh:    isPIInstanceRefreshFunc(client, id, powerinstanceid),
		Delay:      3 * time.Minute,
		MinTimeout: 60 * time.Second,
		Timeout:    timeout,
	}

	return stateConf.WaitForState()
}

func isPIInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		//if pvm.Health.Status == helpers.PIInstanceHealthOk {
		if *pvm.Status == helpers.PIInstanceAvailable && pvm.Health.Status == helpers.PIInstanceHealthOk {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			//return pvm, helpers.PIInstanceHealthOk, nil
			return pvm, helpers.PIInstanceAvailable, nil
			//}
		}

		return pvm, helpers.PIInstanceBuilding, nil
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

func checkBase64(input string) error {
	fmt.Println("Calling the checkBase64")
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Printf("Data is correctly Encoded to Base64 %s", data)
	return err

}
