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
    resource "ibm_pi_network_security_group_rules" "network_security_group" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_name = "name"
        pi_user_tags = ["tag1", "tag2"]
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

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_name` - (Required, String) The name of the Network Security Group.
- `pi_user_tags` - (Optional, List) A list of tags.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network security group's crn.

- `members` - (List) The list of IPv4 addresses and\or network interfaces in the network security group.

    Nested schema for `members`:
  - `id` - (String) The id of the member in a network security group.
  - `mac_address` - (String) The mac address of a network interface included if the type is `network-interface`.
  - `target` - (String) If `ipv4-address` type, then IPv4 address or if `network-interface` type, then network interface id.
  - `type` - (String) The type of member. Supported values are: `ipv4-address`, `network-interface`.

- `network_security_group_id` -(String) The unique identifier of the network security group.
- `rules` - (List) The list of rules in the network security group.

    Nested schema for `rules`:
  - `action` - (String) The action to take if the rule matches network traffic. Supported values are: `allow`, `deny`.
  - `destination_port` - (List) The list of destination port.

        Nested schema for `destination_port`:
        - `maximum` - (Float) The end of the port range, if applicable, If values are not present then all ports are in the range.
        - `minimum` - (Float) The start of the port range, if applicable. If values are not present then all ports are in the range.
  - `direction` - (String) The direction of the network traffic. Supported values are: `inbound`, `outbound`.
  - `id` - (String) The id of the rule in a network security group.
  - `name` - (String) The unique name of the network security group rule.
  - `protocol` - (List) The list of protocol.

        Nested schema for `protocol`:
        - `icmp_types` - (List) If icmp type, the list of ICMP packet types (by numbers) affected by ICMP rules and if not present then all types are matched.
        - `tcp_flags` - (String) If tcp type, the list of TCP flags and if not present then all flags are matched. Supported values are: `syn`, `ack`, `fin`, `rst`, `urg`, `psh`, `wnd`, `chk`, `seq`.
        - `type` - (String) The protocol of the network traffic. Supported values are: `icmp`, `tcp`, `udp`, `all`.
  - `remote` - (List) List of remote.

        Nested schema for `remote`:
        - `id` - (String) The id of the remote network Address group or network security group the rules apply to. Not required for default-network-address-group.
        - `type` - (String) The type of remote group the rules apply to. Supported values are: `network-security-group`, `network-address-group`, `default-network-address-group`.
  - `source_port` - (List) List of source port

        Nested schema for `source_port`:
        - `maximum` - (Float) The end of the port range, if applicable, If values are not present then all ports are in the range.
        - `minimum` - (Float) The start of the port range, if applicable. If values are not present then all ports are in the range.
