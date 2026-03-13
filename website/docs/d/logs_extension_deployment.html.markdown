---
layout: "ibm"
page_title: "IBM : ibm_logs_extension_deployment"
description: |-
  Get information about Extension deployment
subcategory: "Cloud Logs"
---

# ibm_logs_extension_deployment

Provides a read-only data source to retrieve information about an Extension deployment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_extension_deployment" "logs_extension_deployment" {
  instance_id  = ibm_resource_instance.logs_instance.guid
  region       = ibm_resource_instance.logs_instance.location
  logs_extension_deployment_id = "logs_extension_deployment_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String) Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `logs_extension_deployment_id` - (Required, Forces new resource, String) The unique identifier of the extension.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Extension deployment.
* `applications` - (List) Applications that the Extension is deployed for. When this is empty, it is applied to all applications.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.

* `item_ids` - (List) The list of Extension item IDs to deploy.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `1` item.

* `subsystems` - (List) Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.

* `version` - (String) The version of the Extension revision to deploy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

