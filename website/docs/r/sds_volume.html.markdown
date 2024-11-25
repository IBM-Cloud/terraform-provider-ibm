---
layout: "ibm"
page_title: "IBM : ibm_sds_volume"
description: |-
  Manages sds_volume.
subcategory: "sdsaas"
---

# ibm_sds_volume

Create, update, and delete sds_volumes with this resource.

## Example Usage

```hcl
resource "ibm_sds_volume" "sds_volume_instance" {
  capacity = 10
  hostnqnstring = "nqn.2024-07.org:1234"
  name = "my-volume"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `capacity` - (Required, Integer) The capacity of the volume (in gigabytes).
* `hostnqnstring` - (Optional, String) The host nqn.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
* `name` - (Optional, String) The name of the volume.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sds_volume.
* `created_at` - (String) The date and time that the volume was created.
* `host_mappings` - (List) List of host details that volume is mapped to.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
Nested schema for **host_mappings**:
	* `host_id` - (String) Unique identifer of the host.
	* `host_name` - (String) Unique name of the host.
	* `host_nqn` - (String) The NQN of the host configured in customer's environment.
* `resource_type` - (String) The resource type of the volume.
* `status` - (String) The current status of the volume.
* `status_reasons` - (List) Reasons for the current status of the volume.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.


## Import

You can import the `ibm_sds_volume` resource by using `id`. The volume profile id.

# Syntax
<pre>
$ terraform import ibm_sds_volume.sds_volume &lt;id&gt;
</pre>
