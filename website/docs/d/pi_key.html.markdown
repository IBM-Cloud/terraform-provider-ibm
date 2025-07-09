---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_key"
description: |-
  Manages an SSH key in the Power Virtual Server cloud.
---

# ibm_pi_key

Retrieve information about the SSH key that is used for your Power Systems Virtual Server instance. The SSH key is used to access the instance after it is created. For more information, about [generating and using SSH Keys](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-ssh-key).

## Example Usage

```terraform
data "ibm_pi_key" "ds_instance" {
  pi_key_name          = "terraform-test-key"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
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

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_key_name`  - (Required, String) User defined name for the SSH key or SSH key ID.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) User defined name for the SSH key
- `creation_date` - (String) Date of SSH Key creation.
- `description` - (String) Description of the SSH key.
- `name` - (String) Name of SSH key.
- `primary_workspace` - (Boolean) Indicates if the current workspace owns the ssh key or not.
- `ssh_key` - (String) SSH RSA key.
- `ssh_key_id` - (String) Unique ID of SSH key.
- `visibility` - (String) Visibility of the SSH key.
