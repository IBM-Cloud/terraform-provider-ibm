---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_volumes_delete"
description: |-
  Manages volumes in Power Virtual Server cloud.
---

# ibm_pi_volumes_delete

Delete multiple volumes for a Power Systems Virtual Server resources, see [volumes](https://cloud.ibm.com/apidocs/power-cloud#pcloud-v2-pvminstances-volumes-post)

## Example Usage

```terraform
resource "ibm_pi_volumes_delete" "pi_volumes_delete_instance" {
    pi_cloud_instance_id = "<value of the cloud_instance_id>"
    pi_volume_ids        = ["600850d6-4b38-40cf-857d-123456","d783bb14-efeb-4ded-9847-123456"]
}
```

## Notes

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

ibm_pi_volumes_delete provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- `delete` - (Default 30 minutes) Used for deleting volumes in power virtual server.

## Argument Reference

You can specify the following arguments for this resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_ids` - (Required, List) List of volumes to be deleted.
