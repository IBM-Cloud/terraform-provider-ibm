---
layout: "ibm"
page_title: "IBM : ibm_pi_network_address_groups"
description: |-
  Get information about pi_network_address_groups
subcategory: "Power Systems"
---

# ibm_pi_network_address_groups

Retrieves information about a network address groups.

## Example Usage

```terraform
    data "ibm_pi_network_address_groups" "network_address_groups" {
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

- `network_address_groups` - (List) List of network address groups.

  Nested schema for `network_address_groups`:
  - `crn` - (String) The network address group's crn.
  - `id` - (String) The id of the network address group.
  - `members` - (List) The list of IP addresses in CIDR notation in the network address group.

      Nested schema for `members`:
        - `cidr` - (String) The IP addresses in CIDR notation.
        - `id` - (String) The id of the network address group member IP addresses.
  - `name` - (String) The name of the network address group.
  - `user_tags` - (List) List of user tags attached to the resource.
