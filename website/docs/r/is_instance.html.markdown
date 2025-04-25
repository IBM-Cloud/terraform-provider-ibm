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

### Sample for creating an instance in a VPC using virtual network interface and network attachments.

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

resource "ibm_is_virtual_network_interface" "example"{
	name 						            = "example-vni"
	allow_ip_spoofing 			    = false
	enable_infrastructure_nat 	= true
	primary_ip {
		auto_delete 	  = false
    address 		    = "10.240.0.8"
	}
	subnet   = ibm_is_subnet.example.id
}

resource "ibm_is_instance" "example" {
  name                      = "example-instance"
  image                     = ibm_is_image.example.id
  profile                   = "bx2-2x8"
  metadata_service_enabled  = false

  boot_volume {
    encryption = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
  }

  primary_network_attachment {
    name = "vexample-primary-att"
    virtual_network_interface { 
      id = ibm_is_virtual_network_interface.example.id
    }
  }

  network_attachments {
    name = "example-network-att"
    virtual_network_interface {
      name = "example-net-vni"
			auto_delete = true
			enable_infrastructure_nat = true
			primary_ip {
				auto_delete 	= true
				address 		= "10.240.0.6"
			}
			subnet = ibm_is_subnet.example.id
    }
  }
  vpc  = ibm_is_vpc.example.id
  zone = ibm_is_subnet.example.zone
  keys = [ibm_is_ssh_key.example.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

### Sample for creating an instance in a VPC.

```terraform
resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = ibm_is_image.example.id
  profile = "bx2-2x8"
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

### Sample for creating an instance in a VPC with reservation

```terraform
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
  reservation_affinity {
    policy = "manual"
    pool {
      id = ibm_is_reservation.example.id
    }
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

```
### Sample for creating an instance with custom security group rules.

The following example shows how you can create a virtual server instance with custom security group rules. Note that the security group, security group rules, and the virtual server instance must be created in a specific order to meet the dependencies of the individual resources. To force the creation in a specific order, you use the [`depends_on` parameter](https://www.terraform.io/docs/configuration/resources.html). If you do not provide this parameter, all resources are created at the same time which might lead to resource dependency errors during the provisioning of your virtual server, such as `The security group to attach to is not available`.

~>**Conflict** 
 IBM Cloud terraform provider currently provides both a standalone `ibm_is_security_group_target` resource and a `security_groups` block defined in-line in the `ibm_is_instance` resource to attach security group to a network interface target. At this time you cannot use the `security_groups` block inline with `ibm_is_instance` in conjunction with the standalone resource `ibm_is_security_group_target`. </br> Doing so will create a **conflict of security groups** attaching to the network interface and will overwrite it.

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
  profile = "bx2-2x8"

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
    plan_crn = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
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

// Example for provisioning instance from an existing boot volume

resource "ibm_is_volume" "example" {
  name            = "example-volume"
  profile         = "10iops-tier"
  zone            = "us-south-1"
  source_snapshot = ibm_is_snapshot.example.id
}

resource "ibm_is_instance" "example" {
  name    = "example-vsi-restore"
  profile = "cx2-2x4"
  boot_volume {
    volume_id = ibm_is_volume.example.id
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
### Example to create an instance with metadata service configuration ###

```terraform

resource "ibm_is_instance" "example" {
  metadata_service {
    enabled = true
    protocol = "https"
    response_hop_limit = 5
  }
  name    = "example-vsi-catalog"
  profile = "cx2-2x4"
  catalog_offering {
    version_crn = data.ibm_is_images.example.images.0.catalog_offering.0.version.0.crn
    plan_crn = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
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
### Example to create an instance with cluster network attachments ###

```terraform

resource "ibm_is_instance" "is_instance" {
  name    = "example-instance"
  image   = data.ibm_is_image.example.id
  profile = "gx3d-160x1792x8h100"
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  cluster_network_attachments {
    name = "cna-1"
    cluster_network_interface{
      auto_delete = true
      name = "cni-1"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-2"
    cluster_network_interface{
      auto_delete = true
      name = "cni-2"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-3"
    cluster_network_interface{
      auto_delete = true
      name = "cni-3"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-4"
    cluster_network_interface{
      auto_delete = true
      name = "cni-4"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-5"
    cluster_network_interface{
      auto_delete = true
      name = "cni-5"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-5"
    cluster_network_interface{
      auto_delete = true
      name = "cna-6"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-7"
    cluster_network_interface{
      auto_delete = true
      name = "cni-7"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-8"
    cluster_network_interface{
      auto_delete = true
      name = "cni-8"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  vpc  = ibm_is_vpc.example.id
  zone = ibm_is_subnet.example.zone
  keys = [ibm_is_ssh_key.example.id]
}
```
## Timeouts

The `ibm_is_instance` resource provides the following [[Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:


- **create**: The creation of the instance is considered failed when no response is received for 30 minutes.
- **update**: The update of the instance or the attachment of a volume to an instance is considered failed when no response is received for 30 minutes.
- **delete**: The deletion of the instance is considered failed when no response is received for 30 minutes.


## Argument reference
Review the argument references that you can specify for your resource.

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the instance.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `action` - (Optional, String) Action to be taken on the instance. Supported values are `stop`, `start`, or `reboot`.
  
  ~> **Note** 
    `action` allows to start, stop and reboot the instance and it is not recommended to manage the instance from terraform and other clients (UI/CLI) simultaneously, as it would cause unknown behaviour. `start` action can be performed only when the instance is in `stopped` state. `stop` and `reboot` actions can be performed only when the instance is in `running` state. It is also recommended to remove the `action` configuration from terraform once it is applied succesfully, to avoid instability in the terraform configuration later.
- `auto_delete_volume`- (Optional, Bool) If set to **true**, automatically deletes the volumes that are attached to an instance. **Note** Setting this argument can bring some inconsistency in the volume resource, as the volumes is destroyed along with instances.
- `availability_policy_host_failure` - (Optional, String) The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure. Supported values are `restart` and `stop`.
- `boot_volume`  (Optional, List) A list of boot volumes for an instance.

  Nested scheme for `boot_volume`:
  - `auto_delete_volume` - (Optional, String) If set to **true**, when deleting the instance the volume will also be deleted
  - `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.
  - `encryption` - (Optional, String) The type of encryption to use for the boot volume.
  - `name` - (Optional, String) The name of the boot volume.
  - `size` - (Optional, Integer) The size of the boot volume.(The capacity of the volume in gigabytes. This defaults to minimum capacity of the image and maximum to `250`.)

    ~> **NOTE:**
    Supports only expansion on update (must be attached to a running instance and must not be less than the current volume size)
  - `snapshot` - (Optional, Forces new resource, String) The snapshot id of the snapshot to be used for creating boot volume attachment
    
    ~> **Note:**
    `snapshot` conflicts with `image` id, `instance_template` , `catalog_offering`, `boot_volume.volume_id` and `snapshot_crn`
  - `snapshot_crn` - (Optional, Forces new resource, String) The crn of the snapshot to be used for creating boot volume attachment
    
    ~> **Note:**
    `snapshot` conflicts with `image` id, `instance_template` , `catalog_offering`, `boot_volume.volume_id` and `snapshot`
  - `volume_id` - (Optional, Forces new resource, String) The ID of the volume to be used for creating boot volume attachment
    ~> **Note:** 

     - `volume_id` conflicts with `image` id, `instance_template` ,`boot_volume.snapshot`, `catalog_offering`, 
  - `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `catalog_offering` - (Optional, List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.
  Nested scheme for `catalog_offering`:
  - `offering_crn` - (Optional, String) The CRN for this catalog offering. Identifies a catalog offering by this unique property
  - `version_crn` - (Optional, String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property
  - `plan_crn` - (Optional, String) The CRN for this catalog offering version's billing plan. If unspecified, no billing plan will be used (free). Must be specified for catalog offering versions that require a billing plan to be used.
 
    ~> **Note:**
    `offering_crn` conflicts with `version_crn`, both are mutually exclusive. `catalog_offering` and `image` id are mutually exclusive.
    `snapshot` conflicts with `image` id and `instance_template`

- `cluster_network_attachments` - (Optional, List) The cluster network attachments for this virtual server instance.The cluster network attachments are ordered for consistent instance configuration. [About cluster networks](https://cloud.ibm.com/docs/vpc?topic=vpc-about-cluster-network)

  Nested schema for **cluster_network_attachments**:
	- `name` - (Required, String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance. (`name` is a apply once attribute, changing it will not be detected by terraform)
  - `cluster_network_interface` - (Required, List) The cluster network interface for this instance cluster network attachment.
    
      Nested schema for **cluster_network_interface**:
      - `id` - (Required, String) The unique identifier for this cluster network interface.
      - `name` - (Required, String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
      - `primary_ip` - (Required, List) The primary IP for this cluster network interface.
        
          Nested schema for **primary_ip**:
          - `address` - (Required, String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
          - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
            
          - `href` - (Required, String) The URL for this cluster network subnet reserved IP.
          - `id` - (Required, String) The unique identifier for this cluster network subnet reserved IP.
          - `name` - (Required, String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
          - `resource_type` - (Computed, String) The resource type.
      - `subnet` - (Required, List)
        
          Nested schema for **subnet**:
          - `id` - (Required, String) The unique identifier for this cluster network subnet.

  ~> **Note:** 
  **&#x2022;** `cluster_network_attachments` updation requires the instance to be in stopped state. Use `action` attribute or `ibm_is_instance_action` resource accordingly to stop/start the instance.</br>
  **&#x2022;** Using cluster_network_attachments in `ibm_is_instance` and `ibm_is_instance_cluster_network_attachment` resource together would result in changes shown in both resources alternatively, use either of them or use meta lifecycle argument `ignore_changes` on `cluster_network_attachments`</br>

- `confidential_compute_mode` - (Optional, String) The confidential compute mode to use for this virtual server instance.If unspecified, the default confidential compute mode from the profile will be used. **Constraints: Allowable values are: `disabled`, `sgx`, `tdx`** {Select Availability}

  ~>**Note:** The confidential_compute_mode is `Select Availability` feature. Confidential computing with Intel SGX for VPC is available only in the US-South (Dallas) region.

- `dedicated_host` - (Optional, String) The placement restrictions to use the virtual server instance. Unique ID of the dedicated host where the instance id placed.
- `dedicated_host_group` - (Optional, String) The placement restrictions to use for the virtual server instance. Unique ID of the dedicated host group where the instance is placed.

  -> **NOTE:**
  An instance can be moved from one dedicated host or group to another host or group. Moving an instance from public to dedicated host or vice versa is not allowed.

- `default_trusted_profile_auto_link` - (Optional, Forces new resource, Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default value : **true**
- `default_trusted_profile_target` - (Optional, Forces new resource, String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
- `enable_secure_boot` - (Optional, Boolean) Indicates whether secure boot is enabled for this virtual server instance.If unspecified, the default secure boot mode from the profile will be used. {Select Availability}

  ~>**Note:** The enable_secure_boot is `Select Availability` feature.
- `force_action` - (Optional, Boolean) Required with `action`. If set to `true`, the action will be forced immediately, and all queued actions deleted. Ignored for the start action.
- `force_recovery_time` - (Optional, Integer) Define timeout (in minutes), to force the `is_instance` to recover from a perpetual "starting" state, during provisioning. And to force the is_instance to recover from a perpetual "stopping" state, during removal of user access.

  ~>**Note:** The force_recovery_time is used to retry multiple times until timeout.
- `image` - (Required, String) The ID of the virtual server image that you want to use. To list supported images, run `ibmcloud is images` or use `ibm_is_images` datasource.
  
  ~> **Note:**
  `image` conflicts with `boot_volume.0.snapshot` and `catalog_offering`, not required when creating instance using `instance_template` or `catalog_offering`
- `keys` - (Optional, List) A comma-separated list of SSH keys that you want to add to your instance. The public SSH keys for the administrative user of the virtual server instance. Keys will be made available to the virtual server instance as cloud-init vendor data. For cloud-init enabled images, these keys will also be added as SSH authorized keys for the administrative user.

  ~> **Note:**
  For Windows images, the keys of type rsa must be specified, and one will be selected to encrypt the administrator password. Keys are optional for other images, but if no keys are specified, the instance will be inaccessible unless the specified image provides another means of access.
  
  ~> **Note:**
  **&#x2022;** `ed25519` can only be used if the operating system supports this key type.</br>
  **&#x2022;** `ed25519` can't be used with Windows or VMware images.</br>
- `lifecycle_reasons`- (List) The reasons for the current lifecycle_state (if any).

  Nested scheme for `lifecycle_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle state.
    - `message` - (String) An explanation of the reason for this lifecycle state.
    - `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state`- (String) The lifecycle state of the virtual server instance. [ **deleting**, **failed**, **pending**, **stable**, **suspended**, **updating**, **waiting** ]
- `metadata_service_enabled` - (Optional, Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance. Default value : **false**

  ~> **NOTE**
  `metadata_service_enabled` is deprecated and will be removed in the future. Use `metadata_service` instead
- `metadata_service` - (Optional, List) The metadata service configuration. 

  Nested scheme for `metadata_service`:
  - `enabled` - (Optional, Bool) Indicates whether the metadata service endpoint will be available to the virtual server instance. Default is **false**
  - `protocol` - (Optional, String) The communication protocol to use for the metadata service endpoint. Applies only when the metadata service is enabled. Default is **http**
  - `response_hop_limit` - (Optional, Integer) The hop limit (IP time to live) for IP response packets from the metadata service. Default is **1**
- `name` - (Optional, String) The instance name.
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
      - `name` - (Optional, String) The virtual network interface name. The name must not be used by another virtual network interface in the VPC. 
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
- `network_interfaces`  (Optional,  Forces new resource, List) A list of more network interfaces that are set up for the instance.

    -> **Allowed vNIC per profile.** Follow the vNIC count as per the instance profile's `network_interface_count`. For details see  [`is_instance_profile`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/is_instance_profile) or [`is_instance_profiles`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/is_instance_profiles).

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
      - `name` - (Optional, String) The virtual network interface name. The name must not be used by another virtual network interface in the VPC.
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

- `reservation_affinity` - (Optional, List) The reservation affinity for the instance
  Nested scheme for `reservation_affinity`:
  - `policy` - (Optional, String) The reservation affinity policy to use for this virtual server instance.

    ->**policy** 
			&#x2022; disabled: Reservations will not be used
      </br>&#x2022; manual: Reservations in pool will be available for use
  - `pool` - (Optional, String) The pool of reservations available for use by this virtual server instance. Specified reservations must have a status of active, and have the same profile and zone as this virtual server instance. The pool must be empty if policy is disabled, and must not be empty if policy is manual.
    Nested scheme for `pool`:
    - `id` - The unique identifier for this reservation
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the instance.
- `instance_template` - (Optional, String) ID of the instance template to create the instance from. To create an instance template, use `ibm_is_instance_template` resource.
  
  ~> **Note:**
  `instance_template` conflicts with `boot_volume.0.snapshot`. When creating an instance using `instance_template`, [`image `, `primary_network_interface`, `vpc`, `zone`] are not required.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your instance. Tags can help you find your instance more easily later.
- `total_volume_bandwidth` - (Optional, Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `user_data` - (Optional, String) User data to transfer to the instance. For more information, about `user_data`, see [about user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data).
- `volumes`  (Optional, List) A comma separated list of volume IDs to attach to the instance. Mutually exclusive with `volume_prototypes`.
- `volume_prototypes`- (List of Strings) A list of data volumes to attach to the instance. Mutually exclusive with `volumes`.

  Nested scheme for `volume_prototypes`:
  - `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.
  - `delete_volume_on_instance_delete` - (Bool) If set to **true**, automatically deletes the volumes that are attached to an instance. **Note** Setting this argument can bring some inconsistency in the volume resource, as the volumes is destroyed along with instances.
  - `encryption` - (String) The type of encryption that is used for the volume prototype.
  - `iops`- (Integer) The number of input and output operations per second of the volume prototype.
  - `name` - (String) The name of the volume prototype.
  - `profile` - (String) The profile of the volume prototype.
  - `size`- (Integer) The capacity of the volume in gigabytes.
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
- `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.

  Nested scheme for `catalog_offering`:
    - `offering_crn` - (String) The CRN for this catalog offering. Identifies a catalog offering by this unique property
    - `version_crn` - (String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property
    - `plan_crn` - (String) The CRN for this catalog offering version's billing plan
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
		  Nested schema for `deleted`:
      - `more_info`  - (String) Link to documentation about deleted resources.

- `cluster_network_attachments` - (List) The cluster network attachments for this virtual server instance.The cluster network attachments are ordered for consistent instance configuration.
    Nested schema for **cluster_network_attachments**:
    - `href` - (String) The URL for this instance cluster network attachment.
    - `id` - (String) The unique identifier for this instance cluster network attachment.
    - `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
    - `resource_type` - (String) The resource type.


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
- `health_reasons` - (List) The reasons for the current health_state (if any).

    Nested scheme for `health_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this health state.
    - `message` - (String) An explanation of the reason for this health state.
    - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.
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
- `numa_count` - (Integer) The number of NUMA nodes this instance is provisioned on. This property may be absent if the instance's status is not running.
- `network_attachments` - (List) The network attachments list for this virtual server instance.
    Nested schema for **network_attachments**:

    - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

        Nested scheme for `primary_ip`:
        - `auto_delete` - (Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
        - `address` - (String) The IP address of the reserved IP. 
        - `name`- (String) The user-defined or system-provided name for this reserved IP
        - `id`- (String) The unique identifier for this reserved IP.
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
- `primary_network_attachment` - (List) The primary network attachment for this virtual server instance.
    Nested schema for **primary_network_attachment**:

    - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

        Nested scheme for `primary_ip`:
        - `auto_delete` - (Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
        - `address` - (String) The IP address of the reserved IP. 
        - `name`- (String) The user-defined or system-provided name for this reserved IP
        - `id`- (String) The unique identifier for this reserved IP.
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
      - `resource_type` - (string) The resource type.      ```
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
  - `manufacturer`- (String) The VCPU manufacturer.


## Import
The `ibm_is_instance` resource can be imported by using the instance ID.

**Example**

```
$ terraform import ibm_is_instance.example a1aaa111-1111-111a-1a11-a11a1a11a11a
```
