---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agent_registration"
description: |-
  Manages Connector agent registration request.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agent_registration

Create, update, and delete Connector agent registration requests with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_connector_agent_registration" "backup_recovery_connector_agent_registration_instance" {
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connection_name` - (Required, Forces new resource, String) Specifies the name to be associated with the connector agent. This must be unique within the tenant to which this connector agent is registered.
* `join_existing_connection` - (Optional, Forces new resource, Boolean) Whether this agent is joining a connection that was already claimed by a previous registration (e.g. another agent in the same cluster for clustered sources). When true, the server adds this agent to the existing connection instead of rejecting the request as a duplicate. If the connection does not yet exist, a new one is created regardless of this flag.
  * Constraints: The default value is `false`.
* `registration_token` - (Required, Forces new resource, String) The JWT registration token. A single token can be used to register multiple connector agents in that tenant. By default, the token is valid for 24 hours.


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.
* `registration_status` - (String) Used to indicate if a duplicate registration was attempted.
  * Constraints: Allowable values are: `registered`, `already-registered`.


## Import

You can import the `ibm_backup_recovery_connector_agent_registration` resource by using `id`. id.

# Syntax
<pre>
$ terraform import ibm_backup_recovery_connector_agent_registration.backup_recovery_connector_agent_registration &lt;id&gt;
</pre>
