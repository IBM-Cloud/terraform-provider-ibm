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
- `parameters` - (Required, String) JSON object of parameters to send.
- `destroy_parameters` - (Optional, String) JSON object to send on destroy. When omitted, destroy only removes the resource from state.

## Attribute reference

- `id` - Same as `resource_instance_id`.
- `state` - Instance state after the last update or refresh.
