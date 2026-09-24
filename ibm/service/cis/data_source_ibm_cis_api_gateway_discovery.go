// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISApiGatewayDiscovery       = "ibm_cis_api_gateway_discovery"
	cisApiGatewayDiscoveryResultDoc = "result"
)

// DataSourceIBMCISApiGatewayDiscovery returns the schema.Resource for the
// ibm_cis_api_gateway_discovery data source.
func DataSourceIBMCISApiGatewayDiscovery() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISApiGatewayDiscoveryRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS instance CRN",
				ValidateFunc: validate.InvokeDataSourceValidator(
					ibmCISApiGatewayDiscovery,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayDiscoveryResultDoc: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Discovered operations rendered as a JSON-encoded OpenAPI schema document",
			},
		},
	}
}

// DataSourceIBMCISApiGatewayDiscoveryValidator returns the ResourceValidator
// for the ibm_cis_api_gateway_discovery data source.
func DataSourceIBMCISApiGatewayDiscoveryValidator() *validate.ResourceValidator {
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
		ResourceName: ibmCISApiGatewayDiscovery,
		Schema:       validateSchema,
	}
}

func dataSourceIBMCISApiGatewayDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
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

	opt := cisClient.NewGetApiGatewayDiscoveryOptions()
	result, response, err := cisClient.GetApiGatewayDiscoveryWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] GetApiGatewayDiscovery failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	if result != nil && result.Result != nil {
		jsonBytes, err := json.Marshal(result.Result)
		if err != nil {
			return fmt.Errorf("error serialising API Gateway discovery result: %w", err)
		}
		d.Set(cisApiGatewayDiscoveryResultDoc, string(jsonBytes))
	}
	return nil
}
