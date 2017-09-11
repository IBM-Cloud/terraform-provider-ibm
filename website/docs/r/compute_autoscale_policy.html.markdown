---
layout: "ibm"
page_title: "IBM : compute_autoscale_policy"
sidebar_current: "docs-ibm-resource-compute-autoscale-policy"
description: |-
  Manages IBM Compute Auto Scale Policy.
---

# ibm\_compute_autoscale_policy

Provides a resource for auto scaling policies. This allows policies for auto scale groups to be created, updated, and deleted.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API documentation](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Policy).

## Example Usage

```hcl
# Create an auto scaling policy
resource "ibm_compute_autoscale_policy" "test_scale_policy" {
    name = "test_scale_policy_name"
    scale_type = "RELATIVE"         # Other accepted values: ABSOLUTE, PERCENT
    scale_amount = 1
    cooldown = 30                   # If not provided, the scale_group cooldown applies
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

* `name` - (Required, string) Name of the auto scaling policy.
* `scale_type` - (Required, string) The scale type for the auto scaling policy. Accepted values are `ABSOLUTE`, `RELATIVE`, and `PERCENT`.
* `scale_amount` - (Required, integer) A count of the scaling actions to perform upon any trigger hit.
* `cooldown` - (Optional, integer) The duration, expressed in seconds, that the policy waits after the last action date before performing another scaling action.
* `scale_group_id` - (Required, integer) Specify the ID of the auto scale group this policy is on.
* `triggers` - (Optional, array of integers and strings) The triggers to check for this group.
* `tags` - (Optional, array of strings) Set tags on the auto scaling policy instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the auto scaling policy.
