---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_alerts_stats"
description: |-
  Get information about Active Alerts stats response.
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_alerts_stats

Provides a read-only data source to retrieve information about an Active Alerts stats response.. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_alerts_stats" "backup_recovery_manager_get_alerts_stats" {
	end_time_usecs = 12
	start_time_usecs = 14
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `alert_source` - (Optional, String) Specifies a list of alert origination source. If not specified, all alerts from all the sources are considered in the response.
  * Constraints: Allowable values are: `kCluster`, `kHelios`.
* `cluster_ids` - (Optional, List) Specifies the list of cluster IDs.
* `end_time_usecs` - (Required, Integer) Specifies the end time Unix time epoch in microseconds to which the active alerts stats are computed.
* `exclude_stats_by_cluster` - (Optional, Boolean) Specifies if stats of active alerts per cluster needs to be excluded. If set to false (default value), stats of active alerts per cluster is included in the response. If set to true, only aggregated stats summary will be present in the response.
* `region_ids` - (Optional, List) Filter by a list of region ids.
* `service_instance_ids` - (Optional, List) Specifies list of service instance ids to filter alert stats by.
* `start_time_usecs` - (Required, Integer) Specifies the start time Unix time epoch in microseconds from which the active alerts stats are computed.
* `tenant_ids` - (Optional, List) Specifies a list of tenants.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Active Alerts stats response..
* `aggregated_alerts_stats` - (List) Specifies the active alert statistics details.
Nested schema for **aggregated_alerts_stats**:
	* `num_critical_alerts` - (Integer) Specifies the count of active critical Alerts excluding alerts that belong to other bucket.
	* `num_critical_alerts_categories` - (Integer) Specifies the count of active critical alerts categories.
	* `num_data_service_alerts` - (Integer) Specifies the count of active service Alerts.
	* `num_data_service_critical_alerts` - (Integer) Specifies the count of active service critical Alerts.
	* `num_data_service_info_alerts` - (Integer) Specifies the count of active service info Alerts.
	* `num_data_service_warning_alerts` - (Integer) Specifies the count of active service warning Alerts.
	* `num_hardware_alerts` - (Integer) Specifies the count of active hardware Alerts.
	* `num_hardware_critical_alerts` - (Integer) Specifies the count of active hardware critical Alerts.
	* `num_hardware_info_alerts` - (Integer) Specifies the count of active hardware info Alerts.
	* `num_hardware_warning_alerts` - (Integer) Specifies the count of active hardware warning Alerts.
	* `num_info_alerts` - (Integer) Specifies the count of active info Alerts excluding alerts that belong to other bucket.
	* `num_info_alerts_categories` - (Integer) Specifies the count of active info alerts categories.
	* `num_maintenance_alerts` - (Integer) Specifies the count of active Alerts of maintenance bucket.
	* `num_maintenance_critical_alerts` - (Integer) Specifies the count of active other critical Alerts.
	* `num_maintenance_info_alerts` - (Integer) Specifies the count of active other info Alerts.
	* `num_maintenance_warning_alerts` - (Integer) Specifies the count of active other warning Alerts.
	* `num_software_alerts` - (Integer) Specifies the count of active software Alerts.
	* `num_software_critical_alerts` - (Integer) Specifies the count of active software critical Alerts.
	* `num_software_info_alerts` - (Integer) Specifies the count of active software info Alerts.
	* `num_software_warning_alerts` - (Integer) Specifies the count of active software warning Alerts.
	* `num_warning_alerts` - (Integer) Specifies the count of active warning Alerts excluding alerts that belong to other bucket.
	* `num_warning_alerts_categories` - (Integer) Specifies the count of active warning alerts categories.
* `aggregated_cluster_stats` - (List) Specifies the cluster statistics based on active alerts.
Nested schema for **aggregated_cluster_stats**:
	* `num_clusters_with_critical_alerts` - (Integer) Specifies the count of clusters with at least one critical alert.
	* `num_clusters_with_warning_alerts` - (Integer) Specifies the count of clusters with at least one warning category alert and no critical alerts.
	* `num_healthy_clusters` - (Integer) Specifies the count of clusters with no warning or critical alerts.
* `stats_by_cluster` - (List) Specifies the active Alerts stats by clusters.
Nested schema for **stats_by_cluster**:
	* `alerts_stats` - (List) Specifies the active alert statistics details.
	Nested schema for **alerts_stats**:
		* `num_critical_alerts` - (Integer) Specifies the count of active critical Alerts excluding alerts that belong to other bucket.
		* `num_critical_alerts_categories` - (Integer) Specifies the count of active critical alerts categories.
		* `num_data_service_alerts` - (Integer) Specifies the count of active service Alerts.
		* `num_data_service_critical_alerts` - (Integer) Specifies the count of active service critical Alerts.
		* `num_data_service_info_alerts` - (Integer) Specifies the count of active service info Alerts.
		* `num_data_service_warning_alerts` - (Integer) Specifies the count of active service warning Alerts.
		* `num_hardware_alerts` - (Integer) Specifies the count of active hardware Alerts.
		* `num_hardware_critical_alerts` - (Integer) Specifies the count of active hardware critical Alerts.
		* `num_hardware_info_alerts` - (Integer) Specifies the count of active hardware info Alerts.
		* `num_hardware_warning_alerts` - (Integer) Specifies the count of active hardware warning Alerts.
		* `num_info_alerts` - (Integer) Specifies the count of active info Alerts excluding alerts that belong to other bucket.
		* `num_info_alerts_categories` - (Integer) Specifies the count of active info alerts categories.
		* `num_maintenance_alerts` - (Integer) Specifies the count of active Alerts of maintenance bucket.
		* `num_maintenance_critical_alerts` - (Integer) Specifies the count of active other critical Alerts.
		* `num_maintenance_info_alerts` - (Integer) Specifies the count of active other info Alerts.
		* `num_maintenance_warning_alerts` - (Integer) Specifies the count of active other warning Alerts.
		* `num_software_alerts` - (Integer) Specifies the count of active software Alerts.
		* `num_software_critical_alerts` - (Integer) Specifies the count of active software critical Alerts.
		* `num_software_info_alerts` - (Integer) Specifies the count of active software info Alerts.
		* `num_software_warning_alerts` - (Integer) Specifies the count of active software warning Alerts.
		* `num_warning_alerts` - (Integer) Specifies the count of active warning Alerts excluding alerts that belong to other bucket.
		* `num_warning_alerts_categories` - (Integer) Specifies the count of active warning alerts categories.
	* `cluster_id` - (Integer) Specifies the Cluster Id.
	* `region_id` - (String) Specifies the region id of cluster.

