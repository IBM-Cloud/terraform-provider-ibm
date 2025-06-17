---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_catalog_images"
description: |-
  List all images available for copying into the cloud instances
---

# ibm_pi_catalog_images

Retrieve the details of an image that you can use in your Power Systems Virtual Server instance for copying into IBM Cloud instances. For more information, about catalog images, see [provisioning a virtual server instance from a third-party image](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-ordering-3P).

## Example Usage

The following example shows how to retrieve information using `ibm_pi_catalog_images`.

```terraform
data "ibm_pi_catalog_images" "ds_images" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument Reference

Review the argument reference that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `sap` - (Optional, Bool) Set `true` to include SAP images. The default value is `false`.
- `vtl` - (Optional, Bool) Set `true` to include VTL images. The default value is `false`.

## Attribute Reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `images`- (List) Lists all the images in the IBM Power Virtual Server Cloud.

  Nested scheme for `images`:
  - `architecture` - (String) The CPU architecture that the image is designed for.
  - `container_format` - (String) The container format.
  - `creation_date` - (String) Date of image creation.
  - `crn` - (String) The CRN of this resource.
  - `description` - (String) The description of an image.
  - `disk_format` - (String) The disk format.
  - `endianness` - (String) The `Endianness` order.
  - `href` - (String) The `href` of an image.
  - `hypervisor_type` - (String) Hypervisor type.
  - `image_id` - (String) The unique identifier of an image.
  - `image_type` - (String) The identifier of this image type.
  - `last_update_date` - (String) The last updated date of an image.
  - `name` - (String) The name of the image.
  - `operating_system` - (String)  Operating System.
  - `state` - (String) The state of an Operating System.
  - `storage_pool` - (String) Storage pool where image resides.
  - `storage_type` - (String) The storage type of an image.
