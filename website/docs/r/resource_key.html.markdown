---
layout: "ibm"
page_title: "IBM : resource_key"
sidebar_current: "docs-ibm-resource-resource-key"
description: |-
  Manages IBM Resource Key.
---

# ibm\_resource_key

Provides a resource key resource. This allows resource keys to be created, and deleted.

## Example Usage

```hcl
data "ibm_resource_instance" "resource_instance" {
  name = "myobjectsotrage"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "myobjectkey"
  role                 = "Viewer"
  resource_instance_id = "${data.ibm_resource_instance.resource_instance.id}"

  //User can increase timeouts 
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```

## Timeouts

ibm_resource_key provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Key.
* `delete` - (Default 10 minutes) Used for Deleting Key.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify a resource key.
* `role` - (Required, string) Name of the user role. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `parameters` - (Optional, map) Arbitrary parameters to pass. Must be a JSON object.
* `resource_instance_id` - (Optional, string) The id of the resource instance associated with the resource key.  
 **NOTE**: Conflicts with `resource_alias_id`.
* `resource_alias_id` - (Optional, string) The id of the resource alias associated with the resource key.  
 **NOTE**: Conflicts with `resource_instance_id`.
* `tags` - (Optional, array of strings) Tags associated with the resource key instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new resource key.
* `credentials` - The credentials associated with the key.
* `status` - Status of resource key.
