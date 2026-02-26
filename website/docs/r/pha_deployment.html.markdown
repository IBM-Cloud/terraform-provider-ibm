---
layout: "ibm"
page_title: "IBM : ibm_pha_deployment"
description: |-
  Manages pha_deployment.
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_deployment

Create, update, and delete pha_deployments with this resource.

## Example Usage

```hcl
resource "ibm_pha_deployment" "pha_deployment_instance" {
  accept_language = "en-US"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  primary_cluster_nodes {
		agent_status = "RUNNING"
		cores = 8.0
		ip_address = "10.0.2.45"
		memory = 32
		pha_level = "7.2.1"
		region = "us-south"
		vm_id = "vm-3c91af27"
		vm_name = "pha-node-01"
		vm_status = "ACTIVE"
		workspace_id = "workspace-pha-prod"
  }
  primary_location = "us-south"
  primary_workspace = "workspace-primary"
  secondary_location = "us-east"
  secondary_workspace = "workspace-secondary"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, Forces new resource, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.
* `primary_cluster_nodes` - (Optional, Forces new resource, List) List of primary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **primary_cluster_nodes**:
	* `agent_status` - (Optional, String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Optional, Float) Number of CPU cores allocated to the node.
	* `ip_address` - (Optional, String) IP address assigned to the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.:]+$/`.
	* `memory` - (Optional, Integer) Memory allocated to the virtual machine in MB or GB.
	  * Constraints: The maximum value is `64`. The minimum value is `1`.
	* `pha_level` - (Optional, String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9.]+$/`.
	* `region` - (Optional, String) Region where the virtual machine is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (Optional, String) Unique identifier of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (Optional, String) Name of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (Optional, String) Current operational status of the virtual machine.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (Optional, String) Workspace identifier associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_location` - (Optional, Forces new resource, String) Primary cluster location.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace` - (Required, Forces new resource, String) Primary workspace identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_location` - (Optional, Forces new resource, String) Secondary cluster location.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_workspace` - (Optional, Forces new resource, String) Secondary workspace identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_deployment.
* `api_key` - (String) API key used for authentication to the deployment service.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:\\-*]+$/`.
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
* `pha_instance_id` - (String) Provision request identifier.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
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
* `primary_region_name` - (String) name of the primary workspace region.
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

* `etag` - ETag identifier for pha_deployment.

## Import

You can import the `ibm_pha_deployment` resource by using `id`.
The `id` property can be formed from `pha_instance_id`, and `pha_instance_id` in the following format:

<pre>
&lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
* `pha_instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. instance id of instance to provision.
* `pha_instance_id`: A string in the format `prov-9f8a7b6c`. Provision request identifier.

# Syntax
<pre>
$ terraform import ibm_pha_deployment.pha_deployment &lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
