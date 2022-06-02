---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_definition"
description: |-
  Get information about tekton_pipeline_definition
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline_definition

Provides a read-only data source for tekton_pipeline_definition. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_definition" "tekton_pipeline_definition" {
	definition_id = ibm_cd_tekton_pipeline_definition.tekton_pipeline_definition.definition_id
	pipeline_id = ibm_cd_tekton_pipeline_definition.tekton_pipeline_definition.pipeline_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `definition_id` - (Required, Forces new resource, String) The definition ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the tekton_pipeline_definition.
* `scm_source` - (Required, List) Scm source for tekton pipeline defintion.
Nested scheme for **scm_source**:
	* `branch` - (Optional, String) A branch of the repo, branch field doesn't coexist with tag field.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `path` - (Required, String) The path to the definitions yaml files.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `tag` - (Optional, String) A tag of the repo.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,235}$/`.
	* `url` - (Required, String) General href URL.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `service_instance_id` - (Required, String) UUID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

