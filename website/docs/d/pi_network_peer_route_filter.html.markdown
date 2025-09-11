---
layout: "ibm"
page_title: "IBM : ibm_pi_network_peer_route_filter"
description: |-
  Get information about pi_network_peer_route_filter
subcategory: "Power Systems"
---

# ibm_pi_network_peer_route_filter

Provides a read-only data source to retrieve information about a pi_network_peer_route_filter for on-prem locations.

## Example Usage

```terraform
data "ibm_pi_network_peer_route_filter" "pi_network_peer_route_filter" {
    pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
    pi_network_peer_id = "8e047a5c-d24f-11ec-9655-526e74301ad0"
    pi_route_filter_id = "a1f2b3c4-d5e6-7890-fghi-jklmno123456"
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
- `pi_network_peer_id` - (Required, String) Network peer ID.
- `pi_route_filter_id` - (Required, String) Route filter ID.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `action` - (String) Action of the filter.
- `creation_date` - (String) Time stamp for create route filter.
- `direction` - (String) Direction of the filter.
- `error` - (String) Error description.
- `ge` - (Integer) The minimum matching length of the prefix-set.
- `id` - (String) The unique identifier of the pi_network_peer_route_filter.
- `index` - (Integer) Priority or order of the filter.
- `le` - (Integer) The maximum matching length of the prefix-set.
- `prefix` - (String) IP prefix representing an address and mask length of the prefix-set.
- `state` - (String) Status of the route filter.
