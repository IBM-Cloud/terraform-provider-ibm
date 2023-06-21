---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm_is_vpc
Create, update, or delete a Virtual Private Cloud (VPC). VPCs allow you to create your own space in IBM Cloud to run an isolated environment within the public cloud. VPC gives you the security of a private cloud, with the agility and ease of a public cloud. For more information, about VPC, see [getting started with Virtual Private Cloud](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example to create a VPC:

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

```

## Timeouts
The `ibm_is_vpc` resource provides the following [[Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the VPC is considered `failed` when no response is received for 10 minutes. 
- **delete**: The deletion of the VPC is considered `failed` when no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the bare metal server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `address_prefix_management` - (Optional, Forces new resource, String) Indicates whether a default address prefix should be created automatically `auto` or manually `manual` for each zone in this VPC. Default value is `auto`.
- `classic_access` - (Optional, Bool) Specify if you want to create a VPC that can connect to classic infrastructure resources. Enter **true** to set up private network connectivity from your VPC to classic infrastructure resources that are created in the same IBM Cloud account, and **false** to disable this access. If you choose to not set up this access, you cannot enable it after the VPC is created. Make sure to review the [prerequisites](https://cloud.ibm.com/docs/vpc-on-classic-network?topic=vpc-on-classic-setting-up-access-to-your-classic-infrastructure-from-vpc#vpc-prerequisites) before you create a VPC with classic infrastructure access. Note that you can enable one VPC for classic infrastructure access per IBM Cloud account only.
- `default_network_acl_name` - (Optional, String) Enter the name of the default network access control list (ACL).
- `default_security_group_name` - (Optional, String) Enter the name of the default security group.
- `default_routing_table_name` - (Optional, String) Enter the name of the default routing table.
- `name` - (Required, String) Enter a name for your VPC. No.
- `no_sg_acl_rules` - (Optional, Bool) Delete all rules attached to default security group and default network ACL for a new VPC. This attribute has no impact on update.
- `resource_group` - (Optional, Forces new resource, String) Enter the ID of the resource group where you want to create the VPC. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the VPC is created in the `default` resource group. 
- `tags` - (Optional, Array of Strings) Enter any tags that you want to associate with your VPC. Tags might help you find your VPC more easily after it is created. Separate multiple tags with a comma (`,`).


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the VPC.
- `cse_source_addresses`- (List) A list of the cloud service endpoints that are associated with your VPC, including their source IP address and zone.
	- `address` - (String) The IP address of the cloud service endpoint.
	- `zone_name` - (String) The zone where the cloud service endpoint is located.
- `default_security_group_crn` - (String) CRN of the default security group created and attached to the VPC. 
- `default_security_group` - (String) The default security group ID created and attached to the VPC. 
- `default_network_acl_crn`-  (String) CRN of the default network ACL ID created and attached to the VPC.
- `default_network_acl`-  (String) The default network ACL ID created and attached to the VPC.
- `default_routing_table`-  (String) The unique identifier of the VPC default routing table.
- `id` - (String) The unique identifier of the VPC that you created.
- `subnets`- (List of Strings) A list of subnets that are attached to a VPC.

  Nested scheme for `subnets`:
  - `available_ipv4_address_count`-String - Available IPv4 addresses available for the usage in the subnet.
  - `available_ipv4_address_count`- (Integer) The number of IPv4 addresses in the subnet that are available for you to be used.
  - `id` - (String) The ID of the subnet.
  - `name` - (String) The name of the subnet.
  - `status` - (String) The status of the subnet.
  - `total_ipv4_address_count`- (Integer) The total number of IPv4 addresses in the subnet.
  - `zone` - (String) The Zone of the subnet. 
- `status` - (String) The provisioning status of your VPC. 
- `security_group` - (List) A list of security groups attached to VPC. 

  Nested scheme for `security_group`:
  - `group_id` - (String) The security group ID.
  - `group_name` - (String) The name of the security group.
  - `rules` - (List) Set of rules attached to a security group.
  
    Nested scheme for `rules`:
    - `code`- (String) The ICMP traffic code to allow.
	- `direction`- (String) The direction of the traffic either inbound or outbound.
    - `ip_version`-  (String) The IP version: **ipv4**.
    - `remote` -  (String) Security group ID, an IP address, a CIDR block, or a single security group identifier.
	- `rule_id` - (String) The rule ID.
    - `port_min` - (String) The inclusive lower bound of TCP port range.
    - `port_max` - (String) The inclusive upper bound of TCP port range.
	- `type` - (String) The ICMP traffic type to allow.


## Import
The `ibm_is_vpc` resource can be imported by using the VPC ID.

**Syntax**

```
$ terraform import ibm_is_vpc.example <vpc_ID>
```

**Example**

```
$ terraform import ibm_is_vpc.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
