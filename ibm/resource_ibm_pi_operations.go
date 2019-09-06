package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	_ "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/helpers"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.ibm.com/Bluemix/power-go-client/power/models"

	"log"
	"time"
)

/*
Transition states

The server can go from

ACTIVE --> SHUTOFF
ACTIVE --> HARD-REBOOT
ACTIVE --> SOFT-REBOOT
SHUTOFF--> ACTIVE




*/

func resourceIBMPIIOperations() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIOperationsCreate,
		Read:     resourceIBMPIOperationsRead,
		Update:   resourceIBMPIOperationsUpdate,
		Delete:   resourceIBMPIOperationsDelete,
		Exists:   resourceIBMPIOperationsExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PIInstanceOperationStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			helpers.PIInstanceOperationServerName: {
				Type:     schema.TypeString,
				Required: true,
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
					},
				},
			},

			helpers.PIInstanceHealthStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			helpers.PIInstanceOperationType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"start", "stop", "hard-reboot", "soft-reboot"}),
			},

			helpers.PIInstanceOperationProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
		},
	}
}

func resourceIBMPIOperationsCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	operation := d.Get(helpers.PIInstanceOperationType).(string)
	name := d.Get(helpers.PIInstanceOperationServerName).(string)

	body := &models.PVMInstanceAction{Action: ptrToString(operation)}
	log.Printf("Calling the IBM PI Operations with the following attributes %s - %s", powerinstanceid, name)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	/*
		TODO
		To add a check if the action performed is applicable on the current state of the instance
	*/

	pvmoperation, err := client.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{
		Body: body,
	}, name, powerinstanceid)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("Printing the instance info %+v", &pvmoperation)

	if operation == "stop" {
		var targetStatus = "SHUTOFF"
		log.Printf("Calling the check status operation to check for status %s", targetStatus)
		_, err = isWaitForPIInstanceOperationStatus(client, name, d.Timeout(schema.TimeoutCreate), powerinstanceid, operation, targetStatus)
		if err != nil {
			return err
		}

	}

	if operation == "start" || operation == "soft-reboot" || operation == "hard-reboot" {
		var targetStatus = "ACTIVE"
		log.Printf("Calling the check status operation to check for status %s", targetStatus)
		_, err = isWaitForPIInstanceOperationStatus(client, name, d.Timeout(schema.TimeoutCreate), powerinstanceid, operation, targetStatus)
		if err != nil {
			return err
		}

	}

	return resourceIBMPIOperationsRead(d, meta)
}

func resourceIBMPIOperationsRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the PowerOperations Read code..for instance id %s", d.Get(helpers.PICloudInstanceId).(string))

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(d.Id(), powerinstanceid)

	if err != nil {
		return err
	}

	d.Set("status", powervmdata.Status)
	d.Set("progress", powervmdata.Progress)

	if powervmdata.Health != nil {
		d.Set("healthstatus", powervmdata.Health.Status)

	}

	return nil

}

func resourceIBMPIOperationsUpdate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceIBMPIOperationsDelete(data *schema.ResourceData, meta interface{}) error {

	return nil
}

// Exists

func resourceIBMPIOperationsExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	instance, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return instance.PvmInstanceID == &id, nil
}

func isWaitForPIInstanceOperationStatus(client *st.IBMPIInstanceClient, name string, timeout time.Duration, powerinstanceid, operation, targetstatus string) (interface{}, error) {

	log.Printf("Waiting for the Operation( %s) to be performed on the instance with name (%s)", name, operation)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "SHUTOFF", helpers.PIInstanceHealthWarning},
		Target:     []string{targetstatus},
		Refresh:    isPIOperationsRefreshFunc(client, name, powerinstanceid, targetstatus),
		Delay:      1 * time.Minute,
		MinTimeout: 30 * time.Second,
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()

}

func isPIOperationsRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid, targetstatus string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		log.Printf("Waiting for the target status to be %s", targetstatus)
		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if *pvm.Status == targetstatus {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			return pvm, targetstatus, nil
			//}
		}

		return pvm, helpers.PIInstanceHealthWarning, nil
	}
}
