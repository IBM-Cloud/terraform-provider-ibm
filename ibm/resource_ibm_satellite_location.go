// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.ibm.com/ibmcloud/kubernetesservice-go-sdk/kubernetesserviceapiv1"
)

const (
	satLocation = "location"
	sateLocZone = "managed_from"

	isLocationDeleting     = "deleting"
	isLocationDeleteDone   = "done"
	isLocationDeploying    = "deploying"
	isLocationReady        = "action required"
	isLocationDeployFailed = "deploy_failed"
)

func resourceIBMSatelliteLocation() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMSatelliteLocationCreate,
		Read:     resourceIBMSatelliteLocationRead,
		Update:   resourceIBMSatelliteLocationUpdate,
		Delete:   resourceIBMSatelliteLocationDelete,
		Importer: &schema.ResourceImporter{},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			satLocation: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique name for the new Satellite location",
			},
			sateLocZone: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IBM Cloud metro from which the Satellite location is managed",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the new Satellite location",
			},
			"logging_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account ID for IBM Log Analysis with LogDNA log forwarding",
			},
			"cos_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "COSBucket - IBM Cloud Object Storage bucket configuration details",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"cos_credentials": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "COSAuthorization - IBM Cloud Object Storage authorization keys",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The HMAC secret access key ID",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The HMAC secret access key",
						},
					},
				},
			},
			"zones": {
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The names of at least three high availability zones to use for the location",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ID of the resource group.",
			},
		},
	}
}

func resourceIBMSatelliteLocationCreate(d *schema.ResourceData, meta interface{}) error {
	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	createSatLocOptions := &kubernetesserviceapiv1.CreateSatelliteLocationOptions{}
	satLocation := d.Get(satLocation).(string)
	createSatLocOptions.Name = &satLocation
	sateLocZone := d.Get(sateLocZone).(string)
	createSatLocOptions.Location = &sateLocZone

	if v, ok := d.GetOk("cos_config"); ok {
		createSatLocOptions.CosConfig = expandCosConfig(v.([]interface{}))
	}

	if v, ok := d.GetOk("cos_credentials"); ok {
		createSatLocOptions.CosCredentials = expandCosCredentials(v.([]interface{}))
	}

	if v, ok := d.GetOk("logging_account_id"); ok {
		logAccID := v.(string)
		createSatLocOptions.LoggingAccountID = &logAccID
	}

	if v, ok := d.GetOk("description"); ok {
		desc := v.(string)
		createSatLocOptions.Description = &desc
	}

	if v, ok := d.GetOk("zones"); ok {
		z := v.(*schema.Set)
		createSatLocOptions.Zones = flatterSatelliteZones(z)
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		pathParamsMap := map[string]string{
			"X-Auth-Resource-Group": v.(string),
		}
		createSatLocOptions.Headers = pathParamsMap
	}

	instance, response, err := satClient.CreateSatelliteLocation(createSatLocOptions)
	if err != nil {
		return fmt.Errorf("Error Creating Satellite Location: %s\n%s", err, response)
	}

	d.SetId(satLocation)
	log.Printf("[INFO] Created satellite location : %s", satLocation)

	//Wait for location to be in ready state
	_, err = waitForLocationToReady(*instance.ID, d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for location (%s) to reach ready state: %s", *instance.ID, err)
	}

	return resourceIBMSatelliteLocationRead(d, meta)
}

func resourceIBMSatelliteLocationRead(d *schema.ResourceData, meta interface{}) error {
	ID := d.Id()
	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationOptions{
		Controller: &ID,
	}

	instance, response, err := satClient.GetSatelliteLocation(getSatLocOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set(satLocation, ID)
	if instance.Description != nil {
		d.Set("description", *instance.Description)
	}

	if instance.Datacenter != nil {
		d.Set(sateLocZone, *instance.Datacenter)
	}

	if instance.WorkerZones != nil {
		d.Set("zones", instance.WorkerZones)
	}

	if instance.ResourceGroup != nil {
		d.Set("resource_group_id", instance.ResourceGroup)
	}

	return nil
}

func resourceIBMSatelliteLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMSatelliteLocationDelete(d *schema.ResourceData, meta interface{}) error {
	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	removeSatLocOptions := &kubernetesserviceapiv1.RemoveSatelliteLocationOptions{}
	name := d.Get(satLocation).(string)
	removeSatLocOptions.Controller = &name

	response, err := satClient.RemoveSatelliteLocation(removeSatLocOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Deleting Satellite Location: %s\n%s", err, response)
	}

	//Wait for location to delete
	_, err = waitForLocationDelete(name, d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for deleting location instance: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForLocationDelete(location string, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return false, err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{isLocationDeleting, ""},
		Target:  []string{isLocationDeleteDone},
		Refresh: func() (interface{}, string, error) {
			getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationsOptions{}
			locations, response, err := satClient.GetSatelliteLocations(getSatLocOptions)
			if err != nil {
				return nil, "", fmt.Errorf("Error Getting locations list to delete : %s\n%s", err, response)
			}

			isExist := false
			if locations != nil {
				for _, loc := range locations {
					if *loc.ID == location || *loc.Name == location {
						isExist = true
						return "", isLocationDeleting, nil
					}
				}
				if isExist == false {
					return location, isLocationDeleteDone, nil
				}
			}
			return nil, "", fmt.Errorf("Failed to delete location : %s\n%s", err, response)
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForState()
}

func waitForLocationToReady(loc string, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return false, err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{isLocationDeploying},
		Target:  []string{isLocationReady, isLocationDeployFailed},
		Refresh: func() (interface{}, string, error) {
			getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationOptions{
				Controller: ptrToString(loc),
			}
			location, response, err := satClient.GetSatelliteLocation(getSatLocOptions)
			if err != nil {
				return nil, "", fmt.Errorf("Error Getting location : %s\n%s", err, response)
			}

			if location != nil && *location.State == isLocationDeployFailed {
				return location, isLocationDeployFailed, fmt.Errorf("The location is in failed state: %s", d.Id())
			}

			if location != nil && *location.State == isLocationReady {
				return location, isLocationReady, nil
			}
			return location, isLocationDeploying, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForState()
}
