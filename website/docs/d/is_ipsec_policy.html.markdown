---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_ipsec_policy"
description: |-
  Get information about IPsecPolicy
---

# ibm_is_ipsec_policy

Provides a read-only data source for IPsecPolicy. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about managing IBM Cloud VPN Gateway and IPsec policy , see [about site-to-site VPN gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#policy-negotiation).

## Example Usage

```hcl
resource "ibm_is_ipsec_policy" "example" {
  name                     = "my-ipsec-policy"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  pfs                      = "disabled"
}

data "ibm_is_ipsec_policy" "example" {
  ipsec_policy = ibm_is_ipsec_policy.example.id
}
```
```hcl
resource "ibm_is_ipsec_policy" "example" {
  name                     = "my-ipsec-policy"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  pfs                      = "disabled"
}

data "ibm_is_ipsec_policy" "example" {
  name = ibm_is_ipsec_policy.example.name
}
```
## Argument Reference

Review the argument reference that you can specify for your data source.

- `ipsec_policy` - (Optional, String) The IPsec policy identifier.
    ~> **NOTE** : One of `ipsec_policy` or  `name` is required

- `name` - (Optional, String) The name of the ipsec policy
    ~> **NOTE** : One of `ipsec_policy` or  `name` is required

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the IPsecPolicy.
- `authentication_algorithm` - (String) The authentication algorithm.
- `connections` - (List) The VPN gateway connections that use this IPsec policy.
  Nested scheme for **connections**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The VPN connection's canonical URL.
	- `id` - (String) The unique identifier for this VPN gateway connection.
	- `name` - (String) The user-defined name for this VPN connection.
	- `resource_type` - (String) The resource type.
- `created_at` - (String) The date and time that this IPsec policy was created.
- `encapsulation_mode` - (String) The encapsulation mode used. Only `tunnel` is supported.
- `encryption_algorithm` - (String) The encryption algorithm.
- `href` - (String) The IPsec policy's canonical URL.
- `key_lifetime` - (Integer) The key lifetime in seconds.
- `name` - (String) The user-defined name for this IPsec policy.
- `pfs` - (String) Perfect Forward Secrecy.
- `resource_group` - (List) The resource group for this IPsec policy.
  Nested scheme for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.
- `resource_type` - (String) The resource type.
- `transform_protocol` - (String) The transform protocol used. Only `esp` is supported.

