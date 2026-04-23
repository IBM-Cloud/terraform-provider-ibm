---
layout: "ibm"
page_title: "IBM : ibm_pha_deployment"
description: |-
  Get information about pha_deployment
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_deployment

Provides a read-only data source to retrieve information about a pha_deployment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_deployment" "pha_deployment" {
	if_none_match = ibm_pha_deployment.pha_deployment_instance.if_none_match
	instance_id = ibm_pha_deployment.pha_deployment_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier of the provisioned instance.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_deployment.
* `cloud_account_id` - (String) Cloud account identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `connectivity_type` - (String) Type of network connectivity.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `creation_time` - (String) Timestamp expressing creation time.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `custom_network` - (List) List of custom network CIDRs.
  * Constraints: The list items must match regular expression `/^[A-Za-z0-9._:\/-]+$/`. The maximum length is `16` items. The minimum length is `0` items.
* `deprovision_time` - (String) Timestamp expressing deprovision time.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$/`.
* `guid` - (String) Global unique identifier.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `is_duplicate` - (Boolean) Indicates whether deployment is duplicate.
* `plan_id` - (String) Identifier for the service plan.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `plan_name` - (String) Name of service plan.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `powerha_cluster_name` - (String) Name of the PowerHA cluster.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_cluster_type` - (String) Type of PowerHA cluster.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_level` - (String) PowerHA version level.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_cluster_nodes_details` - (List) List of primary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **primary_cluster_nodes_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the node.
	* `ip_address` - (String) IP address assigned to the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.:]+$/`.
	* `memory` - (Integer) Memory allocated to the virtual machine in MB or GB.
	  * Constraints: The maximum value is `64`. The minimum value is `1`.
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.]+$/`.
	* `region` - (String) Region where the virtual machine is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current operational status of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) Workspace identifier associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_location` - (String) Primary cluster location.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_region_name` - (String) name of the primary workspace region.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace` - (String) Primary workspace identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace_name` - (String) name of the primary workspace.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_end_time` - (String) Time stamp provisioning completed.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_start_time` - (String) Time stamp provisioning started.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_status` - (String) Current provision status.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `region_id` - (String) Deployment region identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group` - (String) Name of the resource group.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group_crn` - (String) CRN of associated resource group.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:\/-]+$/`.
* `resource_instance` - (String) Resource instance identifier.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_cluster_nodes` - (List) List of secondary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **secondary_cluster_nodes**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the node.
	* `ip_address` - (String) IP address assigned to the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.:]+$/`.
	* `memory` - (Integer) Memory allocated to the virtual machine in MB or GB.
	  * Constraints: The maximum value is `64`. The minimum value is `1`.
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.]+$/`.
	* `region` - (String) Region where the virtual machine is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current operational status of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) Workspace identifier associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_location` - (String) Secondary cluster location.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_workspace` - (String) Secondary workspace identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `service_description` - (String) Description of provisioned service.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,\\- ]+$/`.
* `service_id` - (String) Identifier for the service.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `service_name` - (String) Name of service.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `standby_region_name` - (String) name of the standby workspace region.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `standby_workspace_name` - (String) name of the standby workspace.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `user_tags` - (String) User defined tags.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,-]+$/`.

