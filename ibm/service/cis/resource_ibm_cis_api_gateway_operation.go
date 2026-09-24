// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISApiGatewayOperation    = "ibm_cis_api_gateway_operation"
	cisApiGatewayOperationID     = "operation_id"
	cisApiGatewayOperationHost   = "host"
	cisApiGatewayOperationMethod = "method"
	cisApiGatewayOperationEndpt  = "endpoint"
)

func ResourceIBMCISApiGatewayOperation() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISApiGatewayOperationCreate,
		Read:     ResourceIBMCISApiGatewayOperationRead,
		Delete:   ResourceIBMCISApiGatewayOperationDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIS Instance CRN",
				ValidateFunc: validate.InvokeValidator(ibmCISApiGatewayOperation,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayOperationEndpt: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "API endpoint path (e.g. /api/v1/chat)",
			},
			cisApiGatewayOperationHost: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "API host (e.g. api.example.com)",
			},
			cisApiGatewayOperationMethod: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)",
			},
			cisApiGatewayOperationID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID of the created API Gateway operation",
			},
		},
	}
}

func ResourceIBMCISApiGatewayOperationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	ibmCISApiGatewayOperationValidator := validate.ResourceValidator{
		ResourceName: ibmCISApiGatewayOperation,
		Schema:       validateSchema,
	}
	return &ibmCISApiGatewayOperationValidator
}

func ResourceIBMCISApiGatewayOperationCreate(d *schema.ResourceData, meta interface{}) error {
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

	opt := cisClient.NewCreateApiGatewayOperationItemOptions()
	opt.SetEndpoint(d.Get(cisApiGatewayOperationEndpt).(string))
	opt.SetHost(d.Get(cisApiGatewayOperationHost).(string))
	opt.SetMethod(d.Get(cisApiGatewayOperationMethod).(string))

	result, response, err := cisClient.CreateApiGatewayOperationItemWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] Create API Gateway Operation failed: %v", response)
		return err
	}

	operationID := *result.Result.OperationID
	d.SetId(flex.ConvertCisToTfThreeVar(operationID, zoneID, crn))
	d.Set(cisApiGatewayOperationID, operationID)
	return ResourceIBMCISApiGatewayOperationRead(d, meta)
}

func ResourceIBMCISApiGatewayOperationRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	operationID, zoneID, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetZoneApiGatewayOperationOptions(operationID)
	result, response, err := cisClient.GetZoneApiGatewayOperationWithContext(context.Background(), opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("[WARN] API Gateway Operation %s not found, removing from state", operationID)
			d.SetId("")
			return nil
		}
		log.Printf("[ERROR] Get API Gateway Operation failed: %v", response)
		return err
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	if result != nil && result.Result != nil {
		r := result.Result
		if r.OperationID != nil {
			d.Set(cisApiGatewayOperationID, *r.OperationID)
		}
		if r.Endpoint != nil {
			d.Set(cisApiGatewayOperationEndpt, *r.Endpoint)
		}
		if r.Host != nil {
			d.Set(cisApiGatewayOperationHost, *r.Host)
		}
		if r.Method != nil {
			d.Set(cisApiGatewayOperationMethod, *r.Method)
		}
	}
	return nil
}

func ResourceIBMCISApiGatewayOperationDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	operationID, zoneID, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewDeleteZoneApiGatewayOperationOptions(operationID)
	response, err := cisClient.DeleteZoneApiGatewayOperationWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] Delete API Gateway Operation failed: %v", response)
		return err
	}

	d.SetId("")
	return nil
}
