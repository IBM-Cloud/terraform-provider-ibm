---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : network_acl_rule"
description: |-
  Manages IBM Network ACL rule.
---

# ibm_is_network_acl_rule

Provides a network ACL rule resource with `icmp`, `tcp`, `udp`, `icmp_tcp_udp`. Protocol `all` in older versions is replaced with `icmp_tcp_udp` from `1.87.0-beta1`.
This allows Network ACL rule to create, update, and delete an existing network ACL. For more information, about managing IBM Cloud Network ACL , see [about network acl](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage (icmp_tcp_udp)

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_network_acl" "example" {
  name = "example-acl"
  vpc  = ibm_is_vpc.example.id
}
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "outbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
  protocol    = "icmp_tcp_udp"
}
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
  name        = "inbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"
  protocol    = "icmp_tcp_udp"
}
```

## Example usage (icmp)

```terraform
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
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
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
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

```

## Example usage (tcp/udp)

```terraform
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "outbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
  # Deprecated block: replaced with 'protocol', 'port_min', 'port_max', `source_port_max` and `source_port_min` arguments
  # tcp {
  #   port_min = 65535
  #   port_max = 1
  #   source_port_max = 60000
  #   source_port_min = 22
  # }
  protocol        = "tcp"
  port_max        = 65535
  port_min        = 1
  source_port_max = 60000
  source_port_min = 22
}
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
  name        = "inbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"
  # Deprecated block: replaced with 'protocol', 'port_min', 'port_max', `source_port_max` and `source_port_min` arguments
  # tcp {
  #   port_min = 65535
  #   port_max = 1
  #   source_port_max = 60000
  #   source_port_min = 22
  # }
  protocol        = "tcp"
  port_max        = 65535
  port_min        = 1
  source_port_max = 60000
  source_port_min = 22
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `action` - (Required, String) Whether to **allow** or **deny** matching traffic. Provide `protocol` mandatory if actions is `deny`, otherwise `any` protocol gets denied which may cause discrepancies in older versions of provider. 
- `before` - (Optional, String) The unique identifier of the rule that this rule is immediately before. If unspecified, this rule will be inserted after all existing rules. While modifying the resource, specify **"null"** (within double quotes) to move this rule after all existing rules.

~> **NOTE:** When using the `before` attribute to specify rule ordering:</br>
    1. Adding a new rule with the `before` attribute will change the position of that rule in the ACL rule list, which may affect the evaluation order of other rules.</br>
    2. Updating the `before` attribute of an existing rule will reposition that rule, potentially causing changes to other rules' relative positions in the evaluation sequence.</br>
    3. Setting `before = "null"` will move the rule to the end of the ACL rule list.</br>
    These position changes are expected and reflect the actual state of your network ACL ruleset, however, they may cause Terraform to show additional changes in other rules during subsequent plan/apply operations.
- `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
- `destination` - (Required, String) The destination IP address or CIDR block.
- `direction` - (Required, String) Whether the traffic to be matched is **inbound** or **outbound**.
- `icmp` - (Optional, DEPRECATED, List) The protocol ICMP. `icmp` is deprecated and use `protocol`, `code`, and `type` argument instead.

   Nested scheme for `icmp`:
   - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
   - `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
- `network_acl` - (Required, String) The ID of the network ACL.
- `name` - (Optional, String) The user-defined name for this rule.
- `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
- `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
- `protocol` - (Optional, String) The name of the network protocol.
- `source` - (Required, String) The source IP address or CIDR block.
- `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
- `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
- `tcp` - (Optional, DEPRECATED, List) TCP protocol. `tcp` is deprecated and use `protocol`, `port_min`, `port_max`, `source_port_max` and `source_port_min` argument instead.

   Nested scheme for `tcp`:
   - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
   - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
- `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
- `udp` - (Optional, DEPRECATED, List) UDP protocol. `udp` is deprecated and use `protocol`, `port_min`, `port_max`,  `source_port_max` and `source_port_min` argument instead.

   Nested scheme for `udp`:
   - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
   - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.

~> **NOTE:** Only one type of protocol out of **icmp**, **tcp**, or **udp** can be used to create a new rule. If none is provided, **all** is selected.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the network ACL rule. The ID is composed of `\network_acl\rule_id`.
- `href` - (String) The URL for this network ACL rule.
- `protocol` - (String) The protocol to enforce.
- `rule_id` - (String) The unique identifier of the rule.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_network_acl_rule` resource by using `id`.
The `id` property can be formed from `<network_acl_id>\<rule_id>`. For example:

```terraform
import {
  to = ibm_is_network_acl.example
  id = "<network_acl_id>\<rule_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_network_acl_rule.example <network_acl_id>\<rule_id>
```