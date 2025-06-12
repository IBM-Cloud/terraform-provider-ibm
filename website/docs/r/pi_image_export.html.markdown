---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_image_export"
description: |-
  Exports IBM Image to IBM Cloud Object Storage in the Power Virtual Server cloud.
---

# ibm_pi_image_export

Export an image to IBM Cloud Object Storage for Power Systems Virtual Server instance. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example enables you to export an image:

```terraform
resource "ibm_pi_image_export" "testacc_image_export"{
  pi_cloud_instance_id   = "<value of the cloud_instance_id>"
  pi_image_id            = "test_image"
  pi_image_access_key    = "dummy-access-key"
  pi_image_bucket_name   = "images-public-bucket"
  pi_image_bucket_region = "us-south"
  pi_image_secret_key    = "dummy-secret-key"
}
```

### Notes

- Ensure the exported file is cleaned up manually from the Cloud Object Storage when no longer needed. Power Systems Virtual Server does not support deleting the exported image. Updating any attribute will result in creating a new Export job.
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
  
## Timeouts

The `ibm_pi_image_export` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) used for exporting image to IBM Cloud Object Storage bucked. Considered failed if no response is received by timeout.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_image_access_key` - (Required, String, Sensitive) The Cloud Object Storage access key; required for buckets with private access.
- `pi_image_bucket_name` - (Required, String) The Cloud Object Storage bucket name; `bucket-name[/optional/folder]`
- `pi_image_bucket_region` - (Required, String) The Cloud Object Storage region. Supported COS regions are:`au-syd`, `br-sao`, `ca-tor`, `che01`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`.
- `pi_image_id` - (Required, String) The Image ID of existing source image; required for image export.
- `pi_image_secret_key` - (Required, String, Sensitive) The Cloud Object Storage secret key; required for buckets with private access.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of an image export resource. The ID is composed of `<image_id>/<bucket_name>/<bucket_region>`.
