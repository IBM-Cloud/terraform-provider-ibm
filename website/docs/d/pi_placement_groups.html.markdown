---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_placement_groups"
description: |-
  Manages placement groups in the Power Virtual Server cloud.
---

# ibm_pi_placement_groups

Retrieve information about all placement groups. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_placement_groups" "example" {
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

- `placement_groups` - (List) List of all the placement groups.

  Nested scheme for `placement_groups`:
  - `id` - (String) The ID of the placement group.
  - `members` - (List) List of server instances IDs that are members of the placement group.
  - `name` - (String) User defined name for the placement group.
  - `policy` - (String) The value of the group's affinity policy. Valid values are affinity and anti-affinity.
