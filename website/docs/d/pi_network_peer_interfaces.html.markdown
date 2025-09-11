---
layout: "ibm"
page_title: "IBM : ibm_pi_network_peer_interfaces"
description: |-
  Get information about IBM Power Virtual Server network peer interfaces
subcategory: "Power Systems"
---

# ibm_pi_network_peer_interfaces

Provides a read-only data source to retrieve information about pi_network_peer_interfaces for on-prem locations.

## Example Usage

```terraform
data "ibm_pi_network_peer_interfaces" "pi_network_peer_interfaces" {
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

- `id` - (String) The unique identifier of the pi_network_peer_interfaces.
- `peer_interfaces` - (List) List of peer interfaces.
  
  Nested schema for `peer_interfaces`:
  - `device_id` - (String) Device ID of the peer interface.
  - `name` - (String) Peer interface name.
  - `peer_interface_id` - (String) Peer interface ID.
  - `peer_type` - (String) Type of peer interface.
  - `port_id` - (String) Port ID of the peer interface.
