---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_alerts_resolution"
description: |-
  Get information about backup_recovery_manager_get_alerts_resolution
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_alerts_resolution

Provides a read-only data source to retrieve information about a backup_recovery_manager_get_alerts_resolution. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_alerts_resolution" "backup_recovery_manager_get_alerts_resolution" {
	max_resolutions = 14
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `max_resolutions` - (Required, Integer) Specifies the max number of Resolutions to be returned, from the latest created to the earliest created.
* `resolution_id` - (Optional, String) Specifies Alert Resolution id to query.
* `resolution_name` - (Optional, String) Specifies Alert Resolution Name to query.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_alerts_resolution.
* `alert_resolutions_list` - (List) List of alert resolutions.
Nested schema for **alert_resolutions_list**:
	* `account_id` - (String) Specifies account id of the user who create the resolution.
	* `created_time_usecs` - (Integer) Specifies unix epoch timestamp (in microseconds) when the resolution is created.
	* `description` - (String) Specifies the full description about the Resolution.
	* `external_key` - (String) Specifies the external key assigned outside of management console, with the form of "clusterid:resolutionid".
	* `resolution_id` - (String) Specifies the unique reslution id assigned in management console.
	* `resolution_name` - (String) Specifies the unique name of the resolution.
	* `resolved_alerts` - (List)
	Nested schema for **resolved_alerts**:
		* `alert_id` - (Integer) Id of the alert.
		* `alert_id_str` - (String) Alert Id with string format.
		* `alert_name` - (String) Name of the alert being resolved.
		* `cluster_id` - (Integer) Id of the cluster which the alert is associated.
		* `first_timestamp_usecs` - (Integer) First occurrence of the alert.
		* `resolved_time_usec` - (Integer)
		* `service_instance_id` - (String) Id of the service instance which the alert is associated.
	* `silence_minutes` - (Integer) Specifies the time duration (in minutes) for silencing alerts. If unspecified or set zero, a silence rule will be created with default expiry time. No silence rule will be created if value < 0.
	* `tenant_id` - (String) Specifies tenant id of the user who create the resolution.

