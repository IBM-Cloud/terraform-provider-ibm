---
layout: "ibm"
page_title: "IBM : ibm_pha_cluster_nodes"
description: |-
  Get information about pha_cluster_nodes
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_cluster_nodes

Retrieves the list of all cluster nodes and their details for the specified PowerHA service instance.

## Example Usage

```hcl
data "ibm_pha_cluster_nodes" "pha_cluster_nodes" {
	if_none_match = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.if_none_match
	instance_id = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.instance_id
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

* `id` - (String) Identifier for this cluster node response.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_-]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `primary_node_details` - (List) Details of the primary cluster nodes.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **primary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `16` items. The minimum length is `0` items. The maximum length of each item is `1048` characters. The minimum length of each item is `7` characters.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `powerha_version_supported` - (Boolean) Indicates whether the installed PowerHA version is supported.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `powerha_version_supported` - (Boolean) Indicates whether the installed PowerHA version is supported.  

* `secondary_node_details` - (List) Details of the secondary cluster nodes.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **secondary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM.
	* `ip_addresses` - (List) List of IP addresses assigned to the VM.
	  * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/`. The maximum length is `16` items. The minimum length is `0` items. The maximum length of each item is `1048` characters. The minimum length of each item is `7` characters.
	* `memory` - (Float) Amount of memory allocated to the VM (in GB).
	* `pha_level` - (String) PowerHA version level installed on the node.
	  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `powerha_version_supported` - (Boolean) Indicates whether the installed PowerHA version is supported.
	* `region` - (String) Region where the VM is deployed.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `vm_id` - (String) Unique identifier of the VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_name` - (String) Name of the VM.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `vm_status` - (String) Current status of the VM.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `workspace_id` - (String) ID of the workspace associated with the VM.
	  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `powerha_version_supported` - (Boolean) Indicates whether the installed PowerHA version is supported.    

