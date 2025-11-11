---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_interface"
description: |-
  Manages ClusterNetworkInterface.
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_interface

Create, update, and delete ClusterNetworkInterfaces with this resource. [About cluster networks](https://cloud.ibm.com/docs/vpc?topic=vpc-about-cluster-network)

## Example Usage

```hcl
resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
  cluster_network_id = "cluster_network_id"
  name = "my-cluster-network-interface"
  primary_ip {
		address = "10.1.0.6"
		name = "my-cluster-network-subnet-reserved-ip"
  }
  subnet {
		id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `cluster_network_id` - (Required, Forces new resource, String) The cluster network identifier.
- `name` - (Optional, String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
- `primary_ip` - (Optional, List) The cluster network subnet reserved IP for this cluster network interface.
	
	Nested schema for **primary_ip**:
	- `address` - (Required, String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	- `href` - (Required, String) The URL for this cluster network subnet reserved IP.
	- `id` - (Required, String) The unique identifier for this cluster network subnet reserved IP.
	- `name` - (Required, String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
	- `resource_type` - (Computed, String) The resource type. Allowable values are: `cluster_network_subnet_reserved_ip`. 
- `subnet` - (Optional, List) The associated cluster network subnet. Required if `primary_ip` does not specify a clusternetwork subnet reserved IP identity.
	
	Nested schema for **subnet**:
	- `href` - (Required, String) The URL for this cluster network subnet.
	- `id` - (Required, String) The unique identifier for this cluster network subnet.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the ClusterNetworkInterface.
- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this cluster network interface. If `false`, source IP spoofing is prevented on this cluster network interface. If `true`, source IP spoofing is allowed on this cluster network interface.
- `auto_delete` - (Boolean) Indicates whether this cluster network interface will be automatically deleted when `target` is deleted.
- `created_at` - (String) The date and time that the cluster network interface was created.
- `enable_infrastructure_nat` - (Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.
- `href` - (String) The URL for this cluster network interface.
- `cluster_network_interface_id` - (String) The unique identifier for this cluster network interface.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
	Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `internal_error`, `resource_suspended_by_provider`. 
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the cluster network interface. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. 
- `mac_address` - (String) The MAC address of the cluster network interface. May be absent if`lifecycle_state` is `pending`.
- `resource_type` - (String) The resource type. llowable values are: `cluster_network_interface`.
- `target` - (List) The target of this cluster network interface.If absent, this cluster network interface is not attached to a target.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	Nested schema for **target**:
	- `href` - (String) The URL for this instance cluster network attachment.
	- `id` - (String) The unique identifier for this instance cluster network attachment.
	- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
	- `resource_type` - (String) The resource type. Allowable values are: `instance_cluster_network_attachment`.
- `vpc` - (List) The VPC this cluster network interface resides in.
	Nested schema for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	- `resource_type` - (String) The resource type. Allowable values are: `vpc`.
- `zone` - (List) The zone this cluster network interface resides in.
	Nested schema for **zone**:
	- `href` - (String) The URL for this zone.
	- `name` - (String) The globally unique name for this zone.

- `etag` - ETag identifier for ClusterNetworkInterface.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_cluster_network_interface` resource by using `id`.
The `id` property can be formed from `cluster_network_id` and `cluster_network_interface_id`. For example:

```terraform
import {
  to = ibm_is_cluster_network_interface.is_cluster_network_interface
  id = "<cluster_network_id>/<cluster_network_interface_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_cluster_network_interface.is_cluster_network_interface <cluster_network_id>/<cluster_network_interface_id>
```