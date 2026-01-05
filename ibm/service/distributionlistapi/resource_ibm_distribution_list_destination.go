// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package distributionlistapi

import (
	"context"
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
)

func ResourceIbmDistributionListDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmDistributionListDestinationCreate,
		ReadContext:   resourceIbmDistributionListDestinationRead,
		DeleteContext: resourceIbmDistributionListDestinationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "account_id"),
				Description:  "The IBM Cloud account ID.",
			},
			"destination_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "destination_type"),
				Description:  "The type of the destination.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GUID of the Event Notifications instance.",
			},
		},
	}
}

func ResourceIbmDistributionListDestinationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "account_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-zA-Z]{32}$`,
			MinValueLength:             32,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "event_notifications",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_distribution_list_destination", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmDistributionListDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListApiClient, err := meta.(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createDistributionListDestinationOptions := &distributionlistapiv1.CreateDistributionListDestinationOptions{}

	if _, ok := d.GetOk("id"); ok {
		bodyModelMap["id"] = d.Get("id")
	}
	createDistributionListDestinationOptions.SetAccountID(d.Get("account_id").(string))
	convertedModel, err := ResourceIbmDistributionListDestinationMapToAddDestinationResponseBodyPrototypeEventNotificationDestination(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "create", "parse-request-body").GetDiag()
	}
	createDistributionListDestinationOptions.AddDestinationResponseBodyPrototype = convertedModel

	addDestinationResponseBodyIntf, _, err := distributionListApiClient.CreateDistributionListDestinationWithContext(context, createDistributionListDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	addDestinationResponseBody := addDestinationResponseBodyIntf.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination)
	d.SetId(fmt.Sprintf("%s/%s", *createDistributionListDestinationOptions.AccountID, *addDestinationResponseBody.ID))

	return resourceIbmDistributionListDestinationRead(context, d, meta)
}

func resourceIbmDistributionListDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListApiClient, err := meta.(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDistributionListDestinationOptions := &distributionlistapiv1.GetDistributionListDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "sep-id-parts").GetDiag()
	}

	getDistributionListDestinationOptions.SetAccountID(parts[0])
	getDistributionListDestinationOptions.SetDestinationID(core.UUIDPtr(strfmt.UUID(parts[1])))

	addDestinationResponseBodyIntf, response, err := distributionListApiClient.GetDistributionListDestinationWithContext(context, getDistributionListDestinationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	addDestinationResponseBody := addDestinationResponseBodyIntf.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination)
	if err = d.Set("destination_type", addDestinationResponseBody.DestinationType); err != nil {
		err = fmt.Errorf("Error setting destination_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_type").GetDiag()
	}
	if err = d.Set("destination_id", addDestinationResponseBody.ID); err != nil {
		err = fmt.Errorf("Error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_id").GetDiag()
	}

	return nil
}

func resourceIbmDistributionListDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListApiClient, err := meta.(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDistributionListDestinationOptions := &distributionlistapiv1.DeleteDistributionListDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "delete", "sep-id-parts").GetDiag()
	}

	deleteDistributionListDestinationOptions.SetAccountID(parts[0])
	deleteDistributionListDestinationOptions.SetDestinationID(core.UUIDPtr(strfmt.UUID(parts[1])))

	_, err = distributionListApiClient.DeleteDistributionListDestinationWithContext(context, deleteDistributionListDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmDistributionListDestinationMapToAddDestinationResponseBodyPrototypeEventNotificationDestination(modelMap map[string]interface{}) (*distributionlistapiv1.AddDestinationResponseBodyPrototypeEventNotificationDestination, error) {
	model := &distributionlistapiv1.AddDestinationResponseBodyPrototypeEventNotificationDestination{}
	model.ID = core.UUIDPtr(strfmt.UUID(modelMap["id"].(string)))
	model.DestinationType = core.StringPtr(modelMap["destination_type"].(string))
	return model, nil
}
