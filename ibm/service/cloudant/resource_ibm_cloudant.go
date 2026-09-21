// Copyright IBM Corp. 2021, 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant

import (
	"context"
	"log"
	"maps"
	"net/url"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

// suppressMissingCorsConfig suppresses a diff when cors_config is absent from
// the configuration. Defaults are applied during update, so an omitted
// cors_config block should not be treated as a removal.
func suppressMissingCorsConfig(k, _, new string, _ *schema.ResourceData) bool {
	return k == "cors_config.#" && new == "0"
}

func ResourceIBMCloudant() *schema.Resource {
	// Clone to avoid mutating the shared resource-controller schema map.
	cloudantResourceInstanceSchema := maps.Clone(resourcecontroller.ResourceIBMResourceInstance().Schema)

	// parameters_json is used internally to ferry nested Gen 2 broker parameters
	// (dataservices.cloudant.*) through the RC layer without TypeMap string corruption.
	// It is not a user-facing attribute: mark it Computed-only so it cannot be set
	// in HCL config.
	pj := cloudantResourceInstanceSchema["parameters_json"]
	cloudantResourceInstanceSchema["parameters_json"] = &schema.Schema{
		Type:        pj.Type,
		Computed:    true,
		Description: pj.Description,
	}

	// Override: service is hardcoded to "cloudantnosqldb" in create/update,
	// so it must be Computed-only here rather than Required as in the base schema.
	cloudantResourceInstanceSchema["service"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The service type of the instance",
		Computed:    true,
	}

	cloudantResourceInstanceSchema["legacy_credentials"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Only available for IBM Cloudant Gen 1. Use both legacy credentials and IAM for authentication.",
		ForceNew:    true,
	}

	cloudantResourceInstanceSchema["environment_crn"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Only available for IBM Cloudant Gen 1. CRN of the IBM Cloudant Dedicated Hardware plan instance.",
		ForceNew:    true,
	}

	cloudantResourceInstanceSchema["include_data_events"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Include data event types in events sent to IBM Cloud Activity Tracker Event Routing for the IBM Cloudant instance. By default only emitted events are of \"management\" type. For Gen 1 instances this is applied via the Activity Tracker API; for Gen 2 instances it is set via the broker parameter `dataservices.cloudant.configuration.audit.data_events`.",
	}

	cloudantResourceInstanceSchema["capacity"] = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		Description:  "A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.",
		ValidateFunc: validation.IntAtLeast(1),
	}

	cloudantResourceInstanceSchema["throughput"] = &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes. This is only available for IBM Cloudant Gen 1.",
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
	}

	cloudantResourceInstanceSchema["enable_cors"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Boolean value to turn CORS on and off.",
	}

	cloudantResourceInstanceSchema["cors_config"] = &schema.Schema{
		Type:             schema.TypeList,
		Optional:         true,
		DiffSuppressFunc: suppressMissingCorsConfig,
		Description:      "Configuration for CORS.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_credentials": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.",
				},
				"origins": {
					Type:        schema.TypeList,
					Required:    true,
					Description: "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
		MinItems: 1,
		MaxItems: 1,
	}

	return &schema.Resource{
		Create: resourceIBMCloudantCreate,
		Read:   resourceIBMCloudantRead,
		Update: resourceIBMCloudantUpdate,
		Delete: resourcecontroller.ResourceIBMResourceInstanceDelete,
		Exists: resourcecontroller.ResourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				if err := resourceIBMCloudantRead(d, meta); err != nil {
					return nil, err
				}
				if _, ok := d.GetOk("legacy_credentials"); !ok {
					if err := d.Set("legacy_credentials", false); err != nil {
						return nil, err
					}
				}
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
			func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
				if err := validateCloudantInstanceCapacity(diff); err != nil {
					return err
				}
				if err := validateCloudantInstanceCors(diff); err != nil {
					return err
				}
				return nil
			},
			// For Gen 2 plans, propagate capacity/CORS changes into parameters_json
			// at plan time so that HasChange("parameters_json") is true during Update
			// and the RC layer actually sends the new parameters to the broker.
			// Also needs to mark the extensions as updating so they become part of
			// the plan delta (see https://github.com/hashicorp/terraform/issues/36462)
			func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
				if !isCloudantGen2PlanFrom(diff) {
					return nil
				}
				if !diff.HasChange("capacity") && !diff.HasChange("enable_cors") && !diff.HasChange("cors_config") && !diff.HasChange("include_data_events") {
					return nil
				}
				if err := diff.SetNewComputed("extensions"); err != nil {
					return err
				}
				paramsJSON, err := cloudantGen2ParamsAsJSON(map[string]interface{}{}, diff)
				if err != nil {
					return err
				}
				return diff.SetNew("parameters_json", paramsJSON)
			},
		),

		Schema: cloudantResourceInstanceSchema,
	}
}

