---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_subnet_reserved_ip_patch"
description: |-
  Manages IBM Subnet reserved IP patch.
---

# ibm_is_subnet_reserved_ip_patch
Update name and/or auto_delete of an existing reserved ip. For more information, about associated reserved IP subnet, see [reserved IP subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-troubleshoot-reserved-ip).
  
~> NOTE: Use this resource with caution, conflicts with `ibm_is_subnet_reserved_ip` resource if it has `name` attribute, using both will show changes on either of the resources alternatively on each apply.

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

resource "ibm_is_subnet_reserved_ip" "example" {
  subnet = ibm_is_subnet.example.id
}

resource "ibm_is_subnet_reserved_ip_patch" "example" {
  subnet      = ibm_is_subnet.example.id
  reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip

  name        = "test-reserved-ip"
  auto_delete = "true"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `auto_delete`- (Optional, Bool)  If reserved IP is auto deleted.
- `name` - (Required, String) The name of the reserved IP. 
  
  ~> **NOTE:** raise  error if name is given with a prefix `ibm- `.
- `subnet` - (Required, Forces new resource, String) The subnet ID for the reserved IP.
- `reserved_ip` - (Required, Forces new resource, string) The ID for the reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The date and time that the reserved IP was created.",
- `href` - (String) The URL for this reserved IP.
- `id` - (String) The combination of the subnet ID and reserved IP ID, separated by **/**.
- `lifecycle_state` - (String) The lifecycle state of the reserved IP. [ deleting, failed, pending, stable, suspended, updating, waiting ]
- `owner` - (String) The owner of a reserved IP, defining whether it is managed by the user or the provider.
- `reserved_ip` - (String) The reserved IP.
- `resource_type` - (String) The resource type.
- `target` - (String) The ID for the target for the reserved IP.
- `target_crn` - (String) The crn of the target for the reserved IP.

## Import
The `ibm_is_subnet_reserved_ip_patch` resource can be imported by using subnet ID and reserved IP ID separated by **/**.

**Syntax**

```
$ terraform import ibm_is_subnet_reserved_ip_patch.example <subnet_ID>/<subnet_reserved_IP_ID>
```

**Example**

```
$ terraform import ibm_is_subnet_reserved_ip_patch.example 0716-13315ad8-d355-4041-bb60-62342000423/0716-617de4d8-5e2f-4d4a-b0d6-1000023
```
