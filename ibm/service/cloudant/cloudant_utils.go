// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var cloudantGen1PlanNames = map[string]struct{}{
	"dedicated-hardware": {},
	"lite":               {},
	"standard":           {},
}

// isCloudantGen2Plan reports whether the given plan name identifies a Gen 2
// Cloudant instance. Any plan not in the known Gen 1 set is treated as Gen 2.
// An empty or whitespace-only plan defaults to Gen 1 behaviour.
func isCloudantGen2Plan(plan string) bool {
	normalized := strings.ToLower(strings.TrimSpace(plan))
	if normalized == "" {
		return false
	}
	_, ok := cloudantGen1PlanNames[normalized]
	return !ok
}

// isCloudantGen2PlanFrom is a convenience wrapper around isCloudantGen2Plan
// that reads the "plan" attribute from rd, accepting either *schema.ResourceData
// or *schema.ResourceDiff via the resourceDataGetter interface.
func isCloudantGen2PlanFrom(rd resourceDataGetter) bool {
	return isCloudantGen2Plan(rd.Get("plan").(string))
}

// GetCloudantClientFromCrn creates an authenticated Cloudant SDK client for a given instance CRN.
func GetCloudantClientFromCrn(instanceCRN string, meta interface{}, resourceName string, operation string) (*cloudantv1.CloudantV1, *flex.TerraformProblem) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "get-resource-controller-client")
	}

	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: flex.PtrToString(instanceCRN),
	}

	instance, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
	if err != nil {
		err = fmt.Errorf("Error retrieving resource instance: %s with resp code: %s", err, resp)
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "get-resource-instance")
	}

	instanceExtensionMap, tfErr := getCloudantExtensions(instance.Extensions, resourceName, operation)
	if tfErr != nil {
		return nil, tfErr
	}

	return getCloudantClientFromExtensions(instanceExtensionMap, meta, resourceName, operation)
}

// GetCloudantClientFromResource creates an authenticated Cloudant SDK client from resource data extensions.
func GetCloudantClientFromResource(d *schema.ResourceData, meta interface{}, resourceName string, operation string) (*cloudantv1.CloudantV1, *flex.TerraformProblem) {
	instanceExtensionMap, tfErr := getCloudantExtensions(d.Get("extensions"), resourceName, operation)
	if tfErr != nil {
		return nil, tfErr
	}

	return getCloudantClientFromExtensions(instanceExtensionMap, meta, resourceName, operation)
}

func getCloudantExtensions(rawExtensions interface{}, resourceName string, operation string) (flex.Map, *flex.TerraformProblem) {
	extensions, ok := rawExtensions.(map[string]interface{})
	if !ok || extensions == nil {
		err := fmt.Errorf("Missing Cloudant extensions")
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "missing-extensions")
	}

	result := flex.Flatten(extensions)
	if len(result) == 0 {
		err := fmt.Errorf("Cloudant extensions are empty — instance may not be fully provisioned")
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "empty-extensions")
	}
	return result, nil
}

// getCloudantClientFromExtensions creates an authenticated Cloudant SDK client from flattened Cloudant extensions.
func getCloudantClientFromExtensions(instanceExtensionMap flex.Map, meta interface{}, resourceName string, operation string) (*cloudantv1.CloudantV1, *flex.TerraformProblem) {
	endpoint, isGen2, err := getCloudantInstanceUrl(instanceExtensionMap, meta)
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "get-instance-url")
	}

	client, err := getCloudantClientForUrl(endpoint, isGen2, meta)
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), resourceName, operation, "get-client")
	}

	return client, nil
}

// getCloudantInstanceUrl retrieves the Cloudant instance URL from flattened extensions.
// It handles both Gen1 and Gen2 endpoint formats and respects the visibility
// configuration (private, public, or public-and-private) from the provider.
// Returns the endpoint URL and a boolean indicating if it's a Gen2 instance.
func getCloudantInstanceUrl(instanceExtensionMap flex.Map, meta interface{}) (string, bool, error) {
	bxSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return "", false, fmt.Errorf("Error getting session: %s", err)
	}

	cloudantInstanceUrl, isGen2 := selectCloudantEndpoint(instanceExtensionMap, bxSession.Config.Visibility)
	if cloudantInstanceUrl == "" {
		if bxSession.Config.Visibility == "private" {
			return "", false, fmt.Errorf("Unable to get private URL for Cloudant instance: no private endpoint available")
		}
		return "", false, fmt.Errorf("Unable to get URL for cloudant instance")
	}

	cloudantInstanceUrl = conns.EnvFallBack([]string{"IBMCLOUD_CLOUDANT_API_ENDPOINT"}, cloudantInstanceUrl)
	return cloudantInstanceUrl, isGen2, nil
}

