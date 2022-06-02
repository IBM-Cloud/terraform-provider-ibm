---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_workers"
description: |-
  Get information about tekton_pipeline_workers
subcategory: "Continuous Delivery Pipeline"
---

# ibm_cd_tekton_pipeline_workers

Provides a read-only data source for tekton_pipeline_workers. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_workers" "tekton_pipeline_workers" {
	pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the tekton_pipeline_workers.
* `workers` - (Required, List) Workers list.
Nested scheme for **workers**:
	* `id` - (Required, String) ID.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]+$/`.
	* `name` - (Optional, String) worker name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,235}$/`.
	* `type` - (Optional, String) worker type.
	  * Constraints: Allowable values are: `private`, `public`.

