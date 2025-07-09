---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_dhcp"
description: |-
  Manages DHCP Server in the Power Virtual Server cloud.
---

# ibm_pi_dhcp

Retrieve information about a DHCP Server. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_dhcp" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_dhcp_id = "0e48e1be-9f54-4a67-ba55-7e31ce98b65a"
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
- `pi_dhcp_id` - (Required, String) ID of the DHCP Server.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `dhcp_id` - (Deprecated, String) ID of the DHCP Server.
- `leases` - (List) List of DHCP Server PVM Instance leases.
  Nested scheme for `leases`:
  - `instance_ip` - (String) IP of the PVM Instance.
  - `instance_mac` - (String) MAC Address of the PVM Instance.
- `network` - (String) ID of the DHCP Server private network (deprecated - replaced by `network_id`).
- `network_id`- (String) ID of the DHCP Server private network.
- `network_name` - (String) Name of the DHCP Server private network.
- `status` - (String) Status of the DHCP Server.
