---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_host"
description: |-
  Manages host in Power Virtual Server Cloud.
---

# ibm_pi_host

Create, update, and delete host with this resource.For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example enables you add a host to an existing host group in your project:

```terraform
resource "ibm_pi_host" "host" {
    pi_cloud_instance_id = "<value of the cloud_instance_id>"
    pi_host {
            display_name = "<value of the display_name>"
            sys_type = "<value of the sys_type>"
    }
    pi_host_group_id = "<value of the host_group_id>"
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

The ibm_pi_host provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) The creating of the host is considered failed if no response is received for 10 minutes.
- **delete** - (Default 10 minutes) The deletion of the host is considered failed if no response is received for 10 minutes.

## Argument Reference

You can specify the following arguments for this resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_host` - (Required, List) Host to add to a host group.
  
    Nested schema for `pi_host`:
  - `display_name` - (Required, String) Name of the host chosen by the user.
  - `sys_type` - (Required, String) System type.
  - `user_tags` - (Optional, List) The user tags attached to this resource.
- `pi_host_group_id` - (Required, String) ID of the host group to which the host should be added.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `capacity` - (List) Capacities of the host.

    Nested schema for `capacity`:
      - `available_core` - (Float) Number of cores currently available.
      - `available_memory` - (Float) Amount of memory currently available (in GB).
      - `reserved_core` - (Float) Number of cores reserved for system use.
      - `reserved_memory` - (Float) Amount of memory reserved for system use (in GB).
      - `total_core` - (Float) Total number of cores of the host.
      - `total_memory` - (Float) Total amount of memory of the host (in GB).
      - `used_core` - (Float) Number of cores in use on the host.
      - `used_memory` - (Float) Amount of memory used on the host (in GB).

- `crn` - (String) The CRN of this resource.
- `host_group` - (Map)  Information about the owning host group.
  
    Nested schema for `host_group`:
      - `access` - (String) Whether the host group is a primary or secondary host group.
      - `href` - (String) Link to the host group resource.
      - `name` - (String) Name of the host group.
  
- `host_id` - (String) ID of the host.
- `id` - The unique identifier of the host. The ID is composed of `<pi_cloud_instance_id>/<host_id>`.
- `state` - (String) State of the host `up` or `down`.
- `status` - (String) Status of the host `enabled` or `disabled`.
- `sys_type` - (String) System type.
- `user_tags` - (Set of String) The user tags attached to this resource.

## Import

The `ibm_pi_host` resource can be imported by using `pi_cloud_instance_id` and `host_id`.

### Example

```bash
terraform import ibm_pi_host.example d7bec597-4726-451f-8a63-e62e6f19c32c/b17a2b7f-77ab-491c-811e-495f8d4c8947
```
