---
layout: "ibm"
page_title: "IBM : ibm_pi_network_address_group_member"
description: |-
  Manages pi_network_address_group_member.
subcategory: "Power Systems"
---

# ibm_pi_network_address_group_member

Add or remove a network address group member.

## Example Usage

```terraform
    resource "ibm_pi_network_address_group_member" "network_address_group_member" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_cidr = "cidr"
        pi_network_address_group_id = "network_address_group_id"
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

## Timeouts

ibm_pi_network_address_group_member provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 5 minutes) Used for creating network address group member.
- **delete** - (Default 5 minutes) Used for deleting network address group member.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cidr` - (Optional, String) The member to add in CIDR format, for example 192.168.1.5/32. Required if `pi_network_address_group_member_id` not provided.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.  
- `pi_network_address_group_id` - (Required, String) network address group id.
- `pi_network_address_group_member_id` - (Optional, String) The network address group member id to remove. Required if `pi_cidr` not provided.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network address group's crn.
- `id` -(String) The unique identifier of the network address group resource. Composed of `<cloud_instance_id>/<network_address_group_id>/network_address_group_member_id`.
- `network_address_group_id` - (String) The unique identifier of the network address group.
- `members` - (List) The list of IP addresses in CIDR notation in the network address group.

    Nested schema for `members`:
  - `cidr` - (String) The IP addresses in CIDR notation
  - `id` - (String) The id of the network address group member IP addresses.
- `name` - (String) The name of the network address group.
- `user_tags` - (List) List of user tags attached to the resource.

## Import

The `ibm_pi_network_address_group_member` resource can be imported by using `cloud_instance_id` and `network_address_group_id`, and `network_address_group_member_id`.

## Example

```bash
terraform import ibm_pi_network_address_group_member.example d7bec597-4726-451f-8a63-e62e6f19c32c/041b186b-9598-4cb9-bf70-966d7b9d1dc8
```
