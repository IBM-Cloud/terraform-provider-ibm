---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_key"
description: |-
  Manages an key in the Power Virtual Server cloud.
---

# ibm_pi_key
Retrieve information about the SSH key that is used for your Power Systems Virtual Server instance. The SSH key is used to access the instance after it is created. For more information, about [configuring your IBM virtual machine (VM)](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-configuring-ibmi).

## Example usage

```terraform
data "ibm_pi_key" "ds_instance" {
  pi_key_name          = "terraform-test-key"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```
  
## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) Cloud Instance ID of a PCloud Instance.
- `pi_key_name`  - (Required, String) User defined name for the SSH key. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) User defined name for the SSH key
- `creation_date` - (String) Date of SSH Key creation. 
- `ssh_key` - (String) SSH RSA key.

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