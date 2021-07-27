---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_cloud_instance"
description: |-
  Provides data for a cloud instance in an IBM Power Virtual Server cloud.
---

# ibm_pi_cloud_instance
Retrieve information about an existing IBM Power Virtual Server Cloud Instance as a read-only data source. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_cloud_instance" "ds_cloud_instance" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`

Example Usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument reference
Review the argument reference that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with an account. 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `capabilities` -  (String) Lists the capabilities for this cloud instance.
- `creation_date` - (String) The date on which the tenant was created.
- `enabled` -  (Bool) Indicates whether the tenant is enabled.
- `id` - (String) The unique identifier for this tenant.
- `tenant_name` -  (String) The name of the tenant.
- `total_instances` -  (String) The count of lpars that belong to this specific cloud instance.
- `total_memory_consumed` -  (String) The total memory consumed by this service instance.
- `total_processors_consumed` -  (String) The total processors consumed by this service instance.
- `total_ssd_storage_consumed` -  (String) The total SSD Storage consumed by this service instance.
- `total_standard_storage_consumed` -  (String) The total Standard Storage consumed by this service instance.
