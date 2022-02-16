---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_network_interface_floating_ip"
description: |-
  Manages IBM floating IP on bare metal server network interface.
---

# ibm\_is_bare_metal_server_network_interface_floating_ip
Adding a floating IP address that you can associate with a Bare Metal Server. You can use the floating IP address to access your server from the public network, independent of whether the subnet is attached to a public gateway. For more information, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example shows how to create a Bare Metal Server for VPC  and associate a floating IP address to a network interface on the server.

```terraform

resource "ibm_is_vpc" "vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.vpc.id
  zone            = "us-south-3"
  ipv4_cidr_block = "10.240.129.0/24"
}

resource "ibm_is_ssh_key" "ssh" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_floating_ip" "fip" {
  name    = "testfip1"
  zone    = "us-south-3"
}

resource "ibm_is_bare_metal_server" "bms" {
  profile 			= "bx2-metal-192x768"
  name 				  = "testserver"
  image 				= "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone 				  = "us-south-3"
  keys 				  = [ibm_is_ssh_key.ssh.id]
  primary_network_interface {
    subnet     		= ibm_is_subnet.subnet.id
  }
  vpc 				  = ibm_is_vpc.vpc.id
}
resource "ibm_is_bare_metal_server_network_interface" "bms_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  
  subnet = ibm_is_subnet.subnet.id
  name   = "eth2"
  allow_ip_spoofing = true
  allowed_vlans = [101, 102]
}

resource "ibm_is_bare_metal_server_network_interface_floating_ip" "bms_nic_fip" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  network_interface = ibm_is_bare_metal_server_network_interface.bms_nic.id
  floating_ip       = ibm_is_floating_ip.fip.id
}

```

## Timeouts
The `ibm_is_bare_metal_server_network_interface_floating_ip` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the floating IP address is considered `failed` if no response is received for 10 minutes. 
- **delete**: The deletion of the floating IP address is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `bare_metal_server` - (Required, String) Bare metal server identifier. 
- `floating_ip` - (Required, String) The unique identifier for a floating IP to associate with the network interface associated with the bare metal server
- `network_interface` - (Required, String) The unique identifier for a  network interface associated with the Bare metal server.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The floating IP address.
- `crn` - (String) The CRN for this floating IP.
- `id` - (String) The unique identifier of the floating IP.
- `status` - (String) Provisioning status of the floating IP address.
- `tags` - (String) The tags associated with VPC.
- `target` - (String) The ID of the network interface used to allocate the floating IP address.
- `zone` - (String) The zone name where to create the floating IP address.


## Import
The `ibm_is_bare_metal_server_network_interface_floating_ip` resource can be imported by using bare metal server ID, network interface ID, floating ip ID.

## Syntax
```
terraform import ibm_is_bare_metal_server_network_interface_floating_ip.example <bare_metal_server_id>/<bare_metal_server_network_interface_id>/<floating_ip_id> 
```

## Example

```
$ terraform import ibm_is_bare_metal_server_network_interface_floating_ip.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c
```
