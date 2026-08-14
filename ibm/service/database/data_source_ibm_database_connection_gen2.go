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

	// Persist the effective selection so acceptance tests and state reflect
	// the actual key used when fallback-to-first-key behavior is exercised.
	if selectedKey.Name != nil {
		if err = d.Set("user_id", *selectedKey.Name); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting user_id: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
	}

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
	// The structure is nested under credentials.connection for Gen2 databases

	// Check if connection information is nested under "connection" key
	var connectionData map[string]interface{}
	if connObj, ok := credentials["connection"].(map[string]interface{}); ok {
		connectionData = connObj
		log.Printf("[DEBUG] Found connection data nested under 'connection' key")
		// Log all available connection types for debugging
		log.Printf("[DEBUG] Available connection types in credentials.connection:")
		for key := range connectionData {
			log.Printf("[DEBUG]   - %s", key)
		}
	} else {
		// Fallback to direct credentials structure (for backward compatibility)
		connectionData = credentials
		log.Printf("[DEBUG] Using direct credentials structure")
		log.Printf("[DEBUG] Available keys in credentials:")
		for key := range connectionData {
			log.Printf("[DEBUG]   - %s", key)
		}
	}

	// Transform and set postgres connection if available
	if postgresConn, ok := connectionData["postgres"].(map[string]interface{}); ok {
		transformedPostgres, err := transformGen2ConnectionToSchema(postgresConn)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error transforming postgres connection: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
		postgres := []map[string]interface{}{transformedPostgres}
		if err = d.Set("postgres", postgres); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting postgres: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
		log.Printf("[DEBUG] Successfully set postgres connection data")
	} else {
		log.Printf("[DEBUG] No postgres connection data found in credentials")
	}

	// Transform and set CLI connection if available
	if cliConn, ok := connectionData["cli"].(map[string]interface{}); ok {
		transformedCli, err := transformGen2CliToSchema(cliConn)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error transforming cli connection: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
		cli := []map[string]interface{}{transformedCli}
		if err = d.Set("cli", cli); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting cli: %s", err), "(Data) ibm_database_connection", "read")
			return tfErr.GetDiag()
		}
		log.Printf("[DEBUG] Successfully set cli connection data")
	} else {
		log.Printf("[DEBUG] No cli connection data found in credentials")
	}

	// Transform and set MongoDB connection if available
	// Try both "mongodb" and "mongo" keys as different services may use different names
	var mongodbConn map[string]interface{}
	var mongodbFound bool
	if conn, ok := connectionData["mongodb"].(map[string]interface{}); ok {
		mongodbConn = conn
		mongodbFound = true
		log.Printf("[DEBUG] Found mongodb connection data under 'mongodb' key")
	} else if conn, ok := connectionData["mongo"].(map[string]interface{}); ok {
		mongodbConn = conn
		mongodbFound = true
		log.Printf("[DEBUG] Found mongodb connection data under 'mongo' key")
	}

	if mongodbFound {
		transformedMongodb, err := transformGen2ConnectionToSchema(mongodbConn)
		if err != nil {
			log.Printf("[DEBUG] Error transforming mongodb connection: %s", err)
		} else {
			mongodb := []map[string]interface{}{transformedMongodb}
			if err = d.Set("mongodb", mongodb); err != nil {
				log.Printf("[DEBUG] Error setting mongodb: %s", err)
			} else {
				log.Printf("[DEBUG] Successfully set mongodb connection data")
			}
		}
	} else {
		log.Printf("[DEBUG] No mongodb connection data found in credentials (tried 'mongodb' and 'mongo' keys)")
	}

	// Transform and set other connection types with flexible key matching
	connectionTypes := map[string][]string{
		"rediss":       {"rediss", "redis"},
		"https":        {"https", "http"},
		"amqps":        {"amqps", "amqp"},
		"mqtts":        {"mqtts", "mqtt"},
		"stomp_ssl":    {"stomp_ssl", "stomp"},
		"mysql":        {"mysql"},
		"grpc":         {"grpc"},
		"bi_connector": {"bi_connector"},
		"analytics":    {"analytics"},
		"ops_manager":  {"ops_manager"},
		"emp":          {"emp"},
	}

	for tfKey, possibleKeys := range connectionTypes {
		var conn map[string]interface{}
		var found bool
		var foundKey string

		for _, key := range possibleKeys {
			if c, ok := connectionData[key].(map[string]interface{}); ok {
				conn = c
				found = true
				foundKey = key
				break
			}
		}

		if found {
			log.Printf("[DEBUG] Found %s connection data under '%s' key", tfKey, foundKey)
			transformed, err := transformGen2ConnectionToSchema(conn)
			if err != nil {
				log.Printf("[DEBUG] Error transforming %s connection: %s", tfKey, err)
				continue
			}
			connList := []map[string]interface{}{transformed}
			if err = d.Set(tfKey, connList); err != nil {
				log.Printf("[DEBUG] Error setting %s: %s", tfKey, err)
				continue
			}
			log.Printf("[DEBUG] Successfully set %s connection data", tfKey)
		}
	}

	return nil
}

