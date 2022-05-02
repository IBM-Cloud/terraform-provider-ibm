---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_subnet_reserved_ip"
description: |-
  Manages IBM Subnet reserved IP.
---

# ibm_is_subnet_reserved_ip
Create, update, or delete a subnet. For more information, about associated reserved IP subnet, see [reserved IP subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-troubleshoot-reserved-ip).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
Sample to create a reserved IP:

```terraform
// Create a VPC
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

// Create a subnet
resource "ibm_is_subnet" "example" {
  name                     = "example-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

// Create the resrved IP in the following ways

// Only with Subnet ID
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet = ibm_is_subnet.example.id
}

// Subnet ID with a given name
resource "ibm_is_subnet_reserved_ip" "example1" {
  subnet = ibm_is_subnet.example.id
  name   = "example-subnet-reserved-ip1"
}

// Subnet ID with auto_delete
resource "ibm_is_subnet_reserved_ip" "example2" {
  subnet      = ibm_is_subnet.example.id
  auto_delete = true
}

// Subnet ID with both name and auto_delete
resource "ibm_is_subnet_reserved_ip" "example3" {
  subnet      = ibm_is_subnet.example.id
  name        = "example-subnet-reserved-ip3"
  auto_delete = true
}

// Subnet ID with address, name and auto_delete
resource "ibm_is_subnet_reserved_ip" "example4" {
  subnet      = ibm_is_subnet.example.id
  address     = "${replace(ibm_is_subnet.example.ipv4_cidr_block, "0/24", "14")}"
  name        = "example-subnet-reserved-ip4"
  auto_delete = true
}

// Create a virtual endpoint gateway and set as a target for reserved IP
resource "ibm_is_virtual_endpoint_gateway" "example" {
  name = "example-endpoint-gateway"
  target {
    name          = "ibm-ntp-server"
    resource_type = "provider_infrastructure_service"
  }
  vpc = ibm_is_vpc.example.id
}
resource "ibm_is_subnet_reserved_ip" "example5" {
  subnet = ibm_is_subnet.example.id
  name   = "example-subnet-reserved-ip5"
  target = ibm_is_virtual_endpoint_gateway.example.id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `address` - (Optional, Forces new resource, String) The IP address.
- `auto_delete`- (Optional, Bool)  If reserved IP is auto deleted.
- `name` - (Optional, String) The name of the reserved IP. ~> **NOTE:** raise  error if name is given with a prefix `ibm- `.
- `subnet` - (Required, Forces new resource, String) The subnet ID for the reserved IP.
- `target` - (Optional, string) The ID for the target endpoint gateway for the reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The date and time that the reserved IP was created.",
- `href` - (String) The URL for this reserved IP.
- `id` - (String) The combination of the subnet ID and reserved IP ID separated by **/**.
- `lifecycle_state` - (String) TThe lifecycle state of the reserved IP. [ deleting, failed, pending, stable, suspended, updating, waiting ]
- `owner` - (String) The owner of a reserved IP, defining whether it is managed by the user or the provider.
- `reserved_ip` - (String) The reserved IP.
- `resource_type` - (String) The resource type.
- `target` - (String) The ID for the target for the reserved IP.

## Import
The `ibm_is_subnet_reserved_ip` and `ibm_is_subnet` resource can be imported by using subnet ID and reserved IP ID separated by **/**.

**Syntax**

```
$ terraform import ibm_is_subnet.example <subnet_ID>/<subnet_reserved_IP_ID>
```

**Example**

```
$ terraform import ibm_is_subnet_reserved_ip.example 0716-13315ad8-d355-4041-bb60-62342000423/0716-617de4d8-5e2f-4d4a-b0d6-1000023
```
