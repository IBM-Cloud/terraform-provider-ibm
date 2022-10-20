---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance"
description: |-
  Manages IBM VPC instance.
---

# ibm_is_instance
Create, update, or delete a Virtual Servers for VPC instance. For more information, about managing VPC instance, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).

**Note**
- IBM Cloud terraform provider currently provides both a standalone `ibm_is_instance_network_interface` resource and a `network_interfaces` block defined in-line in the `ibm_is_instance` resource. At this time you cannot use the `network_interfaces` block inline with `ibm_is_instance` in conjunction with the standalone resource `ibm_is_instance_network_interface`. Doing so will create a conflict of network interfaces and will overwrite it.
- VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

  **provider.tf**

  ```terraform
  provider "ibm" {
    region = "eu-gb"
  }
  ```

## Example usage

### Sample for creating an instance in a VPC.

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
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bc1-2x8"
  metadata_service_enabled  = false

  boot_volume {
    encryption = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
  }

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
    primary_ipv4_address = "10.240.0.6"  // will be deprecated. Use primary_ip.[0].address
    allow_ip_spoofing = true
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
    allow_ip_spoofing = false
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
// primary_ipv4_address deprecation 
output "primary_ipv4_address" {
  # value = ibm_is_instance.example.primary_network_interface.0.primary_ipv4_address // will be deprecated in future
  value = ibm_is_instance.example.primary_network_interface.0.primary_ip.0.address // use this instead 
}

```
### Sample for creating an instance with reserved ip as primary_ip reference.

The following example shows how you can create a virtual server instance using reserved ip as the primary ip reference of the network interface

// Example to provision instance using reserved ip

```terraform
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet    = ibm_is_subnet.example.id
  name      = "example-reserved-ip1"
  address		= "${replace(ibm_is_subnet.example.ipv4_cidr_block, "0/24", "13")}"
}

resource "ibm_is_instance" "example1" {
  name    = "example-instance-reserved-ip"
  image   = ibm_is_image.example.id
  profile = "bc1-2x8"
  metadata_service_enabled  = false

  primary_network_interface {
    name   = "eth0"
    subnet = ibm_is_subnet.example.id
    primary_ip {
      reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip
    }
  }  
  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
    allow_ip_spoofing = false
    primary_ip {
      name = "example-reserved-ip1"
      auto_delete = true
      address = "${replace(ibm_is_subnet.example.ipv4_cidr_block, "0/24", "14")}"
    }
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

```

### Sample for creating an instance with custom security group rules.

The following example shows how you can create a virtual server instance with custom security group rules. Note that the security group, security group rules, and the virtual server instance must be created in a specific order to meet the dependencies of the individual resources. To force the creation in a specific order, you use the [`depends_on` parameter](https://www.terraform.io/docs/configuration/resources.html). If you do not provide this parameter, all resources are created at the same time which might lead to resource dependency errors during the provisioning of your virtual server, such as `The security group to attach to is not available`.

```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_security_group" "example" {
  name = "example-security-group"
  vpc  = ibm_is_vpc.example.id
}

resource "ibm_is_security_group_rule" "example1" {
  group      = ibm_is_security_group.example.id
  direction  = "inbound"
  remote     = "127.0.0.1"
  depends_on = [ibm_is_security_group.example]
}

resource "ibm_is_security_group_rule" "example2" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  icmp {
    code = 20
    type = 30
  }
  depends_on = [ibm_is_security_group_rule.example1]

}

resource "ibm_is_security_group_rule" "example3" {
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  udp {
    port_min = 805
    port_max = 807
  }
  depends_on = [ibm_is_security_group_rule.example2]
}

