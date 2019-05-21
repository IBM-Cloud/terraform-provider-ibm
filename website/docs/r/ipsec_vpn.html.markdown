---
layout: "ibm"
page_title: "IBM : ipsec_vpn"
sidebar_current: "docs-ibm-resource-ipsec-vpn"
description: |-
  Manages IBM IPSec VPN.
---

# ibm\_ipsec_vpn

Provides an IPSec VPN resource.

For additional details, see the [IBM Cloud (SoftLayer) ipsec vpn Request docs](https://softlayer.github.io/reference/datatypes/SoftLayer_Container_Product_Order_Network_Tunnel_Ipsec/)

## Example Usage

In the following example, you can create an IPSec VPN:

```hcl
resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=131}]
	phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=133}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}
```


## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center in which the IPSec VPN resides.
* `phase_one` - (Optional, map) The key-value parameters for phaseOne negotiation 
* `phase_two` - (Optional, map) The key-value parameters for phaseTwo negotiation
* `address_translation` - (Optional, map) The key-value parameters for creating an adress translation
* `Preshared_Key` - (Optional, string) A key used so that peers authenticate each other.  This key is hashed by using the phase one encryption and phase one authentication.
* `Customer_Peer_IP` - (Optional, string) The remote end of a network tunnel. This end of the network tunnel resides on an outside network and will be sending and receiving the IPSec packets.
* `internal_subnet_id` - (Optional, map) The id of the network device on which the vpn configurations have to be applied.When a private subnet is associated, the network tunnel will allow the customer (remote) network to access the private subnet.
* `remote_subnet_id` - (Optional, map) The id of the customer owned device on which the network configuration has to be applied. When a remote subnet is associated, a network tunnel will allow the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel.
* `remote_subnet` - (Optional, map) The key-value parameters for creating a customer subnet
* `service_subnet_id` - (Optional, string) The id of the service subnet which is to be associated to the network tunnel.When a service subnet is associated, a network tunnel will allow the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel.  Service subnets provide access to SoftLayer services such as the customer management portal and the SoftLayer API.

## Attribute Reference

The following attributes are exported:

* `id` - (Computed, string) The id of the IPSec VPN device that is created
* `name` - (Computed, string) The name of the IPSec VPN device that is created
* `internal_peer_ip_address` - (Computed, string) The local  end of a network tunnel. This end of the network tunnel resides on the SoftLayer networks and allows access to remote end of the tunnel to subnets on SoftLayer networks.




