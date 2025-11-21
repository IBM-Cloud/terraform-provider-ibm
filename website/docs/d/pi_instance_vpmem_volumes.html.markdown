---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_instance_vpmem_volumes"
description: |-
  Get information about PVM instance ID vPMEM Volumes
---

# ibm_pi_instance_vpmem_volumes

Retrieves information about a power virtual machine instance vPMEM volumes.

## Example Usage

```terraform
data "ibm_pi_instance_vpmem_volumes" "instance_vpmem_volumes" {
    pi_cloud_instance_id = "098f6bcd-2f7e-470a-a1ab-664e61882371"
    pi_pvm_instance_id   = "11223344-5566-7788-99ab-cdef01234567"
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

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_pvm_instance_id` - (Required, String) PCloud PVM instance ID.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the pi_instance_vpmem_volumes.
- `volumes` - (List) List of vPMEM volumes.
   Nested schema for `volumes`:
  - `creation_date` - (String) The date and time when the volume was created.
  - `crn` - (String) The CRN for this resource.
  - `error_code` - (String) Error code for the vPMEM volume.
  - `href` - (String) Link to vPMEM volume resource.
  - `name` - (String) Volume Name.
  - `pvm_instance_id` - (String) PVM Instance ID which the volume is attached to.
  - `size` - (Float) Volume Size (GB).
  - `status` - (String) Status of the volume.
  - `updated_date` - (String) The date and time when the volume was updated.
  - `user_tags` - (List) List of user tags.
  - `volume_id` - (String) Volume ID.
