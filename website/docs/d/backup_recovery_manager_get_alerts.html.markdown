---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_alerts"
description: |-
  Get information about backup_recovery_manager_get_alerts
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_alerts

Provides a read-only data source to retrieve information about backup_recovery_manager_get_alerts. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_alerts" "backup_recovery_manager_get_alerts" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `alert_categories` - (Optional, List) Filter by list of alert categories.
  * Constraints: Allowable list items are: `kDisk`, `kNode`, `kCluster`, `kChassis`, `kPowerSupply`, `kCPU`, `kMemory`, `kTemperature`, `kFan`, `kNIC`, `kFirmware`, `kNodeHealth`, `kOperatingSystem`, `kDataPath`, `kMetadata`, `kIndexing`, `kHelios`, `kAppMarketPlace`, `kSystemService`, `kLicense`, `kSecurity`, `kUpgrade`, `kClusterManagement`, `kAuditLog`, `kNetworking`, `kConfiguration`, `kStorageUsage`, `kFaultTolerance`, `kBackupRestore`, `kArchivalRestore`, `kRemoteReplication`, `kQuota`, `kCDP`, `kViewFailover`, `kDisasterRecovery`, `kStorageDevice`, `kStoragePool`, `kGeneralSoftwareFailure`, `kAgent`.
* `alert_ids` - (Optional, List) Filter by list of alert ids.
* `alert_name` - (Optional, String) Specifies name of alert to filter alerts by.
* `alert_severities` - (Optional, List) Filter by list of alert severity types.
  * Constraints: Allowable list items are: `kCritical`, `kWarning`, `kInfo`.
* `alert_states` - (Optional, List) Filter by list of alert states.
  * Constraints: Allowable list items are: `kResolved`, `kOpen`, `kNote`, `kSuppressed`.
* `alert_type_buckets` - (Optional, List) Filter by list of alert type buckets.
  * Constraints: Allowable list items are: `kHardware`, `kSoftware`, `kDataService`, `kMaintenance`.
* `alert_types` - (Optional, List) Filter by list of alert types.
* `all_under_hierarchy` - (Optional, Boolean) Filter by objects of all the tenants under the hierarchy of the logged in user's organization.
* `end_time_usecs` - (Optional, Integer) Specifies end time Unix epoch time in microseconds to filter alerts by.
* `max_alerts` - (Optional, Integer) Specifies maximum number of alerts to return.The default value is 100 and maximum allowed value is 1000.
* `property_key` - (Optional, String) Specifies name of the property to filter alerts by.
* `property_value` - (Optional, String) Specifies value of the property to filter alerts by.
* `resolution_ids` - (Optional, List) Specifies alert resolution ids to filter alerts by.
* `start_time_usecs` - (Optional, Integer) Specifies start time Unix epoch time in microseconds to filter alerts by.
* `tenant_ids` - (Optional, List) Filter by tenant ids.
* `x_scope_identifier` - (Optional, String) This field uniquely represents a service        instance. Please specify the values as "service-instance-id: <value>".

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_alerts.
* `alerts` - (List) Specifies the list of alerts.
Nested schema for **alerts**:
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
	* `event_source` - (String) Specifies source where the event occurred.
	* `first_timestamp_usecs` - (Integer) Specifies Unix epoch Timestamp (in microseconds) of the first occurrence of the Alert.
	* `id` - (String) Specifies unique id of the alert.
	* `label_ids` - (List) Specifies the labels for which this alert has been raised.
	* `latest_timestamp_usecs` - (Integer) Specifies Unix epoch Timestamp (in microseconds) of the most recent occurrence of the Alert.
	* `property_list` - (List) List of property key and values associated with alert.
	Nested schema for **property_list**:
		* `key` - (String) Key of the Label.
		* `value` - (String) Value of the Label, multiple values should be joined by '|'.
	* `region_id` - (String) Specifies the region id of the alert.
	* `resolution_details` - (List) Specifies information about the Alert Resolution.
	Nested schema for **resolution_details**:
		* `resolution_details` - (String) Specifies detailed notes about the Resolution.
		* `resolution_id` - (Integer) Specifies the unique resolution id assigned in management console.
		* `resolution_summary` - (String) Specifies short description about the Resolution.
		* `timestamp_usecs` - (Integer) Specifies unix epoch timestamp (in microseconds) when the Alert was resolved.
		* `user_name` - (String) Specifies name of the IBM Cluster user who resolved the Alerts.
	* `resolution_id_string` - (String) Resolution Id String.
	* `resolved_timestamp_usecs` - (Integer) Specifies Unix epoch Timestamps in microseconds when alert is resolved.
	* `severity` - (String) Specifies the alert severity.
	  * Constraints: Allowable values are: `kCritical`, `kWarning`, `kInfo`.
	* `suppression_id` - (Integer) Specifies unique id generated when the Alert is suppressed by the admin.
	* `tenant_ids` - (List) Specifies the tenants for which this alert has been raised.
	* `vaults` - (List) Specifies information about vaults where source object associated with alert is vaulted. This could be empty if alert is not related to any source object or it is not vaulted.
	Nested schema for **vaults**:
		* `global_vault_id` - (String) Specifies Global vault id.
		* `region_id` - (String) Specifies id of region where vault resides.
		* `region_name` - (String) Specifies name of region where vault resides.
		* `vault_name` - (String) Specifies name of vault.

