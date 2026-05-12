---
layout: "ibm"
page_title: "IBM : ibm_code_engine_allowed_outbound_destination"
description: |-
  Get information about code_engine_allowed_outbound_destination
subcategory: "Code Engine"
---

# ibm_code_engine_allowed_outbound_destination

Provides a read-only data source to retrieve information about a code_engine_allowed_outbound_destination. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination" {
  project_id = data.ibm_code_engine_project.code_engine_project.project_id
  name       = "my-allowed-outbound-destination"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your allowed outbound destination.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_allowed_outbound_destination.

* `cidr_block` - (String) The IPv4 address range.
  * Constraints: The maximum length is `18` characters. The minimum length is `0` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.

* `entity_tag` - (String) The version of the allowed outbound destination, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `type` - (Forces new resource, String) Specify the type of the allowed outbound destination. Allowed types are: 'cidr_block'.
  * Constraints: The default value is `cidr_block`. Allowable values are: `cidr_block`. The value must match regular expression `/^(cidr_block)$/`.

