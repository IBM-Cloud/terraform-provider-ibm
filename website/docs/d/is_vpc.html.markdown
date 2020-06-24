---
layout: "ibm"
page_title: "IBM : vpc"
sidebar_current: "docs-ibm-datasources-is-vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm\_is_vpc

Import the details of an existing IBM Virtual Private cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

data "ibm_is_vpc" "ds_vpc" {
  name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the VPC.

## Attribute Reference

The following attributes are exported:

* `crn` - The CRN of VPC.
* `status` - The status of VPC.
* `default_network_acl` - ID of the default network ACL.
* `classic_access` - Indicates whether this VPC is connected to Classic Infrastructure.
* `resource_group` - The resource group ID where the VPC created.
* `tags` - Tags associated with the instance.
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
