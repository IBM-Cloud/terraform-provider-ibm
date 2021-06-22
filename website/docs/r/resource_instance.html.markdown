---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : resource_instance"
description: |-
  Manages IBM resource instance.
---

# ibm_resource_instance
Create, update, or delete an IAM enabled service instance. For more information, about resource instance, see [assigning access to resources](https://cloud.ibm.com/docs/account?topic=account-access-getstarted). 

## Example usage

```terraform
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
### Example to provision a Hyper Protect DBaaS PostgreSQL Service
The following example enables you to create a service instance of IBM Cloud Hyper Protect DBaaS for PostgreSQL. For detailed argument reference, see the tables in the [Hyper Protect DBaaS for PostgreSQL documentation](/docs/hyper-protect-dbaas-for-postgresql?topic=hyper-protect-dbaas-for-postgresql-create-service#cli-create-service), or the [Hyper Protect DBaaS for MongoDB documentation](/docs/hyper-protect-dbaas-for-mongodb?topic=hyper-protect-dbaas-for-monbgodb-create-service#cli-create-service) to create MongoDB service instances.

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_resource_instance" "myhpdbcluster" {
  name = "0001-postgresql"
  service = "hyperp-dbaas-postgresql"
  plan = "postgresql-free"
  location = "us-south"
  resource_group_id = data.ibm_resource_group.group.id

  //User can increase timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }

  parameters = {
    name: "cluster01",
    admin_name: "admin",
    password: "Hyperprotectdbaas0001"
    confirm_password: "Hyperprotectdbaas0001",
    db_version: "10"
  }
}
```

## Timeouts

The `ibm_resource_instance` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for Creating Instance.
- **update** - (Default 10 minutes) Used for Updating Instance.
- **delete** - (Default 10 minutes) Used for Deleting Instance.

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Required, Forces new resource, String) Target location or environment to create the resource instance.
- `parameters` (Optional, Forces new resource, Map) Arbitrary parameters to create instance. The value must be a JSON object.
- `plan` - (Required, String) The name of the plan type supported by service. You can retrieve the value by running the `ibmcloud catalog service <servicename>` command.
- `name` - (Required, String) A descriptive name used to identify the resource instance.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.
- `tags` (Optional, Array of Strings) Tags associated with the instance.
- `service` - (Required, Forces new resource, String) The name of the service offering. You can retrieve the value by installing the `catalogs-management` command line plug-in and running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command. For more information, about IBM Cloud catalog service marketplace, refer [IBM Cloud catalog service marketplace](https://cloud.ibm.com/docs/cli?topic=cli-ibmcloud_catalog#ibmcloud_catalog_service_marketplace).
- `service_endpoints` - (Optional, String) Types of the service endpoints that can be set to a resource instance. Possible values are `public`, `private`, `public-and-private`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account_id` - (String) An alpha-numeric value identifying the account ID.
- `allow_cleanup` - (String) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
- `created_at` - (Timestamp) The date when the instance  created.
- `created_by` - (String) The subject who created the instance.
- `dashboard_url` - (String) The dashboard URL of the new resource instance.
- `deleted_at` - (Timestamp) The date when the instance was deleted.
- `deleted_by` - (String) The subject who deleted the instance.
- `extensions` - (String) The extended metadata as a map associated with the resource instance.
- `guid` - (String) The GUID of the resource instance.
- `id` - (String) The unique identifier of the new resource instance.
- `last_operation` - (String) The status of the last operation requested on the instance.
- `locked` - (String) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
- `plan_history` - (String) The plan history of the instance.
- `resource_group_crn` - (String) The long ID (full CRN) of the resource group.
- `resource_id` - (String) The unique ID of the offering. This value is provided by and stored in the global catalog.
- `resource_plan_id` - (String) The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
- `resource_aliases_url` - (String) The relative path to the resource aliases for the instance.
- `resource_bindings_url` - (String) The relative path to the resource bindings for the instance.
- `resource_keys_url` - (String)  The relative path to the resource keys for the instance.
- `restored_at` - (Timestamp) The date when the instance under reclamation restored.
- `restored_by` - (String) The subject who restored the instance back from reclamation.
- `status` - (String) The status of resource instance.
- `sub_type` - (String) The sub-type of instance, for example, `cfaas`.
- `state` - (String) The current state of the instance. For example, if the instance is deleted, it will return removed.
- `scheduled_reclaim_at` - (Timestamp) The date when the instance scheduled for reclamation.
- `scheduled_reclaim_by` - (String) The subject who initiated the instance reclamation.
- `target_crn` - (String) The full deployment CRN as defined in the global catalog. The Cloud Resource Name (CRN) of the deployment location where the instance is provisioned.
- `type` - (String) The type of the instance. For example, `service_instance`.
- `update_at` - (Timestamp) The date when the instance last updated.
- `update_by` - (String) The subject who updated the instance.
