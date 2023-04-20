---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpcs"
description: |-
  Manages IBM Cloud Virtual Private Cloud.
---

# ibm_is_vpcs
Retrieve information of an existing VPCs. For more information, about VPC, see [getting started with Virtual Private Cloud (VPC)](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started).

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
data "ibm_is_vpcs" "example" {
}

```
## Argument reference

Review the argument references that you can specify for your data source. 

- `resource_group` - (Optional, String) The ID of the Resource group this flow log collector belongs to
- `classic_access` - (Optional, Boolean) Indicates whether this VPC is connected to Classic Infrastructure.

## Attribute reference
You can access the following attribute references after your data source is created. 
- `vpcs` (List) List of all the VPCs.

  Nested scheme for `vpcs`:
    - `access_tags`  - (List) Access management tags associated for the volume.
    - `available_ipv4_address_count`- (Integer) The number of IPv4 addresses in the subnet that are available for you to be used.
    - `classic_access`- (Bool) Indicates whether this VPC is connected to the Classic Infrastructure.
    - `crn` - (String) The CRN of the VPC.
    - `cse_source_addresses`- (List of Cloud Service Endpoints) A list of the cloud service endpoints that are associated with your VPC, including their source IP address and zone.

      Nested scheme for `cse_source_addresses`:
      - `address` - (String) The IP address of the cloud service endpoint.
      - `zone_name` - (String) The zone where the cloud service endpoint is located.
    - `default_network_acl` - (String) The ID of the default network ACL.
    - `default_network_acl_crn` - (String) The CRN of the default network ACL.
    - `default_network_acl_name` - (String) The name of the default network ACL.
    - `default_security_group`-  (String) The unique identifier of the VPC default security group.
    - `default_security_group_crn` - (String) The CRN of the default security group.
    - `default_security_group_name` - (String) The name of the default security group.
    - `default_routing_table`-  (String) The unique identifier of the VPC default routing table.
    - `default_routing_table_name` - (String) The name of the default routing table.
    - `id` - (String) The ID of the VPC.
    - `name` - (String) The name of the VPC.
    - `resource_group` - (String) The resource group ID where the VPC created.
    - `security_group` - (String) A list of security groups attached to VPC. The nested security group block has the following structure:

      Nested scheme for `security_group`:
      - `group_id` - (String) Security group ID.
      - `group_name` - (String) Name of the security group.
      - `rules` -  (String) Set of rules attached to a security group.
      
        Nested scheme for `rules`:
        - `direction` - (String) Direction of the traffic either inbound or outbound.
        - `code` - (String) The ICMP traffic code to allow.
        - `ip_version` - (String) The IP version either **ipv4** or **ipv6**.
        - `port_min` - (String) The inclusive lower bound of TCP port range.
        - `port_max` - (String) The inclusive upper bound of TCP port range.
        - `remote` - (String) The security group ID, an IP address, a CIDR block, or a single security group identifier.
        - `rule_id` - (String) ID of the rule.
        - `type` - (String) The ICMP traffic type to allow.
    - `status` - (String) The status of the VPC.
    - `subnets`- (List) A list of subnets that are attached to a VPC.

      Nested scheme for `subnets`:
      - `id` - (String) The ID of the subnet.
      - `name` - (String) The name of the subnet.
      - `status` - (String) The status of the subnet.
      - `zone` - (String) The zone that the subnet belongs to.
    - `total_ipv4_address_count`- (Integer) The total number of IPv4 addresses in the subnet.
    - `tags` - (String) Tags associated with the instance.
