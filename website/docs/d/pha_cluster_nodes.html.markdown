---
layout: "ibm"
page_title: "IBM : ibm_pha_cluster_nodes"
description: |-
  Get information about pha_cluster_nodes
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_cluster_nodes

Provides a read-only data source to retrieve information about pha_cluster_nodes. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_cluster_nodes" "pha_cluster_nodes" {
	if_none_match = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.if_none_match
	instance_id = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.instance_id
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

* `id` - The unique identifier of the pha_cluster_nodes.
* `primary_node_details` - (List) Details of the primary cluster nodes.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **primary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `16` items. The minimum length is `0` items.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_node_details` - (List) Details of the secondary cluster nodes.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **secondary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `16` items. The minimum length is `0` items.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

