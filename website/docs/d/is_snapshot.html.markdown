---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshot"
description: |-
  Reads IBM Cloud snapshots.
---
# ibm\_is_snapshot

Import the details of existing IBM Cloud Infrastructure snapshot as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.For more information, about infrastructure snapshots, see [viewing snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-view).


## Example Usage

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            			    = "testsubnet"
  vpc             			    = ibm_is_vpc.testacc_vpc.id
  zone            			    = "us-south-2"
  total_ipv4_address_count 	= 16
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testvsi"
  image   = "xxxxx-xxxxx-xxxxx-xxxxxx"
  profile = "bx2-2x8"
  primary_network_interface {
    subnet     = ibm_is_subnet.testacc_subnet.id
  }
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
  network_interfaces {
    subnet = ibm_is_subnet.testacc_subnet.id
    name   = "eth1"
  }
}
resource "ibm_is_snapshot" "testacc_snapshot" {
  name            = "testsnapshot"
  source_volume   = ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
}

data "ibm_is_snapshot" "ds_snapshot1" {
    identifier = ibm_is_snapshot.testacc_snapshot.id
}

```

```terraform

data "ibm_is_snapshot" "ds_snapshot2" {
    name = ibm_is_snapshot.testacc_snapshot.name
}

```


## Argument reference
Review the argument references that you can specify for your data source. 

- `identifier` - (Optional, String) The unique identifier for this snapshot.
- `name` - (Optional, String) The name of the snapshot.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.

- `bootable` - (Bool) Indicates if a boot volume attachment can be created with a volume created from this snapshot.
- `crn` - (String) The CRN for this snapshot.
- `encryption` - (String) The type of encryption used on the source volume(One of [ provider_managed, user_managed ]).
- `href` - (String) The URL for this snapshot.
- `lifecycle_state` - (String) The lifecycle state of this snapshot. Supported values are **deleted**, **deleting**, **failed**, **pending**, **stable**, **updating**, **waiting**, **suspended**.
- `minimum_capacity` - (Integer) The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
- `operating_system` - (String) The globally unique name for the operating system included in this image.
- `resource_type` - (String) The resource type.
- `size` - (Integer) The size of this snapshot rounded up to the next gigabyte.
- `source_image` - (String) If present, the unique identifier for the image from which the data on this volume was most directly provisioned.
