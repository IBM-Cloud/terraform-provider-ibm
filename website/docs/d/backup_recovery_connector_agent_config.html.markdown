---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agent_config"
description: |-
  Get information about Connector agent config
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agent_config

Provides a read-only data source to retrieve information about a Connector agent config. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_agent_config" "backup_recovery_connector_agent_config" {
}
```
## Argument Reference

You can specify the following arguments for this data source.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.
* `registration_token` - (String) Token that is used for authenticating the connector agent with the DataProtect cluster.

