// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsPrivatePathServiceGatewayAccountPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsPrivatePathServiceGatewayAccountPolicyCreate,
		ReadContext:   resourceIBMIsPrivatePathServiceGatewayAccountPolicyRead,
		UpdateContext: resourceIBMIsPrivatePathServiceGatewayAccountPolicyUpdate,
		DeleteContext: resourceIBMIsPrivatePathServiceGatewayAccountPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"private_path_service_gateway": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private path service gateway identifier.",
			},
			"access_policy": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_private_path_service_gateway_account_policy", "access_policy"),
				Description:  "The access policy for the account:- permit: access will be permitted- deny:  access will be denied- review: access will be manually reviewed.",
			},
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The account for this access policy.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the account policy was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this account policy.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the account policy was updated.",
			},
			"account_policy": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this account policy.",
			},
		},
	}
}

func ResourceIBMIsPrivatePathServiceGatewayAccountPolicyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "access_policy",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "deny, permit, review",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_private_path_service_gateway_account_policy", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsPrivatePathServiceGatewayAccountPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_private_path_service_gateway_account_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createPrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.CreatePrivatePathServiceGatewayAccountPolicyOptions{}

	createPrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(d.Get("private_path_service_gateway").(string))
	createPrivatePathServiceGatewayAccountPolicyOptions.SetAccessPolicy(d.Get("access_policy").(string))
	accountId := d.Get("account").(string)
	account := &vpcv1.AccountIdentity{
		ID: &accountId,
	}
	createPrivatePathServiceGatewayAccountPolicyOptions.SetAccount(account)

	privatePathServiceGatewayAccountPolicy, response, err := vpcClient.CreatePrivatePathServiceGatewayAccountPolicyWithContext(context, createPrivatePathServiceGatewayAccountPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error creating PPSG account policy failed: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_account_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createPrivatePathServiceGatewayAccountPolicyOptions.PrivatePathServiceGatewayID, *privatePathServiceGatewayAccountPolicy.ID))

	return resourceIBMIsPrivatePathServiceGatewayAccountPolicyRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayAccountPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_private_path_service_gateway_account_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.GetPrivatePathServiceGatewayAccountPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getPrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(parts[0])
	getPrivatePathServiceGatewayAccountPolicyOptions.SetID(parts[1])

	privatePathServiceGatewayAccountPolicy, response, err := vpcClient.GetPrivatePathServiceGatewayAccountPolicyWithContext(context, getPrivatePathServiceGatewayAccountPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error getting PPSG account policy failed: %s", err.Error()), "ibm_is_private_path_service_gateway_account_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("access_policy", privatePathServiceGatewayAccountPolicy.AccessPolicy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_policy: %s", err), "ibm_is_private_path_service_gateway_account_policy", "read", "set-access_policy").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(privatePathServiceGatewayAccountPolicy.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "ibm_is_private_path_service_gateway_account_policy", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", privatePathServiceGatewayAccountPolicy.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "ibm_is_private_path_service_gateway_account_policy", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", privatePathServiceGatewayAccountPolicy.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "ibm_is_private_path_service_gateway_account_policy", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("account_policy", privatePathServiceGatewayAccountPolicy.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_policy: %s", err), "ibm_is_private_path_service_gateway_account_policy", "read", "set-account_policy").GetDiag()
	}

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayAccountPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_private_path_service_gateway_account_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updatePrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.UpdatePrivatePathServiceGatewayAccountPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updatePrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(parts[0])
	updatePrivatePathServiceGatewayAccountPolicyOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.PrivatePathServiceGatewayAccountPolicyPatch{}

	if d.HasChange("access_policy") {
		newAccessPolicy := d.Get("access_policy").(string)
		patchVals.AccessPolicy = &newAccessPolicy
		hasChange = true
	}

	if hasChange {
		updatePrivatePathServiceGatewayAccountPolicyOptions.PrivatePathServiceGatewayAccountPolicyPatch, _ = patchVals.AsPatch()
		if err != nil {
			return flex.TerraformErrorf(err, fmt.Sprintf("Error calling AsPatch for PrivatePathServiceGatewayAccountPolicyPatch %s", err.Error()), "ibm_is_private_path_service_gateway_account_policy", "update").GetDiag()
		}
		_, response, err := vpcClient.UpdatePrivatePathServiceGatewayAccountPolicyWithContext(context, updatePrivatePathServiceGatewayAccountPolicyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error updating PPSG account policy: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_account_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsPrivatePathServiceGatewayAccountPolicyRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayAccountPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_private_path_service_gateway_account_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePrivatePathServiceGatewayAccountPolicyOptions := &vpcv1.DeletePrivatePathServiceGatewayAccountPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deletePrivatePathServiceGatewayAccountPolicyOptions.SetPrivatePathServiceGatewayID(parts[0])
	deletePrivatePathServiceGatewayAccountPolicyOptions.SetID(parts[1])

	response, err := vpcClient.DeletePrivatePathServiceGatewayAccountPolicyWithContext(context, deletePrivatePathServiceGatewayAccountPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error deleting PPSG account policy: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_account_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
