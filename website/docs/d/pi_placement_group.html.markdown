---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_placement_group"
description: |-
  Manages a placement group in the Power Virtual Server cloud.
---

# ibm_pi_placement_group

Retrieve information about a placement group. For more information, about placement groups, see [Managing server placement groups](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-placement-groups).

## Example Usage

```terraform
data "ibm_pi_placement_group" "ds_placement_group" {
  pi_placement_group_name   = "my-pg"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_placement_group_name` - (Required, String) The name of the placement group.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `crn` - (String) The CRN of this resource.
- `id` - (String) The ID of the placement group.
- `members` - (List) List of server instances IDs that are members of the placement group.
- `policy` - (String) The value of the group's affinity policy. Valid values are affinity and anti-affinity.
- `user_tags` - (List) List of user tags attached to the resource.
