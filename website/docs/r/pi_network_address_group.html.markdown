---
layout: "ibm"
page_title: "IBM : ibm_pi_network_address_group"
description: |-
  Manages network address group.
subcategory: "Power Systems"
---

# ibm_pi_network_address_group

Create, update, and delete a network address group.

## Example Usage

The following example creates a network address group.

```terraform
    resource "ibm_pi_network_address_group" "network_address_group_instance" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_name = "name"
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

ibm_pi_network_address_group provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 5 minutes) Used for deleting network address group.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_name` - (Required, String) The name of the Network Address Group.
- `pi_user_tags` - (Optional, List) List of user tags attached to the resource.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network address group's crn.
- `id` - (String) The unique identifier of this resource. The ID is composed of `<pi_cloud_instance_id>/<network_address_group_id>`.
- `network_address_group_id` - (String) The unique identifier of the network address group.
- `members` - (List) The list of IP addresses in CIDR notation in the network address group.

  Nested schema for `members`:
  - `cidr` - (String) The IP addresses in CIDR notation.
  - `id` - (String) The id of the network address group member IP addresses.

## Import

The `ibm_pi_network_address_group` resource can be imported by using `cloud_instance_id` and `network_address_group_id`.

## Example

```bash
terraform import ibm_pi_network_address_group.example d7bec597-4726-451f-8a63-e62e6f19c32c/041b186b-9598-4cb9-bf70-966d7b9d1dc8
```
