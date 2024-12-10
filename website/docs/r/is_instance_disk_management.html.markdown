---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_disk_management"
description: |-
  Manages IBM instance disk management.
---

# ibm_is_instance_disk_management
Create, update, or delete an IBM instance disk management. For more information, about instance disk management, see [managing instance storage](https://cloud.ibm.com/docs/vpc?topic=vpc-instance-storage-provisioning).

**Note:** VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bx2-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

data "ibm_is_instance" "example" {
  name        = ibm_is_instance.example.name
  private_key = file("~/.ssh/id_rsa")
  passphrase  = ""
}

data "ibm_is_instance_disk" "example" {
  instance = data.ibm_is_instance.example.id
  disk     = data.ibm_is_instance.example.disks.0.id
}

resource "ibm_is_instance_disk_management" "example" {
  instance = data.ibm_is_instance.example.id
  disks {
    name = "example-disk"
    id   = data.ibm_is_instance.example.disks.0.id
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
