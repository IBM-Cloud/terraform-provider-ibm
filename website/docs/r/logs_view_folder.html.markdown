---
layout: "ibm"
page_title: "IBM : ibm_logs_view_folder"
description: |-
  Manages logs_view_folder.
subcategory: "Cloud Logs"
---


# ibm_logs_view_folder

Create, update, and delete logs_view_folders with this resource.

## Example Usage

```hcl
resource "ibm_logs_view_folder" "logs_view_folder_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-view-folder"
}

```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `name` - (Required, String) Folder name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_view_folder resource.
* `view_folder_id` - The unique identifier of the logs_view_folder.


## Import

You can import the `ibm_logs_view_folder` resource by using `id`. `id` combination of `region`, `instance_id` and `view_folder_id`.

# Syntax
<pre>
$ terraform import ibm_logs_view_folder.logs_view_folder < region >/< instance_id >/< view_folder_id >;
</pre>

# Example
```
$ terraform import ibm_logs_view_folder.logs_view_folder eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/3dc02998-0b50-4ea8-b68a-4779d716fa1f
```
