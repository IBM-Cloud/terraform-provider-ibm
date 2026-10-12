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
	ibmCISApiGatewayOperations        = "ibm_cis_api_gateway_operations"
	cisApiGatewayOperationsOps        = "operations"
	cisApiGatewayOperationsOpMethod   = "method"
	cisApiGatewayOperationsOpHost     = "host"
	cisApiGatewayOperationsOpEndpoint = "endpoint"
	cisApiGatewayOperationsOpID       = "operation_id"
)

// ResourceIBMCISApiGatewayOperations returns the schema.Resource for managing a
// set of API Gateway operations (bulk create / bulk delete) in a CIS zone.
//
// This resource wraps:
//   - POST /api_gateway/operations  (CreateZoneApiGatewayOperation)
//   - DELETE /api_gateway/operations (DeleteZoneApiGatewayOperations)
//
// Each operation in the list is identified by the server-assigned operation_id.
// Updates diff the desired list against the last-known list and issue individual
// single-operation deletes (DeleteZoneApiGatewayOperation) for removals followed
// by a bulk create for additions.
func ResourceIBMCISApiGatewayOperations() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISApiGatewayOperationsCreate,
		Read:     ResourceIBMCISApiGatewayOperationsRead,
		Update:   ResourceIBMCISApiGatewayOperationsUpdate,
		Delete:   ResourceIBMCISApiGatewayOperationsDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIS Instance CRN",
				ValidateFunc: validate.InvokeValidator(ibmCISApiGatewayOperations,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayOperationsOps: {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "List of API Gateway operations to create in Endpoint Management",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisApiGatewayOperationsOpMethod: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)",
						},
						cisApiGatewayOperationsOpHost: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "RFC3986-compliant host (e.g. api.example.com)",
						},
						cisApiGatewayOperationsOpEndpoint: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Endpoint path — must start with / (e.g. /v1/chat/completions)",
						},
						cisApiGatewayOperationsOpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server-assigned UUID for the operation",
						},
					},
				},
			},
		},
	}
}

// ResourceIBMCISApiGatewayOperationsValidator returns the validator for ibm_cis_api_gateway_operations.
func ResourceIBMCISApiGatewayOperationsValidator() *validate.ResourceValidator {
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
		ResourceName: ibmCISApiGatewayOperations,
		Schema:       validateSchema,
	}
}

// ResourceIBMCISApiGatewayOperationsCreate bulk-creates the requested operations and
// stores the server-assigned operation_ids back into state.
func ResourceIBMCISApiGatewayOperationsCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	crn, zoneID := cisCredentials(d)
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	ops := expandApiGatewayOperations(d.Get(cisApiGatewayOperationsOps).([]interface{}))
	if len(ops) == 0 {
		return fmt.Errorf("at least one operation must be specified")
	}

	opt := cisClient.NewCreateZoneApiGatewayOperationOptions()
	opt.SetApiGatewayOperation(ops)

	result, response, err := cisClient.CreateZoneApiGatewayOperationWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] CreateZoneApiGatewayOperation failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	if result != nil {
		if err := d.Set(cisApiGatewayOperationsOps, flattenBulkOperationsResult(result.Result)); err != nil {
			return fmt.Errorf("error setting operations: %w", err)
		}
	}
	return nil
}

// ResourceIBMCISApiGatewayOperationsRead reconciles state by listing operations from the
// API and matching them by (method, host, endpoint) against the stored list.
func ResourceIBMCISApiGatewayOperationsRead(d *schema.ResourceData, meta interface{}) error {
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

	opt := cisClient.NewListZoneApiGatewayOperationsOptions()
	result, response, err := cisClient.ListZoneApiGatewayOperationsWithContext(context.Background(), opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("[WARN] API Gateway operations for zone %s not found, removing from state", zoneID)
			d.SetId("")
			return nil
		}
		log.Printf("[ERROR] ListZoneApiGatewayOperations failed: %v", response)
		return err
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	if result == nil {
		return nil
	}

	// Build a lookup map from the live API: (method+host+endpoint) → operation_id.
	liveOps := make(map[string]string)
	for _, item := range result.Result {
		if item.Method != nil && item.Host != nil && item.Endpoint != nil && item.OperationID != nil {
			key := operationKey(*item.Method, *item.Host, *item.Endpoint)
			liveOps[key] = *item.OperationID
		}
	}

	// Reconcile against the stored list — update operation_ids from live data.
	stored := d.Get(cisApiGatewayOperationsOps).([]interface{})
	for _, raw := range stored {
		m := raw.(map[string]interface{})
		method, _ := m[cisApiGatewayOperationsOpMethod].(string)
		host, _ := m[cisApiGatewayOperationsOpHost].(string)
		endpoint, _ := m[cisApiGatewayOperationsOpEndpoint].(string)
		key := operationKey(method, host, endpoint)
		if id, ok := liveOps[key]; ok {
			m[cisApiGatewayOperationsOpID] = id
		}
	}
	d.Set(cisApiGatewayOperationsOps, stored)
	return nil
}

