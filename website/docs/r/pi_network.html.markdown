---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network"
description: |-
  Manages networks in the IBM Power Virtual Server cloud.
---

# ibm_pi_network
Create, update, or delete a network connection for your Power Systems Virtual Server instance. For more information, about power virtual server instance network, see [setting up an IBM i network install server](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-preparing-install-server).

## Example usage
The following example creates a network connection for your Power Systems Virtual Server instance.

```terraform
resource "ibm_pi_network" "power_networks" {
  count                = 1
  pi_network_name      = "power-network"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_network_type      = "vlan"
  pi_cidr              = "<Network in CIDR notation (192.168.0.0/24)>"
  pi_dns               = [<"DNS Servers">]
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

## Timeouts

The `ibm_pi_network` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating a network.
- **update** - (Default 60 minutes) Used for updating a network.
- **delete** - (Default 60 minutes) Used for deleting a network.

## Argument reference 
Review the argument references that you can specify for your resource. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_network_name` - (Required, String) The name of the network.
- `pi_network_type` - (Required, String) The type of network that you want to create, such as `pub-vlan` or `vlan`.
- `pi_dns` - (Optional, Set of String) The DNS Servers for the network. Required for `vlan` network type.
- `pi_cidr` - (Optional, String) The network CIDR. Required for `vlan` network type.
- `pi_network_jumbo` - (Optional, Bool) MTU Jumbo option of the network.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the network. The ID is composed of `<power_instance_id>/<network_id>`.
- `networkid` - (String) The unique identifier of the network.
- `vlanid` - (Integer) The ID of the VLAN that your network is attached to. 

## Import
The `ibm_pi_network` resource can be imported by using `power_instance_id` and `network_id`.

**Example**

```
$ terraform import ibm_pi_network.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
