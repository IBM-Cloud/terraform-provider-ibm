---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_disaster_recovery_location"
description: |-
  Manages a disaster recovery location in the Power Virtual Server cloud.
---

# ibm_pi_disaster_recovery_location

Retrieves information about disaster recovery location. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about the disaster recovery location present in Power Systems Virtual Server.

```terraform
data "ibm_pi_disaster_recovery_location" "ds_disaster_recovery_location" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
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

- `location` - (String) The region zone of a site.
- `replication_sites` - (List) List of replication sites.

  Nested scheme for `replication_sites`:
  - `is_active` - (Boolean) Indicates the location is active or not, `true` if location is active , otherwise it is `false`.
  - `location` - (String) The region zone of the location.
  - `replication_pool_map` - (List) List of replication pool map.

    Nested scheme for `replication_pool_map`:
    - `remote_pool` - (String) Remote pool.
    - `volume_pool` - (String) Volume pool.
