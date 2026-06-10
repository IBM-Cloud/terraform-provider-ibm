---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshots"
description: |-
  Reads IBM Cloud snapshots.
---
# ibm_is_snapshots

Import the details of an existing IBM Cloud infrastructure snapshot collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax, see [viewing snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-view).

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

data "ibm_is_snapshots" "example" {
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Optional, String) Filter snapshot collection by name of the snapshot.
- `resource_group` - (Optional, String) Filter snapshot collection by resource group id of the snapshot.
- `source_image` - (Optional, String) Filter snapshot collection by source image of the snapshot.
- `source_volume` - (Optional, String) Filter snapshot collection by source volume of the snapshot.
- `backup_policy_plan_tag` - (Optional, String)Filters the collection to resources with the exact tag value.
- `backup_policy_plan_id` - (Optional, String)Filters the collection to backup policy jobs with the backup plan with the specified identifier
- `snapshot_consistency_group_id` - (Optional, String)Filters the collection to snapshots with snapshot consistency group with the specified identifier.
- `snapshot_consistency_group_crn` - (Optional, String)Filters the collection to snapshots with snapshot consistency group with the specified identifier.
- `snapshot_copies_id` - Filters the collection to snapshots with copies with the specified identifier.
- `snapshot_copies_name` - Filters the collection to snapshots with copies with the exact specified name.
- `snapshot_copies_crn` - Filters the collection to snapshots with copies with the specified CRN.
- `snapshot_copies_remote_region_name` - Filters the collection to snapshots with copies with the exact remote region name.
- `source_snapshot_id` - Filters the collection to resources with the source snapshot with the specified identifier
- `source_snapshot_remote_region_name` - Filters the collection to snapshots with a source snapshot with the exact remote region name.
- `snapshot_source_volume_remote_region_name` - Filters the collection to snapshots with a source volume with the exact remote region name.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `snapshots` - (List) List of snapshots in the IBM Cloud Infrastructure.
  
  Nested scheme for `snapshots`:
  - `access_tags`  - (Array of Strings) Access management tags associated with the snapshot.
  - `allowed_use` - (List) The usage constraints to be matched against the requested instance properties to determine compatibility. While bare metal servers cannot be provisioned from snapshots, an image or volume created from this snapshot will inherit its allowed_use value. Only present on bootable snapshots. The value of this property will be inherited from the source volume or source snapshot at snapshot creation, but can be changed.
    
     Nested schema for `allowed_use`:
     - `api_version` - (String) The API version with which to evaluate the expressions.
	  
     - `bare_metal_server` - (String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this snapshot. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.
    
     ~> **NOTE** </br> the following variable is supported, corresponding to the `BareMetalServer` property: </br>
       **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled.
	  
     - `instance` - (String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this snapshot. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
    
     ~> **NOTE** </br> In addition, the following variables are supported, corresponding to `Instance` properties: </br>
       **&#x2022;** `gpu.count` - (integer) The number of GPUs. </br>
       **&#x2022;** `gpu.manufacturer` - (string) The GPU manufacturer. </br>
       **&#x2022;** `gpu.memory` - (integer) The overall amount of GPU memory in GiB (gibibytes). </br>
       **&#x2022;** `gpu.model` - (string) The GPU. </br>
       **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled. </br>  
  - `id` - (String) The unique identifier for this snapshot.
  - `backup_policy_plan` - (List) If present, the backup policy plan which created this snapshot.
  
   Nested scheme for `backup_policy_plan`:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
   
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this backup policy plan.
    - `id` - (String) The unique identifier for this backup policy plan.
    - `name` - (String) The unique user defined name for this backup policy plan. If unspecified, the name will be a hyphenated list of randomly selected words.
    - `resource_type` - (String) The type of resource referenced.
  - `bootable` - (Bool) Indicates if a boot volume attachment can be created with a volume created from this snapshot.
  - `catalog_offering` - (List) The catalog offering inherited from the snapshot's source. If a virtual server instance is provisioned with a source_snapshot specifying this snapshot, the virtual server instance will use this snapshot's catalog offering, including its pricing plan. If absent, this snapshot is not associated with a catalog offering.
  
    Nested scheme for `catalog_offering`:
    - `version_crn` - (String) The CRN for this version of a catalog offering
    - `plan_crn` - (String) The CRN for this catalog offering version's billing plan
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
    
      Nested schema for `deleted`:
        - `more_info`  - (String) Link to documentation about deleted resources.
  - `copies` - (List) The copies of this snapshot in other regions.
  
      Nested scheme for `copies`:
      - `crn` - (String) The CRN for the copied snapshot.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

          Nested scheme for `deleted`:
          - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for the copied snapshot.
      - `id` - (String) The unique identifier for the copied snapshot.
      - `name` - (String) The name for the copied snapshot. The name is unique across all snapshots in the copied snapshot's native region.
      - `remote` - (List) If present, this property indicates the referenced resource is remote to this region,and identifies the native region.
        Nested scheme for `remote`:
        - `href` - (String) The URL for this region.
        - `name` - (String) The globally unique name for this region.
      - `resource_type` - (String) The resource type.
  - `clones` - (List) The list of zones where clones of this snapshot exist.
  - `crn` - (String) The CRN for this snapshot.
  - `encryption` - (String) The type of encryption used on the source volume. Supported values are **provider_managed**, **user_managed** ]).
  - `encryption_key` - (String) The CRN of the `Key Protect Root Key` or `Hyper Protect Crypto Services Root Key` for this resource. The root key used to wrap the data encryption key for the source volume. This property will be present for volumes with an encryption type of `user_managed`.
  - `href` - (String) The URL for this snapshot.
  - `lifecycle_state` - (String) The lifecycle state of this snapshot. Supported values are **deleted**, **deleting**, **failed**, **pending**, **stable**, **updating**, **waiting**, **suspended**.
  - `minimum_capacity` - (Integer) The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
  - `operating_system` - (String) The globally unique name for the operating system included in this image.  
  - `resource_type` - (String) The resource type.
  - `service_tags` - (List) The [service tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags) prefixed with `is.snapshot:` associated with this snapshot.
  - `size` - (Integer) The size of this snapshot rounded up to the next gigabyte.
  - `snapshot_consistency_group` - (List) The snapshot consistency group which created this snapshot.

    Nested scheme for `snapshot_consistency_group`:
    - `crn` - (String) The CRN of this snapshot consistency group.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
    
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for the snapshot consistency group.
    - `id` - (String) The unique identifier for the snapshot consistency group.
    - `name` - (String) TThe name for the snapshot consistency group. The name is unique across all snapshot consistency groups in the region.
    - `resource_type` - (String) The resource type.
	- `software_attachments` - (List) The software attachments for this snapshot.
	  Nested schema for **software_attachments**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) A link to documentation about deleted resources.
		- `href` - (String) The URL for this snapshot software attachment.
		- `id` - (String) The unique identifier for this snapshot software attachment.
		- `name` - (String) The name for this snapshot software attachment. The name is unique across all software attachments for the snapshot.
		- `resource_type` - (String) The resource type.
  - `source_image` - (String) If present, the unique identifier for the image from which the data on this volume was most directly provisioned.
  - `source_snapshot` - (String) If present, the source snapshot this snapshot was created from.
    
     Nested scheme for `source_snapshot`:
      - `crn` - (String) The CRN of the source snapshot.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
    
          Nested scheme for `deleted`:
          - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for the source snapshot.
      - `id` - (String) The unique identifier for the source snapshot.
      - `name` - (String) The name for the source snapshot. The name is unique across all snapshots in the source snapshot's native region.
      - `remote` - (List) If present, this property indicates the referenced resource is remote to this region,and identifies the native region.
    
          Nested scheme for `remote`:
          - `href` - (String) The URL for this region.
          - `name` - (String) The globally unique name for this region.
      - `resource_type` - (String) The resource type.
  - `captured_at` - (String) The date and time that this snapshot was captured.
  - `tags` - (String) Tags associated with the snapshot.


