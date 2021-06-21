---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_disk_management"
description: |-
  Manages IBM instance disk management.
---

# ibm_is_instance_disk_management
Create, update, or delete an IBM instance disk management. For more information, about instance disk management, see [managing instance storage](https://cloud.ibm.com/docs/vpc?topic=vpc-instance-storage-provisioning).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `instance` - (Required, Forces new resource, String) The unique-identifier of the instance.
- `disks` - (Required, List) Disks that needs to be updated. Nested `disks` blocks have the following structure:

  Nested scheme for `disks`:
  - `id` - (Required, String) The unique-identifier of the instance disk.
  - `name` - (Required, String) The unique user defined name for the instance disk.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

* `id` - (String) The unique-identifier of an instance disk management.

## Import

The `ibm_is_instance_disk_management` resource can be imported byusing Instance disk management ID.

**Example**

```
$ terraform import ibm_is_instance_disk_management.example 0716-111172bb2-decc-4555-b1a6-5d128c62806c
```
