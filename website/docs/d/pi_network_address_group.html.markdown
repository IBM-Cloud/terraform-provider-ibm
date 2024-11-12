---
layout: "ibm"
page_title: "IBM : ibm_pi_network_address_group"
description: |-
  Get information about pi_network_address_group
subcategory: "Power Systems"
---

# ibm_pi_network_address_group

Retrieves information about a network address group.

## Example Usage

```terraform
    data "ibm_pi_network_address_group" "network_address_group" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>" 
        pi_network_address_group_id = "<value of the network_address_group_id>"
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

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_network_address_group_id` - (Required, String) The network address group id.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `crn` - (String) The network address group's crn.
- `id` - The unique identifier of the network address group.
- `members` - (List) The list of IP addresses in CIDR notation in the network address group.

  Nested schema for `members`:
  - `cidr` - (String) The IP addresses in CIDR notation.
  - `id` - (String) The id of the network address group member IP addresses.
- `name` - (String) The name of the network address group.
- `user_tags` - (List) List of user tags attached to the resource.
