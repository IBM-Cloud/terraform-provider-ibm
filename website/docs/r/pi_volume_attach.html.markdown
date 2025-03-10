---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_attach"
description: |-
  Manages IBM Volume Attach in the Power Virtual Server cloud.
---

# ibm_pi_volume_attach

Attaches and Detaches a volume to a Power Systems Virtual Server instance. For more information, about managing volume, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

The following example attaches volume to a power systems virtual server instance.

```terraform
resource "ibm_pi_volume_attach" "testacc_volume_attach"{
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_volume_id = "<id of the volume to attach>"
  pi_instance_id = "<pvm instance id>"
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

ibm_pi_volume_attach provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 15 minutes) Used for attaching volume.
- **delete** - (Default 15 minutes) Used for detaching volume.

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, Forces new resource, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Required, Forces new resource, String) The ID of the pvm instance to attach the volume to.
- `pi_volume_id` - (Required, Forces new resource, String) The ID of the volume to attach.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the volume attach. The ID is composed of `<pi_cloud_instance_id>/<instance_id>/<volume_id>`.
- `status` - (String) The status of the volume.

## Import

The `ibm_pi_volume_attach` resource can be imported by using `pi_cloud_instance_id`, `instance_id` and `volume_id`.

### Example

```bash
terraform import ibm_pi_volume_attach.example d7bec597-4726-451f-8a63-e62e6f19c32c/49fba6c9-23f8-40bc-9899-aca322ee7d5b/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
