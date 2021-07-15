---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance Volume Attachment"
description: |-
  Manages IBM Cloud infrastructure instance volume attachment.
---

# ibm_is_instance_volume_attachment
Retrieve information of an existing IBM Cloud Infrastructure instance volume attachment as a read-only data source. For more information, about VPC virtual server instances, see [Managing virtual server instances](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-virtual-server-instances).


## Example usage

```terraform

data "ibm_is_instance_volume_attachment" "ds_vsi_va" {
  instance = "xx-x-x-x-xxxxx
  name = "test-volume"
}

```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the volume attachment.
- `instance` - (Required, String) The id of the instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `delete_volume_on_instance_delete` - (Boolean) If set to true, when deleting the instance the volume will also be deleted.
- `device`-  (String) A unique identifier for the device which is exposed to the instance operating system.
- `href` - (String) The URL for this volume attachment.
- `id` - (String) The ID of the instance volume attachment. The ID is composed of `<instance_id>/<volume_attachment_id>`.
- `name`-  (String) The user-defined name for this volume attachment.
- `status` - (String) The status of this volume attachment [ attached, attaching, deleting, detaching ].
- `type` - (String) The type of volume attachment [ boot, data ].
- `volume_attachment_id` - (String) The unique identifier for this volume attachment.
- `volume` - (List) The attached volume.

  Nested scheme for `volume`:
  - `crn` - (String) The CRN for this volume.
  - `deleted` - (String) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
  - `href` - (String) The URL for this volume.
  - `id` - (String) The unique identifier for this volume.
  - `name` - (String) The unique user-defined name for this volume.
