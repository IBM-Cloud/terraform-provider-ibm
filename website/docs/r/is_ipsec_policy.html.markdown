---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ipsec_policy"
description: |-
  Manages IBM IPsec policy.
---

# ibm_is_ipsec_policy
Create, update, or delete an ipsec policy resource. For more information, about ipsec policy, see [creating an IPsec policy](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-ipsec-policy).


## Example usage
In the following example, you can create a IPsec policy:

```terraform
resource "ibm_is_ipsec_policy" "example" {
  name                     = "test"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  pfs                      = "disabled"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `authentication_algorithm` - (Required, String) Enter the algorithm that you want to use to authenticate `IPSec` peers. Available options are `md5`, `sha1`, or `sha256`.
- `encryption_algorithm` - (Required, String) Enter the algorithm that you want to use to encrypt data. Available options are: `triple_des`, `aes128`, or `aes256`. No.
- `key_lifetime`  - (Optional, Integer) Enter the time in seconds that your encryption key can be used before it expires. You must enter a number between 300 and 86400. If you do not specify this option, 3600 seconds is used.
- `name` - (Required, String) Enter the name for your IPSec policy.
- `pfs` - (Required, String) Enter the Perfect Forward Secrecy protocol that you want to use during a session. Available options are `disabled`, `group_2`, `group_5`, and `group_14`.
- `resource_group` - (Optional, Forces new resource, String) Enter the ID of the resource group where you want to create the IPSec policy. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the IPSec policy is created in the `default` resource group. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `encapsulation_mode` - (String) The encapsulation mode that was set for your IPSec policy. Only `tunnel` is supported.
- `id` - (String) The unique identifier of the IPSec policy that you created.
- `transform_protocol` - (String) The transform protocol that is used in your IPSec policy. Only the `esp` protocol is supported that uses the triple DES (3DES) encryption algorithm to encrypt your data.
- `vpn_connections`- (List) A collection of VPN connections that use the IPSec policy. 

  Nested scheme for `vpn_connections`:
  - `href` - (String) The VPN connection's canonical URL.
  - `id` -  (String) The unique identifier of a VPN connection.
  - `name` - (String) The name given to this VPN connection.

## Import

The `ibm_is_ipsec_policy` resource can be imported by using ID.

**Example**

```
$ terraform import ibm_is_ipsec_policy.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
