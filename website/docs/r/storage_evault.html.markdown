---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: storage_evault"
description: |-
  Manages IBM Cloud storage EVault.
---
# ibm_storage_evault
Create, delete, and update a EVault storage resource. For more information, about storage evalut [getting started with IBM Cloud backup](https://cloud.ibm.com/docs/infrastructure/Backup/index.html#getting-started-with-evault-backup-services).

## Example usage
In the following example, you can create 20G of EVault storage 

```terraform
resource "ibm_storage_evault" "test" {
  datacenter          = "dal05"
  capacity            = "20"
  virtual_instance_id = "62870765"
}
```

## Timeouts
The `ibm_storage_evault` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating instance.
- **update** - (Default 10 minutes) Used for updating instance.


## Argument reference
Review the argument references that you can specify for your resource.

- `datacenter` - (Required, Forces new resource, String) The data center where you want to provision the EVault  storage instance.
- `capacity` - (Required, Integer) The amount of storage capacity that you want to allocate, specified in gigabytes.
- `hardware_instance_id` - (Optional, Forces new resource, Integer) The ID of the hardware instance. **Note** Conflicts with `virtual_instance_id`.
- `tags` - (Optional, Array of string) Tags associated with the storage EVault instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `virtual_instance_id` - (Optional, Forces new resource, Integer) The ID of the virtual instance. **Note** Conflicts with `hardware_instance_id`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id`- (String) The unique identifier of the EVault.
- `password`- (String) The password of the EVault.
- `service_resource_name`- (String) The name of an EVault storage network resource.
- `username`- (String) The username of the EVault.
