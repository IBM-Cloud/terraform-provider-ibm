---

layout: "ibm"
page_title: "IBM : ibm_pdr_dr_summary_response"
description: |-
Get information about pdr_dr_summary_response
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
--------------------------------------------------------------------------------

# ibm_pdr_dr_summary_response

Retrieves the disaster recovery (DR) summary details for the specified service instance, including key configuration, status information and managed vm details.

## Example Usage

```hcl
data "ibm_pdr_dr_summary_response" "pdr_dr_summary_response" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `managed_vm_list` - (List) Array of mamagedVMs for the instanceID.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

Nested schema for **managed_vm_list**:
	* `core` - (String) The Number of cores assigned to the managed vitual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+(\\.\\d+)?$/`.
	* `dr_average_time` - (String) The DR operation average time(in minutes) for the managed virtual machine.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+$/`.
	* `dr_region` - (String) The name of the region where the virtual machine is recovered.
	* Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `memory` - (String) The amount of memory (in GB) assigned to the managed virtual machine.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d+(\\.\\d+)?$/`.
	* `region` - (String) The source region where the managed virtual machine is deployed.
	* Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `vm_id` - (String) The id of the irtual machine.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `vm_name` - (String) The name of the managed virtual machine.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `workgroup_name` - (String) The name of the workgroup where the managed virtual machine is added for disaster recovery.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `workspace_name` - (String) The Name of the power virtual server workspace.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.

* `orchestrator_details` - (List) Contains details about the orchestrator configuration.

Nested schema for **orchestrator_details**:
	* `last_updated_orchestrator_deployment_time` - (String) The deployment time of primary orchestrator VM.
	* Constraints: The maximum length is `1048` characters. The minimum length is `20` characters.
	* `last_updated_standby_orchestrator_deployment_time` - (String) The deployment time of StandBy orchestrator VM.
	* Constraints: The maximum length is `1048` characters. The minimum length is `20` characters.
	* `latest_orchestrator_time` - (String) Latest Orchestrator Time in COS.
	* Constraints: The maximum length is `1048` characters. The minimum length is `20` characters.
	* `location_id` - (String) The unique identifier of location.
	* Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `mfa_enabled` - (String) indicates if Multi Factor Authentication is enabled or not.
	* Constraints: The maximum length is `5` characters. The minimum length is `4` characters. The value must match regular expression `/^(true|false)$/`.
	* `orch_ext_connectivity_status` - (String) The external connectivity status of the orchestrator.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orch_standby_node_addition_status` - (String) The status of standby node addition.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_cluster_message` - (String) The message regarding orchestrator cluster status.
	* Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `orchestrator_config_status` - (String) The configuration status of the orchestrator.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_group_leader` - (String) The leader node of the orchestrator group.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_location_type` - (String) The type of orchestrator Location.
	* Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_name` - (String) The name of the primary orchestrator.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_status` - (String) The status of the primary orchestrator.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_workspace_name` - (String) The name of the orchestrator workspace.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `proxy_ip` - (String) The IP address of the proxy.
	* Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-_:]+$/`.
	* `schematic_workspace_name` - (String) The name of the schematic workspace.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_.]+$/`.
	* `schematic_workspace_status` - (String) The status of the schematic workspace.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `ssh_key_name` - (String) SSH key name used for the orchestrator.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_name` - (String) The name of the standby orchestrator.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_status` - (String) The status of the standby orchestrator.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_workspace_name` - (String) The name of the standby orchestrator workspace.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_ssh_key_name` - (String) SSH key name used for the standby orchestrator.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `transit_gateway_name` - (String) The name of the transit gateway.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `vpc_name` - (String) The name of the VPC.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.

* `service_details` - (List) Contains details about the DR automation service.

Nested schema for **service_details**:
	* `crn` - (String) The deployment crn.
	* Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
	* `deployment_name` - (String) The name of the deployment.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `description` - (String) The Service description.
	* Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `orchestrator_ha` - (Boolean) The flag indicating whether orchestartor HA is enabled.
	* `plan_name` - (String) The plan name.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_ ]+$/`.
	* `primary_ip_address` - (String) The service Orchestator primary IP address.
	* Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
	* `primary_orchestrator_dashboard_url` - (String) The Primary Orchestrator Dashboard URL.
	* Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `recovery_location` - (String) The disaster recovery location.
	* Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `resource_group` - (String) The Resource group name.
	* Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_description` - (String) The standby orchestrator current status details.
	* Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `standby_ip_address` - (String) The service Orchestator standby IP address.
	* Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
	* `standby_orchestrator_dashboard_url` - (String) The Standby Orchestrator Dashboard URL.
	* Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `standby_status` - (String) The standby orchestrator current status.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `status` - (String) The Status of the service.
	* Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
