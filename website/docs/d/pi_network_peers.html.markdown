---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_network_peers"
description: |-
  Get information about IBM Power Virtual Server cloud network peers.
---

# ibm_pi_network_peers

Provides a read-only data source to retrieve information about pi_network_peers for on-prem locations.

## Example Usage

```terraform
data "ibm_pi_network_peers" "pi_network_peers" {
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

- `id` - (String) The unique identifier of the pi_network_peers.
- `network_peers` - (List) List of network peers.
  
  Nested schema for `network_peers`:
      - `creation_date` - (String) Time stamp for create network peer.
      - `customer_asn` - (Integer) ASN number at customer network side.
      - `customer_cidr` - (String) IP address used for configuring customer network interface with network subnet mask.
      - `default_export_route_filter` - (String) Default action for export route filter.
      - `default_import_route_filter` - (String) Default action for import route filter.
      - `error` - (String) Error description.
      - `export_route_filters` - (List) List of export route filters.

        Nested schema for `export_route_filters`:
         - `action` - (String) Action of the filter.
         - `creation_date` - (String) Time stamp for create route filter.
         - `direction` - (String) Direction of the filter.
         - `error` - (String) Error description.
         - `ge` - (Integer) The minimum matching length of the prefix-set.
         - `index` - (Integer) Priority or order of the filter.
         - `le` - (Integer) The maximum matching length of the prefix-set.
         - `prefix` - (String) IP prefix representing an address and mask length of the prefix-set.
         - `route_filter_id` - (String) Route filter ID.
         - `state` - (String) Status of the route filter.
      - `ibm_asn` - (Integer) ASN number at IBM PowerVS side.
      - `ibm_cidr` - (String) IP address used for configuring IBM network interface with network subnet mask.
      - `id` - (String) ID of the network peer.
      - `import_route_filters` - (List) List of import route filters.
          
          Nested schema for `import_route_filters`:
          - `action` - (String) Action of the filter.
          - `creation_date` - (String) Time stamp for create route filter.
          - `direction` - (String) Direction of the filter.
          - `error` - (String) Error description.
          - `ge` - (Integer) The minimum matching length of the prefix-set.
          - `index` - (Integer) Priority or order of the filter.
          - `le` - (Integer) The maximum matching length of the prefix-set.
          - `prefix` - (String) IP prefix representing an address and mask length of the prefix-set.
          - `route_filter_id` - (String) Route filter ID.
          - `state` - (String) Status of the route filter.
      - `name` - (String) User defined name.
      - `peer_interface_id` - (String) Peer interface id.
      - `state` - (String) Status of the network peer.
      - `type` - (String) Type of the peer network.
      - `updated_date` - (String) Time stamp for update network peer.
      - `vlan` - (Integer) A vlan configured at the customer network.
