---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : resource_key"
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

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the new resource key.
* `credentials` - The credentials associated with the key.
* `status` - Status of resource key.
* `guid` - This GUID is a unique internal identifier managed by the resource controller that corresponds to the key.
* `crn` - The full Cloud Resource Name (CRN) associated with the key.
* `url` - When you created a new key, a relative URL path is created identifying the location of the key.
* `account_id` - An alpha-numeric value identifying the account ID.
* `resource_group_id` - The short ID of the resource group.
* `source_crn` - The CRN of resource instance or alias associated to the key.
* `state` - The state of the key.
* `iam_compatible` - Specifies whether the key’s credentials support IAM.
* `resource_instance_url` - The relative path to the resource.
* `created_at` - The date when the key was created.
* `updated_at` - The date when the key was last updated.
* `deleted_at` - The date when the key was deleted.
* `created_by` - The subject who created the key.
* `updated_by` - The subject who updated the key.
* `deleted_by` - The subject who deleted the key.