func resourceIBMCloudantCreate(d *schema.ResourceData, meta interface{}) error {
	d.Set("service", "cloudantnosqldb")

	if err := cloudantToResourceInstance(d); err != nil {
		return err
	}
	if err := resourcecontroller.ResourceIBMResourceInstanceCreate(d, meta); err != nil {
		return err
	}

	if !isCloudantGen2PlanFrom(d) {

		client, tfErr := GetCloudantClientFromResource(d, meta, "ibm_cloudant", "create")
		if tfErr != nil {
			return tfErr
		}

		// if matches an instance creation default skip request
		if d.Get("include_data_events").(bool) {
			if err := updateCloudantActivityTrackerEvents(client, d); err != nil {
				return flex.FmtErrorf("[ERROR] Error updating activity tracker events: %s", err)
			}
		}

		// if matches an instance creation default skip request
		if d.Get("capacity").(int) > 1 {
			if err := updateCloudantInstanceCapacity(client, d); err != nil {
				return flex.FmtErrorf("[ERROR] Error retrieving capacity throughput information: %s", err)
			}
		}

		if err := updateCloudantInstanceCors(client, d); err != nil {
			return flex.FmtErrorf("[ERROR] Error updating CORS settings: %s", err)
		}
	}

	return resourceIBMCloudantRead(d, meta)
}

func resourceIBMCloudantRead(d *schema.ResourceData, meta interface{}) error {
	if err := resourcecontroller.ResourceIBMResourceInstanceRead(d, meta); err != nil {
		return err
	}

	if err := setCloudantResourceControllerURL(d, meta); err != nil {
		return err
	}

	resourceInstanceToCloudant(d)

	if !isCloudantGen2PlanFrom(d) {

		// Gen 1: API calls to fix up Cloudant-specific values.

		client, tfErr := GetCloudantClientFromResource(d, meta, "ibm_cloudant", "read")
		if tfErr != nil {
			return tfErr
		}

		if err := setCloudantActivityTrackerEvents(client, d); err != nil {
			return err
		}
		if err := setCloudantInstanceCapacity(client, d); err != nil {
			return err
		}
		if err := setCloudantInstanceCors(client, d); err != nil {
			return err
		}
	}

	return nil
}

func resourceIBMCloudantUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Set("service", "cloudantnosqldb")
	isGen2 := isCloudantGen2PlanFrom(d)

	if err := cloudantToResourceInstance(d); err != nil {
		return err
	}
	if err := resourcecontroller.ResourceIBMResourceInstanceUpdate(d, meta); err != nil {
		return err
	}

	if !isGen2 {
		client, tfErr := GetCloudantClientFromResource(d, meta, "ibm_cloudant", "update")
		if tfErr != nil {
			return tfErr
		}

		if d.HasChange("include_data_events") {
			err := updateCloudantActivityTrackerEvents(client, d)
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error updating activity tracker events: %s", err)
			}
		}

		if d.HasChange("capacity") {
			err := updateCloudantInstanceCapacity(client, d)
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error retrieving capacity throughput information: %s", err)
			}
		}

		if d.HasChange("enable_cors") || d.HasChange("cors_config") {
			err := updateCloudantInstanceCors(client, d)
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error updating CORS settings: %s", err)
			}
		}
	}

	return resourceIBMCloudantRead(d, meta)
}

func setCloudantResourceControllerURL(d *schema.ResourceData, meta interface{}) error {
	crn := d.Get(flex.ResourceCRN).(string)
	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/cloudantnosqldb/"+url.QueryEscape(crn))

	return nil
}

func validateCloudantGen2UnsupportedFields(diff *schema.ResourceDiff) error {
	if !isCloudantGen2PlanFrom(diff) {
		return nil
	}

	if diff.Get("legacy_credentials").(bool) {
		return flex.FmtErrorf("[ERROR] Setting legacy_credentials is not supported for Gen 2 Cloudant instances")
	}

	if _, ok := diff.GetOk("environment_crn"); ok {
		return flex.FmtErrorf("[ERROR] Setting environment_crn is not supported for Gen 2 Cloudant instances")
	}

	return nil
}

func setCloudantActivityTrackerEvents(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	activityTrackerEvents, err := readCloudantActivityTrackerEvents(client)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving activity tracker events: %s", err)
	}
	if activityTrackerEvents.Types != nil {
		includeDataEvents := false
		for _, t := range activityTrackerEvents.Types {
			if t == "data" {
				includeDataEvents = true
			}
		}
		d.Set("include_data_events", includeDataEvents)
	}
	return nil
}

func readCloudantActivityTrackerEvents(client *cloudantv1.CloudantV1) (*cloudantv1.ActivityTrackerEvents, error) {
	opts := client.NewGetActivityTrackerEventsOptions()

	activityTrackerEvents, response, err := client.GetActivityTrackerEvents(opts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving activity tracker events: %s\n%s", err, response)
	}
	return activityTrackerEvents, err
}

func updateCloudantActivityTrackerEvents(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	auditEventTypes := []string{"management"}
	if d.Get("include_data_events").(bool) {
		auditEventTypes = append(auditEventTypes, "data")
	}

	opts := client.NewPostActivityTrackerEventsOptions(auditEventTypes)

	_, response, err := client.PostActivityTrackerEvents(opts)
	if err != nil {
		log.Printf("[DEBUG] Error updating activity tracker events: %s\n%s", err, response)
	}
	return err
}

