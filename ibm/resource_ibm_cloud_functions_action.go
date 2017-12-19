package ibm

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMCloudFunctionsAction() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCloudFunctionsActionCreate,
		Read:     resourceIBMCloudFunctionsActionRead,
		Update:   resourceIBMCloudFunctionsActionUpdate,
		Delete:   resourceIBMCloudFunctionsActionDelete,
		Exists:   resourceIBMCloudFunctionsActionExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of action.",
				ValidateFunc: validateActionName,
			},
			"limits": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     60000,
							Description: "The timeout LIMIT in milliseconds after which the action is terminated.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     256,
							Description: "The maximum memory LIMIT in MB for the action (default 256.",
						},
						"log_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     10,
							Description: "The maximum log size LIMIT in MB for the action.",
						},
					},
				},
			},
			"exec": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "Container image name when kind is 'blackbox'.",
							ConflictsWith: []string{"exec.components", "exec.code"},
						},
						"init": {
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "Optional zipfile reference.",
							ConflictsWith: []string{"exec.image", "exec.components"},
						},
						"code": {
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "The code to execute when kind is not ‘blackbox’.",
							ConflictsWith: []string{"exec.image", "exec.components"},
						},
						"kind": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "The type of action. Possible values:php:7.1, nodejs:8, swift:3, nodejs, blackbox, java, sequence, nodejs:6, python:3, python, python:2, swift, swift:3.1.1.",
							ValidateFunc: validateActionKind,
						},
						"main": {
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "The name of the action entry point (function or fully-qualified method name when applicable).",
							ConflictsWith: []string{"exec.image", "exec.components"},
						},
						"components": {
							Type:          schema.TypeList,
							Optional:      true,
							Elem:          &schema.Schema{Type: schema.TypeString},
							Description:   "The List of fully qualified action.",
							ConflictsWith: []string{"exec.image", "exec.code"},
						},
					},
				},
			},
			"publish": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Action visibilty.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item.",
			},
			"user_defined_annotations": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "[]",
				Description:      "Annotation values in KEY VALUE format.",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
			"user_defined_parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "[]",
				Description:      "Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the action.",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
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

func resourceIBMCloudFunctionsActionCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	actionService := wskClient.Actions

	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Action{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}

	exec := d.Get("exec").(*schema.Set)
	payload.Exec = expandExec(exec)

	userDefinedAnnotations := d.Get("user_defined_annotations").(string)
	payload.Annotations, err = expandAnnotations(userDefinedAnnotations)
	if err != nil {
		return err
	}

	userDefinedParameters := d.Get("user_defined_parameters").(string)
	payload.Parameters, err = expandParameters(userDefinedParameters)
	if err != nil {
		return err
	}

	if v := d.Get("limits").(*schema.Set); len(v.List()) > 0 {
		payload.Limits = expandLimits(v.List())
	}

	if publish, ok := d.GetOk("publish"); ok {
		p := publish.(bool)
		payload.Publish = &p
	}

	log.Println("[INFO] Creating IBM Cloud Functions Action")
	action, _, err := actionService.Insert(&payload, false)
	if err != nil {
		return fmt.Errorf("Error creating IBM Cloud Functions Action: %s", err)
	}

	temp := strings.Split(action.Namespace, "/")

	if len(temp) == 2 {
		d.SetId(fmt.Sprintf("%s/%s", temp[1], action.Name))
	} else {
		d.SetId(action.Name)
	}

	return resourceIBMCloudFunctionsActionRead(d, meta)
}

func resourceIBMCloudFunctionsActionRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}

	actionService := wskClient.Actions
	id := d.Id()

	action, _, err := actionService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Functions Action %s : %s", id, err)
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

	temp := strings.Split(action.Namespace, "/")

	if len(temp) == 2 {
		d.Set("name", fmt.Sprintf("%s/%s", temp[1], action.Name))
		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: wskClient.Namespace,
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		pkg, _, err := c.Packages.Get(temp[1])

		if err != nil {
			return fmt.Errorf("Error retrieving package IBM Cloud Functions package %s : %s", temp[1], err)
		}

		userAnnotations, err := flattenAnnotations(filterInheritedAnnotations(pkg.Annotations, pkg.Annotations))
		if err != nil {
			return err
		}
		d.Set("user_defined_annotations", userAnnotations)

		userParameters, err := flattenParameters(filterInheritedParameters(pkg.Parameters, pkg.Parameters))
		if err != nil {
			return err
		}
		d.Set("user_defined_parameters", userParameters)
	} else {
		d.Set("name", action.Name)
		userDefinedAnnotations, err := filterActionAnnotations(action.Annotations)
		if err != nil {
			return err
		}
		d.Set("user_defined_annotations", userDefinedAnnotations)

		userDefinedParameters, err := filterActionParameters(action.Parameters)
		if err != nil {
			return err
		}
		d.Set("user_defined_parameters", userDefinedParameters)
	}
	return nil
}

func resourceIBMCloudFunctionsActionUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	actionService := wskClient.Actions

	id := d.Id()

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	payload := whisk.Action{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}

	ischanged := false

	if d.HasChange("publish") {
		p := d.Get("publish").(bool)
		payload.Publish = &p
	}

	if d.HasChange("user_defined_parameters") {
		var err error
		payload.Parameters, err = expandParameters(d.Get("user_defined_parameters").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("user_defined_annotations") {
		var err error
		payload.Annotations, err = expandAnnotations(d.Get("user_defined_annotations").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("exec") {
		exec := d.Get("exec").(*schema.Set)
		payload.Exec = expandExec(exec)
		ischanged = true
	}

	if ischanged {
		log.Println("[INFO] Update IBM Cloud Functions Action")

		_, _, err = actionService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating IBM Cloud Functions Action: %s", err)
		}
	}

	return resourceIBMCloudFunctionsActionRead(d, meta)
}

func resourceIBMCloudFunctionsActionDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	actionService := wskClient.Actions
	id := d.Id()

	_, err = actionService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting IBM Cloud Functions Action: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMCloudFunctionsActionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return false, err
	}
	actionService := wskClient.Actions
	id := d.Id()

	action, resp, err := actionService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with IBM Cloud Functions Client : %s", err)
	}

	temp := strings.Split(action.Namespace, "/")
	var name string

	if len(temp) == 2 {
		name = fmt.Sprintf("%s/%s", temp[1], action.Name)
	} else {
		name = action.Name
	}

	return name == id, nil
}
