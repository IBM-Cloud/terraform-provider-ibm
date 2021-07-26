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

- `pi_image_name` - (Required, String) The name of an image.
- `pi_image_id` - (Required, String) The ID of an image. 
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of an image.
- `image_id` - (String) The unique identifier of an image.

## Import

The `ibm_pi_image` can be imported by using `power_instance_id` and `image_id`.

**Example**

```
$ terraform import ibm_pi_image.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
