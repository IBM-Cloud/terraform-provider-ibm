---
layout: "ibm"
page_title: "IBM : ibm_pdr_ibm_maintained_orch_details"
description: |-
Get information about pdr_ibm_maintained_orch_details
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
--------------------------------------------------------------------------------

# ibm_pdr_ibm_maintained_orch_details

Retrieves information about IBM-maintained Orchestrator details for the specified service instance.


## Example Usage

```hcl
data "ibm_pdr_ibm_maintained_orch_details" "pdr_ibm_maintained_orch_details" {
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
* `orchestrator_details` - (List) Contains details about the orchestrator details.
Nested schema for **orchestrator_details**:
	* `orchestrator_gui_status` - (String) Status of the orchestrator gui.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_name` - (String) The name of the orchestrator.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_status` - (String) The status of the orchestrator.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_ip` - (String) The IP of the orchestrator.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
	* `orchestrator_id` - (String) The ID of the orchestrator.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_gui_url` - (String) Orchestrator URL to access the VMRM GUI.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `orchestrator_cluster_status` - (String) The configuration status of the orchestrator.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_cluster_config_message` - (String) The message regarding orchestrator cluster creation.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `orchestrator_location_type` - (String) The type of orchestrator Location.
	  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_workspace_name` - (String) The name of the orchestrator workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `orchestrator_description` - (String) Indicates the progress details of orchestrator creation.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `orchestrator_username` - (String) Denotes the username for the orchestrator VMRM GUI.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_gui_status` - (String) Status of the standby_orchestrator gui.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_name` - (String) The name of the standby orchestrator VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_id` - (String) The ID of the standby orchestrator VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_ip` - (String) The IP address of the standby orchestrator VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
	* `standby_orchestrator_status` - (String) The current status of the standby orchestrator VM.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_description` - (String) Description of the standby orchestrator.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
	* `standby_orchestrator_gui_url` - (String) GUI URL for accessing the standby orchestrator.
     * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `standby_orchestrator_username` - (String) Admin username for the standby orchestrator.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `standby_orchestrator_node_addition_status` - (String) Status of the standby orchestrator node.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `service_details` - (List) Contains details about the IBM Maintained Orchestrator service details.
Nested schema for **service_details**:

	* `status` - (String) The Status of the service.
      * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `deployment_name` - (String) The name of the deployment.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `resource_group` - (String) The Resource group name.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `crn` - (String) The deployment crn.
      * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
	* `plan_name` - (String) The plan name of the specified instance.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_ ]+$/`.
	* `region` - (String) Region where the orchestrator is created.
      * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
