---
layout: "ibm"
page_title: "IBM : ibm_pi_network_security_group_rule"
description: |-
  Manages pi_network_security_group_rule.
subcategory: "Power Systems"
---

# ibm_pi_network_security_group_rule

Add or remove a network security group rule.

## Example Usage

```terraform
    resource "ibm_pi_network_security_group_rule" "network_security_group_rule" {
      pi_cloud_instance_id         = "<value of the cloud_instance_id>"
      pi_network_security_group_id = "<value of network_security_group_id>"
      pi_action                    = "allow"
      pi_destination_ports {
        minimum = 1200
        maximum = 37466
      }
      pi_source_ports {
        minimum = 1000
        maximum = 19500
      }
      pi_protocol {
        tcp_flags {
          flag = "ack"
        }
        tcp_flags {
          flag = "syn"
        }
        tcp_flags {
          flag = "psh"
        }
        type       = "tcp"
      }
      pi_remote {
        id   = "<value of remote_id>"
        type = "network-security-group"
      }
    }
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Timeouts

The `ibm_pi_network_security_group` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating a network security group rule.
- **delete** - (Default 10 minutes) Used for deleting a network security group rule.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_action` - (Optional, String) The action to take if the rule matches network traffic. Supported values are: `allow`, `deny`. Required if `pi_network_security_group_rule_id` is not provided.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_destination_port` - (Optional, List) The list of destination port.

    Nested schema for `pi_destination_port`:
      - `maximum` - (Optional, Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
      - `minimum` - (Optional, Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
- `pi_destination_ports` - (Deprecated, Optional, List) The list of destination port. Deprecated, please use `pi_destination_port`.

    Nested schema for `pi_destination_ports`:
      - `maximum` - (Optional, Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
      - `minimum` - (Optional, Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
- `pi_network_security_group_id` - (Required, String) The unique identifier of the network security group.
- `pi_network_security_group_rule_id` - (Optional, String) The network security group rule id to remove. Required if none of the other optional fields are provided.
- `pi_protocol` - (Optional, List) The list of protocol. Required if `pi_network_security_group_rule_id` is not provided.

    Nested schema for `pi_protocol`:
      - `icmp_type` - (Optional, String) If icmp type, a ICMP packet type affected by ICMP rules and if not present then all types are matched. Supported values are: `all`, `destination-unreach`, `echo`, `echo-reply`, `source-quench`, `time-exceeded`.
      - `tcp_flags` - (Optional, String) If tcp type, the list of TCP flags and if not present then all flags are matched. Supported values are: `syn`, `ack`, `fin`, `rst`.
      - `type` - (Required, String) The protocol of the network traffic. Supported values are: `icmp`, `tcp`, `udp`, `all`.
- `pi_remote` - (Optional, List) List of remote. Required if `pi_network_security_group_rule_id` is not provided.

    Nested schema for `pi_remote`:
      - `id` - (Optional, String) The id of the remote network address group or network security group the rules apply to. Not required for default-network-address-group.
      - `type` - (Optional, String) The type of remote group the rules apply to. Supported values are: `network-security-group`, `network-address-group`, `default-network-address-group`.
- `pi_source_port` - (Optional, List) List of source port

    Nested schema for `pi_source_port`:
      - `maximum` - (Optional, Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
      - `minimum` - (Optional, Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
- `pi_source_ports` - (Deprecated, Optional, List) List of source port. Deprecated, please use `pi_source_port`.

    Nested schema for `pi_source_ports`:
      - `maximum` - (Optional, Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
      - `minimum` - (Optional, Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.

- `pi_name` - (Optional, String) The name of the network security group rule. Required if `pi_network_security_group_rule_id` is not provided.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network security group's crn.
- `default` - (Boolean) Indicates if the network security group is the default network security group in the workspace.
- `id` - (String) The unique identifier of the network security group resource. Composed of `<cloud_instance_id>/<network_security_group_id/rule_id>`
- `members` - (List) The list of IPv4 addresses and\or network interfaces in the network security group.

    Nested schema for `members`:
  - `id` - (String) The id of the member in a network security group.
  - `mac_address` - (String) The mac address of a network interface included if the type is `network-interface`.
  - `network_interface_id` - (String) The network ID of a network interface included if the type is `network-interface`.
  - `target` - (String) If `ipv4-address` type, then IPv4 address or if `network-interface` type, then network interface id.
  - `type` - (String) The type of member. Supported values are: `ipv4-address`, `network-interface`.

- `network_security_group_id` -(String) The unique identifier of the network security group.
- `rules` - (List) The list of rules in the network security group.

    Nested schema for `rules`:
  - `action` - (String) The action to take if the rule matches network traffic. Supported values are: `allow`, `deny`.
  - `destination_port` - (List) The list of destination port.

        Nested schema for `destination_port`:
          - `maximum` - (Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
          - `minimum` - (Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
  - `id` - (String) The id of the rule in a network security group.
  - `name` - (String) The unique name of the network security group rule.
  - `protocol` - (List) The list of protocol.

        Nested schema for `protocol`:
          - `icmp_type` - (String) If icmp type, a ICMP packet type affected by ICMP rules and if not present then all types are matched. Supported values are: `all`, `destination-unreach`, `echo`, `echo-reply`, `source-quench`, `time-exceeded`.
          - `tcp_flags` - (String) If tcp type, the list of TCP flags and if not present then all flags are matched. Supported values are: `syn`, `ack`, `fin`, `rst`.
          - `type` - (String) The protocol of the network traffic. Supported values are: `icmp`, `tcp`, `udp`, `all`.
  - `remote` - (List) List of remote.

        Nested schema for `remote`:
        - `id` - (String) The id of the remote network address group or network security group the rules apply to. Not required for default-network-address-group.
        - `type` - (String) The type of remote group the rules apply to. Supported values are: `network-security-group`, `network-address-group`, `default-network-address-group`.
  - `source_port` - (List) List of source port

        Nested schema for `source_port`:
        - `maximum` - (Int) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
        - `minimum` - (Int) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
- `user_tags` - (List) List of user tags attached to the resource.

## Import

The `ibm_pi_network_security_group_rule` resource can be imported by using `cloud_instance_id`, `network_security_group_id` and `network_security_group_rule_id`.

## Example

```bash
terraform import ibm_pi_network_security_group_rule.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
