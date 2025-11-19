---
layout: "ibm"
page_title: "IBM : ibm_pi_network_peer"
description: |-
  Manages network peer in Power Virtual Server cloud.
subcategory: "Power Systems"
---

# ibm_pi_network_peer

Create, update, and delete network peers with this resource for on-prem locations.

## Example Usage

The following example creates a network peer for your Power Systems Virtual Server.

```terraform
resource "ibm_pi_network_peer" "network_peer_instance" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_customer_asn = 64512
  pi_customer_cidr = "192.168.91.2/30"
  pi_default_export_route_filter = "allow"
  pi_default_import_route_filter = "allow"
  pi_ibm_asn = 64512
  pi_ibm_cidr = "192.168.91.1/30"
  pi_name = "newPeerNetwork"
  pi_peer_interface_id = "031ab7da-bca6-493f-ac55-1a2a26f19160"
  pi_type = "dcnetwork_bgp"
  pi_vlan = 2000
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

## Timeouts

The `ibm_pi_network_peer` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 2 minutes) Used for creating a network peer.
- **delete** - (Default 2 minutes) Used for deleting a network peer.
- **update** - (Default 2 minutes) Used for updating a network peer.

## Argument Reference

You can specify the following arguments for this resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_customer_asn` - (Required, Integer) ASN number at customer network side.
- `pi_customer_cidr` - (Required, String) IP address used for configuring customer network interface with network subnet mask. customerCidr and ibmCidr must have matching network and subnet mask values.
- `pi_default_export_route_filter` - (Optional, String) Default action for export route filter,the default value is `allow`. Allowable values are: `allow`, `deny`.
- `pi_default_import_route_filter` - (Optional, String) Default action for import route filter, the default value is `allow`. Allowable values are: `allow`, `deny`.
- `pi_ibm_asn` - (Required, Integer) ASN number at IBM PowerVS side.
- `pi_ibm_cidr` - (Required, String) IP address used for configuring IBM network interface with network subnet mask. customerCidr and ibmCidr must have matching network and subnet mask values.
- `pi_name` - (Required, String) User defined name.
- `pi_peer_interface_id` - (Required, String) Peer interface id. Use datasource `ibm_pi_network_peer_interfaces` to get a list of valid peer interface id.
- `pi_type` - (Optional, String) Type of the peer network - dcnetwork_bgp: broader gateway protocol is used to share routes between two autonomous network. The default value is `dcnetwork_bgp`. Allowable values are: `dcnetwork_bgp`.
- `pi_vlan` - (Required, Integer) A vlan configured at the customer network.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `creation_date` - (String) Time stamp for create network peer.
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
- `id` - (String) The unique identifier of the network peer resource. Composed of `<pi_cloud_instance_id>/<peer_id>`
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
- `peer_id` - (String) The unique identifier of the pi_network_peer.
- `state` - (String) Status of the network peer.
- `updated_date` - (String) Time stamp for update network peer.

## Import

The `ibm_pi_network_peer` resource can be imported by using `pi_cloud_instance_id` and `peer_id`.

### Example

```bash
terraform import ibm_pi_network_peer.pi_network_peer 49fba6c9-23f8-40bc-9899-aca322ee7d5b/8a9b1c2d-3e4f-5g6h-7i8j-9k0l1m2n3o4p
```
