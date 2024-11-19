---
layout: "ibm"
page_title: "IBM : ibm_logs_dashboard_folders"
description: |-
  Get information about logs_dashboard_folders
subcategory: "Cloud Logs"
---


# ibm_logs_dashboard_folders

Provides a read-only data source to retrieve information about logs_dashboard_folders. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_dashboard_folders" "logs_dashboard_folders" {
	instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
    region      = ibm_logs_e2m.logs_e2m_instance.region
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_dashboard_folders.
* `folders` - (List) The list of folders.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **folders**:
	* `id` - (String) The dashboard folder ID, uuid.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `name` - (String) The dashboard folder name, required.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `parent_id` - (String) The dashboard folder parent ID, optional. If not set, the folder is a root folder, if set, the folder is a subfolder of the parent folder and needs to be a uuid.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

