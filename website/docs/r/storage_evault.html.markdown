---
layout: "ibm"
page_title: "IBM: storage_evault"
sidebar_current: "docs-ibm-resource-storage-evault"
description: |-
  Manages IBM Storage Evault.
---
# ibm\_storage_evault

Provides a evault storage resource. This allows [Evault Backup](https://console.bluemix.net/docs/infrastructure/Backup/index.html#getting-started-with-evault-backup-services) storage to be created, updated, and deleted.

## Example Usage

In the following example, you can create 20G of evault storage 

```hcl
resource "ibm_storage_evault" "test" {
  datacenter          = "dal05"
  capacity            = "20"
  virtual_instance_id = "62870765"
}
```

## Timeouts

ibm_storage_evault provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Instance.
* `update` - (Default 10 minutes) Used for Updating Instance.

## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center where you want to provision the evault storage instance.
* `capacity` - (Required, integer) The amount of storage capacity you want to allocate, specified in gigabytes.
* `virtual_instance_id` - (Optional, integer) The id of the virtual instance.
    **NOTE**: Conflicts with `hardware_instance_id`.
* `hardware_instance_id` - (Optional, integer) The id of the hardware instance.
    **NOTE**: Conflicts with `virtual_instance_id`.
* `tags` - (Optional, array of strings) Tags associated with the storage evault instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the evault.
* `username` - The username of the evault.
* `password` - The password of the evault.
* `service_resource_name` - The name of a evault storage network resource.
