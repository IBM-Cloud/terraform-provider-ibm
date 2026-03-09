---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_data_source_connection"
description: |-
  Manages Data-Source Connection.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_data_source_connection

Create, update, and delete Data-Source Connections with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_data_source_connection" "backup_recovery_data_source_connection_instance" {
  connection_name = "connection_name"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connection_env_type` - (Optional, String) Specifies the environment type of the connection.
  * Constraints: Allowable values are: `kRoksVpc`, `kRoksClassic`, `kIksVpc`, `kIksClassic`.
* `connection_name` - (Required, String) Specifies the name of the connection. For a given tenant, different connections can't have the same name. However, two (or more) different tenants can each have a connection with the same name.
* `x_ibm_tenant_id` - (Optional, String) Id of the tenant accessing the cluster.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Data-Source Connection.
* `connection_id` - The unique identifier of Connection.
* `connector_ids` - (List) Specifies the IDs of the connectors in this connection.
* `registration_token` - (String) Specifies a token that can be used to register a connector against this connection.
* `tenant_id` - (String) Specifies the tenant ID of the connection.
* `upgrading_connector_id` - (String) Specifies the connector ID that is currently in upgrade.


### Import
You can import the `ibm_backup_recovery_data_source_connection` resource by using `id`. The ID is formed using tenantID and resourceId.
`id = <tenantId>::<connection_id>`. 

#### Syntax
```
import {
	to = <ibm_backup_recovery_resource>
	id = "<tenantId>::<connection_id>"
}
```

#### Example
```
resource "ibm_backup_recovery_data_source_connection" "backup_recovery_data_source_connection_instance" {
	x_ibm_tenant_id = "jhxqx715r9/"
	connection_name = "terraform-conn"
}

import {
	to = ibm_backup_recovery_data_source_connection.backup_recovery_data_source_connection_instance
	id = "jhxqx715r9/::3309023926479362560"
}
```