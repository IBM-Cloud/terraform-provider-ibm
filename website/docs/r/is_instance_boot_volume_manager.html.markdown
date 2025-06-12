---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_instance_boot_volume_manager"
description: |-
  Manages IBM VPC instance boot volume.
---

# ibm_is_instance_boot_volume_manager

Provides a resource to manage the boot volume of a VPC virtual server instance. This resource allows you to manage boot volumes that may have been orphaned or need to be managed independently of their instance lifecycle.

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_instance_boot_volume_manager` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

This resource is particularly useful for managing boot volumes that:
- Have been orphaned after instance deletion
- Need independent lifecycle management from their instance
- Require specific configuration changes that affect the boot volume properties

For more information, about VPC virtual server instances, see [getting started with Virtual Private Cloud](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started). For more information about VPC storage, see [About Block Storage for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-about).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Basic boot volume management
```terraform
resource "ibm_is_instance_boot_volume_manager" "example" {
  boot_volume = "r006-10cc2ee8-f395-47a1-b043-4e7a855a6dd0"
  name        = "my-managed-boot-volume"
}
```

### Boot volume with performance configuration
```terraform
resource "ibm_is_instance_boot_volume_manager" "example" {
  boot_volume = "r006-10cc2ee8-f395-47a1-b043-4e7a855a6dd0"
  name        = "high-performance-boot"
  profile     = "10iops-tier"
  capacity    = 120
  iops        = 3000
}
```

### Boot volume with tags and conditional deletion
```terraform
resource "ibm_is_instance_boot_volume_manager" "example" {
  boot_volume             = "r006-10cc2ee8-f395-47a1-b043-4e7a855a6dd0"
  name                    = "managed-boot-volume"
  profile                 = "10iops-tier"
  tags                    = ["env:production", "managed:terraform"]
  access_tags             = ["project:web-app"]
  delete_volume           = true
  delete_all_snapshots    = true
}
```

### Managing boot volume from instance
```terraform
resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_boot_volume_manager" "example" {
  boot_volume      = ibm_is_instance.example.boot_volume[0].volume_id
  name             = "managed-boot-volume"
  profile          = "10iops-tier"
  delete_volume    = false  # Don't delete when resource is destroyed
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `access_tags` - (Optional, List of Strings) A list of access management tags to attach to the boot volume.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `boot_volume` - (Required, Forces new resource, String) The unique identifier for the boot volume.
- `capacity` - (Optional, Integer) The capacity of the volume in gigabytes. The minimum capacity is 10 GB and the maximum capacity is 16,000 GB.

  ~> **NOTE:** Supports only expansion on update (must be attached to a running instance and must not be less than the current volume capacity). Can be updated only if volume is attached to a running virtual server instance. Stopped instance will be started automatically on update of capacity.
- `delete_all_snapshots` - (Optional, Boolean) If set to true, all snapshots created from this volume will be deleted when the volume is deleted. Default value is `false`.
- `delete_volume` - (Optional, Boolean) If set to true, the boot volume will be deleted when this resource is destroyed. Default value is `false`.
- `iops` - (Optional, Integer) The maximum I/O operations per second (IOPS) for the volume. This value is required for `custom` storage profiles only.

  ~> **NOTE:** `iops` value can be upgraded and downgraded if volume is attached to a running virtual server instance. Stopped instances will be started automatically on update of volume.

  This table shows how storage size affects the `iops` ranges:

  |   Size range (GB)  |   IOPS range   |
  |--------------------|----------------|
  |    10  -     39    |  100  -   1000 |
  |    40  -     79    |  100  -   2000 |
  |    80  -     99    |  100  -   4000 |
  |   100  -    499    |  100  -   6000 |
  |   500  -    999    |  100  -  10000 |
  |  1000  -   1999    |  100  -  20000 |
  |  2000  -   3999    |  100  -  40000 |
  |  4000  -   7999    |  100  -  40000 |
  |  8000  -   9999    |  100  -  48000 |
  | 10000  -  16000    |  100  -  48000 |

- `name` - (Optional, String) The user-defined name for this boot volume. If unspecified, the name will be automatically assigned.
- `profile` - (Optional, String) The globally unique name of the volume profile to use for this volume. Valid profiles are `general-purpose`, `5iops-tier`, `10iops-tier`, and `custom`.

  ~> **NOTE:** Tiered profiles [`general-purpose`, `5iops-tier`, `10iops-tier`] can be upgraded and downgraded into each other if volume is attached to a running virtual server instance. Stopped instances will be started automatically on update of volume.
