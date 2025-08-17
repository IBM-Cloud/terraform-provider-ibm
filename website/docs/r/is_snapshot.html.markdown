---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshot"
description: |-
  Manages IBM snapshot.
---

# ibm_is_snapshot

Create, update, or delete a snapshot. For more information, about subnet, see [creating snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-create).

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
  name                     = "example-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "example" {
  name    = "example-vsi"
  image   = ibm_is_image.example.id
  profile = "bx2-2x8"
  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.example.id]
  network_interfaces {
    subnet = ibm_is_subnet.example.id
    name   = "eth1"
  }
}
resource "ibm_is_snapshot" "example" {
  name          = "example-snapshot"
  source_volume = ibm_is_instance.example.volume_attachments[0].volume_id

  //User can configure timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```

## Example usage (clones)
```
resource "ibm_is_snapshot" "example_clones" {
  name            = "example-snapshot"
  source_volume   = ibm_is_instance.example.volume_attachments[0].volume_id
  clones          = ["us-south-1", "us-south-2"]
  //User can configure timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}  
 ``` 

## Example usage (Source snasphot - cross region snapshot crn)
```terraform
resource "ibm_is_snapshot" "example_copy" {
  name                = "example-snapshot"
  source_snapshot_crn = "crn:v1:bluemix:public:is:us-south:a/xxxxxxxxxxxxxxxxxxxxxxxx::snapshot:r006-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"
}  
 ``` 

 ## Example usage (allowed use)
```terraform
resource "ibm_is_snapshot" "example_allowed_use" {
  name          = "example-snapshot"
  source_volume = ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
  allowed_use {
    api_version       = "2025-03-31"
    bare_metal_server = "true"
    instance          = "true"
  }
}
``` 

## Timeouts
The `ibm_is_snapshot` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Snapshot.
- **delete** - (Default 10 minutes) Used for deleting Snapshot.


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the bare metal server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
 - `allowed_use` - (Optional, List) The usage constraints to match against the requested instance or bare metal server properties to determine compatibility. Can only be specified for bootable snapshots.
    
    Nested schema for `allowed_use`:
    - `api_version` - (Optional, String) The API version with which to evaluate the expressions.
	  
    - `bare_metal_server` - (Optional, String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this snapshot. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.
    
    ~> **NOTE** </br> the following variable is supported, corresponding to the `BareMetalServer` property: </br>
      **&#x2022;** `enable_secure_boot` - (boolean)Indicates whether secure boot is enabled.
	  
    - `instance` - (Optional, String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this snapshot. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.
    
    ~> **NOTE** </br> In addition, the following variables are supported, corresponding to `Instance` properties: </br>
      **&#x2022;** `gpu.count` - (integer) The number of GPUs. </br>
      **&#x2022;** `gpu.manufacturer` - (string) The GPU manufacturer. </br>
      **&#x2022;** `gpu.memory` - (integer) The overall amount of GPU memory in GiB (gibibytes). </br>
      **&#x2022;** `gpu.model` - (string) The GPU. </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled. </br>  
- `clones` - (Optional, List) The list of zones to create a clone of this snapshot.
- `encryption_key` - (String) A reference CRN to the root key used to wrap the data encryption key for the source snapshot.
- `name` - (Optional, String) The name of the snapshot.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the snapshot is to be created
- `source_volume` - (Optional, Forces new resource, String) The unique identifier for the volume for which snapshot is to be created.
- `source_snapshot_crn` - (Optional, Forces new resource, String) The CRN for source snapshot.

  -> **Note** `source_volume` and `source_snapshot_crn` are mutually exclusive, you can create snapshot either by a source volume or using another snapshot as a source.

- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your snapshot. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.
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
    - `crn` - (String) The CRN for this snapshot.
- `crn` - (String) The CRN for this snapshot.
- `encryption` - (String) The type of encryption used on the source volume. Supported values are **provider_managed**, **user_managed**.
- `encryption_key` - (String) The CRN of the `Key Protect Root Key` or `Hyper Protect Crypto Services Root Key` for this resource. The root key used to wrap the data encryption key for the source volume. This property will be present for volumes with an encryption type of `user_managed`.
- `href` - (String) The URL for this snapshot.
- `id` - (String) The unique identifier for this snapshot.
- `lifecycle_state` - (String) The lifecycle state of this snapshot. Supported values are **deleted**, **deleting**, **failed**, **pending**, **stable**, **updating**, **waiting**, **suspended**.
- `minimum_capacity` - (Integer) The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
- `operating_system` - (String) The globally unique name for an Operating System included in this image.
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


## Import

The `ibm_is_snapshot` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_snapshot.example < id >
```

**Example**

```
$ terraform import ibm_is_snapshot.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
