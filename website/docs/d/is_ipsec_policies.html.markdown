---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_ipsec_policies"
description: |-
  Get information about IPsecPolicyCollection
---

# ibm_is_ipsec_policies

Provides a read-only data source for IPsecPolicyCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about managing IBM Cloud VPN Gateway and IPsec policy , see [about site-to-site VPN gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#policy-negotiation).

## Example Usage

```hcl
data "ibm_is_ipsec_policies" "example" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `ipsec_policies` - (List) Collection of IPsec policies.
  Nested scheme for **ipsec_policies**:
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
	- `id` - (String) The unique identifier for this IPsec policy.
	- `key_lifetime` - (Integer) The key lifetime in seconds.
	- `name` - (String) The user-defined name for this IPsec policy.
	- `pfs` - (String) Perfect Forward Secrecy.
	- `resource_group` - (List) The resource group object, for this IPsec policy.
	  Nested scheme for **resource_group**:
		- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The user-defined name for this resource group.
	- `resource_type` - (String) The resource type.
	- `transform_protocol` - (String) The transform protocol used. Only `esp` is supported.
