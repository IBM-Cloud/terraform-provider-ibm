---
layout: "ibm"
page_title: "IBM : ibm_logs_view"
description: |-
  Get information about logs_view
subcategory: "Cloud Logs"
---


# ibm_logs_view

Provides a read-only data source to retrieve information about a logs_view. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_view" "logs_view_instance" {
	instance_id = ibm_logs_view.logs_view_instance.instance_id
	region      = ibm_logs_view.logs_view_instance.region
	logs_view_id = ibm_logs_view.logs_view_instance.view_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_view_id` - (Required, Forces new resource, Integer) View id.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_view.
* `filters` - (List) View selected filters.
Nested schema for **filters**:
	* `filters` - (List) Selected filters.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **filters**:
		* `name` - (String) Filter name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `selected_values` - (Map) Filter selected values.

* `folder_id` - (String) View folder ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

* `name` - (String) View name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `search_query` - (List) View search query.
Nested schema for **search_query**:
	* `query` - (String) View search query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `time_selection` - (List) View time selection.
Nested schema for **time_selection**:
	* `custom_selection` - (List) Custom time selection.
	Nested schema for **custom_selection**:
		* `from_time` - (String) Custom time selection start timestamp.
		* `to_time` - (String) Custom time selection end timestamp.
	* `quick_selection` - (List) Quick time selection.
	Nested schema for **quick_selection**:
		* `caption` - (String) Quick time selection caption.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `seconds` - (Integer) Quick time selection amount of seconds.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.

