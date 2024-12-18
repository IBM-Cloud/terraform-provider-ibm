---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2"
description: |-
  Manages IBM Db2 SaaS instance.
---

# ibm_db2

Create or delete an IBM Db2 SaaS on IBM Cloud instance. The `ibmcloud_api_key` that are used by Terraform should grant IAM rights to create and modify IBM Cloud Db2 Databases and have access to the resource group the Db2 SaaS instance is associated with. For more information, see [documentation](https://cloud.ibm.com/docs/Db2onCloud?topic=Db2onCloud-getting-started) to manage Db2 SaaS instances.


Configuration of an Db2 SaaS resource requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the target Db2 SaaS `location/region`. If the Terraform configuration needs to deploy resources into multiple regions, provider alias can be used. For more information, see [Terraform provider configuration](https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances).


## Example usage
To find an example for provisioning and configuring a Db2 SaaS instance , see [here](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-db2).

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }
}

```

**provider.tf**
Please make sure to target right region in the provider block. If database is created in region other than `us-south` , please specify it in provider block.

```terraform
provider "ibm" {
  ibmcloud_api_key      = var.ibmcloud_api_key
}
```


## Timeouts
The following timeouts are defined for this resource.

* `Create` The creation of an instance is considered failed when no response is received for 720 minutes.
* `Delete` The deletion of an instance is considered failed when no response is received for 30 minutes.

Db2 SaaS create instance typically takes between 30 minutes to 45 minutes. Delete and update takes a minute. Provisioning time are unpredictable, if the apply fails due to a timeout, import the database resource once the create is completed.


## Argument reference
Review the argument reference that you can specify for your resource.


- `location` - (Required, String) The location where you want to deploy your instance. The location must match the `region` parameter that you specify in the `provider` block of your  Terraform configuration file. Currently, supported regions are `us-south`, `us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `mon01`, `br-sao`, `ca-tor`, `mil01`.
- `name` - (Required, String) A descriptive name that is used to identify the database instance. The name must not include spaces.
- `plan` - (Required, Forces new resource, String) The name of the service plan to use when provisioning.  Currently the only supported option is `performance`.
- `resource_group_id` - (Optional, Forces new resource, String)  The ID of the resource group where you want to create the instance. To retrieve this value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `service` - (Required, Forces new resource, String) The type of Cloud Db2 SaaS that you want to create. Only the following services are currently accepted: `dashdb-for-transactions` only.
- `service_endpoints` - (Required, String) Specify whether you want to enable the public, private, or both service endpoints. Supported values are `public`, `private`, or `public-and-private`.
- `tags` - (Optional, Array of Strings) A list of tags that you want to add to your instance.
- `high_availability` - (Optional, String) By default, it is `no`.if you want please change to `yes`
- `backup_location` - (Optional, String) Cross Regional backups can be stored across multiple regions in a zone. Regional backups are stored in only one specific region.
- `instance_type` - (Optional, String) The hosting infrastructure identifier.By Default `bx2.1x4` taken automatically. With this identifier, minimum resource configurations apply. Alternatively, setting the identifier to any of the following host sizes places your database on the specified host size with no other tenants.
          - `bx2.4x16`
          - `bx2.8x32`
          - `bx2.16x64`
          - `bx2.32x128`
          - `bx2.48x192`
          - `mx2.4x32`
          - `mx2.16x128`
          - `mx2.128x1024`
- `disk_encryption_instance_crn` - (Optional, String) Please ensure Databases for Db2 has been authorized to access the selected KMS instance.
- `disk_encryption_key_crn` - (Optional, String) Warning: deleting this key will result in the loss of all data stored in this Db2 instance.
- `oracle_compatibility` - (Optional, String) If you require Oracle compatibility, please choose this option(YES/NO).
- `subscription_id` - (Optional, String) ID which is required for subscription plans, for example: PerformanceSubscription.
- `parameters_json` - (Optional, JSON) Parameters to create Db2 SaaS instance. The value must be a JSON string.

  Nested scheme for `parameters_json`:
  - `backup_encryption_key_crn` -  (Optional, Forces new resource, String) The CRN of a key protect key, that you want to use for encrypting disk that holds deployment backups. A key protect CRN is in the format `crn:v1:<...>:key:`. `backup_encryption_key_crn` can be added only at the time of creation and no update support  are available.


## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The CRN of the database instance.
- `status` - (String) The status of the instance.
- `version` - (String) The database version.

## Import
The database instance can be imported by using the ID, that is formed from the CRN. To import the resource, you must specify the `region` parameter in the `provider` block of your  Terraform configuration file. If the region is not specified, `us-south` is used by default. An  Terraform refresh or apply fails, if the database instance is not in the same region as configured in the provider or its alias.

CRN is a 120 digit character string of the form -  `crn:v1:bluemix:public:dashdb-for-transactions:us-south:a/60970f92286548d8a64cbb45bce39bc1:deae06ff-3966-4534-bfa0-4b42281e7cef::`

**Syntax**

```
$ terraform import ibm_db2.my_db <crn>
```

**Example**

```
$ terraform import ibm_db2.my_db crn:v1:bluemix:public:dashdb-for-transactions:us-south:a/60970f92286548d8a64cbb45bce39bc1:deae06ff-3966-4534-bfa0-4b42281e7cef::
```

Import requires a minimal Terraform config file to allow importing.

```terraform
resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
}
```

Run `terraform state show ibm_db2.<your_database>` after import to retrieve the more values to be included in the resource config file. It does not export any more user IDs and passwords that are configured on the instance. These values must be retrieved from an alternative source.