// transformGen2ConnectionToSchema transforms Gen2 API connection structure to Terraform schema format
func transformGen2ConnectionToSchema(conn map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Copy simple fields
	if v, ok := conn["type"]; ok {
		result["type"] = v
	}
	if v, ok := conn["composed"]; ok {
		result["composed"] = v
	}
	if v, ok := conn["scheme"]; ok {
		result["scheme"] = v
	}
	if v, ok := conn["path"]; ok {
		result["path"] = v
	}
	if v, ok := conn["database"]; ok {
		result["database"] = v
	}
	if v, ok := conn["ssl"]; ok {
		result["ssl"] = v
	}
	if v, ok := conn["browser_accessible"]; ok {
		result["browser_accessible"] = v
	}

	// Transform hosts array
	if hostsRaw, ok := conn["hosts"].([]interface{}); ok {
		hosts := make([]map[string]interface{}, 0, len(hostsRaw))
		for _, hostRaw := range hostsRaw {
			if hostMap, ok := hostRaw.(map[string]interface{}); ok {
				host := make(map[string]interface{})
				if hostname, ok := hostMap["hostname"]; ok {
					host["hostname"] = hostname
				}
				if port, ok := hostMap["port"]; ok {
					// Convert port to int64 if it's a float64
					switch p := port.(type) {
					case float64:
						host["port"] = int64(p)
					case int:
						host["port"] = int64(p)
					case int64:
						host["port"] = p
					default:
						host["port"] = port
					}
				}
				hosts = append(hosts, host)
			}
		}
		result["hosts"] = hosts
	}

	// Transform authentication
	if authRaw, ok := conn["authentication"].(map[string]interface{}); ok {
		auth := make(map[string]interface{})
		if method, ok := authRaw["method"]; ok {
			auth["method"] = method
		}
		if username, ok := authRaw["username"]; ok {
			auth["username"] = username
		}
		if password, ok := authRaw["password"]; ok {
			auth["password"] = password
		}
		result["authentication"] = []map[string]interface{}{auth}
	}

	// Transform certificate
	if certRaw, ok := conn["certificate"].(map[string]interface{}); ok {
		cert := make(map[string]interface{})
		if name, ok := certRaw["name"]; ok {
			cert["name"] = name
		}
		if certBase64, ok := certRaw["certificate_base64"]; ok {
			cert["certificate_base64"] = certBase64
		}
		result["certificate"] = []map[string]interface{}{cert}
	}

	// Transform query_options - convert all values to strings as Terraform schema expects
	if queryOpts, ok := conn["query_options"].(map[string]interface{}); ok {
		convertedOpts := make(map[string]interface{})
		for key, value := range queryOpts {
			// Convert boolean values to strings
			switch v := value.(type) {
			case bool:
				convertedOpts[key] = fmt.Sprintf("%t", v)
			case float64:
				// Convert numbers to strings
				convertedOpts[key] = fmt.Sprintf("%v", v)
			case int, int64:
				convertedOpts[key] = fmt.Sprintf("%v", v)
			default:
				convertedOpts[key] = value
			}
		}
		result["query_options"] = convertedOpts
	}

	return result, nil
}

// transformGen2CliToSchema transforms Gen2 API CLI structure to Terraform schema format
func transformGen2CliToSchema(cli map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Copy simple fields
	if v, ok := cli["type"]; ok {
		result["type"] = v
	}
	if v, ok := cli["composed"]; ok {
		result["composed"] = v
	}
	if v, ok := cli["bin"]; ok {
		result["bin"] = v
	}
	if v, ok := cli["arguments"]; ok {
		result["arguments"] = v
	}

	// Transform environment variables
	if envRaw, ok := cli["environment"].(map[string]interface{}); ok {
		result["environment"] = envRaw
	}

	// Transform certificate
	if certRaw, ok := cli["certificate"].(map[string]interface{}); ok {
		cert := make(map[string]interface{})
		if name, ok := certRaw["name"]; ok {
			cert["name"] = name
		}
		if certBase64, ok := certRaw["certificate_base64"]; ok {
			cert["certificate_base64"] = certBase64
		}
		result["certificate"] = []map[string]interface{}{cert}
	}

	return result, nil
}
