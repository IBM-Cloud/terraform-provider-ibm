---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_image"
description: |-
  Manages IBM Image in the Power Virtual Server cloud.
---

# ibm_pi_image
Create, update, or delete for a Power Systems Virtual Server image. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage
The following example enables you to create a image:

```terraform
resource "ibm_pi_image" "testacc_image  "{
  pi_image_name       = "7200-03-02"
  pi_image_id         = <"image id obtained from the datasource">
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

```terraform
resource "ibm_pi_image" "testacc_image  "{
  pi_image_name       = "test_image"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_image_bucket_name = "images-public-bucket"
  pi_image_bucket_access = "public"
  pi_image_bucket_region = "us-south"
  pi_image_bucket_file_name = "rhcos-48-07222021.ova.gz"
  pi_image_storage_type = "tier1"
}
```

**Note**
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  
  Example usage:
  
  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Timeouts

The   ibm_pi_image   provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The creation of the image is considered failed if no response is received for 60 minutes. 
- **Delete** The deletion of the image is considered failed if no response is received for 60 minutes. 

## Argument reference
Review the argument references that you can specify for your resource. 

- `pi_affinity_instance` - (Optional, String) PVM Instance (ID or Name) to base storage affinity policy against; required if requesting `affinity` and `pi_affinity_volume` is not provided.
- `pi_affinity_policy` - (Optional, String) Affinity policy for image; ignored if `pi_image_storage_pool` provided; for policy affinity requires one of `pi_affinity_instance` or `pi_affinity_volume` to be specified; for policy anti-affinity requires one of `pi_anti_affinity_instances` or `pi_anti_affinity_volumes` to be specified; Allowable values: `affinity`, `anti-affinity`
- `pi_affinity_volume`- (Optional, String) Volume (ID or Name) to base storage affinity policy against; required if requesting `affinity` and `pi_affinity_instance` is not provided.
- `pi_anti_affinity_instances` - (Optional, String) List of pvmInstances to base storage anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_volumes` is not provided.
- `pi_anti_affinity_volumes`- (Optional, String) List of volumes to base storage anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_instances` is not provided.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_image_name` - (Required, String) The name of an image.
- `pi_image_id` - (Optional, String) Image ID of existing source image; required for copy image.
  - Either `pi_image_id` or `pi_image_bucket_name` is required.
- `pi_image_bucket_name` - (Optional, String) Cloud Object Storage bucket name; `bucket-name[/optional/folder]`
  - Either `pi_image_bucket_name` or `pi_image_id` is required.
- `pi_image_access_key` - (Optional, String, Sensitive) Cloud Object Storage access key; required for buckets with private access.
  - `pi_image_access_key` is required with `pi_image_secret_key`
- `pi_image_bucket_access` - (Optional, String) Indicates if the bucket has public or private access. The default value is `public`.
- `pi_image_bucket_file_name` - (Optional, String) Cloud Object Storage image filename
  - `pi_image_bucket_file_name` is required with `pi_image_bucket_name`
- `pi_image_bucket_region` - (Optional, String) Cloud Object Storage region
  - `pi_image_bucket_region` is required with `pi_image_bucket_name`
- `pi_image_secret_key` - (Optional, String, Sensitive) Cloud Object Storage secret key; required for buckets with private access.
  - `pi_image_secret_key` is required with `pi_image_access_key`
- `pi_image_storage_pool` - (Optional, String) Storage pool where the image will be loaded, if provided then `pi_image_storage_type` and `pi_affinity_policy` will be ignored.
- `pi_image_storage_type` - (Optional, String) Type of storage. Will be ignored if `pi_image_storage_pool` or `pi_affinity_policy` is provided. If only using `pi_image_storage_type` for storage selection then the storage pool with the most available space will be selected.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of an image. The ID is composed of `<pi_cloud_instance_id>/<image_id>`. 
- `image_id` - (String) The unique identifier of an image.

## Import

The `ibm_pi_image` can be imported by using `pi_cloud_instance_id` and `image_id`.

**Example**

```
$ terraform import ibm_pi_image.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
