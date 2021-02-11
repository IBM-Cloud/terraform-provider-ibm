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
  resource_instance_id = data.ibm_resource_instance.resource_instance.id

  //User can increase timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```

**Note** The current `ibm_resource_key` resource doesn't have support for service_id argument but the service_id can be passesd as one of the parameter.

## Example Usage with serviceID 

```hcl
data "ibm_resource_instance" "resource_instance" {
  name = "myobjectsotrage"
}

resource "ibm_iam_service_id" "serviceID" {
  name        = "test"
  description = "New ServiceID"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "myobjectkey"
  role                 = "Viewer"
  resource_instance_id = data.ibm_resource_instance.resource_instance.id
  parameters = {
    "serviceid_crn" = ibm_iam_service_id.serviceID.crn
  }

  //User can increase timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```
## Example Usage with HMAC 

```hcl
data "ibm_resource_group" "group" {
    name ="Default"
}
resource "ibm_resource_instance" "resource_instance" {
  name              = "test-21"
  service           = "cloud-object-storage"
  plan              = "lite"
  location          = "global"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]
  
  //User can increase timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
resource "ibm_resource_key" "resourceKey" {
  name                 = "my-cos-bucket-xx-key"
  resource_instance_id = ibm_resource_instance.resource_instance.id
  parameters           = { "HMAC" = true }
  role                 = "Manager"
}

```

## Timeouts

ibm_resource_key provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Key.
* `delete` - (Default 10 minutes) Used for Deleting Key.

## Argument Reference

The following arguments are supported:

* `name` - (Required,Forces new resource, string) A descriptive name used to identify a resource key.
* `role` - (Required,Forces new resource, string) Name of the user role. Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.
* `parameters` - (Optional,Forces new resource, map) Arbitrary parameters to pass. Must be a JSON object.
* `resource_instance_id` - (Optional, Forces new resource,string) The id of the resource instance associated with the resource key.  
 **NOTE**: Conflicts with `resource_alias_id`.
* `resource_alias_id` - (Optional,Forces new resource,string) The id of the resource alias associated with the resource key.  
 **NOTE**: Conflicts with `resource_instance_id`.
* `tags` - (Optional, array of strings) Tags associated with the resource key instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new resource key.
* `credentials` - The credentials associated with the key.
* `status` - Status of resource key.
