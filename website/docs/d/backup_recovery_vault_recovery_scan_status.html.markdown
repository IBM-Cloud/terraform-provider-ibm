---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_vault_recovery_scan_status"
description: |-
  Get information about Get Batch Vault Recovery Scan Status
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_vault_recovery_scan_status

Provides a read-only data source to retrieve information about Get Batch Vault Recovery Scan Status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_vault_recovery_scan_status" "backup_recovery_vault_recovery_scan_status" {
	cloud_type = "ibm"
	x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cloud_type` - (Required, String) Specifies the cloud environment type where the Backup and Recovery instance is used. Currently, only 'ibm' is supported for recover scans.
  * Constraints: Allowable values are: `ibm`.
* `vault_ids` - (Optional, List) Specifies the unique ids of the Backup and Recovery instances for which the latest recovery scan status is to be fetched.
* `x_ibm_tenant_id` - (Required, String) Id of the tenant accessing the cluster.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Get Batch Vault Recovery Scan Status.
* `error_message` - (String) Specifies the error message if the batch recovery scan status retrieval failed.
* `recovery_scan_statuses` - (List) Array of recovery scan statuses for the specified Backup and Recovery instances.
Nested schema for **recovery_scan_statuses**:
	* `status` - (List) Specifies the status of a Recovery Scan.
	Nested schema for **status**:
		* `end_time_usecs` - (Integer) Specifies the end time of the recovery scan in microseconds since epoch.
		* `error_message` - (String) Specifies the error message if the recovery scan failed.
		* `start_time_usecs` - (Integer) Specifies the start time of the recovery scan in microseconds since epoch.
		* `status` - (String) Specifies the current status of the recovery scan.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
		* `uid` - (String) Specifies the unique id of the recovery scan.
	* `vault_id` - (Integer) Specifies the unique id of the Backup and Recovery instance.

