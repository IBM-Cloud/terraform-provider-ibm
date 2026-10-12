// Copyright IBM Corp. 2026 All Rights Reserved.
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

func DataSourceIBMCISApiGatewayOperation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISApiGatewayOperationRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance CRN",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					ibmCISApiGatewayOperation,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain ID",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayOperationID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of the API Gateway operation",
			},
			cisApiGatewayOperationEndpt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API endpoint path (e.g. /api/v1/chat)",
			},
			cisApiGatewayOperationHost: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API host (e.g. api.example.com)",
			},
			cisApiGatewayOperationMethod: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)",
			},
		},
	}
}

func DataSourceIBMCISApiGatewayOperationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	ibmCISApiGatewayOperationDataSourceValidator := validate.ResourceValidator{
		ResourceName: ibmCISApiGatewayOperation,
		Schema:       validateSchema,
	}
	return &ibmCISApiGatewayOperationDataSourceValidator
}

func dataSourceIBMCISApiGatewayOperationRead(d *schema.ResourceData, meta interface{}) error {
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
	operationID := d.Get(cisApiGatewayOperationID).(string)

	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetZoneApiGatewayOperationOptions(operationID)
	result, response, err := cisClient.GetZoneApiGatewayOperationWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] Get API Gateway Operation failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfThreeVar(operationID, zoneID, crn))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	if result != nil && result.Result != nil {
		r := result.Result
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
