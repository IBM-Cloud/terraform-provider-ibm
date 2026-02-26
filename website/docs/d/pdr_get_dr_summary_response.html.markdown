---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_dr_summary_response"
description: |-
  Get information about pdr_get_dr_summary_response
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_dr_summary_response

Retrieves the disaster recovery (DR) summary details for the specified service instance, including key configuration, status information and managed vm details.

## Example Usage

```hcl
data "ibm_pdr_get_dr_summary_response" "pdr_get_dr_summary_response" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9:"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document (Required, Forces new resource, String) (ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, Forces new resource, String) ID of the service instance.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_dr_summary_response.
* `managed_vm_list` - (Map) A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).
* `orchestrator_details` - (List) Contains details about the orchestrator configuration.
Nested schema for **orchestrator_details**:
	* `last_updated_orchestrator_deployment_time` - (String) The deployment time of primary orchestrator VM.
	* `last_updated_standby_orchestrator_deployment_time` - (String) The deployment time of StandBy orchestrator VM.
	* `latest_orchestrator_time` - (String) Latest Orchestrator Time in COS.
	* `location_id` - (String) The unique identifier of location.
	* `mfa_enabled` - (String) indicates if Multi Factor Authentication is enabled or not.
	* `orch_ext_connectivity_status` - (String) The external connectivity status of the orchestrator.
	* `orch_standby_node_addition_status` - (String) The status of standby node addition.
	* `orchestrator_cluster_message` - (String) The message regarding orchestrator cluster status.
	* `orchestrator_config_status` - (String) The configuration status of the orchestrator.
	* `orchestrator_group_leader` - (String) The leader node of the orchestrator group.
	* `orchestrator_location_type` - (String) The type of orchestrator Location.
	* `orchestrator_name` - (String) The name of the primary orchestrator.
	* `orchestrator_status` - (String) The status of the primary orchestrator.
	* `orchestrator_workspace_name` - (String) The name of the orchestrator workspace.
	* `proxy_ip` - (String) The IP address of the proxy.
	* `schematic_workspace_name` - (String) The name of the schematic workspace.
	* `schematic_workspace_status` - (String) The status of the schematic workspace.
	* `ssh_key_name` - (String) SSH key name used for the orchestrator.
	* `standby_orchestrator_name` - (String) The name of the standby orchestrator.
	* `standby_orchestrator_status` - (String) The status of the standby orchestrator.
	* `standby_orchestrator_workspace_name` - (String) The name of the standby orchestrator workspace.
	* `transit_gateway_name` - (String) The name of the transit gateway.
	* `vpc_name` - (String) The name of the VPC.
	* `api_key` - (String) api key.  
	* `standby_ssh_key_name` - (String) SSH key name used for the standby orchestrator.
* `service_details` - (List) Contains details about the DR automation service.
Nested schema for **service_details**:
	* `crn` - (String) The deployment crn.
	* `deployment_name` - (String) The name of the deployment.
	* `description` - (String) The Service description.
	* `orchestrator_ha` - (Boolean) The flag indicating whether orchestartor HA is enabled.
	* `plan_name` - (String) The plan name.
	* `primary_ip_address` - (String) The service Orchestator primary IP address.
	* `primary_orchestrator_dashboard_url` - (String) The Primary Orchestrator Dashboard URL.
	* `recovery_location` - (String) The disaster recovery location.
	* `resource_group` - (String) The Resource group name.
	* `standby_description` - (String) The standby orchestrator current status details.
	* `standby_ip_address` - (String) The service Orchestator standby IP address.
	* `standby_orchestrator_dashboard_url` - (String) The Standby Orchestrator Dashboard URL.
	* `standby_status` - (String) The standby orchestrator current status.
	* `status` - (String) The Status of the service.