resource "ibm_is_security_group_rule" "example3" {
  group     = ibm_is_security_group.example.id
  direction = "outbound"
  remote    = "127.0.0.1"
  tcp {
    port_min = 8080
    port_max = 8080
  }
  depends_on = [ibm_is_security_group_rule.example2]
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bc1-2x8"

  primary_network_interface {
    subnet          = ibm_is_subnet.example.id
    security_groups = [ibm_is_security_group.example.id]
  }

  vpc        = ibm_is_vpc.example.id
  zone       = "us-south-1"
  keys       = [ibm_is_ssh_key.example.id]
  depends_on = [ibm_is_security_group_rule.example3]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

resource "ibm_resource_group" "example" {
  name = "example-resource-group" 
}

resource "ibm_is_dedicated_host_group" "example" {
  family         = "compute"
  class          = "cx2"
  zone           = "us-south-1"
  name           = "example-dh-group-01"
  resource_group = ibm_resource_group.example.id
}

resource "ibm_is_dedicated_host" "example" {
  profile        = "bx2d-host-152x608"
  name           = "example-dedicated-host-01"
  host_group     = ibm_is_dedicated_host_group.example.id
  resource_group = ibm_resource_group.example.id
}

// Example to provision an instance in a dedicated host
resource "ibm_is_instance" "example1" {
  name    = "example-instance-1"
  image   = ibm_is_image.example.id
  profile = "cx2-2x4"

  primary_network_interface {
    subnet          = ibm_is_subnet.example.id
    security_groups = [ibm_is_security_group.example.id]
  }
  dedicated_host = ibm_is_dedicated_host.example.id
  vpc            = ibm_is_vpc.example.id
  zone           = "us-south-1"
  keys           = [ibm_is_ssh_key.example.id]
  depends_on     = [ibm_is_security_group_rule.example3]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

// Example to provision an instance in a dedicated host that belongs to the provided dedicated host group
resource "ibm_is_instance" "example2" {
  name    = "example-instance-2"
  image   = ibm_is_image.example.id
  profile = "cx2-2x4"

  primary_network_interface {
    subnet          = ibm_is_subnet.example.id
    security_groups = [ibm_is_security_group.example.id]
  }
  dedicated_host_group = ibm_is_dedicated_host_group.example.id
  vpc                  = ibm_is_vpc.example.id
  zone                 = "us-south-1"
  keys                 = [ibm_is_ssh_key.example.id]
  depends_on           = [ibm_is_security_group_rule.example3]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

// Example to provision an instance from a snapshot, restoring boot volume from an existing snapshot

resource "ibm_is_snapshot" "example" {
  name          = "example-snapshot"
  source_volume = ibm_is_instance.example.volume_attachments[0].volume_id
}

resource "ibm_is_instance" "example" {
  name    = "example-vsi-restore"
  profile = "cx2-2x4"
  boot_volume {
    name     = "boot-restore"
    snapshot = ibm_is_snapshot.example.id
    tags     = ["tags-0"]
  }
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
  network_interfaces {
    subnet = ibm_is_subnet.example.id
    name   = "eth1"
  }
}


// Example to provision an instance using an enterprise managed catalog image

data ibm_is_images example {
  catalog_managed = true
}

resource "ibm_is_instance" "example" {
  name    = "example-vsi-catalog"
  profile = "cx2-2x4"
  catalog_offering {
    version_crn = data.ibm_is_images.example.images.0.catalog_offering.0.version.0.crn
  }
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
  network_interfaces {
    subnet = ibm_is_subnet.example.id
    name   = "eth1"
  }
}
```

## Timeouts

The `ibm_is_instance` resource provides the following [[Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:


- **create**: The creation of the instance is considered failed when no response is received for 30 minutes.
- **update**: The update of the instance or the attachment of a volume to an instance is considered failed when no response is received for 30 minutes.
- **delete**: The deletion of the instance is considered failed when no response is received for 30 minutes.


## Argument reference
Review the argument references that you can specify for your resource.
- `action` - (Optional, String) Action to be taken on the instance. Supported values are `stop`, `start`, or `reboot`.
  
  ~> **Note** 
    `action` allows to start, stop and reboot the instance and it is not recommended to manage the instance from terraform and other clients (UI/CLI) simultaneously, as it would cause unknown behaviour. `start` action can be performed only when the instance is in `stopped` state. `stop` and `reboot` actions can be performed only when the instance is in `running` state. It is also recommended to remove the `action` configuration from terraform once it is applied succesfully, to avoid instability in the terraform configuration later.
- `auto_delete_volume`- (Optional, Bool) If set to **true**, automatically deletes the volumes that are attached to an instance. **Note** Setting this argument can bring some inconsistency in the volume resource, as the volumes is destroyed along with instances.
- `availability_policy_host_failure` - (Optional, String) The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure. Supported values are `restart` and `stop`.
- `boot_volume`  (Optional, List) A list of boot volumes for an instance.

  Nested scheme for `boot_volume`:
  - `encryption` - (Optional, String) The type of encryption to use for the boot volume.
  - `name` - (Optional, String) The name of the boot volume.
  - `size` - (Optional, Integer) The size of the boot volume.(The capacity of the volume in gigabytes. This defaults to minimum capacity of the image and maximum to `250`.

    ~> **NOTE:**
    Supports only expansion on update (must be attached to a running instance and must not be less than the current volume size)
  - `snapshot` - (Optional, Forces new resource, String) The snapshot id of the volume to be used for creating boot volume attachment
    
    ~> **Note:**
    `snapshot` conflicts with `image` id, `instance_template` and `catalog_offering`
  - `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `catalog_offering` - (Optional, List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.
  Nested scheme for `catalog_offering`:
  - `offering_crn` - (Optional, String) The CRN for this catalog offering. Identifies a catalog offering by this unique property
  - `version_crn` - (Optional, String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property
 
    ~> **Note:**
    `offering_crn` conflicts with `version_crn`, both are mutually exclusive. `catalog_offering` and `image` id are mutually exclusive.
    `snapshot` conflicts with `image` id and `instance_template`
- `dedicated_host` - (Optional, String) The placement restrictions to use the virtual server instance. Unique ID of the dedicated host where the instance id placed.
- `dedicated_host_group` - (Optional, String) The placement restrictions to use for the virtual server instance. Unique ID of the dedicated host group where the instance is placed.

  -> **NOTE:**
  An instance can be moved from one dedicated host or group to another host or group. Moving an instance from public to dedicated host or vice versa is not allowed.

- `default_trusted_profile_auto_link` - (Optional, Forces new resource, Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default value : **true**
- `default_trusted_profile_target` - (Optional, Forces new resource, String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
- `force_action` - (Optional, Boolean) Required with `action`. If set to `true`, the action will be forced immediately, and all queued actions deleted. Ignored for the start action.
- `force_recovery_time` - (Optional, Integer) Define timeout (in minutes), to force the `is_instance` to recover from a perpetual "starting" state, during provisioning. And to force the is_instance to recover from a perpetual "stopping" state, during removal of user access.

  ~>**Note:** The force_recovery_time is used to retry multiple times until timeout.
- `image` - (Required, String) The ID of the virtual server image that you want to use. To list supported images, run `ibmcloud is images` or use `ibm_is_images` datasource.
  
  ~> **Note:**
  `image` conflicts with `boot_volume.0.snapshot` and `catalog_offering`, not required when creating instance using `instance_template` or `catalog_offering`
- `keys` - (Required, List) A comma-separated list of SSH keys that you want to add to your instance.
- `lifecycle_reasons`- (List) The reasons for the current lifecycle_state (if any).

  Nested scheme for `lifecycle_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle state.
    - `message` - (String) An explanation of the reason for this lifecycle state.
    - `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state`- (String) The lifecycle state of the virtual server instance. [ **deleting**, **failed**, **pending**, **stable**, **suspended**, **updating**, **waiting** ]
- `metadata_service_enabled` - (Optional, Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance. Default value : **false**
- `name` - (Optional, String) The instance name.
- `network_interfaces`  (Optional,  Forces new resource, List) A list of more network interfaces that are set up for the instance.

    -> **Allowed vNIC per profile.** 
    **&#x2022;** 2-16 vCPUs: Up to 5 vNICs </br> **&#x2022;** 17-48 vCPUs: Up to 10 vNICs </br> **&#x2022;** 49+ vCPUs: Up to 15 vNICs

  Nested scheme for `network_interfaces`:
  - `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on the interface. If **false**, IP spoofing is prevented on the interface. If **true**, IP spoofing is allowed on the interface.
      
      ~> **NOTE:**
      `allow_ip_spoofing` requires **IP spoofing operator** access under VPC infrastructure Services. As the **IP spoofing operator**, you can enable or disable the IP spoofing check on virtual server instances. Use this only if you have **IP spoofing operator** access.

  - `name` - (Optional, String) The name of the network interface.
  - `primary_ip` - (Optional, List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
      - `address` - (Optional, String) The IP address of the reserved IP. This is same as `network_interfaces.[].primary_ipv4_address`
      - `name`- (Optional, String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP.
  - `primary_ipv4_address` - (Optional, Deprecated, Forces new resource, String) The IPV4 address of the interface. `primary_ipv4_address` will be deprecated, use `primary_ip.[0].address` instead.
  - `subnet` - (Required, String) The ID of the subnet.
  - `security_groups`- (Optional, List of strings)A comma separated list of security groups to add to the primary network interface.
- `placement_group` - (Optional, string) Unique Identifier of the Placement Group for restricting the placement of the instance
- `primary_network_interface` - (Required, List) A nested block describes the primary network interface of this instance. Only one primary network interface can be specified for an instance. When using `instance_template`, `primary_network_interface` is not required.

  Nested scheme for `primary_network_interface`:
  - `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on the interface. If **false**, IP spoofing is prevented on the interface. If **true**, IP spoofing is allowed on the interface.

    ~> **NOTE:**
    `allow_ip_spoofing` requires **IP spoofing operator** access under VPC infrastructure Services. As the **IP spoofing operator**, you can enable or disable the IP spoofing check on virtual server instances. Use this only if you have **IP spoofing operator** access.

  - `name` - (Optional, String) The name of the network interface.
  - `port_speed` - (Deprecated, Integer) Speed of the network interface.
  - `primary_ip` - (Optional, List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
      - `address` - (Optional, String) The IP address of the reserved IP. This is same as `primary_network_interface.[0].primary_ipv4_address` which will be deprecated.
      - `name`- (Optional, String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP
  - `primary_ipv4_address` - (Optional, Deprecated, Forces new resource, String) The IPV4 address of the interface.`primary_ipv4_address` will be deprecated, use `primary_ip.[0].address` instead.
  - `subnet` - (Required, String) The ID of the subnet.
  - `security_groups`-List of strings-Optional-A comma separated list of security groups to add to the primary network interface.
- `profile` - (Required, String) The name of the profile that you want to use for your instance. Not required when using `instance_template`. To list supported profiles, run `ibmcloud is instance-profiles` or `ibm_is_instance_profiles` datasource.

  **NOTE:**
  When the `profile` is changed, the VSI is restarted. The new profile must:
    1. Have matching instance disk support. Any disks associated with the current profile will be deleted, and any disks associated with the requested profile will be created.        
    2. Be compatible with any placement_target(`dedicated_host`, `dedicated_host_group`, `placement_group`) constraints. For example, if the instance is placed on a dedicated host, the requested profile family must be the same as the dedicated host family.
  
  ~> **NOTE**
   Changing a `profile` without disk to a `profile` with disk or vise versa will result in recreating(forcenew) the resource.
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the instance.
- `instance_template` - (Optional, String) ID of the instance template to create the instance from. To create an instance template, use `ibm_is_instance_template` resource.
  
  ~> **Note:**
  `instance_template` conflicts with `boot_volume.0.snapshot`. When creating an instance using `instance_template`, [`image `, `primary_network_interface`, `vpc`, `zone`] are not required.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your instance. Tags can help you find your instance more easily later.
- `total_volume_bandwidth` - (Optional, Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `user_data` - (Optional, String) User data to transfer to the instance. For more information, about `user_data`, see [about user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data).
- `volumes`  (Optional, List) A comma separated list of volume IDs to attach to the instance.
- `vpc` - (Required, Forces new resource, String) The ID of the VPC where you want to create the instance. When using `instance_template`, `vpc` is not required.
- `zone` - (Required, Forces new resource, String) The name of the VPC zone where you want to create the instance. When using `instance_template`, `zone` is not required.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.
- `bandwidth` - The total bandwidth (in megabits per second) shared across the instance's network interfaces and storage volumes
- `boot_volume`- (List of Strings) A list of boot volumes that the instance uses.

  Nested scheme for `boot_volume`:
  - `encryption` - (String) The type of encryption that is used for the boot volume.
  - `iops`- (Integer) The number of input and output operations per second of the volume.
  - `name` - (String) The name of the boot volume.
  - `profile` - (String) The profile of the volume.
  - `size`- (Integer) The capacity of the volume in gigabytes.
- `crn` - (String) The CRN of the instance.
- `disks` - (List of Strings) The collection of the instance's disks. Nested `disks` blocks have the following structure:

  Nested scheme for `disks`:
  - `created_at` - (Timestamp) The date and time that the disk was created.
  - `href` - (String) The URL for the instance disk.
  - `id` - (String) The unique identifier for the instance disk.
  - `interface_type` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing the property, check for the unknown log values. Optionally stop processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
  - `name` - (String) The user defined name for the disk.
  - `resource_type` - (String) The resource type.
  - `size` - (String) The size of the disk in GB (gigabytes).
- `gpu`- (List of Strings) A list of GPUs that are assigned to the instance.

  Nested scheme for `gpu`:
  - `count`- (Integer) The count of the GPU.
  - `manufacture` - (String) The manufacturer of the GPU.
  - `memory`- (Integer) The amount of memory of the GPU in gigabytes.
  - `model` - (String) The model of the GPU.
- `placement_target` - The placement restrictions for the virtual server instance.
  - `crn` - The CRN of the placement target
  - `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    `more_info` - Link to documentation about deleted resources.
  - `href` - The URL for this placement target
  - `id` - The unique identifier for this placement target
  - `name` - The unique user-defined name for this placement target
  - `resource_type` - (String) The resource type.
- `id` - (String) The ID of the instance.
- `memory`- (Integer) The amount of memory that is allocated to the instance in gigabytes.
- `network_interfaces`- (List of Strings) A list of more network interfaces that are attached to the instance.

  Nested scheme for `network_interfaces`:
  - `allow_ip_spoofing` - (Bool) Indicates whether IP spoofing is allowed on the interface.
  - `id` - (String) The ID of the network interface.
  - `name` - (String) The name of the network interface.
  - `subnet` - (String) The ID of the subnet.
  - `security_groups`- (List of Strings) A list of security groups that are used in the network interface.
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
      - `address` - (String) The IP address of the reserved IP. This is same as `network_interfaces.[].primary_ipv4_address` which will be deprecated.
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
  - `primary_ipv4_address` - (String, Deprecated) The primary IPv4 address. Same as `primary_ip.[0].address`
- `primary_network_interface`- (List of Strings) A list of primary network interfaces that are attached to the instance.

  Nested scheme for `primary_network_interface`:
  - `allow_ip_spoofing` - (Bool) Indicates whether IP spoofing is allowed on the interface.
  - `id` - (String) The ID of the primary network interface.
  - `name` - (String) The name of the primary network interface.
  - `subnet` - (String) The ID of the subnet that the primary network interface is attached to.
  - `security_groups`-List of strings-A list of security groups that are used in the primary network interface.
  - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
      - `address` - (String) The IP address of the reserved IP. This is same as `primary_network_interface.[0].primary_ipv4_address` which will be deprecated.
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP.
  - `primary_ipv4_address` - (String, Deprecated) The primary IPv4 address. Same as `primary_ip.[0].address`
      ```terraform
      // primary_ipv4_address deprecation 
      output "primary_ipv4_address" {
        # value = ibm_is_instance.example.primary_network_interface.0.primary_ipv4_address // will be deprecated in future
        value = ibm_is_instance.example.primary_network_interface.0.primary_ip.0.address // use this instead 
      }
      ```
- `status` - (String) The status of the instance.
- `status_reasons` - (List) Array of reasons for the current status.

  Nested scheme for `status_reasons`:
  - `code` - (String) A string with an underscore as a special character identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) Link to documentation about this status reason
- `total_network_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance network interfaces.
- `volume_attachments`- (List of Strings) A list of volume attachments for the instance.

  Nested scheme for `volume_attachements`:
  - `id` - (String) The ID of the volume attachment.
  - `name` - (String) The name of the volume attachment.
  - `volume_id` - (String) The ID of the volume that is used in the volume attachment.
  - `volume_name` - (String) The name of the volume that is used in the volume attachment.
  - `volume_crn` - (String) The CRN of the volume that is used in the volume attachment.
- `vcpu`- (List of Strings) A list of virtual CPUs that are allocated to the instance.

  Nested scheme for `vcpu`:
  - `architecture` - (String) The architecture of the CPU.
  - `count`- (Integer) The number of virtual CPUS that are assigned to the instance.


## Import
The `ibm_is_instance` resource can be imported by using the instance ID.

**Example**

```
$ terraform import ibm_is_instance.example a1aaa111-1111-111a-1a11-a11a1a11a11a
```
