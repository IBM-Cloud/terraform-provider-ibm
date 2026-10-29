---
subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : ibm_resource_instance_parameters"
description: |-
  Applies service parameters to an existing resource instance.
---

# ibm_resource_instance_parameters

Applies service-specific parameters to an existing resource instance through the Resource Controller update API (`PATCH /v2/resource_instances/{id}`). Use it when the parameters depend on resources created after the instance, which would otherwise be a dependency cycle on `ibm_resource_instance`. 

Drift is checked only on the top-level keys in `parameters`. Server-generated fields are ignored.

When the Resource Controller runs the update asynchronously (`last_operation.async = true`), create and update wait until the operation finishes and the instance is `active` again. A failed operation fails the apply with the operation's description.

## Example usage

```terraform
resource "ibm_resource_instance_parameters" "cspm" {
  resource_instance_id = ibm_resource_instance.scc_wp.guid
  parameters = jsonencode({
    enable_cspm = true
    target_accounts = [{
      account_id         = ibm_resource_instance.scc_wp.account_id
      account_type       = "account"
      config_crn         = var.app_config_crn
      trusted_profile_id = module.trusted_profile.profile_id
    }]
  })
  destroy_parameters = jsonencode({ enable_cspm = false, target_accounts = [] })
}
```

## Argument reference

- `resource_instance_id` - (Required, Forces new resource, String) GUID or CRN of the instance.
- `parameters` - (Required, String) JSON object of parameters to send. Checked at plan time.
- `destroy_parameters` - (Optional, String) JSON object to send on destroy. Checked at plan time. When omitted, destroy only removes the resource from state. If the instance is already deleted (`404 Not Found` or `410 Gone`), destroy succeeds without sending anything.

## Attribute reference

- `id` - Same as `resource_instance_id`.
- `state` - Instance state after the last update or refresh.

## Timeouts

- `create` - (Default 30 minutes) Waiting for an asynchronous update.
- `update` - (Default 30 minutes) Waiting for an asynchronous update.
- `delete` - (Default 30 minutes) Waiting for an asynchronous update with `destroy_parameters`.

## Import

Import with the instance GUID or CRN. Only the keys you then set in `parameters` are managed, and the first apply sends them again (the update is idempotent).

```terraform
import {
  to = ibm_resource_instance_parameters.cspm
  id = ibm_resource_instance.scc_wp.guid
}
```

```
$ terraform import ibm_resource_instance_parameters.cspm <instance_guid>
```
