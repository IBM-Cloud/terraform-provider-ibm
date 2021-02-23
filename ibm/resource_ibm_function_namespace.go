/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/namespace-go-sdk/ibmcloudfunctionsnamespaceapiv1"
)

const (
	funcNamespaceName      = "name"
	funcNamespaceResGrpId  = "resource_group_id"
	funcNamespaceResPlanId = "resource_plan_id"
	funcNamespaceDesc      = "description"
	funcNamespaceLoc       = "location"
)

func resourceIBMFunctionNamespace() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMFunctionNamespaceCreate,
		Read:     resourceIBMFunctionNamespaceRead,
		Update:   resourceIBMFunctionNamespaceUpdate,
		Delete:   resourceIBMFunctionNamespaceDelete,
		Exists:   resourceIBMFunctionNamespaceExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			funcNamespaceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of namespace.",
				ValidateFunc: InvokeValidator("ibm_function_namespace", funcNamespaceName),
			},
			funcNamespaceDesc: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Namespace Description.",
			},
			funcNamespaceResGrpId: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Resource Group ID.",
				ValidateFunc: InvokeValidator("ibm_function_namespace", funcNamespaceResGrpId),
			},
			funcNamespaceLoc: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Namespace Location.",
			},
		},
	}
}

func resourceIBMFuncNamespaceValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 funcNamespaceName,
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString,
			Required:                   true})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 funcNamespaceResGrpId,
			ValidateFunctionIdentifier: ValidateNoZeroValues,
			Type:                       TypeString,
			Required:                   true})

	ibmFuncNamespaceResourceValidator := ResourceValidator{ResourceName: "ibm_function_namespace", Schema: validateSchema}
	return &ibmFuncNamespaceResourceValidator
}

func resourceIBMFunctionNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	createNamespaceOptions := &ibmcloudfunctionsnamespaceapiv1.CreateNamespaceOptions{}

	name := d.Get(funcNamespaceName).(string)
	createNamespaceOptions.Name = &name
	resource_group_id := d.Get(funcNamespaceResGrpId).(string)
	createNamespaceOptions.ResourceGroupID = &resource_group_id
	resource_plan_id := "functions-base-plan"
	createNamespaceOptions.ResourcePlanID = &resource_plan_id

	if _, ok := d.GetOk(funcNamespaceDesc); ok {
		description := d.Get(funcNamespaceDesc).(string)
		createNamespaceOptions.Description = &description
	}

	namespace, response, err := nsClient.CreateNamespace(createNamespaceOptions)
	if err != nil {
		return fmt.Errorf("Error Creating Namespace: %s\n%s", err, response)
	}

	d.SetId(*namespace.ID)
	log.Printf("[INFO] Created namespace (IAM) : %s", *namespace.Name)

	return resourceIBMFunctionNamespaceRead(d, meta)
}

func resourceIBMFunctionNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	ID := d.Id()

	getOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespaceOptions{
		ID: &ID,
	}
	instance, response, err := nsClient.GetNamespace(getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
	}

	if instance.ID != nil {
		d.SetId(*instance.ID)
	}
	if instance.Name != nil {
		d.Set(funcNamespaceName, *instance.Name)
	}

	if instance.ResourceGroupID != nil {
		d.Set(funcNamespaceResGrpId, *instance.ResourceGroupID)
	}

	if instance.Location != nil {
		d.Set(funcNamespaceLoc, *instance.Location)
	}
	if instance.Description != nil {
		d.Set(funcNamespaceDesc, *instance.Description)
	}

	return nil
}

func resourceIBMFunctionNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	ID := d.Id()

	updateNamespaceOptions := &ibmcloudfunctionsnamespaceapiv1.UpdateNamespaceOptions{}

	if d.HasChange(funcNamespaceName) {
		name := d.Get(funcNamespaceName).(string)
		updateNamespaceOptions.Name = &name
	}

	if d.HasChange(funcNamespaceDesc) {
		description := d.Get(funcNamespaceDesc).(string)
		updateNamespaceOptions.Description = &description
	}

	updateNamespaceOptions.ID = &ID
	namespace, response, err := nsClient.UpdateNamespace(updateNamespaceOptions)
	if err != nil {
		return fmt.Errorf("Error Updating Namespace: %s\n%s", err, response)
	}

	log.Printf("[INFO] Updated namespace (IAM) : %s", *namespace.Name)

	return resourceIBMFunctionNamespaceRead(d, meta)
}

func resourceIBMFunctionNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	ID := d.Id()

	delOptions := &ibmcloudfunctionsnamespaceapiv1.DeleteNamespaceOptions{
		ID: &ID,
	}
	response, err := nsClient.DeleteNamespace(delOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Deleting Namespace: %s\n%s", err, response)
	}

	d.SetId("")
	return nil
}

func resourceIBMFunctionNamespaceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return false, err
	}

	ID := d.Id()

	getOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespaceOptions{
		ID: &ID,
	}
	_, response, err := nsClient.GetNamespace(getOptions)
	if err != nil && response.StatusCode == 404 {
		d.SetId("")
		return false, fmt.Errorf("Error Getting Namesapce (IAM): %s\n%s", err, response)
	}

	return true, nil

}
