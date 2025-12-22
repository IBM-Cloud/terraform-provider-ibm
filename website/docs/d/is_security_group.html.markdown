---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group"
description: |-
  Reads IBM Cloud security group.
---

# ibm_is_security_group
Retrieve information about a security group as a read-only data source. For more information, about managing IBM Cloud security group , see [about security group](https://cloud.ibm.com/docs/vpc?topic=vpc-using-security-groups).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example allows to create a different types of protocol rules `ALL`, `ICMP`, `UDP`, `TCP` and read the security group.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_security_group" "example" {
  name = "example-sg"
  vpc  = ibm_is_vpc.example.id
}

resource "ibm_is_security_group_rule" "example" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
}

resource "ibm_is_security_group_rule" "example" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  icmp {
    code = 20
    type = 30
  }
}

resource "ibm_is_security_group_rule" "example" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  udp {
    port_min = 805
    port_max = 807
  }
}

resource "ibm_is_security_group_rule" "example" {
  group     = ibm_is_security_group.example.id
  direction = "egress"
  remote    = "127.0.0.1"
  tcp {
    port_min = 8080
    port_max = 8080
  }
}

data "ibm_is_security_group" "example" {
  name = ibm_is_security_group.example.name
}

data "ibm_is_security_group" "examplevpc" {
  name = ibm_is_security_group.example.name
  vpc  = ibm_is_vpc.example.id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The name of the security group.
- `vpc` - (Optional, String) The identifier of the vpc where this security group resides. (Useful when two security groups have same name across different VPCs)
- `vpc_name` - (Optional, String) The name of the vpc where this security group resides. (Useful when two security groups have same name across different VPCs)
- `resource_group` - (Optional, String) The identifier of the resource group where this security group resides.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `access_tags`  - (List) Access management tags associated for the security group.
- `crn` - The CRN of the security group.
- `id` - (String) The ID of the security group.
- `rules` - (List of Objects) The rules associated with security group. Each rule has following attributes.

  Nested scheme for `rules`:
  - `rule_id`-  (String) ID of the rule.
  - `direction` - (String) Direction of traffic to enforce, either inbound or outbound.
  - `local` - (String) 	The local IP address or range of local IP addresses to which this rule will allow inbound traffic (or from which, for outbound traffic). A CIDR block of 0.0.0.0/0 allows traffic to all local IP addresses (or from all local IP addresses, for outbound rules). an IP address, a `CIDR` block.
  - `ip_version` - (String) IP version: IPv4
  - `protocol` - (String) The type of the protocol `all`, `icmp`, `tcp`, `udp`.
  - `type` - (String) The traffic type to allow.
  - `code` - (String) The traffic code to allow.
  - `port_max`- (Integer) The TCP/UDP port range that includes the maximum bound.
  - `port_min`- (Integer) The TCP/UDP port range that includes the minimum bound.
  - `remote`- (Integer)  Security group ID, an IP address, a CIDR block, or a single security group identifier.
- `tags` - Tags associated with the security group.
  