func validateCloudantInstanceCapacity(diff *schema.ResourceDiff) error {
	if err := validateCloudantGen2UnsupportedFields(diff); err != nil {
		return err
	}

	plan := diff.Get("plan").(string)
	capacity := diff.Get("capacity").(int)
	if capacity > 1 && plan == "lite" {
		return flex.FmtErrorf("[ERROR] Setting capacity is not supported for your instance's plan")
	}
	return nil
}

func setCloudantInstanceCapacity(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	capacityThroughputInformation, err := readCloudantInstanceCapacity(client)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving capacity throughput information: %s", err)
	}

	if capacityThroughputInformation.Current != nil && capacityThroughputInformation.Current.Throughput != nil {
		currentThroughput := capacityThroughputInformation.Current.Throughput
		targetThroughput := currentThroughput
		if capacityThroughputInformation.Target != nil && capacityThroughputInformation.Target.Throughput != nil {
			targetThroughput = capacityThroughputInformation.Target.Throughput
		}
		// lite plan doesn't have "blocks" attr on broker's response
		if d.Get("plan").(string) == "lite" || currentThroughput.Blocks == nil {
			d.Set("capacity", 1)
		} else {
			blocks := int(*targetThroughput.Blocks)
			d.Set("capacity", blocks)
		}
		throughput := map[string]int{
			"query": int(*targetThroughput.Query),
			"read":  int(*targetThroughput.Read),
			"write": int(*targetThroughput.Write),
		}
		d.Set("throughput", throughput)
	}
	return nil
}

func readCloudantInstanceCapacity(client *cloudantv1.CloudantV1) (*cloudantv1.CapacityThroughputInformation, error) {
	opts := client.NewGetCapacityThroughputInformationOptions()

	capacityThroughputInformation, response, err := client.GetCapacityThroughputInformation(opts)
	if err != nil {
		log.Printf("[DEBUG] Error getting capacity throughput information: %s\n%s", err, response)
		return nil, err
	}
	return capacityThroughputInformation, nil
}

func updateCloudantInstanceCapacity(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	blocks := int64(d.Get("capacity").(int))

	putOpts := client.NewPutCapacityThroughputConfigurationOptions(blocks)

	_, response, err := client.PutCapacityThroughputConfiguration(putOpts)
	if err != nil {
		log.Printf("[DEBUG] Error updating capacity throughput: %s\n%s", err, response)
		return err
	}

	return nil
}

func validateCloudantInstanceCors(diff *schema.ResourceDiff) error {
	enableCors := diff.Get("enable_cors").(bool)
	corsConfigRaw := diff.Get("cors_config").([]interface{})
	if !enableCors && len(corsConfigRaw) > 0 {
		corsConfig := corsConfigRaw[0].(map[string]interface{})
		allowCredentials := corsConfig["allow_credentials"].(bool)
		origins := corsConfig["origins"].([]interface{})
		if !allowCredentials || len(origins) > 0 {
			return flex.FmtErrorf("[ERROR] Setting \"cors_config\" conflicts with enable_cors set to false")
		}
	}
	return nil
}

func setCloudantInstanceCors(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	corsInformation, err := readCloudantInstanceCors(client)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving CORS config: %s", err)
	}
	if corsInformation != nil {
		d.Set("enable_cors", corsInformation.EnableCors)

		if *corsInformation.EnableCors {
			corsConfig := []map[string]interface{}{
				map[string]interface{}{
					"allow_credentials": corsInformation.AllowCredentials,
					"origins":           corsInformation.Origins,
				},
			}
			d.Set("cors_config", corsConfig)
		}
	}
	return nil
}

func readCloudantInstanceCors(client *cloudantv1.CloudantV1) (*cloudantv1.CorsInformation, error) {
	opts := client.NewGetCorsInformationOptions()

	corsInformation, response, err := client.GetCorsInformation(opts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving CORS config: %s\n%s", err, response)
	}
	return corsInformation, err
}

func updateCloudantInstanceCors(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	enableCors := d.Get("enable_cors").(bool)
	allowCredentials := true
	origins := make([]string, 0)
	corsConfigRaw := d.Get("cors_config").([]interface{})
	if enableCors && len(corsConfigRaw) > 0 {
		corsConfig := corsConfigRaw[0].(map[string]interface{})
		allowCredentials = corsConfig["allow_credentials"].(bool)
		origins = flex.ExpandStringList(corsConfig["origins"].([]interface{}))
	}

	opts := client.NewPutCorsConfigurationOptions(origins)
	opts.SetEnableCors(enableCors)
	opts.SetAllowCredentials(allowCredentials)

	_, response, err := client.PutCorsConfiguration(opts)
	if err != nil {
		log.Printf("[DEBUG] Error updating CORS settings: %s\n%s", err, response)
	}
	return err
}
