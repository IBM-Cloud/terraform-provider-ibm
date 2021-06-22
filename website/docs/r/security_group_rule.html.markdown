---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: security_group_rule"
description: |-
  Manages IBM Cloud Security Group rule.
---

# ibm_security_group_rule
Create, delete, and update a rule for a security group. You can set the IP range to manage incoming (ingress) and outgoing (egress) traffic to a virtual server instance. To create the security group, use the `security_group` resource. For more information, about security group rule, see [about security group](https://cloud.ibm.com/docs/security-groups?topic=security-groups-about-ibm-security-groups).

**Note**

For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API documentation](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_SecurityGroup_Rule).

## Example usage

```terraform
resource "ibm_security_group_rule" "allow_port_8080" {
    direction = "ingress"
    ether_type = "IPv4"
    port_range_min = 8080
    port_range_max = 8080
    protocol = "tcp"
    security_group_id = 123456
}
```

## Argument reference 
Review the argument references that you can specify for your resource.

- `direction` - (Required, String) The direction of traffic. Accepted values: `ingress` or `egress`.
- `ether_type`- (Optional, String) The IP version. Accepted values  (case-sensitive): `IPv4` or `IPv6`. Default value is `IPv4`.
- `port_range_min` - (Optional, Integer) The start of the port range for allowed traffic.
- `port_range_max` - (Optional, Integer) The end of the port range for allowed traffic.
- `protocol`- (Optional, String) The IP protocol type. Accepted values (case-sensitive): **icmp**,**tcp**, or **udp**.
- `remote_group_id` - (Optional, Integer) The ID of the remote security group allowed as part of the rule.  **Note** Conflicts with `remote_ip`.
- `remote_ip`- (Optional, String) The CIDR or IP address for allowed connections. **Note** Conflicts with `remote_group_id`.
- `security_group_id` - (Required,  Forces new resource, Integer) The ID of the security group this rule belongs to.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the security group rule.
