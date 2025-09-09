// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package mqcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func ResourceIbmMqcloudUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmMqcloudUserCreate,
		ReadContext:   resourceIbmMqcloudUserRead,
		UpdateContext: resourceIbmMqcloudUserUpdate,
		DeleteContext: resourceIbmMqcloudUserDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_mqcloud_user", "service_instance_guid"),
				Description:  "The GUID that uniquely identifies the MQ SaaS service instance.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_mqcloud_user", "name"),
				Description:  "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.",
			},
			"email": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_mqcloud_user", "email"),
				Description:  "The email of the user.",
			},
			"iam_service_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the user.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the user details.",
			},
			"user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the user which was allocated on creation, and can be used for delete calls.",
			},
		},
	}
}

func ResourceIbmMqcloudUserValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "service_instance_guid",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z][-a-z0-9]*$`,
			MinValueLength:             1,
			MaxValueLength:             12,
		},
		validate.ValidateSchema{
			Identifier:                 "email",
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			Required:                   true,
			MinValueLength:             5,
			MaxValueLength:             253,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_mqcloud_user", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmMqcloudUserCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createUserOptions := &mqcloudv1.CreateUserOptions{}

	createUserOptions.SetServiceInstanceGuid(d.Get("service_instance_guid").(string))
	createUserOptions.SetEmail(d.Get("email").(string))
	createUserOptions.SetName(d.Get("name").(string))

	userDetails, _, err := mqcloudClient.CreateUserWithContext(context, createUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateUserWithContext failed: %s", err.Error()), "ibm_mqcloud_user", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createUserOptions.ServiceInstanceGuid, *userDetails.ID))

	return resourceIbmMqcloudUserRead(context, d, meta)
}

func resourceIbmMqcloudUserRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getUserOptions := &mqcloudv1.GetUserOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "sep-id-parts").GetDiag()
	}

	getUserOptions.SetServiceInstanceGuid(parts[0])
	getUserOptions.SetUserID(parts[1])

	userDetails, response, err := mqcloudClient.GetUserWithContext(context, getUserOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetUserWithContext failed: %s", err.Error()), "ibm_mqcloud_user", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("service_instance_guid", parts[0]); err != nil {
		err = fmt.Errorf("Error setting service_instance_guid: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-service_instance_guid").GetDiag()
	}
	if err = d.Set("name", userDetails.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-name").GetDiag()
	}
	if err = d.Set("email", userDetails.Email); err != nil {
		err = fmt.Errorf("Error setting email: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-email").GetDiag()
	}
	if err = d.Set("iam_service_id", userDetails.IamServiceID); err != nil {
		err = fmt.Errorf("Error setting iam_service_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-iam_service_id").GetDiag()
	}
	if err = d.Set("href", userDetails.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-href").GetDiag()
	}
	if err = d.Set("user_id", userDetails.ID); err != nil {
		err = fmt.Errorf("Error setting user_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "read", "set-user_id").GetDiag()
	}

	return nil
}

func resourceIbmMqcloudUserUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	setUserNameOptions := &mqcloudv1.SetUserNameOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "update", "sep-id-parts").GetDiag()
	}

	setUserNameOptions.SetServiceInstanceGuid(parts[0])
	setUserNameOptions.SetUserID(parts[1])

	hasChange := false

	if d.HasChange("service_instance_guid") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "service_instance_guid")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_mqcloud_user", "update", "service_instance_guid-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		setUserNameOptions.SetName(d.Get("name").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = mqcloudClient.SetUserNameWithContext(context, setUserNameOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetUserNameWithContext failed: %s", err.Error()), "ibm_mqcloud_user", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmMqcloudUserRead(context, d, meta)
}

func resourceIbmMqcloudUserDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteUserOptions := &mqcloudv1.DeleteUserOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_mqcloud_user", "delete", "sep-id-parts").GetDiag()
	}

	deleteUserOptions.SetServiceInstanceGuid(parts[0])
	deleteUserOptions.SetUserID(parts[1])

	_, err = mqcloudClient.DeleteUserWithContext(context, deleteUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteUserWithContext failed: %s", err.Error()), "ibm_mqcloud_user", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
