// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIBMPdrIBMMaintainedOrchDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrIBMMaintainedOrchDetailsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service Instance ID.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"service_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the IBM Maintained Orchestrator service details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status of the service.",
						},
						"deployment_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the deployment.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Resource group name.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deployment crn.",
						},
						"plan_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The plan name of the specified instance.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the orchestrator is created.",
						},
					},
				},
			},
			"orchestrator_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the orchestrator details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"orchestrator_gui_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the orchestrator gui.",
						},
						"orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the orchestrator.",
						},
						"orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the orchestrator.",
						},
						"orchestrator_ip": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP of the orchestrator.",
						},
						"orchestrator_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the orchestrator.",
						},
						"orchestrator_gui_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Orchestrator URL to access the VMRM GUI.",
						},
						"orchestrator_cluster_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration status of the orchestrator.",
						},
						"orchestrator_cluster_config_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message regarding orchestrator cluster creation.",
						},
						"orchestrator_location_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of orchestrator Location.",
						},
						"orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the orchestrator workspace.",
						},
						"orchestrator_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the progress details of orchestrator creation.",
						},
						"orchestrator_username": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Denotes the username for the orchestrator VMRM GUI.",
						},
						"standby_orchestrator_gui_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the standby_orchestrator gui.",
						},
						"standby_orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the standby orchestrator VM.",
						},
						"standby_orchestrator_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the standby orchestrator VM.",
						},
						"standby_orchestrator_ip": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the standby orchestrator VM.",
						},
						"standby_orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current status of the standby orchestrator VM.",
						},
						"standby_orchestrator_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the standby orchestrator.",
						},
						"standby_orchestrator_gui_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GUI URL for accessing the standby orchestrator.",
						},
						"standby_orchestrator_username": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Admin username for the standby orchestrator.",
						},
						"standby_orchestrator_node_addition_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the standby orchestrator node.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPdrIBMMaintainedOrchDetailsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_ibm_maintained_orch_details", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getIBMMaintainedDeploymentOptions := &drautomationservicev1.GetIBMMaintainedDeploymentOptions{}

	getIBMMaintainedDeploymentOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getIBMMaintainedDeploymentOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	ibmOrchestratorGetDeploymentResponse, response, err := drAutomationServiceClient.GetIBMMaintainedDeploymentWithContext(context, getIBMMaintainedDeploymentOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetIBMMaintainedDeploymentWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetIBMMaintainedDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_ibm_maintained_orch_details", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrIBMMaintainedOrchDetailsID(d))

	if !core.IsNil(ibmOrchestratorGetDeploymentResponse.ServiceDetails) {
		serviceDetails := []map[string]interface{}{}
		serviceDetailsMap, err := DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchServiceDetailsToMap(ibmOrchestratorGetDeploymentResponse.ServiceDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_ibm_maintained_orch_details", "read", "service_details-to-map").GetDiag()
		}
		serviceDetails = append(serviceDetails, serviceDetailsMap)
		if err = d.Set("service_details", serviceDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_details: %s", err), "(Data) ibm_pdr_ibm_maintained_orch_details", "read", "set-service_details").GetDiag()
		}
	}

	if !core.IsNil(ibmOrchestratorGetDeploymentResponse.OrchestratorDetails) {
		orchestratorDetails := []map[string]interface{}{}
		orchestratorDetailsMap, err := DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchOrchestratorDetailsToMap(ibmOrchestratorGetDeploymentResponse.OrchestratorDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_ibm_maintained_orch_details", "read", "orchestrator_details-to-map").GetDiag()
		}
		orchestratorDetails = append(orchestratorDetails, orchestratorDetailsMap)
		if err = d.Set("orchestrator_details", orchestratorDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_details: %s", err), "(Data) ibm_pdr_ibm_maintained_orch_details", "read", "set-orchestrator_details").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrIBMMaintainedOrchDetailsID returns a reasonable ID for the list.
func dataSourceIBMPdrIBMMaintainedOrchDetailsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchServiceDetailsToMap(model *drautomationservicev1.IBMMaintainedOrchServiceDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.DeploymentName != nil {
		modelMap["deployment_name"] = *model.DeploymentName
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = *model.ResourceGroup
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.PlanName != nil {
		modelMap["plan_name"] = *model.PlanName
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	return modelMap, nil
}

func DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchOrchestratorDetailsToMap(model *drautomationservicev1.IBMMaintainedOrchOrchestratorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["orchestrator_gui_status"] = *model.OrchestratorGuiStatus
	modelMap["orchestrator_name"] = *model.OrchestratorName
	modelMap["orchestrator_status"] = *model.OrchestratorStatus
	modelMap["orchestrator_ip"] = *model.OrchestratorIP
	modelMap["orchestrator_id"] = *model.OrchestratorID
	modelMap["orchestrator_gui_url"] = *model.OrchestratorGuiURL
	modelMap["orchestrator_cluster_status"] = *model.OrchestratorClusterStatus
	modelMap["orchestrator_cluster_config_message"] = *model.OrchestratorClusterConfigMessage
	if model.OrchestratorLocationType != nil {
		modelMap["orchestrator_location_type"] = *model.OrchestratorLocationType
	}
	if model.OrchestratorWorkspaceName != nil {
		modelMap["orchestrator_workspace_name"] = *model.OrchestratorWorkspaceName
	}
	modelMap["orchestrator_description"] = *model.OrchestratorDescription
	modelMap["orchestrator_username"] = *model.OrchestratorUsername
	modelMap["standby_orchestrator_gui_status"] = *model.StandbyOrchestratorGuiStatus
	modelMap["standby_orchestrator_name"] = *model.StandbyOrchestratorName
	modelMap["standby_orchestrator_id"] = *model.StandbyOrchestratorID
	modelMap["standby_orchestrator_ip"] = *model.StandbyOrchestratorIP
	modelMap["standby_orchestrator_status"] = *model.StandbyOrchestratorStatus
	modelMap["standby_orchestrator_description"] = *model.StandbyOrchestratorDescription
	modelMap["standby_orchestrator_gui_url"] = *model.StandbyOrchestratorGuiURL
	modelMap["standby_orchestrator_username"] = *model.StandbyOrchestratorUsername
	modelMap["standby_orchestrator_node_addition_status"] = *model.StandbyOrchestratorNodeAdditionStatus
	return modelMap, nil
}
