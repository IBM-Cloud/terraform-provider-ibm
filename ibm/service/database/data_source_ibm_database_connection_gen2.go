// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// Gen2 backend implementation for database connections
// This uses Resource Controller APIs to retrieve connection information from resource keys
type dataSourceIBMDatabaseConnectionGen2Backend struct{}

func newDataSourceIBMDatabaseConnectionGen2Backend() dataSourceIBMDatabaseConnectionBackend {
	return &dataSourceIBMDatabaseConnectionGen2Backend{}
}

func (g *dataSourceIBMDatabaseConnectionGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// NOTE - Edge case: potential stale values for unsupported Gen2 attributes.
	// If this data source was previously resolved to a Classic instance, all
	// attributes (including ones not supported by Gen2) would have been set.
	// If the same filters later resolve to a Gen2 instance (e.g., deployment_id),
	// Terraform will not automatically clear attributes that are no longer set,
	// unlike a resource which would ForceNew on such changes.
	// As a result, the Gen2 read path may only set supported attributes while
	// previously populated Classic-only attributes remain stale in state.
	// There is no clean mechanism to fully reset datasource state, and doing so
	// is generally considered an anti-pattern.
	// If this becomes an issue, unsupported attributes could be explicitly set
	// to null via d.Set() to ensure stale values are cleared.

	// Gen2 databases use Resource Controller API, not CloudDatabasesV5
	// Get the resource controller client to fetch instance details
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deploymentID := d.Get("deployment_id").(string)
	userID := d.Get("user_id").(string)
	userType := d.Get("user_type").(string)
	endpointType := d.Get("endpoint_type").(string)

	// Get the instance to verify it exists and is accessible
	instance, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &deploymentID,
	})
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetResourceInstance failed: %s\n%s", err.Error(), response), "(Data) ibm_database_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Verify this is actually a Gen2 instance
	if instance.ResourcePlanID != nil && !isGen2Plan(*instance.ResourcePlanID) {
		tfErr := flex.TerraformErrorf(
			fmt.Errorf("instance %s is not a Gen2 database", deploymentID),
			"Instance is not a Gen2 database",
			"(Data) ibm_database_connection",
			"read",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// For Gen2 databases, connection information is retrieved through resource keys
	// List all resource keys for this instance
	listKeysOptions := &rc.ListResourceKeysForInstanceOptions{
		ID: &deploymentID,
	}

	keysList, response, err := rsConClient.ListResourceKeysForInstance(listKeysOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceKeysForInstance failed: %s\n%s", err.Error(), response), "(Data) ibm_database_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Find a key that matches the user_id or use the first available key
	var selectedKey *rc.ResourceKey
	for _, key := range keysList.Resources {
		if key.Name != nil && *key.Name == userID {
			selectedKey = &key
			break
		}
	}

	// If no matching key found by name, use the first available key
	if selectedKey == nil && len(keysList.Resources) > 0 {
		selectedKey = &keysList.Resources[0]
		log.Printf("[DEBUG] No resource key found with name '%s', using first available key: %s", userID, *selectedKey.Name)
	}

	if selectedKey == nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "No resource keys found for Gen2 database",
				Detail: fmt.Sprintf(
					"No resource keys found for Gen2 database (deployment_id: %s). "+
						"Gen2 databases require resource keys to access connection information. "+
						"Please create a resource key using ibm_resource_key resource first.",
					deploymentID,
				),
			},
		}
	}

	d.SetId(DataSourceIBMDatabaseConnectionID(d))

	// Extract connection information from the resource key credentials
	if selectedKey.Credentials == nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Resource key credentials not available",
				Detail: fmt.Sprintf(
					"Resource key '%s' does not contain credentials. "+
						"This may be due to insufficient permissions or the key being in an invalid state.",
					*selectedKey.Name,
				),
			},
		}
	}

	// Parse credentials to extract connection information
	credBytes, err := json.Marshal(selectedKey.Credentials)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "Failed to marshal credentials", "(Data) ibm_database_connection", "read")
		return tfErr.GetDiag()
	}

	var credentials map[string]interface{}
	if err := json.Unmarshal(credBytes, &credentials); err != nil {
		tfErr := flex.TerraformErrorf(err, "Failed to unmarshal credentials", "(Data) ibm_database_connection", "read")
		return tfErr.GetDiag()
	}

	log.Printf("[DEBUG] Gen2 database connection information retrieved from resource key: %s", *selectedKey.Name)
	log.Printf("[DEBUG] User type: %s, Endpoint type: %s", userType, endpointType)

	// Gen2 databases provide connection information in the credentials
	// Extract and set connection details based on what's available in the credentials
	// The structure is similar to Classic but comes from resource key credentials

	// Set postgres connection if available
	if postgresConn, ok := credentials["postgres"].(map[string]interface{}); ok {
		postgres := []map[string]interface{}{postgresConn}
		if err = d.Set("postgres", postgres); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting postgres: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set CLI connection if available
	if cliConn, ok := credentials["cli"].(map[string]interface{}); ok {
		cli := []map[string]interface{}{cliConn}
		if err = d.Set("cli", cli); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting cli: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set rediss connection if available
	if redissConn, ok := credentials["rediss"].(map[string]interface{}); ok {
		rediss := []map[string]interface{}{redissConn}
		if err = d.Set("rediss", rediss); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting rediss: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set https connection if available
	if httpsConn, ok := credentials["https"].(map[string]interface{}); ok {
		https := []map[string]interface{}{httpsConn}
		if err = d.Set("https", https); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting https: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set amqps connection if available
	if amqpsConn, ok := credentials["amqps"].(map[string]interface{}); ok {
		amqps := []map[string]interface{}{amqpsConn}
		if err = d.Set("amqps", amqps); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting amqps: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set mqtts connection if available
	if mqttsConn, ok := credentials["mqtts"].(map[string]interface{}); ok {
		mqtts := []map[string]interface{}{mqttsConn}
		if err = d.Set("mqtts", mqtts); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting mqtts: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set stomp_ssl connection if available
	if stompSslConn, ok := credentials["stomp_ssl"].(map[string]interface{}); ok {
		stompSsl := []map[string]interface{}{stompSslConn}
		if err = d.Set("stomp_ssl", stompSsl); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting stomp_ssl: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set mongodb connection if available
	if mongodbConn, ok := credentials["mongodb"].(map[string]interface{}); ok {
		mongodb := []map[string]interface{}{mongodbConn}
		if err = d.Set("mongodb", mongodb); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting mongodb: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	// Set mysql connection if available
	if mysqlConn, ok := credentials["mysql"].(map[string]interface{}); ok {
		mysql := []map[string]interface{}{mysqlConn}
		if err = d.Set("mysql", mysql); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting mysql: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}
