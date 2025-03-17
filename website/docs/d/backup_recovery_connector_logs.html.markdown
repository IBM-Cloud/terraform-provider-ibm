---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_logs"
description: |-
  Get information about Data-Source Connector Logs
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_logs

Provides a read-only data source to retrieve information about Data-Source Connector Logs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_logs" "backup_recovery_connector_logs" {
	access_token = resource.ibm_backup_recovery_connector_access_token.name.access_token
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `access_token` - (Required, String) Specify an access token for connector authentication.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Data-Source Connector Logs.
* `connector_logs` - (List) Specifies the data-source connector logs.
Nested schema for **connector_logs**:
	* `message` - (String) Specifies the message of this event.
	* `timestamp_msecs` - (Integer) Specifies the time stamp in milliseconds of the event.
	* `type` - (String) Specifies the severity of the event.
	  * Constraints: Allowable values are: `Info`, `Warning`, `Error`.

