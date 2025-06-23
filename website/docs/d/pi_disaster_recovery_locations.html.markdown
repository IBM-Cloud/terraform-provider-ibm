---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_disaster_recovery_locations"
description: |-
  Manages a disaster recovery location in the Power Virtual Server cloud.
---

# ibm_pi_disaster_recovery_locations

Retrieves information about disaster recovery locations. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example usage

The following example retrieves information about the disaster recovery locations present in Power Systems Virtual Server.

```terraform
data "ibm_pi_disaster_recovery_locations" "ds_disaster_recovery_locations" {}
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

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `disaster_recovery_locations` - List of Disaster Recovery Locations.

  Nested scheme for `disaster_recovery_locations`:
  - `location` - (String) The region zone of a site.
  - `replication_sites` - List of Replication Sites.

        Nested scheme for `replication_sites`:
        - `is_active` - (Boolean) Indicates the location is active or not, `true` if location is active, otherwise it is `false`.
        - `location` - (String) The region zone of the location.
        - `replication_pool_map` - (List) List of replication pool maps.

          Nested scheme for `replication_pool_map`:
          - `remote_pool` - (String) Remote pool.
          - `volume_pool` - (String) Volume pool.
