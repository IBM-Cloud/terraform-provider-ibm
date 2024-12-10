// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package db2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"log"
	"net/url"
	"reflect"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMDb2Instance() *schema.Resource {
	riSchema := resourcecontroller.DataSourceIBMResourceInstance().Schema

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

	riSchema["autoscaling_config"] = &schema.Schema{
		Description: "Autoscaling configurations of the created db2 instance",
		Optional:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"auto_scaling_enabled": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Indicates if automatic scaling is enabled or not.",
				},
				"auto_scaling_allow_plan_limit": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Indicates the maximum number of scaling actions that are allowed within a specified time period.",
				},
				"auto_scaling_max_storage": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "The maximum limit for automatically increasing storage capacity to handle growing data needs.",
				},
				"auto_scaling_over_time_period": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Defines the time period over which auto-scaling adjustments are monitored and applied.",
				},
				"auto_scaling_pause_limit": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Specifies the duration to pause auto-scaling actions after a scaling event has occurred.",
				},
				"auto_scaling_threshold": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Specifies the resource utilization level that triggers an auto-scaling.",
				},
				"storage_unit": &schema.Schema{
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Specifies the unit of measurement for storage capacity.",
				},
				"storage_utilization_percentage": &schema.Schema{
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Represents the percentage of total storage capacity currently in use.",
				},
				"support_auto_scaling": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Indicates whether a system or service can automatically adjust resources based on demand.",
				},
			},
		},
	}

	riSchema["whitelist_config"] = &schema.Schema{
		Description: "Whitelists configurations of the created db2 instance",
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

	riSchema["connectioninfo_config"] = &schema.Schema{
		Description: "Connection info configurations of the created db2 instance",
		Optional:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"public": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"hostname": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"database_name": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"ssl_port": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"ssl": &schema.Schema{
								Type:     schema.TypeBool,
								Computed: true,
							},
							"database_version": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
				"private": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"hostname": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"database_name": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"ssl_port": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"ssl": &schema.Schema{
								Type:     schema.TypeBool,
								Computed: true,
							},
							"database_version": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"private_service_name": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"cloud_service_offering": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"vpe_service_crn": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"db_vpc_endpoint_service": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
		},
	}

	return &schema.Resource{
		Read:   dataSourceIBMDb2InstanceRead,
		Schema: riSchema,
	}
}

func getInstancesNext(next *string) (string, error) {
	if reflect.ValueOf(next).IsNil() {
		return "", nil
	}
	u, err := url.Parse(*next)
	if err != nil {
		return "", err
	}
	q := u.Query()
	return q.Get("next_url"), nil
}

