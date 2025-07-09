---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_volume_attachment"
description: |-
  Manages IBM Cloud infrastructure instance volume attachment.
---

# ibm_is_instance_volume_attachment
Create, update, or delete a volume attachment on an existing instance. For more information, about VPC virtual server instances, see [Managing virtual server instances](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-virtual-server-instances).

~> **NOTE**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage (using capacity)

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

resource "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id

  name                               = "example-vol-att-1"
  profile                            = "general-purpose"
  capacity                           = "20"
  delete_volume_on_attachment_delete = true
  delete_volume_on_instance_delete   = true
  volume_name                        = "example-vol-1"

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
resource "ibm_is_volume" "example" {
  name    = "example-vol"
  profile = "10iops-tier"
  zone    = "us-south-2"
}

resource "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id

  name   = "example-col-att-2"
  volume = ibm_is_volume.example.id

  // it is recommended to keep the delete_volume_on_attachment_delete as false here otherwise on deleting attachment, existing volume will also get deleted

  delete_volume_on_attachment_delete = false
  delete_volume_on_instance_delete   = false
}

```
## Example usage (creating new volume)

```terraform
resource "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id

  name                                = "example-col-att-3"
  iops                                = 100
  capacity                            = 50
  delete_volume_on_attachment_delete  = true
  delete_volume_on_instance_delete    = true
  volume_name                         = "example-vol-3"

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

```

## Example usage (restoring using snapshot)

```terraform
resource "ibm_is_instance_volume_attachment" "example-vol-3" {
  instance = ibm_is_instance.example.id

  name = "test-col-att-4"
  profile = "general-purpose"
  snapshot = xxxx-xx-x-xxxxx
  delete_volume_on_attachment_delete = true
  delete_volume_on_instance_delete = true
  volume_name = "example-vol-4"

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

## Timeouts

The `ibm_is_instance_volume_attachment` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:


- **create**: The creation of the instance volume attachment is considered failed when no response is received for 10 minutes.
- **update**: The update of the instance volume attachment or the attachment of a volume to an instance is considered failed when no response is received for 10 minutes.
- **delete**: The deletion of the instance volume attachment is considered failed when no response is received for 10 minutes.

## Argument reference
Review the argument references that you can specify for your resource.

- `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.
- `capacity` - (Optional, Integer) The capacity of the volume in gigabytes.

  ~> **NOTE**
  **&#x2022;** The specified minimum and maximum capacity values for creating or updating volumes may expand in the future. Accepted value is in [10-16000]. </br>**&#x2022;** If unspecified, the capacity will be the source snapshot's `minimum_capacity` when `snapshot` is provided.</br>**&#x2022;** Supports only expansion on update (must not be less than the current volume capacity)</br>**&#x2022;** Can be updated only if volume is attached to an running virtual server instance.</br>**&#x2022;** Stopped instance will be started on update of capacity of the volume.</br>

- `delete_volume_on_attachment_delete` - (Optional, Bool) If set to **true**, when deleting the attachment, the volume will also be deleted. By default it is **true**
- `delete_volume_on_instance_delete` - (Optional, Bool) If set to **true**, when deleting the instance, the volume will also be deleted. By default it is **false**
- `encryption_key` - (Optional, String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource. If this property is not provided but the image is encrypted, the image's encryption_key will be used. Otherwise, the encryption type for the volume will be `provider_managed`.
- `instance` - (Required, String) The id of the instance.
- `iops` - (Optional, Integer) The bandwidth for the new volume.  This value is required for `custom` storage profiles only.

  ~> **NOTE**
      **&#x2022;** `iops` value can be upgraded and downgraged if volume is attached to an running virtual server instance. Stopped instances will be started on update of volume.</br>
      **&#x2022;** This table shows how storage size affects the `iops` ranges:

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
      </br>

- `name` - (Required, String) The name of the volume attachment.
- `profile` - (Optional, String) The globally unique name for this volume profile.

  ~> **NOTE**
        **&#x2022;**  Allowed values are : [`general-purpose`, `5iops-tier`, `10iops-tier`, `custom`].</br>
        **&#x2022;** If `iops` is not present, `general-purpose` is taken as the volume profile.</br>
        **&#x2022;** If `iops` is present, `custom` is taken as the volume profile.</br>
        **&#x2022;** Tiered profiles [`general-purpose`, `5iops-tier`, `10iops-tier`] can be upgraded and downgraded into each other.</br>
        **&#x2022;** Can be updated only if volume is attached to an running virtual server instance.</br>
        **&#x2022;** Stopped instances will be started on update of volume.</br>
- `snapshot` - (Optional, String) The unique identifier for this snapshot from which to clone the new volume. 

  ~> **NOTE**
        **&#x2022;** one of `capacity` or `snapshot` must be present for volume creation.</br>
        **&#x2022;** If `capacity` is not present or less than `minimum_capacity` of the snapshot, `minimum_capacity` is taken as the volume capacity.</br>
- `snapshot_crn` - (Optional, String) The CRN for this snapshot from which to clone the new volume. 

  ~> **NOTE**
        **&#x2022;** one of `capacity` or `snapshot` must be present for volume creation.</br>
        **&#x2022;** If `capacity` is not present or less than `minimum_capacity` of the snapshot, `minimum_capacity` is taken as the volume capacity.</br>
- `volume` - (Optional, String) The unique identifier for the existing volume
- `volume_name` - (Optional, String) The unique user-defined name for this new volume.
- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `device`-  (String) A unique identifier for the device which is exposed to the instance operating system.
- `href` - (String) The URL for this volume attachment.
- `id` - (String) The ID of the instance volume attachment. The ID is composed of `<instance_id>/<volume_attachment_id>`.
- `status` - (String) The status of this volume attachment. Supported values are **attached**, **attaching**, **deleting**, **detaching**.
- `type` - (String) The type of volume attachment. Supported values are **boot**, **data**.
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
