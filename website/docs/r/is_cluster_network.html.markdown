---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network"
description: |-
  Manages ClusterNetwork.
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network

Create, update, and delete ClusterNetworks with this resource. [About cluster networks](https://cloud.ibm.com/docs/vpc?topic=vpc-about-cluster-network)

## Example Usage

```hcl
resource "ibm_is_cluster_network" "example" {
  name = "my-cluster-network"
  profile = "h100"
  resource_group = "fee82deba12e4c0fb69c3b09d1f12345"
  subnet_prefixes {
		cidr = "10.0.0.0/24"
  }
  vpc {
		id = "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
  }
  zone  = "us-east-3"
}

```

## Argument Reference

You can specify the following arguments for this resource.

- `name` - (Optional, String) The name for this cluster network. The name must not be used by another cluster network in the region.
- `profile` - (Required, String) The profile (globally unique name for the cluster network profile) for this cluster network.
- `resource_group` - (Optional, String) The resource group (unique identifier for the resource group) for this cluster network.
- `subnet_prefixes` - (Optional, List) The IP address ranges available for subnets for this cluster network.(The maximum length is `1` item. The minimum length is `1` item.)
	Nested schema for **subnet_prefixes**:
	- `cidr` - (Required, String) The CIDR block for this prefix.
- `vpc` - (Required, List) The VPC this cluster network resides in.
	Nested schema for **vpc**:
	- `id` - (Required, String) The unique identifier for this VPC.
- `zone` - (Required, List)  The zone (globally unique name for this zone) this cluster network resides in.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the ClusterNetwork.
- `created_at` - (String) The date and time that the cluster network was created.
- `crn` - (String) The CRN for this cluster network.
- `href` - (String) The URL for this cluster network.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
	Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `internal_error`, `resource_suspended_by_provider`.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the cluster network. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `resource_type` - (String) The resource type. Allowable values are: `cluster_network`.
- `etag` - ETag identifier for ClusterNetwork.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_cluster_network` resource by using `id`.
The `id` property can be formed using the cluster_network id. For example:

```terraform
import {
  to = ibm_is_cluster_network.is_cluster_network
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_cluster_network.is_cluster_network <id>
```