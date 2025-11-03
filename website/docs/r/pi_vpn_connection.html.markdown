---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_vpn_connection"
description: |-
  Manages IBM VPN Connections in the Power Virtual Server cloud.
---

# ibm_pi_vpn_connection

~> This resource is deprecated and will be removed in the next major version. This resource has reached end of life.

Update or delete a VPN connection. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

~> **Deprecated:**
  Create VPN connection is deprecated and no longer supported. Existing `pi_vpn_connection` will still have support for `update` and `delete`. See [a new method for creating a VPN connection](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-VPN-connections)
  
## Example Usage

The following example creates a VPN Connection.

```terraform
  resource "ibm_pi_vpn_connection" "example" {
    pi_cloud_instance_id    = "<value of the cloud_instance_id>"
    pi_vpn_connection_name  = "test"
    pi_ike_policy_id        = ibm_pi_ike_policy.policy.policy_id
    pi_ipsec_policy_id      = ibm_pi_ipsec_policy.policy.policy_id
    pi_vpn_connection_mode  = "policy"
    pi_networks             = [ibm_pi_network.private_network1.network_id]
    pi_peer_gateway_address = "1.22.124.1"
    pi_peer_subnets         = ["107.0.0.0/24"]
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

ibm_pi_vpn_connection provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **update** - (Default 20 minutes) Used for updating VPN connection.
- **delete** - (Default 20 minutes) Used for deleting VPN connection.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_ike_policy_id` - (Required, String) Unique identifier of IKE Policy selected for this VPN Connection.
- `pi_ipsec_policy_id`- (Required, String) Unique identifier of IPSec Policy selected for this VPN Connection.
- `pi_networks` - (Required, Set of String) Set of network IDs to attach to this VPN connection.
- `pi_peer_gateway_address` - (Required, String) Peer Gateway address.
- `pi_peer_subnets`  - (Required, Set of String) Set of CIDR of peer subnets.
- `pi_vpn_connection_mode` - (Required, String) Mode used by this VPN Connection, either `policy` or `route`.
- `pi_vpn_connection_name` - (Required, String) Name of the VPN Connection.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VPN Connection. The ID is composed of `<power_instance_id>/<vpn_connection_id>`.
- `connection_id` - (String) VPN Connection ID.
- `connection_status` - (String) Status of the VPN connection.
- `dead_peer_detections` - (Map) Dead Peer Detection.

  Nested scheme for `dead_peer_detections`:
  - `action` - (String) Action to take when a Peer Gateway stops responding.
  - `interval` - (String) How often to test that the Peer Gateway is responsive.
  - `threshold` - (String) The number of attempts to connect before tearing down the connection.
- `gateway_address` - (String) Public IP address of the VPN Gateway (vSRX) attached to this VPN Connection.
- `local_gateway_address` - (String) Local Gateway address, only in `route` mode.

## Import

The `ibm_pi_vpn_connection` resource can be imported by using `power_instance_id` and `vpn_connection_id`.

### Example

```bash
terraform import ibm_pi_vpn_connection.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf451f
```
