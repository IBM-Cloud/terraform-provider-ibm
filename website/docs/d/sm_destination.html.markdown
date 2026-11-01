---
layout: "ibm"
page_title: "IBM : ibm_sm_destination"
description: |-
  Get information about sm_destination
subcategory: "IBM Cloud Secrets Manager Instance Management API"
---

# ibm_sm_destination

Provides a read-only data source to retrieve information about a sm_destination. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_destination" "sm_destination" {
	destination_id = ibm_sm_destination.sm_destination_instance.destination_id
	instance_id = ibm_sm_destination.sm_destination_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_destination.
* `created_at` - (String) Timestamp when the destination was created.
* `created_by` - (String) IAM identity that created the destination.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\S]+$/`.
* `crn` - (Forces new resource, String) IBM Cloud Database service instance CRN.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `description` - (String) Optional description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `href` - (String) The URL of the destination resource.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character.
* `message` - (String) Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character.
* `name` - (String) Destination name.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `state` - (String) Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress.
  * Constraints: Allowable values are: `not_started`, `provisioning`, `succeeded`, `failed`, `deleting`.
* `type` - (String) Destination type.
  * Constraints: Allowable values are: `ibm_cloud_database`.
* `updated_at` - (String) Timestamp when the destination was last updated.

