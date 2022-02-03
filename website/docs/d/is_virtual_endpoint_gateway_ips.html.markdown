---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_ips"
description: |-
  Manages IBM Cloud infrastructure virtual endpoint gateway IPs.
---

# ibm_is_virtual_endpoint_gateway_ips
Retrieve information of an existing IBM Cloud infrastructure virtual endpoint gateway IPs as a read only data source.  For more information, about the VPC endpoint gateways, see [about VPC gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-about-vpe).

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
data "ibm_is_virtual_endpoint_gateway_ips" "example" {
  gateway = ibm_is_virtual_endpoint_gateway.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `gateway` - (Required, String) The endpoint gateway ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `address` - (String) The endpoint gateway IP address.
- `auto_delete` - (String) The endpoint gateway IP auto delete.
- `created_at` - (Timestamp) The created date and time of the endpoint gateway IP.
- `id` - (String) The endpoint gateway reserved IP ID.
- `ips` - (String) Endpoint gateway reserved IP id
- `name` - (String) The endpoint gateway IP name.
- `reserved_ip` - (String) The endpoint gateway reserved IP ID.
- `resource_type` - (String) The endpoint gateway IP resource type.
- `target` - (List) The endpoint gateway target details.

  Nested scheme for `target`:
	- `id` - (String) The IPs target ID.
	- `name` - (String) The IPs target name.
	- `resource_type` - (String) The endpoint gateway resource type.
