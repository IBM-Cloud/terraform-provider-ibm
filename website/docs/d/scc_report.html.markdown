---
layout: "ibm"
page_title: "IBM : ibm_scc_report"
description: |-
  Get information about scc_report
subcategory: "Results"
---

# ibm_scc_report

Provides a read-only data source to retrieve information about a scc_report. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_report" "scc_report" {
	report_id = "report_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report.
* `account` - (List) The account that is associated with a report.
Nested schema for **account**:
	* `id` - (String) The account ID.
	* `name` - (String) The account name.
	* `type` - (String) The account type.

* `attachment` - (List) The attachment that is associated with a report.
Nested schema for **attachment**:
	* `description` - (String) The description of the attachment.
	* `id` - (String) The attachment ID.
	* `name` - (String) The name of the attachment.
	* `schedule` - (String) The attachment schedule.
	* `scope` - (List) The scope of the attachment.
	  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
	Nested schema for **scope**:
		* `environment` - (String) The environment that relates to this scope.
		* `id` - (String) The unique identifier for this scope.
		* `properties` - (List) The properties that are supported for scoping by this environment.
		  * Constraints: The maximum length is `99999` items. The minimum length is `0` items.
		Nested schema for **properties**:
			* `name` - (String) The property name.
			* `value` - (String) The property value.

* `cos_object` - (String) The Cloud Object Storage object that is associated with the report.

* `created_on` - (String) The date when the report was created.

* `group_id` - (String) The group ID that is associated with the report. The group ID combines profile, scope, and attachment IDs.

* `id` - (String) The ID of the report.

* `instance_id` - (String) Instance ID.

* `profile` - (List) The profile information.
Nested schema for **profile**:
	* `id` - (String) The profile ID.
	* `name` - (String) The profile name.
	* `version` - (String) The profile version.

* `scan_time` - (String) The date when the scan was run.

* `type` - (String) The type of the scan.

