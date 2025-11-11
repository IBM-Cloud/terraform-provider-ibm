---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_alerts_summary"
description: |-
  Get information about backup_recovery_manager_get_alerts_summary
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_alerts_summary

Provides a read-only data source to retrieve information about a backup_recovery_manager_get_alerts_summary. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_alerts_summary" "backup_recovery_manager_get_alerts_summary" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `end_time_usecs` - (Optional, Integer) Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time.
* `include_tenants` - (Optional, Boolean) IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user's organization should be used to compute summary.
* `start_time_usecs` - (Optional, Integer) Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day.
* `states_list` - (Optional, List) Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary.
  * Constraints: Allowable list items are: `kResolved`, `kOpen`, `kNote`, `kSuppressed`.
* `tenant_ids` - (Optional, List) TenantIds contains ids of the tenants for which alerts are to be used to compute summary.
* `x_scope_identifier` - (Optional, String) This field uniquely represents a service        instance. Please specify the values as "service-instance-id: <value>".

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_alerts_summary.
* `alerts_summary` - (List) Specifies a list of alerts summary grouped by category.
Nested schema for **alerts_summary**:
	* `category` - (String) Category of alerts by which summary is grouped.
	  * Constraints: Allowable values are: `kDisk`, `kNode`, `kCluster`, `kChassis`, `kPowerSupply`, `kCPU`, `kMemory`, `kTemperature`, `kFan`, `kNIC`, `kFirmware`, `kNodeHealth`, `kOperatingSystem`, `kDataPath`, `kMetadata`, `kIndexing`, `kHelios`, `kAppMarketPlace`, `kSystemService`, `kLicense`, `kSecurity`, `kUpgrade`, `kClusterManagement`, `kAuditLog`, `kNetworking`, `kConfiguration`, `kStorageUsage`, `kFaultTolerance`, `kBackupRestore`, `kArchivalRestore`, `kRemoteReplication`, `kQuota`, `kCDP`, `kViewFailover`, `kDisasterRecovery`, `kStorageDevice`, `kStoragePool`, `kGeneralSoftwareFailure`, `kAgent`.
	* `critical_count` - (Integer) Specifies count of critical alerts.
	* `info_count` - (Integer) Specifies count of info alerts.
	* `total_count` - (Integer) Specifies count of total alerts.
	* `type` - (String) Type/bucket that this alert category belongs to.
	* `warning_count` - (Integer) Specifies count of warning alerts.

