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

## Timeouts
The `ibm_is_snapshot` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Snapshot.
- **delete** - (Default 10 minutes) Used for deleting Snapshot.


## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Optional, String) The name of the snapshot.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the snapshot is to be created
- `source_volume` - (Required, Forces new resource, String) The unique identifier for the volume for which snapshot is to be created. 
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
- `crn` - (String) The CRN for this snapshot.
- `encryption` - (String) The type of encryption used on the source volume. Supported values are **provider_managed**, **user_managed**.
- `href` - (String) The URL for this snapshot.
- `id` - (String) The unique identifier for this snapshot.
- `lifecycle_state` - (String) The lifecycle state of this snapshot. Supported values are **deleted**, **deleting**, **failed**, **pending**, **stable**, **updating**, **waiting**, **suspended**.
- `minimum_capacity` - (Integer) The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
- `operating_system` - (String) The globally unique name for an Operating System included in this image.
- `resource_type` - (String) The resource type.
- `size` - (Integer) The size of this snapshot rounded up to the next gigabyte.
- `source_image` - (String) If present, the unique identifier for the image from which the data on this volume was most directly provisioned.

## Import

The `ibm_is_snapshot` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_snapshot.example <id>
```

**Example**

```
$ terraform import ibm_is_snapshot.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
