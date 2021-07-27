---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_disk"
description: |-
  Get information about instance disk.
---

# ibm_is_instance_disk
Retrieve information about an instance disk. For more information about instance disk, see [managing instance storage](https://cloud.ibm.com/docs/vpc?topic=vpc-instance-storage-provisioning).

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
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `instance` - (Required, String) The instance identifier.
- `disk` - (Required, String) The instance disk identifier.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (Timestamp) The date and time that the disk was created.
- `href` - (String) The URL for this instance disk.
- `id` - (String) The unique identifier of the instance disk.
- `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the resource on which the unexpected property value is used.
- `name` - (String) The user defined name for this disk.
- `resource_type` - (String) The resource type.
- `size` - (String) The size of the disk in GB (gigabytes).
