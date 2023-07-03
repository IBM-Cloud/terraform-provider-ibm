---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateway"
description: |-
  Manages IBM Cloud public gateway.
---

# ibm_is_public_gateway
Retrieve information of an existing public gateway data source as a read only data source. For more information, about an VPC public gateway, see [about networking](https://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc).

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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_public_gateway" "example" {
  name = "example-gateway"
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
}

data "ibm_is_public_gateway" "example" {
  name = ibm_is_public_gateway.example.name
}
```

## Argument reference
Review the argument references that you can specify for your data source. 
 
- `name` - (Required, String) The name of the gateway.
- `resource_group` - (Optional, String) The resource group ID of the public gateway. **Note** This parameter is supported only for VPC Generation 2 infrastructure.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `access_tags`  - (List) Access management tags associated for the public gateway.
- `crn` - (String) The CRN for this public gateway.
- `floating_ip` - (List) List of the nested block describes the floating IP of the gateway with the **id** and **address** details.
	
  Nested scheme for `floating_ip`:
  - `id` - (String) The ID of the floating IP that is bound to the public gateway.
  - `address` - (String) The IP address of the floating IP that is bound to the public gateway.
- `id` - (String) The ID of the public gateway.
- `name` - (String) The name of the public gateway.
- `status` - (String) The status of the gateway.
- `tags` - (String) Tags associated with the Public gateway.
- `vpc` - (String) The VPC ID of the gateway.
- `zone` - (String) The public gateway zone name.
