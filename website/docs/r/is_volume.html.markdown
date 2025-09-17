---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : volume"
description: |-
  Manages IBM volume.
---

# ibm_is_volume
Create, update, or delete a VPC block storage volume. For more information, about the VPC block storage volume, see [getting started with VPC](https://cloud.ibm.com/docs/vpc).

~> **NOTE:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example creates a volume with 10 IOPs tier.

```terraform
resource "ibm_is_volume" "example" {
  name    = "example-volume"
  profile = "10iops-tier"
  zone    = "us-south-1"
}
```
The following example creates a custom volume.

```terraform
resource "ibm_is_volume" "example" {
  name           = "example-volume"
  profile        = "custom"
  zone           = "us-south-1"
  iops           = 1000
  capacity       = 200
  encryption_key = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
}

```

The following example creates a volume from snapshot.
```terraform
resource "ibm_is_volume" "storage" {
  name            = "example-volume"
  profile         = "general-purpose"
  zone            = "us-south-1"
  source_snapshot = ibm_is_snapshot.example.id
}
```

The following example creates a volume from snapshot with allowed_use.
```terraform
resource "ibm_is_volume" "storage" {
  name            = "example-volume"
  profile         = "general-purpose"
  zone            = "us-south-1"
  source_snapshot = ibm_is_snapshot.example.id
  allowed_use {
    api_version       = "2025-07-02"
    bare_metal_server = "enable_secure_boot == true"
    instance          = "enable_secure_boot == true"
  }
}
```

## Timeouts
The `ibm_is_volume` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

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
- `adjustable_capacity_states` - (List) The attachment states that support adjustable capacity for this volume. Allowable list items are: `attached`, `unattached`, `unusable`. 
- `adjustable_iops_states` - (List) The attachment states that support adjustable IOPS for this volume. Allowable list items are: `attached`, `unattached`, `unusable`.
- `allowed_use` - (Optional, List) The usage constraints to be matched against the requested instance or bare metal server properties to determine compatibility. Can only be specified if `source_snapshot` is bootable. If not specified, the value of this property will be inherited from the `source_snapshot`.
    
    Nested schema for `allowed_use`:
    - `api_version` - (Optional, String) The API version with which to evaluate the expressions.
	  
    - `bare_metal_server` - (Optional, String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
   
    ~> **NOTE** </br> In addition, the following property is supported, corresponding to the `BareMetalServer` property: </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled.
	 
    - `instance` - (Optional, String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.
    
    ~> **NOTE** </br> In addition, the following variables are supported, corresponding to `Instance` properties:  </br>
      **&#x2022;** `gpu.count` - (integer) The number of GPUs. </br>
      **&#x2022;** `gpu.manufacturer` - (string) The GPU manufacturer. </br>
      **&#x2022;** `gpu.memory` - (integer) The overall amount of GPU memory in GiB (gibibytes). </br>
      **&#x2022;** `gpu.model` - (string) The GPU model. </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled. </br>
- `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.
- `capacity` - (Optional, Integer) (The capacity of the volume in gigabytes. This defaults to `100`, minimum to `10 ` and maximum to `16000`.

  ~> **NOTE:** Supports only expansion on update (must be attached to a running instance and must not be less than the current volume capacity). Can be updated only if volume is attached to an running virtual server instance. Stopped instance will be started on update of capacity of the volume.If `source_snapshot` is provided `capacity` must be at least the snapshot's minimum_capacity. The maximum value may increase in the future and If unspecified, the capacity will be the source snapshot's minimum_capacity.

- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume
- `delete_all_snapshots` - (Optional, Bool) Deletes all snapshots created from this volume.
- `encryption_key` - (Optional, Forces new resource, String) The key to use for encrypting this volume.
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
- `resource_group` - (Optional, Forces new resource, String) The resource group ID for this volume.
- `resource_controller_url` - (Optional, Forces new resource, String) The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.
- `source_snapshot` - The ID of snapshot from which to clone the volume.
- `source_snapshot_crn` - The CRN of snapshot from which to clone the volume.
- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `zone` - (Required, Forces new resource, String) The location of the volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.
- `catalog_offering` - (List) The catalog offering this volume was created from. If a virtual server instance is provisioned with a boot_volume_attachment specifying this volume, the virtual server instance will use this volume's catalog offering, including its pricing plan.If absent, this volume was not created from a catalog offering.

  Nested scheme for `catalog_offering`:
    - `version_crn` - (String) The CRN for this version of a catalog offering
    - `plan_crn` - (String) The CRN for this catalog offering version's billing plan
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
    
      Nested schema for `deleted`:
        - `more_info`  - (String) Link to documentation about deleted resources.
- `crn` - (String) The CRN for the volume.
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
- `storage_generation` - (Int) The storage generation indicates which generation the profile family belongs to. For the custom and tiered profiles, this value is 1. For the sdp profile, this value is 2.

## Import
The `ibm_is_volume` resource can be imported by using volume ID.

**Example**

```
$ terraform import ibm_is_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
