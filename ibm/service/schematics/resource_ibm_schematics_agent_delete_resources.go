// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func ResourceIbmSchematicsAgentDeleteResources() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSchematicsAgentDeleteResourcesCreate,
		ReadContext:   resourceIbmSchematicsAgentDeleteResourcesRead,
		UpdateContext: resourceIbmSchematicsAgentDeleteResourcesUpdate,
		DeleteContext: resourceIbmSchematicsAgentDeleteResourcesDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"agent_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Agent ID to get the details of agent.",
			},
			"force": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Equivalent to -force options in the command line, default is false.",
			},
			"job_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job Id.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The agent resources destroy job updation time.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who ran the agent resources destroy job.",
			},
			"agent_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Agent version.",
			},
			"status_code": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Final result of the agent resources destroy job.",
			},
			"status_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The outcome of the agent resources destroy job, in a formatted log string.",
			},
			"log_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the full agent resources destroy job logs.",
			},
		},
	}
}

func resourceIbmSchematicsAgentDeleteResourcesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}
	session, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	iamAccessToken := session.Config.IAMAccessToken
	iamRefreshToken := session.Config.IAMRefreshToken

	deleteAgentResourcesOptions := &schematicsv1.DeleteAgentResourcesOptions{}
	ff := map[string]string{
		"Authorization": iamAccessToken,
		"refresh_token": iamRefreshToken,
	}
	deleteAgentResourcesOptions.Headers = ff
	deleteAgentResourcesOptions.RefreshToken = core.StringPtr(iamRefreshToken)
	deleteAgentResourcesOptions.SetAgentID(d.Get("agent_id").(string))

	response, err := schematicsClient.DeleteAgentResourcesWithContext(context, deleteAgentResourcesOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteAgentResourcesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteAgentResourcesWithContext failed %s\n%s", err, response))
	}

	getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
		Profile: core.StringPtr("detailed"),
	}
	getAgentDataOptions.SetAgentID(d.Get("agent_id").(string))
	getAgentDataOptions.Headers = ff

	agentData, response, err := schematicsClient.GetAgentDataWithContext(context, getAgentDataOptions)
	if err != nil {
		log.Printf("[DEBUG] GetAgentDataWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetAgentDataWithContext failed %s\n%s", err, response))
	}
	if agentData.RecentDestroyJob != nil {
		d.SetId(fmt.Sprintf("%s/%s", *deleteAgentResourcesOptions.AgentID, *agentData.RecentDestroyJob.JobID))
	} else {
		d.SetId(fmt.Sprintf("%s/%s", *deleteAgentResourcesOptions.AgentID, time.Now().UTC().String()))
	}
	log.Printf("[INFO] Agent : %s", *deleteAgentResourcesOptions.AgentID)

	_, err = isWaitForAgentDestroyResources(context, schematicsClient, *deleteAgentResourcesOptions.AgentID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Waiting for agent resources to be destroyed, failed %s", err))
	}

	return resourceIbmSchematicsAgentDeleteResourcesRead(context, d, meta)
}

func isWaitForAgentDestroyResources(context context.Context, schematicsClient *schematicsv1.SchematicsV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for agent (%s) resources to be destroyed.", id)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", agentProvisioningStatusCodeJobInProgress, agentProvisioningStatusCodeJobPending, agentProvisioningStatusCodeJobReadyToExecute, agentProvisioningStatusCodeJobStopInProgress},
		Target:     []string{agentProvisioningStatusCodeJobFinished, agentProvisioningStatusCodeJobFailed, agentProvisioningStatusCodeJobCancelled, agentProvisioningStatusCodeJobStopped, ""},
		Refresh:    agentDestroyRefreshFunc(schematicsClient, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(context)
}
func agentDestroyRefreshFunc(schematicsClient *schematicsv1.SchematicsV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
			AgentID: core.StringPtr(id),
			Profile: core.StringPtr("detailed"),
		}

		agent, response, err := schematicsClient.GetAgentData(getAgentDataOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Agent: %s\n%s", err, response)
		}
		if agent.RecentDestroyJob.StatusCode != nil {
			return agent, *agent.RecentDestroyJob.StatusCode, nil
		}
		return agent, agentProvisioningStatusCodeJobPending, nil
	}
}

func resourceIbmSchematicsAgentDeleteResourcesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
		Profile: core.StringPtr("detailed"),
	}

	getAgentDataOptions.SetAgentID(parts[0])
	agentData, response, err := schematicsClient.GetAgentDataWithContext(context, getAgentDataOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetAgentDataWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetAgentDataWithContext failed %s\n%s", err, response))
	}
	if agentData.RecentDestroyJob != nil {

		if err = d.Set("agent_id", parts[0]); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting agent_id: %s", err))
		}
		if err = d.Set("job_id", agentData.RecentDestroyJob.JobID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting job_id: %s", err))
		}
		if err = d.Set("updated_at", flex.DateTimeToString(agentData.RecentDestroyJob.UpdatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
		}
		if err = d.Set("updated_by", agentData.RecentDestroyJob.UpdatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
		}
		if err = d.Set("agent_version", agentData.Version); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting agent_version: %s", err))
		}
		if err = d.Set("status_code", agentData.RecentDestroyJob.StatusCode); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status_code: %s", err))
		}
		if err = d.Set("status_message", agentData.RecentDestroyJob.StatusMessage); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status_message: %s", err))
		}
		if err = d.Set("log_url", agentData.RecentDestroyJob.LogURL); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting log_url: %s", err))
		}

	}
	return nil
}

func resourceIbmSchematicsAgentDeleteResourcesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}
	session, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	iamAccessToken := session.Config.IAMAccessToken
	iamRefreshToken := session.Config.IAMRefreshToken
	deleteAgentResourcesOptions := &schematicsv1.DeleteAgentResourcesOptions{}
	ff := map[string]string{
		"Authorization": iamAccessToken,
		"refresh_token": iamRefreshToken,
	}
	deleteAgentResourcesOptions.Headers = ff
	deleteAgentResourcesOptions.RefreshToken = core.StringPtr(iamRefreshToken)

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteAgentResourcesOptions.SetAgentID(parts[0])

	hasChange := false

	if d.HasChange("agent_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "agent_id"))
	}

	if hasChange {
		response, err := schematicsClient.DeleteAgentResourcesWithContext(context, deleteAgentResourcesOptions)
		if err != nil {
			log.Printf("[DEBUG] DeleteAgentResourcesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("DeleteAgentResourcesWithContext failed %s\n%s", err, response))
		}
		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
			Profile: core.StringPtr("detailed"),
		}
		getAgentDataOptions.SetAgentID(d.Get("agent_id").(string))
		getAgentDataOptions.Headers = ff

		agentData, response, err := schematicsClient.GetAgentDataWithContext(context, getAgentDataOptions)
		if err != nil {
			log.Printf("[DEBUG] GetAgentDataWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetAgentDataWithContext failed %s\n%s", err, response))
		}
		if agentData.RecentDestroyJob != nil {
			d.SetId(fmt.Sprintf("%s/%s", *deleteAgentResourcesOptions.AgentID, *agentData.RecentDestroyJob.JobID))
		} else {
			d.SetId(fmt.Sprintf("%s/%s", *deleteAgentResourcesOptions.AgentID, time.Now().UTC().String()))
		}

		_, err = isWaitForAgentDestroyResources(context, schematicsClient, parts[0], d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Waiting for agent resources to be destroyed, failed %s", err))
		}
	}

	return resourceIbmSchematicsAgentDeleteResourcesRead(context, d, meta)
}

func resourceIbmSchematicsAgentDeleteResourcesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