func dataSourceIBMDb2InstanceRead(d *schema.ResourceData, meta interface{}) error {
	var instance rc.ResourceInstance
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()
	if _, ok := d.GetOk("name"); ok {
		name := d.Get("name").(string)
		resourceInstanceListOptions := rc.ListResourceInstancesOptions{
			Name: &name,
		}

		if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
			rg := rsGrpID.(string)
			resourceInstanceListOptions.ResourceGroupID = &rg
		}

		if service, ok := d.GetOk("service"); ok {

			serviceOff, err := rsCatRepo.FindByName(service.(string), true)
			if err != nil {
				return fmt.Errorf("[ERROR] Error retrieving service offering: %s", err)
			}
			resourceId := serviceOff[0].ID
			resourceInstanceListOptions.ResourceID = &resourceId
		}

		next_url := ""
		var instances []rc.ResourceInstance
		for {
			if next_url != "" {
				resourceInstanceListOptions.Start = &next_url
			}
			listInstanceResponse, resp, err := rsConClient.ListResourceInstances(&resourceInstanceListOptions)
			if err != nil {
				return fmt.Errorf("[ERROR] Error retrieving resource instance: %s with resp code: %s", err, resp)
			}
			next_url, err = getInstancesNext(listInstanceResponse.NextURL)
			if err != nil {
				return fmt.Errorf("[DEBUG] ListResourceInstances failed. Error occurred while parsing NextURL: %s", err)

			}
			instances = append(instances, listInstanceResponse.Resources...)
			if next_url == "" {
				break
			}
		}

		var filteredInstances []rc.ResourceInstance
		var location string

		if loc, ok := d.GetOk("location"); ok {
			location = loc.(string)
			for _, instance := range instances {
				if flex.GetLocationV2(instance) == location {
					filteredInstances = append(filteredInstances, instance)
				}
			}
		} else {
			filteredInstances = instances
		}

		if len(filteredInstances) == 0 {
			return fmt.Errorf("[ERROR] No resource instance found with name [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
		}
		if len(filteredInstances) > 1 {
			return fmt.Errorf("[ERROR] More than one resource instance found with name matching [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
		}
		instance = filteredInstances[0]
	} else if _, ok := d.GetOk("identifier"); ok {
		instanceGUID := d.Get("identifier").(string)
		getResourceInstanceOptions := &rc.GetResourceInstanceOptions{
			ID: &instanceGUID,
		}
		instances, res, err := rsConClient.GetResourceInstance(getResourceInstanceOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] No resource instance found with id [%s\n%v]", instanceGUID, res)
		}
		instance = *instances
		d.Set("name", instance.Name)
	}
	d.SetId(*instance.ID)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	serviceOff, err := rsCatRepo.GetServiceName(*instance.ResourceID)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving service offering: %s", err)
	}

	d.Set("service", serviceOff)

	d.Set(flex.ResourceName, instance.Name)
	d.Set(flex.ResourceCRN, instance.CRN)
	d.Set(flex.ResourceStatus, instance.State)
	// ### Modifiction : Setting the onetime credientials
	d.Set("onetime_credentials", instance.OnetimeCredentials)
	if instance.Parameters != nil {
		params, err := json.Marshal(instance.Parameters)
		if err != nil {
			return fmt.Errorf("[ERROR] Error marshalling instance parameters: %s", err)
		}
		if err = d.Set("parameters_json", string(params)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance parameters json: %s", err)
		}
	}
	rMgtClient, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return err
	}
	GetResourceGroup := rg.GetResourceGroupOptions{
		ID: instance.ResourceGroupID,
	}
	resourceGroup, resp, err := rMgtClient.GetResourceGroup(&GetResourceGroup)
	if err != nil || resourceGroup == nil {
		log.Printf("[ERROR] Error retrieving resource group: %s %s", err, resp)
	}
	if resourceGroup != nil && resourceGroup.Name != nil {
		d.Set(flex.ResourceGroupName, resourceGroup.Name)
	}
	d.Set("guid", instance.GUID)
	if len(instance.Extensions) == 0 {
		d.Set("extensions", instance.Extensions)
	} else {
		d.Set("extensions", flex.Flatten(instance.Extensions))
	}

	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/")

	servicePlan, err := rsCatRepo.GetServicePlanName(*instance.ResourcePlanID)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving plan: %s", err)
	}
	d.Set("plan", servicePlan)
	d.Set("crn", instance.CRN)
	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf(
			"Error on get of resource instance tags (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)

	d.Set("high_availability", instance.Parameters["high_availability"])
	d.Set("instance_type", instance.Parameters["instance_type"])
	d.Set("backup_location", instance.Parameters["backup_location"])
	d.Set("disk_encryption_instance_crn", instance.Parameters["disk_encryption_instance_crn"])
	d.Set("disk_encryption_crn", instance.Parameters["disk_encryption_crn"])
	d.Set("oracle_compatibility", instance.Parameters["oracle_compatibility"])
	d.Set("subscription_id", instance.Parameters["subscription_id"])

	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		log.Printf("[ERROR] Error retrieving db2saas client: %s", err)
	}

	//appending whitelist ips config
	if _, ok := d.GetOk("whitelist_config"); ok {
		getDb2SaasWhitelistOptions := &db2saasv1.GetDb2SaasWhitelistOptions{}

		getDb2SaasWhitelistOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

		successGetWhitelistIPs, _, err := db2saasClient.GetDb2SaasWhitelistWithContext(context.Background(), getDb2SaasWhitelistOptions)
		if err != nil {
			log.Printf("[ERROR] Error retrieving db2saas whitelist: %s", err)
		}

		d.SetId(dataSourceIbmDb2SaasWhitelistID(d))

		ipAddresses := []map[string]interface{}{}
		for _, ipAddressesItem := range successGetWhitelistIPs.IpAddresses {
			ipAddressesItemMap, err := DataSourceIbmDb2SaasWhitelistIpAddressToMap(&ipAddressesItem) // #nosec G601
			if err != nil {
				log.Printf("[ERROR] Error converting ip addresses to map: %s", err)
			}
			ipAddresses = append(ipAddresses, ipAddressesItemMap)
		}
		if err = d.Set("ip_addresses", ipAddresses); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance ip addresses: %s", err)
		}
	}

	//append autoscale config
	if _, ok := d.GetOk("autoscaling_config"); ok {
		getDb2SaasAutoscaleOptions := &db2saasv1.GetDb2SaasAutoscaleOptions{}

		getDb2SaasAutoscaleOptions.SetXDbProfile(d.Get("x_db_profile").(string))

		successAutoScaling, _, err := db2saasClient.GetDb2SaasAutoscaleWithContext(context.Background(), getDb2SaasAutoscaleOptions)
		if err != nil {
			log.Printf("[ERROR] Error retrieving db2saas autoscale: %s", err)
		}

		d.SetId(dataSourceIbmDb2SaasAutoscaleID(d))

		if err = d.Set("auto_scaling_allow_plan_limit", successAutoScaling.AutoScalingAllowPlanLimit); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_allow_plan_limit: %s", err)
		}

		if err = d.Set("auto_scaling_enabled", successAutoScaling.AutoScalingEnabled); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_enabled: %s", err)
		}

		if err = d.Set("auto_scaling_max_storage", flex.IntValue(successAutoScaling.AutoScalingMaxStorage)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_max_storage: %s", err)
		}

		if err = d.Set("auto_scaling_over_time_period", flex.IntValue(successAutoScaling.AutoScalingOverTimePeriod)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_over_time_period: %s", err)
		}

		if err = d.Set("auto_scaling_pause_limit", flex.IntValue(successAutoScaling.AutoScalingPauseLimit)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_pause_limit: %s", err)
		}

		if err = d.Set("auto_scaling_threshold", flex.IntValue(successAutoScaling.AutoScalingThreshold)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance auto_scaling_threshold: %s", err)
		}

		if err = d.Set("storage_unit", successAutoScaling.StorageUnit); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance storage_unit: %s", err)
		}

		if err = d.Set("storage_utilization_percentage", flex.IntValue(successAutoScaling.StorageUtilizationPercentage)); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance storage_utilization_percentage: %s", err)
		}

		if err = d.Set("support_auto_scaling", successAutoScaling.SupportAutoScaling); err != nil {
			return fmt.Errorf("[ERROR] Error setting instance support_auto_scaling: %s", err)
		}
	}

	//appending connectioninfo config
	if _, ok := d.GetOk("connectioninfo_config"); ok {
		getDb2SaasConnectionInfoOptions := &db2saasv1.GetDb2SaasConnectionInfoOptions{}

		getDb2SaasConnectionInfoOptions.SetDeploymentID(d.Get("deployment_id").(string))
		getDb2SaasConnectionInfoOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

		successConnectionInfo, _, err := db2saasClient.GetDb2SaasConnectionInfoWithContext(context.Background(), getDb2SaasConnectionInfoOptions)
		if err != nil {
			log.Printf("[ERROR] Error retrieving connection info: %s", err)
		}

		d.SetId(dataSourceIbmDb2SaasConnectionInfoID(d))

		if !core.IsNil(successConnectionInfo.Public) {
			public := []map[string]interface{}{}
			publicMap, err := DataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPublicToMap(successConnectionInfo.Public)
			if err != nil {
				log.Printf("[ERROR] Error converting public connection info to map: %s", err)
			}
			public = append(public, publicMap)
			if err = d.Set("public", public); err != nil {
				return fmt.Errorf("[ERROR] Error setting instance public: %s", err)
			}
		}

		if !core.IsNil(successConnectionInfo.Private) {
			private := []map[string]interface{}{}
			privateMap, err := DataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPrivateToMap(successConnectionInfo.Private)
			if err != nil {
				log.Printf("[ERROR] Error converting private connection info to map: %s", err)
			}
			private = append(private, privateMap)
			if err = d.Set("private", private); err != nil {
				return fmt.Errorf("[ERROR] Error setting instance private: %s", err)
			}
		}
	}

	return nil
}
