// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	registryv1 "github.com/IBM-Cloud/bluemix-go/api/container/registryv1"
)

func resourceIBMContainerRegistryNamespace() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerRegistryNamespaceCreate,
		Read:     resourceIBMContainerRegistryNamespaceRead,
		Delete:   resourceIBMContainerRegistryNamespaceDelete,
		Exists:   resourceIBMContainerRegistryNamespaceExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Container Registry Namespace",
				ValidateFunc: InvokeValidator("ibm_cr_namespace", "name"),
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Resource Group to which namespace has to be assigned",
				ForceNew:    true,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of the Namespace",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created Date",
			},
			"updated_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Updated Date",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},
		},
	}
}
func resourceIBMCrNamespaceValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[a-z0-9]+[a-z0-9\-\_]+[a-z0-9]+$`,
			MinValueLength:             4,
			MaxValueLength:             30})

	ibmCrNamespaceResourceValidator := ResourceValidator{ResourceName: "ibm_cr_namespace", Schema: validateSchema}
	return &ibmCrNamespaceResourceValidator
}
func resourceIBMContainerRegistryNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	log.Print("accountID", accountID)

	crClient, err := meta.(ClientSession).ContainerRegistryAPI()
	if err != nil {
		return err
	}
	crAPI := crClient.Namespaces()
	namespace := d.Get("name").(string)
	target := registryv1.NamespaceTargetHeader{
		AccountID: accountID,
	}
	if rg, ok := d.GetOk("resource_group_id"); ok {
		target.ResourceGroup = rg.(string)
	} else {
		defaultRg, err := defaultResourceGroup(meta)
		if err != nil {
			return err
		}
		target.ResourceGroup = defaultRg
	}

	response, err := crAPI.AddNamespace(namespace, target)
	if err != nil {
		return err
	}
	d.SetId(response.Namespace)
	return resourceIBMContainerRegistryNamespaceRead(d, meta)
}
func resourceIBMContainerRegistryNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount

	crClient, err := meta.(ClientSession).ContainerRegistryAPI()
	if err != nil {
		return err
	}
	crAPI := crClient.Namespaces()
	namespace := d.Id()
	target := registryv1.NamespaceTargetHeader{
		AccountID: accountID,
	}

	err = crAPI.DeleteNamespace(namespace, target)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return err
	}
	return nil
}

func resourceIBMContainerRegistryNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount

	crClient, err := meta.(ClientSession).ContainerRegistryAPI()
	if err != nil {
		return err
	}
	namespace := d.Id()
	target := registryv1.NamespaceTargetHeader{
		AccountID: accountID,
	}

	crAPI := crClient.Namespaces()

	response, err := crAPI.GetDetailedNamespaces(target)
	if err != nil {
		return err
	}
	found := false
	for _, ns := range response {
		if ns.Name == namespace {
			found = true
			d.Set("name", ns.Name)
			d.Set("resource_group_id", ns.ResourceGroup)
			d.Set("crn", ns.CRN)
			d.Set("created_on", ns.CreatedDate)
			d.Set("updated_on", ns.UpdatedDate)
			d.Set(ResourceControllerURL, ns.CRN)
			d.Set(ResourceName, ns.Name)

		}
	}
	if !found {
		d.SetId("")
	}

	return nil
}

func resourceIBMContainerRegistryNamespaceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	accountID := userDetails.userAccount

	crClient, err := meta.(ClientSession).ContainerRegistryAPI()
	if err != nil {
		return false, err
	}
	namespace := d.Id()
	target := registryv1.NamespaceTargetHeader{
		AccountID: accountID,
	}

	crAPI := crClient.Namespaces()

	response, err := crAPI.GetDetailedNamespaces(target)
	if err != nil {
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	found := false
	for _, ns := range response {
		if ns.Name == namespace {
			found = true
		}
	}
	return found, nil
}
