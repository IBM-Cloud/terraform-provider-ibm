// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
)

func resourceIBMContainerAlbCreate() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerClassicAlbCreate,
		Read:     resourceIBMContainerClassicAlbRead,
		Update:   resourceIBMContainerClassicAlbUpdate,
		Delete:   resourceIBMContainerClassicAlbDelete,
		Exists:   resourceIBMContainerClassicAlbExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			//post req
			"enable_by_default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If set to true, the ALB is enabled by default.",
			},
			"ingress_image": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The type of Ingress image that you want to use for your ALB deployment.",
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				//ForceNew:    true,
				Description: "The IP address that you want to assign to the ALB.",
			},
			"nlb_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The version of the network load balancer that you want to use for the ALB.",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew:    true,
				Description: "The type of ALB that you want to create.",
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew:    true,
				Description: "The VLAN ID that you want to use for your ALBs.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone where you want to deploy the ALB.",
			},
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cluster that the ALB belongs to.",
			},

			//response
			"alb_id": {
				Type: schema.TypeString,
				//Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the application load balancer (ALB).",
			},
		},
	}
}

func resourceIBMContainerClassicAlbCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("CREATE")
	fmt.Println("CREATE")
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	albAPI := albClient.Albs()

	// "cluster":"string" //mandatory
	// "enableByDefault": true,
	// "ingressImage": "string",
	// "ip": "string",
	// "nlbVersion": "string",
	// "type": "string", //mandatory
	// "vlanID": "string", //mandatory
	// "zone": "string" //mandatory

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

	var vlanID string
	if v, ok := d.GetOkExists("vlan_id"); ok {
		vlanID = v.(string)
	} else {
		return fmt.Errorf("Provide `vlanID`")
	}

	var zone string
	if v, ok := d.GetOkExists("zone"); ok {
		zone = v.(string)
	} else {
		return fmt.Errorf("Provide `zone`")
	}

	enableByDefault := d.Get("enable_by_default").(bool)
	ingressImage := d.Get("ingress_image").(string)
	ip := d.Get("ip").(string)
	nlbVersion := d.Get("nlb_version").(string)

	params := v1.CreateALB{
		Zone:            zone,
		VlanID:          vlanID,
		Type:            albType,
		EnableByDefault: enableByDefault,
		IP:              ip,
		NLBVersion:      nlbVersion,
		IngressImage:    ingressImage,
	}

	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return err
	}

	//v1.AlbCreateResp
	albResp, err := albAPI.CreateALB(params, cluster, targetEnv)
	if err != nil {
		return fmt.Errorf("Error creating ALb to the cluster %s", err)
	}

	d.SetId(albResp.Alb)
	return nil
}

func resourceIBMContainerClassicAlbRead(d *schema.ResourceData, meta interface{}) error {
	// log.Println("READ")
	// fmt.Println("READ")

	// albClient, err := meta.(ClientSession).ContainerAPI()
	// if err != nil {
	// 	return err
	// }

	// albID := d.Id()

	// albAPI := albClient.Albs()
	// targetEnv, err := getAlbTargetHeader(d, meta)
	// if err != nil {
	// 	return err
	// }
	// albConfig, err := albAPI.GetALB(albID, targetEnv)
	// if err != nil {
	// 	return err
	// }

	// // "cluster":"string" //mandatory
	// // "enableByDefault": true,
	// // "ingressImage": "string",
	// // "ip": "string",
	// // "nlbVersion": "string",
	// // "type": "string", //mandatory
	// // "vlanID": "string", //mandatory
	// // "zone": "string" //mandatory

	// d.Set("type", albConfig.ALBType)
	// d.Set("cluster", albConfig.ClusterID)
	// d.Set("ip", albConfig.ALBIP)
	// d.Set("zone", albConfig.Zone)

	return nil
}

func resourceIBMContainerClassicAlbUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerClassicAlbDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerClassicAlbExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}
