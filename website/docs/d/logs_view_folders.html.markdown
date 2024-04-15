---
layout: "ibm"
page_title: "IBM : ibm_logs_view_folders"
description: |-
  Get information about logs_view_folders
subcategory: "Cloud Logs"
---

# ibm_logs_view_folders

Provides a read-only data source to retrieve information about logs_view_folders. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_view_folders" "logs_view_folders" {
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_view_folders.
* `view_folders` - (List) List of view folders.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **view_folders**:
	* `id` - (String) Folder id.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (String) Folder name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

