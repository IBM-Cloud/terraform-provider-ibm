---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_instance_vpmem_volume"
description: |-
  Get information about PVM instance ID vPMEM Volume
---

# ibm_pi_instance_vpmem_volume

Retrieves information about a power virtual machine instance vPMEM volume.

## Example Usage

```terraform
data "ibm_pi_instance_vpmem_volume" "instance_vpmem_volume" {
    pi_cloud_instance_id = "098f6bcd-2f7e-470a-a1ab-664e61882371"
    pi_pvm_instance_id   = "11223344-5566-7788-99ab-cdef01234567"
    pi_vpmem_volume_id   = "a1b2c3d4-e5f6-7g8h-9i0j-1k2l3m4n5o6p"
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
- `pi_vpmem_volume_id` - (Required, String) vPMEM volume ID.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the pvm instance vpmem volume.
- `creation_date` - (String) The date and time when the volume was created.
- `crn` - (String) The CRN for this resource.
- `error_code` - (String) Error code for the vPMEM volume.
- `href` - (String) Link to vPMEM volume resource.
- `name` - (String) Volume name.
- `pvm_instance_id` - (String) PVM Instance ID which the volume is attached to.
- `reason` - (String) Reason for error.
- `size` - (Float) Volume size (GB).
- `status` - (String) Status of the volume.
- `updated_date` - (String) The date and time when the volume was updated.
- `user_tags` - (List) List of user tags.
