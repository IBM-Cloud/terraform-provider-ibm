---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_public_network"
description: |-
  Manages a public network in the IBM Power Virtual Server cloud.
---

# ibm_pi_public_network
Retrieve the details about a public network that is used for your Power Systems Virtual Server instance. For more information, about public network in IBM power virutal server, see [adding or removing a public network
](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-modifying-server#adding-removing-network).

## Example usage

```terraform
data "ibm_pi_public_network" "ds_public_network" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

**Note**
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`

 Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  

## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_network_name` - (Deprecated, string) The name of the network.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The ID of the network.
- `name` - (String) The name of the network.
- `type` - (String) The type of VLAN that the network is connected to.
- `vlan_id` - (String) The ID of the VLAN that the network is connected to.
