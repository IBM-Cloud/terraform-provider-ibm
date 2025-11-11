---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disk_management"
description: |-
  Manages IBM Dedicated host disk management.
---

# ibm_is_dedicated_host_disk_management

Create, update, delete and suspend the dedicated host disk management resource. For more information, about dedicated host disk management, see [migrating a dedicated host instance to another host](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-migrating-dedicated-host).

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
resource "ibm_is_dedicated_host_group" "example" {
  family = "memory"
  class  = "beta"
  zone   = "us-south-1"
}
data "ibm_is_dedicated_host_group" "example" {
  name = ibm_is_dedicated_host_group.example.name
}
resource "ibm_is_dedicated_host" "example" {
  profile    = "dh2-56x464"
  host_group = ibm_is_dedicated_host_group.example.id
  name       = "example-dedicated-host"
}
data "ibm_is_dedicated_host" "example" {
  host_group = ibm_is_dedicated_host_group.example.id
  name       = "example-dedicated-host"
}
data "ibm_is_dedicated_host_disks" "example" {
  dedicated_host = data.ibm_is_dedicated_host.example.id
}
data "ibm_is_dedicated_host_disk" "example" {
  dedicated_host = data.ibm_is_dedicated_host.example.id
  disk           = ibm_is_dedicated_host_disk_management.example.disks.0.id
}
resource "ibm_is_dedicated_host_disk_management" "example" {
  dedicated_host = data.ibm_is_dedicated_host.example.id
  disks {
    name = "example-disks"
    id   = data.ibm_is_dedicated_host.example.disks.0.id
  }
}
```

## Argument reference
Review the argument reference that you can specify for your resource.

- `dedicated_host` - (Required, Force New Resource, String) The unique-identifier of the dedicated host.
- `disks` - (Required, List) Disks that needs to be updated. Nested `disks` blocks have the following structure:
  
  Nested scheme for `disks`:
  - `id` - (Required, String) The unique-identifier of the dedicated host disk.
  - `name` - (Required, String) The unique user defined name for the dedicated host disk.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique-identifier of the dedicated host disk management.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_dedicated_host_disk_management` resource by using `id`.
The `id` property can be formed from `dedicated host ID`. For example:

```terraform
import {
  to = ibm_is_dedicated_host_disk_management.example
  id = "<dedicated_host_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_dedicated_host_disk_management.example <dedicated_host_id>
```