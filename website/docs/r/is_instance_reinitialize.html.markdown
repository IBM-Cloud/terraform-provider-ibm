---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_reinitialize"
description: |-
  Manages IBM Cloud VPC instance reinitialization.
---

# ibm_is_instance_reinitialize

Reinitialize a virtual server instance for VPC. This resource allows you to reinitialize an instance with a new image, boot volume, or snapshot. For more information, about managing VPC instance reinitialization, see [reinitializing virtual server instances](https://cloud.ibm.com/docs/vpc?topic=vpc-reinitializing-instances).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Reinitialize by image

In the following example, you can reinitialize an instance using a new image:

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
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "bx2-2x8"
  
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_reinitialize" "example" {
  depends_on = [ibm_is_instance.example]
  
  instance_id = ibm_is_instance.example.id
  image       = "r006-8f7c5c3d-5b8c-4b8c-8b8c-8b8c8b8c8b8c"
}
```

### Reinitialize by boot volume attachment

```terraform
resource "ibm_is_volume" "example" {
  name       = "example-volume"
  profile    = "general-purpose"
  zone       = "us-south-1"
  capacity   = 100
}

resource "ibm_is_instance_reinitialize" "example" {
  depends_on = [ibm_is_instance.example, ibm_is_volume.example]
  
  instance_id = ibm_is_instance.example.id
  
  boot_volume_attachment {
    name = "reinit-boot-volume"
    volume {
      id   = ibm_is_volume.example.id
      name = "reinit-volume"
    }
    delete_volume_on_instance_delete = false
  }
}
```

### Reinitialize with custom SSH keys and user data

```terraform
resource "ibm_is_instance_reinitialize" "example" {
  depends_on = [ibm_is_instance.example]
  
  instance_id = ibm_is_instance.example.id
  image       = "r006-8f7c5c3d-5b8c-4b8c-8b8c-8b8c8b8c8b8c"
  keys        = ["r006-12345678-1234-1234-1234-123456789012", "r006-87654321-4321-4321-4321-210987654321"]
  
  user_data = <<-EOF
    #!/bin/bash
    echo "Instance reinitialized at $(date)" > /var/log/reinit.log
    apt-get update
    apt-get install -y htop
  EOF
}
```

### Reinitialize with trusted profile

```terraform
resource "ibm_iam_trusted_profile" "example" {
  name = "example-trusted-profile"
}

resource "ibm_is_instance_reinitialize" "example" {
  depends_on = [ibm_is_instance.example, ibm_iam_trusted_profile.example]
  
  instance_id = ibm_is_instance.example.id
  image       = "r006-8f7c5c3d-5b8c-4b8c-8b8c-8b8c8b8c8b8c"
  
  default_trusted_profile {
    auto_link = "true"
    target {
      id = ibm_iam_trusted_profile.example.id
    }
  }
}
```

### Reinitialize with source snapshot

```terraform
resource "ibm_is_snapshot" "example" {
  name     = "example-snapshot"
  volume   = ibm_is_volume.example.id
}

