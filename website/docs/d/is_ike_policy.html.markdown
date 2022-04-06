---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_ike_policy"
description: |-
  Get information about IKEPolicy
---

# ibm_is_ike_policy

Provides a read-only data source for IKEPolicy. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about managing IBM Cloud VPN Gateway and IKE policy , see [about site-to-site VPN gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#policy-negotiation).

## Example Usage

```hcl
resource "ibm_is_ike_policy" "example" {
  name                     = "my-ike-policy"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  dh_group                 = 2
  ike_version              = 1
}

data "ibm_is_ike_policy" "example" {
  ike_policy = ibm_is_ike_policy.example.id
}
```

```hcl
resource "ibm_is_ike_policy" "example" {
  name                     = "my-ike-policy"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  dh_group                 = 2
  ike_version              = 1
}

data "ibm_is_ike_policy" "example" {
  name = ibm_is_ike_policy.example.name
}
```
## Argument Reference

Review the argument reference that you can specify for your data source.

- `ike_policy` - (Optional, String) The IKE policy identifier.
    ~> **NOTE** : One of `ike_policy` or  `name` is required

- `name` - (Optional, String) The name of the IKE policy
    ~> **NOTE** : One of `ike_policy` or  `name` is required

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the IKEPolicy.
- `authentication_algorithm` - (String) The authentication algorithm. 
- `connections` - (List) The VPN gateway connections that use this IKE policy.
  Nested scheme for **connections**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The VPN connection's canonical URL.
	- `id` - (String) The unique identifier for this VPN gateway connection.
	- `name` - (String) The user-defined name for this VPN connection.
	- `resource_type` - (String) The resource type.
- `created_at` - (String) The date and time that this IKE policy was created.
- `dh_group` - (Integer) The Diffie-Hellman group.
- `encryption_algorithm` - (String) The encryption algorithm.
- `href` - (String) The IKE policy's canonical URL.
- `ike_version` - (Integer) The IKE protocol version.
- `key_lifetime` - (Integer) The key lifetime in seconds.
- `name` - (String) The user-defined name for this IKE policy.
- `negotiation_mode` - (String) The IKE negotiation mode. Only `main` is supported.
- `resource_group` - (List) The resource group for this IKE policy.
  Nested scheme for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.
- `resource_type` - (String) The resource type.

