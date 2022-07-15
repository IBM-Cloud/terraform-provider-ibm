---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_dhcp"
description: |-
  Manages DHCP Server in the Power Virtual Server cloud.
---

# ibm_pi_dhcp

Retrieve information about a DHCP Server. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_dhcp" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_dhcp_id = "0e48e1be-9f54-4a67-ba55-7e31ce98b65a"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) Cloud Instance ID of a PCloud Instance.
- `pi_dhcp_id` - (Required, String) The ID of the DHCP Server.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID of the DHCP Server.
- `leases` - (List) The list of DHCP Server PVM Instance leases.
  Nested scheme for `leases`:
  - `instance_ip` - (String) The IP of the PVM Instance.
  - `instance_mac` - (String) The MAC Address of the PVM Instance.
- `network` - (String) The ID of the DHCP Server private network (deprecated - replaced by `network_id`).
- `network_id`- (String) The ID of the DHCP Server private network.
- `network_name` - The name of the DHCP Server private network.
- `status` - (String) The status of the DHCP Server.

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