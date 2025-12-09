---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group_rule"
description: |-
  Manages IBM security group rule.
---

# ibm_is_security_group_rule
Create, update, or delete a security group rule. When you want to create a security group and security group rule for a virtual server instance in your VPC, you must create these resources in a specific order to avoid errors during the creation of your virtual server instance. For more information, about security group rule, see [security in your VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-security-in-your-vpc).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you create a different type of protocol rules.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_security_group" "example" {
  name = "example-security-group"
  vpc  = ibm_is_vpc.example.id
}

resource "ibm_is_security_group_rule" "example" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
}

resource "ibm_is_security_group_rule" "example1" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  # Deprecated block: replaced with 'protocol', 'code', and 'type' arguments
  # icmp {
  #   code = 20
  #   type = 30
  # }
  protocol  = "icmp"
  code      = 20
  type      = 30
}

resource "ibm_is_security_group_rule" "example2" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  # Deprecated block: replaced with 'protocol', 'port_min', and 'port_max' arguments
  # udp {
  #   port_min = 805
  #   port_max = 807
  # }
  protocol  = "udp"
  port_min = 805
  port_max = 807
}

resource "ibm_is_security_group_rule" "example3" {
  group     = ibm_is_security_group.example.id
  direction = "egress"
  remote    = "127.0.0.1"
  # Deprecated block: replaced with 'protocol', 'port_min', and 'port_max' arguments
  # tcp {
  #  port_min = 8080
  #  port_max = 8080
  # }
  protocol  = "tcp"
  port_min = 8080
  port_max = 8080
}

resource "ibm_is_security_group_rule" "example4" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  protocol   = "any"
}

resource "ibm_is_security_group_rule" "example_security_group_rule_icmp" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  # Deprecated block: replaced with 'protocol' argument
  # icmp {
  # }
  protocol  = "icmp"
}


resource "ibm_is_security_group_rule" "example_security_group_rule_udp" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  # Deprecated block: replaced with 'protocol' argument
  # udp {
  # }
  protocol  = "udp"
}

resource "ibm_is_security_group_rule" "example_security_group_rule_tcp" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  # Deprecated block: replaced with 'protocol' argument
  # tcp {
  # }
  protocol  = "tcp"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 
- `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed.
- `direction` - (Required, String) The direction of the traffic either `inbound` or `outbound`.
- `group` - (Required, Forces new resource, String) The security group ID.
- `local` - (String) 	The local IP address or range of local IP addresses to which this rule will allow inbound traffic (or from which, for outbound traffic). A CIDR block of 0.0.0.0/0 allows traffic to all local IP addresses (or from all local IP addresses, for outbound rules). an IP address, a `CIDR` block.
- `ip_version` - (Optional, String) The IP version to enforce. The format of local.address, remote.address, local.cidr_block or remote.cidr_block must match this property, if they are used. If remote references a security group, then this rule only applies to IP addresses (network interfaces) in that group matching this IP version. Supported value is [`ipv4`].
- `icmp` - (Optional, DEPRECATED, List) A nested block describes the `icmp` protocol of this security group rule. `icmp` is deprecated and use `protocol`, `code`, and `type` argument instead.

  Nested scheme for `icmp`:
  - `type`- (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all codes are allowed. 
  - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed.
- `port_min`- (Required, Integer) The TCP port range that includes the minimum bound. Valid values are from 1 to 65535.
- `port_max`- (Required, Integer) The TCP port range that includes the maximum bound. Valid values are from 1 to 65535.
- `protocol` - (Optional, String) The name of the network protocol.
- `remote` - (Optional, String) Security group ID, an IP address, a CIDR block, or a single security group identifier.
- `tcp` - (Optional, DEPRECATED, List) A nested block describes the `tcp` protocol of this security group rule. `tcp` is deprecated and use `protocol`, `port_min`, and `port_max` argument instead.

  Nested scheme for `tcp`:
  - `port_min`- (Required, Integer) The TCP port range that includes the minimum bound. Valid values are from 1 to 65535.
  - `port_max`- (Required, Integer) The TCP port range that includes the maximum bound. Valid values are from 1 to 65535.
- `type`- (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all codes are allowed.
- `udp` - (Optional, DEPRECATED, List) A nested block describes the `udp` protocol of this security group rule. `udp` is deprecated and use `protocol`, `port_min`, and `port_max` argument instead.

  Nested scheme for `udp`:
  - `port_min`- (Required, Integer) The UDP port range that includes minimum bound. Valid values are from 1 to 65535.
  - `port_max`- (Required, Integer) The UDP port range that includes maximum bound. Valid values are from 1 to 65535.

~> **Note:** Note: If no `protocol` block is specified; it creates a rule with protocol `icmp_tcp_udp`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the security group rule. The ID is composed of `<security_group_id>.<security_group_rule_id>`.
- `rule_id` - (String) The unique identifier of the rule.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_security_group_rule` resource by using `id`.
The `id` property can be formed from `security group ID`, and `security group rule ID`. For example:

```terraform
import {
  to = ibm_is_security_group_rule.example
  id = "<security_group_id>/<security_group_rule_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_security_group_rule.example <security_group_id>/<security_group_rule_id>
```