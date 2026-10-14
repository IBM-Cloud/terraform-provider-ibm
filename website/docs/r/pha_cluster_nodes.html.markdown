---
layout: "ibm"
page_title: "IBM : ibm_pha_cluster_nodes"
description: |-
  Manages pha_cluster_nodes.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_cluster_nodes

Add and delete cluster nodes from PowerHA deployments.

## Example Usage

```hcl
resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
  accept_language = "en-US"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  primary_cluster_nodes = ["8eefautr-4c02-0009-0086-vdvds_cdvcesdvcscv"]
  secondary_cluster_nodes = ["8eefautr-4c02-0009-0086-vdvds_cdvcesdvcscv"]
}
```
## Adding a VM

To add a new VM to the cluster:

Add the VM ID to the primary_cluster_nodes list in your Terraform configuration.

Run:
```hcl

terraform apply
```

Terraform will update the resource and include the new VM in the cluster.

## Removing a VM

To remove a VM from the cluster:

Remove the VM ID from the primary_cluster_nodes list in your Terraform configuration.

Run:

```hcl

terraform apply
```

Terraform will update the resource and remove the VM from the cluster.

## Important Notes
Changes to VM membership are fully controlled by the primary_cluster_nodes field.
Terraform compares the desired state (configuration) with the current state and performs updates accordingly.

Running:

```hcl

terraform destroy
```

will only remove the resource from the Terraform state and delete the managed resource, not selectively remove individual VMs.

### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete. 
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `6` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`. 

## Argument Reference

You can specify the following arguments for this resource.

* `primary_cluster_nodes` - (Optional, List) List of primary cluster node VM IDs.
  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
* `secondary_cluster_nodes` - (Optional, List) List of secondary cluster node VM IDs.
  * Constraints: The maximum length is `100` items. The minimum length is `1` item.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_cluster_nodes.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `primary_node_details` - (List) Details of the primary cluster nodes.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **primary_node_details**:
	* `agent_status` - (String) Status of the PHA agent running on the node.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `cores` - (Float) Number of CPU cores allocated to the VM. 
    * `ip_addresses` - (List) List of IP addresses assigned to the VM.
      * Constraints: The list items must match regular expression `/^(?:\\d{1,3}\\.){3}\\d{1,3}$/.  The maximum length is `16` items. The minimum length is `0` items.
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
      * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

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
      * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.


* `etag` - ETag identifier for pha_cluster_nodes.

## Import

You can import the `ibm_pha_cluster_nodes` resource by using `id`.
The `id` property can be formed from `pha_instance_id`, and `pha_instance_id` in the following format:

<pre>
&lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
* `pha_instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. Unique identifier of the provisioned instance.
* `pha_instance_id`: A string in the format `04044004-123`. Identifier for this cluster node response.

# Syntax
<pre>
$ terraform import ibm_pha_cluster_nodes.pha_cluster_nodes &lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
