---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM : ibm_cr_plan_upgrade"
description: |-
  Upgrades the IBM Cloud Container Registry plan of the account.
---

# ibm_cr_plan_upgrade (Action)

Upgrades the Container Registry plan of the account for one registry region. The action reads the current plan first and does nothing when the plan already matches.

## Example usage

```terraform
resource "terraform_data" "registry" {
  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.ibm_cr_plan_upgrade.standard]
    }
  }
}

action "ibm_cr_plan_upgrade" "standard" {
  config {
    endpoint = "https://us.icr.io"
  }
}
```

## Argument reference

- `plan` - (Optional, String) Target plan. Default `Standard`, the only value the API accepts today.
- `endpoint` - (Optional, String) Regional Container Registry endpoint, for example `https://de.icr.io`. Defaults to the provider endpoint.
