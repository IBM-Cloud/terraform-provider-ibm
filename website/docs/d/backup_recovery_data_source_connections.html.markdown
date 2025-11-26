---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_data_source_connections"
description: |-
  Get information about Data-Source Connections
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_data_source_connections

Provides a read-only data source to retrieve information about Data-Source Connections. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_data_source_connections" "backup_recovery_data_source_connections" {
	x_ibm_tenant_id = ibm_backup_recovery_data_source_connection.backup_recovery_data_source_connection_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `connection_ids` - (Optional, List) Specifies the unique IDs of the connections which are to be fetched.
* `connection_names` - (Optional, List) Specifies the names of the connections which are to be fetched.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Data-Source Connections.
* `connections` - (List) 
Nested schema for **connections**:
	* `connection_id` - (String) Specifies the unique ID of the connection.
	* `connection_name` - (String) Specifies the name of the connection. For a given tenant, different connections can't have the same name. However, two (or more) different tenants can each have a connection with the same name.
	* `connector_ids` - (List) Specifies the IDs of the connectors in this connection.
	* `registration_token` - (String) Specifies a token that can be used to register a connector against this connection.
	* `tenant_id` - (String) Specifies the tenant ID of the connection.
	* `upgrading_connector_id` - (String) Specifies the connector ID that is currently in upgrade.

