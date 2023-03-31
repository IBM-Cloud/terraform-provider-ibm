---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scope"
description: |-
  Manages scopes.
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scope

Provides a resource for scopes. This allows scopes to be created, updated and deleted. Before creation of the scope, we need to create credential and collector.

## Example Usage

```hcl
resource "ibm_scc_posture_scope" "scopes" {
  credential_id = "5"
  credential_type = "on_premise"
  description = "IBMSchema"
  interval = 10
  is_discovery_scheduled = true
  name = "IBMSchema-new-048-test"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `collector_ids` - (Required, List) The unique IDs of the collectors that are attached to the scope.
  * Constraints: The list items must match regular expression `/^[0-9]*$/`.
* `credential_id` - (Required, String) The unique identifier of the credential.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
* `credential_type` - (Required, String) The environment that the scope is targeted to.
  * Constraints: Allowable values are: `ibm`, `aws`, `azure`, `on_premise`, `hosted`, `services`, `openstack`, `gcp`.
* `description` - (Required, String) A detailed description of the scope.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
* `interval` - (Optional, Integer) Stores the value of Frequency. This is used in case of on-prem Scope if the user wants to schedule a discovery task.The unit is seconds. Example if a user wants to trigger discovery every hour, this value will be set to 3600.
* `is_discovery_scheduled` - (Optional, Boolean) Stores the value of Discovery Scheduled.This is used in case of on-prem Scope if the user wants to schedule a discovery task.
  * Constraints: The default value is `false`.
* `name` - (Required, String) A unique name for your scope.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the scopes.

## Import

You can import the `ibm_scc_posture_scope` resource by using `id`. An identifier of the scope.

# Syntax
```
$ terraform import ibm_scc_posture_scope.scopes <id>
```

# Example
```
$ terraform import ibm_scc_posture_scope.scopes 1
```

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scope is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
