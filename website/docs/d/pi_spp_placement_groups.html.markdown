---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_spp_placement_groups"
description: |-
  Manages the shared processor pool placement groups in the Power Virtual Server cloud.
---

# ibm_pi_spp_placement_groups

Retrieve information about all shared processor pool placement groups. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_spp_placement_groups" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
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

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `spp_placement_groups` - (List) List of all the shared processor pool placement groups.

  Nested scheme for `spp_placement_groups`:
  - `crn` - (String) The CRN of this resource.
  - `members` - (List) The list of shared processor pool IDs that are members of the shared processor pool placement group.
  - `name` - (String) User defined name for the shared processor pool placement group.
  - `policy` - (String) The value of the group's affinity policy. Valid values are affinity and anti-affinity.
  - `spp_placement_group_id` - (String) The ID of the shared processor pool placement group.
  - `user_tags` - (List) List of user tags attached to the resource.
