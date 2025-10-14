---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_data_source_connector_patch"
description: |-
  Manages Data-Source Connector.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_data_source_connector_patch

Update, and delete Data-Source Connectors with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_data_source_connector_patch" "backup_recovery_data_source_connector_patch_instance" {
  connector_id = "connector_id"
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connector_id` - (Required, Forces new resource, String) Specifies the unique ID of the connector which is to be deleted.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   
* `connector_name` - (Optional, Forces new resource, String) Specifies the name of the connector. The name of a connector need not be unique within a tenant or across tenants. The name of the connector can be updated as needed.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Data-Source Connector.
* `cluster_side_ip` - (String) Specifies the IP of the connector's NIC facing the cluster.
* `connection_id` - (String) Specifies the ID of the connection to which this connector belongs.
* `connectivity_status` - (List) Specifies status information for the data-source connector. For example if it's currently connected to the cluster, when it last connected to the cluster successfully, etc.
Nested schema for **connectivity_status**:
	* `is_connected` - (Boolean) Specifies whether the connector is currently connected to the cluster.
	* `last_connected_timestamp_secs` - (Integer) Specifies the last timestamp in UNIX time (seconds) when the connector had successfully connected to the cluster. This property can be present even if the connector is currently disconnected.
	* `message` - (String) Specifies error message when the connector is unable to connect to the cluster.
* `software_version` - (String) Specifies the connector's software version.
* `tenant_side_ip` - (String) Specifies the IP of the connector's NIC facing the sources of the tenant to which the connector belongs.
* `upgrade_status` - (List) Specifies upgrade status for the data-source connector. For example when the upgrade started, current status of the upgrade, errors for upgrade failure etc.
Nested schema for **upgrade_status**:
	* `last_status_fetched_timestamp_msecs` - (Integer) Specifies the last timestamp in UNIX time (milliseconds) when the connector upgrade status was fetched.
	* `message` - (String) Specifies error message for upgrade failure.
	* `start_timestamp_m_secs` - (Integer) Specifies the last timestamp in UNIX time (milliseconds) when the connector upgrade was triggered.
	* `status` - (String) Specifies the last fetched upgrade status of the connector.
	  * Constraints: Allowable values are: `NotStarted`, `InProgress`, `Succeeded`, `Failed`.


## Import
Not Supported
