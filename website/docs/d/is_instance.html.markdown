---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance"
description: |-
  Manages IBM Cloud virtual server instance.
---

# ibm\_is_instance

Import the details of an existing IBM Cloud virtual server instance  as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


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

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name for this virtual server instance .
* `private_key` - (Optional, string) The private key of the ssh key used in the creation of virtual server instance in PEM Format. It is used to decrypt Windows administrator default password for the virtual server instance if the image used is of type `windows`.
* `passphrase` - (Optional, string) The passphrase used in the creation of encrypted ssh key pair. If non encrypted ssh key pair is used in the creation of the virtual server instance, this field can be omitted.

## Attribute Reference

The following attributes can be exported:

* `id` - The id of the instance.
* `memory` - Memory of the instance.
* `status` - Status of the instance.
* `image` - Image used in the instance.
* `zone` - zone of the instance.
* `vpc` - vpc id of the instance.
* `resource_group` - resource group id of the instance.
* `disks` - Collection of the instance's disks. Nested `disks` blocks have the following structure:
	* `created_at` - The date and time that the disk was created.
	* `href` - The URL for this instance disk.
	* `id` - The unique identifier for this instance disk.
	* `interface_type` - The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	* `name` - The user-defined name for this disk.
	* `resource_type` - The resource type.
	* `size` - The size of the disk in GB (gigabytes).
* `vcpu` - A nested block describing the VCPU configuration of this instance.
Nested `vcpu` blocks have the following structure:
  * `architecture` - The architecture of the instance.
  * `count` - The number of VCPUs assigned to the instance.
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
  * `subnet` -  ID of the subnet.
  * `security_groups` -  List of security groups.
  * `primary_ipv4_address` - The primary IPv4 address.
* `network_interfaces` - A nested block describing the additional network interface of this instance.
Nested `network_interfaces` blocks have the following structure:
  * `id` - The id of the network interface.
  * `name` - The name of the network interface.
  * `subnet` -  ID of the subnet.
  * `security_groups` -  List of security groups.
  * `primary_ipv4_address` - The primary IPv4 address.
* `boot_volume` - A nested block describing the boot volume.
Nested `boot_volume` blocks have the following structure:
  * `id` -  The id of the boot volume attachment.
  * `name` - The name of the boot volume.
  * `device` -  The boot volume device Name.
  * `volume_id` - The id of the boot volume attachment's volume
  * `volume_crn` - The CRN/encryption of the boot volume attachment's volume
* `volume_attachments` - A nested block describing the volume attachments.  
Nested `volume_attachments` block have the following structure:
  * `id` - The id of the volume attachment
  * `name` -  The name of the volume attachment
  * `volume_id` - The id of the volume attachment's volume
  * `volume_name` -  The name of the volume attachment's volume
  * `volume_crn` -  The CRN of the volume attachment's volume
* `resource_controller_url` - The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.
* `password` - The password to this instance
* `keys` - A nested block describing the keys used the creation of this instance.  
Nested `keys` block have the following structure:
  * `id` - The id of the key used in this instance creation
  * `name` -  The name of the key used in this instance creation
