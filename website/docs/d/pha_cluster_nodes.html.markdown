---
layout: "ibm"
page_title: "IBM : ibm_pha_cluster_nodes"
description: |-
  Get information about pha_cluster_nodes
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_cluster_nodes

Retrieves the list of all cluster nodes and their details for the specified PowerHA service instance.

## Example Usage

```hcl
data "ibm_pha_cluster_nodes" "pha_cluster_nodes" {
	instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document. (ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, Forces new resource, String) Unique identifier of the provisioned instance.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_cluster_nodes.
* `primary_node_details` - (List) Details of the primary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **primary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `48` items. The minimum length is `0` items.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_node_details` - (List) Details of the secondary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **secondary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `48` items. The minimum length is `0` items.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

