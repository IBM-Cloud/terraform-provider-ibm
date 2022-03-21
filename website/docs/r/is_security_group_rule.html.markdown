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
In the following example, you create a different type of protocol rules `ALL`, `ICMP`, `UDP`, `TCP` and `ANY`.

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
  icmp {
    code = 20
    type = 30
  }
}

resource "ibm_is_security_group_rule" "example2" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  udp {
    port_min = 805
    port_max = 807
  }
}

resource "ibm_is_security_group_rule" "example3" {
  group     = ibm_is_security_group.example.id
  direction = "egress"
  remote    = "127.0.0.1"
  tcp {
    port_min = 8080
    port_max = 8080
  }
}

resource "ibm_is_security_group_rule" "example_security_group_rule_icmp_any" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  icmp {
  }
}

resource "ibm_is_security_group_rule" "example_security_group_rule_udp_any" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  udp {
  }
}

resource "ibm_is_security_group_rule" "example_security_group_rule_tcp_any" {
  group      = ibm_is_security_group.example_security_group.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  tcp {
  }
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `direction` - (Required, String) The direction of the traffic either `inbound` or `outbound`.
- `group` - (Required, Forces new resource, String) The security group ID.
- `ip_version` - (Optional, String) The IP version either `IPv4` or `IPv6`. Default `IPv4`.
- `icmp` - (Optional, List) A nested block describes the `icmp` protocol of this security group rule.

  Nested scheme for `icmp`:
  - `type`- (Required, Integer) The ICMP traffic type to allow. Valid values from 0 to 254.
  - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255.
- `remote` - (Optional, String) Security group ID, an IP address, a CIDR block, or a single security group identifier.
- `tcp` - (Optional, List) A nested block describes the `tcp` protocol of this security group rule.

  Nested scheme for `tcp`:
  - `port_min`- (Required, Integer) The TCP port range that includes the minimum bound. Valid values are from 1 to 65535.
  - `port_max`- (Required, Integer) The TCP port range that includes the maximum bound. Valid values are from 1 to 65535.
- `udp` - (Optional, List) A nested block describes the `udp` protocol of this security group rule.

  Nested scheme for `udp`:
  - `port_min`- (Required, Integer) The UDP port range that includes minimum bound. Valid values are from 1 to 65535.
  - `port_max`- (Required, Integer) The UDP port range that includes maximum bound. Valid values are from 1 to 65535.

~> **Note:** 

If any of the `icmp` , `tcp`, or `udp` is not specified it creates a rule with protocol `ALL`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the security group rule. The ID is composed of `<security_group_id>.<security_group_rule_id>`.
- `rule_id` - (String) The unique identifier of the rule.


## Import
The `ibm_is_security_group_rule` resource can be imported by using security group ID and security group rule ID.

**Example**

```
$ terraform import ibm_is_security_group_rule.example d7bec597-4726-451f-8a63-e62e6f19c32c.cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```


