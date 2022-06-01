---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_template"
description: |-
  Manages IBM VPC instance template.
---

# ibm_is_instance_template
Create, update, or delete an instance template on VPC. For more information, about instance template, see [managing an instance template](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-template).

**Note:** 
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
      profile  = "general-purpose"
      capacity = 200
    }
  }
}

// Template with volume attachment that attaches existing storage volume. This template cannot be used with instance groups
resource "ibm_is_instance_template" "example1" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"
  metadata_service_enabled = false
  
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

## Argument reference
Review the argument references that you can specify for your resource. 
- `availability_policy_host_failure` - (Optional, String) The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure. Supported values are `restart` and `stop`.

- `boot_volume` - (Optional, List) A nested block describes the boot volume configuration for the template.

  Nested scheme for `boot_volume`:
	- `delete_volume_on_instance_delete` - (Optional, Bool) You can configure to delete the boot volume based on instance deletion.
	- `encryption` - (Optional, String) The encryption key CRN to encrypt the boot volume attached.
	- `name` - (Optional, String) The name of the boot volume.
- `total_volume_bandwidth` - (Optional, int) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `dedicated_host` - (Optional, Force new resource,String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host where the instance is placed.

  **Note:**
    - only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used

- `dedicated_host_group` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host group where the instance is placed.

  **Note:**
    - only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used

- `default_trusted_profile_auto_link` - (Optional, Forces new resource, Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default value : **true**
- `default_trusted_profile_target` - (Optional, Forces new resource, String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
- `image` - (Required, String) The ID of the image to create the template.
- `keys` - (Required, List) List of SSH key IDs used to allow log in user to the instances.
- `metadata_service_enabled` - (Optional, Forces new resource, Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.  Default value : **false**
- `name` - (Optional, String) The name of the instance template.
- `placement_group` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the placement group where the instance is placed.

  **Note:**
    - only one of [**dedicated_host**, **dedicated_host_group**, **placement_group**] can be used
- `profile` - (Required, String) The number of instances created in the instance group.
- `primary_network_interfaces` (Required, List) A nested block describes the primary network interface for the template.

  Nested scheme for `primary_network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the primary network interface.
  - `security_groups`- (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Force new resource, String) The VPC subnet to assign to the interface.
- `network_interfaces` - (Optional, List) A nested block describes the network interfaces for the template.

  Nested scheme for `network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the network interface.
  - `security_groups` - (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Forces new resource, String) The VPC subnet to assign to the interface.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID.
- `volume_attachments` - (Optional, List) A nested block describes the storage volume configuration for the template.

  Nested scheme for `volume_attachments`:
	- `name` - (Required, String) The name of the boot volume.
	- `volume` - (Required, String) The storage volume ID created in VPC.
  - `delete_volume_on_instance_delete`- (Required, Bool) You can configure to delete the storage volume to delete based on instance deletion.
  - `volume_prototype` - (Optional, Force new resource, List)

    Nested scheme for `volume_prototype`:
    - `capacity` - (Optional, Integer) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
    - `encryption_key` - (Optional, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for the resource.
    - `iops` - (Optional, Integer) The maximum input and output operations per second (IOPS) for the volume.
    - `profile` - (Optional, String) The global unique name for the volume profile to use for the volume.
    
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
