---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_subnet_reserved_ip"
description: |-
  Manages IBM Subnet reserved IP.
---

# ibm_is_subnet_reserved_ip
Create, update, or delete a subnet. For more information, about associated reserved IP subnet, see [reserved IP subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-troubleshoot-reserved-ip).

## Example usage
Sample to create a reserved IP:

```terraform
    // Create a VPC
    resource "ibm_is_vpc" "vpc1" {
        name = "my-vpc"
    }

    // Create a subnet
    resource "ibm_is_subnet" "subnet1" {
        name                     = "my-subnet"
        vpc                      = ibm_is_vpc.vpc1.id
        zone                     = "us-south-1"
        total_ipv4_address_count = 256
    }

    // Create the resrved IP in the following ways

    // Only with Subnet ID
    resource "ibm_is_subnet_reserved_ip" "res_ip" {
        subnet = ibm_is_subnet.subnet1.id
    }

    // Subnet ID with a given name
    resource "ibm_is_subnet_reserved_ip" "res_ip_name" {
        subnet = ibm_is_subnet.subnet1.id
        name = "my-subnet-reserved-ip"
    }

    // Subnet ID with auto_delete
    resource "ibm_is_subnet_reserved_ip" "res_ip_auto_delete" {
        subnet = ibm_is_subnet.subnet1.id
        auto_delete = true
    }

    // Subnet ID with both name and auto_delete
    resource "ibm_is_subnet_reserved_ip" "res_ip_auto_delete_name" {
        subnet = ibm_is_subnet.subnet1.id
        name = "my-subnet-reserved-ip"
        auto_delete = true
    }

        // Create a virtual endpoint gateway and set as a target for reserved IP
    resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
        name = "my-endpoint-gateway-1"
        target {
            name          = "ibm-ntp-server"
            resource_type = "provider_infrastructure_service"
        }
        vpc = ibm_is_vpc.vpc1.id
    }
    resource "ibm_is_subnet_reserved_ip" "reserved_ip_1" {
        subnet = ibm_is_subnet.subnet1.id
        name = "my-subnet-reserved-ip1"
        target = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
    }
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `auto_delete`- (Optional, Bool)  If reserved IP is auto deleted.
- `name` - (Optional, String) The name of the reserved IP. **NOTE** raise  error if name is given with a prefix `ibm- `.
- `subnet` - (Required, Forces new resource, String) The subnet ID for the reserved IP.
- `target` - (Optional, string) The ID for the target endpoint gateway for the reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The IP address.
- `created_at` - (Timestamp) The date and time that the reserved IP was created.",
- `href` - (String) The URL for this reserved IP.
- `id` - (String) The combination of the subnet ID and reserved IP ID separated by **/**.
- `owner` - (String) The owner of a reserved IP, defining whether it is managed by the user or the provider.
- `reserved_ip` - (String) The reserved IP.
- `resource_type` - (String) The resource type.
- `target` - (String) The ID for the target endpoint gateway for the reserved IP.

## Import
The `ibm_is_subnet_reserved_ip` and `ibm_is_subnet` resource can be imported by using subnet ID and reserved IP ID separated by **/**.

**Syntax**

```
$ terraform import ibm_is_subnet.example <reserved_subnet_IP>
```

**Example**

```
$ terraform import ibm_is_subnet_reserved_ip.example 0716-13315ad8-d355-4041-bb60-62342000423/0716-617de4d8-5e2f-4d4a-b0d6-1000023
```
