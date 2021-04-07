---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : resource_instance"
description: |-
  Manages IBM Resource Instance.
---

# ibm\_resource_instance

Provides a Resource Instance resource. This allows Resource Instances to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_resource_instance" "resource_instance" {
  name              = "test"
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
```

## Timeouts

ibm_resource_instance provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Instance.
* `update` - (Default 10 minutes) Used for Updating Instance.
* `delete` - (Default 10 minutes) Used for Deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the resource instance.
* `service` - (Required,Forces new resource, string) The name of the service offering. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `plan` - (Required, string) The name of the plan type supported by service. You can retrieve the value by running the `ibmcloud catalog service <servicename>` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `location` - (Required,Forces new resource, string) Target location or environment to create the resource instance.
* `resource_group_id` - (Optional,Forces new resource,string) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.
* `tags` - (Optional, array of strings) Tags associated with the instance.
* `parameters` - (Optional,Forces new resource,map) Arbitrary parameters to create instance. The value must be a JSON object.
* `service_endpoints` - (Optional, string) Types of the service endpoints that can be set to a resource instance. Possible values are 'public', 'private', 'public-and-private'.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the new resource instance.
* `status` - Status of resource instance.
* `guid`- Guid of the resource instance.
* `dashboard_url`- The dashboard url of the new resource instance.
* `extensions` - The extended metadata as a map associated with the resource instance.
* `plan_history` - The plan history of the instance.
* `account_id` - An alpha-numeric value identifying the account ID.
* `resource_group_crn` - The long ID (full CRN) of the resource group.
* `resource_id` - The unique ID of the offering. This value is provided by and stored in the global catalog.
* `resource_plan_id` - The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
* `target_crn` - The full deployment CRN as defined in the global catalog. The Cloud Resource Name (CRN) of the deployment location where the instance is provisioned.
* `state` - The current state of the instance. For example, if the instance is deleted, it will return removed.
* `type` - The type of the instance, e.g. service_instance.
* `sub_type` - The sub-type of instance, e.g. cfaas .
* `allow_cleanup` - A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
* `locked` - A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
* `last_operation` - The status of the last operation requested on the instance.
* `resource_aliases_url` - The relative path to the resource aliases for the instance.
* `resource_bindings_url` - The relative path to the resource bindings for the instance.
* `resource_keys_url` - The relative path to the resource keys for the instance.
* `created_at` - The date when the instance was created.
* `created_by` - The subject who created the instance.
* `update_at` - The date when the instance was last updated.
* `update_by` - The subject who updated the instance.
* `deleted_at` - The date when the instance was deleted.
* `deleted_by` - The subject who deleted the instance.
* `scheduled_reclaim_at` - The date when the instance was scheduled for reclamation.
* `scheduled_reclaim_by` - The subject who initiated the instance reclamation.
* `restored_at` - The date when the instance under reclamation was restored.
* `restored_by` - The subject who restored the instance back from reclamation.

