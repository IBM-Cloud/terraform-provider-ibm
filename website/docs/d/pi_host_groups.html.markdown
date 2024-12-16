---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: ibm_pi_host_groups"
description: |-
  Get information about host groups.
---

# ibm_pi_host_groups

Provides a read-only data source to retrieve information about host groups. you can use in Power Systems Virtual Server. For more information, about ower Systems Virtual Server host group, see [host groups](https://cloud.ibm.com/apidocs/power-cloud#endpoint).

## Example Usage

```terraform
data "ibm_pi_host_groups" "hostGroup" {
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

- `host_groups` - (List) List of host groups.
  - Nested sheme for `host_groups`:
    - `creation_date` - (String) Date/Time of host group creation.
    - `hosts` - (List) List of hosts.
    - `id` - (String) The unique identifier of the host group.
    - `name` - (String) Name of the host group.
    - `primary` - (String) ID of the workspace owning the host group.
    - `secondaries` - (List) IDs of workspaces the host group has been shared with.
