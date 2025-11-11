---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_dr_summary_response"
description: |-
  Get information about pdr_get_dr_summary_response
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_dr_summary_response

Provides a read-only data source to retrieve information about a pdr_get_dr_summary_response. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_dr_summary_response" "pdr_get_dr_summary_response" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_dr_summary_response.
* `managed_vm_list` - (Map) A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).
* `orchestrator_details` - (List) Contains details about the orchestrator configuration.
Nested schema for **orchestrator_details**:
	* `last_updated_orchestrator_deployment_time` - (String) Deployment time of primary orchestrator VM.
	* `last_updated_standby_orchestrator_deployment_time` - (String) Deployment time of StandBy orchestrator VM.
	* `latest_orchestrator_time` - (String) Latest Orchestrator Time in COS.
	* `location_id` - (String) Location identifier.
	* `mfa_enabled` - (String) Multi Factor Authentication Enabled or not.
	* `orch_ext_connectivity_status` - (String) External connectivity status of the orchestrator.
	* `orch_standby_node_addition_status` - (String) Status of standby node addition.
	* `orchestrator_cluster_message` - (String) Message regarding orchestrator cluster status.
	* `orchestrator_config_status` - (String) Configuration status of the orchestrator.
	* `orchestrator_group_leader` - (String) Leader node of the orchestrator group.
	* `orchestrator_location_type` - (String) Type of orchestrator Location.
	* `orchestrator_name` - (String) Name of the primary orchestrator.
	* `orchestrator_status` - (String) Status of the primary orchestrator.
	* `orchestrator_workspace_name` - (String) Name of the orchestrator workspace.
	* `proxy_ip` - (String) IP address of the proxy.
	* `schematic_workspace_name` - (String) Name of the schematic workspace.
	* `schematic_workspace_status` - (String) Status of the schematic workspace.
	* `ssh_key_name` - (String) SSH key name used for the orchestrator.
	* `standby_orchestrator_name` - (String) Name of the standby orchestrator.
	* `standby_orchestrator_status` - (String) Status of the standby orchestrator.
	* `standby_orchestrator_workspace_name` - (String) Name of the standby orchestrator workspace.
	* `transit_gateway_name` - (String) Name of the transit gateway.
	* `vpc_name` - (String) Name of the VPC.
* `service_details` - (List) Contains details about the DR automation service.
Nested schema for **service_details**:
	* `crn` - (String) Cloud Resource Name identifier.
	* `deployment_name` - (String) Name of the deployment.
	* `description` - (String) Description of the primary service.
	* `is_ksys_ha` - (Boolean) Flag indicating if KSYS HA is enabled.
	* `plan_name` - (String) plan name.
	* `primary_ip_address` - (String) IP address of the primary service.
	* `primary_orchestrator_dashboard_url` - (String) Primary Orchestrator Dashboard URL.
	* `recovery_location` - (String) Location for disaster recovery.
	* `resource_group` - (String) Resource group name.
	* `standby_description` - (String) Description of the standby service.
	* `standby_ip_address` - (String) IP address of the standby service.
	* `standby_orchestrator_dashboard_url` - (String) Standby Orchestrator Dashboard URL.
	* `standby_status` - (String) Current status of the standby service.
	* `status` - (String) Current status of the primary service.

