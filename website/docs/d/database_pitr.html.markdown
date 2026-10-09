---
layout: "ibm"
page_title: "IBM : ibm_database_point_in_time_recovery"
description: |-
  Get information about database_pitr
subcategory: "Cloud Databases"
---

# ibm_database_point_in_time_recovery

Provides a read-only data source for database_pitr. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

The data source reads the instance's plan and picks the Classic or Gen2 behavior from it.

## Example Usage

```hcl
data "ibm_database_point_in_time_recovery" "database_pitr" {
	deployment_id = data.ibm_database.database.id
}
```

### Gen2 restorable window

For a Gen2 instance, the data source reads the instance's restorable window: the range of times that a point-in-time restore of it can target, the periods inside that range that cannot be restored to, and the state of its change archiving. The values are read on every plan. While archiving is healthy, `latest_point_in_time_recovery_time` is the time of the read, unless the last entry of `unavailable_periods` has an empty `until`: it then stays at the millisecond before that entry's `from` until the source's next backup completes, so a restore to it leaves out every change after that time.

```hcl
data "ibm_database_point_in_time_recovery" "source" {
  deployment_id = ibm_database.postgres_gen2.id
}

resource "ibm_database" "postgres_gen2_restore" {
  name              = "my-postgres-gen2-restore"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service           = "databases-for-postgresql"
  service_endpoints = "private"

  point_in_time_recovery_deployment_id = ibm_database.postgres_gen2.id
  point_in_time_recovery_time          = data.ibm_database_point_in_time_recovery.source.latest_point_in_time_recovery_time

  timeouts {
    create = "180m"
  }
}
```

The restore arguments of `ibm_database` are applied only at creation, so later plans ignore the moving `latest_point_in_time_recovery_time`. `latest_point_in_time_recovery_time` is empty whenever `not_restorable_reason` is set: `archive_delayed` while archiving is delayed, and `archive_not_started` or `no_eligible_backup` for a source that has had no backup since it started archiving, such as one created minutes ago. `ibm_database` then refuses the restore at plan time, or during apply when the source itself has pending changes, because Terraform then reads the data source only during apply. Check `not_restorable_reason`, and retry once archiving is healthy or the source's next backup completes. `earliest_point_in_time_recovery_time` can move forward at any time as old backups expire, so a restore to it can be refused; choose a later time for a margin.

Once the restored instance exists, the data source still reads the source on every plan. Before you remove the source or the data source, set the restore's `point_in_time_recovery_deployment_id` and `point_in_time_recovery_time` to the literal values that `terraform state show ibm_database.postgres_gen2_restore` prints, or remove both arguments. Both are applied only at creation, so this edit plans no change. Then remove the data source.

Reading the Gen2 window needs the following:

* An IAM access token with a role that grants `deployment-backup.list` on the instance, such as Viewer. The window omits ranges that depend on a backup the caller cannot read.
* The URL of the instance's window in its Resource Controller extensions, at `dataservices.point_in_time_recovery_status.restorable_window_url`. If it is missing where point-in-time recovery is available, update the instance's parameters or plan and wait for the update to finish, which refreshes its extensions (a tag change does not), or ask IBM Cloud support to sync it.
* Network access, from where Terraform runs, to the host in `restorable_window_url`, which is a public endpoint.

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, Forces new resource, String) Deployment ID.

  **Classic:** The database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:us-south:a/<account_id>:<instance_id>::
  ```

  **Gen2:** The Gen2 database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:<region>:a/<account_id>:<instance_id>::
  ```
  The Gen2 deployment CRN can be retrieved from:
  - The `id` attribute of an `ibm_database` resource configured with a Gen2 plan.
  - The IBM Cloud UI under **Databases → your instance → Overview**.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `deployment_id` - The unique identifier of the database_pitr.
* `earliest_point_in_time_recovery_time` - (String) - The earliest point in time recovery. **Gen2:** The earliest time that a restore can target, in UTC with millisecond precision, for example `2026-09-20T09:30:01.000Z`. Empty when the instance cannot be restored.
* `latest_point_in_time_recovery_time` - (String) **Gen2 only.** The latest time that a restore can target when the data source was read, in the same format. Empty whenever `not_restorable_reason` is set. While the last entry of `unavailable_periods` is open, it stays at the millisecond before that entry's `from`.
* `retention_days` - (Integer) **Gen2 only.** The days of history the instance keeps for point-in-time recovery.
* `archiving_status` - (String) **Gen2 only.** The state of the instance's change archiving: `healthy`, `delayed` or `unknown`.
* `unavailable_periods` - (List) **Gen2 only.** Periods inside the window that cannot be restored to.
Nested scheme for **unavailable_periods**:
	* `from` - (String) The start of the period, rounded down to the millisecond, so `from` itself may still be restorable; treat it as unavailable.
	* `until` - (String) The first time after the period that can be restored to again. Empty while the period is still open.
* `not_restorable_reason` - (String) **Gen2 only.** Why the instance cannot be restored now: `archive_not_started`, `no_eligible_backup`, `archive_delayed` or `archive_gap`. Empty when it can be restored. The data source returns a warning whenever it is set.
