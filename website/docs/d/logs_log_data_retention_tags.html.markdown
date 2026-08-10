---
layout: "ibm"
page_title: "IBM : ibm_logs_log_data_retention_tags"
description: |-
  Get information about logs_log_data_retention_tags
subcategory: "Cloud Logs"
---

# ibm_logs_log_data_retention_tags

Provides a read-only data source to retrieve information about log data retention tags. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> **Important:** An archive bucket must be configured and attached to your Cloud Logs instance before retention tags can be retrieved. If no archive bucket is configured, the data source will fail.

## Example Usage

```hcl
data "ibm_logs_log_data_retention_tags" "logs_data_retention_tags_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String) Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_data_retention_tags resource, in the format `{region}/{instance_id}`.
* `tags` - (List) List of editable archive retention tags, excluding non-editable tags such as Default.
  * Constraints: The list contains exactly `3` items. Each item is between `1` and `256` characters. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.