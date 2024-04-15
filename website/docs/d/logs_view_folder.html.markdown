---
layout: "ibm"
page_title: "IBM : ibm_logs_view_folder"
description: |-
  Get information about logs_view_folder
subcategory: "Cloud Logs"
---

# ibm_logs_view_folder

Provides a read-only data source to retrieve information about a logs_view_folder. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_view_folder" "logs_view_folder" {
	logs_view_folder_id = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `logs_view_folder_id` - (Required, Forces new resource, String) Folder id.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_view_folder.
* `name` - (String) Folder name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

