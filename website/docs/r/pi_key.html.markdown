---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_key"
description: |-
  Manages IBM SSH keys in the Power Virtual Server Cloud.
---

# ibm\_pi_key

Provides a SSH key resource. This allows SSH Keys to be created, updated, and cancelled in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a ssh key to be used during creation of a pvminstance:

```hcl
resource "ibm_pi_key" "testacc_sshkey" {
  pi_key_name          = "testkey"
  pi_ssh_key           = "ssh-rsa <value>"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```
## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```hcl
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Timeouts

ibm_pi_key provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a SSH key.
* `delete` - (Default 60 minutes) Used for deleting a SSH key.

## Argument Reference

The following arguments are supported:

* `pi_key_name` - (Required, int) The key name.
* `pi_ssh_key` - (Required, string) The value of the ssh key.
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account
* `pi_creation_date` - (Optional, string) Date of sshkey creation.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the key.The id is composed of \<power_instance_id\>/\<key_name\>.
* `key_id` -  The unique identifier of the key.

## Import

ibm_pi_key can be imported using `power_instance_id` and `ssh_key_name`, eg

```
$ terraform import ibm_pi_key.example d7bec597-4726-451f-8a63-e62e6f19c32c/mykey
```