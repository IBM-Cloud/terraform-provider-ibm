---
layout: "ibm"
page_title: "IBM : ibm_onboarding_product"
description: |-
  Manages onboarding_product.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_product

Create, update, and delete onboarding_products with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_product" "onboarding_product_instance" {
  primary_contact {
		name = "name"
		email = "email"
  }
  support {
		escalation_contacts {
			name = "name"
			email = "email"
			role = "role"
		}
  }
  type = "service"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `eccn_number` - (Optional, String) The Export Control Classification Number of your product.
* `ero_class` - (Optional, String) The ERO class of your product.
* `primary_contact` - (Required, List) The primary contact for your product.
Nested schema for **primary_contact**:
	* `email` - (Required, String) The email address of the primary contact for your product.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+@\\S+$/`.
	* `name` - (Required, String) The name of the primary contact for your product.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character.
* `support` - (Optional, List) The support information that is not displayed in the catalog, but available in ServiceNow.
Nested schema for **support**:
	* `escalation_contacts` - (Optional, List) The list of contacts in case of support escalations.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **escalation_contacts**:
		* `email` - (Optional, String) The email address of the support escalation contact.
		* `name` - (Optional, String) The name of the support escalation contact.
		* `role` - (Optional, String) The role of the support escalation contact.
* `tax_assessment` - (Optional, String) The tax assessment type of your product.
* `type` - (Required, String) The type of the product.
  * Constraints: Allowable values are: `software`, `service`, `professional_service`.
* `unspsc` - (Optional, Float) The United Nations Standard Products and Services Code of your product.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_product.
* `account_id` - (String) The IBM Cloud account ID of the provider.
* `approver_resource_id` - (String) The ID of the approval workflow of your product.
* `global_catalog_offering_id` - (String) The ID of a global catalog object.
  * Constraints: The value must match regular expression `/^\\S*$/`.
* `iam_registration_id` - (String) IAM registration identifier.
  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/^\\S*$/`.
* `private_catalog_id` - (String) The ID of the private catalog that contains the product. Only applicable for software type products.
  * Constraints: The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `private_catalog_offering_id` - (String) The ID of the linked private catalog product. Only applicable for software type products.
  * Constraints: The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `staging_global_catalog_offering_id` - (String) The ID of a global catalog object.
  * Constraints: The value must match regular expression `/^\\S*$/`.


## Import

You can import the `ibm_onboarding_product` resource by using `id`. The ID of a product in Partner Center - Sell.

# Syntax
<pre>
$ terraform import ibm_onboarding_product.onboarding_product &lt;id&gt;
</pre>
