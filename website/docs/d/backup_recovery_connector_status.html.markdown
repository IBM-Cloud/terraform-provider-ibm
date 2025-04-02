---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_status"
description: |-
  Get information about Data-Source Connector Status
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_status

Provides a read-only data source to retrieve information about Data-Source Connector Status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_status" "backup_recovery_connector_status" {
	access_token = resource.ibm_backup_recovery_connector_access_token.name.access_token
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `access_token` - (Required, String) Specify an access token for connector authentication.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Data-Source Connector Status.
* `cluster_connection_status` - (List) Specifies the data-source connector-cluster connectivity status.
Nested schema for **cluster_connection_status**:
	* `is_active` - (Boolean) Specifies if the connection to the cluster is active.
	* `last_connected_timestamp_msecs` - (Integer) Specifies last known connectivity status time in milliseconds.
	* `message` - (String) Specifies possible connectivity error message.
* `is_certificate_valid` - (Boolean) Flag to indicate if connector certificate is valid.
* `registration_status` - (List) Specifies the data-source connector registration status.
Nested schema for **registration_status**:
	* `message` - (String) Specifies the message corresponding the registration.
	* `status` - (String) Specifies the registration status.
	  * Constraints: Allowable values are: `NotDone`, `InProgress`, `Success`, `Failed`.

