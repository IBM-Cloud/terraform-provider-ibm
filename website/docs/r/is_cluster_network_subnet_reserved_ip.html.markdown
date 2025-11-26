---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_subnet_reserved_ip"
description: |-
  Manages ClusterNetworkSubnetReservedIP.
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_subnet_reserved_ip

Create, update, and delete ClusterNetworkSubnetReservedIPs with this resource. [About cluster networks](https://cloud.ibm.com/docs/vpc?topic=vpc-about-cluster-network)

## Example Usage

```hcl
resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
  address = "192.168.3.4"
  cluster_network_id = "cluster_network_id"
  cluster_network_subnet_id = "cluster_network_subnet_id"
  name = "my-cluster-network-subnet-reserved-ip"
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `address` - (Optional, String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
- `cluster_network_id` - (Required, Forces new resource, String) The cluster network identifier.
- `cluster_network_subnet_id` - (Required, Forces new resource, String) The cluster network subnet identifier.
- `name` - (Optional, String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the ClusterNetworkSubnetReservedIP.
- `auto_delete` - (Boolean) Indicates whether this cluster network subnet reserved IP member will be automatically deleted when either `target` is deleted, or the cluster network subnet reserved IP is unbound.
- `created_at` - (String) The date and time that the cluster network subnet reserved IP was created.
- `href` - (String) The URL for this cluster network subnet reserved IP.
- `cluster_network_subnet_reserved_ip_id` - (String) The unique identifier for this cluster network subnet reserved IP.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
	Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `internal_error`, `resource_suspended_by_provider`.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the cluster network subnet reserved IP. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `owner` - (String) The owner of the cluster network subnet reserved IPThe enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `provider`, `user`. 
- `resource_type` - (String) The resource type. Allowable values are: `cluster_network_subnet_reserved_ip`.
- `target` - (List) The target this cluster network subnet reserved IP is bound to.If absent, this cluster network subnet reserved IP is provider-owned or unbound.
	Nested schema for **target**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this cluster network interface.
	- `id` - (String) The unique identifier for this cluster network interface.
	- `name` - (String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
	- `resource_type` - (String) The resource type. Allowable values are: `cluster_network_interface`. 
- `etag` - ETag identifier for ClusterNetworkSubnetReservedIP.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_cluster_network_subnet_reserved_ip` resource by using `id`.
The `id` property can be formed from `cluster_network_id`, `cluster_network_subnet_id`, and `cluster_network_subnet_reserved_ip_id`. For example:

```terraform
import {
  to = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip
  id = "<cluster_network_id>/<cluster_network_subnet_id>/<cluster_network_subnet_reserved_ip_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip <cluster_network_id>/<cluster_network_subnet_id>/<cluster_network_subnet_reserved_ip_id>
```