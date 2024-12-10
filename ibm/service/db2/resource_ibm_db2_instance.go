// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package db2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	RsInstanceSuccessStatus       = "active"
	RsInstanceProgressStatus      = "in progress"
	RsInstanceProvisioningStatus  = "provisioning"
	RsInstanceInactiveStatus      = "inactive"
	RsInstanceFailStatus          = "failed"
	RsInstanceRemovedStatus       = "removed"
	RsInstanceReclamation         = "pending_reclamation"
	RsInstanceUpdateSuccessStatus = "succeeded"
)

func ResourceIBMDb2Instance() *schema.Resource {
	riSchema := resourcecontroller.ResourceIBMResourceInstance().Schema

	riSchema["high_availability"] = &schema.Schema{
		Description: "If you require high availability, please choose this option",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["instance_type"] = &schema.Schema{
		Description: "Available machine type flavours (default selection will assume smallest configuration)",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["backup_location"] = &schema.Schema{
		Description: "Cross Regional backups can be stored across multiple regions in a zone. Regional backups are stored in only specific region.",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["disk_encryption_instance_crn"] = &schema.Schema{
		Description: "Cross Regional disk encryption crn",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["disk_encryption_crn"] = &schema.Schema{
		Description: "Cross Regional disk encryption crn",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["oracle_compatibility"] = &schema.Schema{
		Description: "Indicates whether is has compatibility for oracle or not",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["subscription_id"] = &schema.Schema{
		Description: "Subscription ID",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["autoscale_config"] = &schema.Schema{
		Description: "Autoscaling configurations of created Db2 instance",
		Optional:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"auto_scaling_allow_plan_limit": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Optional:    true,
					Description: "Indicates the maximum number of scaling actions that are allowed within a specified time period.",
				},
				"auto_scaling_enabled": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Optional:    true,
					Description: "Indicates if automatic scaling is enabled or not.",
				},
				"auto_scaling_over_time_period": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Optional:    true,
					Description: "Defines the time period over which auto-scaling adjustments are monitored and applied.",
				},
				"auto_scaling_pause_limit": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Optional:    true,
					Description: "Specifies the duration to pause auto-scaling actions after a scaling event has occurred.",
				},
				"auto_scaling_threshold": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Optional:    true,
					Description: "Specifies the resource utilization level that triggers an auto-scaling.",
				},
			},
		},
	}

	riSchema["whitelist_config"] = &schema.Schema{
		Description: "Whitelists configurations of created Db2 instance",
		Optional:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_addresses": &schema.Schema{
					Type:        schema.TypeList,
					Computed:    true,
					Description: "List of IP addresses.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"address": &schema.Schema{
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The IP address, in IPv4/ipv6 format.",
							},
							"description": &schema.Schema{
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Description of the IP address.",
							},
						},
					},
				},
			},
		},
	}

	return &schema.Resource{
		Create:   resourceIBMDb2InstanceCreate,
		Read:     resourcecontroller.ResourceIBMResourceInstanceRead,
		Update:   resourcecontroller.ResourceIBMResourceInstanceUpdate,
		Delete:   resourcecontroller.ResourceIBMResourceInstanceDelete,
		Exists:   resourcecontroller.ResourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),

		Schema: riSchema,
	}
}

