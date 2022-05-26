// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func ResourceIBMContainerDedicatedHostPool() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerDedicatedHostPoolCreate,
		Read:     resourceIBMContainerDedicatedHostPoolRead,
		Delete:   resourceIBMContainerDedicatedHostPoolDelete,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the dedicated host pool",
			},
			"metro": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The metro to create the dedicated host pool in",
			},
			"flavor_class": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The flavor class of the dedicated host pool",
			},
			"host_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The count of the hosts under the dedicated host pool",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the dedicated host pool",
			},
			"zones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The zones of the dedicated host pool",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capacity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"memory_bytes": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"vcpu": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"host_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"zone": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"worker_pools": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The worker pools of the dedicated host pool",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"worker_pool_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMContainerDedicatedHostPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	dedicatedHostPoolAPI := client.DedicatedHostPool()
	targetEnv := v2.ClusterTargetHeader{}

	params := v2.CreateDedicatedHostPoolRequest{
		FlavorClass: d.Get("flavor_class").(string),
		Metro:       d.Get("metro").(string),
		Name:        d.Get("name").(string),
	}

	res, err := dedicatedHostPoolAPI.CreateDedicatedHostPool(params, targetEnv)
	if err != nil {
		return fmt.Errorf("[ERROR] Error creating host pool %v", err)
	}

	d.SetId(res.ID)

	return resourceIBMContainerDedicatedHostPoolRead(d, meta)
}

func resourceIBMContainerDedicatedHostPoolRead(d *schema.ResourceData, meta interface{}) error {
	err := getIBMContainerDedicatedHostPool(d.Id(), d, meta)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				d.SetId("")
				return nil
			}
		}
		return fmt.Errorf("[ERROR] Error retrieving host pool details %v", err)
	}
	return nil
}

func getIBMContainerDedicatedHostPool(hostPoolID string, d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	dedicatedHostPoolAPI := client.DedicatedHostPool()
	targetEnv := v2.ClusterTargetHeader{}

	dedicatedHostPool, err := dedicatedHostPoolAPI.GetDedicatedHostPool(hostPoolID, targetEnv)
	if err != nil {
		return err
	}

	d.Set("name", dedicatedHostPool.Name)
	d.Set("metro", dedicatedHostPool.Metro)
	d.Set("flavor_class", dedicatedHostPool.FlavorClass)
	d.Set("host_count", dedicatedHostPool.HostCount)
	d.Set("state", dedicatedHostPool.State)

	zones := make([]map[string]interface{}, len(dedicatedHostPool.Zones))
	for i, zone := range dedicatedHostPool.Zones {
		zones[i] = map[string]interface{}{
			"capacity": []interface{}{map[string]interface{}{
				"memory_bytes": zone.Capacity.MemoryBytes,
				"vcpu":         zone.Capacity.VCPU,
			}},
			"host_count": zone.HostCount,
			"zone":       zone.Zone,
		}
	}
	d.Set("zones", zones)

	workerpools := make([]map[string]interface{}, len(dedicatedHostPool.WorkerPools))
	for i, wpool := range dedicatedHostPool.WorkerPools {
		workerpools[i] = map[string]interface{}{
			"cluster_id":     wpool.ClusterID,
			"worker_pool_id": wpool.WorkerPoolID,
		}
	}
	d.Set("worker_pools", workerpools)

	return nil
}

func resourceIBMContainerDedicatedHostPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	dedicatedHostPoolAPI := client.DedicatedHostPool()
	targetEnv := v2.ClusterTargetHeader{}

	params := v2.RemoveDedicatedHostPoolRequest{
		HostPoolID: d.Id(),
	}

	if err := dedicatedHostPoolAPI.RemoveDedicatedHostPool(params, targetEnv); err != nil {
		return fmt.Errorf("[ERROR] Error removing host pool %v", err)
	}

	return nil
}
