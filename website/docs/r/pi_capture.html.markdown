---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_capture"
description: |-
  Manages IBM Capture instance in the Power Virtual Server cloud.
---

# ibm_pi_capture
Create or delete for Capture Power System Virtual Server Instance

**Note:-**
If `pi_capture_destination` is `Cloud-Storage` then delete bucket object functionality not supported by this resource , hence user need to delete bucket object manually from `Cloud Storage bucket`.

For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).
## Example usage
The following example creates a Capture Power System Virtual Server Instance.

```terraform
resource "ibm_pi_capture" "test_capture" {
        pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
        pi_capture_name = "terraform-test-capture"
        pi_instance_name = "terraform-test-instance"
        pi_capture_destination = "image-catalog"
}
```
```terraform
resource "ibm_pi_capture" "test_capture" {
	pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
	pi_capture_name = "test-capture"
	pi_instance_name = "test-vm"
	pi_capture_destination = "cloud-storage"
	pi_capture_cloud_storage_region = "us-east"
	pi_capture_cloud_storage_access_key = "<Cloud Storage Access key>"
	pi_capture_cloud_storage_secret_key = "<Cloud Storage Secret key>"
	pi_capture_storage_image_path = "test-bucket"
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

ibm_pi_capture provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 75 minutes) Used for creating capture instance .
- **delete** - (Default 50 minutes) Used for deleting capture instance.

## Argument reference 
Review the argument references that you can specify for your resource. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_capture_name` - (Required, String) Name of the deployable image created for the captured PVMInstance.
- `pi_instance_name` - (Required, String) The name of the instance.
- `pi_capture_destination`- (Required, String) Destination for the deployable image.
`[image-catalog,cloud-storage,both]`
- `pi_capture_volume_ids`- (Optional, List of String)  List of Data volume IDs to include in the captured   PVMInstance.
- `pi_capture_cloud_storage_region`- (Optional,String) Cloud Storage Region
- `pi_capture_cloud_storage_access_key`- (Optional,String) Cloud Storage Access key
- `pi_capture_cloud_storage_secret_key`- (Optional,String) Cloud Storage Secret key
- `pi_capture_storage_image_path` - (Optional,String) Cloud Storage Image Path (bucket-name [/folder/../..])


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The image id of the capture instance. The ID is composed of `<pi_cloud_instance_id>/<pi_capture_name>/<pi_capture_destination>`.
- `image_id` - (String) The image id of the capture instance.


## Import

The `ibm_pi_capture` resource can be imported by using `pi_cloud_instance_id` `pi_capture_name` and `pi_capture_destination`.

**Example**
```
$ terraform import ibm_pi_capture.example d7bec597-4726-451f-8a63-e62e6f19c32c/test-capture/image-catalog

```

