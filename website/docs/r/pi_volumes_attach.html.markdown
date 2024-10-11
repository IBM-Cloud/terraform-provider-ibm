---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_volumes_attach"
description: |-
  Manages volumes in Power Virtual Server cloud.
---

# ibm_pi_volumes_attach

Attach volumes to a Power Systems Virtual Server instance, see [volumes](https://cloud.ibm.com/apidocs/power-cloud#pcloud-v2-pvminstances-volumes-post)

## Example Usage

The following example attach volumes to power virtual instance.

```terraform
resource "ibm_pi_volumes_attach" "pi_volumes_attach_instance" {
    pi_cloud_instance_id = "<value of the cloud_instance_id>"
    pi_instance_id       = "<value of the instance_id>"
    pi_volume_ids        = ["volume ids"]
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

## Timeouts

`pi_volumes_attach` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- `create` - (Default 10 minutes) Used for attaching volumes to a power virtual server instance.
- `delete` - (Default 10 minutes) Used for detaching volumes to a power virtual server instance.

## Argument Reference

You can specify the following arguments for this resource.

- `pi_boot_volume_id` - (Optional, String) Primary Boot Volume ID.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Required, String)  The unique identifier of an instance.
- `pi_volume_ids` - (Required, List) List of volumes to be attached to a `pi_instance`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The unique identifier of volumes attach. The ID is composed of `<pi_cloud_instance_id>/<pi_instance_id>/<pi_volume_ids>`
