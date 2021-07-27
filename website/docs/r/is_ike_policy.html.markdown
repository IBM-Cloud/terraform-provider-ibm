---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ike_policy"
description: |-
  Manages IBM IKE policy.
---

# ibm_is_ike_policy
Create, update, or cancel an Internet Key Exchange (IKE) policy. IKE is an IPSec (Internet Protocol Security) standard protocol that is used to ensure secure communication over the VPC VPN service. For more information, see [Using VPC with your VPC](https://cloud.ibm.com/docs/vpc-on-classic-network?topic=vpc-on-classic-networkusing-vpn-with-your-vpc).

## Example usage

In the following example, you can create a IKE policy:

```terraform
resource "ibm_is_ike_policy" "example" {
  name                     = "test"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
  dh_group                 = 2
  ike_version              = 1
}

```


## Argument reference
Review the argument references that you can specify for your resource. 

- `authentication_algorithm` - (Required, String) Enter the algorithm that you want to use to authenticate `IPSec` peers. Available options are `md5`, `sha1`, or `sha256`.
- `dh_group`  - (Required, Integer) Enter the Diffie-Hellman group that you want to use for the encryption key. Available enumeration type are `2`, `5`, `14`, or `19`.
- `encryption_algorithm` - (Required, String) Enter the algorithm that you want to use to encrypt data. Available options are: `triple_des`, `aes128`, or `aes256`.
- `ike_version`  - (Optional, Integer) Enter the IKE protocol version that you want to use. Available options are `1`, or `2`.
- `key_lifetime`  - (Optional, Integer)The key lifetime in seconds. `Maximum: 86400`, `Minimum: 1800`. Default is `28800`. 
- `name` - (Required, String) Enter a name for your IKE policy.
- `resource_group` - (Optional, Forces new resource, String) Enter the ID of the resource group where you want to create the IKE policy. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the IKE policy is created in the `default` resource group.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `href`-  (String) The canonical URL that was assigned to your IKE policy. 
- `id` - (String) The unique identifier of the IKE policy that you created.
- `negotiation_mode` - (String) The IKE negotiation mode that was set for your IKE policy. Only `main` is supported. 
- `vpn_connections`- List - A collection of VPN connections that use the IKE policy.

  Nested scheme for `vpn_connections`:
  - `name`-String - The name given to the VPN connection.
  - `id`-  (String) The unique identifier of a VPN connection.
	- `href` - (String) The VPN connection's canonical URL.

## Import
The `ibm_is_ike_policy` resource can be imported by using IKE Policy ID.

**Example**

```
$ terraform import ibm_is_ike_policy.example d7bec597-4726-451f-8a63-e62e6f19c32c
```

