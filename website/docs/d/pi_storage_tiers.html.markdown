---
layout: "ibm"
page_title: "IBM : ibm_pi_storage_tiers"
description: |-
  Get information about a storage tiers for a pi cloud instance.
subcategory: "Power Systems"
---

# ibm_pi_storage_tiers

Retrieve information about supported storage tiers for a pi cloud instance. For more information, see [storage tiers docs](https://cloud.ibm.com/apidocs/power-cloud#pcloud-cloudinstances-storagetiers-getall).

## Example Usage

```terraform
    data "ibm_pi_storage_tiers" "pi_storage_tiers" {
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

- `id` - (String) The unique identifier of the storage tiers.
- `region_storage_tiers` - (List) An array of storage tiers supported in a region.

    Nested schema for `region_storage_tiers`:
  - `description` - (String) Description of the storage tier label.
  - `name` - (String) Name of the storage tier.
  - `state` - (String) State of the storage tier, `active` or `inactive`.
