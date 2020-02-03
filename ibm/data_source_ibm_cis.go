package ibm

import (
	"fmt"
	"net/url"

	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMCISInstance() *schema.Resource {
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
				Type:        schema.TypeMap,
				Computed:    true,
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
			},
		},
	}
}

func dataSourceIBMCISInstanceRead(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	rsAPI := rsConClient.ResourceServiceInstance()
	name := d.Get("name").(string)

	rsInstQuery := controller.ServiceInstanceQuery{
		Name: name,
	}

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rsInstQuery.ResourceGroupID = rsGrpID.(string)
	} else {
		rsMangClient, err := meta.(ClientSession).ResourceManagementAPIv2()
		if err != nil {
			return err
		}
		resourceGroupQuery := managementv2.ResourceGroupQuery{
			Default: true,
		}
		grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
		if err != nil {
			return err
		}
		rsInstQuery.ResourceGroupID = grpList[0].ID
	}

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	if service, ok := d.GetOk("service"); ok {

		serviceOff, err := rsCatRepo.FindByName(service.(string), true)
		if err != nil {
			return fmt.Errorf("Error retrieving service offering: %s", err)
		}

		rsInstQuery.ServiceID = serviceOff[0].ID
	}

	var instances []models.ServiceInstance

	instances, err = rsAPI.ListInstances(rsInstQuery)
	if err != nil {
		return err
	}
	var filteredInstances []models.ServiceInstance
	var location string

	if loc, ok := d.GetOk("location"); ok {
		location = loc.(string)
		for _, instance := range instances {
			if getLocation(instance) == location {
				filteredInstances = append(filteredInstances, instance)
			}
		}
	} else {
		filteredInstances = instances
	}

	if len(filteredInstances) == 0 {
		return fmt.Errorf("No resource instance found with name [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
	}

	var instance models.ServiceInstance

	if len(filteredInstances) > 1 {
		return fmt.Errorf(
			"More than one resource instance found with name matching [%s]\nIf not specified please specify more filters like resource_group_id if instance doesn't exists in default group, location or service", name)
	}
	instance = filteredInstances[0]

	d.SetId(instance.ID)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	serviceOff, err := rsCatRepo.GetServiceName(instance.ServiceID)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}

	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(instance.ServicePlanID)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	d.Set("plan", servicePlan)

	d.Set(ResourceName, instance.Name)
	d.Set(ResourceCRN, instance.Crn.String())
	d.Set(ResourceStatus, instance.State)
	d.Set(ResourceGroupName, instance.ResourceGroupName)

	rcontroller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, rcontroller+"/internet-svcs/"+url.QueryEscape(instance.Crn.String()))

	return nil
}
