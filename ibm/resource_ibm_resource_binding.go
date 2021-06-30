// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func resourceIbmResourceBinding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmResourceBindingCreate,
		ReadContext:   resourceIbmResourceBindingRead,
		UpdateContext: resourceIbmResourceBindingUpdate,
		DeleteContext: resourceIbmResourceBindingDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The short or long ID of resource alias.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The CRN of application to bind to in a specific environment, for example, Dallas YP, CFEE instance.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_resource_binding", "name"),
				Description:  "The name of the binding. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Configuration options represented as key-value pairs. Service defined options are passed through to the target resource brokers, whereas platform defined options are not.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serviceid_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An optional platform defined option to reuse an existing IAM serviceId for the role assignment.",
						},
					},
				},
			},
			"role": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "Writer",
				Description: "The role name or it's CRN.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you create a new binding, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the binding.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you provision a new binding, a relative URL path is created identifying the location of the binding.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the binding was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the binding was last updated.",
			},
			"deleted_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the binding was deleted.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who created the binding.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who updated the binding.",
			},
			"deleted_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who deleted the binding.",
			},
			"source_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of resource alias associated to the binding.",
			},
			"target_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of target resource, for example, application, in a specific environment.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full Cloud Resource Name (CRN) associated with the binding. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).",
			},
			"region_binding_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the binding in the specific target environment, for example, `service_binding_id` in a given IBM Cloud environment.",
			},
			"region_binding_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the binding in the specific target environment.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An alpha-numeric value identifying the account ID.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the resource group.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the binding.",
			},
			"credentials": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The credentials for the binding. Additional key-value pairs are passed through from the resource brokers.  For additional details, see the service’s documentation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apikey": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The API key for the credentials.",
							Sensitive:   true,
						},
						"iam_apikey_description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The optional description of the API key.",
						},
						"iam_apikey_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the API key.",
						},
						"iam_role_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Cloud Resource Name for the role of the credentials.",
						},
						"iam_serviceid_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Cloud Resource Name for the service ID of the credentials.",
						},
					},
				},
			},
			"iam_compatible": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the binding’s credentials support IAM.",
			},
			"resource_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of the offering. This value is provided by and stored in the global catalog.",
			},
			"migrated": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "A boolean that dictates if the alias was migrated from a previous CF instance.",
			},
			"resource_alias_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The relative path to the resource alias that this binding is associated with.",
			},
		},
	}
}

func resourceIbmResourceBindingValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexp,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$`,
			MaxValueLength:             180,
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_resource_binding", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmResourceBindingCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	createResourceBindingOptions := &resourcecontrollerv2.CreateResourceBindingOptions{}

	createResourceBindingOptions.SetSource(d.Get("source").(string))
	createResourceBindingOptions.SetTarget(d.Get("target").(string))
	if _, ok := d.GetOk("name"); ok {
		createResourceBindingOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		parameters := resourceIbmResourceBindingMapToResourceBindingPostParameters(d.Get("parameters.0").(map[string]interface{}))
		createResourceBindingOptions.SetParameters(&parameters)
	}
	if _, ok := d.GetOk("role"); ok {
		createResourceBindingOptions.SetRole(d.Get("role").(string))
	}

	log.Printf("[DEBUG] Creating Resource Binding : %s", *createResourceBindingOptions.Name)
	resourceBinding, response, err := resourceControllerClient.CreateResourceBindingWithContext(context, createResourceBindingOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateResourceBindingWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateResourceBindingWithContext failed %s\n%s", err, response))
	}

	log.Println("[DEBUG] Resource Binding created successfully")
	d.SetId(*resourceBinding.ID)

	return resourceIbmResourceBindingRead(context, d, meta)
}

func resourceIbmResourceBindingMapToResourceBindingPostParameters(resourceBindingPostParametersMap map[string]interface{}) resourcecontrollerv2.ResourceBindingPostParameters {
	resourceBindingPostParameters := resourcecontrollerv2.ResourceBindingPostParameters{}

	if resourceBindingPostParametersMap["serviceid_crn"] != nil {
		resourceBindingPostParameters.ServiceidCRN = core.StringPtr(resourceBindingPostParametersMap["serviceid_crn"].(string))
	}

	return resourceBindingPostParameters
}

func resourceIbmResourceBindingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	getResourceBindingOptions := &resourcecontrollerv2.GetResourceBindingOptions{}
	getResourceBindingOptions.SetID(d.Id())

	log.Printf("[DEBUG] Reading Resource Binding : %s", *getResourceBindingOptions.ID)
	resourceBinding, response, err := resourceControllerClient.GetResourceBindingWithContext(context, getResourceBindingOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetResourceBindingWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetResourceBindingWithContext failed %s\n%s", err, response))
	}

	if _, ok := d.GetOk("source"); !ok {
		// when resource is imported, source can be extracted from ResourceAliasURL
		resourceAliasURL := *resourceBinding.ResourceAliasURL
		resourceAliasID := resourceAliasURL[strings.LastIndex(resourceAliasURL, "/")+1:]
		if err = d.Set("source", resourceAliasID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting source: %s", err))
		}
	}
	if _, ok := d.GetOk("target"); !ok {
		// when resource is imported, target can be set as TargetCRN
		if err = d.Set("target", resourceBinding.TargetCRN); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting target: %s", err))
		}
	}
	if _, ok := d.GetOk("role"); !ok {
		// when resource is imported, role can be extracted from Credentials.IamRoleCRN
		iamRoleCRN := *resourceBinding.Credentials.IamRoleCRN
		roleName := iamRoleCRN[strings.LastIndex(iamRoleCRN, ":")+1:]
		if err = d.Set("role", roleName); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting role: %s", err))
		}
	}
	if err = d.Set("name", resourceBinding.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("guid", resourceBinding.GUID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting guid: %s", err))
	}
	if err = d.Set("url", resourceBinding.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}
	if err = d.Set("created_at", dateTimeToString(resourceBinding.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", dateTimeToString(resourceBinding.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("deleted_at", dateTimeToString(resourceBinding.DeletedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting deleted_at: %s", err))
	}
	if err = d.Set("created_by", resourceBinding.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("updated_by", resourceBinding.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}
	if err = d.Set("deleted_by", resourceBinding.DeletedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting deleted_by: %s", err))
	}
	if err = d.Set("source_crn", resourceBinding.SourceCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_crn: %s", err))
	}
	if err = d.Set("target_crn", resourceBinding.TargetCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_crn: %s", err))
	}
	if err = d.Set("crn", resourceBinding.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("region_binding_id", resourceBinding.RegionBindingID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_binding_id: %s", err))
	}
	if err = d.Set("region_binding_crn", resourceBinding.RegionBindingCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_binding_crn: %s", err))
	}
	if err = d.Set("account_id", resourceBinding.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("resource_group_id", resourceBinding.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("state", resourceBinding.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}
	if resourceBinding.Credentials != nil {
		credentialsMap := resourceIbmResourceBindingCredentialsToMap(*resourceBinding.Credentials)
		if err = d.Set("credentials", []map[string]interface{}{credentialsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting credentials: %s", err))
		}
	}
	if err = d.Set("iam_compatible", resourceBinding.IamCompatible); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting iam_compatible: %s", err))
	}
	if err = d.Set("resource_id", resourceBinding.ResourceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_id: %s", err))
	}
	if err = d.Set("migrated", resourceBinding.Migrated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting migrated: %s", err))
	}
	if err = d.Set("resource_alias_url", resourceBinding.ResourceAliasURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_alias_url: %s", err))
	}

	log.Println("[DEBUG] Resource Binding read successfully")
	return nil
}

func resourceIbmResourceBindingCredentialsToMap(credentials resourcecontrollerv2.Credentials) map[string]interface{} {
	credentialsMap := map[string]interface{}{}

	if credentials.Apikey != nil {
		credentialsMap["apikey"] = credentials.Apikey
	}
	if credentials.IamApikeyDescription != nil {
		credentialsMap["iam_apikey_description"] = credentials.IamApikeyDescription
	}
	if credentials.IamApikeyName != nil {
		credentialsMap["iam_apikey_name"] = credentials.IamApikeyName
	}
	if credentials.IamRoleCRN != nil {
		credentialsMap["iam_role_crn"] = credentials.IamRoleCRN
	}
	if credentials.IamServiceidCRN != nil {
		credentialsMap["iam_serviceid_crn"] = credentials.IamServiceidCRN
	}

	return credentialsMap
}

func resourceIbmResourceBindingUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	updateResourceBindingOptions := &resourcecontrollerv2.UpdateResourceBindingOptions{}
	updateResourceBindingOptions.SetID(d.Id())

	hasChange := false
	if d.HasChange("name") {
		updateResourceBindingOptions.SetName(d.Get("name").(string))
		hasChange = true
	}

	if hasChange {
		log.Printf("[DEBUG] Updating Resource Binding name : %s", *updateResourceBindingOptions.Name)
		_, response, err := resourceControllerClient.UpdateResourceBindingWithContext(context, updateResourceBindingOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateResourceBindingWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateResourceBindingWithContext failed %s\n%s", err, response))
		}
		log.Println("[DEBUG] Resource Binding updated successfully")
	}

	return resourceIbmResourceBindingRead(context, d, meta)
}

func resourceIbmResourceBindingDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteResourceBindingOptions := &resourcecontrollerv2.DeleteResourceBindingOptions{}
	deleteResourceBindingOptions.SetID(d.Id())

	log.Printf("[DEBUG] Deleting Resource Binding : %s", *deleteResourceBindingOptions.ID)
	response, err := resourceControllerClient.DeleteResourceBindingWithContext(context, deleteResourceBindingOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteResourceBindingWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteResourceBindingWithContext failed %s\n%s", err, response))
	}

	log.Println("[DEBUG] Resource Binding deleted successfully")
	d.SetId("")

	return nil
}
