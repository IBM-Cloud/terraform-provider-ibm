---
layout: "ibm"
page_title: "IBM : ibm_sm_admin_token"
description: |-
  Manages Vault admin token for an SM dedicated instance.
subcategory: "IBM Cloud Secrets Manager Instance Management API"
---

# ibm_sm_admin_token

A resource for generating a Vault admin token for authenticating to your Vault Dedicated cluster. The token is valid for 1 hour and grants administrative privileges. The token is automatically refreshed if it expired or it is about to expire.

## Example Usage

```hcl
data "ibm_sm_admin_token" "sm_admin_token" {
	instance_id = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The service instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_admin_token.
* `created_at` - (String) The date that the admin token was created. The date format follows RFC 3339.
* # `token` - (String) The token value.
  * Constraints: The maximum length is `4096` characters. The minimum length is `16` characters. The value must match regular expression `/^hvs.[\\s\\S]+$/`.

