---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_disk"
description: |-
  Get information about InstanceDisk
---

# ibm\_is_instance_disk

Provides a read-only data source for InstanceDisk. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

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
```

## Argument Reference

The following arguments are supported:

* `instance` - (Required, string) The instance identifier.
* `disk` - (Required, string) The instance disk identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the InstanceDisk.
* `created_at` - The date and time that the disk was created.

* `href` - The URL for this instance disk.

* `interface_type` - The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.

* `name` - The user-defined name for this disk.

* `resource_type` - The resource type.

* `size` - The size of the disk in GB (gigabytes).

