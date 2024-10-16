---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_volumes_detach"
description: |-
  Manages volumes in Power Virtual Server cloud.
---

# ibm_pi_volumes_detach

Detach volumes to a Power Systems Virtual Server instance, see [volumes](https://cloud.ibm.com/apidocs/power-cloud#pcloud-v2-pvminstances-volumes-post)

## Example Usage

The following example detach volumes from a power virtual instance.

```terraform
resource "ibm_pi_volumes_detach" "pi_volumes_detach_instance" {
    pi_cloud_instance_id  = "<value of the cloud_instance_id>"
    pi_instance_id        = "<value of the instance_id>"
    pi_detach_all_volumes = true
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

`pi_volumes_detach` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- `delete` - (Default 10 minutes) Used for detaching volumes from a power virtual server instance.

## Argument Reference

You can specify the following arguments for this resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_detach_all_volumes` - (Optional, Boolean) Indicates if all volumes, except primary boot volume, attached to the `pi_instance` should be detached (default=`false`); required if `pi_volume_ids` is not provided.
- `pi_detach_primary_boot_volume` - (Optional, Boolean) Indicates if primary boot volume attached to the `pi_instance` should be detached (default=`false`).
- `pi_instance_id` - (Required, String) The unique identifier of an instance.
- `pi_volume_ids` - (Optional, List) List of volumes to be detached from a `pi_instance`; required if `pi_detach_all_volumes` is not provided.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The unique identifier of volumes detach. The ID is composed of `<pi_cloud_instance_id>/<pi_instance_id>`