- `tags` - (Optional, List of Strings) User tags for the boot volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
- `crn` - (String) The CRN for the volume.
- `encryption_key` - (String) The CRN of the encryption key used to encrypt this volume.
- `encryption_type` - (String) The type of encryption used on the volume.
- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health state of this volume.
- `id` - (String) The unique identifier of the volume.
- `resource_group` - (List) The resource group for this volume.

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The user-defined name for this resource group.
- `source_snapshot` - (String) The unique identifier for the snapshot from which this volume was created.
- `status` - (String) The status of volume. Supported values are **available**, **failed**, **pending**, **unusable**, or **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.

  Nested scheme for `status_reasons`:
  - `code` - (String) A string with an underscore as a special character identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) Link to documentation about this status reason.
- `volume_attachments` - (List) The volume attachments for this volume.

  Nested scheme for `volume_attachments`:
  - `delete_volume_on_instance_delete` - (Boolean) If set to true, this volume will be deleted when the instance is deleted.
  - `device` - (List) Information about how the volume is exposed to the instance operating system.
    
    Nested scheme for `device`:
    - `id` - (String) A unique identifier for the device which is exposed to the instance operating system.
  - `href` - (String) The URL for this volume attachment.
  - `id` - (String) The unique identifier for this volume attachment.
  - `instance` - (List) The attached instance.
    
    Nested scheme for `instance`:
    - `crn` - (String) The CRN for this virtual server instance.
    - `href` - (String) The URL for this virtual server instance.
    - `id` - (String) The unique identifier for this virtual server instance.
    - `name` - (String) The user-defined name for this virtual server instance.
  - `name` - (String) The user-defined name for this volume attachment.
  - `type` - (String) The type of volume attachment.
- `zone` - (String) The zone where this volume resides.

## Timeouts
The `ibm_is_instance_boot_volume_manager` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for adopting the boot volume.
- **update** - (Default 10 minutes) Used for updating boot volume properties.
- **delete** - (Default 10 minutes) Used for deleting boot volume (when `delete_volume` is true).

## Import
The `ibm_is_instance_boot_volume_manager` resource can be imported by using the boot volume ID.

**Syntax**

```
$ terraform import ibm_is_instance_boot_volume_manager.example <volume_id>
```

**Example**

```
$ terraform import ibm_is_instance_boot_volume_manager.example d7bec597-4726-451f-8a63-e62e6f19c32c
```

## Usage scenarios

### Orphaned boot volume management
When an instance is deleted but the boot volume remains (either by design or accident), you can use this resource to manage the orphaned volume:

```terraform
resource "ibm_is_instance_boot_volume_manager" "orphaned_volume" {
  boot_volume      = "r006-orphaned-volume-id"
  name             = "recovered-boot-volume"
  profile          = "10iops-tier"
  delete_volume    = true  # Clean up when done
}
```

### Instance-independent management
Manage boot volume properties independently of the instance lifecycle:

```terraform
# Instance with minimal boot volume configuration
resource "ibm_is_instance" "app_server" {
  name    = "app-server"
  # ... other configuration
}

# Separate management of the boot volume
resource "ibm_is_instance_boot_volume_manager" "boot_config" {
  boot_volume             = ibm_is_instance.app_server.boot_volume[0].volume_id
  name                    = "app-server-boot-managed"
  profile                 = "10iops-tier"
  tags                    = ["env:production", "managed:terraform"]
  delete_volume           = false  # Keep volume when resource is destroyed
  delete_all_snapshots    = false  # Preserve snapshots
}
```

## Important notes

- **Instance state management**: When updating volume properties like capacity, profile, or IOPS, the attached instance must be running. This resource will automatically start stopped instances when necessary for updates.
- **Boot volume lifecycle**: Boot volumes cannot be detached from running instances. This resource manages the volume properties while it remains attached.
- **Deletion behavior**: By default, this resource does not delete the actual boot volume when destroyed (only removes it from Terraform state). Set `delete_volume = true` to actually delete the volume.
- **Snapshot management**: Set `delete_all_snapshots = true` to automatically clean up all snapshots when the volume is deleted.
- **Performance updates**: Volume performance characteristics (profile, IOPS, capacity) can be modified while the volume is in use, but may require brief instance restart.
- **Profile compatibility**: Some profile changes require specific IOPS values. Refer to the IOPS table for valid combinations.

## Error handling

Common scenarios and resolutions:

- **Volume not found**: Ensure the boot volume ID is correct and the volume exists in the target region.
- **Volume attached to running instance**: Some operations require the instance to be running. The resource will attempt to start stopped instances automatically.
- **Invalid IOPS for profile**: Ensure IOPS values are within the valid range for the selected profile and capacity.
- **Capacity decrease**: Boot volume capacity can only be increased, never decreased.
