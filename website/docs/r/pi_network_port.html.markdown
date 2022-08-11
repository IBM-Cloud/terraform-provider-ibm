---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network_port"
description: |-
  Manages an Network Port in the Power Virtual Server Cloud. A network port is equivalent to reserving an IP in the subnet.
  When the port is created the status will be "DOWN". This network port however is not attached to an instance. 
---

# ibm_pi_network_port
Creates or updates network port in the Power Virtual Server Cloud. For more information, about network in IBM power virutal server, see [adding or removing a public network
](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-modifying-server#adding-removing-network)..

## Example usage

In the following example, you can create an network_port:

```terraform
resource "ibm_pi_network_port" "test-network-port" {
    pi_network_name             = "Zone1-CFN"
    pi_cloud_instance_id        = "51e1879c-bcbe-4ee1-a008-49cdba0eaf60"
    pi_network_port_description = "IP Reserved for Oracle RAC "
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

ibm_pi_network_port provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating a network_port.
- **delete** - (Default 60 minutes) Used for deleting a network_port.

## Argument reference
Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_network_name` - (Required, String) Network ID or name.
- `pi_network_port_description` - (Optional, String) The description for the Network Port.
- `pi_network_port_ipaddress` - (Optional, String) The requested ip address of this port.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the instance. The ID is composed of `<pi_cloud_instance_id>/<power_network_port_id>/<id>`.
- `macaddress` - (String) The MAC address of the port.
- `portid` - (String) The ID of the port.
- `public_ip` - (String) The public IP associated with the port.
- `status` - (String) The status of the port.


## Import

The `ibm_pi_network_port` resource can be imported by using `power_instance_id`, `port_id` and `pi_network_name`.

**Example**

```
$ terraform import ibm_pi_network_port.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/network-name
```
