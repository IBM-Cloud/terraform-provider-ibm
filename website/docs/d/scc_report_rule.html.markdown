---
layout: "ibm"
page_title: "IBM : ibm_scc_report_rule"
description: |-
  Get information about scc_report_rule
subcategory: "Results"
---

# ibm_scc_report_rule

Provides a read-only data source to retrieve information about a scc_report_rule. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_report_rule" "scc_report_rule" {
	report_id = "report_id"
	rule_id = "rule-8d444f8c-fd1d-48de-bcaa-f43732568761"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `rule_id` - (Required, Forces new resource, String) The ID of a rule in a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_rule.
* `account_id` - (String) The rule account ID.

* `created_by` - (String) The ID of the user who created the rule.

* `created_on` - (String) The date when the rule was created.

* `description` - (String) The rule description.

* `id` - (String) The rule ID.

* `labels` - (List) The rule labels.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

* `type` - (String) The rule type.

* `updated_by` - (String) The ID of the user who updated the rule.

* `updated_on` - (String) The date when the rule was updated.

* `version` - (String) The rule version.

