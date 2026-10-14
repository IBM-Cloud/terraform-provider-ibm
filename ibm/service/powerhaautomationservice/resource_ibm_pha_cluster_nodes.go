// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"regexp"

	// "regexp"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPhaClusterNodes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPhaClusterNodesCreate,
		ReadContext:   resourceIBMPhaClusterNodesRead,
		UpdateContext: resourceIBMPhaClusterNodesUpdate,
		DeleteContext: resourceIBMPhaClusterNodesDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pha_cluster_nodes", "pha_instance_id"),
				Description: "Unique identifier of the provisioned instance.",
			},
			"accept_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pha_cluster_nodes", "accept_language"),
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pha_cluster_nodes", "if_none_match"),
				Description: "ETag for conditional requests (optional).",
			},

			"primary_cluster_nodes": {
				Type: schema.TypeSet,
				// Optional: true,
				Required: true,
				// ForceNew:         true,
				MinItems: 0,
				MaxItems: 8,
				// DiffSuppressFunc: flex.ApplyOnce,
				Description: "List of primary cluster node VM IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z0-9._:-]+$`),
						"must contain only alphanumeric characters, dot, underscore, colon or dash",
					),
				},
			},

			"secondary_cluster_nodes": {
				Type: schema.TypeSet,
				// Optional: true,
				Required: true,
				// ForceNew:         true,
				MinItems: 0,
				MaxItems: 7,
				// DiffSuppressFunc: flex.ApplyOnce,
				Description: "List of primary cluster node VM IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z0-9._:-]+$`),
						"must contain only alphanumeric characters, dot, underscore, colon or dash",
					),
				},
			},

			"primary_node_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details of the primary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique identifier of the VM.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the VM.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Region where the VM is deployed.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "ID of the workspace associated with the VM.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Number of CPU cores allocated to the VM.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Amount of memory allocated to the VM (in GB).",
						},
						"ip_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of IP addresses assigned to the VM.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Current status of the VM.",
						},
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"powerha_version_supported": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the installed PowerHA version is supported.",
						},
					},
				},
			},
			"secondary_node_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details of the secondary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique identifier of the VM.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the VM.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Region where the VM is deployed.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "ID of the workspace associated with the VM.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Number of CPU cores allocated to the VM.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Amount of memory allocated to the VM (in GB).",
						},
						"ip_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of IP addresses assigned to the VM.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Current status of the VM.",
						},
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"powerha_version_supported": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the installed PowerHA version is supported.",
						},
					},
				},
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier for this cluster node response.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMPhaClusterNodesValidator() *validate.ResourceValidator {
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
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pha_cluster_nodes", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPhaClusterNodesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createClusterNodeOptions := &powerhaautomationservicev1.CreateClusterNodeOptions{}

	createClusterNodeOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	// var primaryClusterNodes []string
	if v, ok := d.GetOk("primary_cluster_nodes"); ok {
		set := v.(*schema.Set)
		var primaryClusterNodes []string
		for _, item := range set.List() {
			primaryClusterNodes = append(primaryClusterNodes, item.(string))
		}
		createClusterNodeOptions.SetPrimaryClusterNodes(primaryClusterNodes)
	}
	// if _, ok := d.GetOk("secondary_cluster_nodes"); ok {
	// 	var secondaryClusterNodes []string
	// 	for _, v := range d.Get("secondary_cluster_nodes") {
	// 		secondaryClusterNodesItem := v.(string)
	// 		secondaryClusterNodes = append(secondaryClusterNodes, secondaryClusterNodesItem)
	// 	}
	// 	createClusterNodeOptions.SetSecondaryClusterNodes(secondaryClusterNodes)
	// }

	if j, ok := d.GetOk("secondary_cluster_nodes"); ok {
		set := j.(*schema.Set)
		var secondaryClusterNodes []string
		for _, item := range set.List() {
			secondaryClusterNodes = append(secondaryClusterNodes, item.(string))
		}
		createClusterNodeOptions.SetSecondaryClusterNodes(secondaryClusterNodes)
	}

	// secondaryNodes := make([]string, 0)

	// if v, ok := d.GetOk("secondary_cluster_nodes"); ok {
	// 	nodes, ok := v.([]interface{})
	// 	if !ok {
	// 		return diag.Errorf("secondary_cluster_nodes must be a list")
	// 	}

	// 	for _, node := range nodes {
	// 		secondaryNodes = append(secondaryNodes, node.(string))
	// 	}

	// }
	// if len(secondaryNodes) > 0 {
	// 	createClusterNodeOptions.SetSecondaryClusterNodes(secondaryNodes)
	// }

	if _, ok := d.GetOk("accept_language"); ok {
		createClusterNodeOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		createClusterNodeOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	clusterNodeResponse, response, err := powerhaAutomationServiceClient.CreateClusterNodeWithContext(context, createClusterNodeOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateClusterNodeWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateClusterNodeWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_cluster_nodes", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *clusterNodeResponse.ID))

	return resourceIBMPhaClusterNodesRead(context, d, meta)
}

func resourceIBMPhaClusterNodesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getClusterNodeOptions := &powerhaautomationservicev1.GetClusterNodeOptions{}

	getClusterNodeOptions.SetPhaInstanceID(d.Id())
	if _, ok := d.GetOk("if_none_match"); ok {
		getClusterNodeOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	clusterNodeResponse, response, err := powerhaAutomationServiceClient.GetClusterNodeWithContext(context, getClusterNodeOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetClusterNodeWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil && response.StatusCode == 404 {
			detailedMsg = fmt.Sprintf(
				"GetClusterNodeWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_cluster_nodes", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	primaryNodeDetails := []map[string]interface{}{}
	for _, primaryNodeDetailsItem := range clusterNodeResponse.PrimaryNodeDetails {
		primaryNodeDetailsItemMap, err := ResourceIBMPhaClusterNodesNodeDetailToMap(&primaryNodeDetailsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "primary_node_details-to-map").GetDiag()
		}
		primaryNodeDetails = append(primaryNodeDetails, primaryNodeDetailsItemMap)
	}
	if err = d.Set("primary_node_details", primaryNodeDetails); err != nil {
		err = fmt.Errorf("Error setting primary_node_details: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "set-primary_node_details").GetDiag()
	}
	// Extract vm_ids from primary_node_details
	primaryNodeDetailsresp := d.Get("primary_node_details").([]interface{})

	var primaryClusterNodeIDs []string

	for _, node := range primaryNodeDetailsresp {
		if node == nil {
			continue
		}

		nodeMap := node.(map[string]interface{})

		if vmID, ok := nodeMap["vm_id"].(string); ok && vmID != "" {
			primaryClusterNodeIDs = append(primaryClusterNodeIDs, vmID)
		}
	}

	// Set into primary_cluster_nodes (TypeSet)
	if err := d.Set("primary_cluster_nodes", schema.NewSet(schema.HashString, convertToInterfaceSlice(primaryClusterNodeIDs))); err != nil {
		return diag.FromErr(err)
	}
	secondaryNodeDetails := []map[string]interface{}{}
	for _, secondaryNodeDetailsItem := range clusterNodeResponse.SecondaryNodeDetails {
		secondaryNodeDetailsItemMap, err := ResourceIBMPhaClusterNodesNodeDetailToMap(&secondaryNodeDetailsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "secondary_node_details-to-map").GetDiag()
		}
		secondaryNodeDetails = append(secondaryNodeDetails, secondaryNodeDetailsItemMap)
	}
	if err = d.Set("secondary_node_details", secondaryNodeDetails); err != nil {
		err = fmt.Errorf("Error setting secondary_node_details: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "set-secondary_node_details").GetDiag()
	}
	// Extract vm_ids from secondary_node_details
	secondaryNodeDetailsresp := d.Get("secondary_node_details").([]interface{})

	var secondaryClusterNodeIDs []string

	for _, node := range secondaryNodeDetailsresp {
		if node == nil {
			continue
		}

		nodeMap := node.(map[string]interface{})

		if vmID, ok := nodeMap["vm_id"].(string); ok && vmID != "" {
			secondaryClusterNodeIDs = append(secondaryClusterNodeIDs, vmID)
		}
	}

	// Set into secondary_cluster_nodes (TypeList)
	if err := d.Set("secondary_cluster_nodes", schema.NewSet(schema.HashString, convertToInterfaceSlice(secondaryClusterNodeIDs))); err != nil {
		return diag.FromErr(err)
	}
	if !core.IsNil(clusterNodeResponse.ID) {
		if err = d.Set("pha_instance_id", clusterNodeResponse.ID); err != nil {
			err = fmt.Errorf("Error setting pha_instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "read", "set-pha_instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pha_cluster_nodes", "read", "set-etag").GetDiag()
	}

	return nil
}

func convertToInterfaceSlice(input []string) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}
func resourceIBMPhaClusterNodesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceID := d.Id()

	// Only run if there is a change
	if !d.HasChange("primary_cluster_nodes") {
		return resourceIBMPhaClusterNodesRead(ctx, d, meta)
	}

	oldRaw, newRaw := d.GetChange("primary_cluster_nodes")

	oldSet := oldRaw.(*schema.Set)
	newSet := newRaw.(*schema.Set)

	// -------------------------
	//  1. DELETE removed VMs
	// -------------------------
	removed := oldSet.Difference(newSet)
	if len(removed.List()) > 1 {
		msg := fmt.Sprintf("Only 1 VM can be deleted at a time")
		return diag.FromErr(fmt.Errorf(msg))
	}

	for _, item := range removed.List() {
		vmID := item.(string)

		opts := &powerhaautomationservicev1.DeleteClusterNodeOptions{}
		opts.SetPhaInstanceID(instanceID)
		opts.SetVMID(vmID)

		_, response, err := client.DeleteClusterNodeWithContext(ctx, opts)
		if err != nil {
			msg := fmt.Sprintf("Failed to delete VM %s: %s", vmID, err.Error())
			if response != nil {
				msg = fmt.Sprintf("Failed to delete VM %s: %s (status: %d, response: %s)",
					vmID, err.Error(), response.StatusCode, response.Result)
			}
			return diag.FromErr(fmt.Errorf(msg))
		}
	}

	// -------------------------
	//  2. ADD new VMs
	// -------------------------
	added := newSet.Difference(oldSet)
	if added.Len() > 0 {
		var vmIDs []string

		for _, item := range added.List() {
			vmIDs = append(vmIDs, item.(string))
		}

		opts := &powerhaautomationservicev1.CreateClusterNodeOptions{}
		opts.SetPhaInstanceID(instanceID)
		opts.SetPrimaryClusterNodes(vmIDs)

		_, response, err := client.CreateClusterNodeWithContext(ctx, opts)
		if err != nil {
			msg := fmt.Sprintf("Failed to add VMs: %s", err.Error())
			if response != nil {
				msg = fmt.Sprintf("Failed to add VMs: %s (status: %d, response: %s)",
					err.Error(), response.StatusCode, response.Result)
			}
			return diag.FromErr(fmt.Errorf(msg))
		}
	}

	// -------------------------
	// 3. Refresh state
	// -------------------------
	return resourceIBMPhaClusterNodesRead(ctx, d, meta)
}

func resourceIBMPhaClusterNodesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	// if err != nil {
	// 	tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "delete", "initialize-client")
	// 	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	// 	return tfErr.GetDiag()
	// }

	// deleteClusterNodeOptions := &powerhaautomationservicev1.DeleteClusterNodeOptions{}

	// parts, err := flex.SepIdParts(d.Id(), ":")
	// if err != nil {
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_cluster_nodes", "delete", "sep-id-parts").GetDiag()
	// }

	// deleteClusterNodeOptions.SetPhaInstanceID(parts[7])
	// if _, ok := d.GetOk("if_none_match"); ok {
	// 	deleteClusterNodeOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	// }

	// _, _, err = powerhaAutomationServiceClient.DeleteClusterNodeWithContext(context, deleteClusterNodeOptions)
	// if err != nil {
	// 	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteClusterNodeWithContext failed: %s", err.Error()), "ibm_pha_cluster_nodes", "delete")
	// 	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	// 	return tfErr.GetDiag()
	// }

	d.SetId("")

	return nil
}

func ResourceIBMPhaClusterNodesNodeDetailToMap(model *powerhaautomationservicev1.NodeDetail) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VMID != nil {
		modelMap["vm_id"] = *model.VMID
	}
	if model.VMName != nil {
		modelMap["vm_name"] = *model.VMName
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = *model.WorkspaceID
	}
	if model.Cores != nil {
		modelMap["cores"] = flex.Float64Value(model.Cores)
	}
	if model.Memory != nil {
		modelMap["memory"] = flex.Float64Value(model.Memory)
	}
	modelMap["ip_addresses"] = model.IPAddresses
	if model.VMStatus != nil {
		modelMap["vm_status"] = *model.VMStatus
	}
	if model.AgentStatus != nil {
		modelMap["agent_status"] = *model.AgentStatus
	}
	if model.PhaLevel != nil {
		modelMap["pha_level"] = *model.PhaLevel
	}
	modelMap["powerha_version_supported"] = *model.PowerhaVersionSupported
	return modelMap, nil
}
