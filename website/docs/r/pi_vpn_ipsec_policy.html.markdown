---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_ipsec_policy"
description: |-
  Manages IBM IPSec Policy in the Power Virtual Server cloud.
---

# ibm_pi_ipsec_policy
Create, update, or delete a IPSec Policy. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage
The following example creates a IPSec Policy.

```terraform
	resource "ibm_pi_ipsec_policy" "example" {
		pi_cloud_instance_id    = "<value of the cloud_instance_id>"
		pi_policy_name          = "test"
		pi_policy_dh_group = 1
		pi_policy_encryption = "3des-cbc"
		pi_policy_key_lifetime = 180
		pi_policy_pfs = true
		pi_policy_authentication = "hmac-md5-96"
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

ibm_pi_ipsec_policy provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating IPSec Policy.
- **update** - (Default 10 minutes) Used for updating IPSec Policy.
- **delete** - (Default 10 minutes) Used for deleting IPSec Policy.

## Argument reference 
Review the argument references that you can specify for your resource. 
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_policy_authentication`  - (Optional, String) Authentication for the IPSec Policy. Supported values are `none`(Default), `sha-256`, `sha-384`, and `sha1`.
- `pi_policy_dh_group` - (Required, Integer) DH group of the IPSec Policy. Supported values are `1`,`2`,`5`,`14`,`19`,`20`,`24`.
- `pi_policy_encryption`- (Required, String) Encryption of the IPSec Policy. Supported values are `3des-cbc`,`aes-128-cbc`,`aes-128-gcm`,`aes-192-cbc`,`aes-256-cbc`,`aes-256-gcm`,`des-cbc`.
- `pi_policy_key_lifetime` - (Required, Integer) Policy key lifetime. Supported values:  `180` ≤ value ≤ `86400`.
- `pi_policy_name` - (Required, String) Name of the IPSec Policy.
- `pi_policy_pfs` - (Required, Boolean) Perfect Forward Secrecy.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the IPSec Policy. The ID is composed of `<power_instance_id>/<policy_id>`.
- `policy_id` - (String) IPSec Policy ID.

## Import

The `ibm_pi_ipsec_policy` resource can be imported by using `power_instance_id` and `policy_id`.

**Example**

```
$ terraform import ibm_pi_ipsec_policy.example d7bec597-4726-451f-8a63-e62e6f19c32c/ffag151a-bc0a-4438-9f8a-b0760bbf4u1u
```
