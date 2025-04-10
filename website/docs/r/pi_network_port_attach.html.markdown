---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network_port_attach"
description: |-
  Manages an Network Port Attachments in the Power Virtual Server Cloud.
---

# ibm_pi_network_port_attach

~> This resource is deprecated and will be removed in the next major version. Use `ibm_pi_network_interface` resource instead.

Attaches a network port to a Power Systems Virtual Server instance. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

In the following example, you can create an network_port_attach resource:

```terraform
resource "ibm_pi_network_port_attach" "test-network-port-attach" {
    pi_cloud_instance_id        = "<value of the cloud_instance_id>"
    pi_instance_id              = "<pvm instance id>"
    pi_network_name             = "<network name>"
    pi_network_port_description = "<description>"
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
  
## Timeouts

ibm_pi_network_port_attach provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for attaching a network port.
- **delete** - (Default 60 minutes) Used for detaching a network port.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Required, String) Instance id to attach the network port to.
- `pi_network_name` - (Required, String) The network ID or name.
- `pi_network_port_description` - (Optional, String) The description for the Network Port.
- `pi_network_port_ipaddress` - (Optional, String) The requested ip address of this port.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the instance. The ID is composed of `<pi_cloud_instance_id>/<pi_network_name>/<network_port_id>`.
- `mac_address` - (String) The MAC address of the instance.
- `macaddress` - (String) The MAC address of the instance. Deprecated please use `mac_address` instead.
- `network_port_id` - (String) The ID of the port.
- `public_ip` - (String) The public IP associated with the port.
- `status` - (String) The status of the port.

## Import

The `ibm_pi_network_port` resource can be imported by using `power_instance_id`, `pi_network_name`  and `network_port_id`.

### Example

```bash
terraform import ibm_pi_network_port_attach.example d7bec597-4726-451f-8a63-e62e6f19c32c/pi_network_name/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
