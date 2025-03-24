---
layout: "ibm"
page_title: "IBM : ibm_onboarding_registration"
description: |-
  Manages onboarding_registration.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_registration

Create, update, and delete onboarding_registrations with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_registration" "onboarding_registration_instance" {
  account_id = "4a5c3c51b97a446fbb1d0e1ef089823b"
  company_name = "Beautiful Company"
  primary_contact {
		name = "name"
		email = "email"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, String) The ID of your account.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
* `company_name` - (Required, String) The name of your company that is displayed in the IBM Cloud catalog.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character.
* `default_private_catalog_id` - (Optional, String) The default private catalog in which products are created.
  * Constraints: The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `primary_contact` - (Required, List) The primary contact for your product.
Nested schema for **primary_contact**:
	* `email` - (Required, String) The email address of the primary contact for your product.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+@\\S+$/`.
	* `name` - (Required, String) The name of the primary contact for your product.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character.
* `provider_access_group` - (Optional, String) The onboarding access group for your team.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^AccessGroupId-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_registration.
* `created_at` - (String) The time when the registration was created.
* `updated_at` - (String) The time when the registration was updated.


## Import

You can import the `ibm_onboarding_registration` resource by using `id`. The ID of your registration, which is the same as your billing and metering (BSS) account ID.

# Syntax
<pre>
$ terraform import ibm_onboarding_registration.onboarding_registration &lt;id&gt;
</pre>
