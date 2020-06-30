---
layout: "ibm"
page_title: "IBM : vpc"
sidebar_current: "docs-ibm-resource-is-vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm\_is_vpc

Provides a vpc resource. This allows VPC to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPC:

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

```

## Timeouts

ibm_is_vpc provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating vpc.
* `delete` - (Default 10 minutes) Used for Deleting vpc.


## Argument Reference

The following arguments are supported:

* `default_network_acl` - (Deprecated, string) ID of the default network ACL.
* `is_default` - (Removed, bool) This field is removed.
* `address_prefix_management` - (Optional, string) Indicates whether a default address prefix should be automatically created for each zone in this VPC. Default value `auto`
* `classic_access` -(Optional, bool) Indicates whether this VPC should be connected to Classic Infrastructure. If true, This VPC's resources will have private network connectivity to the account's Classic Infrastructure resources. Only one VPC on an account may be connected in this way. 
* `name` - (Required, string) The name of the VPC.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID where the VPC to be created
* `tags` - (Optional, array of strings) Tags associated with the instance.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPC.
* `crn` - The CRN of VPC.
* `default_security_group` - The unique identifier of the VPC default security group.
* `status` - The status of VPC.
* `cse_source_addresses` - A list describing the cloud service endpoint source ip adresses and zones. The nested cse_source_addresses block have the following structure:
  * `address` - Ip Address of the cloud service endpoint.
  * `zone_name` - Zone associated with the IP Address.
* `subnets` - A list of subnets attached to VPC. The nested subnets block have the following structure:
  * `name` - Name of the subnet.
  * `id` - ID of the subnet.
  * `status` -  Status of the subnet.
  * `zone` -  Zone of the subnet.
  * `total_ipv4_address_count` - Total IPv4 addresses under the subnet.
  * `available_ipv4_address_count` - Available IPv4 addresses available for the usage in the subnet.  


## Import

ibm_is_vpc can be imported using ID, eg

```
$ terraform import ibm_is_vpc.example d7bec597-4726-451f-8a63-e62e6f19c32c
```