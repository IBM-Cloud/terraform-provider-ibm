---
layout: "ibm"
page_title: "IBM : ibm_pi_network_interface"
description: |-
  Manages pi_network_interface.
subcategory: "Power Systems"
---

# ibm_pi_network_interface

Create, update, and delete a network interface.

## Example Usage

```terraform
  resource "ibm_pi_network_interface" "network_interface" {
    pi_cloud_instance_id = "<value of the cloud_instance_id>"
    pi_network_id = "network_id"
    pi_name = "network-interface-name"
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

The `ibm_pi_network` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating a network interface.
- **update** - (Default 60 minutes) Used for updating a network interface.
- **delete** - (Default 60 minutes) Used for deleting a network interface.
  
## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Optional, String) If supplied populated it attaches to the instance ID, if empty detaches from the instance ID.
- `pi_ip_address` - (Optional, String) The requested IP address of this network interface.
- `pi_name` - (Optional, String) Name of the network interface.
- `pi_network_id` - (Required, String) network id.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The network interface's crn.
- `id` - (String) The unique identifier of the network interface resource. The ID is composed of `<cloud_instance_id>/<network_id>/<network_interface_id>`.
- `instance` - (List) The attached instance to this network interface.

  Nested scheme for `instance`:
  - `href` - (String) Link to instance resource.
  - `instance_id` - (String) The attached instance id.
- `ip_address` - (String) The ip address of this network interface.
- `mac_address` - (String) The mac address of the network interface.
- `name` - (String) Name of the network interface (not unique or indexable).
- `network_interface_id` - (String) The unique identifier of the network interface.
- `network_security_group_id` - (Deprecated, String) ID of the network security group the network interface will be added to. Please use network_security_group_ids instead.
- `network_security_group_ids` - (List) List of network security groups that the network interface is a member of.
- `status` - (String) The status of the network address group.

## Import

The `ibm_pi_network_interface` resource can be imported by using `cloud_instance_id`, `network_id` and `network_interface_id`.

## Example

```bash
terraform import ibm_pi_network_interface.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/041b186b-9598-4cb9-bf70-966d7b9d1dc8
```
