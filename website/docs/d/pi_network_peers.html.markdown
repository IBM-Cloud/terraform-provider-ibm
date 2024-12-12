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
  
## Argument reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
  
## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the pi_network_peers.
- `network_peers` - (List) List of network peers.
  
   Nested schema for `network_peers`:
  - `description` - (String) Description of the network peer.
  - `id` - (String) ID of the network peer.
  - `name` - (String) Name of the network peer.
  - `type` - (String) Type of the network peer.
