---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_dhcp"
description: |-
  Manages IBM DHCP Server in the Power Virtual Server cloud.
---

# ibm_pi_dhcp

Create, update, or delete DHCP Server for your Power Systems Virtual Server instance. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example enables you to create a DHCP Server:

```terraform
resource "ibm_pi_dhcp" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cidr` - (Optional, String) The CIDR for the DHCP private network.
- `pi_cloud_connection_id` - (Optional, String) The Cloud Connection UUID to connect with the DHCP private network.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_dhcp_name` - (Optional, String) The name of the DHCP Service that will be prefixed by the DHCP identifier.
- `pi_dhcp_snat_enabled` - (Optional, Bool) Indicates if SNAT will be enabled for the DHCP service. The default value is **true**.
- `pi_dns_server` - (Optional, String) The DNS Server for the DHCP service.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `dhcp_id` - (String) The ID of the DHCP Server.
- `id` - (String) The unique identifier of the DHCP Server. The ID is composed of `<pi_cloud_instance_id>/<dhcp_id>`.
- `leases` - (List) The list of DHCP Server PVM Instance leases.
  Nested scheme for `leases`:
  - `instance_ip` - (String) The IP of the PVM Instance.
  - `instance_mac` - (String) The MAC Address of the PVM Instance.
- `network_id`- (String) The ID of the DHCP Server private network.
- `network_name` - The name of the DHCP Server private network.
- `status` - (String) The status of the DHCP Server.

## Timeouts

The `ibm_pi_dhcp` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating a DHCP Server.
- **delete** - (Default 10 minutes) Used for deleting a DHCP Server.

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

## Import

The `ibm_pi_dhcp` resource can be imported by using `pi_cloud_instance_id` and `dhcp_id`.

```bash
terraform import ibm_pi_dhcp.example d7bec597-4726-451f-8a63-e62e6f19c32c/0e48e1be-9f54-4a67-ba55-7e31ce98b65a
```
