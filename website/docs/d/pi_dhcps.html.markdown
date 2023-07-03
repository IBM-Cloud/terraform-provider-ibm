---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_dhcps"
description: |-
  Manages DHCP Servers in the Power Virtual Server cloud.
---

# ibm_pi_dhcps

Retrieve information about all DHCP Servers. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_dhcp" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) Cloud Instance ID of a PCloud Instance.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `servers` - (List) The list of all the DHCP Servers.

  Nested scheme for `servers`:
  - `dhcp_id` - (String) The ID of the DHCP Server.
  - `network` - (String) The ID of the DHCP Server private network (deprecated - replaced by `network_id`).
  - `network_id`- (String) The ID of the DHCP Server private network.
  - `network_name` - The name of the DHCP Server private network.
  - `status` - (String) The status of the DHCP Server.

**Notes**

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