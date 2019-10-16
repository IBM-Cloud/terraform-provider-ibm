package ibm

import (
	"fmt"
	"log"
	"net/http"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const feedLifeCycleEvent = "lifecycleEvent"
const feedTriggerName = "triggerName"
const feedAuthKey = "authKey"
const feedCreate = "CREATE"
const feedDelete = "DELETE"

func resourceIBMFunctionTrigger() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMFunctionTriggerCreate,
		Read:     resourceIBMFunctionTriggerRead,
		Update:   resourceIBMFunctionTriggerUpdate,
		Delete:   resourceIBMFunctionTriggerDelete,
		Exists:   resourceIBMFunctionTriggerExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of Trigger.",
				ValidateFunc: validateFunctionName,
			},
			"feed": {
				Type:        schema.TypeList,
				ForceNew:    true,
				Optional:    true,
				MaxItems:    1,
				Description: "Trigger feed",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Trigger feed ACTION_NAME.",
						},
						"parameters": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "[]",
							Description:  "Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the action invoke.",
							ValidateFunc: validateJSONString,
							DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
								if o == "" && n == "" {
									return false
								}
								if o == "[]" {
									return true
								}
								return false
							},
							StateFunc: func(v interface{}) string {
								json, _ := normalizeJSONString(v)
								return json
							},
						},
					},
				},
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Trigger visbility.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item.",
			},
			"user_defined_annotations": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Annotation values in KEY VALUE format.",
				Default:          "[]",
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
				Description:      "Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the trigger.",
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
				Description: "All annotations set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All parameters set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
		},
	}
}

func resourceIBMFunctionTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers
	feed := false
	feedPayload := map[string]interface{}{}
	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}

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

	if v, ok := d.GetOk("feed"); ok {
		feed = true
		value := v.([]interface{})[0].(map[string]interface{})
		feedPaylod := whisk.KeyValue{
			Key:   "feed",
			Value: value["name"],
		}
		feedArray := make([]whisk.KeyValue, 0, 1)
		feedArray = append(feedArray, feedPaylod)
		payload.Annotations = payload.Annotations.AppendKeyValueArr(feedArray)
	}

	log.Println("[INFO] Creating IBM Cloud Function trigger")
	result, _, err := triggerService.Insert(&payload, false)
	if err != nil {
		return fmt.Errorf("Error creating IBM Cloud Function trigger: %s", err)
	}

	d.SetId(result.Name)

	if feed {
		feed := d.Get("feed").([]interface{})[0].(map[string]interface{})
		actionName := feed["name"].(string)
		parameters := feed["parameters"].(string)
		var err error
		feedParameters, err := expandParameters(parameters)
		if err != nil {
			return err
		}
		for _, value := range feedParameters {
			feedPayload[value.Key] = value.Value
		}
		var feedQualifiedName = new(QualifiedName)

		if feedQualifiedName, err = NewQualifiedName(actionName); err != nil {
			_, _, delerr := triggerService.Delete(name)
			if delerr != nil {
				return fmt.Errorf("Error creating IBM Cloud Function trigger with feed: %s", err)
			}
			return NewQualifiedNameError(actionName, err)
		}

		feedPayload[feedLifeCycleEvent] = feedCreate
		feedPayload[feedAuthKey] = wskClient.Config.AuthToken
		feedPayload[feedTriggerName] = fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), name)

		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: feedQualifiedName.GetNamespace(),
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		actionService := c.Actions
		_, _, err = actionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, true)
		if err != nil {
			_, _, delerr := triggerService.Delete(name)
			if delerr != nil {
				return fmt.Errorf("Error creating IBM Cloud Function trigger with feed: %s", err)
			}
			d.SetId("")
			return fmt.Errorf("Error creating IBM Cloud Function trigger with feed: %s", err)
		}
	}

	d.SetId(result.Name)

	return resourceIBMFunctionTriggerRead(d, meta)
}

func resourceIBMFunctionTriggerRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers
	id := d.Id()

	trigger, _, err := triggerService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Function Trigger %s : %s", id, err)
	}

	d.SetId(trigger.Name)
	d.Set("name", trigger.Name)
	d.Set("publish", trigger.Publish)
	d.Set("version", trigger.Version)
	annotations, err := flattenAnnotations(trigger.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(trigger.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)
	d.Set("user_defined_parameters", parameters)

	userDefinedAnnotations, err := filterTriggerAnnotations(trigger.Annotations)
	if err != nil {
		return err
	}
	d.Set("user_defined_annotations", userDefinedAnnotations)

	found := trigger.Annotations.FindKeyValue("feed")

	if found >= 0 {
		d.Set("feed", flattenFeed(trigger.Annotations.GetValue("feed").(string)))
	}

	return nil
}

func resourceIBMFunctionTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers

	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false

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

	if ischanged {
		log.Println("[INFO] Update IBM Cloud Function Trigger")

		_, _, err = triggerService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating IBM Cloud Function Trigger: %s", err)
		}
	}

	return resourceIBMFunctionTriggerRead(d, meta)
}

func resourceIBMFunctionTriggerDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers
	id := d.Id()

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}
	trigger, _, err := triggerService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Function Trigger %s : %s", id, err)
	}
	found := trigger.Annotations.FindKeyValue("feed")
	if found >= 0 {
		actionName := trigger.Annotations.GetValue("feed").(string)
		var feedQualifiedName = new(QualifiedName)

		if feedQualifiedName, err = NewQualifiedName(actionName); err != nil {
			return NewQualifiedNameError(actionName, err)
		}

		feedPayload := map[string]interface{}{
			feedLifeCycleEvent: feedDelete,
			feedAuthKey:        wskClient.Config.AuthToken,
			feedTriggerName:    fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), id),
		}
		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: feedQualifiedName.GetNamespace(),
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		actionService := c.Actions
		_, _, err = actionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, true)
		if err != nil {
			return fmt.Errorf("Error deleting IBM Cloud Function trigger with feed: %s", err)
		}
		wskClient.Namespace = qualifiedName.GetNamespace()
	}

	_, _, err = triggerService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting IBM Cloud Function Trigger: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMFunctionTriggerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return false, err
	}
	triggerService := wskClient.Triggers
	id := d.Id()
	trigger, resp, err := triggerService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with IBM Cloud Function Client : %s", err)
	}
	return trigger.Name == id, nil
}
