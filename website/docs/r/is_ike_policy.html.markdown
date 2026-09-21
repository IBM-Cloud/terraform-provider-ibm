---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ike_policy"
description: |-
  Manages IBM IKE policy.
---

# ibm_is_ike_policy
Create, update, or cancel an Internet Key Exchange (IKE) policy. IKE is an IPSec (Internet Protocol Security) standard protocol that is used to ensure secure communication over the VPC VPN service. For more information, see [Using VPC with your VPC](https://cloud.ibm.com/docs/vpc-on-classic-network?topic=vpc-on-classic-networkusing-vpn-with-your-vpc).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Multiple algorithms

```terraform
resource "ibm_is_ike_policy" "example_multi" {
  name                      = "example-ike-policy-multi"
  authentication_algorithms = ["sha384", "sha512"]
  encryption_algorithms     = ["aes256", "aes192"]
  dh_groups                 = [15, 14]
  ike_version               = 2
  key_lifetime              = 1800
}
```

### Deprecated single-algorithm

```terraform
resource "ibm_is_ike_policy" "example" {
  name                     = "example-ike-policy"
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes128"
  dh_group                 = 14
  ike_version              = 1
}
```



## Argument reference
Review the argument references that you can specify for your resource.

- `authentication_algorithm` - (Optional, String) **Deprecated.** Use `authentication_algorithms` instead. The authentication algorithm. Available options are `sha256`, `sha512`, `sha384`. Existing configurations using this field continue to work without any changes. Exactly one of `authentication_algorithm` or `authentication_algorithms` must be specified.
- `authentication_algorithms` - (Optional, List) The authentication algorithms to use for IKE negotiation. The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it. Exactly one of `authentication_algorithm` or `authentication_algorithms` must be specified.
  * Constraints: Allowable list items are: `sha256`, `sha384`, `sha512`. The maximum length is `3` items. The minimum length is `1` item.
- `dh_group`  - (Optional, Integer) **Deprecated.** Use `dh_groups` instead. The Diffie-Hellman group. Available enumeration type are `14`, `19`, `15`, `16` ,`17` ,`18` ,`20` ,`21` ,`22` ,`23` ,`24` ,`31`. Exactly one of `dh_group` or `dh_groups` must be specified.
- `dh_groups` - (Optional, List) The Diffie-Hellman groups to use for IKE negotiation. The order of the Diffie-Hellman groups in this array indicates their priority for negotiation, with each Diffie-Hellman group having priority over the one after it. Exactly one of `dh_group` or `dh_groups` must be specified.
  * Constraints: Allowable list items are: `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `31`. The maximum length is `12` items. The minimum length is `1` item.
- `encryption_algorithm` - (Optional, String) **Deprecated.** Use `encryption_algorithms` instead. The encryption algorithm. Available options are: `aes128`, `aes192`, `aes256`. Exactly one of `encryption_algorithm` or `encryption_algorithms` must be specified.
- `encryption_algorithms` - (Optional, List) The encryption algorithms to use for IKE negotiation. The order of the algorithms in this array indicates their priority for negotiation, with each algorithm having priority over the one after it. Exactly one of `encryption_algorithm` or `encryption_algorithms` must be specified.
  * Constraints: Allowable list items are: `aes128`, `aes192`, `aes256`. The maximum length is `3` items. The minimum length is `1` item.
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

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_ike_policy` resource by using `id`.
The `id` property can be formed from `IKE Policy ID`. For example:

```terraform
import {
  to = ibm_is_ike_policy.example
  id = "<ike_policy_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_ike_policy.example <ike_policy_id>
```