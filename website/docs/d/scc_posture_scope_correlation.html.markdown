---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scope_correlation"
description: |-
  Get information about scope_correlation
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scope_correlation

Provides a read-only data source for scope_correlation. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_scope_correlation" "scope_correlation" {
	correlation_id = "correlation_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `correlation_id` - (Required, Forces new resource, String) A correlation_Id is created when a scope is created and discovery task is triggered or when a validation is triggered on a Scope. This is used to get the status of the task(discovery or validation).
  * Constraints: The maximum length is `50` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scope_correlation.
* `last_heartbeat` - (String) Returns the time that the scope was last updated. This value exists when collector is installed and running.

* `start_time` - (String) Returns the time that task started.

* `status` - (String) Returns the current status of a task.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scope_correlation is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
