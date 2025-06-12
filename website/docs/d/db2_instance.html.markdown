---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2"
description: |-
  Get Information about IBM Db2 SaaS instance.
---

# ibm_db2

Retrieve information about an existing [IBM Db2 SaaS Instance](https://cloud.ibm.com/docs/Db2onCloud).

**Note**
Configuration of an IBM Db2 SaaS on IBM Cloud Instance `data_source` requires that the `region` parameter is set for the IBM provider in the `provider.tf`. The region must be the same as the `location` that the IBM Cloud Databases instance is deployed into.A `terraform refresh` of the `data_source` fails if the region and the location differ.

## Example usage
The following example retrieves information about the `db2_instance` instance in `us-south`.

```terraform
data "ibm_db2" "db2_instance" {
  name              = "<your_database_name>"
  resource_group_id = data.ibm_resource_group.group.id
  location          = "us-south"
  service           = "dashdb-for-transactions"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the IBM Db2 SaaS on IBM Cloud instance. IBM Cloud does not enforce that service names are unique and it is possible that duplicate service names exist. The first located service instance is used by  Terraform. The name must not include spaces.
- `location` - (Optional, String) The location where the IBM Db2 SaaS on IBM Cloud instance is deployed into.
- `resource_group_id`- (Optional, String) The ID of the resource group where the IBM Db2 SaaS on IBM Cloud instance is deployed into. The default is `default`.
- `service` - (Optional, String) The service type of the instance. To retrieve this value, run `ibmcloud catalog service-marketplace` or `ibmcloud catalog search`.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your data source is created. 

- `guid` - (String) The unique identifier of the IBM Db2 SaaS on IBM Cloud instance.
- `resource_crn` - (String) The unique identifier(CRN) of the IBM Db2 SaaS on IBM Cloud instance.
- `plan` - (String)  The service plan of the IBM Db2 SaaS on IBM Cloud instance.
- `location` - (String)  The location where the IBM Db2 SaaS on IBM Cloud instance is deployed into.
- `status` - (String)  The status of the IBM Db2 SaaS on IBM Cloud instance.
- `version` - (String) The database version.
- `platform_options`-  (String) The CRN of key protect key.
   
   Nested scheme for `platform_options`:
   - `disk_encryption_key_crn`-  (String) The CRN of disk encryption key.
   - `backup_encryption_key_crn`-  (String) The CRN of backup encryption key.
   - `subscription_id` - (String) ID which is required for subscription plans, for example: PerformanceSubscription.

