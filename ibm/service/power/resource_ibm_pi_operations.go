// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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

func ResourceIBMPIIOperations() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPIOperationsCreate,
		Read:   resourceIBMPIOperationsRead,
		Update: resourceIBMPIOperationsUpdate,
		Delete: resourceIBMPIOperationsDelete,
		//Exists:   resourceIBMPIOperationsExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI Cloud instnce id",
			},

			helpers.PIInstanceOperationStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI instance operation status",
			},
			helpers.PIInstanceOperationServerName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI instance Operation server name",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI instance health status",
			},

			helpers.PIInstanceOperationType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"start", "stop", "hard-reboot", "soft-reboot", "immediate-shutdown"}),
				Description:  "PI instance operation type",
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

	log.Printf("Now in the Power Operations Code")
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	operation := d.Get(helpers.PIInstanceOperationType).(string)
	name := d.Get(helpers.PIInstanceOperationServerName).(string)

	body := &models.PVMInstanceAction{Action: flex.PtrToString(operation)}
	log.Printf("Calling the IBM PI Operations [ %s ] with on the instance with name [ %s ]", operation, name)
	client := st.NewIBMPIInstanceClient(context.Background(), sess, powerinstanceid)

	/*
		TODO
		To add a check if the action performed is applicable on the current state of the instance
	*/

	err = client.Action(name, body)
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return fmt.Errorf("[ERROR] Failed to perform the operation on  the instance %v", err)

	} else {
		log.Printf("Executed the stop operation on the lpar")
	}

	if operation == "stop" || operation == "immediate-shutdown" {
		var targetStatus = "SHUTOFF"
		log.Printf("Calling the check opertion that was invoked [%s]  to check for status [ %s ]", operation, targetStatus)
		_, err = isWaitForPIInstanceOperationStatus(client, name, d.Timeout(schema.TimeoutCreate), powerinstanceid, operation, targetStatus)
		if err != nil {
			return err
		} else {
			log.Printf("Executed the start operation on the lpar")
		}

	}

	if operation == "start" || operation == "soft-reboot" || operation == "hard-reboot" {
		var targetStatus = "ACTIVE"
		log.Printf("Calling the check opertion that was invoked [%s]  to check for status [ %s ]", operation, targetStatus)
		_, err = isWaitForPIInstanceOperationStatus(client, name, d.Timeout(schema.TimeoutCreate), powerinstanceid, operation, targetStatus)
		if err != nil {
			return err
		}

	}

	return resourceIBMPIOperationsRead(d, meta)
}

func resourceIBMPIOperationsRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the PowerOperations Read code..for instance name %s", d.Get(helpers.PIInstanceOperationServerName).(string))

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIInstanceOperationServerName).(string)
	powerC := st.NewIBMPIInstanceClient(context.Background(), sess, powerinstanceid)
	powervmdata, err := powerC.Get(name)

	if err != nil {
		return err
	}

	d.Set("status", powervmdata.Status)
	d.Set("progress", powervmdata.Progress)

	if powervmdata.Health != nil {
		d.Set("healthstatus", powervmdata.Health.Status)

	}

	pvminstanceid := *powervmdata.PvmInstanceID
	d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, pvminstanceid))

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

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIInstanceClient(context.Background(), sess, powerinstanceid)

	instance, err := client.Get(d.Id())
	if err != nil {

		return false, err
	}
	return instance.PvmInstanceID == &id, nil
}

func isWaitForPIInstanceOperationStatus(client *st.IBMPIInstanceClient, name string, timeout time.Duration, powerinstanceid, operation, targetstatus string) (interface{}, error) {

	log.Printf("Waiting for the Operation [ %s ] to be performed on the instance with name [ %s ]", operation, name)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "SHUTOFF", "WARNING"},
		Target:     []string{targetstatus},
		Refresh:    isPIOperationsRefreshFunc(client, name, powerinstanceid, targetstatus),
		Delay:      1 * time.Minute,
		MinTimeout: 2 * time.Minute,
		Timeout:    120 * time.Minute,
	}

	return stateConf.WaitForState()

}

func isPIOperationsRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid, targetstatus string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		log.Printf("Waiting for the target status to be [ %s ]", targetstatus)
		pvm, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if *pvm.Status == targetstatus && pvm.Health.Status == helpers.PIInstanceHealthOk {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			return pvm, targetstatus, nil
			//}
		}

		return pvm, helpers.PIInstanceHealthWarning, nil
	}
}
