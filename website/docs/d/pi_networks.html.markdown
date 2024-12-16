---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_networks"
description: |-
  Manages networks in the IBM Power Virtual Server cloud.
---

# ibm_pi_networks

Retrieve a list of networks that you can use in your Power Systems Virtual Server instance. For more information, about power virtual server instance networks, see [setting up an IBM network install server](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-configuring-subnet).

## Example usage

```terraform
data "ibm_pi_networks" "ds_network" {
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

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `networks` - (List) List of all networks.

  Nested scheme for `networks`:
  - `access_config` - (Deprecated, String) The network communication configuration option of the network (for on-prem locations only). Use `peer_id` instead.
  - `crn` - (String) The CRN of this resource.
  - `dhcp_managed` - (Boolean) Indicates if the network DHCP Managed.
  - `href` - (String) The hyper link of a network.
  - `mtu` - (Boolean) Maximum Transmission Unit option of the network.
  - `name` - (String) The name of a network.
  - `network_id` - (String) The ID of the network.
  - `peer_id` - (String) Network peer ID (for on-prem locations only).
  - `type` - (String) The type of network.
  - `user_tags` - (List) List of user tags attached to the resource.
  - `vlan_id` - (String) The VLAN ID that the network is connected to.
