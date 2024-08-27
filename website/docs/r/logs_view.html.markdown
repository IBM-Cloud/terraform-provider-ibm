---
layout: "ibm"
page_title: "IBM : ibm_logs_view"
description: |-
  Manages logs_view.
subcategory: "Cloud Logs"
---


# ibm_logs_view

Create, update, and delete logs_views with this resource.

## Example Usage

```hcl
resource "ibm_logs_view" "logs_view_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-view"
  filters {
    filters {
      name = "applicationName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "subsystemName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "operationName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "serviceName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "severity"
      selected_values = {
        demo = true
      }
    }
  }
  search_query {
    query = "logs"
  }
  time_selection {
    custom_selection {
      from_time = "2024-01-25T11:31:43.152Z"
      to_time   = "2024-01-25T11:37:13.238Z"
    }
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `filters` - (Optional, List) View selected filters.
Nested schema for **filters**:
	* `filters` - (Optional, List) Selected filters.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **filters**:
		* `name` - (Required, String) Filter name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `selected_values` - (Required, Map) Filter selected values.
* `folder_id` - (Optional, String) View folder ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `name` - (Required, String) View name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
* `search_query` - (Optional, List) View search query.
Nested schema for **search_query**:
	* `query` - (Required, String) View search query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
* `time_selection` - (Required, List) View time selection.
Nested schema for **time_selection**:
	* `custom_selection` - (Optional, List) Custom time selection.
	Nested schema for **custom_selection**:
		* `from_time` - (Required, String) Custom time selection start timestamp.
		* `to_time` - (Required, String) Custom time selection end timestamp.
	* `quick_selection` - (Optional, List) Quick time selection.
	Nested schema for **quick_selection**:
		* `caption` - (Required, String) Quick time selection caption.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `seconds` - (Required, Integer) Quick time selection amount of seconds.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_view resource.
* `view_id` - The unique identifier of the logs_view.


## Import

You can import the `ibm_logs_view` resource by using `id`. `id` combination of `region`, `instance_id` and `view_id`.

# Syntax
<pre>
$ terraform import ibm_logs_view.logs_view < region >/< instance_id >/< view_id >;
</pre>

# Example
```
$ terraform import ibm_logs_view.logs_view eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/52
```
