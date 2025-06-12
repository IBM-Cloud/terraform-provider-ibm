---
layout: "ibm"
page_title: "IBM : ibm_schematics_agents"
description: |-
  Get information about schematics_agents
subcategory: "Schematics"
---

# ibm_schematics_agents

Provides a read-only data source for schematics_agents. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_agents" "schematics_agents" {
	name = "MyDevAgent"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Optional, String) The name of the agent (must be unique, for an account).

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_agents.
* `agents` - (List) The list of agents in the account.
Nested scheme for **agents**:
	* `agent_crn` - (String) The Agent crn, obtained from the Schematics Agent deployment configuration.
	* `agent_location` - (String) The location where agent is deployed in the user environment.
	* `connection_state` - (List) Connection status of the agent.
	Nested scheme for **connection_state**:
		* `checked_at` - (String) When the connection state is modified.
		* `state` - (String) Agent Connection Status  * `Connected` When Schematics is able to connect to the agent.  * `Disconnected` When Schematics is able not connect to the agent.
		  * Constraints: Allowable values are: `Connected`, `Disconnected`.
	* `description` - (String) Agent description.
	* `id` - (String) The Agent registration id.
	* `location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
	  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
	* `name` - (String) The name of the agent.
	* `profile_id` - (String) The IAM trusted profile id, used by the Agent instance.
	* `registered_at` - (String) The Agent registration date-time.
	* `registered_by` - (String) The email address of an user who registered the Agent.
	* `resource_group` - (String) The resource-group name for the agent.  By default, Agent will be registered in Default Resource Group.
	* `system_state` - (List) Computed state of the agent.
	Nested scheme for **system_state**:
		* `message` - (String) The Agent status message.
		* `state` - (String) Agent Status.
		  * Constraints: Allowable values are: `error`, `normal`, `in_progress`, `pending`, `draft`.
	* `tags` - (List) Tags for the agent.
	* `updated_at` - (String) The Agent registration updation time.
	* `updated_by` - (String) Email address of user who updated the Agent registration.
	* `user_state` - (List) User defined status of the agent.
	Nested scheme for **user_state**:
		* `set_at` - (String) When the User who set the state of the Object.
		* `set_by` - (String) Name of the User who set the state of the Object.
		* `state` - (String) User-defined states  * `enable`  Agent is enabled by the user.  * `disable` Agent is disbaled by the user.
		  * Constraints: Allowable values are: `enable`, `disable`.

* `limit` - (Integer) The number of records returned.

* `offset` - (Integer) The skipped number of records.

* `total_count` - (Integer) The total number of records.

