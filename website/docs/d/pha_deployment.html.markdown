---
layout: "ibm"
page_title: "IBM : ibm_pha_deployment"
description: |-
  Get information about pha_deployment
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_deployment

Retrieves details of the specified PowerHA deployment, including its configuration, operational status, and related metadata.

## Example Usage

```hcl
data "ibm_pha_deployment" "pha_deployment" {
	if_none_match = ibm_pha_deployment.pha_deployment_instance.if_none_match
	pha_instance_id = ibm_pha_deployment.pha_deployment_instance.pha_instance_id
}
```
### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Provision request identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `cloud_account_id` - (String) Cloud account identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `connectivity_type` - (String) Type of network connectivity.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `creation_time` - (String) Timestamp expressing creationtime.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `deprovision_time` - (String) Timestamp expressing deprovision time.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$/`.
* `guid` - (String) Global unique identifier.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `is_duplicate` - (Boolean) Indicates whether deployment is duplicate.
* `plan_id` - (String) Identifier for the service plan.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `plan_name` - (String) Name of service plan.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `powerha_cluster_name` - (String) Name of the PowerHA cluster.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_cluster_type` - (String) Type of PowerHA cluster.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_level` - (String) PowerHA version level.
  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

* `primary_cluster_nodes_details` - (List) List of primary cluste nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **primary_cluster_nodes_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the node.
	  * Constraints: The maximum value is `16`. The minimum value is `1`.
	* `ip_address` - (String) IP address assigned to the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9.:]+$/`.
	* `memory` - (Integer) Memory allocated to the virtual machine in MB.
	  * Constraints: The maximum value is `1048576`. The minimum value is `1`.
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9.]+$/`.
	* `region` - (String) Region where the virtual machine is deployed.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `vm_id` - (String) Identifier of the Power virtual server instance.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current operational status of the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) Workspace identifier associated with the node.
	  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_location` - (String) Primary cluster location.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_region_name` - (String) name of the primary workspace region.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace` - (String) Primary workspace identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace_crn` - (String) CRN of the primary workspace.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[A-Za-z0-9._:-]+(:[A-Za-z0-9._:-]+)*$/`.
* `primary_workspace_name` - (String) name of the primary workspace.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_end_time` - (String) Time stamp provisioning completed.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_start_time` - (String) Time stamp provisioning started.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_status` - (String) Current provision status.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `region_id` - (String) Deployment region identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group` - (String) Name of the resource group.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group_crn` - (String) CRN of associated resource group.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:\/-]+$/`.
* `resource_instance` - (String) Resource instance identifier.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

* `secondary_cluster_nodes_details` - (List) List of secondary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **secondary_cluster_nodes_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the node.
	  * Constraints: The maximum value is `16`. The minimum value is `1`.
	* `ip_address` - (String) IP address assigned to the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9.:]+$/`.
	* `memory` - (Integer) Memory allocated to the virtual machine in MB.
	  * Constraints: The maximum value is `1048576`. The minimum value is `1`.
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9.]+$/`.
	* `region` - (String) Region where the virtual machine is deployed.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `vm_id` - (String) Identifier of the Power virtual server instance.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current operational status of the virtual machine.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) Workspace identifier associated with the node.
	  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_location` - (String) Secondary cluster location.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_workspace` - (String) Secondary workspace identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_workspace_crn` - (String) CRN of the secondary workspace.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[A-Za-z0-9._:-]+(:[A-Za-z0-9._:-]+)*$/`.
* `service_description` - (String) Description of provisioned service.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,\\- ]+$/`.
* `service_id` - (String) Identifier for the service.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `service_name` - (String) Name of service.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `standby_region_name` - (String) name of the standby worksapce region.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `standby_workspace_name` - (String) name of the standby worksapce.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `user_tags` - (String) User defined tags.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,-]+$/`.

