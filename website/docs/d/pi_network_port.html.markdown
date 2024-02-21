---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network_port"
description: |-
  Manages an Network Port in the Power Virtual Server Cloud. 
---

# ibm_pi_network_port
Retrieve information about a network port in the Power Virtual Server Cloud. For more information, about networks in IBM power virtual server, see [adding or removing a public network](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-modifying-server#adding-removing-network).

## Example usage
```terraform
data "ibm_pi_network_port" "test-network-port" {
    pi_network_name             = "Zone1-CFN"
    pi_cloud_instance_id        = "51e1879c-bcbe-4ee1-a008-49cdba0eaf60"
}
```

**Notes**
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
- `pi_network_name` - (Required, String) The unique identifier or name of a network.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.

- `network_ports` - (List) List of all in use network ports for a network.

  Nested scheme for `network_ports`:
  - `description` - (String) The description for the network port.
  - `href` - (String) Network port href.
  - `ipaddress` - (String) The IP address of the port.
  - `macaddress` - (String) The MAC address of the port.
  - `portid` - (String) The ID of the port.
  - `public_ip`- (String) The public IP associated with the port.
  - `status` - (String) The status of the port.