// selectCloudantEndpoint selects the appropriate Cloudant endpoint based on
// visibility configuration and available endpoint formats (Gen1 vs Gen2).
// Returns the selected endpoint URL and a boolean indicating if it's a Gen2 instance.
//
// Visibility modes:
// - "private": Only use private endpoints
// - "public-and-private": Prefer private, fallback to public
// - default: Use public endpoints
func selectCloudantEndpoint(instanceExtensionMap flex.Map, visibility string) (string, bool) {
	// Determine which generation is available and set public/private URLs accordingly
	var publicURL, privateURL string
	var isGen2 bool

	// Check Gen2 format: extensions.dataservices.connection
	gen2Public := instanceExtensionMap["dataservices.connection.public_endpoint_url"]
	gen2Private := instanceExtensionMap["dataservices.connection.vpe_url"]

	if gen2Public != "" || gen2Private != "" {
		// Gen2 instance
		publicURL = gen2Public
		privateURL = gen2Private
		isGen2 = true
	} else {
		// Gen1 instance: extensions.endpoints
		publicURL = normalizeEndpoint(instanceExtensionMap["endpoints.public"])
		privateURL = normalizeEndpoint(instanceExtensionMap["endpoints.private"])
		isGen2 = false
	}

	// Select endpoint based on visibility configuration
	var selectedURL string
	switch visibility {
	case "private":
		selectedURL = privateURL
	case "public-and-private":
		if privateURL != "" {
			selectedURL = privateURL
		} else {
			selectedURL = publicURL
		}
	default:
		selectedURL = publicURL
	}

	return selectedURL, isGen2
}

// normalizeEndpoint ensures the endpoint has the https:// scheme prefix.
// Gen2 endpoints already include the scheme, but Gen1 endpoints do not.
func normalizeEndpoint(endpoint string) string {
	if endpoint == "" {
		return ""
	}
	if strings.HasPrefix(endpoint, "https://") || strings.HasPrefix(endpoint, "http://") {
		return endpoint
	}
	return "https://" + endpoint
}

// getCloudantClientForUrl creates an authenticated Cloudant SDK client for the given endpoint.
// It handles both bearer token and API key authentication, and configures the IAM URL
// based on visibility settings and instance generation.
//
// Gen 2 instances use the VPE private IAM URL: https://private.iam.cloud.ibm.com
// Gen 1 instances use regional private IAM URLs: https://private.{region}.iam.cloud.ibm.com
//
// This function is shared between cloudant_resource and cloudant_resource_database.
func getCloudantClientForUrl(endpoint string, isGen2 bool, meta interface{}) (*cloudantv1.CloudantV1, error) {
	session, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}

	var authenticator core.Authenticator
	token := session.Config.IAMAccessToken

	if token != "" {
		token = strings.Replace(token, "Bearer ", "", -1)
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: token,
		}
	} else {
		apiKey := session.Config.BluemixAPIKey
		region := session.Config.Region
		visibility := session.Config.Visibility
		iamURL := iamidentityv1.DefaultServiceURL
		if visibility == "private" || visibility == "public-and-private" {
			if isGen2 {
				// Gen 2 uses VPE private IAM URL
				iamURL = conns.ContructEndpoint("private.iam", "cloud.ibm.com")
			} else {
				// Gen 1 uses regional private IAM URLs: private.{region}.iam.cloud.ibm.com
				iamURL = conns.ContructEndpoint(fmt.Sprintf("private.%s.iam", region), "cloud.ibm.com")
			}
		}
		authenticator = &core.IamAuthenticator{
			ApiKey: apiKey,
			URL:    conns.EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL) + "/identity/token",
		}
	}

	client, err := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
		Authenticator: authenticator,
		URL:           endpoint,
	})
	if err != nil {
		return nil, flex.FmtErrorf("[ERROR] Error occured while configuring Cloudant service: %q", err)
	}
	client.Service.SetUserAgent("cloudant-terraform/" + version.Version)

	return client, nil
}

