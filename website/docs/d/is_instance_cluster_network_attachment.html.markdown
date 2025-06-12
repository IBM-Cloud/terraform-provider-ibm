---
layout: "ibm"
page_title: "IBM : ibm_is_instance_cluster_network_attachment"
description: |-
  Get information about InstanceClusterNetworkAttachment
subcategory: "VPC infrastructure"
---

# ibm_is_instance_cluster_network_attachment

Provides a read-only data source to retrieve information about an InstanceClusterNetworkAttachment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment" {
	instance_id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance.instance_id
	instance_cluster_network_attachment_id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance.instance_cluster_network_attachment_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `instance_id` - (Required, Forces new resource, String) The virtual server instance identifier.
- `instance_cluster_network_attachment_id` - (Required, Forces new resource, String) The instance cluster network attachment identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the InstanceClusterNetworkAttachment.
- `before` - (List) The instance cluster network attachment that is immediately before. If absent, this is thelast instance cluster network attachment.
	Nested schema for **before**:
	- `href` - (String) The URL for this instance cluster network attachment.
	- `id` - (String) The unique identifier for this instance cluster network attachment.
	- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
	- `resource_type` - (String) The resource type.
- `cluster_network_interface` - (List) The cluster network interface for this instance cluster network attachment.
	Nested schema for **cluster_network_interface**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this cluster network interface.
	- `id` - (String) The unique identifier for this cluster network interface.
	- `name` - (String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
	- `primary_ip` - (List) The primary IP for this cluster network interface.
		Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.

			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet reserved IP.
		- `id` - (String) The unique identifier for this cluster network subnet reserved IP.
		- `name` - (String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List)
		Nested schema for **subnet**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet.
		- `id` - (String) The unique identifier for this cluster network subnet.
		- `name` - (String) The name for this cluster network subnet. The name is unique across all cluster network subnets in the cluster network.
		- `resource_type` - (String) The resource type.
- `href` - (String) The URL for this instance cluster network attachment.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
	Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the instance cluster network attachment.
- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
- `resource_type` - (String) The resource type.

