---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_volume_attachment"
description: |-
  Manages IBM Cloud infrastructure instance volume attachment.
---

# ibm_is_vpc_route
Create, update, or delete a Volume attachment on an instance. For more information, about VPC virtual server instances, see [Managing virtual server instances](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-virtual-server-instances).

## Example usage (using capacity)

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testvsi1"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "bc1-2x8"
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

resource "ibm_is_instance_volume_attachment" "testacc_att1" {
  instance = ibm_is_instance.testacc_instance.id

  name = "test-vol-att-1"
  profile = "general-purpose"
  capacity = "20"
  delete_volume_on_attachment_delete = true
  delete_volume_on_instance_delete = true
  volume_name = "testvol1"

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

```
## Example usage (using existing volume)

```terraform
resource "ibm_is_volume" "testacc_vol" {
  name    = "testvol2"
  profile = "10iops-tier"
  zone    = "us-south-2"
}

resource "ibm_is_instance_volume_attachment" "testacc_att2" {
  instance = ibm_is_instance.testacc_instance.id

  name = "test-col-att-2"
  volume = ibm_is_volume.testacc_vol.id
  delete_volume_on_attachment_delete = false
  delete_volume_on_instance_delete = false
}

```
## Example usage (restoring using snapshot)

```terraform
resource "ibm_is_instance_volume_attachment" "testacc_att3" {
  instance = ibm_is_instance.testacc_instance.id

  name = "test-col-att-3"
  profile = "general-purpose"
  snapshot = xxxx-xx-x-xxxxx
  delete_volume_on_attachment_delete = true
  delete_volume_on_instance_delete = true
  volume_name = "testvol3"

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

```

## Timeouts

The `ibm_is_instance_volume_attachment` resource provides the following [[Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:


- **create**: The creation of the instance volume attachment is considered failed when no response is received for 10 minutes.
- **update**: The update of the instance volume attachment or the attachment of a volume to an instance is considered failed when no response is received for 10 minutes.
- **delete**: The deletion of the instance volume attachment is considered failed when no response is received for 10 minutes.



## Argument reference
Review the argument references that you can specify for your resource. 

- `capacity` - (Optional, Integer) The capacity of the volume in gigabytes. **NOTE** The specified minimum and maximum capacity values for creating or updating volumes may expand in the future. Accepted value is in [10-2000].
  - If this property is not provided or less than the minimum_capacity, minimum_capacity of the snapshot will be used as the capacity for the volume.

- `delete_volume_on_attachment_delete` - (Optional, Bool) If set to true, when deleting the attachment, the volume will also be deleted
- `delete_volume_on_instance_delete` - (Optional, Bool) If set to true, when deleting the instance, the volume will also be deleted
- `encryption_key` - (Optional, String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource. If this property is not provided but the image is encrypted, the image's encryption_key will be used. Otherwise, the encryption type for the volume will be `provider_managed`.
- `instance` - (Required, String) The id of the instance.
- `iops` - (Optional, Integer) The bandwidth for the new volume
- `name` - (Required, String) The name of the volume attachment.
- `profile` - (Optional, String) The globally unique name for this volume profile
- `snapshot` - (Optional, String) The unique identifier for this snapshot from which to clone the new volume. 

  **NOTE**
  - one of `capacity` or `snapshot` must be present for volume creation
  - if `capacity` is not present or less than `minimum_capacity` of the snapshot, `minimum_cpacity` is taken as the volume capacity.
- `volume` - (Optional, String) The unique identifier for the existing volume
- `volume_name` - (Optional, String) he unique user-defined name for this new volume.                     
           
## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `device`-  (String) A unique identifier for the device which is exposed to the instance operating system.
- `href` - (String) The URL for this volume attachment.
- `id` - (String) The ID of the instance volume attachment. The ID is composed of `<instance_id>/<volume_attachment_id>`.
- `status` - (String) The status of this volume attachment [ attached, attaching, deleting, detaching ].
- `type` - (String) The type of volume attachment [ boot, data ].
- `volume_attachment_id` - (String) The unique identifier for this volume attachment.
- `volume_crn` - (String) The CRN for this volume.
- `volume_deleted` - (String) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
- `volume_href` - (String) The URL for this volume.


## Import
The `ibm_is_instance_volume_attachment` resource can be imported by using the instance id and volume attachment id. 

**Syntax**

```
$ terraform import ibm_is_instance_volume_attachment.example <instance_id>/<volume_attachment_id>
```
