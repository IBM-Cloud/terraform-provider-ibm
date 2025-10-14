---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_data_source_connectors"
description: |-
  Get information about Data-Source Connectors
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_data_source_connectors

Provides a read-only data source to retrieve information about Data-Source Connectors. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_data_source_connectors" "backup_recovery_data_source_connectors" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `connection_id` - (Optional, String) Specifies the ID of the connection, connectors belonging to which are to be fetched.
* `connector_ids` - (Optional, List) Specifies the unique IDs of the connectors which are to be fetched.
* `connector_names` - (Optional, List) Specifies the names of the connectors which are to be fetched.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Data-Source Connectors.
* `connectors` - (List) 
Nested schema for **connectors**:
	* `cluster_side_ip` - (String) Specifies the IP of the connector's NIC facing the cluster.
	* `connection_id` - (String) Specifies the ID of the connection to which this connector belongs.
	* `connectivity_status` - (List) Specifies status information for the data-source connector. For example if it's currently connected to the cluster, when it last connected to the cluster successfully, etc.
	Nested schema for **connectivity_status**:
		* `is_connected` - (Boolean) Specifies whether the connector is currently connected to the cluster.
		* `last_connected_timestamp_secs` - (Integer) Specifies the last timestamp in UNIX time (seconds) when the connector had successfully connected to the cluster. This property can be present even if the connector is currently disconnected.
		* `message` - (String) Specifies error message when the connector is unable to connect to the cluster.
	* `connector_id` - (String) Specifies the unique ID of the connector.
	* `connector_name` - (String) Specifies the name of the connector. The name of a connector need not be unique within a tenant or across tenants. The name of the connector can be updated as needed.
	* `software_version` - (String) Specifies the connector's software version.
	* `tenant_side_ip` - (String) Specifies the IP of the connector's NIC facing the sources of the tenant to which the connector belongs.
	* `upgrade_status` - (List) Specifies upgrade status for the data-source connector. For example when the upgrade started, current status of the upgrade, errors for upgrade failure etc.
	Nested schema for **upgrade_status**:
		* `last_status_fetched_timestamp_msecs` - (Integer) Specifies the last timestamp in UNIX time (milliseconds) when the connector upgrade status was fetched.
		* `message` - (String) Specifies error message for upgrade failure.
		* `start_timestamp_m_secs` - (Integer) Specifies the last timestamp in UNIX time (milliseconds) when the connector upgrade was triggered.
		* `status` - (String) Specifies the last fetched upgrade status of the connector.
		  * Constraints: Allowable values are: `NotStarted`, `InProgress`, `Succeeded`, `Failed`.

