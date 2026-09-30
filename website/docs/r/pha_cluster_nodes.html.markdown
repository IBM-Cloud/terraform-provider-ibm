---
layout: "ibm"
page_title: "IBM : ibm_pha_cluster_nodes"
description: |-
  Manages pha_cluster_nodes.
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_cluster_nodes

Use this resource to create and delete PowerHA cluster nodes (pha_cluster_nodes). It supports managing cluster nodes by adding new nodes and removing nodes when no longer required.

## Example Usage

```hcl
resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
  accept_language = "en-US"
  instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  primary_cluster_nodes = ["sdthautr-6762-028t-0975-8bd4sdfrb6"]
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

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) The language requested for the return document. (ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, String) Unique identifier of the provisioned instance.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.
* `primary_cluster_nodes` - (Required, List of String) List of primary cluster node VM IDs.
  * Constraints: Minimum items are 1. Maximum items allowed are 8. Each value must match /^[A-Za-z0-9._:-]+$/. Length between 1 and 36 characters
* `secondary_cluster_nodes` - (Optional, List of String) List of secondary cluster node VM IDs.
  * Constraints: Minimum items are 1. Maximum items allowed are 8. Each value must match /^[A-Za-z0-9._:-]+$/. Length between 1 and 36 characters


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_cluster_nodes.
* `instance_id` - (String) Identifier for this cluster node response.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_-]+$/`.
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

* `etag` - ETag identifier for pha_cluster_nodes.

## Import

You can import the `ibm_pha_cluster_nodes` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_id` in the following format:

<pre>
&lt;instance_id&gt;
</pre>
* `instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. Unique identifier of the provisioned instance.

# Syntax
<pre>
$ terraform import ibm_pha_cluster_nodes.pha_cluster_nodes &lt;instance_id&gt;
</pre>
