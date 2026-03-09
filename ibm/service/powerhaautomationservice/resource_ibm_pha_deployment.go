// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func ResourceIBMPhaDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPhaDeploymentCreate,
		ReadContext:   resourceIBMPhaDeploymentRead,
		DeleteContext: resourceIBMPhaDeploymentDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "pha_instance_id"),
				Description:  "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "accept_language"),
				Description:  "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "if_none_match"),
				Description:  "ETag for conditional requests (optional).",
			},
			"cluster_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Type of PowerHA cluster being deployed.",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9._:-]+$`), "invalid format"),
			},

			"configure_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Configuration type for the deployment.",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9._:-]+$`), "invalid format"),
			},

			"location_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Identifier for the deployment location.",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9._:-]+$`), "invalid format"),
			},

			"primary_cluster_nodes": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "List of primary cluster node VM IDs.",
				MaxItems:    50,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 36),
				},
			},

			"standby_cluster_nodes": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "List of standby cluster node VM IDs.",
				MaxItems:    50,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(0, 36),
				},
			},
			"primary_cluster_node": &schema.Schema{
				Type: schema.TypeList,
				// Optional:    true,
				// ForceNew:    true,
				Computed:    true,
				Description: "List of primary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "Number of CPU cores allocated to the node.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address assigned to the virtual machine.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory allocated to the virtual machine in MB or GB.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Region where the virtual machine is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the virtual machine.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the virtual machine.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Current operational status of the virtual machine.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Workspace identifier associated with the node.",
						},
					},
				},
			},
			"primary_location": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "primary_location"),
				Description:  "Primary cluster location.",
			},
			"primary_workspace": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "primary_workspace"),
				Description:  "Primary workspace identifier.",
			},
			"secondary_location": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "secondary_location"),
				Description:  "Secondary cluster location.",
			},
			"secondary_workspace": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_deployment", "secondary_workspace"),
				Description:  "Secondary workspace identifier.",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API key used for authentication to the deployment service.",
			},
			"cloud_account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud account identifier.",
			},
			"connectivity_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of network connectivity.",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp expressing creation time.",
			},
			"custom_network": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of custom network CIDRs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"deprovision_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp expressing deprovision time.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Global unique identifier.",
			},
			"is_duplicate": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether deployment is duplicate.",
			},
			"plan_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier for the service plan.",
			},
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of service plan.",
			},
			"powerha_cluster_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the PowerHA cluster.",
			},
			"powerha_cluster_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of PowerHA cluster.",
			},
			"powerha_level": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PowerHA version level.",
			},
			"primary_region_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the primary workspace region.",
			},
			"primary_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the primary workspace.",
			},
			"provision_end_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp provisioning completed.",
			},
			"provision_start_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp provisioning started.",
			},
			"provision_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current provision status.",
			},
			"region_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Deployment region identifier.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the resource group.",
			},
			"resource_group_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of associated resource group.",
			},
			"resource_instance": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource instance identifier.",
			},
			"secondary_cluster_nodes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of secondary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Number of CPU cores allocated to the node.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "IP address assigned to the virtual machine.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Memory allocated to the virtual machine in MB or GB.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Region where the virtual machine is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique identifier of the virtual machine.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the virtual machine.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Current operational status of the virtual machine.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Workspace identifier associated with the node.",
						},
					},
				},
			},
			"service_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of provisioned service.",
			},
			"service_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier for the service.",
			},
			"service_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of service.",
			},
			"standby_region_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the standby workspace region.",
			},
			"standby_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the standby workspace.",
			},
			"user_tags": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User defined tags.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Provision request identifier.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMPhaDeploymentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "pha_instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9-]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "accept_language",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "if_none_match",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "primary_location",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9._:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             16,
		},
		validate.ValidateSchema{
			Identifier:                 "primary_workspace",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9._:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             2048,
		},
		validate.ValidateSchema{
			Identifier:                 "secondary_location",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9._:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             16,
		},
		validate.ValidateSchema{
			Identifier:                 "secondary_workspace",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9._:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             2048,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pha_deployment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPhaDeploymentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createPhaDeploymentOptions := &powerhaautomationservicev1.CreatePhaDeploymentOptions{}

	createPhaDeploymentOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	createPhaDeploymentOptions.SetLocationID(d.Get("location_id").(string))
	createPhaDeploymentOptions.SetPrimaryWorkspace(d.Get("primary_workspace").(string))
	if _, ok := d.GetOk("cluster_type"); ok {
		createPhaDeploymentOptions.SetClusterType(d.Get("cluster_type").(string))
	}
	if _, ok := d.GetOk("configure_type"); ok {
		createPhaDeploymentOptions.SetConfigureType(d.Get("configure_type").(string))
	}
	if _, ok := d.GetOk("primary_cluster_nodes"); ok {
		var primaryClusterNodes []string
		for _, v := range d.Get("primary_cluster_nodes").([]interface{}) {
			primaryClusterNodesItem := v.(string)
			primaryClusterNodes = append(primaryClusterNodes, primaryClusterNodesItem)
		}
		createPhaDeploymentOptions.SetPrimaryClusterNodes(primaryClusterNodes)
	}
	if _, ok := d.GetOk("standby_cluster_nodes"); ok {
		var standbyClusterNodes []string
		for _, v := range d.Get("standby_cluster_nodes").([]interface{}) {
			standbyClusterNodesItem := v.(string)
			standbyClusterNodes = append(standbyClusterNodes, standbyClusterNodesItem)
		}
		createPhaDeploymentOptions.SetStandbyClusterNodes(standbyClusterNodes)
	}
	if _, ok := d.GetOk("primary_location"); ok {
		createPhaDeploymentOptions.SetPrimaryLocation(d.Get("primary_location").(string))
	}
	if _, ok := d.GetOk("secondary_location"); ok {
		createPhaDeploymentOptions.SetSecondaryLocation(d.Get("secondary_location").(string))
	}
	if _, ok := d.GetOk("secondary_workspace"); ok {
		createPhaDeploymentOptions.SetSecondaryWorkspace(d.Get("secondary_workspace").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createPhaDeploymentOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		createPhaDeploymentOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	_, response, err := powerhaAutomationServiceClient.CreatePhaDeploymentWithContext(context, createPhaDeploymentOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreatePhaDeploymentWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreatePhaDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_deployment", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()

		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePhaDeploymentWithContext failed: %s", err.Error()), "ibm_pha_deployment", "create")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *createPhaDeploymentOptions.PhaInstanceID))

	return resourceIBMPhaDeploymentRead(context, d, meta)
}

func resourceIBMPhaDeploymentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPhaDeploymentOptions := &powerhaautomationservicev1.GetPhaDeploymentOptions{}
	fmt.Println(d.Id())
	getPhaDeploymentOptions.SetPhaInstanceID(d.Id())
	if _, ok := d.GetOk("if_none_match"); ok {
		getPhaDeploymentOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaDeploymentResponse, response, err := powerhaAutomationServiceClient.GetPhaDeploymentWithContext(context, getPhaDeploymentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		detailedMsg := fmt.Sprintf("GetPhaDeploymentWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetPhaDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_deployment", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()

		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPhaDeploymentWithContext failed: %s", err.Error()), "ibm_pha_deployment", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryClusterNodes) {
		primaryClusterNodes := []map[string]interface{}{}
		for _, primaryClusterNodesItem := range phaDeploymentResponse.PrimaryClusterNodes {
			primaryClusterNodesItemMap, err := ResourceIBMPhaDeploymentClusterNodeInfoToMap(&primaryClusterNodesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "primary_cluster_nodes-to-map").GetDiag()
			}
			primaryClusterNodes = append(primaryClusterNodes, primaryClusterNodesItemMap)
		}
		if err = d.Set("primary_cluster_node", primaryClusterNodes); err != nil {
			err = fmt.Errorf("Error setting primary_cluster_nodes: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-primary_cluster_nodes").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PrimaryLocation) {
		if err = d.Set("primary_location", phaDeploymentResponse.PrimaryLocation); err != nil {
			err = fmt.Errorf("Error setting primary_location: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-primary_location").GetDiag()
		}
	}
	if err = d.Set("primary_workspace", phaDeploymentResponse.PrimaryWorkspace); err != nil {
		err = fmt.Errorf("Error setting primary_workspace: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-primary_workspace").GetDiag()
	}
	if !core.IsNil(phaDeploymentResponse.SecondaryLocation) {
		if err = d.Set("secondary_location", phaDeploymentResponse.SecondaryLocation); err != nil {
			err = fmt.Errorf("Error setting secondary_location: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-secondary_location").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.SecondaryWorkspace) {
		if err = d.Set("secondary_workspace", phaDeploymentResponse.SecondaryWorkspace); err != nil {
			err = fmt.Errorf("Error setting secondary_workspace: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-secondary_workspace").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.APIKey) {
		if err = d.Set("api_key", phaDeploymentResponse.APIKey); err != nil {
			err = fmt.Errorf("Error setting api_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-api_key").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.CloudAccountID) {
		if err = d.Set("cloud_account_id", phaDeploymentResponse.CloudAccountID); err != nil {
			err = fmt.Errorf("Error setting cloud_account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-cloud_account_id").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ConnectivityType) {
		if err = d.Set("connectivity_type", phaDeploymentResponse.ConnectivityType); err != nil {
			err = fmt.Errorf("Error setting connectivity_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-connectivity_type").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.CreationTime) {
		if err = d.Set("creation_time", phaDeploymentResponse.CreationTime); err != nil {
			err = fmt.Errorf("Error setting creation_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-creation_time").GetDiag()
		}
	}
	if err = d.Set("custom_network", phaDeploymentResponse.CustomNetwork); err != nil {
		err = fmt.Errorf("Error setting custom_network: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-custom_network").GetDiag()
	}
	if !core.IsNil(phaDeploymentResponse.DeprovisionTime) {
		if err = d.Set("deprovision_time", phaDeploymentResponse.DeprovisionTime); err != nil {
			err = fmt.Errorf("Error setting deprovision_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-deprovision_time").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.GUID) {
		if err = d.Set("guid", phaDeploymentResponse.GUID); err != nil {
			err = fmt.Errorf("Error setting guid: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-guid").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.IsDuplicate) {
		if err = d.Set("is_duplicate", phaDeploymentResponse.IsDuplicate); err != nil {
			err = fmt.Errorf("Error setting is_duplicate: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-is_duplicate").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PlanID) {
		if err = d.Set("plan_id", phaDeploymentResponse.PlanID); err != nil {
			err = fmt.Errorf("Error setting plan_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-plan_id").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PlanName) {
		if err = d.Set("plan_name", phaDeploymentResponse.PlanName); err != nil {
			err = fmt.Errorf("Error setting plan_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-plan_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PowerhaClusterName) {
		if err = d.Set("powerha_cluster_name", phaDeploymentResponse.PowerhaClusterName); err != nil {
			err = fmt.Errorf("Error setting powerha_cluster_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-powerha_cluster_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PowerhaClusterType) {
		if err = d.Set("powerha_cluster_type", phaDeploymentResponse.PowerhaClusterType); err != nil {
			err = fmt.Errorf("Error setting powerha_cluster_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-powerha_cluster_type").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PowerhaLevel) {
		if err = d.Set("powerha_level", phaDeploymentResponse.PowerhaLevel); err != nil {
			err = fmt.Errorf("Error setting powerha_level: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-powerha_level").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PrimaryRegionName) {
		if err = d.Set("primary_region_name", phaDeploymentResponse.PrimaryRegionName); err != nil {
			err = fmt.Errorf("Error setting primary_region_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-primary_region_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.PrimaryWorkspaceName) {
		if err = d.Set("primary_workspace_name", phaDeploymentResponse.PrimaryWorkspaceName); err != nil {
			err = fmt.Errorf("Error setting primary_workspace_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-primary_workspace_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ProvisionEndTime) {
		if err = d.Set("provision_end_time", phaDeploymentResponse.ProvisionEndTime); err != nil {
			err = fmt.Errorf("Error setting provision_end_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-provision_end_time").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ProvisionStartTime) {
		if err = d.Set("provision_start_time", phaDeploymentResponse.ProvisionStartTime); err != nil {
			err = fmt.Errorf("Error setting provision_start_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-provision_start_time").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ProvisionStatus) {
		if err = d.Set("provision_status", phaDeploymentResponse.ProvisionStatus); err != nil {
			err = fmt.Errorf("Error setting provision_status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-provision_status").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.RegionID) {
		if err = d.Set("region_id", phaDeploymentResponse.RegionID); err != nil {
			err = fmt.Errorf("Error setting region_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-region_id").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ResourceGroup) {
		if err = d.Set("resource_group", phaDeploymentResponse.ResourceGroup); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-resource_group").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ResourceGroupCRN) {
		if err = d.Set("resource_group_crn", phaDeploymentResponse.ResourceGroupCRN); err != nil {
			err = fmt.Errorf("Error setting resource_group_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-resource_group_crn").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ResourceInstance) {
		if err = d.Set("resource_instance", phaDeploymentResponse.ResourceInstance); err != nil {
			err = fmt.Errorf("Error setting resource_instance: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-resource_instance").GetDiag()
		}
	}
	secondaryClusterNodes := []map[string]interface{}{}
	for _, secondaryClusterNodesItem := range phaDeploymentResponse.SecondaryClusterNodes {
		secondaryClusterNodesItemMap, err := ResourceIBMPhaDeploymentClusterNodeInfoToMap(&secondaryClusterNodesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "secondary_cluster_nodes-to-map").GetDiag()
		}
		secondaryClusterNodes = append(secondaryClusterNodes, secondaryClusterNodesItemMap)
	}
	if err = d.Set("secondary_cluster_nodes", secondaryClusterNodes); err != nil {
		err = fmt.Errorf("Error setting secondary_cluster_nodes: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-secondary_cluster_nodes").GetDiag()
	}
	if !core.IsNil(phaDeploymentResponse.ServiceDescription) {
		if err = d.Set("service_description", phaDeploymentResponse.ServiceDescription); err != nil {
			err = fmt.Errorf("Error setting service_description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-service_description").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ServiceID) {
		if err = d.Set("service_id", phaDeploymentResponse.ServiceID); err != nil {
			err = fmt.Errorf("Error setting service_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-service_id").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ServiceName) {
		if err = d.Set("service_name", phaDeploymentResponse.ServiceName); err != nil {
			err = fmt.Errorf("Error setting service_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-service_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.StandbyRegionName) {
		if err = d.Set("standby_region_name", phaDeploymentResponse.StandbyRegionName); err != nil {
			err = fmt.Errorf("Error setting standby_region_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-standby_region_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.StandbyWorkspaceName) {
		if err = d.Set("standby_workspace_name", phaDeploymentResponse.StandbyWorkspaceName); err != nil {
			err = fmt.Errorf("Error setting standby_workspace_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-standby_workspace_name").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.UserTags) {
		if err = d.Set("user_tags", phaDeploymentResponse.UserTags); err != nil {
			err = fmt.Errorf("Error setting user_tags: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-user_tags").GetDiag()
		}
	}
	if !core.IsNil(phaDeploymentResponse.ID) {

		if err = d.Set("pha_instance_id", extractInstanceIDFromCRN(*phaDeploymentResponse.ID)); err != nil {
			err = fmt.Errorf("Error setting pha_instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_deployment", "read", "set-pha_instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pha_deployment", "read", "set-etag").GetDiag()
	}

	return nil
}

func extractInstanceIDFromCRN(crn string) string {
	parts := strings.Split(crn, ":")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			return parts[i]
		}
	}
	return ""
}

func resourceIBMPhaDeploymentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIBMPhaDeploymentClusterNodeInfoToMap(model *powerhaautomationservicev1.ClusterNodeInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AgentStatus != nil {
		modelMap["agent_status"] = *model.AgentStatus
	}
	if model.Cores != nil {
		modelMap["cores"] = flex.Float64Value(model.Cores)
	}
	if model.IPAddress != nil {
		modelMap["ip_address"] = *model.IPAddress
	}
	if model.Memory != nil {
		modelMap["memory"] = flex.IntValue(model.Memory)
	}
	if model.PhaLevel != nil {
		modelMap["pha_level"] = *model.PhaLevel
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.VMID != nil {
		modelMap["vm_id"] = *model.VMID
	}
	if model.VMName != nil {
		modelMap["vm_name"] = *model.VMName
	}
	if model.VMStatus != nil {
		modelMap["vm_status"] = *model.VMStatus
	}
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = *model.WorkspaceID
	}
	return modelMap, nil
}
