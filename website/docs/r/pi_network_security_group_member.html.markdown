---
layout: "ibm"
page_title: "IBM : ibm_pi_network_security_group_member"
description: |-
  Manages pi_network_security_group_member.
subcategory: "Power Systems"
---

# ibm_pi_network_security_group_member

Add or remove a network security group member.

## Example Usage

```terraform
    resource "ibm_pi_network_security_group_member" "network_security_group_member" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_network_security_group_id = "network_security_group_id"
        pi_target = "target"
        pi_type = "ipv4-address"
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

The `ibm_pi_network_security_group_member` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 5 minutes) Used for deleting a network security group member.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_network_security_group_id` - (Required, String) Network security group ID.
- `pi_network_security_group_member_id` - (Optional, String) Network security group member ID. This conflicts with `pi_target` and `pi_type`.
- `pi_target` - (Optional, String) The target member to add. Required with `pi_type`. This conflicts with `pi_network_security_group_member_id`.
- `pi_type` - (Optional, String) The type of member. Supported values are: `ipv4-address`, `network-interface`. Required with `pi_target`. This conflicts with `pi_network_security_group_member_id`.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network security group's crn.
- `default` - (Boolean) Indicates if the network security group is the default network security group in the workspace.
- `id` - (String) The unique identifier of the network security group resource. Composed of `<cloud_instance_id>/<network_security_group_id>/<network_security_group_member_id>`
- `members` - (List) The list of IPv4 addresses and\or network interfaces in the network security group.

    Nested schema for `members`:
      - `id` - (String) The id of the member in a network security group.
      - `mac_address` - (String) The mac address of a network interface included if the type is `network-interface`.
      - `network_interface_id` - (String) The network ID of a network interface included if the type is `network-interface`.
      - `target` - (String) If `ipv4-address` type, then IPv4 address or if `network-interface` type, then network interface id.
      - `type` - (String) The type of member. Supported values are: `ipv4-address`, `network-interface`.

- `name` - (String) The name of the network security group.
- `network_security_group_member_id` - (String) The unique identifier of the network security group resource.
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
          - `tcp_flags` - (String) If tcp type, the list of TCP flags and if not present then all flags are matched. Supported values are: `syn`, `ack`, `fin`, `rst`.
          - `type` - (String) The protocol of the network traffic. Supported values are: `icmp`, `tcp`, `udp`, `all`.
      - `remote` - (List) List of remote.
        
        Nested schema for `remote`:
          - `id` - (String) The id of the remote network Address group or network security group the rules apply to. Not required for default-network-address-group.
          - `type` - (String) The type of remote group the rules apply to. Supported values are: `network-security-group`, `network-address-group`, `default-network-address-group`.
      - `source_port` - (List) List of source port

          Nested schema for `source_port`:
          - `maximum` - (Integer) The end of the port range, if applicable. If the value is not present then the default value of 65535 will be the maximum port number.
          - `minimum` - (Integer) The start of the port range, if applicable. If the value is not present then the default value of 1 will be the minimum 
- `user_tags` - (List) List of user tags attached to the resource.

## Import

The `ibm_pi_network_security_group_member` resource can be imported by using `cloud_instance_id`, `network_security_group_id` and `network_security_group_member_id`.

## Example

```bash
terraform import ibm_pi_network_security_group_member.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
