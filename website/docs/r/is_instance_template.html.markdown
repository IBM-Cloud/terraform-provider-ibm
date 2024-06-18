---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_template"
description: |-
  Manages IBM VPC instance template.
---

# ibm_is_instance_template
Create, update, or delete an instance template on VPC. For more information, about instance template, see [managing an instance template](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-template).

~>**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example creates an instance template in a VPC generation-2 infrastructure.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "SSH KEY"
}

resource "ibm_resource_group" "example" {
  name = "example-resource-group"
}

resource "ibm_is_dedicated_host_group" "example" {
  family         = "compute"
  class          = "cx2"
  zone           = "us-south-1"
  name           = "example-dedicated-host-group-01"
  resource_group = ibm_resource_group.example.id
}

resource "ibm_is_dedicated_host" "example" {
  profile        = "bx2d-host-152x608"
  name           = "example-dedicated-host"
  host_group     = ibm_is_dedicated_host_group.example.id
  resource_group = ibm_resource_group.example.id
}

resource "ibm_is_volume" "example" {
  name           = "example-data-vol1"
  resource_group = ibm_resource_group.example.id
  zone           = "us-south-2"

  profile  = "general-purpose"
  capacity = 50
}


// Create a new volume with the volume attachment. This template format can be used with instance groups
resource "ibm_is_instance_template" "example" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.example.id]

  boot_volume {
    name                             = "example-boot-volume"
    delete_volume_on_instance_delete = true
  }
  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "example-volume-att-01"
    volume_prototype {
      iops     = 3000
      profile  = "custom"
      capacity = 200
    }
  }
}

// Template with volume attachment that attaches existing storage volume. This template cannot be used with instance groups
resource "ibm_is_instance_template" "example1" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"
  
  primary_network_interface {
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.example.id]

  boot_volume {
    name                             = "example-boot-volume"
    delete_volume_on_instance_delete = true
  }
  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "example-volume-att"
    volume                           = ibm_is_volume.example.id
  }
}

resource "ibm_is_instance_template" "example3" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = true
  }

  dedicated_host_group = ibm_is_dedicated_host_group.example.id
  vpc                  = ibm_is_vpc.example.id
  zone                 = "us-south-2"
  keys                 = [ibm_is_ssh_key.example.id]

  boot_volume {
    name                             = "example-boot-volume"
    delete_volume_on_instance_delete = true
  }
}

resource "ibm_is_instance_template" "example4" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = true
  }

  dedicated_host = ibm_is_dedicated_host.example.id
  vpc            = ibm_is_vpc.vpc2.id
  zone           = "us-south-2"
  keys           = [ibm_is_ssh_key.example.id]

  boot_volume {
    name                             = "example-boot-volume"
    delete_volume_on_instance_delete = true
  }
}

```

```
resource "ibm_is_instance_template" "example4" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"

  metadata_service {
    enabled = true
    protocol = "https"
    response_hop_limit = 5
  }
  
  primary_network_interface {
    subnet            = ibm_is_subnet.example.id
    allow_ip_spoofing = true
  }

  dedicated_host = ibm_is_dedicated_host.example.id
  vpc            = ibm_is_vpc.vpc2.id
  zone           = "us-south-2"
  keys           = [ibm_is_ssh_key.example.id]

  boot_volume {
    name                             = "example-boot-volume"
    delete_volume_on_instance_delete = true
  }
}
```
## Argument reference
Review the argument references that you can specify for your resource. 
- `availability_policy_host_failure` - (Optional, String) The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure. Supported values are `restart` and `stop`.

- `boot_volume` - (Optional, List) A nested block describes the boot volume configuration for the template.

  Nested scheme for `boot_volume`:
	- `delete_volume_on_instance_delete` - (Optional, Bool) You can configure to delete the boot volume based on instance deletion.
	- `encryption` - (Optional, String) The encryption key CRN to encrypt the boot volume attached.
	- `name` - (Optional, String) The name of the boot volume.
  - `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)

