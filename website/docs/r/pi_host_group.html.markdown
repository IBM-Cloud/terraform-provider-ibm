---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_host_group"
description: |-
  Manages host group in power virtual server cloud.
---

# ibm_pi_host_group

Create, update, and delete host group with this resource. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example enables you to create a host group:

```terraform
resource "ibm_pi_host_group" "hostGroup" {
    pi_cloud_instance_id  = "<value of the cloud_instance_id>"
    pi_hosts =  {
        display_name = "display_name"
        sys_type = "sys_type"
    }
    pi_name = "name"
    pi_secondaries = {
        name = "name"
        workspace = "workspace"
    }
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

The `ibm_pi_host_group` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 10 minutes) Used for deleting a host group.
  
## Argument Reference

You can specify the following arguments for this resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

- `pi_hosts` - (Required, List) List of hosts to add to the group.
  
  Nested schema for `pi_hosts`:
      - `display_name` - (Required, String) Name of the host chosen by the user.
      - `sys_type` - (Required, String) System type.
      - `user_tags` - (Optional, List) The user tags attached to this resource. Please avoid reading user tags from this attribute as environment tags will not be included. Please use appropriate data sources such as `ibm_pi_host` and `ibm_pi_hosts`.

- `pi_name` - (Required, String) Name of the host group to create.
- `pi_remove` - (Optional, String) A workspace ID to stop sharing the host group with.
- `pi_secondaries` - (Optional, List) List of workspaces to share the host group with.
  
  Nested schema for `pi_secondaries`:
      - `name` - (Optional, String) Name of the host group to create in the secondary workspace.
      - `workspace` - (Required, String) ID of the workspace to share the host group with.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `creation_date` - (String) Date/Time of host group creation.
- `host_group_id` - (String) The id of the created host group.
- `hosts` - (List) List of hosts.
- `id` - (String) The unique identifier of the host group. The ID is composed of `<pi_cloud_instance_id>/<host_group_id>`.
- `primary` - (String) The ID of the workspace owning the host group.
- `secondaries` - (List) IDs of workspaces the host group has been shared with.
