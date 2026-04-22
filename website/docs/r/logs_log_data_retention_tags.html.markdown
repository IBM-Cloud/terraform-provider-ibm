---
layout: "ibm"
page_title: "IBM : ibm_logs_log_data_retention_tags"
description: |-
  Manages logs_data_retention_tags.
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

# ibm_logs_log_data_retention_tags

Create, update, and delete log data retention tags with this resource.

Log data retention tags configuration manages the three editable archive retention tags used for log retention policies. The 'Default' tag is system-managed and cannot be modified through this resource.

~> **Important:** An archive bucket must be configured and attached to your Cloud Logs instance before you can configure data retention tags. The API will return an error if you attempt to configure retention tags without an archive bucket.

## Example Usage

```hcl
resource "ibm_logs_log_data_retention_tags" "logs_data_retention_tags_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  tags        = ["Short", "Medium", "Long"]
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `tags` - (Required, List) List of editable archive retention tags, excluding non-editable tags such as Default.
  * Constraints: The list must contain exactly `3` items. Each item must be between `1` and `256` characters. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_data_retention_tags resource, in the format `{region}/{instance_id}`.

## Import

You can import the `ibm_logs_log_data_retention_tags` resource by using `id`. The unique identifier of the logs_data_retention_tags resource.

# Syntax
<pre>
$ terraform import ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance <region>/<instance_id>
</pre>

# Example
```
$ terraform import ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f
```

## Notes

* This resource cannot be deleted. When you run `terraform destroy`, the resource is only removed from the Terraform state. The retention tags remain active in the Cloud Logs instance.
* Retention tags can only be deactivated by removing the archive bucket configuration from the Cloud Logs instance.
* The first successful creation (PUT request) of this resource will activate retention tags for the instance.
* An archive bucket must be configured before retention tags can be activated. If no archive bucket is configured, the resource creation will fail with a 412 Precondition Failed error.