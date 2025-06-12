---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: ibm_pi_host_group"
description: |-
  Get information about a host group.
---

# ibm_pi_host_group

Provides a read-only data source to retrieve information about a host group you can use in Power Systems Virtual Server. For more information, about Power Systems Virtual Server host group, see [host groups](https://cloud.ibm.com/apidocs/power-cloud#endpoint).

## Example Usage

```terraform
data "ibm_pi_host_group" "ds_host_group" {
    pi_cloud_instance_id    = "<value of the cloud_instance_id>"
    pi_host_group_id         = "<value of the host_group_id>"
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

- `pi_host_group_id` - (Required, String) Host group ID.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `creation_date` - (String) Date/Time of host group creation.

- `hosts` - (List) List of hosts.
  
- `id` - The unique identifier of the host group.

- `name` - (String) Name of the host group.

- `primary` - (String) ID of the workspace owning the host group.

- `secondaries` - (List) IDs of workspaces the host group has been shared with.
