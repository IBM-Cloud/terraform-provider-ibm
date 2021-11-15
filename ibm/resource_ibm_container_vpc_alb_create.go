// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMContainerVpcAlbCreateNew() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVpcAlbCreate,
		Read:     resourceIBMContainerVpcALBRead,
		Update:   resourceIBMContainerVpcALBUpdate,
		Delete:   resourceIBMContainerVpcALBDelete,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{

			//post req
			// "enable_by_default": {
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	Default:     true,
			// 	Description: "If set to true, the ALB is enabled by default.",
			// },
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of ALB that you want to create.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone where you want to deploy the ALB.",
			},
			"ingress_image": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The type of Ingress image that you want to use for your ALB deployment.",
			},
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cluster that the ALB belongs to.",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group.",
			},

			//response
			"alb_id": {
				Type: schema.TypeString,
				//Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the application load balancer (ALB).",
			},

			//get
			"alb_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the ALB",
			},
			"enable": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"disable_deployment"},
				Description:   "Enable the ALB instance in the cluster",
			},
			"disable_deployment": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"enable"},
				Description:   "Disable the ALB instance in the cluster",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ALB name",
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
		},
	}
}

func resourceIBMContainerVpcAlbCreate(d *schema.ResourceData, meta interface{}) error {

	albClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	albAPI := albClient.Albs()

	var cluster string
	if v, ok := d.GetOkExists("cluster"); ok {
		cluster = v.(string)
	} else {
		return fmt.Errorf("Provide `clusterIDorName`")
	}

	var albType string
	if v, ok := d.GetOkExists("type"); ok {
		albType = v.(string)
	} else {
		return fmt.Errorf("Provide `type`")
	}

	var zone string
	if v, ok := d.GetOkExists("zone"); ok {
		zone = v.(string)
	} else {
		return fmt.Errorf("Provide `zone`")
	}

	enableByDefault := d.Get("enable").(bool)

	params := v2.AlbCreateReq{
		ZoneAlb:         zone,
		Type:            albType,
		EnableByDefault: enableByDefault,
		Cluster:         cluster,
	}

	targetEnv, _ := getVpcClusterTargetHeader(d, meta)

	//v2.AlbCreateResp
	albResp, err := albAPI.CreateAlb(params, targetEnv)
	if err != nil {
		return err
	}

	d.SetId(albResp.Alb)
	return nil
}
