---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_disk_management"
description: |-
  Manages IBM Instance Disk Management.
---

# ibm\_is_instance_disk_management

Provides a resource for Instance Disk Management. This allows Instance disk names to be updated

## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "a7a0626c-f97e-4180-afbe-0331ec62f32a"
  profile = "bc1-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
}

data "ibm_is_instance" "ds_instance" {
  name        = "${ibm_is_instance.testacc_instance.name}"
  private_key = file("~/.ssh/id_rsa")
  passphrase  = ""
}

data "is_instance_disk" "is_instance_disk" {
  instance = data.ibm_is_instance.ds_instance.id
  disk = data.ibm_is_instance.ds_instance.disks.0.id
}

resource "ibm_is_instance_disk_management" "disks"{
  instance = data.ibm_is_instance.ds_instance.id
  disks {
    name = "mydisk01"
    id = data.ibm_is_instance.ds_instance.disks.0.id
  }
}
```

## Argument Reference

The following arguments are supported:


* `instance` - (Required, string, ForceNew) The unique-identifier of the instance
* `disks` - (Required, string) Disks that needs to be updated. Nested `disks` blocks have the following structure:
	* `id` - (Required, string) The unique-identifier of the instance disk.
	* `name` - (Required, string) The unique user defined name for the instance disk

## Attribute Reference

The following attributes are exported:

* `id` - The unique-identifier of the Instance disk management

## Import

ibm_is_instance_disk_management can be imported using Instance disk management ID, eg

```
$ terraform import ibm_is_instance_disk_management.example 0716-1c372bb2-decc-4555-b1a6-5d128c62806c
```