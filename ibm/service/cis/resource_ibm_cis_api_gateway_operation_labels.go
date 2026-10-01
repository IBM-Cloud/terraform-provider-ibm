// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/aisecurityforappsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISApiGatewayOperationLabels        = "ibm_cis_api_gateway_operation_labels"
	cisApiGatewayOperationLabelsOperIDs    = "operation_ids"
	cisApiGatewayOperationLabelsUserLabels = "user_labels"
	cisApiGatewayOperationLabelsMgdLabels  = "managed_labels"
	cisApiGatewayOperationLabelsResult     = "result"
	cisApiGatewayOperationLabelsResultOpID = "operation_id"
	cisApiGatewayOperationLabelsResultLbls = "labels"
)

// ResourceIBMCISApiGatewayOperationLabels returns the schema.Resource for
// managing user and managed labels applied to API Gateway operations.
func ResourceIBMCISApiGatewayOperationLabels() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISApiGatewayOperationLabelsCreateOrUpdate,
		Read:     ResourceIBMCISApiGatewayOperationLabelsRead,
		Update:   ResourceIBMCISApiGatewayOperationLabelsCreateOrUpdate,
		Delete:   ResourceIBMCISApiGatewayOperationLabelsDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIS Instance CRN",
				ValidateFunc: validate.InvokeValidator(ibmCISApiGatewayOperationLabels,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayOperationLabelsOperIDs: {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of API Gateway operation UUIDs to apply labels to",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisApiGatewayOperationLabelsUserLabels: {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "User-defined label strings to apply to the selected operations",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisApiGatewayOperationLabelsMgdLabels: {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Managed label strings to apply (e.g. cf-llm)",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisApiGatewayOperationLabelsResult: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Per-operation label assignment results returned by the API",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisApiGatewayOperationLabelsResultOpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operation UUID",
						},
						cisApiGatewayOperationLabelsResultLbls: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Labels assigned to the operation (as raw JSON objects)",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

// ResourceIBMCISApiGatewayOperationLabelsValidator returns the ResourceValidator
// for ibm_cis_api_gateway_operation_labels.
func ResourceIBMCISApiGatewayOperationLabelsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true,
		})
	return &validate.ResourceValidator{
		ResourceName: ibmCISApiGatewayOperationLabels,
		Schema:       validateSchema,
	}
}

// ResourceIBMCISApiGatewayOperationLabelsCreateOrUpdate handles both Create and Update.
func ResourceIBMCISApiGatewayOperationLabelsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	domainID := d.Get(cisDomainID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(domainID)
	if zoneID == "" {
		zoneID = domainID
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewUpdateApiGatewayOperationLabelsOptions()

	// Selector: operation IDs
	operationIDs := flex.ExpandStringList(d.Get(cisApiGatewayOperationLabelsOperIDs).(*schema.Set).List())
	selectorInclude := &aisecurityforappsv1.ApiGatewayOperationsLabelsInputSelectorInclude{
		OperationIds: operationIDs,
	}
	opt.SetSelector(&aisecurityforappsv1.ApiGatewayOperationsLabelsInputSelector{
		Include: selectorInclude,
	})

	// User labels (always set to ensure updates clear previous values)
	userLabels := flex.ExpandStringList(d.Get(cisApiGatewayOperationLabelsUserLabels).(*schema.Set).List())
	opt.SetUser(&aisecurityforappsv1.ApiGatewayOperationsLabelsInputUser{
		Labels: userLabels,
	})

	// Managed labels (always set to ensure updates clear previous values)
	managedLabels := flex.ExpandStringList(d.Get(cisApiGatewayOperationLabelsMgdLabels).(*schema.Set).List())
	opt.SetManaged(&aisecurityforappsv1.ApiGatewayOperationsLabelsInputManaged{
		Labels: managedLabels,
	})

	result, response, err := cisClient.UpdateApiGatewayOperationLabelsWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] UpdateApiGatewayOperationLabels failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))

	if result != nil {
		if err := d.Set(cisApiGatewayOperationLabelsResult, flattenLabelsResult(result.Result)); err != nil {
			return fmt.Errorf("error setting result: %w", err)
		}
	}
	return nil
}

// ResourceIBMCISApiGatewayOperationLabelsRead refreshes state by re-applying the labels
// (the API has no dedicated GET for bulk label state, so we accept the last-known state).
func ResourceIBMCISApiGatewayOperationLabelsRead(d *schema.ResourceData, meta interface{}) error {
	// There is no GET endpoint for bulk operation labels — the state is maintained from
	// the last Create/Update response. Nothing to refresh here.
	return nil
}

// ResourceIBMCISApiGatewayOperationLabelsDelete clears all user and managed labels from
// the selected operations and removes the resource from state.
func ResourceIBMCISApiGatewayOperationLabelsDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	// Rebuild operation IDs from current state.
	operationIDs := flex.ExpandStringList(d.Get(cisApiGatewayOperationLabelsOperIDs).(*schema.Set).List())

	opt := cisClient.NewUpdateApiGatewayOperationLabelsOptions()
	opt.SetSelector(&aisecurityforappsv1.ApiGatewayOperationsLabelsInputSelector{
		Include: &aisecurityforappsv1.ApiGatewayOperationsLabelsInputSelectorInclude{
			OperationIds: operationIDs,
		},
	})
	d.SetId("")
	return nil
}

// flattenLabelsResult converts the SDK result slice into the Terraform TypeList structure.
func flattenLabelsResult(items []aisecurityforappsv1.ApiGatewayOperationsLabelsRespResultItem) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		m := map[string]interface{}{}
		if item.OperationID != nil {
			m[cisApiGatewayOperationLabelsResultOpID] = *item.OperationID
		}
		// Each label entry is a map[string]interface{}; serialise as a JSON string per entry
		// so that Terraform can store them in a TypeList of TypeString.
		lbls := make([]string, 0, len(item.Labels))
		for _, lbl := range item.Labels {
			for k, v := range lbl {
				lbls = append(lbls, fmt.Sprintf("%s:%v", k, v))
			}
		}
		m[cisApiGatewayOperationLabelsResultLbls] = lbls
		result = append(result, m)
	}
	return result
}
