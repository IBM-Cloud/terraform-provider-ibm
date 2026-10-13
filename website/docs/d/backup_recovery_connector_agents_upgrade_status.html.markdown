---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agents_upgrade_status"
description: |-
  Get information about backup_recovery_connector_agents_upgrade_status
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agents_upgrade_status

Provides a read-only data source to retrieve information about backup_recovery_connector_agents_upgrade_status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_agents_upgrade_status" "backup_recovery_connector_agents_upgrade_status" {
	upgrade_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `connection_ids` - (Optional, List) Filters the connectionStatuses within each returned tenant to only these connections. Connections not part of this upgrade run are silently ignored. If omitted, every connection under each returned tenant is returned.
* `tenant_ids` - (Optional, List) Filters the returned tenantStatuses to only these tenants. Tenants not part of this upgrade run are silently ignored. If omitted, every tenant in the run is returned.
* `upgrade_id` - (Required, Forces new resource, String) Specifies the ID of the upgrade run, as returned by POST /connector-agents/upgrade.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_connector_agents_upgrade_status.
* `status` - (String) Specifies the overall state of this upgrade run, derived from the states of every tenant in the run. 'NotStarted' indicates the run has not yet begun. 'InProgress' indicates at least one tenant is still upgrading. 'Failed' indicates no tenant is in progress and at least one has failed. 'Succeeded' indicates every tenant has succeeded. Unlike tenantStatuses, this field is not affected by the tenantIds or connectionIds filters.
  * Constraints: Allowable values are: `NotStarted`, `InProgress`, `Succeeded`, `Failed`.
* `tenant_statuses` - (List) Specifies the upgrade status for every tenant included in this upgrade run, filtered by the tenantIds query parameter, if specified in the request.
Nested schema for **tenant_statuses**:
	* `info` - (List) Specifies the upgrade progress and outcome for a tenant.
	Nested schema for **info**:
		* `connection_statuses` - (List) Specifies the upgrade status of every individual connection for this tenant included in this upgrade run, filtered by the connectionIds query parameter, if specified in the request. Absent if the run has no connections.
		Nested schema for **connection_statuses**:
			* `attempt_count` - (Integer) Specifies the number of upgrade attempts made so far for this connection in the current upgrade run. Absent if no attempt has been made.
			* `connection_id` - (Integer) Specifies the ID of the connection.
			* `error_message` - (String) Specifies the error message when the status is 'Failed'. Absent otherwise.
			* `last_attempt_timestamp_msecs` - (Integer) Specifies the timestamp in UNIX time (milliseconds) of the most recent upgrade attempt for this connection. Absent if no attempt has been made.
			* `last_status_fetched_timestamp_msecs` - (Integer) Specifies the timestamp in UNIX time (milliseconds) when this connection's upgrade status is last fetched from the connector agent. Absent if the status has not yet been fetched.
			* `status` - (String) Specifies the last fetched upgrade status of the connector agent on this connection. 'NotStarted' indicates the upgrade has not yet begun. 'InProgress' indicates the upgrade is underway. 'Succeeded' indicates the connection upgraded successfully. 'Failed' indicates the upgrade attempt failed.
			  * Constraints: Allowable values are: `NotStarted`, `InProgress`, `Succeeded`, `Failed`.
		* `created_timestamp_usecs` - (Integer) Specifies the timestamp in UNIX time (microseconds) when this tenant's upgrade marker is created. Absent if not yet recorded.
		* `failed_connections` - (Integer) Specifies the number of connections for this tenant that exhaust their retry budget without upgrading successfully. Not affected by the connectionIds filter. Absent until the run is initialized.
		* `finished_timestamp_usecs` - (Integer) Specifies the timestamp in UNIX time (microseconds) when every connection for this tenant reached a terminal state. Absent while this tenant's upgrade is still in progress.
		* `retry_forever` - (Boolean) Specifies whether this tenant's upgrade retries indefinitely (true) or is capped at a fixed number of attempts (false), as set on the originating upgrade request. Absent if not specified on the originating upgrade request.
		* `state` - (String) Specifies the state of the upgrade run for this tenant. 'NotStarted' indicates the run has not yet begun. 'InProgress' indicates at least one connection is still upgrading. 'Failed' indicates every connection has reached a terminal state and at least one has failed. 'Succeeded' indicates every connection has succeeded.
		  * Constraints: Allowable values are: `NotStarted`, `InProgress`, `Succeeded`, `Failed`.
		* `succeeded_connections` - (Integer) Specifies the number of connections for this tenant that upgrade successfully to the target version. Not affected by the connectionIds filter. Absent until the run is initialized.
		* `target_version` - (String) Specifies the version every connector agent for this tenant is being upgraded to. Absent if not specified on the originating upgrade request.
		* `total_connections` - (Integer) Specifies the total number of connections in scope for this tenant's upgrade run. Not affected by the connectionIds filter. Absent until the run is initialized.
	* `tenant_id` - (String) Specifies the ID of the tenant.

