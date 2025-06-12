---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm_is_vpc
Retrieve information of an existing IBM Virtual Private cloud. For more information, about VPC, see [getting started with Virtual Private Cloud (VPC)](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started).

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
data "ibm_is_vpc" "example" {
  name = "example-vpc"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

~> **Note:** `name` and `identifier` are mutually exclusive. One of them is required.
- `name` - (Optional, String) The name of the VPC.
- `identifier` - (Optional, String) The id of the VPC.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 
- `access_tags`  - (List) Access management tags associated for the volume.
- `available_ipv4_address_count`- (Integer) The number of IPv4 addresses in the subnet that are available for you to be used.
- `classic_access`- (Bool) Indicates whether this VPC is connected to Classic Infrastructure.

  ~> **Note:** 
    `classic_access` is deprecated. Use [Transit Gateway](https://cloud.ibm.com/docs/transit-gateway) with Classic as a spoke/connection.
- `crn` - (String) The CRN of the VPC.
- `cse_source_addresses`-List of Cloud Service Endpoints-A list of the cloud service endpoints that are associated with your VPC, including their source IP address and zone.

  Nested scheme for `cse_source_addresses`:
	- `address` - (String) The IP address of the cloud service endpoint.
	- `zone_name` - (String) The zone where the cloud service endpoint is located.
- `default_address_prefixes` - (Map) A map of default address prefixes for each zone in the VPC. The keys are the zone names, and the values are the corresponding address prefixes.
  Example:
  ```hcl
    default_address_prefixes    = {
        "us-south-1" = "10.240.0.0/18"
        "us-south-2" = "10.240.64.0/18"
        "us-south-3" = "10.240.128.0/18"
        "us-south-4" = "10.240.192.0/18"
    }
  ```
- `default_network_acl` - (String) The ID of the default network ACL.
- `default_network_acl_crn` - (String)  The CRN of the default network ACL.
- `default_network_acl_name` - (String)  The name of the default network ACL.
- `default_security_group`-  (String) The unique identifier of the VPC default security group.
- `default_security_group_crn` - (String) The CRN of the default security group.
- `default_security_group_name` - (String) The name of the default security group.
- `default_routing_table`-  (String) The unique identifier of the VPC default routing table.
- `default_routing_table_name` - (String) The name of the default routing table.

- `dns` - (List) The DNS configuration for this VPC.
  
  Nested scheme for `dns`:
  - `enable_hub` - (Boolean) Indicates whether this VPC is enabled as a DNS name resolution hub.
  - `resolver` - (List) The zone list this backup policy plan will create snapshot clones in.
    
    Nested scheme for `resolver`:
      - `manual_servers` - (Integer) The DNS servers to use for this VPC, replacing any existing servers. All the DNS servers must either: **have a unique zone_affinity**, or **not have a zone_affinity**.  
      - `type` - (String) The type of the DNS resolver to use.
      - `vpc` - (String) The VPC to provide DNS server addresses for this VPC. The specified VPC must be configured with a DNS Services custom resolver and must be in one of this VPC's DNS resolution bindings.
- `health_reasons` - (List) The reasons for the current `health_state` (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
  Nested schema for **health_reasons**:
	- `code` - (String) A snake case string succinctly identifying the reason for this health state.
	- `message` - (String) An explanation of the reason for this health state.
	- `more_info` - (String) Link to documentation about the reason for this health state.

- `health_state` - (String) The health of this resource.- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.[`degraded`, `faulted`, `inapplicable`, `ok`]
- `resource_group` - (String) The resource group ID where the VPC created.
- `security_group` - (String) A list of security groups attached to VPC. The nested security group block has the following structure:

  Nested scheme for `security_group`:
  - `group_id` - (String) Security group ID.
  - `group_name` - (String) Name of the security group.
  - `rules` -  (String) Set of rules attached to a security group.
   
    Nested scheme for `rules`:
    - `direction` - (String) Direction of the traffic either inbound or outbound.
    - `code` - (String) The ICMP traffic code to allow.
    - `ip_version` - (String) The IP version **ipv4**
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
