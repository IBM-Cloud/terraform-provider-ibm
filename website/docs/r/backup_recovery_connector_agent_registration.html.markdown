---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agent_registration"
description: |-
  Manages Connector agent registration request.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agent_registration

Create Connector agent registration requests with this resource.

~> **NOTE:** This resource must be executed on the source VSI (Virtual Server Instance) where the connector agent is installed, as it calls localhost APIs. The user needs to copy the Terraform provider binary to the VSI and run Terraform from there.

## Example Usage

```hcl
resource "ibm_backup_recovery_connector_agent_registration" "backup_recovery_connector_agent_registration_instance" {
  registration_token = "eyJ"
	connection_name = "terra-conn-register-Connector-1"
	join_existing_connection = false
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

Not supported