// cloudantToResourceInstance mutates d in-place to prepare it for an RC CRUD
// call. Cloudant-only schema attributes are mapped into the RC transport fields:
//
//   - Gen 1: legacy_credentials and environment_crn are merged into parameters
//     alongside any user-supplied keys. Typed schema attributes win over
//     any conflicting user-supplied values.
//   - Gen 2: capacity, enable_cors, and cors_config are encoded as a nested JSON
//     object and written to parameters_json, which the RC layer unmarshals
//     into the correct nested broker payload. User-supplied parameters are
//     folded into the same JSON object; typed schema attributes take
//     precedence over any conflicting user-supplied dataservices.cloudant.*
//     keys.
func cloudantToResourceInstance(d *schema.ResourceData) error {
	if isCloudantGen2PlanFrom(d) {
		// Seed from user-supplied parameters so that typed fields written below
		// always take precedence over any conflicting user-supplied values.
		paramsJSON, err := cloudantGen2ParamsAsJSON(seedParamsFromExisting(d), d)
		if err != nil {
			return err
		}
		if err := d.Set("parameters_json", paramsJSON); err != nil {
			return fmt.Errorf("error setting parameters_json: %s", err)
		}
	} else {
		// Gen 1: merge Cloudant-only fields into parameters. Start with any
		// user-supplied keys so that typed fields written below take precedence.
		params := seedParamsFromExisting(d)
		// Always send legacyCredentials explicitly — the server default is true,
		// so omitting it when the user set false would silently enable it.
		params["legacyCredentials"] = fmt.Sprintf("%t", d.Get("legacy_credentials").(bool))
		if environmentCRN, ok := d.GetOk("environment_crn"); ok {
			params["environment_crn"] = environmentCRN
		}
		if err := d.Set("parameters", params); err != nil {
			return fmt.Errorf("error setting parameters: %s", err)
		}
	}
	return nil
}

// resourceInstanceToCloudant mutates d in-place after an RC CRUD call, mapping
// RC response fields back into Cloudant-specific schema attributes and cleaning
// up the transport fields so they do not leak into state:
//
//   - Gen 1: extracts legacy_credentials and environment_crn from the parameters
//     map (populated before the RC call by cloudantToResourceInstance),
//     then removes those keys from parameters so only user-supplied keys
//     remain in state.
//   - Gen 2: reads capacity, enable_cors, and cors_config from the extensions
//     flat TypeMap (populated by the RC read);
//     throughput is Gen 1-only so it is set to an empty map.
func resourceInstanceToCloudant(d *schema.ResourceData) {
	if isCloudantGen2PlanFrom(d) {
		if err := d.Set("throughput", map[string]interface{}{}); err != nil {
			log.Printf("[WARN] error clearing throughput for Gen 2 instance: %s", err)
		}

		ext, ok := d.Get("extensions").(map[string]interface{})
		if !ok || len(ext) == 0 {
			return
		}

		if capacityStr, ok := extStr(ext, "dataservices.cloudant.capacity_units"); ok {
			if capacity, err := strconv.Atoi(capacityStr); err == nil {
				if err := d.Set("capacity", capacity); err != nil {
					log.Printf("[WARN] error setting capacity: %s", err)
				}
			}
		}

		if dataEventsStr, ok := extStr(ext, "dataservices.cloudant.configuration.audit.data_events"); ok {
			includeDataEvents := dataEventsStr == "true"
			if err := d.Set("include_data_events", includeDataEvents); err != nil {
				log.Printf("[WARN] error setting include_data_events: %s", err)
			}
		}

		if enabledStr, ok := extStr(ext, "dataservices.cloudant.configuration.cors.enabled"); ok {
			enabled := enabledStr == "true"
			if err := d.Set("enable_cors", enabled); err != nil {
				log.Printf("[WARN] error setting enable_cors: %s", err)
			}
			if enabled {
				corsState := map[string]interface{}{
					"allow_credentials": false,
					"origins":           []string{},
				}
				if acStr, ok := extStr(ext, "dataservices.cloudant.configuration.cors.allowCredentials"); ok {
					corsState["allow_credentials"] = acStr == "true"
				}
				if countStr, ok := extStr(ext, "dataservices.cloudant.configuration.cors.origins.#"); ok {
					if count, err := strconv.Atoi(countStr); err == nil && count > 0 {
						origins := make([]string, 0, count)
						for i := 0; i < count; i++ {
							key := fmt.Sprintf("dataservices.cloudant.configuration.cors.origins.%d", i)
							if origin, ok := extStr(ext, key); ok {
								origins = append(origins, origin)
							}
						}
						corsState["origins"] = origins
					}
				}
				if err := d.Set("cors_config", []map[string]interface{}{corsState}); err != nil {
					log.Printf("[WARN] error setting cors_config: %s", err)
				}
			}
		}
	} else {
		// Gen 1: extract Cloudant-specific fields from parameters and remove
		// them so only user-supplied keys remain in state.
		parameters, ok := d.Get("parameters").(map[string]interface{})
		if !ok || parameters == nil {
			return
		}

		if legacyCredentials, ok := parameters["legacyCredentials"]; ok {
			var val bool
			switch v := legacyCredentials.(type) {
			case bool:
				val = v
			case string:
				val = v == "true"
			}
			if err := d.Set("legacy_credentials", val); err != nil {
				log.Printf("[WARN] error setting legacy_credentials: %s", err)
			}
			delete(parameters, "legacyCredentials")
		}

		if crn, ok := parameters["environment_crn"].(string); ok {
			if err := d.Set("environment_crn", crn); err != nil {
				log.Printf("[WARN] error setting environment_crn: %s", err)
			}
			delete(parameters, "environment_crn")
		}

		// Write back the cleaned parameters map (user keys only).
		// Always call Set so the SDK sees the mutation; use an empty map
		// rather than nil so the TypeMap is explicitly cleared.
		if err := d.Set("parameters", parameters); err != nil {
			log.Printf("[WARN] error setting parameters: %s", err)
		}
	}
}