- `catalog_offering` - (Optional, Forces new resource, List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.

  ~> **Note:**
  `catalog_offering` conflicts with `image`

  Nested scheme for `catalog_offering`:
    - `offering_crn` - (Optional, Force new resource, String) The CRN for this catalog offering. Identifies a catalog offering by this unique property. Conflicts with `catalog_offering.0.version_crn`
    - `version_crn` - (Optional, Force new resource, String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property. Conflicts with `catalog_offering.0.offering_crn`
   
- `dedicated_host` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host where the instance is placed.

  ~>**Note:** 
    only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used

- `dedicated_host_group` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host group where the instance is placed.

  ~>**Note:** 
    only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used

- `default_trusted_profile_auto_link` - (Optional, Forces new resource, Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default value : **true**
- `default_trusted_profile_target` - (Optional, Forces new resource, String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
- `image` - (Required, String) The ID of the image to create the template. Conflicts when using `catalog_offering`

  ~> **Note:**
  `image` conflicts with `catalog_offering`

- `keys` - (Required, List) List of SSH key IDs used to allow log in user to the instances.
- `metadata_service_enabled` - (Optional, Forces new resource, Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.  Default value : **false**

  ~> **NOTE**
  `metadata_service_enabled` is deprecated and will be removed in the future. Use `metadata_service` instead
- `metadata_service` - (Optional, List) The metadata service configuration. 

  Nested scheme for `metadata_service`:
  - `enabled` - (Optional, Forces new resource, Boolean) Indicates whether the metadata service endpoint will be available to the virtual server instance.  Default is **false**
  - `protocol` - (Optional, Forces new resource, String) The communication protocol to use for the metadata service endpoint. Applies only when the metadata service is enabled. Default is **http**
  - `response_hop_limit` - (Optional, Forces new resource, Integer) The hop limit (IP time to live) for IP response packets from the metadata service. Default is **1**
- `name` - (Optional, String) The name of the instance template.
- `placement_group` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the placement group where the instance is placed.

  ~>**Note:** 
    only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used
- `profile` - (Required, String) The number of instances created in the instance group.
- `primary_network_attachment` - (Optional, List) The primary network attachment for this virtual server instance.
  Nested schema for **primary_network_attachment**:
	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (Required, String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (Required, String) The name of this network attachment
  - `resource_type` - (String) The resource type.
	- `virtual_network_interface` - (Required, List(1)) The details of the virtual network interface for this network attachment. It can either accept an `id` or properties of `virtual_network_interface`
      Nested schema for **virtual_network_interface**: 
      - `id` - (Optional, List) The `id` of the virtual network interface, id conflicts with other properties of virtual network interface
      - `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
      - `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when target is deleted
      - `enable_infrastructure_nat` - (Optional, Boolean) If true: The VPC infrastructure performs any needed NAT operations and floating_ips must not have more than one floating IP. If false: Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations, allow_ip_spoofing must be false, can only be attached to a target with a resource_type of bare_metal_server_network_attachment.
      - `name` - (Optional, String) The resource type.
      - `ips` - (Optional, Array of String) Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the primary IP's subnet.
      - `resource_group` - (Optional, String) The resource type.
      - `security_groups` - (Optional, Array of String) The resource type.
      - `primary_ip` - (Required, List) The primary IP address of the virtual network interface for the network attachment.
          Nested schema for **primary_ip**:
          - `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
          - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
            Nested schema for **deleted**:
            - `more_info` - (Required, String) Link to documentation about deleted resources.
          - `href` - (Required, String) The URL for this reserved IP.
          - `id` - (Required, String) The unique identifier for this reserved IP.
          - `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
          - `resource_type` - (String) The resource type.
      - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

            ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
      - `resource_type` - (String) The resource type.
      - `subnet` - (Required, String) The subnet id of the virtual network interface for the network attachment.

- `primary_network_interfaces` (Required, List) A nested block describes the primary network interface for the template.

  Nested scheme for `primary_network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the primary network interface.
  - `security_groups`- (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Force new resource, String) The VPC subnet to assign to the interface.

- `network_attachments` - (Optional, List) The network attachments for this virtual server instance, including the primary network attachment. Adding and removing of network attachments must be done from the rear end to avoid unwanted differences and changes in terraform.
  Nested schema for **network_attachments**:
	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (Required, String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (Optional, String) Name of the attachment.
	- `virtual_network_interface` - (Required, List(1)) The details of the virtual network interface for this network attachment. It can either accept an `id` or properties of `virtual_network_interface`
      Nested schema for **virtual_network_interface**:
      - `id` - (Optional, String) The `id` of the virtual network interface, id conflicts with other properties of virtual network interface
      - `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
      - `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when target is deleted
      - `enable_infrastructure_nat` - (Optional, Boolean) If true: The VPC infrastructure performs any needed NAT operations and floating_ips must not have more than one floating IP. If false: Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations, allow_ip_spoofing must be false, can only be attached to a target with a resource_type of bare_metal_server_network_attachment.
      - `name` - (Optional, String) The resource type.
      - `ips` - (Optional, Array of String) Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the primary IP's subnet.
        ~> **NOTE** to add `ips` only existing `reserved_ip` is supported, new reserved_ip creation is not supported as it leads to unmanaged(dangling) reserved ips. Use `ibm_is_subnet_reserved_ip` to create a reserved_ip
      - `resource_group` - (Optional, String) The resource type.
      - `security_groups` - (Optional, Array of String) The resource type.
      - `primary_ip` - (Required, List) The primary IP address of the virtual network interface for the network attachment.
          Nested schema for **primary_ip**:
          - `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
          - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
            Nested schema for **deleted**:
            - `more_info` - (Required, String) Link to documentation about deleted resources.
          - `href` - (Required, String) The URL for this reserved IP.
          - `id` - (Required, String) The unique identifier for this reserved IP.
          - `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
          - `resource_type` - (String) The resource type.
      - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

            ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
      - `resource_type` - (String) The resource type.
      - `subnet` - (Required, String) The subnet id of the virtual network interface for the network attachment.

- `network_interfaces` - (Optional, List) A nested block describes the network interfaces for the template.

  Nested scheme for `network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the network interface.
  - `security_groups` - (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Forces new resource, String) The VPC subnet to assign to the interface.
- `reservation_affinity` - (Optional, List) The reservation affinity for the instance
  Nested scheme for `reservation_affinity`:
  - `policy` - (Optional, String) The reservation affinity policy to use for this virtual server instance.

     ->**policy** 
			&#x2022; disabled: Reservations will not be used
      </br>&#x2022; manual: Reservations in pool will be available for use
  - `pool` - (Optional, String) The pool of reservations available for use by this virtual server instance. Specified reservations must have a status of active, and have the same profile and zone as this virtual server instance. The pool must be empty if policy is disabled, and must not be empty if policy is manual.
    Nested scheme for `pool`:
    - `id` - The unique identifier for this reservation
- `resource_group` - (Optional, Forces new resource, String) The resource group ID.
- `total_volume_bandwidth` - (Optional, int) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `volume_attachments` - (Optional, Force new resource, List) A nested block describes the storage volume configuration for the template. 

  Nested scheme for `volume_attachments`:
	- `delete_volume_on_instance_delete`- (Required, Bool) You can configure to delete the storage volume to delete based on instance deletion.
  - `name` - (Required, String) The name of the boot volume.
	- `volume` - (Optional, Forces new resource, String) The storage volume ID created in VPC.
  - `volume_prototype` - (Optional, Forces new resource, List)

      Nested scheme for `volume_prototype`:
      - `capacity` - (Required, Forces new resource, Integer) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
      - `encryption_key` - (Optional, Forces new resource, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for the resource.
      - `iops` - (Optional, Forces new resource, Integer) The maximum input and output operations per second (IOPS) for the volume.
      - `profile` - (Required, Forces new resource, String) The global unique name for the volume profile to use for the volume. Allowed values areFor more information, about volume profiles, see [volume profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles)
      - `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
      
      ~>**Note:** 
      `volume_attachments` provides either `volume` with a storage volume ID, or `volume_prototype` to create a new volume. If you plan to use this template with instance group, provide the `volume_prototype`. Instance group does not support template with existing storage volume IDs.
- `vpc` - (Required, String) The VPC ID that the instance templates needs to be created.
- `user_data` -  (Optional, String) The user data provided for the instance.
- `zone` - (Required, String) The name of the zone.

## Attribute reference
In addition to all arguments listed, you can access the following attribute references after your resource is created.

- `crn` - (String) The CRN for this instance template.
- `id` - (String) The ID of an instance template.
- `placement_target` - (List) The placement restrictions to use for the virtual server instance.
  Nested scheme for `placement_target`:
    - `crn` - (String) The unique identifier for this placement target.
    - `href` - (String) The CRN for this placement target.
    - `id` - (String) The URL for this placement target.

## Import
The `ibm_is_instance_template` resource can be imported by using instance template ID.

**Example**

```
$ terraform import ibm_is_instance_template.template r006-14140f94-fcc4-1349-96e7-a71212125115
```
