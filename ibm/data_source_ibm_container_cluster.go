package ibm

import (
	"fmt"
	"strings"

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
			"worker_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_per_zone": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hardware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						"zones": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_vlan": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_vlan": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"worker_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
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
			"alb_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "all",
				ValidateFunc: validateAllowedStringValue([]string{"private", "public", "all"}),
			},
			"albs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alb_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_of_instances": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alb_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resize": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disable_deployment": {
							Type:     schema.TypeBool,
							Computed: true,
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
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The cluster region",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group.",
				Computed:    true,
			},
			"public_service_endpoint": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"private_service_endpoint": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"public_service_endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"private_service_endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
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
	workerPoolsAPI := csClient.WorkerPools()
	albsAPI := csClient.Albs()

	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}
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
	workerPools, err := workerPoolsAPI.ListWorkerPools(name, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving worker pools of the cluster %s: %s", name, err)
	}

	albs, err := albsAPI.ListClusterALBs(name, targetEnv)
	if err != nil && !strings.Contains(err.Error(), "The specified cluster is a lite cluster.") {
		return fmt.Errorf("Error retrieving alb's of the cluster %s: %s", name, err)
	}

	filterType := d.Get("alb_type").(string)
	filteredAlbs := flattenAlbs(albs, filterType)

	d.SetId(clusterFields.ID)
	d.Set("worker_count", clusterFields.WorkerCount)
	d.Set("workers", workers)
	d.Set("region", clusterFields.Region)
	d.Set("bounded_services", boundedServices)
	d.Set("vlans", flattenVlans(clusterFields.Vlans))
	d.Set("is_trusted", clusterFields.IsTrusted)
	d.Set("worker_pools", flattenWorkerPools(workerPools))
	d.Set("albs", filteredAlbs)
	d.Set("resource_group_id", clusterFields.ResourceGroupID)
	d.Set("public_service_endpoint", clusterFields.PublicServiceEndpointEnabled)
	d.Set("private_service_endpoint", clusterFields.PrivateServiceEndpointEnabled)
	d.Set("public_service_endpoint_url", clusterFields.PublicServiceEndpointURL)
	d.Set("private_service_endpoint_url", clusterFields.PrivateServiceEndpointURL)

	return nil
}
