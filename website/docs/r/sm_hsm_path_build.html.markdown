---
layout: "ibm"
page_title: "IBM : ibm_sm_hsm_path_build"
description: |-
  Manages sm_hsm_path_build.
subcategory: "Secrets Manager"
---

# ibm_sm_hsm_path_build

Create, update, and delete sm_hsm_path_builds with this resource.

## Example Usage

```hcl
resource "ibm_sm_hsm_path_build" "sm_hsm_path_build_instance" {
  destination_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  instance_id = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
  path_id = "fa3d1234-0000-0000-0000-000000000001"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `destination_id` - (Required, Forces new resource, String) Destination ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `instance_id` - (Required, Forces new resource, String) Secrets Manager instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `path_id` - (Required, Forces new resource, String) Path ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sm_hsm_path_build.
* `path_id` - (String) ID of the path being provisioned.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `provisioning_state` - (String) Always `provisioning` on a 202 response.
  * Constraints: Allowable values are: `provisioning`.


## Import

You can import the `ibm_sm_hsm_path_build` resource by using `id`.
The `id` property can be formed from `instance_id`, `destination_id`, `path_id`, and `path_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;destination_id&gt;/&lt;path_id&gt;/&lt;path_id&gt;
</pre>
* `instance_id`: A string in the format `fe5c5c27-3e44-474a-9484-55892e12d6d1`. Secrets Manager instance ID.
* `destination_id`: A string in the format `b2c3d4e5-f6a7-8901-bcde-f12345678901`. Destination ID.
* `path_id`: A string in the format `fa3d1234-0000-0000-0000-000000000001`. Path ID.
* `path_id`: A string in the format `fa3d1234-0000-0000-0000-000000000001`. ID of the path being provisioned.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
<pre>
$ terraform import ibm_sm_hsm_path_build.sm_hsm_path_build &lt;instance_id&gt;/&lt;destination_id&gt;/&lt;path_id&gt;/&lt;path_id&gt;
</pre>
