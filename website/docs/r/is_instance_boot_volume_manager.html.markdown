---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_boot_volume_manager"
description: |-
  Manages IBM instance boot volume.
---

# ibm_is_instance_boot_volume_manager
Provides a resource to manage a boot volume of a VPC virtual server instance. Manage the boot volume of a VPC virtual server instance created alongwith the instance creation with this resource.

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_instance_boot_volume_manager` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

Every Virtual server instance has a boot volume that needs to be managed but not destroyed. When Terraform first adopts a instance_boot_volume_manager, 

For more information, about VPC, see [getting started with Virtual Private Cloud](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started). For more information, about updating default security group, see [updating a VPC's default security group rules](https://cloud.ibm.com/docs/vpc?topic=vpc-updating-the-default-security-group&interface=ui).

~> **NOTE:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example manages a boot volume with 10 IOPs tier.

```terraform
resource "ibm_is_instance_boot_volume_manager" "example" {
  volume_id   = "r006-10cc2ee8-f395-47a1-b043-4e7a855a6dd0"
  name        = "example-volume"
  profile     = "10iops-tier"
}
```

## Timeouts
The `ibm_is_instance_boot_volume_manager` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating instance.
- **delete** - (Default 10 minutes) Used for deleting instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the bare metal server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `capacity` - (Optional, Integer) (The capacity of the volume in gigabytes. This defaults to `100`, minimum to `10 ` and maximum to `16000`.

  ~> **NOTE:** Supports only expansion on update (must be attached to a running instance and must not be less than the current volume capacity). Can be updated only if volume is attached to an running virtual server instance. Stopped instance will be started on update of capacity of the volume.If `source_snapshot` is provided `capacity` must be at least the snapshot's minimum_capacity. The maximum value may increase in the future and If unspecified, the capacity will be the source snapshot's minimum_capacity.

- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume
- `delete_all_snapshots` - (Optional, Bool) Deletes all snapshots created from this volume.
- `iops` - (Optional, Integer) The total input/ output operations per second (IOPS) for your storage. This value is required for `custom` storage profiles only.

  ~> **NOTE:** `iops` value can be upgraded and downgraged if volume is attached to an running virtual server instance. Stopped instances will be started on update of volume.

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
            |  4000  -   1999    |  100  -  40000 |
            |  8000  -   1999    |  100  -  48000 |
            | 10000  -  16000    |  100  -  48000 |

- `name` - (Required, String) The user-defined name for this volume.No.
- `profile` - (Required, String) The profile to use for this volume.

  ~> **NOTE:**  tiered profiles [`general-purpose`, `5iops-tier`, `10iops-tier`] can be upgraded and downgraded into each other if volume is attached to an running virtual server instance. Stopped instances will be started on update of volume.

- `source_snapshot` - The ID of snapshot from which to clone the volume.
- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `volume_id` - (Required, Forces new resource, String) The volume id of the volume from boot volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.
- `encryption_type` - (String) The type of encryption used in the volume [**provider_managed**, **user_managed**].
- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.
- `id` - (String) The unique identifier of the volume.
- `status` - (String) The status of volume. Supported values are **available**, **failed**, **pending**, **unusable**, or **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.

  Nested scheme for `status_reasons`:
  - `code` - (String) A string with an underscore as a special character identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) Link to documentation about this status reason
- `crn` - (String) The CRN for the volume.
- `encryption_key` - (String) The key to use for encrypting this volume.
- `resource_group` - (String) The resource group ID for this volume.
- `zone` - (String) The location of the volume.

## Import
The `ibm_is_instance_boot_volume_manager` resource can be imported by using volume ID.

**Example**

```
$ terraform import ibm_is_instance_boot_volume_manager.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
