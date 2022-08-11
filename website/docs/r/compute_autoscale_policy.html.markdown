---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_autoscale_policy"
description: |-
  Manages IBM Cloud compute auto scale policy.
---

# ibm_compute_autoscale_policy
Create, update, or delete a policy for an auto scaling policy. For more information, about compute auto scale policy, see [auto scale](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-about-auto-scale).

**Note**
For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Scale_Policy).

## Example usage

In the following example, you can create an auto scaling policy:

```terraform
/* Deprecated in terraform v0.12 hence not updated */

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

## Argument reference
Review the argument references that you can specify for your resource. 

- `cooldown` - (Optional, Integer) The duration, expressed in seconds, that the policy waits after the last action date before performing another scaling action. If you do not provide a value, the `scale_group` cool down applies.
- `name` - (Required, String) The name of the autoscaling policy.
- `scale_type` - (Required, String) The scale type for the autoscaling policy. Accepted values are `ABSOLUTE`, `RELATIVE`, and `PERCENT`.
- `scale_amount`- (Required, Integer) A count of the scaling actions to perform upon any trigger hit.
- `scale_group_id`- (Required, Integer) The ID of the autoscaling group that is associated with the policy.Yes.
- `triggers` (Optional, Array of Strings) The triggers to check for this group.
- `tags` (Optional, Array of Strings) The tags that you want to add to the autoscaling policy. Tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the autoscaling policy.
