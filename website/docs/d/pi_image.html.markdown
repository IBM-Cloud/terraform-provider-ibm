---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_image"
description: |-
  Manages an image in the Power Virtual Server cloud.
---

# ibm_pi_image

Import the details of an existing IBM Power Virtual Server Cloud image as a read-only data source. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_image" "ds_image" {
  pi_image_name        = "7200-03-03"
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

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_image_name` - (Required, String) The ID of the image. To find supported images, run the `ibmcloud pi images` command.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `architecture` - (String) The CPU architecture that the image is designed for.
- `crn` - (String) The CRN of this resource.
- `hypervisor` - (String) Hypervisor type.
- `id` - (String) The unique identifier of the image.
- `image_type` - (String) The identifier of this image type.
- `operating_system` - (String) The operating system that is installed with the image.
- `size` - (String) The size of the image in megabytes.
- `source_checksum` - (String) Checksum of the image.
- `state` - (String) The state for this image.
- `storage_pool` - (String) Storage pool where image resides.
- `storage_type` - (String) The storage type for this image.
- `user_tags` - (List) List of user tags attached to the resource.
