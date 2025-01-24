---
layout: "ibm"
page_title: "IBM : ibm_code_engine_allowed_outbound_destination"
description: |-
  Manages code_engine_allowed_outbound_destination.
subcategory: "Code Engine"
---

# ibm_code_engine_allowed_outbound_destination

Create, update, and delete code_engine_allowed_outbound_destinations with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  type = "cidr_block"
  name = "my-cidr-block-1"
  cidr_block = "192.68.3.0/24"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cidr_block` - (Optional, String) The IPv4 address range.
  * Constraints: The maximum length is `18` characters. The minimum length is `0` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.
* `name` - (Optional, String) The name of the CIDR block.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `type` - (Required, Forces new resource, String) Specify the type of the allowed outbound destination. Allowed types are: 'cidr_block'.
  * Constraints: The default value is `cidr_block`. Allowable values are: `cidr_block`. The value must match regular expression `/^(cidr_block)$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_allowed_outbound_destination.

* `entity_tag` - (String) The version of the allowed outbound destination, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `etag` - ETag identifier for code_engine_allowed_outbound_destination.

## Import

You can import the `ibm_code_engine_allowed_outbound_destination` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string. The name of the CIDR block.

# Syntax
<pre>
$ terraform import ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination &lt;project_id&gt;/&lt;name&gt;
</pre>
