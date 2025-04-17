---
layout: "ibm"
page_title: "IBM : ibm_pi_volume_snapshots"
description: |-
  Get information about all your volume snapshots in Power Virtual Server.
subcategory:  "Power Systems"
---

# ibm_pi_volume_snapshots

Retrieve information about your volume snapshots.

## Example Usage

```terraform
data "ibm_pi_volume_snapshots" "snapshots" {
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

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `volume_snapshots` - (List) The list of volume snapshots.

  Nested schema for `volume_snapshots`:
  - `creation_date` - (String) The date and time when the volume snapshot was created.
  - `crn` - (Deprecated, String) The CRN of the volume snapshot.
  - `id` - (String) The volume snapshot UUID.
  - `name` - (String) The volume snapshot name.
  - `size` - (Float) The size of the volume snapshot, in gibibytes (GiB).
  - `status` - (String) The status for the volume snapshot.
  - `updated_date` - (String) The date and time when the volume snapshot was last updated.
  - `volume_id` - (String) The volume UUID associated with the snapshot.
