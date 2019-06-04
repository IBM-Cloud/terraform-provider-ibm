---
layout: "ibm"
page_title: "IBM : ike_policy"
sidebar_current: "docs-ibm-resource-is-ike-policy"
description: |-
  Manages IBM ike policy.
---

# ibm\_is_ike_policy

Provides a ike policy resource. This allows ike policy to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a ike policy:

```hcl
resource "ibm_is_ike_policy" "example" {
	name = "test"
	authentication_algorithm = "md5"
	encryption_algorithm = "3des"
	dh_group = 2
	ike_version = 1
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the IKE policy.
* `authentication_algorithm` - (Required, string)  The authentication algorithm. Enumeration type: md5, sha1, sha256.
* `encryption_algorithm` - (Required, string) The encryption algorithm. Enumeration type: 3des, aes128, aes256.
* `dh_group` - (Required, int) The Diffie-Hellman group. Enumeration type: 2, 5, 14.
* `ike_version` - (Optional,int) The IKE protocol version. Enumeration type: 1, 2.
* `key_lifetime` - (Optional, int) The key lifetime in seconds. Maximum: 86400, Minimum: 300. Default is 28800.
* `resource_group` - (Optional, string) The resource group where the ike policy to be created.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the IKE Policy.
* `negotiation_mode` - The IKE negotiation mode. Only main is supported.
* `href` - The IKE policy's canonical URL.
* `vpn_connections` - Collection of references to VPN connections that use this IKE policy. Nested connections is
	* `name` - The name given to this VPN connection.
	* `id` -  The unique identifier of a VPN connection.
	* `href` - The VPN connection's canonical URL.