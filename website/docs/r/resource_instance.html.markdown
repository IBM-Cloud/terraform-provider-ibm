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

## Example to provision a Cloud Logs instance
```
resource "ibm_resource_instance" "logs_instance" {
  name     = "logs-instance"
  service  = "logs"
  plan     = "standard"
  location = "eu-de"
  parameters = {
    retention_period        = "14"
    logs_bucket_crn         = "crn:v1:bluemix:public:cloud-object-storage:global:a/4448261269a14562b839e0a3019ed980:f8b3176e-af8e-4e14-a2f9-7f82634e7f0b:bucket:logs-bucket"
    logs_bucket_endpoint    = "s3.direct.eu-de.cloud-object-storage.appdomain.cloud"
    metrics_bucket_crn      = "crn:v1:bluemix:public:cloud-object-storage:global:a/4448261269a14562b839e0a3019ed980:f8b3176e-af8e-4e14-a2f9-7f82634e7f0b:bucket:metrics-bucket"
    metrics_bucket_endpoint = "s3.direct.eu-de.cloud-object-storage.appdomain.cloud"
  }
}
```

### Example to provision a Hyper Protect DBaaS service instance
The following example enables you to create a service instance of IBM Cloud Hyper Protect DBaaS for MongoDB. For detailed argument reference, see the tables in the [Hyper Protect DBaaS for MongoDB documentation](https://cloud.ibm.com/docs/hyper-protect-dbaas-for-mongodb?topic=hyper-protect-dbaas-for-mongodb-create-service&interface=cli#cli-create-service), or the [Hyper Protect DBaaS for PostgreSQL documentation](https://cloud.ibm.com/docs/hyper-protect-dbaas-for-postgresql?topic=hyper-protect-dbaas-for-postgresql-create-service&interface=cli#cli-create-service) to create PostgreSQL service instances.

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}
resource "ibm_resource_instance" "myhpdbcluster" {
  name = "0001-mongodb"
  service = "hyperp-dbaas-mongodb"
  plan = "mongodb-flexible"
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
   db_version: "4.4"
   cpu: "1",
   kms_instance: "crn:v1:staging:public:kms:us-south:a/23a24a3e3fe7a115473f07be1c44bdb5:9eeb285a-88e4-4378-b7cf-dbdcd97b5e4e::",
   kms_key: "ee742940-d87c-48de-abc9-d26a6184ba5a",
   memory: "2gib",
   private_endpoint_type: "vpe",
   service-endpoints: "public-and-private",
   storage: "5gib"
 }
}
```

### Example to provision a Watson Query service instance
The following example enables you to create a service instance of IBM Watson Query. For detailed argument reference, see the tables in the [Watson Query documentation](https://cloud.ibm.com/docs/data-virtualization?topic=data-virtualization-provisioning).

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}
resource "ibm_resource_instance" "wq_instance_1" {
  name              = "terraform-integration-1"
  service           = "data-virtualization"
  plan              = "data-virtualization-enterprise" # "data-virtualization-enterprise-dev","data-virtualization-enterprise-preprod","data-virtualization-enterprise-dev-stable"
  location          = "us-south" # "eu-gb", "eu-de", "jp-tok"
  resource_group_id = data.ibm_resource_group.group.id

  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }

}
```


### Example to provision a Analytics Engine using parameters_json argument
```terraform
resource "ibm_resource_instance" "instance" {
  name            = "MyServiceInstance"
  plan            = "standard-hourly"
  location        = "us-south"
  service         = "ibmanalyticsengine"
  parameters_json = <<PARAMETERS_JSON
    {
      "num_compute_nodes": "1",
      "hardware_config": "default",
      "software_package": "ae-1.2-hadoop-spark",
      "autoscale_policy": {
      "task_nodes": {
        "num_min_nodes": 1,
        "num_max_nodes": 10,
        "scaleup_rule": {
          "sustained_demand_period_minutes": "10",
          "percentage_of_demand": "50"
        },
        "scaledown_rule": {
          "sustained_excess_period_minutes": "20",
          "percentage_of_excess": "25"
        }
      }
    }
  }
    PARAMETERS_JSON
  tags = [
    "my-tag"
  ]
  timeouts {
    create = "30m"
    update = "15m"
    delete = "15m"
  }
}
```

### Example to provision an OpenPages service instance
The following example enables you to create a service instance of OpenPages. 

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}
resource "ibm_resource_instance" "openpages_instance" {
  name              = "openpages-instance-1"
  service           = "openpages"
  plan              = "essentials"
  location          = "global"
  resource_group_id = data.ibm_resource_group.default_group.id
  parameters_json   = <<EOF
    {
      "aws_region": "us-east-1",
      "baseCurrency": "USD",
      "selectedSolutions": ["ORM"]
    }
  EOF

  timeouts {
    create = "200m"
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
- `parameters` (Optional, Map) Arbitrary parameters to create instance. The value must be a JSON object. Conflicts with `parameters_json`.
- `parameters_json` (Optional,String) Arbitrary parameters to create instance. The value must be a JSON string. Conflicts with `parameters`.
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
- `resource_aliases_url` - (String, Deprecated) The relative path to the resource aliases for the instance. 
- `resource_bindings_url` - (String, Deprecated) The relative path to the resource bindings for the instance.
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
- `onetime_credentials` - (Bool) A boolean that dictates if the onetime_credentials is true or false.
