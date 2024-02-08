// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"log"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMContainerVpcALB() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVpcALBEnable,
		Read:     resourceIBMContainerVpcALBRead,
		Update:   resourceIBMContainerVpcALBUpdate,
		Delete:   resourceIBMContainerVpcALBDelete,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"alb_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ALB ID",
			},
			"alb_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the ALB",
			},
			"cluster": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cluster id",
			},
			"enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Enable the ALB instance in the cluster",
			},
			"disable_deployment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable the ALB instance in the cluster",
				Deprecated:  "Remove this attribute's configuration as it no longer is used, use enable instead",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ALB name",
				Deprecated:  "Remove this attribute's configuration as it no longer is used",
			},
			"load_balancer_hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer host name",
			},
			"resize": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "boolean value to resize the albs",
				Deprecated:  "Remove this attribute's configuration as it no longer is used",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ALB state",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the ALB",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone info.",
			},
			"ingress_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of Ingress image that you want to use for your ALB deployment.",
			},
			"resource_group_id": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "ID of the resource group.",
			},
		},
	}
}

func resourceIBMContainerVpcALBEnable(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	var enable bool
	albID := d.Get("alb_id").(string)
	if v, ok := d.GetOkExists("enable"); ok {
		enable = v.(bool)
	} else {
		return fmt.Errorf("[ERROR] Missing `enable` argument")
	}

	_, err = waitForVpcClusterAvailable(d, meta, albID, schema.TimeoutCreate)
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for cluster resource availabilty (%s) : %s", d.Id(), err)
	}

	params := v2.AlbConfig{
		AlbID:  albID,
		Enable: enable,
	}

	albAPI := albClient.Albs()
	targetEnv, _ := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	if enable {
		err = albAPI.EnableAlb(params, targetEnv)
		if err != nil {
			return err
		}
	} else {
		err = albAPI.DisableAlb(params, targetEnv)
		if err != nil {
			return err
		}
	}

	d.SetId(albID)
	_, err = waitForVpcContainerALB(d, meta, albID, schema.TimeoutCreate, enable)
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for create resource alb (%s) : %s", d.Id(), err)
	}

	return resourceIBMContainerVpcALBRead(d, meta)
}

func resourceIBMContainerVpcALBRead(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	albID := d.Id()

	albAPI := albClient.Albs()
	targetEnv, _ := getVpcClusterTargetHeader(d, meta)

	albConfig, err := albAPI.GetAlb(albID, targetEnv)
	if err != nil {
		return err
	}

	d.Set("alb_type", albConfig.AlbType)
	d.Set("cluster", albConfig.Cluster)
	d.Set("name", albConfig.Name)
	d.Set("enable", albConfig.Enable)
	d.Set("disable_deployment", albConfig.DisableDeployment)
	d.Set("alb_id", albID)
	d.Set("resize", albConfig.Resize)
	d.Set("zone", albConfig.ZoneAlb)
	d.Set("status", albConfig.Status)
	d.Set("state", albConfig.State)
	d.Set("load_balancer_hostname", albConfig.LoadBalancerHostname)
	d.Set("ingress_image", albConfig.AlbBuild)

	return nil
}

func resourceIBMContainerVpcALBUpdate(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	albAPI := albClient.Albs()

	if d.HasChange("enable") {
		enable := d.Get("enable").(bool)
		albID := d.Id()

		_, err = waitForVpcClusterAvailable(d, meta, albID, schema.TimeoutCreate)
		if err != nil {
			return fmt.Errorf("[ERROR] Error waiting for cluster resource availabilty (%s) : %s", d.Id(), err)
		}

		params := v2.AlbConfig{
			AlbID:  albID,
			Enable: enable,
		}

		targetEnv, _ := getVpcClusterTargetHeader(d, meta)

		if enable {
			err = albAPI.EnableAlb(params, targetEnv)
			if err != nil {
				return err
			}
		} else {
			err = albAPI.DisableAlb(params, targetEnv)
			if err != nil {
				return err
			}
		}

		_, err = waitForVpcContainerALB(d, meta, albID, schema.TimeoutUpdate, enable)
		if err != nil {
			return fmt.Errorf("[ERROR] Error waiting for updating resource alb (%s) : %s", d.Id(), err)
		}

	}
	return resourceIBMContainerVpcALBRead(d, meta)
}

func waitForVpcContainerALB(d *schema.ResourceData, meta interface{}, albID, timeout string, enable bool) (interface{}, error) {
	albClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}
	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"active"},
		Refresh: func() (interface{}, string, error) {
			targetEnv, _ := getVpcClusterTargetHeader(d, meta)
			alb, err := albClient.Albs().GetAlb(albID, targetEnv)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("[ERROR] The resource alb %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if enable {
				if !alb.Enable {
					return alb, "pending", nil
				}
			} else {
				if alb.Enable {
					return alb, "pending", nil
				}
			}
			return alb, "active", nil
		},
		Timeout:    d.Timeout(timeout),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMContainerVpcALBDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return nil
}

func waitForVpcClusterAvailable(d *schema.ResourceData, meta interface{}, albID, timeout string) (interface{}, error) {
	albClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}
	createStateConf := &resource.StateChangeConf{
		Pending: []string{deployRequested, deployInProgress},
		Target:  []string{ready},
		Refresh: func() (interface{}, string, error) {
			targetEnv, _ := getVpcClusterTargetHeader(d, meta)
			albInfo, err := albClient.Albs().GetAlb(albID, targetEnv)
			if err == nil {
				cluster := albInfo.Cluster
				workerPools, err := albClient.WorkerPools().ListWorkerPools(cluster, targetEnv)
				if err != nil {
					return workerPools, deployInProgress, err
				}
				for _, wpool := range workerPools {
					workers, err := albClient.Workers().ListByWorkerPool(cluster, wpool.ID, false, targetEnv)
					if err != nil {
						return wpool, deployInProgress, err
					}
					healthCounter := 0

					for _, worker := range workers {
						log.Println("worker: ", worker.ID)
						log.Println("worker health state:  ", worker.Health.State)

						if worker.Health.State == normal {
							healthCounter++
						}
					}
					if healthCounter != len(workers) {
						log.Println("all the worker nodes are not in normal state")
						return wpool, deployInProgress, nil
					}
				}
			} else {
				log.Println("ALB info not available")
				return albInfo, deployInProgress, err
			}
			return albInfo, ready, nil
		},
		Timeout:    d.Timeout(timeout),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}
	return createStateConf.WaitForState()
}
