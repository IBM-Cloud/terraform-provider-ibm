---
layout: "ibm"
page_title: "IBM : ibm_logs_dashboard_folder"
description: |-
  Manages logs_dashboard_folder.
subcategory: "Cloud Logs"
---


# ibm_logs_dashboard_folder

Create, update, and delete logs_dashboard_folders with this resource.

## Example Usage

```hcl
resource "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "My Folder"
  parent_id   = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `name` - (Required, String) The dashboard folder name, required.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `parent_id` - (Optional, String) The dashboard folder parent ID, optional. If not set, the folder is a root folder, if set, the folder is a subfolder of the parent folder and needs to be a uuid.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_dashboard_folder resource.
* `dashboard_folder_id` - The unique identifier of the logs dashboard folder.


## Import

You can import the `ibm_logs_dashboard_folder` resource by using `id`. `id` combination of `region`, `instance_id` and `dashboard_folder_id`.

# Syntax
<pre>
$ terraform import ibm_logs_dashboard_folder.logs_dashboard_folder < region >/< instance_id >/< dashboard_folder_id >;
</pre>

# Example
```
$ terraform import ibm_logs_dashboard_folder.logs_dashboard_folder eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/d6a3658e-78d2-47d0-9b81-b2c551f01b09
```
