package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMContainerCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerClusterRead,

		Schema: map[string]*schema.Schema{
			"cluster_name_id": {
				Description: "Name or id of the cluster",
				Type:        schema.TypeString,
				Required:    true,
			},
			"worker_count": {
				Description: "Number of workers",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"workers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"is_trusted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"bounded_services": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_key_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vlans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ips": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"is_public": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_byoip": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"cidr": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	csAPI := csClient.Clusters()
	wrkAPI := csClient.Workers()

	targetEnv := getClusterTargetHeader(d)
	name := d.Get("cluster_name_id").(string)

	clusterFields, err := csAPI.Find(name, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving cluster: %s", err)
	}
	workerFields, err := wrkAPI.List(name, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving workers for cluster: %s", err)
	}
	workers := make([]string, len(workerFields))
	for i, worker := range workerFields {
		workers[i] = worker.ID
	}
	servicesBoundToCluster, err := csAPI.ListServicesBoundToCluster(name, "", targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving services bound to cluster: %s", err)
	}
	boundedServices := make([]map[string]interface{}, 0)
	for _, service := range servicesBoundToCluster {
		boundedService := make(map[string]interface{})
		boundedService["service_name"] = service.ServiceName
		boundedService["service_id"] = service.ServiceID
		boundedService["service_key_name"] = service.ServiceKeyName
		boundedService["namespace"] = service.Namespace
		boundedServices = append(boundedServices, boundedService)
	}

	d.SetId(clusterFields.ID)
	d.Set("worker_count", clusterFields.WorkerCount)
	d.Set("workers", workers)
	d.Set("bounded_services", boundedServices)
	d.Set("vlans", flattenVlans(clusterFields.Vlans))
	d.Set("is_trusted", clusterFields.IsTrusted)

	return nil
}
