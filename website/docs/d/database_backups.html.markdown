---
layout: "ibm"
page_title: "IBM : ibm_database_backups"
description: |-
  Get information about Backups
subcategory: "Cloud Databases"
---

# ibm_database_backups

Provides a read-only data source for Backups. Supports both Classic and Gen2 (Independent Backup) instances. For more information, refer to [IBM Cloud Databases Gen2 Independent Backups](https://cloud.ibm.com/docs/cloud-databases-gen2?topic=cloud-databases-gen2-independent-backups&interface=ui). You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

### Classic

```hcl
data "ibm_database_backups" "database_backups" {
	deployment_id = "<crn>"
}
```

### Gen2

For Gen2 instances, the `deployment_id` is a Gen2 database instance CRN. The data source lists all Independent Backups associated with that instance.

```hcl
data "ibm_database_backups" "database_backups" {
	deployment_id = "crn:v1:bluemix:public:databases-for-mysql:<region>:a/<account_id>:<instance_id>::"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, String) ID of the deployment this backup relates to.

  **Classic:** The database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:us-south:a/<account_id>:<instance_id>::
  ```

  **Gen2:** The Gen2 database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-mysql:<region>:a/<account_id>:<instance_id>::
  ```
  The Gen2 deployment CRN can be retrieved from:
  - The `id` attribute of an `ibm_database` resource configured with a Gen2 plan.
  - The IBM Cloud UI under **Databases → your instance → Overview**.
  - The IBM Cloud CLI: `ibmcloud resource service-instance <instance_name> --output json | jq -r '.[0].crn'`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `backups` - (Optional, List) An array of backups.
Nested scheme for **backups**:
	* `created_at` - (Optional, String) Date and time when this backup was created.
	* `deployment_id` - (Optional, String) ID of the deployment this backup relates to.
	* `download_link` - (Optional, String) URI which is currently available for file downloading.
	* `backup_id` - (Optional, String) ID of this backup.
	* `is_downloadable` - (Optional, Boolean) Is this backup available to download?.
	* `is_restorable` - (Optional, Boolean) Can this backup be used to restore an instance?.
	* `status` - (Optional, String) The status of this backup.
	  * Constraints: Allowable values are: `running`, `completed`, `failed`.
	* `type` - (Optional, String) The type of backup.
	  * Constraints: Allowable values are: `scheduled`, `on_demand`.

