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

resource "ibm_is_reservation" "example" {
  capacity {
    total = 5
  }
  committed_use {
    term = "one_year"
  }
  profile {
    name          = "ba2-2x8"
    resource_type = "instance_profile"
  }
  zone = "us-east-3"
  name = "reservation-name"
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bx2-2x8"
  metadata_service_enabled  = false
  
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
  }

  reservation_affinity {
    policy = "manual"
    pool {
      id = ibm_is_reservation.example.id
    }
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
- `access_tags`  - (List) Access management tags associated for instance.
- `availability_policy_host_failure` - (String) The availability policy for this virtual server instance. The action to perform if the compute host experiences a failure. 
- `bandwidth` - (Integer) The total bandwidth (in megabits per second) shared across the instance's network interfaces and storage volumes
- `boot_volume` - (List of Objects) A list of boot volumes that were created for the instance.

  Nested scheme for `boot_volume`:
  - `id` - (String) The ID of the boot volume attachment.
  - `name` - (String) The name of the boot volume.
  - `device` - (String) The name of the device that is associated with the boot volume.
  - `volume_id` - (String) The ID of the volume that is associated with the boot volume attachment.
  - `volume_crn` - (String) The CRN of the volume that is associated with the boot volume attachment.

- `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.

  Nested scheme for `catalog_offering`:
    - `offering_crn` - (String) The CRN for this catalog offering. Identifies a catalog offering by this unique property
    - `version_crn` - (String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property
    - `plan_crn` - (String) The CRN for this catalog offering version's billing plan
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
		  Nested schema for `deleted`:
        - `more_info`  - (String) Link to documentation about deleted resources.

- `cluster_network` - (List) If present, the cluster network that this virtual server instance resides in.
  Nested schema for **cluster_network**:
	- `crn` - (String) The CRN for this cluster network.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this cluster network.
	- `id` - (String) The unique identifier for this cluster network.
	- `name` - (String) The name for this cluster network. The name must not be used by another cluster network in the region.
	- `resource_type` - (String) The resource type.
- `cluster_network_attachments` - (List) The cluster network attachments for this virtual server instance.The cluster network attachments are ordered for consistent instance configuration.
  Nested schema for **cluster_network_attachments**:
	- `href` - (String) The URL for this instance cluster network attachment.
	- `id` - (String) The unique identifier for this instance cluster network attachment.
	- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
	- `resource_type` - (String) The resource type.

- `confidential_compute_mode` - (String) The confidential compute mode to use for this virtual server instance.If unspecified, the default confidential compute mode from the profile will be used. 
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
- `enable_secure_boot` - (Boolean) Indicates whether secure boot is enabled for this virtual server instance.If unspecified, the default secure boot mode from the profile will be used.
- `gpu`- (List) A list of graphics processing units that are allocated to the instance.

  Nested scheme for `gpu`:
  - `count`- (Integer) The number of GPUs that are allocated to the instance.
  - `manufacture` - (String) The manufacturer of the GPU.
  - `memory`- (Integer) The amount of memory that was allocated to the GPU.
  - `model` - (String) The model of the GPU. 
- `health_reasons` - (List) The reasons for the current health_state (if any).

    Nested scheme for `health_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this health state.
    - `message` - (String) An explanation of the reason for this health state.
    - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.
- `id` - (String) The ID that was assigned to the Virtual Servers for VPC instance.
- `image` - (String) The ID of the virtual server image that is used in the instance.
- `keys`- (List) A list of SSH keys that were added to the instance during creation.

  Nested scheme for `keys`:
  - `id` - (String) The ID of the SSH key.
  - `name` - (String) The name of the SSH key that you entered when you uploaded the key to IBM Cloud.
- `memory`- (Integer) The amount of memory that was allocated to the instance.
- `lifecycle_reasons`- (List) The reasons for the current lifecycle_state (if any).

  Nested scheme for `lifecycle_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle state.
    - `message` - (String) An explanation of the reason for this lifecycle state.
    - `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state`- (String) The lifecycle state of the virtual server instance. 
 
  ->**lifecycle states** 
    </br>&#x2022; deleting
    </br>&#x2022; failed
    </br>&#x2022; pending
    </br>&#x2022; stable
    </br>&#x2022; suspended
    </br>&#x2022; updating
    </br>&#x2022; waiting

- `metadata_service_enabled` - (Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.

	~> **NOTE**
	`metadata_service_enabled` is deprecated and will be removed in the future. Refer `metadata_service` instead
- `metadata_service` - (List) The metadata service configuration. 

     Nested scheme for `metadata_service`:
     - `enabled` - (Boolean) Indicates whether the metadata service endpoint will be available to the virtual server instance.
     - `protocol` - (String) The communication protocol to use for the metadata service endpoint.
     - `response_hop_limit` - (Integer) The hop limit (IP time to live) for IP response packets from the metadata service.

- `network_attachments` - (List) The network attachments for this virtual server instance, including the primary network attachment.
  Nested schema for **network_attachments**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (String)
	- `primary_ip` - (List) The primary IP address of the virtual network interface for the network attachment.
	  Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List) The subnet of the virtual network interface for the network attachment.
	  Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		- `resource_type` - (String) The resource type.
  - `virtual_network_interface` - (List) The virtual network interface for this bare metal server network attachment.
  Nested schema for **virtual_network_interface**:
    - `crn` - (String) The CRN for this virtual network interface.
    - `href` - (String) The URL for this virtual network interface.
    - `id` - (String) The unique identifier for this virtual network interface.
    - `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
    - `resource_type` - (String) The resource type.
    
- `network_interfaces`- (List) A list of more network interfaces that the instance uses.

  Nested scheme for `network_interfaces`:
  - `id` - (String) The ID of the more network interface.
  - `name` - (String) The name of the more network interface.
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `address` - (String) The IP address of the reserved IP. Same as `network_interfaces.[].primary_ipv4_address`
      - `href`- (String) The URL for this reserved IP
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type`- (String) The resource type.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
  - `subnet` - (String) The ID of the subnet that is used in the more network interface.
  - `security_groups` (List)A list of security groups that were created for the interface.
- `numa_count` - (Integer) The number of NUMA nodes this virtual server instance is provisioned on. This property may be absent if the instance's `status` is not `running`.
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
- `primary_network_attachment` - (List) The primary network attachment for this virtual server instance.
  Nested schema for **primary_network_attachment**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (String)
	- `primary_ip` - (List) The primary IP address of the virtual network interface for the network attachment.
	  Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List) The subnet of the virtual network interface for the network attachment.
	  Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		- `resource_type` - (String) The resource type.
  - `virtual_network_interface` - (List) The virtual network interface for this bare metal server network attachment.
  Nested schema for **virtual_network_interface**:
    - `crn` - (String) The CRN for this virtual network interface.
    - `href` - (String) The URL for this virtual network interface.
    - `id` - (String) The unique identifier for this virtual network interface.
    - `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
    - `resource_type` - (String) The resource type.
    

- `primary_network_interface`- (List) A list of primary network interfaces that were created for the instance. 

  Nested scheme for `primary_network_interface`:
  - `id` - (String) The ID of the primary network interface.
  - `name` - (String) The name of the primary network interface.
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
  
      Nested scheme for `primary_ip`:
      - `address` - (String) The IP address of the reserved IP. Same as `primary_ipv4_address`
      - `href`- (String) The URL for this reserved IP
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type`- (String) The resource type.
  - `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
  - `subnet` - (String) The ID of the subnet that is used in the primary network interface.
  - `security_groups` (List)A list of security groups that were created for the interface.
- `reservation`- (List) The reservation used by this virtual server instance. 

  Nested scheme for `reservation`:
  - `crn` - (String) The CRN for this reservation.
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
        
      Nested `deleted` blocks have the following structure: 
      - `more_info` - (String) Link to documentation about deleted resources.
  - `href` - (String) The URL for this reservation.
  - `id` - (String) The unique identifier for this reservation.
  - `name` - (string) The name for this reservation. The name is unique across all reservations in the region.
  - `resource_type` - (string) The resource type.
- `reservation_affinity`- (List) The instance reservation affinity. 

  Nested scheme for `reservation_affinity`:
  - `policy` - (String) The reservation affinity policy to use for this virtual server instance.
  - `pool` - (List) The pool of reservations available for use by this virtual server instance.
        
    Nested `pool` blocks have the following structure: 
    - `crn` - (String) The CRN for this reservation.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested `deleted` blocks have the following structure:
      - `more_info` - (String) Link to documentation about deleted resources. 
    - `href` - (String) The URL for this reservation.
    - `id` - (String) The unique identifier for this reservation.
    - `name` - (string) The name for this reservation. The name is unique across all reservations in the region.
    - `resource_type` - (string) The resource type.
- `resource_controller_url` - (String) The URL of the IBM Cloud dashboard that you can use to see details for your instance.  
- `resource_group` - (String) The resource group id, where the instance was created.
- `status` - (String) The status of the instance.
- `status_reasons` - (List) Array of reasons for the current status. 
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
  - `more_info` - (String) Link to documentation about this status reason
- `total_volume_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `total_network_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance network interfaces.
- `vpc` - (String) The ID of the VPC that the instance belongs to.
- `vcpu` - (List) The virtual server instance VCPU configuration.
  Nested schema for **vcpu**:
	- `architecture` - (String) The VCPU architecture.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `amd64`, `s390x`.
	- `burst` - (List)
	  Nested schema for **burst**:
		- `limit` - (Integer) The maximum percentage the virtual server instance will exceed its allocated share of VCPU time.The maximum value for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `200`. The maximum value is `800`. The minimum value is `100`.
	- `count` - (Integer) The number of VCPUs assigned.
	- `manufacturer` - (String) The VCPU manufacturer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Allowable values are: `amd`, `ibm`, `intel`. 
	- `percentage` - (Integer) The percentage of VCPU time allocated to the virtual server instance.The virtual server instance `vcpu.percentage` will be `100` when:- The virtual server instance `placement_target` is a dedicated host or dedicated  host group.- The virtual server instance `reservation_affinity.policy` is `disabled`. The maximum value is `100`. The minimum value is `1`.
- `volume_attachments`- (List) A list of volume attachments that were created for the instance. 

  Nested scheme for `volume_attachments`:
  - `volume_crn` - (String) The CRN of the volume that is associated with the volume attachment.
  - `id` - (String) The ID of the volume attachment.
  - `name` - (String) The name of the volume attachment.
  - `volume_id` - (String) The ID of the volume that is associated with the volume attachment.
  - `volume_name` - (String) The name of the volume that is associated with the volume attachment.
- `volume_bandwidth_qos_mode` - (String) The volume bandwidth QoS mode to use for this virtual server instance.  
- `zone` - (String) The zone where the instance was created.
