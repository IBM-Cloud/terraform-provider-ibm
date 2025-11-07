---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_network"
description: |-
  Retrieve information about a network attached to a Power Systems Virtual Server instance.
---

# ibm_pi_instance_network

Retrieve information about a specific network on a Power Systems Virtual Server instance. For more information about Power Virtual Server instances, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_instance_network" "ds_instance_network" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_instance_id       = "cea6651a-bc0a-4438-9f8a-a0770b112ebb"
  pi_network_id        = "52b7c0b1-1df1-495a-9c2d-8b7a6c5ef9aa"
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
- `pi_instance_id` - (Required, String) The PVM instance id.
- `pi_network_id` - (Required, String) The unique identifier or name of the instance.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the data source in the form <pi_cloud_instance_id>/<pi_instance_id>/<pi_network_id>.
- `external_ip` - (String) TThe external IP address of the network (for pub-VLAN networks).
- `ip_address` - (String) The IP address of the network interface.
- `mac_address` - (String) The MAC address of the network interface.
- `network_id` - (String) The network ID.
- `network_interface_id` - (String) ID of the network interface.
- `network_name` - (String) The network name.
- `network_security_group_ids` - (List) IDs of the network security groups that the network interface is a member of.
- `network_security_groups_href` - (List) Links to the network security groups that the network interface is a member of.
- `type` - (String) The address type (for example, fixed or dynamic).
- `href` - (String) Link to this PVM instance network.
- `version` - (Float) Version of the network information.