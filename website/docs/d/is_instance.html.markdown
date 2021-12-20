---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance"
description: |-
  Manages IBM Cloud virtual server instance.
---

# ibm_is_instance
Retrieve information of an existing IBM Cloud virtual server instance  as a read-only data source. For more information, about managing VPC instance, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).

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

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the Virtual Servers for VPC instance that you want to retrieve.
- `private_key` - (Optional, String) The private key of an SSH key that you want to add to your Virtual Servers for VPC instance during creation in PEM format. It is used to decrypt the default password of the Windows administrator for the virtual server instance if the image is used of type `windows`.
- `passphrase` - (Optional, String) The passphrase that you used when you created your SSH key. If you did not enter a passphrase when you created the SSH key, do not provide this input parameter.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 
- availability_policy - (List) The availability policy for this virtual server instance

  Nested scheme for `availability_policy`:
  - `host_failure`- (String) The action to perform if the compute host experiences a failure. 
    - `restart` - Automatically restart the virtual server instance after host failure
    - `stop` -  Leave the virtual server instance stopped after host failure
- `bandwidth` - (Integer) The total bandwidth (in megabits per second) shared across the instance's network interfaces and storage volumes
- `boot_volume` - (List of Objects) A list of boot volumes that were created for the instance.

  Nested scheme for `boot_volume`:
  - `id` - (String) The ID of the boot volume attachment.
  - `name` - (String) The name of the boot volume.
  - `device` - (String) The name of the device that is associated with the boot volume.
  - `volume_id` - (String) The ID of the volume that is associated with the boot volume attachment.
  - `volume_crn` - (String) The CRN of the volume that is associated with the boot volume attachment.
- `crn` - (String) The CRN of the instance.
- `disks` - (List) Collection of the instance's disks. Nested `disks` blocks has the following structure:

  Nested scheme for `disks`:
  - `created_at` - (Timestamp) The creation date and time of the disk.
  - `href` - (String) The URL for this instance disk.
  - `id` - (String) The unique identifier for this instance disk.
  - `interface_type` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
  - `name` - (String) The user-defined name for this disk.
  - `resource_type` - (String) The resource type.
  - `size` - (String) The size of the disk in GB.
- `gpu`- (List) A list of graphics processing units that are allocated to the instance.

  Nested scheme for `gpu`:
  - `count`- (Integer) The number of GPUs that are allocated to the instance.
  - `manufacture` - (String) The manufacturer of the GPU.
  - `memory`- (Integer) The amount of memory that was allocated to the GPU.
  - `model` - (String) The model of the GPU. 
- `id` - (String) The ID that was assigned to the Virtual Servers for VPC instance.
- `image` - (String) The ID of the virtual server image that is used in the instance.
- `keys`- (List) A list of SSH keys that were added to the instance during creation.

  Nested scheme for `keys`:
  - `id` - (String) The ID of the SSH key.
  - `name` - (String) The name of the SSH key that you entered when you uploaded the key to IBM Cloud.
- `memory`- (Integer) The amount of memory that was allocated to the instance.
- `network_interfaces`- (List) A list of more network interfaces that the instance uses.

  Nested scheme for `network_interfaces`:
  - `id` - (String) The ID of the more network interface.
  - `name` - (String) The name of the more network interface.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses.
  - `subnet` - (String) The ID of the subnet that is used in the more network interface.
  - `security_groups` (List)A list of security groups that were created for the interface.
- `password` - (String) The password that you can use to access your instance.
- `placement_target`- (List) The placement restrictions for the virtual server instance.

  Nested scheme for `placement_target`: 
  - `crn` - (String) The CRN for this placement target resource.
  - `deleted` - (String) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    - `more_info` -  (String) Link to documentation about deleted resources. 
  - `href` - (String) The URL for this placement target resource.
  - `id` - (String) The unique identifier for this placement target resource.
  - `name` - (String) The unique user-defined name for this placement target resource. If unspecified, the name will be a hyphenated list of randomly-selected words.
  - `resource_type` - (String) The type of resource referenced.
- `primary_network_interface`- (List) A list of primary network interfaces that were created for the instance. 

  Nested scheme for `primary_network_interface`:
  - `id` - (String) The ID of the primary network interface.
  - `name` - (String) The name of the primary network interface.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses.
  - `subnet` - (String) The ID of the subnet that is used in the primary network interface.
  - `security_groups` (List)A list of security groups that were created for the interface.
- `resource_controller_url` - (String) The URL of the IBM Cloud dashboard that you can use to see details for your instance.  
- `resource_group` - (String) The name of the resource group where the instance was created.
- `status` - (String) The status of the instance.
- `status_reasons` - (List) Array of reasons for the current status. 
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
- `total_volume_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `total_network_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance network interfaces.
- `vpc` - (String) The ID of the VPC that the instance belongs to.
- `vcpu`- (List) A list of virtual CPUs that were allocated to the instance.

  Nested scheme for `vcpu`:
  - `architecture` - (String) The architecture of the virtual CPU.
  - `count`- (Integer) The number of virtual CPUs that are allocated to the instance.
- `volume_attachments`- (List) A list of volume attachments that were created for the instance. 

  Nested scheme for `volume_attachments`:
  - `volume_crn` - (String) The CRN of the volume that is associated with the volume attachment.
  - `id` - (String) The ID of the volume attachment.
  - `name` - (String) The name of the volume attachment.
  - `volume_id` - (String) The ID of the volume that is associated with the volume attachment.
  - `volume_name` - (String) The name of the volume that is associated with the volume attachment.
- `zone` - (String) The zone where the instance was created.
