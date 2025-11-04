---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_networks"
description: |-
  Retrieve information about all networks attached to a Power Systems Virtual Server instance.
---

# ibm_pi_instance_networks

Retrieve information about all networks on a Power Systems Virtual Server instance. For more information about Power Virtual Server instances, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_instance_networks" "ds_instance_networks" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_instance_id       = "cea6651a-bc0a-4438-9f8a-a0770b112ebb"
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
- `pi_instance_id` - (Required, String) The unique identifier or name of the instance.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the data source in the form <pi_cloud_instance_id>/<pi_instance_id>.
- `networks` - (List) List of networks associated with this instance.
      Nested scheme for networks:
      - `external_ip` - (String) The external IP address of the network (for pub-VLAN networks).
      - `ip_address` - (String) The IP address of the network interface.
      - `mac_address` - (String) The MAC address of the network interface.
      - `network_id` - (String) The network ID.