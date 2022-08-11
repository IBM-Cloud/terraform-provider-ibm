// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMFunctionTrigger() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMFunctionTriggerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Trigger.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the namespace.",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Trigger Visibility.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the trigger.",
			},

			"annotations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All annotations set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All parameters set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
			"trigger_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMFunctionTriggerRead(d *schema.ResourceData, meta interface{}) error {
	functionNamespaceAPI, err := meta.(conns.ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	bxSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	namespace := d.Get("namespace").(string)
	wskClient, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
	if err != nil {
		return err

	}

	triggerService := wskClient.Triggers
	name := d.Get("name").(string)

	trigger, _, err := triggerService.Get(name)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving IBM Cloud Function Trigger %s : %s", name, err)
	}

	d.SetId(trigger.Name)
	d.Set("name", trigger.Name)
	d.Set("namespace", namespace)
	d.Set("publish", trigger.Publish)
	d.Set("version", trigger.Version)
	d.Set("trigger_id", trigger.Name)
	annotations, err := flex.FlattenAnnotations(trigger.Annotations)
	if err != nil {
		log.Printf(
			"An error occured during reading of trigger (%s) annotations : %s", d.Id(), err)

	}
	d.Set("annotations", annotations)
	parameters, err := flex.FlattenParameters(trigger.Parameters)
	if err != nil {
		log.Printf(
			"An error occured during reading of trigger (%s) parameters : %s", d.Id(), err)
	}
	d.Set("parameters", parameters)

	return nil
}
