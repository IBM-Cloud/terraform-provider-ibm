---
layout: "ibm"
page_title: "IBM : ibm_pi_instance_vpmem_volumes"
description: |-
  Manages pi_instance_vpmem_volumes.
subcategory: "Power Systems"
---

# ibm_pi_instance_vpmem_volumes

Create, update, and delete pi_instance_vpmem_volumes with this resource.

## Example Usage

```terraform
resource "ibm_pi_instance_vpmem_volumes" "instance_vpmem_volumes" {
  pi_cloud_instance_id = "cloud_instance_id"
  pi_pvm_instance_id   = "pvm_instance_id"
  pi_vpmem_volumes {
      name = "name"
      size = 1
  }
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

The `ibm_pi_instance_vpmem_volumes` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for attaching a vpmem volume to an instance.
- **delete** - (Default 10 minutes) Used for dettaching a vpmem volume to an instance.
  
## Argument Reference

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, Forces new resource, String) Cloud Instance ID of a PCloud Instance.
- `pi_pvm_instance_id` - (Required, Forces new resource, String) PCloud PVM Instance ID.
- `pi_user_tags` - (Optional, Forces new resource, List) List of user tags.
- `pi_vpmem_volumes` - (Required, Forces new resource, List)
   Nested schema for `pi_vpmem_volumes`:
  - `name` - (Required, String) Volume base name.
  - `size` - (Required, Integer) Volume size (GB).

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the pi_instance_vpmem_volumes.
- `volumes` - (List) List of vPMEM volumes.
   Nested schema for `volumes`:
  - `creation_date` - (String) The date and time when the volume was created.
  - `crn` - (String) The CRN for this resource.
  - `href` - (String) Link to vPMEM volume resource.
  - `name` - (String) Volume Name.
  - `pvm_instance_id` - (String) PVM Instance ID which the volume is attached to.
  - `size` - (Float) Volume Size (GB).
  - `status` - (String) Status of the volume.
  - `updated_date` - (String) The date and time when the volume was updated.
  - `user_tags` - (List) List of user tags.
  - `volume_id` - (String) Volume ID.

## Import

The `ibm_pi_instance_vpmem_volumes` resource can be imported by using `pi_cloud_instance_id` `pi_pvm_instance_id`, and `volume_id`.

### Example

```bash
terraform import ibm_pi_instance_vpmem_volumes.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/cea6651a-4726-451f-8a63--e62e6f19c32c
```
