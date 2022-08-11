---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance"
description: |-
  Manages IBM Cloud virtual server instance.
---

# ibm_is_instance
Retrieve information of an existing IBM Cloud virtual server instance  as a read-only data source. For more information, about managing VPC instance, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

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

resource "ibm_is_image" "example" {
  name               = "example-image"
  href               = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
  operating_system   = "ubuntu-16-04-amd64"
  encrypted_data_key = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
  encryption_key     = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"

}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bc1-2x8"
  metadata_service_enabled  = false
  
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
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the Virtual Servers for VPC instance that you want to retrieve.
- `private_key` - (Optional, String) The private key of an SSH key that you want to add to your Virtual Servers for VPC instance during creation in PEM format. It is used to decrypt the default password of the Windows administrator for the virtual server instance if the image is used of type `windows`.
- `passphrase` - (Optional, String) The passphrase that you used when you created your SSH key. If you did not enter a passphrase when you created the SSH key, do not provide this input parameter.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 
- `availability_policy_host_failure` - (String) The availability policy for this virtual server instance. The action to perform if the compute host experiences a failure. 
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
- `metadata_service_enabled` - (Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.
- `network_interfaces`- (List) A list of more network interfaces that the instance uses.

  Nested scheme for `network_interfaces`:
  - `id` - (String) The ID of the more network interface.
  - `name` - (String) The name of the more network interface.
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `href`- (String) The URL for this reserved IP
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type`- (String) The resource type.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
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
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
  
      Nested scheme for `primary_ip`:
      - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `href`- (String) The URL for this reserved IP
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type`- (String) The resource type.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
  - `subnet` - (String) The ID of the subnet that is used in the primary network interface.
  - `security_groups` (List)A list of security groups that were created for the interface.
- `resource_controller_url` - (String) The URL of the IBM Cloud dashboard that you can use to see details for your instance.  
- `resource_group` - (String) The name of the resource group where the instance was created.
- `status` - (String) The status of the instance.
- `status_reasons` - (List) Array of reasons for the current status. 
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
  - `more_info` - (String) Link to documentation about this status reason
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
