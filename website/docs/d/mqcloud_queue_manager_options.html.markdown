---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_queue_manager_options"
description: |-
  Get information about mqcloud_queue_manager_options
subcategory: "MQaaS"
---

# ibm_mqcloud_queue_manager_options

Provides a read-only data source to retrieve information about mqcloud_queue_manager_options. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_queue_manager_options" "mqcloud_queue_manager_options" {
	service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_queue_manager_options.
* `latest_version` - (String) The latest Queue manager version.
  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9]+.[0-9]+.[0-9]+_[0-9]+$/`.
* `locations` - (List) List of deployment locations.
  * Constraints: The list items must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`. The maximum length is `20` items. The minimum length is `1` item.
* `sizes` - (List) List of queue manager sizes.
  * Constraints: Allowable list items are: `xsmall`, `small`, `medium`, `large`. The maximum length is `20` items. The minimum length is `1` item.
* `versions` - (List) List of queue manager versions.
  * Constraints: The maximum length is `12` items. The minimum length is `1` item.

