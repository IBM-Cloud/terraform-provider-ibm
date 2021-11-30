---
layout: "ibm"
page_title: "IBM : ibm_cbr_rule"
description: |-
  Manages cbr_rule.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_rule

Provides a resource for cbr_rule. This allows cbr_rule to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cbr_rule" "cbr_rule" {
  description = "this is an example of rule"
  contexts {
    attributes {
      name = "networkZoneId"
      value = "322af80e125f6842cded8ba7a1008370"
    }
  }
  resources {
    attributes {
      name = "serviceName"
      value = "user-management"
    }
  }
}

```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `contexts` - (Required, List) The contexts this rule applies to.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` item.
Nested scheme for **contexts**:
    * `attributes` - (Required, List) The attributes.
      * Constraints: The minimum length is `1` item.
    Nested scheme for **attributes**:
        * `name` - (Required, String) The attribute name.
          * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9]+$`.
        * `value` - (Required, String) The attribute value.
          * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `^[\S\s]+$`.
* `description` - (Optional, String) The description of the rule.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `^[\x20-\xFE]*$`.
* `resources` - (Required, List) The resources this rule apply to.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested scheme for **resources**:
    * `attributes` - (Required, List) The resource attributes.
      * Constraints: The minimum length is `1` item.
    Nested scheme for **attributes**:
        * `name` - (Required, String) The attribute name.
          * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9]+$`.
        * `operator` - (Optional, String) The attribute operator.
          * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9]+$`.
        * `value` - (Required, String) The attribute value.
          * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `^[\S\s]+$`.
    * `tags` - (Optional, List) The optional resource tags.
      * Constraints: The maximum length is `10` items. The minimum length is `1` item.
    Nested scheme for **tags**:
        * `name` - (Required, String) The tag attribute name.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9 _.-]+$`.
        * `operator` - (Optional, String) The attribute operator.
          * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9]+$`.
        * `value` - (Required, String) The tag attribute value.
          * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `^[a-zA-Z0-9 _*?.-]+$`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cbr_rule.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The rule CRN.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `version` - Version of the cbr_rule.

## Import

You can import the `ibm_cbr_rule` resource by using `id`. The globally unique ID of the rule.

# Syntax
```
$ terraform import ibm_cbr_rule.cbr_rule <id>
```
