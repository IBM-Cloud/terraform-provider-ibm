---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_network_acl"
description: |-
  Manages IBM network ACL.
---

# ibm_is_network_acl
Create, update, or delete a network access control list (ACL). For more information, about network ACL, see [setting up network ACLs](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_network_acl" "example" {
  name = "example-acl"
  vpc  = ibm_is_vpc.example.id
  rules {
    name        = "outbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    # Deprecated block: replaced with 'protocol', 'code', and 'type' arguments
    # icmp {
    #   code = 1
    #   type = 1
    # }
    protocol = "icmp"
    code     = 1
    type     = 1
  }
  rules {
    name        = "inbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    # Deprecated block: replaced with 'protocol', 'code', and 'type' arguments
    # icmp {
    #   code = 1
    #   type = 1
    # }
    protocol = "icmp"
    code     = 1
    type     = 1
    }
}
```

## Argument reference
Review the argument references that you can specify for your resource. 
 
- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the network acl.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `name` - (Optional, String) The name of the network ACL. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the network ACL.
- `rules`- (Optional, Array of Strings) A list of rules for a network ACL. The order in which the rules are added to the list determines the priority of the rules. For example, the first rule that you want to enforce must be specified as the first rule in this list.

  Nested scheme for `rules`:
  - `name` - (Optional, String) The user-defined name for this rule.
  - `action` - (Required, String)  `Allow` or `deny` matching network traffic.
  - `source` - (Required, String) The source IP address or CIDR block.
  - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
  - `destination` - (Required, String) The destination IP address or CIDR block.
  - `direction` - (Required, String) Indicates whether the traffic to be matched is `inbound` or `outbound`.
  - `icmp`- (Optional, DEPRECATED, List) The protocol ICMP. `icmp` is deprecated and use `protocol`, `code`, and `type` argument instead.

    Nested scheme for `icmp`:
    - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
    - `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
  - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
  - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.  
  - `protocol` - (Optional, String) The name of the network protocol.  
  - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
  - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
  - `tcp`- (Optional, DEPRECATED, List) TCP protocol. `tcp` is deprecated and use `protocol`, `port_min`, `port_max`, `source_port_max` and `source_port_min` argument instead.

    Nested scheme for `tcp`:
    - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched, if unspecified, 1 is used as default.
    - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used as default.
    - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used as default.
  - `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
  - `udp` - (Optional, DEPRECATED, List) UDP protocol. `udp` is deprecated and use `protocol`, `port_min`, `port_max`,  `source_port_max` and `source_port_min` argument instead.

    Nested scheme for `udp`:
    - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
    - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
- `tags`- (Optional, List of Strings) Tags associated with the network ACL.
- `vpc` - (Optional, Forces new resource, String) The VPC ID. This parameter is required if you want to create a network ACL for a Generation 2 VPC.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the network ACL.
- `id` - (String) The ID of the network ACL.
- `rules`- (List) The rules for a network ACL.

  Nested scheme for `rules`:
  - `id` - (String) The rule ID.
  - `ip_version` - (String) The IP version of the rule.
  - `subnets` - (String) The subnets for the ACL rule.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_network_acl` resource by using `id`.
The `id` property can be formed from `network ACL ID`. For example:

```terraform
import {
  to = ibm_is_network_acl.example
  id = "<network_acl_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_network_acl.example <network_acl_id>
```