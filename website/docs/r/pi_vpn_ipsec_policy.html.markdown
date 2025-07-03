---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_ipsec_policy"
description: |-
  Manages IBM IPSec Policy in the Power Virtual Server cloud.
---

# ibm_pi_ipsec_policy

~> This resource is deprecated and will be removed in the next major version. This resource has reached end of life.

Create, update, or delete a IPSec Policy. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example creates a IPSec Policy.

```terraform
  resource "ibm_pi_ipsec_policy" "example" {
    pi_cloud_instance_id    = "<value of the cloud_instance_id>"
    pi_policy_name          = "test"
    pi_policy_dh_group = 1
    pi_policy_encryption = "aes-256-cbc"
    pi_policy_key_lifetime = 28800
    pi_policy_pfs = true
    pi_policy_authentication = "hmac-sha-256-128"
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

ibm_pi_ipsec_policy provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating IPSec Policy.
- **update** - (Default 10 minutes) Used for updating IPSec Policy.
- **delete** - (Default 10 minutes) Used for deleting IPSec Policy.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_policy_authentication`  - (Optional, String) Authentication for the IPSec Policy. Supported values are `none`(Default), `hmac-sha-256-128`, `hmac-sha1-96`.
- `pi_policy_dh_group` - (Required, Integer) DH group of the IPSec Policy. Supported values are `1`,`2`,`5`,`14`,`19`,`20`,`24`.
- `pi_policy_encryption`- (Required, String) Encryption of the IPSec Policy. Supported values are `aes-256-cbc`, `aes-192-cbc`, `aes-128-cbc`, `aes-256-gcm`, `aes-128-gcm`, `3des-cbc`.
- `pi_policy_key_lifetime` - (Required, Integer) Policy key lifetime. Supported values:  `180` ≤ value ≤ `86400`.
- `pi_policy_name` - (Required, String) Name of the IPSec Policy.
- `pi_policy_pfs` - (Required, Boolean) Perfect Forward Secrecy.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the IPSec Policy. The ID is composed of `<power_instance_id>/<policy_id>`.
- `policy_id` - (String) IPSec Policy ID.

## Import

The `ibm_pi_ipsec_policy` resource can be imported by using `power_instance_id` and `policy_id`.

### Example

```bash
terraform import ibm_pi_ipsec_policy.example d7bec597-4726-451f-8a63-e62e6f19c32c/ffag151a-bc0a-4438-9f8a-b0760bbf4u1u
```