func resourceIBMDb2InstanceCreate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}

	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := rc.CreateResourceInstanceOptions{
		Name: &name,
	}

	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving service offering: %s", err)
	}

	if metadata, ok := serviceOff[0].Metadata.(*models.ServiceResourceMetadata); ok {
		if !metadata.Service.RCProvisionable {
			return fmt.Errorf("%s cannot be provisioned by resource controller", serviceName)
		}
	} else {
		return fmt.Errorf("[ERROR] Cannot create instance of resource %s\nUse 'ibm_service_instance' if the resource is a Cloud Foundry service", serviceName)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving plan: %s", err)
	}
	rsInst.ResourcePlanID = &servicePlan

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving deployment for plan %s : %s", plan, err)
	}
	if len(deployments) == 0 {
		return fmt.Errorf("[ERROR] No deployment found for service plan : %s", plan)
	}
	deployments, supportedLocations := resourcecontroller.FilterDeployments(deployments, location)

	if len(deployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		return fmt.Errorf("[ERROR] No deployment found for service plan %s at location %s.\nValid location(s) are: %q.\nUse 'ibm_service_instance' if the service is a Cloud Foundry service", plan, location, locationList)
	}

	rsInst.Target = &deployments[0].CatalogCRN

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rg := rsGrpID.(string)
		rsInst.ResourceGroup = &rg
	} else {
		defaultRg, err := flex.DefaultResourceGroup(meta)
		if err != nil {
			return err
		}
		rsInst.ResourceGroup = &defaultRg
	}

	params := map[string]interface{}{}

	if serviceEndpoints, ok := d.GetOk("service_endpoints"); ok {
		params["service-endpoints"] = serviceEndpoints.(string)
	}
	if highAvailability, ok := d.GetOk("high_availability"); ok {
		params["high_availability"] = highAvailability.(string)
	}
	if instanceType, ok := d.GetOk("instance_type"); ok {
		params["instance_type"] = instanceType.(string)
	}
	if backupLocation, ok := d.GetOk("backup_location"); ok {
		params["backup-locations"] = backupLocation.(string)
	}

	if diskEncryptionInstanceCrn, ok := d.GetOk("disk_encryption_instance_crn"); ok {
		params["disk_encryption_instance_crn"] = diskEncryptionInstanceCrn.(string)
	}

	if diskEncryptionKeyCrn, ok := d.GetOk("disk_encryption_key_crn"); ok {
		params["disk_encryption_key_crn"] = diskEncryptionKeyCrn.(string)
	}

	if oracleCompatibility, ok := d.GetOk("oracle_compatibility"); ok {
		params["oracle_compatibility"] = oracleCompatibility.(string)
	}

	if subscriptionId, ok := d.GetOk("subscription_id"); ok {
		params["subscription_id"] = subscriptionId.(string)
	}

	if parameters, ok := d.GetOk("parameters"); ok {
		temp := parameters.(map[string]interface{})
		for k, v := range temp {
			if v == "true" || v == "false" {
				b, _ := strconv.ParseBool(v.(string))
				params[k] = b
			} else if strings.HasPrefix(v.(string), "[") && strings.HasSuffix(v.(string), "]") {
				//transform v.(string) to be []string
				arrayString := v.(string)
				result := []string{}
				trimLeft := strings.TrimLeft(arrayString, "[")
				trimRight := strings.TrimRight(trimLeft, "]")
				if len(trimRight) == 0 {
					params[k] = result
				} else {
					array := strings.Split(trimRight, ",")
					for _, a := range array {
						result = append(result, strings.Trim(a, "\""))
					}
					params[k] = result
				}
			} else {
				params[k] = v
			}
		}

	}

	if s, ok := d.GetOk("parameters_json"); ok {
		json.Unmarshal([]byte(s.(string)), &params)
	}

	rsInst.Parameters = params

	//Start to create resource instance
	instance, resp, err := rsConClient.CreateResourceInstance(&rsInst)
	if err != nil {
		log.Printf(
			"Error when creating resource instance: %s, Instance info  NAME->%s, LOCATION->%s, GROUP_ID->%s, PLAN_ID->%s",
			err, *rsInst.Name, *rsInst.Target, *rsInst.ResourceGroup, *rsInst.ResourcePlanID)
		return fmt.Errorf("[ERROR] Error when creating resource instance: %s with resp code: %s", err, resp)
	}

	d.SetId(*instance.ID)

	_, err = waitForResourceInstanceCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for create resource instance (%s) to be succeeded: %s", d.Id(), err)
	}

	log.Printf("Instance ID %s", *instance.ID)
	log.Printf("Instance CRN %s", *instance.CRN)
	log.Printf("Instance URL %s", *instance.URL)
	log.Printf("Instance DashboardURL %s", *instance.DashboardURL)

	//Create client for db2SaasV1
	db2SaasV1Client, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		return err
	}

	if whitelistConfigRaw, ok := d.GetOk("whitelist_config"); ok {
		if whitelistConfigRaw == nil || reflect.ValueOf(whitelistConfigRaw).IsNil() {
			fmt.Println("No whitelisting config provided; skipping.")
		} else {
			whitelistConfig := whitelistConfigRaw.([]interface{})[0].(map[string]interface{})
			fmt.Println(whitelistConfig)
			fmt.Println(whitelistConfig["ip_addresses"].([]interface{}))

			ipAddress := make([]db2saasv1.IpAddress, 0, len(whitelistConfig["ip_addresses"].([]interface{})))

			for _, ip := range ipAddress {
				if err = validateIPAddress(ip); err != nil {
					return err
				}
			}

			input := &db2saasv1.PostDb2SaasWhitelistOptions{
				XDeploymentID: core.StringPtr(*instance.CRN),
				IpAddresses:   ipAddress,
			}

			result, response, err := db2SaasV1Client.PostDb2SaasWhitelist(input)
			if err != nil {
				log.Printf("Error when posting whitelist to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}
	}

	if autoscaleConfigRaw, ok := d.GetOk("autoscale_config"); ok {
		if autoscaleConfigRaw == nil || reflect.ValueOf(autoscaleConfigRaw).IsNil() {
			fmt.Println("No autoscale config provided; skipping.")
		} else {
			autoscalingConfig := autoscaleConfigRaw.([]interface{})[0].(map[string]interface{})
			fmt.Println(autoscalingConfig)

			autoScalingThreshold, err := strconv.Atoi(autoscalingConfig["auto_scaling_threshold"].(string))
			if err != nil {
				return err
			}

			autoScalingOverTimePeriod, err := strconv.Atoi(autoscalingConfig["auto_scaling_over_time_period"].(string))
			if err != nil {
				return err
			}

			autoScalingPauseLimit, err := strconv.Atoi(autoscalingConfig["auto_scaling_pause_limit"].(string))
			if err != nil {
				return err
			}

			if len(autoscalingConfig) > 1 {
				autoscalingConfig["auto_scaling_enabled"] = true
			} else {
				autoscalingConfig["auto_scaling_enabled"] = false
			}

			input := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDbProfile:                core.StringPtr(*instance.CRN),
				AutoScalingEnabled:        core.StringPtr(autoscalingConfig["auto_scaling_enabled"].(string)),
				AutoScalingAllowPlanLimit: core.StringPtr(autoscalingConfig["auto_scaling_allow_plan_limit"].(string)),
				AutoScalingThreshold:      core.Int64Ptr(int64(autoScalingThreshold)),
				AutoScalingOverTimePeriod: core.Float64Ptr(float64(autoScalingOverTimePeriod)),
				AutoScalingPauseLimit:     core.Int64Ptr(int64(autoScalingPauseLimit)),
			}

			result, response, err := db2SaasV1Client.PutDb2SaasAutoscale(input)
			if err != nil {
				log.Printf("Error when posting whitelist to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateTagsUsingCRN(oldList, newList, meta, *instance.CRN)
		if err != nil {
			log.Printf(
				"Error on create of resource instance (%s) tags: %s", d.Id(), err)
		}
	}

	return resourcecontroller.ResourceIBMResourceInstanceRead(d, meta)
}

func waitForResourceInstanceCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()
	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	stateConf := &retry.StateChangeConf{
		Pending: []string{RsInstanceProgressStatus, RsInstanceInactiveStatus, RsInstanceProvisioningStatus},
		Target:  []string{RsInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
			if err != nil {
				if resp != nil && resp.StatusCode == 404 {
					return nil, "", fmt.Errorf("[ERROR] The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", fmt.Errorf("[ERROR] Get the resource instance %s failed with resp code: %s, err: %v", d.Id(), resp, err)
			}
			if *instance.State == RsInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("[ERROR] The resource instance '%s' creation failed: %v", d.Id(), err)
			}
			return instance, *instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      30 * time.Second,
		MinTimeout: 30 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func validateIPAddress(ip db2saasv1.IpAddress) error {
	if ip.Address == nil || *ip.Address == "" {
		return fmt.Errorf("[ERROR] IP address is required")
	}

	if ip.Description == nil || *ip.Description == "" {
		return fmt.Errorf("[ERROR] IP address is required")
	}

	return nil
}
