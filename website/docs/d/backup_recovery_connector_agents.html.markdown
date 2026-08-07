---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agents"
description: |-
  Get information about backup_recovery_connector_agents
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agents

Provides a read-only data source to retrieve information about backup_recovery_connector_agents. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_agents" "backup_recovery_connector_agents" {
	tenant_id = "tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `connection_ids` - (Optional, List) Specifies the connection IDs whose connector agents are to be fetched.
* `connection_names` - (Optional, List) Specifies the connection names whose connector agents are to be fetched.
* `tenant_id` - (Required, String) Specifies the ID of the tenant for which the connector agents are to be fetched.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_connector_agents.
* `connector_agents` - (List) 
Nested schema for **connector_agents**:
	* `connection_id` - (String) Specifies the ID of the connection to which this connector agent belongs.
	* `connection_name` - (String) Specifies the name of the connection to which this connector agent belongs.
	* `connectivity_status` - (List) Specifies connector agent connectivity statusinformation like current connectivity status to cluster,when it last connected to the cluster successfully and fromwhen it has been continuously connected to the cluster without anyinterruptions.
	Nested schema for **connectivity_status**:
		* `connected_since_timestamp_secs` - (Integer) This denotes the timestamp in UNIX seconds since when this connector agent has been connected to its cluster without any interruptions. This property will not be present if this connector agent is not currently connected to its cluster.
		* `is_connected` - (Boolean) Specifies whether the connector agent is currently connected to the cluster.
		* `last_known_health_ok_timestamp_secs` - (Integer) Specifies the most recent known timestamp in UNIX seconds at which this connector agent passed the health checks. This property can be present even if this connector agent is not currently connected to its cluster.
		* `message` - (String) Specifies error message when the connector agent is unable to connect to the cluster.
	* `connector_agent_id` - (String) Specifies the unique ID of the connector agent.
	* `connector_agent_name` - (String) Specifies the name of the connector agent.
	* `software_version` - (String) Specifies the connector agent's software version.

