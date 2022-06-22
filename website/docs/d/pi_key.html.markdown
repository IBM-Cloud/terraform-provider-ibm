---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_key"
description: |-
  Manages an key in the Power Virtual Server cloud.
---

# ibm_pi_key
Retrieve information about the SSH key that is used for your Power Systems Virtual Server instance. The SSH key is used to access the instance after it is created. For more information, about [configuring your IBM i virtual machine (VM)](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-configuring-ibmi).

## Example usage

```terraform
data "ibm_pi_key" "ds_instance" {
  pi_key_name          = "terraform-test-key"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

**Notes**

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
  
## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_key_name` - (Required, String) The name of the key.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `creation_date` - Timestamp - The timestamp when the SSH key was created.
- `id` - (String) The unique identifier of the SSH key.
- `ssh_key` - (String) The public SSH key value.
