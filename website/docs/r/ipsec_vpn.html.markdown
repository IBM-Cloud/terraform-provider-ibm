---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ipsec_vpn"
description: |-
  Manages IBM IPSec VPN.
---

# ibm_ipsec_vpn
Create, update, or delete an IPSec VPN resource. For more information, about IPSec VPN, see [setting up an IPsec VPN connection](https://cloud.ibm.com/docs/iaas-vpn?topic=iaas-vpn-setup-ipsec-vpn).

**Note**

For more information, see the [IBM Cloud (SoftLayer) IPSec VPN Request](https://softlayer.github.io/reference/datatypes/SoftLayer_Container_Product_Order_Network_Tunnel_Ipsec/)

## Example usage
In the following example, you can create an IPSec VPN:

```terraform
resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=131}]
	phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=133}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}
```


## Argument reference 
Review the argument references that you can specify for your resource. 

- `address_translation` (Optional, Map) The key-value parameters for creating an address translation.
- `Customer_Peer_IP` - (Optional, String) The remote end of a network tunnel. This end of the network tunnel resides on an outside network and be sending and receiving the IPSec packets.
- `datacenter` - (Required, String) The data center in which the IPSec VPN resides.
- `internal_subnet_id` (Optional, Map) The ID of the network device on which the VPN configurations have to be applied. When a private subnet is associated, the network tunnel will allow the customer (remote) network to access the private subnet.
- `phase_one` (Optional, Map) The key-value parameters for phase One negotiation.
- `phase_two` (Optional, Map) The key-value parameters for phase Two negotiation.
- `Preshared_Key` - (Optional, String) A key used so that peers authenticate each other. This key is hashed by using the phase one encryption and phase one authentication.
- `remote_subnet_id` (Optional, Map) The ID of the customer owned device on which the network configuration has to be applied. When a remote subnet is associated, a network tunnel allows the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel.
- `remote_subnet` (Optional, Map) The key-value parameters for creating a customer subnet.
- `service_subnet_id` - (Optional, String) The ID of the service subnet which is to be associated to the network tunnel. When a service subnet is associated, a network tunnel allows the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel. Service subnets provide access to SoftLayer services such as the customer management portal and the SoftLayer API.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The computed ID of the IPSec VPN device that is created.
- `internal_peer_ip_address` - (String) The local end of a network tunnel. This end of the network tunnel resides on the SoftLayer networks and allows access to remote end of the tunnel to subnets on SoftLayer networks.
- `name` - (String) The computed name of the IPSec VPN device that is created.
 