// resourceDataGetter is satisfied by both *schema.ResourceData and
// *schema.ResourceDiff, which share a Get(string) interface{} method.
type resourceDataGetter interface {
	Get(string) interface{}
}

// cloudantGen2ParamsAsJSON builds the nested Cloudant Gen 2 broker payload
// (capacity_units and cors configuration) from rd, merges it with any
// user-supplied parameters already present in params, and returns the result
// as a marshalled JSON string. Typed schema attributes always overwrite
// any conflicting user-supplied dataservices.cloudant.* keys.
//
// This function is called both from cloudantToResourceInstance (at CRUD time)
// and from the ResourceIBMCloudant CustomizeDiff closure (at plan time) so that
// the two places never drift out of sync.
func cloudantGen2ParamsAsJSON(params map[string]interface{}, rd resourceDataGetter) (string, error) {
	dataservices, _ := params["dataservices"].(map[string]interface{})
	if dataservices == nil {
		dataservices = map[string]interface{}{}
		params["dataservices"] = dataservices
	}
	cloudant, _ := dataservices["cloudant"].(map[string]interface{})
	if cloudant == nil {
		cloudant = map[string]interface{}{}
		dataservices["cloudant"] = cloudant
	}
	cloudant["capacity_units"] = rd.Get("capacity").(int)

	configuration, _ := cloudant["configuration"].(map[string]interface{})
	if configuration == nil {
		configuration = map[string]interface{}{}
		cloudant["configuration"] = configuration
	}
	enableCors := rd.Get("enable_cors").(bool)
	corsConfig := map[string]interface{}{"enabled": enableCors}
	if enableCors {
		corsConfig["allowCredentials"] = true
		corsConfig["origins"] = []string{}
		if corsConfigRaw := rd.Get("cors_config").([]interface{}); len(corsConfigRaw) > 0 {
			m := corsConfigRaw[0].(map[string]interface{})
			corsConfig["allowCredentials"] = m["allow_credentials"].(bool)
			corsConfig["origins"] = flex.ExpandStringList(m["origins"].([]interface{}))
		}
	}
	configuration["cors"] = corsConfig

	auditConfig := map[string]interface{}{
		"data_events": rd.Get("include_data_events").(bool),
	}
	configuration["audit"] = auditConfig

	b, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("error marshalling Gen 2 parameters: %s", err)
	}
	return string(b), nil
}

// seedParamsFromExisting returns a new map pre-seeded with any user-supplied
// parameters from d, so that typed schema attributes written afterwards
// always win over any conflicting user-provided values.
func seedParamsFromExisting(d *schema.ResourceData) map[string]interface{} {
	params := make(map[string]interface{})
	if existing, ok := d.GetOk("parameters"); ok {
		for k, v := range existing.(map[string]interface{}) {
			params[k] = v
		}
	}
	return params
}

// extStr returns the string value for key from a flat extensions map,
// reporting false if the key is absent or its value is not a non-empty string.
func extStr(ext map[string]interface{}, key string) (string, bool) {
	v, ok := ext[key]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok && s != ""
}

// Made with Bob
