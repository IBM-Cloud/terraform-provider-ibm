---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_snapshot"
description: |-
  Manages an instance snapshot in the Power Virtual Server cloud.
---

# ibm_pi_instance_snapshot

Retrieve information about a Power Systems Virtual Server instance snapshot. For more information, about Power Virtual Server instance snapshot, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_instance_snapshot" "ds_instance_snapshot" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_snapshot_id = "2ce22124-4c32-6456-bfce-803e0658ab12"
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

## Argument reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_snapshot_id` - (Required, String) The unique identifier of the Power Systems Virtual Machine instance snapshot.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `action` - (String) Action performed on the instance snapshot.
- `creation_date` - (String) Date of snapshot creation.
- `crn` - (String) The CRN of this resource.
- `description` - (String) The description of the snapshot.
- `last_updated_date` - (String) Date of last update.
- `name` - (String) The name of the Power Systems Virtual Machine instance snapshot.
- `percent_complete` - (Integer) The snapshot completion percentage.
- `status` - (String) The status of the Power Virtual Machine instance snapshot.
- `user_tags` - (List) List of user tags attached to the resource.
- `volume_snapshots` - (Map) A map of volume snapshots included in the Power Virtual Machine instance snapshot.
