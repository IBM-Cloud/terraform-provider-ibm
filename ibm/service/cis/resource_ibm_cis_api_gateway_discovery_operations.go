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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	ibmCISApiGatewayDiscoveryOperations       = "ibm_cis_api_gateway_discovery_operations"
	cisApiGatewayDiscoveryOpStates            = "operation_states"
	cisApiGatewayDiscoveryOpResult            = "result"
	cisApiGatewayDiscoveryOpResultID          = "id"
	cisApiGatewayDiscoveryOpResultState       = "state"
	cisApiGatewayDiscoveryOpResultEndpoint    = "endpoint"
	cisApiGatewayDiscoveryOpResultHost        = "host"
	cisApiGatewayDiscoveryOpResultMethod      = "method"
	cisApiGatewayDiscoveryOpResultOrigin      = "origin"
	cisApiGatewayDiscoveryOpResultLastUpdated = "last_updated"
)

// ResourceIBMCISApiGatewayDiscoveryOperations returns the schema.Resource for
// bulk-updating the state (saved/ignored/review) of discovered API Gateway operations.
func ResourceIBMCISApiGatewayDiscoveryOperations() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISApiGatewayDiscoveryOperationsCreateOrUpdate,
		Read:     ResourceIBMCISApiGatewayDiscoveryOperationsRead,
		Update:   ResourceIBMCISApiGatewayDiscoveryOperationsCreateOrUpdate,
		Delete:   ResourceIBMCISApiGatewayDiscoveryOperationsDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIS Instance CRN",
				ValidateFunc: validate.InvokeValidator(ibmCISApiGatewayDiscoveryOperations,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			// Map of operation_id -> state ("saved" | "ignored" | "review").
			cisApiGatewayDiscoveryOpStates: {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Map of discovered operation UUIDs to their desired state (saved, ignored, or review)",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"saved", "ignored", "review"}, false),
				},
			},
			cisApiGatewayDiscoveryOpResult: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of updated discovery operations returned by the API",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisApiGatewayDiscoveryOpResultID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID of the discovered operation",
						},
						cisApiGatewayDiscoveryOpResultEndpoint: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint path",
						},
						cisApiGatewayDiscoveryOpResultHost: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Host",
						},
						cisApiGatewayDiscoveryOpResultMethod: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP method",
						},
						cisApiGatewayDiscoveryOpResultState: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State after update (saved/ignored/review)",
						},
						cisApiGatewayDiscoveryOpResultOrigin: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Discovery engine origins",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						cisApiGatewayDiscoveryOpResultLastUpdated: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of last update (RFC3339)",
						},
					},
				},
			},
		},
	}
}

// ResourceIBMCISApiGatewayDiscoveryOperationsValidator returns the ResourceValidator
// for ibm_cis_api_gateway_discovery_operations.
func ResourceIBMCISApiGatewayDiscoveryOperationsValidator() *validate.ResourceValidator {
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
		ResourceName: ibmCISApiGatewayDiscoveryOperations,
		Schema:       validateSchema,
	}
}

// ResourceIBMCISApiGatewayDiscoveryOperationsCreateOrUpdate handles both Create and Update.
func ResourceIBMCISApiGatewayDiscoveryOperationsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
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

	// Build request body: map of operation_id -> state as map[string]interface{}.
	rawStates := d.Get(cisApiGatewayDiscoveryOpStates).(map[string]interface{})
	if len(rawStates) > 0 {
		requestBody := make(map[string]interface{}, len(rawStates))
		for opID, state := range rawStates {
			requestBody[opID] = state
		}

		opt := cisClient.NewUpdateZoneApiGatewayDiscoveryOperationOptions()
		opt.SetRequestBody(requestBody)

		result, response, err := cisClient.UpdateZoneApiGatewayDiscoveryOperationWithContext(context.Background(), opt)
		if err != nil {
			log.Printf("[ERROR] UpdateZoneApiGatewayDiscoveryOperation failed: %v", response)
			return err
		}

		if result != nil {
			if err := d.Set(cisApiGatewayDiscoveryOpResult, flattenDiscoveryOperationsList(result.Result)); err != nil {
				return fmt.Errorf("error setting result: %w", err)
			}
		}
	} else {
		d.Set(cisApiGatewayDiscoveryOpResult, []map[string]interface{}{})
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	return nil
}

// ResourceIBMCISApiGatewayDiscoveryOperationsRead is a no-op: there is no GET endpoint
// for the bulk discovery operation state map; state is maintained from the last apply.
func ResourceIBMCISApiGatewayDiscoveryOperationsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// ResourceIBMCISApiGatewayDiscoveryOperationsDelete resets all tracked operations
// back to "review" (the default discovery state) and removes the resource from state.
func ResourceIBMCISApiGatewayDiscoveryOperationsDelete(d *schema.ResourceData, meta interface{}) error {
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

	// Reset every tracked operation back to "review" if any were tracked.
	rawStates := d.Get(cisApiGatewayDiscoveryOpStates).(map[string]interface{})
	if len(rawStates) > 0 {
		requestBody := make(map[string]interface{}, len(rawStates))
		for opID := range rawStates {
			requestBody[opID] = "review"
		}

		opt := cisClient.NewUpdateZoneApiGatewayDiscoveryOperationOptions()
		opt.SetRequestBody(requestBody)

		_, response, err := cisClient.UpdateZoneApiGatewayDiscoveryOperationWithContext(context.Background(), opt)
		if err != nil {
			log.Printf("[ERROR] Resetting discovery operation states on delete failed: %v", response)
			return err
		}
	}

	d.SetId("")
	return nil
}

// flattenDiscoveryOperationsList converts a slice of DiscoveryOperation SDK objects
// into the Terraform TypeList structure used by cisApiGatewayDiscoveryOpResult.
func flattenDiscoveryOperationsList(ops []aisecurityforappsv1.DiscoveryOperation) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(ops))
	for _, op := range ops {
		m := map[string]interface{}{}
		if op.ID != nil {
			m[cisApiGatewayDiscoveryOpResultID] = *op.ID
		}
		if op.Endpoint != nil {
			m[cisApiGatewayDiscoveryOpResultEndpoint] = *op.Endpoint
		}
		if op.Host != nil {
			m[cisApiGatewayDiscoveryOpResultHost] = *op.Host
		}
		if op.Method != nil {
			m[cisApiGatewayDiscoveryOpResultMethod] = *op.Method
		}
		if op.State != nil {
			m[cisApiGatewayDiscoveryOpResultState] = *op.State
		}
		if op.LastUpdated != nil {
			m[cisApiGatewayDiscoveryOpResultLastUpdated] = op.LastUpdated.String()
		}
		m[cisApiGatewayDiscoveryOpResultOrigin] = op.Origin
		result = append(result, m)
	}
	return result
}