resource "ibm_is_instance_reinitialize" "example" {
  depends_on = [ibm_is_instance.example, ibm_is_snapshot.example]
  
  instance_id = ibm_is_instance.example.id
  
  boot_volume_attachment {
    volume {
      source_snapshot {
        id = ibm_is_snapshot.example.id
      }
      name = "reinit-from-snapshot"
    }
  }
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `instance_id` - (Required, Forces new resource, String) The unique identifier of the virtual server instance to reinitialize.
- `image` - (Optional, Forces new resource, String) The unique identifier of the image to use when reinitializing the instance. Conflicts with `boot_volume_attachment`.
- `boot_volume_attachment` - (Optional, Forces new resource, List) The boot volume attachment configuration for reinitialization. Conflicts with `image`. The maximum length is `1`. 
  
  Nested `boot_volume_attachment` blocks have the following structure:
  
  - `name` - (Optional, Forces new resource, String) The name of the boot volume attachment.
  - `delete_volume_on_instance_delete` - (Optional, Forces new resource, Bool) Indicates whether the volume will be deleted when the instance is deleted.
  - `volume` - (Optional, Forces new resource, List) The boot volume configuration. The maximum length is `1`.
    
    Nested `volume` blocks have the following structure:
    
    - `id` - (Optional, Forces new resource, String) The unique identifier of the volume to attach as boot volume. Conflicts with `source_snapshot`.
    - `source_snapshot` - (Optional, Forces new resource, List) The snapshot to use as a source for the volume's data. The specified snapshot may be in a different account, subject to IAM policies. Conflicts with `id`.
      
      Nested `source_snapshot` blocks have the following structure:
      
      - `id` - (Required, Forces new resource, String) The unique identifier of the snapshot.
      
    - `allowed_use` - (Optional, Forces new resource, List) The allowed use configuration for this volume. The maximum length is `1`.
      
      Nested `allowed_use` blocks have the following structure:
      
      - `api_version` - (Required, Forces new resource, String) The API version with which to evaluate the expressions.
      - `bare_metal_server` - (Optional, Forces new resource, String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.
      - `instance` - (Optional, Forces new resource, String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.
      
    - `bandwidth` - (Optional, Forces new resource, Integer) The maximum bandwidth (in megabits per second) for the volume.
    - `capacity` - (Optional, Forces new resource, Integer) The capacity to use for the volume (in gigabytes). The specified value must be at least the image's `minimum_provisioned_size`, at most 250 gigabytes for `storage_generation: 1` or at most 32,000 gigabytes for `storage_generation: 2`, and within the `boot_capacity` range of the volume's profile.
    - `encryption_key` - (Optional, Forces new resource, List) The root key to use to wrap the data encryption key for the volume. The maximum length is `1`.
      
      Nested `encryption_key` blocks have the following structure:
      
      - `crn` - (Required, Forces new resource, String) The CRN of the Key Protect Root Key for this resource.
      
    - `iops` - (Optional, Forces new resource, Integer) The maximum I/O operations per second (IOPS) to use for this volume.
    - `name` - (Optional, Forces new resource, String) The name for this volume. The name must not be used by another volume in the region. If unspecified, the name will be a hyphenated list of randomly-selected words.
    - `profile` - (Optional, Forces new resource, List) The profile for this volume. The maximum length is `1`.
      
      Nested `profile` blocks have the following structure:
      
      - `name` - (Required, Forces new resource, String) The globally unique name for this volume profile.
      
    - `resource_group` - (Optional, Forces new resource, String) The resource group to use for this volume. If unspecified, the instance's resource group will be used.
    - `user_tags` - (Optional, Forces new resource, Set of Strings) The user tags associated with this volume.

- `default_trusted_profile` - (Optional, Forces new resource, List) The default trusted profile configuration to use for this virtual server instance. The maximum length is `1`.
  
  Nested `default_trusted_profile` blocks have the following structure:
  
  - `auto_link` - (Required, Forces new resource, String) If set to `true`, the system will create a link to the specified target trusted profile during instance reinitialization. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted.
  - `target` - (Required, Forces new resource, List) The default IAM trusted profile to use for this virtual server instance. The maximum length is `1`.
    
    Nested `target` blocks have the following structure:
    
    - `id` - (Optional, Forces new resource, String) The unique identifier for this trusted profile.
    - `crn` - (Optional, Forces new resource, String) The CRN for this trusted profile.

- `keys` - (Optional, Forces new resource, Set of Strings) The SSH key identifiers to use for the instance after reinitialization. If not specified, the instance will use the existing keys.
- `user_data` - (Optional, Forces new resource, String) User data to inject into the instance during reinitialization. The instance user data is replaced with this value.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the instance that was reinitialized.
- `status` - (String) The status of the instance after reinitialization. 

  **Supported status values**
  - `deleting`
  - `failed`
  - `pending`
  - `restarting`
  - `resuming`
  - `running`
  - `starting`
  - `stopped`
  - `stopping`
  - `suspended`
  - `suspending`
  - `updating`
  - `waiting`

## Important notes

- **One-time operation**: This resource performs a one-time reinitialization operation. After creation, the resource can be imported but any changes to arguments will force a new resource.
- **Instance state**: The instance must be in a `stopped` state before reinitialization. The resource will automatically stop the instance if it's running and restart it after reinitialization.
- **Boot source conflicts**: You must specify either `image` or `boot_volume_attachment`, but not both.
- **Volume ID vs Snapshot**: Within `boot_volume_attachment.volume`, you must specify either `id` or `source_snapshot`, but not both.
- **Force new**: All configuration arguments except `status` force the creation of a new resource when changed.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_instance_reinitialize` resource by using `instance_id`. The `id` property can be formed from `instance ID`. For example:

```terraform
import {
  to = ibm_is_instance_reinitialize.example
  id = "r006-12345678-1234-1234-1234-123456789012"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_instance_reinitialize.example r006-12345678-1234-1234-1234-123456789012
```

**Note**: The imported resource will only have the `instance_id` and `status` attributes populated. The boot source configuration (`image` or `boot_volume_attachment`) and other optional parameters (`keys`, `user_data`, `default_trusted_profile`) are not preserved during import.
