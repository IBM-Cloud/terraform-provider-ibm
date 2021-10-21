// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
)

// resource "ibm_container_classic_alb_create" "alb"{
// 	enableByDefault	=boolean,
// 	ingressImage=	string,
// 	ip=	string,
// 	nlbVersion=	string,
// 	type=	string,
// 	vlanID=	string,
// 	zone=	string,
// }

// POST

// "enableByDefault": true,
// "ingressImage": "string",
// "ip": "string",
// "nlbVersion": "string",
// "type": "string",
// "vlanID": "string",
// "zone": "string"

// GET

// ALBConfig represents the ALB configuration.{
// 	albBuild	string
// 	The build number of the ALB.
// 	albID	string
// 	The ID of the application load balancer (ALB).

// 	albType	string
// 	The type of ALB.
// 	albip	string
// 	The public IP address that exposes the ALB.
// 	authBuild	string
// 	The auth build of the ALB.
// 	clusterID	string
// 	The ID of the cluster that the ALB belongs to.
// 	createdDate	string
// 	The date the ALB was created.
// 	disableDeployment	boolean
// 	If set to true, the deployment of the ALB is disabled.

// 	enable	boolean
// 	Set to true to enable the ALB, or false to disable the ALB for the cluster.

// 	name	string
// 	The name of the cluster that the ALB belongs to.
// 	nlbVersion	string
// 	The version of network load balancer that the ALB uses.
// 	numOfInstances	string
// 	Desired number of ALB replicas that you want in your cluster.
// 	resize	boolean
// 	If set to true, resizing of the ALB is done.

// 	state	string
// 	The state of the ALB.
// 	status	string
// 	The status of the ALB.
// 	vlanID	string
// 	The VLAN ID that the ALB is attached to.
// 	zone	string
// 	The zone where you want to add ALBs.
// 	}

func resourceIBMContainerAlbCreate() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerClassicIngressAlbCreate,
		Read:     resourceIBMContainerClassicIngressAlbRead,
		Update:   resourceIBMContainerClassicIngressAlbUpdate,
		Delete:   resourceIBMContainerClassicIngressAlbDelete,
		Exists:   resourceIBMContainerClassicIngressAlbExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			//post req
			"enableByDefault": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "If set to true, the ALB is enabled by default.",
			},
			"ingressImage": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The type of Ingress image that you want to use for your ALB deployment.",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The IP address that you want to assign to the ALB.",
			},
			"nlbVersion": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The version of the network load balancer that you want to use for the ALB.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of ALB that you want to create.",
			},
			"vlanID": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VLAN ID that you want to use for your ALBs.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone where you want to deploy the ALB.",
			},

			//response
			"albBuild": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The build number of the ALB.",
				//DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"albID": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the application load balancer (ALB).",
			},
			"albType": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The type of ALB.",
			},
			"albip": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The public IP address that exposes the ALB.",
			},
			"authBuild": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The auth build of the ALB.",
			},
			"clusterID": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    true,
				Description: "The ID of the cluster that the ALB belongs to.",
			},
			"createdDate": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Description: "The date the ALB was created.",
			},
			"disableDeployment": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    true,
				Description: "The public IP address that exposes the ALB.",
			},
			"enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Set to true to enable the ALB, or false to disable the ALB for the cluster.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The name of the cluster that the ALB belongs to.",
			},
			"numOfInstances": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "Desired number of ALB replicas that you want in your cluster.",
			},
			"resize": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "If set to true, resizing of the ALB is done.",
			},
			"state": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The state of the ALB.",
			},
			"status": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				ForceNew:    true,
				Description: "The status of the ALB.",
			},
		},
	}
}

func resourceIBMContainerClassicIngressAlbCreate(d *schema.ResourceData, meta interface{}) error {

	//create incgressALB config

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
	if v, ok := d.GetOkExists("vlanID"); ok {
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

	enableByDefault := d.Get("enableByDefault").(bool)
	ingressImage := d.Get("ingressImage").(string)
	ip := d.Get("ip").(string)
	nlbVersion := d.Get("nlbVersion").(string)

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
	albResp, err := albAPI.CreateALB(params, "testCluster", targetEnv)

	d.SetId(albResp.alb)
	return nil
}

func resourceIBMContainerClassicIngressAlbRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerClassicIngressAlbUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerClassicIngressAlbDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerClassicIngressAlbExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}

// //v1
// {
// 	"cluster":"string" //mandatory
// 	"enableByDefault": true,
// 	"ingressImage": "string",
// 	"ip": "string",
// 	"nlbVersion": "string",
// 	"type": "string", //mandatory
// 	"vlanID": "string", //mandatory
// 	"zone": "string" //mandatory
//   }

// //v2
//   {
// 	"cluster": "string",  //mandatory
// 	"enableByDefault": true,
// 	"ingressImage": "string",
// 	"type": "string", //mandatory
// 	"zone": "string" //mandatory
//   }

//https://github.com/IBM-Cloud/bluemix-go/blob/master/api/container/containerv1/alb.go
// add support for create ALB
// probably create a model struct for create alb  - armada-model/model/ingress  type ALB struct
// https://github.ibm.com/alchemy-containers/armada-model/blob/0d0e847de4a281f82a23044a4c1d57f16517ad87/model/ingress/ingress.go#L376

// https://github.ibm.com/alchemy-containers/armada-cli/blob/master/plugin/commands/cmdalb/alb.go
//

//
