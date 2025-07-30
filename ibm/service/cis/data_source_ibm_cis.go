// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"net/url"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMCISInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISInstanceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Resource instance name for example, my cis instance",
				Type:        schema.TypeString,
				Required:    true,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the resource group in which the cis instance is present",
			},

			"guid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of resource instance",
			},

			"location": {
				Description: "The location or the environment in which cis instance exists",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"service": {
				Description: "The name of the Cloud Internet Services offering, 'internet-svcs'",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"plan": {
				Description: "The plan type of the cis instance",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"status": {
				Description: "The resource instance status",
				Type:        schema.TypeString,
				Computed:    true,
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
			},
		},
	}
}

func dataSourceIBMCISInstanceRead(d *schema.ResourceData, meta interface{}) error {
	var instance rc.ResourceInstance
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)

	resourceInstanceListOptions := rc.ListResourceInstancesOptions{
		Name: &name,
	}

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rg := rsGrpID.(string)
		resourceInstanceListOptions.ResourceGroupID = &rg
	} else {
		defaultRg, err := flex.DefaultResourceGroup(meta)
		if err != nil {
			return err
		}
		resourceInstanceListOptions.ResourceGroupID = &defaultRg
	}

	if service, ok := d.GetOk("service"); ok {
		name := service.(string)
		resourceInstanceListOptions.ResourceID = &name
	}
	next_url := ""
	var instances []rc.ResourceInstance
	for {
		if next_url != "" {
			resourceInstanceListOptions.Start = &next_url
		}
		listInstanceResponse, resp, err := rsConClient.ListResourceInstances(&resourceInstanceListOptions)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error retrieving resource instance: %s with resp code: %s", err, resp)
		}
		next_url, err = getInstancesNext(listInstanceResponse.NextURL)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error retrieving service offering: %s", err)
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
		return flex.FmtErrorf("[ERROR] No resource instance found with name [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
	}

	if len(filteredInstances) > 1 {
		return flex.FmtErrorf("[ERROR] More than one resource instance found with name matching [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
	}
	instance = filteredInstances[0]

	d.SetId(*instance.ID)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	d.Set("guid", instance.GUID)
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return err
	}
	options := globalcatalogv1.GetCatalogEntryOptions{

		ID: instance.ResourceID,
	}
	service, _, err := globalClient.GetCatalogEntry(&options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving service offering: %s", err)
	}
	d.Set("service", service.Name)
	planOptions := globalcatalogv1.GetCatalogEntryOptions{

		ID: instance.ResourcePlanID,
	}
	plan, _, err := globalClient.GetCatalogEntry(&planOptions)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving plan: %s", err)
	}
	d.Set("plan", plan.Name)

	d.Set(flex.ResourceName, instance.Name)
	d.Set(flex.ResourceCRN, instance.CRN)
	d.Set(flex.ResourceStatus, instance.State)

	rMgtClient, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return err
	}
	GetResourceGroup := rg.GetResourceGroupOptions{
		ID: instance.ResourceGroupID,
	}
	resourceGroup, resp, err := rMgtClient.GetResourceGroup(&GetResourceGroup)
	if err != nil || resourceGroup == nil {
		flex.FmtErrorf("[ERROR] Error retrieving resource group: %s %s", err, resp)
	}
	if resourceGroup != nil && resourceGroup.Name != nil {
		d.Set(flex.ResourceGroupName, resourceGroup.Name)
	}

	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/internet-svcs/"+url.QueryEscape(*instance.CRN))

	return nil
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
