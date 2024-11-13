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
	data "ibm_logs_view_folders" "logs_view_folders_instance" {
		instance_id = ibm_logs_view_folder.logs_view_folder_instance.instance_id
		region      = ibm_logs_view_folder.logs_view_folder_instance.region
	}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_view_folders.
* `view_folders` - (List) List of view folders.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **view_folders**:
	* `id` - (String) Folder ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (String) Folder name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

