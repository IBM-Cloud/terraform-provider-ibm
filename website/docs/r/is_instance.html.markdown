---
layout: "ibm"
page_title: "IBM : instance"
sidebar_current: "docs-ibm-resource-is-instance"
description: |-
  Manages IBM IS Instance.
---

# ibm\_is_instance

Provides a instance resource. This allows instance to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = "${ibm_is_vpc.testacc_vpc.id}"
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface = {
    port_speed = "100"
    subnet     = "${ibm_is_subnet.testacc_subnet.id}"
  }

  vpc  = "${ibm_is_vpc.testacc_vpc.id}"
  zone = "us-south-1"
  keys = ["${ibm_is_ssh_key.testacc_sshkey.id}"]

  //User can configure timeouts
  	timeouts {
      	create = "90m"
      	delete = "30m"
    }
}

```

## Timeouts

ibm_is_instance provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating Instance.
* `delete` - (Default 60 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The instance name.
* `vpc` - (Required, string) The vpc id. 
* `zone` - (Required, string) Name of the zone. 
* `profile` - (Required, string) The profile name. 
* `image` - (Optional, string) ID of the image. 
  **NOTE**: Conflicts with `boot_volume`.
* `boot_volume` - (Optional, string) ID of the boot volume. 
  **NOTE**: Conflicts with `image`.
* `keys` - (Required, list) Comma separated IDs of ssh keys. 
* `generation` - (Optional, string) Generation of the server instance. valid values are gc, gt. Defaults to gc. 
* `primary_network_interface` - (Required, list) A nested block describing the primary network interface of this instance. We can have only one primary network interface. We can add multiple network interface using [ibm_is_instance_nic](../r/is_instance_nic.html.markdown).
Nested `primary_network_interface` block have the following structure:
  * `name` - (Optional, string) The name of the network interface.
  * `port_speed` - (Required, int) Speed of the network interface.
  * `subnet` -  (Required, string) ID of the subnet.
  * `security_groups` - (Optional, string) Comma separated IDs of security groups.

* `volumes` - (Optional, list) Comma separated IDs of volumes. 
* `user-data` - (Optional, string) User data to transfer to the server instance. 

## Attribute Reference

The following attributes are exported:

* `id` - The id of the instance.
* `memory` - Memory of the instance.
* `status` - Status of the instance.
* `cpu` - A nested block describing the cpu of this instance.
Nested `cpu` blocks have the following structure:
  * `architecture` - The architecture of the instance.
  * `cores` - Number of cores of the instance.
  * `frequency` - Frequency of the instance.
* `gpu` - A nested block describing the gpu of this instance.
Nested `gpu` blocks have the following structure:
  * `cores` - The cores of the gpu.
  * `count` - Count of the gpu.
  * `manufacture` - Manufacture of the gpu.
  * `memory` - Memory of the gpu.
  * `model` - Model of the gpu.
* `primary_network_interface` - A nested block describing the primary network interface of this instance.
Nested `primary_network_interface` blocks have the following structure:
  * `id` - The id of the network interface.
  * `name` - The name of the network interface.
  * `port_speed` - Speed of the network interface.
  * `subnet` -  ID of the subnet.
  * `security_groups` -  List of security groups.
  * `primary_ipv4_address` - The primary IPv4 address.

## Import

ibm_is_instance can be imported using instanceID, eg

```
$ terraform import ibm_is_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
