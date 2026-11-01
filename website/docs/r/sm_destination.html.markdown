---
layout: "ibm"
page_title: "IBM : ibm_sm_destination"
description: |-
  Manages sm_destination.
subcategory: "IBM Cloud Secrets Manager Instance Management API"
---

# ibm_sm_destination

Create, update, and delete sm_destinations with this resource.

## Example Usage

```hcl
resource "ibm_sm_destination" "sm_destination_instance" {
  crn = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"
  description = "Production PostgreSQL database"
  instance_id = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
  name = "name"
  type = "ibm_cloud_database"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `crn` - (Required, Forces new resource, String) IBM Cloud Database service instance CRN.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `description` - (Optional, String) Optional description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `name` - (Required, String) Destination name.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$/`.
* `type` - (Required, String) Destination type.
  * Constraints: Allowable values are: `ibm_cloud_database`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_destination.
* `created_at` - (String) Timestamp when the destination was created.
* `created_by` - (String) IAM identity that created the destination.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\S]+$/`.
* `destination_id` - (String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `href` - (String) The URL of the destination resource.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character.
* `message` - (String) Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character.
* `state` - (String) Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress.
  * Constraints: Allowable values are: `not_started`, `provisioning`, `succeeded`, `failed`, `deleting`.
* `updated_at` - (String) Timestamp when the destination was last updated.


## Import

You can import the `ibm_sm_destination` resource by using `id`.
The `id` property can be formed from `instance_id`, and `destination_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;
</pre>
* `instance_id`: A string in the format `bfc50c2e-d66d-4f37-9ccf-9713f8325b39`. Secrets Manager instance ID.
* `destination_id`: A strfmt.UUID. Destination ID.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_destination.sm_destination &lt;instance_id&gt;/&lt;destination_id&gt;
</pre>