// ResourceIBMCISApiGatewayOperationsUpdate handles additions and removals within the
// operations list.  There is no bulk-update API; we delete removed items individually
// and bulk-create new items.
func ResourceIBMCISApiGatewayOperationsUpdate(d *schema.ResourceData, meta interface{}) error {
	if !d.HasChange(cisApiGatewayOperationsOps) {
		return nil
	}

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

	oldRaw, newRaw := d.GetChange(cisApiGatewayOperationsOps)
	oldOps := oldRaw.([]interface{})
	newOps := newRaw.([]interface{})

	// Build a key→operation_id map for the old list.
	oldMap := make(map[string]string)
	for _, raw := range oldOps {
		m := raw.(map[string]interface{})
		key := operationKey(
			m[cisApiGatewayOperationsOpMethod].(string),
			m[cisApiGatewayOperationsOpHost].(string),
			m[cisApiGatewayOperationsOpEndpoint].(string),
		)
		if id, ok := m[cisApiGatewayOperationsOpID].(string); ok && id != "" {
			oldMap[key] = id
		}
	}

	// Build a key set for the new list.
	newKeys := make(map[string]bool)
	var toCreate []aisecurityforappsv1.ApiGatewayOperation
	for _, raw := range newOps {
		m := raw.(map[string]interface{})
		method := m[cisApiGatewayOperationsOpMethod].(string)
		host := m[cisApiGatewayOperationsOpHost].(string)
		endpoint := m[cisApiGatewayOperationsOpEndpoint].(string)
		key := operationKey(method, host, endpoint)
		newKeys[key] = true
		if _, exists := oldMap[key]; !exists {
			// New operation — needs to be created.
			toCreate = append(toCreate, aisecurityforappsv1.ApiGatewayOperation{
				Method:   core.StringPtr(method),
				Host:     core.StringPtr(host),
				Endpoint: core.StringPtr(endpoint),
			})
		}
	}

	// Delete operations that are in old but not in new.
	for key, opID := range oldMap {
		if !newKeys[key] {
			delOpt := cisClient.NewDeleteZoneApiGatewayOperationOptions(opID)
			if _, delErr := cisClient.DeleteZoneApiGatewayOperationWithContext(context.Background(), delOpt); delErr != nil {
				log.Printf("[ERROR] DeleteZoneApiGatewayOperation(%s) failed: %v", opID, delErr)
				return delErr
			}
		}
	}

	// Bulk-create new operations.
	if len(toCreate) > 0 {
		createOpt := cisClient.NewCreateZoneApiGatewayOperationOptions()
		createOpt.SetApiGatewayOperation(toCreate)
		result, response, createErr := cisClient.CreateZoneApiGatewayOperationWithContext(context.Background(), createOpt)
		if createErr != nil {
			log.Printf("[ERROR] CreateZoneApiGatewayOperation (update) failed: %v", response)
			return createErr
		}
		if result != nil {
			if err := d.Set(cisApiGatewayOperationsOps, flattenBulkOperationsResult(result.Result)); err != nil {
				return fmt.Errorf("error setting operations after update: %w", err)
			}
		}
	}

	return ResourceIBMCISApiGatewayOperationsRead(d, meta)
}

// ResourceIBMCISApiGatewayOperationsDelete deletes each stored operation individually.
// The bulk DELETE /api_gateway/operations endpoint returns an empty response body on
// success which causes SDK unmarshal failures; the per-operation DELETE is reliable.
func ResourceIBMCISApiGatewayOperationsDelete(d *schema.ResourceData, meta interface{}) error {
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

	ops := d.Get(cisApiGatewayOperationsOps).([]interface{})
	for _, raw := range ops {
		m := raw.(map[string]interface{})
		opID, ok := m[cisApiGatewayOperationsOpID].(string)
		if !ok || opID == "" {
			continue
		}
		delOpt := cisClient.NewDeleteZoneApiGatewayOperationOptions(opID)
		if _, delErr := cisClient.DeleteZoneApiGatewayOperationWithContext(context.Background(), delOpt); delErr != nil {
			log.Printf("[ERROR] DeleteZoneApiGatewayOperation(%s) failed: %v", opID, delErr)
			return delErr
		}
	}

	d.SetId("")
	return nil
}

// ── helpers ──────────────────────────────────────────────────────────────────

// cisCredentials extracts the CRN and resolved zoneID from resource data.
func cisCredentials(d *schema.ResourceData) (crn, zoneID string) {
	crn = d.Get(cisID).(string)
	domainID := d.Get(cisDomainID).(string)
	zoneID, _, _ = flex.ConvertTftoCisTwoVar(domainID)
	if zoneID == "" {
		zoneID = domainID
	}
	return
}

// operationKey builds a stable lookup key for an (method, host, endpoint) triple.
func operationKey(method, host, endpoint string) string {
	return method + "|" + host + "|" + endpoint
}

// expandApiGatewayOperations converts the Terraform list value into the SDK slice.
func expandApiGatewayOperations(raw []interface{}) []aisecurityforappsv1.ApiGatewayOperation {
	ops := make([]aisecurityforappsv1.ApiGatewayOperation, 0, len(raw))
	for _, item := range raw {
		m := item.(map[string]interface{})
		ops = append(ops, aisecurityforappsv1.ApiGatewayOperation{
			Method:   core.StringPtr(m[cisApiGatewayOperationsOpMethod].(string)),
			Host:     core.StringPtr(m[cisApiGatewayOperationsOpHost].(string)),
			Endpoint: core.StringPtr(m[cisApiGatewayOperationsOpEndpoint].(string)),
		})
	}
	return ops
}

// flattenBulkOperationsResult converts the SDK bulk-create response items back into
// the Terraform TypeList structure, preserving method/host/endpoint alongside the
// server-assigned operation_id.
func flattenBulkOperationsResult(items []aisecurityforappsv1.ApiGatewayOperationsRespResultItem) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		m := map[string]interface{}{}
		if item.OperationID != nil {
			m[cisApiGatewayOperationsOpID] = *item.OperationID
		}
		if item.Method != nil {
			m[cisApiGatewayOperationsOpMethod] = *item.Method
		}
		if item.Host != nil {
			m[cisApiGatewayOperationsOpHost] = *item.Host
		}
		if item.Endpoint != nil {
			m[cisApiGatewayOperationsOpEndpoint] = *item.Endpoint
		}
		result = append(result, m)
	}
	return result
}
