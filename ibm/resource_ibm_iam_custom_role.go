// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	iamCRDisplayName = "display_name"
	iamCRName        = "name"
	iamCRDescription = "description"
	iamCRActions     = "actions"
	iamCRServiceName = "service"
)

func resourceIBMIAMCustomRole() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMCustomRoleCreate,
		Read:     resourceIBMIAMCustomRoleRead,
		Update:   resourceIBMIAMCustomRoleUpdate,
		Delete:   resourceIBMIAMCustomRoleDelete,
		Exists:   resourceIBMIAMCustomRoleExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			iamCRDisplayName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Display Name of the Custom Role",
				ValidateFunc: InvokeValidator("ibm_iam_custom_role", iamCRDisplayName),
			},

			iamCRName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the custom Role",
				ForceNew:     true,
				ValidateFunc: InvokeValidator("ibm_iam_custom_role", iamCRName),
			},
			iamCRDescription: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The description of the role",
				ValidateFunc: InvokeValidator("ibm_iam_custom_role", iamCRDescription),
			},
			iamCRServiceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Service Name",
				ForceNew:    true,
			},
			iamCRActions: {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The actions of the role",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "crn of the Custom Role",
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
			},
		},
	}
}

func resourceIBMIAMCustomRoleValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 iamCRName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Z]{1}[A-Za-z0-9]{0,29}$`,
			MinValueLength:             1,
			MaxValueLength:             30})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 iamCRDisplayName,
			ValidateFunctionIdentifier: StringLenBetween,
			Type:                       TypeString,
			Optional:                   true,
			MinValueLength:             1,
			MaxValueLength:             50})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 iamCRDescription,
			ValidateFunctionIdentifier: StringLenBetween,
			Type:                       TypeString,
			Optional:                   true,
			MinValueLength:             1,
			MaxValueLength:             250})

	ibmIAMCustomRoleResourceValidator := ResourceValidator{ResourceName: "ibm_iam_custom_role", Schema: validateSchema}
	return &ibmIAMCustomRoleResourceValidator
}

func resourceIBMIAMCustomRoleCreate(d *schema.ResourceData, meta interface{}) error {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}

	displayName := d.Get(iamCRDisplayName).(string)
	name := d.Get(iamCRName).(string)
	description := d.Get(iamCRDescription).(string)
	serviceName := d.Get(iamCRServiceName).(string)
	actionList := expandStringList(d.Get(iamCRActions).([]interface{}))

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	createRoleReq := iampapv2.CreateRoleRequest{
		Name:        name,
		DisplayName: displayName,
		Description: description,
		ServiceName: serviceName,
		Actions:     actionList,
		AccountID:   userDetails.userAccount,
	}

	response, err := iampapv2Client.IAMRoles().Create(createRoleReq)
	if err != nil {
		return fmt.Errorf("Error creating Custom Roles: %s", err)
	}

	d.SetId(response.ID)

	return resourceIBMIAMCustomRoleRead(d, meta)
}

func resourceIBMIAMCustomRoleRead(d *schema.ResourceData, meta interface{}) error {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}

	roleID := d.Id()

	role, _, err := iampapv2Client.IAMRoles().Get(roleID)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return fmt.Errorf("Error retrieving Custom Roles: %s", err)
	} else if err != nil && strings.Contains(err.Error(), "404") {
		d.SetId("")

		return nil
	}

	d.Set(iamCRDisplayName, role.DisplayName)
	d.Set(iamCRName, role.Name)
	d.Set(iamCRDescription, role.Description)
	d.Set(iamCRServiceName, role.ServiceName)
	d.Set(iamCRActions, role.Actions)
	d.Set("crn", role.Crn)

	d.Set(ResourceName, role.Name)
	d.Set(ResourceCRN, role.Crn)
	rcontroller, err := getBaseController(meta)
	if err != nil {
		return err
	}

	d.Set(ResourceControllerURL, rcontroller+"/iam/roles")

	return nil
}

func resourceIBMIAMCustomRoleUpdate(d *schema.ResourceData, meta interface{}) error {

	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}
	roleID := d.Id()

	updateReq := iampapv2.UpdateRoleRequest{
		Description: d.Get(iamCRDescription).(string),
		Actions:     expandStringList(d.Get(iamCRActions).([]interface{})),
		DisplayName: d.Get(iamCRDisplayName).(string),
	}

	if d.HasChange("display_name") || d.HasChange("desciption") || d.HasChange("actions") {
		_, etag, err := iampapv2Client.IAMRoles().Get(roleID)
		if err != nil {
			return fmt.Errorf("Error retrieving Custom Role: %s", err)
		}

		_, err = iampapv2Client.IAMRoles().Update(updateReq, roleID, etag)
		if err != nil {
			return fmt.Errorf("Error updating Custom Role: %s", err)
		}
	}

	return resourceIBMIAMCustomRoleRead(d, meta)
}

func resourceIBMIAMCustomRoleDelete(d *schema.ResourceData, meta interface{}) error {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}

	roleID := d.Id()

	err = iampapv2Client.IAMRoles().Delete(roleID)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return fmt.Errorf("Error deleting Custom Roles: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMCustomRoleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return false, err
	}
	roleID := d.Id()

	role, _, err := iampapv2Client.IAMRoles().Get(roleID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return role.ID == roleID, nil
}
