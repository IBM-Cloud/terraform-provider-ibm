---
layout: "ibm"
page_title: "IBM : ibm_pi_network_security_group"
description: |-
  Get information about pi_network_security_group
subcategory: "Power Systems"
---

# ibm_pi_network_security_group

Retrieves information about a network security group.

## Example Usage

```terraform
    data "ibm_pi_network_security_group" "network_security_group" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_network_security_group_id = "<value of the network_security_group_id>"
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

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_network_security_group_id` - (Required, String) network security group id or name.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `crn` - (String) The network security group's crn.

- `members` - (List) The list of IPv4 addresses and\or network interfaces in the network security group.

  Nested schema for `members`:
  - `id` - (String) The id of the member in a network security group.
  - `mac_address` - (String) The mac address of a network interface included if the type is `network-interface`.
  - `target` - (String) If `ipv4-address` type, then IPv4 address or if `network-interface` type, then network interface id.
  - `type` - (String) The type of member. Supported values are: `ipv4-address`, `network-interface`.

- `name` - (String) The name of the network security group.

- `rules` - (List) The list of rules in the network security group.

  Nested schema for `rules`:
  - `action` - (String) The action to take if the rule matches network traffic. Supported values are: `allow`, `deny`.
  - `destination_port` - (List) The list of destination port.

      Nested schema for `destination_port`:
        - `maximum` - (Integer) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
        - `minimum` - (Integer) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
  - `id` - (String) The id of the rule in a network security group.
  - `protocol` - (List) The list of protocol.

      Nested schema for `protocol`:
        - `icmp_type` - (String) If icmp type, a ICMP packet type affected by ICMP rules and if not present then all types are matched.
        - `tcp_flags` - (List) If tcp type, the list of TCP flags and if not present then all flags are matched. Supported values are: `syn`, `ack`, `fin`, `rst`, `urg`, `psh`, `wnd`, `chk`, `seq`.
         Nested schema for `icmp_types`:
         - `flags` - (String) TCP flag.
        - `type` - (String) The protocol of the network traffic. Supported values are: `icmp`, `tcp`, `udp`, `all`.
  - `remote` - (List) List of remote.

      Nested schema for `remote`:
        - `id` - (String) The id of the remote network Address group or network security group the rules apply to. Not required for default-network-address-group.
        - `type` - (String) The type of remote group the rules apply to. Supported values are: `network-security-group`, `network-address-group`, `default-network-address-group`.
  - `source_port` - (List) List of source port.

      Nested schema for `source_port`:
        - `maximum` - (Integer) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
        - `minimum` - (Integer) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum port number.
- `user_tags` - (List) List of user tags attached to the resource.
