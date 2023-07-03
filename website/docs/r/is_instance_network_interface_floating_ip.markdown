---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_network_interface_floating_ip"
description: |-
  Manages IBM floating IP on virtual server instance network interface.
---

# ibm\_is_instance_network_interface_floating_ip
Associating an existing floating IP address with an instance network interface. You can use the floating IP address to access your server from the public network, independent of whether the subnet is attached to a public gateway. For more information, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example shows how to create a Virtual server instance for VPC, create a floating IP address and then associate it to a network interface on the server.

```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-3"
  ipv4_cidr_block = "10.240.129.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_floating_ip" "example" {
  name    = "example-fip1"
  zone    = "us-south-3"
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "a7a0626c-f97e-4180-afbe-0331ec62f32a"
  profile = "bx2-2x8"
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_network_interface_floating_ip" "example" {
  instance          = ibm_is_instance.example.id
  network_interface = ibm_is_instance.example.primary_network_interface[0].id
  floating_ip       = ibm_is_floating_ip.example.id
}

```

## Timeouts
The `ibm_is_instance_network_interface_floating_ip` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the floating IP address is considered `failed` if no response is received for 10 minutes. 
- **delete**: The deletion of the floating IP address is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `instance` - (Required, String) Instance identifier. 
- `floating_ip` - (Required, String) The unique identifier for a floating IP to associate with the network interface associated with the virtual server instance
- `network_interface` - (Required, String) The unique identifier for a  network interface associated with the virtual server instance.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The floating IP address.
- `crn` - (String) The CRN for this floating IP.
- `id` - (String) The unique identifier of the resource, combination of vsi ID, network interface ID, floating ip ID `<vsi_id>/<vsi_network_interface_id>/<floating_ip_id>`.
- `status` - (String) Provisioning status of the floating IP address.
- `tags` - (String) The tags associated with VPC.
- `target` - (String) The ID of the network interface used to allocate the floating IP address.
- `zone` - (String) The zone name where to create the floating IP address.


## Import
The `ibm_is_instance_network_interface_floating_ip` resource can be imported by using vsi ID, network interface ID, floating ip ID.

## Syntax
```
terraform import ibm_is_instance_network_interface_floating_ip.example <vsi_id>/<vsi_network_interface_id>/<floating_ip_id> 
```

## Example

```
$ terraform import ibm_is_instance_network_interface_floating_ip.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c
```
