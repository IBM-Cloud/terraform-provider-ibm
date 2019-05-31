---
layout: "ibm"
page_title: "IBM : ipsec_policy"
sidebar_current: "docs-ibm-resource-is-ipsec-policy"
description: |-
  Manages IBM ipsec policy.
---

# ibm\_is_ipsec_policy

Provides a ipsec policy resource. This allows ipsec policy to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a ip sec policy:

```hcl
resource "ibm_is_ipsec_policy" "example" {
	name = "test"
	authentication_algorithm = "md5"
	encryption_algorithm = "3des"
	pfs = "disabled"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the IPsec policy.
* `authentication_algorithm` - (Required, string)  The authentication algorithm. Enumeration type: md5, sha1, sha256.
* `encryption_algorithm` - (Required, string) The encryption algorithm. Enumeration type: 3des, aes128, aes256.
* `pfs` - (Required, string) Perfect Forward Secrecy. Enumeration type: disabled, group_2, group_5, group_14.
* `key_lifetime` - (Optional, int) The key lifetime in seconds. Maximum: 86400, Minimum: 300. Default is 3600.
* `resource_group` - (Optional, string) The resource group where the ip sec policy to be created

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Ip Sec Policy.
* `encapsulation_mode` - The encapsulation mode used. Only tunnel is supported.
* `transform_protocol` - The transform protocol used. Only esp is supported.
* `vpn_connections` - Collection of references to VPN connections that use this IPsec policy. Nested connections is
	* `name` - The name given to this VPN connection.
	* `id` -  The unique identifier of a VPN connection.
	* `href` - The VPN connection's canonical URL.

## Import

ibm_is_ipsec_policy can be imported using ID, eg

```
$ terraform import ibm_is_ipsec_policy.example d7bec597-4726-451f-8a63-e62e6f19c32c
```