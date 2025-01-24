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

## Example usage
The following example to create a VPC with dns:

```terraform
// manual type resolver
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
  dns {
		enable_hub = true
		resolver {
			manual_servers {
				address = "192.168.3.4"
			}
		}
	}
}

resource "ibm_is_vpc" "example_vpc_manual" {
	name = "example-vpc-manual"
	dns {
		enable_hub = true
		resolver {
			manual_servers {
				address ="192.168.0.4"
				zone_affinity= "au-syd-1"
			}
			manual_servers {
				address =  "192.168.64.4"
				zone_affinity = "au-syd-2"
			}
			manual_servers {
				address= "192.168.128.4"
				zone_affinity ="au-syd-3"
			}
		}
	}
}

// system type resolver
resource "ibm_is_vpc" "example-system" {
	name = "example-system-vpc"
	dns {
		enable_hub = false

    // uncommenting/patching vpc with below code would make the resolver type delegated
    # resolver {
		# 	type = "delegated"
		# 	vpc_id = ibm_is_vpc.example.id
		# }
	}
}

// delegated type resolver

resource "ibm_is_vpc" "example-delegated" {
  // required : add a dependency on ibm dns custom resolver of the hub vpc
	depends_on = [ ibm_dns_custom_resolver.example-hub ]
	name = "example-hub-false-delegated"
	dns {
		enable_hub = false
		resolver {
			type = "delegated"
			vpc_id = ibm_is_vpc.example.id
			dns_binding_name = "example-vpc-binding"
		}
	}
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

  ~> **Note:** 
    `classic_access` is deprecated. Use [Transit Gateway](https://cloud.ibm.com/docs/transit-gateway) with Classic as a spoke/connection.
- `default_network_acl_name` - (Optional, String) Enter the name of the default network access control list (ACL).
- `default_security_group_name` - (Optional, String) Enter the name of the default security group.
- `default_routing_table_name` - (Optional, String) Enter the name of the default routing table.

- `dns` - (Optional, List) The DNS configuration for this VPC.
  
  Nested scheme for `dns`:
  - `enable_hub` - (Optional, Boolean) Indicates whether this VPC is enabled as a DNS name resolution hub.
  - `resolver` - (Optional, List) The zone list this backup policy plan will create snapshot clones in.
    Nested scheme for `resolver`:

      - `dns_binding_id` - (String) The VPC dns binding id whose DNS resolver provides the DNS server addresses for this VPC. (If any)
      - `dns_binding_name` - (Optional, String) The VPC dns binding name whose DNS resolver provides the DNS server addresses for this VPC. Only applicable for `delegated`, providing value would create binding with this name.

        ~> **Note:** 
          `manual_servers` must be set if and only if `dns.resolver.type` is manual.
      - `manual_servers` - (Optional, List) The DNS servers to use for this VPC, replacing any existing servers. All the DNS servers must either: **have a unique zone_affinity**, or **not have a zone_affinity**.

          Nested schema for **manual_servers**:

          - `address` - (Required, String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
          - `zone_affinity` - (Optional, List) If present, DHCP configuration for this zone will have this DNS server listed first.

        ~> **Note:** 
          While using `zone_affinity`, if fewer DNS servers are specified than the number of zones, then default servers will be created and terraform would show change. Its advised to provide `address` for all `zone_affinity`.


      - `type` - (Optional, String) The type of the DNS resolver to use. To update the resolver type, specify the `type` explicitly.

        ~> **Note:** 
          `delegated`: DNS server addresses will be provided by the resolver for the VPC specified in dns.resolver.vpc. Requires dns.enable_hub to be false.<br/>
          `manual`: DNS server addresses are specified in `manual_servers`.<br/>
          `system`: DNS server addresses will be provided by the system and depend on the configuration.

        ~> **Note:** 
              Updating from `manual` requires dns resolver `manual_servers` to be specified as null.<br/>
              Updating to `manual` requires dns resolver `manual_servers` to be specified and not empty.<br/>
              Updating from `delegated` requires `dns.resolver.vpc` to be specified as null. If type is `delegated` while creation then `vpc_id` is required
      - `vpc_id` - (Optional, List) (update only) The VPC ID to provide DNS server addresses for this VPC. The specified VPC must be configured with a DNS Services custom resolver and must be in one of this VPC's DNS resolution bindings. Mutually exclusive with `vpc_crn`

        ~> **Note:** 
          Specify "null" string to remove an existing VPC.<br/>
          This property must be set if and only if dns resolver type is `delegated`.
      - `vpc_crn` - (Optional, List) (update only) The VPC CRN to provide DNS server addresses for this VPC. The specified VPC must be configured with a DNS Services custom resolver and must be in one of this VPC's DNS resolution bindings. Mutually exclusive with `vpc_id`

        ~> **Note:** 
          Specify "null" string to remove an existing VPC.<br/>
          This property must be set if and only if dns resolver type is `delegated`.


- `name` - (Required, String) Enter a name for your VPC. No.
- `no_sg_acl_rules` - (Optional, Bool) If set to true, delete all rules attached to default security group and default network ACL for a new VPC. This attribute has no impact on update. default false.
- `resource_group` - (Optional, Forces new resource, String) Enter the ID of the resource group where you want to create the VPC. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the VPC is created in the `default` resource group. 
- `tags` - (Optional, Array of Strings) Enter any tags that you want to associate with your VPC. Tags might help you find your VPC more easily after it is created. Separate multiple tags with a comma (`,`).


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the VPC.
- `cse_source_addresses`- (List) A list of the cloud service endpoints that are associated with your VPC, including their source IP address and zone.
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
- `default_security_group_crn` - (String) CRN of the default security group created and attached to the VPC. 
- `default_security_group` - (String) The default security group ID created and attached to the VPC. 
- `default_network_acl_crn`-  (String) CRN of the default network ACL ID created and attached to the VPC.
- `default_network_acl`-  (String) The default network ACL ID created and attached to the VPC.
- `default_routing_table`-  (String) The unique identifier of the VPC default routing table.
- `default_routing_table_crn`-  (String) CRN of the default routing table.
- `health_reasons` - (List) The reasons for the current `health_state` (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
  Nested schema for **health_reasons**:
	- `code` - (String) A snake case string succinctly identifying the reason for this health state.
	- `message` - (String) An explanation of the reason for this health state.
	- `more_info` - (String) Link to documentation about the reason for this health state.

- `health_state` - (String) The health of this resource.- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.[`degraded`, `faulted`, `inapplicable`, `ok`]
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
