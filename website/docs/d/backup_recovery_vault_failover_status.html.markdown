---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_vault_failover_status"
description: |-
  Get information about Get Batch Vault Failover Status
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_vault_failover_status

Provides a read-only data source to retrieve information about Get Batch Vault Failover Status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_vault_failover_status" "backup_recovery_vault_failover_status" {
	cloud_type = "ibm"
	vault_ids = [ 8 ]
	x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cloud_type` - (Required, String) Specifies the cloud environment type where the Backup and Recovery instance is used. Currently, only 'ibm' is supported for failovers.
  * Constraints: Allowable values are: `ibm`.
* `vault_ids` - (Required, List) Specifies the unique ids of the Backup and Recovery instances i.e. vaults for which the latest failover status is to be fetched.
* `x_ibm_tenant_id` - (Required, String) Id of the tenant accessing the cluster.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Get Batch Vault Failover Status.
* `error_message` - (String) Specifies the error message if the batch failover status retrieval failed.
* `failover_statuses` - (List) Array of failover statuses for the specified Backup and Recovery instances.
Nested schema for **failover_statuses**:
	* `status` - (List) Specifies the status of a vault Failover.
	Nested schema for **status**:
		* `end_time_usecs` - (Integer) Specifies the end time of the failover in microseconds since epoch.
		* `error_message` - (String) Specifies the error message if the vault failover failed.
		* `start_time_usecs` - (Integer) Specifies the start time of the failover in microseconds since epoch.
		* `status` - (String) Specifies the current status of the failover.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
		* `uid` - (String) Specifies the unique id of the failover.
	* `vault_id` - (Integer) Specifies the unique id of the Backup and Recovery instance.

