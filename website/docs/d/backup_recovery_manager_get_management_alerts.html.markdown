---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_management_alerts"
description: |-
  Get information about backup_recovery_manager_get_management_alerts
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_management_alerts

Provides a read-only data source to retrieve information about backup_recovery_manager_get_management_alerts. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_management_alerts" "backup_recovery_manager_get_management_alerts" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `alert_category_list` - (Optional, List) Filter by list of alert categories.
  * Constraints: Allowable list items are: `kDisk`, `kNode`, `kCluster`, `kChassis`, `kPowerSupply`, `kCPU`, `kMemory`, `kTemperature`, `kFan`, `kNIC`, `kFirmware`, `kNodeHealth`, `kOperatingSystem`, `kDataPath`, `kMetadata`, `kIndexing`, `kHelios`, `kAppMarketPlace`, `kSystemService`, `kLicense`, `kSecurity`, `kUpgrade`, `kClusterManagement`, `kAuditLog`, `kNetworking`, `kConfiguration`, `kStorageUsage`, `kFaultTolerance`, `kBackupRestore`, `kArchivalRestore`, `kRemoteReplication`, `kQuota`, `kCDP`, `kViewFailover`, `kDisasterRecovery`, `kStorageDevice`, `kStoragePool`, `kGeneralSoftwareFailure`, `kAgent`.
* `alert_id_list` - (Optional, List) Filter by list of alert ids.
* `alert_name` - (Optional, String) Specifies name of alert to filter alerts by.
* `alert_property_key_list` - (Optional, List) Specifies list of the alert property keys to query.
* `alert_property_value_list` - (Optional, List) Specifies list of the alert property value, multiple values for one key should be joined by '|'.
* `alert_severity_list` - (Optional, List) Filter by list of alert severity types.
  * Constraints: Allowable list items are: `kCritical`, `kWarning`, `kInfo`.
* `alert_state_list` - (Optional, List) Filter by list of alert states.
* `alert_type_bucket_list` - (Optional, List) Filter by list of alert type buckets.
  * Constraints: Allowable list items are: `kHardware`, `kSoftware`, `kDataService`, `kMaintenance`.
* `alert_type_list` - (Optional, List) Filter by list of alert types.
* `cluster_identifiers` - (Optional, List) Filter by list of cluster ids.
* `end_date_usecs` - (Optional, Integer) Specifies the end time of the alerts to be returned. All the alerts returned are raised before the specified end time. This value should be in Unix timestamp epoch in microseconds.
* `max_alerts` - (Optional, Integer) Specifies maximum number of alerts to return.
* `region_ids` - (Optional, List) Filter by list of region ids.
* `service_instance_ids` - (Optional, List) Specifies services instance ids to filter alerts for IBM customers.
* `start_date_usecs` - (Optional, Integer) Specifies the start time of the alerts to be returned. All the alerts returned are raised after the specified start time. This value should be in Unix timestamp epoch in microseconds.
* `tenant_ids` - (Optional, List) Filter by tenant ids.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_management_alerts.
* `alerts_list` - (List) Specifies the list of alerts.
Nested schema for **alerts_list**:
	* `alert_category` - (String) Specifies the alert category.
	  * Constraints: Allowable values are: `kDisk`, `kNode`, `kCluster`, `kChassis`, `kPowerSupply`, `kCPU`, `kMemory`, `kTemperature`, `kFan`, `kNIC`, `kFirmware`, `kNodeHealth`, `kOperatingSystem`, `kDataPath`, `kMetadata`, `kIndexing`, `kHelios`, `kAppMarketPlace`, `kSystemService`, `kLicense`, `kSecurity`, `kUpgrade`, `kClusterManagement`, `kAuditLog`, `kNetworking`, `kConfiguration`, `kStorageUsage`, `kFaultTolerance`, `kBackupRestore`, `kArchivalRestore`, `kRemoteReplication`, `kQuota`, `kCDP`, `kViewFailover`, `kDisasterRecovery`, `kStorageDevice`, `kStoragePool`, `kGeneralSoftwareFailure`, `kAgent`.
	* `alert_code` - (String) Specifies a unique code that categorizes the Alert, for example: CE00200014, where CE stands for IBM Error, the alert state next 3 digits is the id of the Alert Category (e.g. 002 for 'kNode') and the last 5 digits is the id of the Alert Type (e.g. 00014 for 'kNodeHighCpuUsage').
	* `alert_document` - (List) Specifies the fields of alert document.
	Nested schema for **alert_document**:
		* `alert_cause` - (String) Specifies the cause of alert.
		* `alert_description` - (String) Specifies the description of alert.
		* `alert_help_text` - (String) Specifies the help text for alert.
		* `alert_name` - (String) Specifies the name of alert.
		* `alert_summary` - (String) Short description for the alert.
	* `alert_state` - (String) Specifies the alert state.
	  * Constraints: Allowable values are: `kResolved`, `kOpen`, `kNote`, `kSuppressed`.
	* `alert_type` - (Integer) Specifies the alert type.
	* `alert_type_bucket` - (String) Specifies the Alert type bucket.
	  * Constraints: Allowable values are: `kHardware`, `kSoftware`, `kDataService`, `kMaintenance`.
	* `cluster_id` - (Integer) Id of the cluster which the alert is associated.
	* `cluster_name` - (String) Specifies the name of cluster which alert is raised from.
	* `dedup_count` - (Integer) Specifies the dedup count of alert.
	* `dedup_timestamps` - (List) Specifies Unix epoch Timestamps (in microseconds) for the last 25 occurrences of duplicated Alerts that are stored with the original/primary Alert. Alerts are grouped into one Alert if the Alerts are the same type, are reporting on the same Object and occur within one hour. 'dedupCount' always reports the total count of duplicated Alerts even if there are more than 25 occurrences. For example, if there are 100 occurrences of this Alert, dedupTimestamps stores the timestamps of the last 25 occurrences and dedupCount equals 100.
	* `first_timestamp_usecs` - (Integer) SpeSpecifies Unix epoch Timestamp (in microseconds) of the first occurrence of the Alert.
	* `id` - (String) Specifies unique id of the alert.
	* `latest_timestamp_usecs` - (Integer) SpeSpecifies Unix epoch Timestamp (in microseconds) of the most recent occurrence of the Alert.
	* `property_list` - (List) List of property key and values associated with alert.
	Nested schema for **property_list**:
		* `key` - (String) Key of the Label.
		* `value` - (String) Value of the Label, multiple values should be joined by '|'.
	* `region_id` - (String) Specifies the region id of the alert.
	* `resolution_id_string` - (String) Specifies the resolution id of the alert if its resolved.
	* `service_instance_id` - (String) Id of the serrvice instance which the alert is associated.
	* `severity` - (String) Specifies the alert severity.
	  * Constraints: Allowable values are: `kCritical`, `kWarning`, `kInfo`.
	* `vaults` - (List) Specifies information about vaults where source object associated with alert is vaulted. This could be empty if alert is not related to any source object or it is not vaulted.
	Nested schema for **vaults**:
		* `global_vault_id` - (String) Specifies Global vault id.
		* `region_id` - (String) Specifies id of region where vault resides.
		* `region_name` - (String) Specifies name of region where vault resides.
		* `vault_name` - (String) Specifies name of vault.

