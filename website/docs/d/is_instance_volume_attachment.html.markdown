---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance Volume Attachment"
description: |-
  Manages IBM Cloud infrastructure instance volume attachment.
---

# ibm_is_instance_volume_attachment
Retrieve information of an existing IBM Cloud infrastructure instance volume attachment as a read-only data source. For more information, about VPC virtual server instances, see [Managing virtual server instances](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-virtual-server-instances).

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

data "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id
  name = "example-instance-volume-attachment"
}

```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the volume attachment.
- `instance` - (Required, String) The ID of the instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.
- `delete_volume_on_instance_delete` - (Boolean) If set to true, when deleting the instance the volume will also be deleted.
- `device`-  (String) A unique identifier for the device which is exposed to the instance operating system.
- `href` - (String) The URL for this volume attachment.
- `id` - (String) The ID of the instance volume attachment. The ID is composed of `<instance_id>/<volume_attachment_id>`.
- `name`-  (String) The user-defined name for this volume attachment.
- `status` - (String) The status of this volume attachment. 

  -> **Supported Status** 
    &#x2022; attached</br>
    &#x2022; attaching</br>
    &#x2022; deleting</br>
    &#x2022; detaching

- `type` - (String) The type of volume attachment. Supported values are **boot**, **data**.
- `volume_attachment_id` - (String) The unique identifier for this volume attachment.
- `volume` - (List) The attached volume.

  Nested scheme for `volume`:
  - `crn` - (String) The CRN for this volume.
  - `deleted` - (String) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
  - `href` - (String) The URL for this volume.
  - `id` - (String) The unique identifier for this volume.
  - `name` - (String) The unique user-defined name for this volume.
