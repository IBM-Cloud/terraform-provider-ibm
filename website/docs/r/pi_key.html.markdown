---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_key"
description: |-
  Manages IBM SSH keys in the Power Virtual Server cloud.
---

# ibm_pi_key
Create, update, or delete an SSH key for your Power Systems Virtual Server instance. The SSH key is used to access the instance after it is created. For more information, about SSH keys in Power Virtual Server, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage
The following example enables you to create a SSH key to be used during creation of a power virtual server instance:

```terraform
resource "ibm_pi_key" "testacc_sshkey" {
  pi_key_name          = "testkey"
  pi_ssh_key           = "ssh-rsa <value>"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

**Note**
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

## Timeouts

The `ibm_pi_key` resource provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- **create** - (Default 60 minutes) Used for creating a SSH key.
- **delete** - (Default 60 minutes) Used for deleting a SSH key.


## Argument reference
Review the argument references that you can specify for your resource. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_creation_date` - (Optional, String) The date when the SSH key was created. 
- `pi_key_name`  - (Required, Integer) The name of the SSH key that you uploaded to IBM Cloud. 
- `pi_ssh_key` - (Required, String) The value of the public SSH key. 


## Attribute reference
 In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the key. The ID is composed of `<power_instance_id>/<key_name>`.
- `key_id` - (String) The unique identifier of the key.

## Import

The `ibm_pi_key` resource can be imported by using `power_instance_id` and `ssh_key_name`.

**Example**

```
$ terraform import ibm_pi_key.example d7bec597-4726-451f-8a63-e62e6f19c32c/mykey
```
