---
layout: "ibm"
page_title: "IBM : compute_autoscale_policy"
sidebar_current: "docs-ibm-resource-compute-autoscale-policy"
description: |-
  Manages IBM Compute Auto Scale Policy.
---

# ibm\_compute_autoscale_policy

Provides an auto scaling policy resource. This allows policies for auto scale groups to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Policy).

## Example Usage

In the following example, you can create an auto scaling policy:

```hcl
resource "ibm_compute_autoscale_policy" "test_scale_policy" {
    name = "test_scale_policy_name"
    scale_type = "RELATIVE"
    scale_amount = 1
    cooldown = 30
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {
                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "80"
                    period = 120
        }
    }
    triggers = {
        type = "ONE_TIME"
        date = "2016-07-30T23:55:00-00:00"
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED *"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the auto scaling policy.
* `scale_type` - (Required, string) The scale type for the auto scaling policy. Accepted values are `ABSOLUTE`, `RELATIVE`, and `PERCENT`.
* `scale_amount` - (Required, integer) A count of the scaling actions to perform upon any trigger hit.
* `cooldown` - (Optional, integer) The duration, expressed in seconds, that the policy waits after the last action date before performing another scaling action. If you do not provide a value, the `scale_group` cooldown applies.
* `scale_group_id` - (Required, integer) The ID of the auto scale group associated with the policy.
* `triggers` - (Optional, array of integers and strings) The triggers to check for this group.
* `tags` - (Optional, array of strings) Tags associated with the auto scaling policy instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the auto scaling policy.
