---
layout: "ibm"
page_title: "IBM : ibm_database_backup"
description: |-
  Get information about Backup
subcategory: "Cloud Databases"
---

# ibm_database_backup

Provides a read-only data source for Backup. Supports both Classic and Gen2 (Independent Backup) instances. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

### Classic

```hcl
data "ibm_database_backup" "database_backup" {
  backup_id = "<backup_crn>"
}
```

### Gen2

For Gen2 instances, the `backup_id` is an Independent Backup CRN with service name `databases-independent-backups`. This can be obtained from the `backup_id` field of the `ibm_database_backups` data source.

```hcl
data "ibm_database_backup" "database_backup" {
  backup_id = "crn:v1:bluemix:public:databases-independent-backups:<region>:a/<account_id>:<backup_id>::"
}

```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `backup_id` - (Required, Forces new resource, String) Backup ID.

  **Classic:** The backup CRN returned by the ICD API, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:us-south:a/<account_id>:<instance_id>:backup:<backup_id>
  ```

  **Gen2:** The Independent Backup CRN, identifiable by the service name `databases-independent-backups`, for example:
  ```
  crn:v1:bluemix:public:databases-independent-backups:<region>:a/<account_id>:<backup_id>::
  ```
  The Gen2 backup CRN can be retrieved from:
  - The `backup_id` field in the `ibm_database_backups` data source (when `deployment_id` is a Gen2 instance CRN).
  - The IBM Cloud UI under **Databases → your instance → Backups**.
  - The IBM Cloud CLI: `ibmcloud cdb backups <instance_name_or_crn>`.

  **Note:** Passing a Classic (coupled) backup CRN for a Gen2 instance returns an error. Use the Independent Backup CRN instead.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `backup_id` - The unique identifier of the Backup.
* `created_at` - (Optional, String) Date and time when this backup was created.

* `deployment_id` - (Optional, String) ID of the deployment this backup relates to.

* `download_link` - (Optional, String) URI which is currently available for file downloading.

* `is_downloadable` - (Optional, Boolean) Is this backup available to download?.

* `is_restorable` - (Optional, Boolean) Can this backup be used to restore an instance?.

* `status` - (Optional, String) The status of this backup.
  * Constraints: Allowable values are: `running`, `completed`, `failed`.

* `type` - (Optional, String) The type of backup.
  * Constraints: Allowable values are: `scheduled`, `on_demand`.

### Gen2 Independent Backups and S2S Authorization

Gen2 database instances support **Independent Backups** — automated backups managed independently of the database instance lifecycle. When a Gen2 instance is configured to use Independent Backups, it requires a service-to-service (S2S) IAM authorization between the database service and the backup storage.

If this authorization is missing or incomplete, Terraform emits a **non-blocking warning** during `plan` or `apply`:

```
╷
│ Warning: Database backup authorization required
│
│   with data.ibm_database_backup.<name>,
│
│ This database uses Independent Backups.
│ Existing backups remain available for 30 days from their creation date.
│ Backup creation and management are unavailable until the required service authorization is completed.
│
│ Complete the required service authorization to enable backup operations.
╵
```

The warning appears only when the instance has Independent Backups configured **and** the required S2S authorizations (`independent_backups` and `resource_group`) are not both `true`. It is suppressed for Classic plans and Gen2 instances not enrolled in Independent Backups.

To resolve the warning, create the required IAM service-to-service authorization between the database service and `databases-independent-backups`. Once both authorizations are in place, the warning will no longer appear.
