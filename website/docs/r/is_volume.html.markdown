---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : volume"
description: |-
  Manages IBM volume.
---

# ibm_is_volume
Create, update, or delete a VPC block storage volume. For more information, about the VPC block storage volume, see [getting started with VPC](https://cloud.ibm.com/docs/vpc).

## Example usage
The following example creates a volume with 10 IOPs tier.

```terraform
resource "ibm_is_volume" "testacc_volume" {
  name     = "test_volume"
  profile  = "10iops-tier"
  zone     = "us-south-1"
}

```
The following example creates a custom volume.

```terraform
resource "ibm_is_volume" "testacc_volume" {
  name     = "test_volume"
  profile  = "custom"
  zone     = "us-south-1"
  iops     = 1000
  capacity = 200
  encryption_key = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
}

```

## Timeouts
The `ibm_is_volume` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating instance.
- **delete** - (Default 10 minutes) Used for deleting instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags` - (Optional, List of Strings) A list of access management tags for the volume. **Note** Currently we are supporting only the existing tags attachement.
- `capacity` - (Optional, Forces new resource, Integer) (The capacity of the volume in gigabytes. This defaults to `100`.
- `delete_all_snapshots` - (Optional, Bool) Deletes all snapshots created from this volume.
- `encryption_key` - (Optional, Forces new resource, String) The key to use for encrypting this volume.
- `iops` - (Optional, Forces new resource, Integer) The total input/ output operations per second (IOPS) for your storage. This value is required for `custom` storage profiles only.
- `name` - (Required, String) The user-defined name for this volume.No.
- `profile` - (Required, Forces new resource, String) The profile to use for this volume.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID for this volume.
- `resource_controller_url` - (Optional, Forces new resource, String) The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.
- `tags`- (Optional, Array of Strings) A list of tags that you want to add to your volume. Tags can help you find your volume more easily later.No.
- `zone` - (Required, Forces new resource, String) The location of the volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the volume.
- `source_snapshot` - ID of the snapshot, if volume was created from it.
- `status` - (String) The status of volume. Supported values are **available**, **failed**, **pending**, **unusable**, or **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.

  Nested scheme for `status_reasons`:
  - `code` - (String) A string with an underscore as a special character identifying the status reason.
  - `message` - (String) An explanation of the status reason.
- `crn` - (String) The CRN for the volume.

## Import
The `ibm_is_volume` resource can be imported by using volume ID.

**Example**

```
$ terraform import ibm_is_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
