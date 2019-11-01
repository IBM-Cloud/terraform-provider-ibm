package ibm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMFunctionAction() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMFunctionActionRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of action.",
			},
			"limits": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout LIMIT in milliseconds after which the action is terminated.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum memory LIMIT in MB for the action (default 256.",
						},
						"log_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum log size LIMIT in MB for the action.",
						},
					},
				},
			},
			"exec": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Container image name when kind is 'blackbox'.",
						},
						"init": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional zipfile reference.",
						},
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The code to execute when kind is not 'blackbox'.",
						},
						"kind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of action. Possible values:php:7.1, nodejs:8, swift:3, nodejs, blackbox, java, sequence, nodejs:6, python:3, python, python:2, swift, swift:3.1.1.",
						},
						"main": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the action entry point (function or fully-qualified method name when applicable)",
						},
						"components": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The List of fully qualified action",
						},
					},
				},
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Action visibilty.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item.",
			},
			"annotations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All annotations set on action by user and those set by the IBM Cloud Function backend/API.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All paramters set on action by user and those set by the IBM Cloud Function backend/API.",
			},
		},
	}
}

func dataSourceIBMFunctionActionRead(d *schema.ResourceData, meta interface{}) error {

	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}

	actionService := wskClient.Actions
	name := d.Get("name").(string)

	action, _, err := actionService.Get(name)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Function Action %s : %s", name, err)
	}

	temp := strings.Split(action.Namespace, "/")

	if len(temp) == 2 {
		d.SetId(fmt.Sprintf("%s/%s", temp[1], action.Name))
		d.Set("name", fmt.Sprintf("%s/%s", temp[1], action.Name))
	} else {
		d.SetId(action.Name)
		d.Set("name", action.Name)
	}
	d.Set("limits", flattenLimits(action.Limits))
	d.Set("exec", flattenExec(action.Exec))
	d.Set("publish", action.Publish)
	d.Set("version", action.Version)
	annotations, err := flattenAnnotations(action.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(action.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)
	return nil
}
