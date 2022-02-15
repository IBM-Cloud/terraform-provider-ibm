---
layout: "ibm"
page_title: "IBM : ibm_tekton_pipeline_definition"
description: |-
  Manages tekton_pipeline_definition.
subcategory: "Continuous Delivery Pipeline"
---

# ibm_tekton_pipeline_definition

Provides a resource for tekton_pipeline_definition. This allows tekton_pipeline_definition to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_tekton_pipeline_definition" "tekton_pipeline_definition" {
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  scm_source = {"path":".tekton","url":"https://github.com/IBM/tekton-tutorial.git","branch":"master"}
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `scm_source` - (Optional, List) Scm source for tekton pipeline defintion.
Nested scheme for **scm_source**:
	* `branch` - (Required, String) The branch of the repo.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `path` - (Required, String) The path to the definitions yaml files.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `url` - (Required, String) General href URL.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the tekton_pipeline_definition.
* `definition_id` - (Required, String) UUID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `service_instance_id` - (Required, String) UUID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Import

You can import the `ibm_tekton_pipeline_definition` resource by using `id`.
The `id` property can be formed from `pipeline_id`, and `definition_id` in the following format:

```
<pipeline_id>/<definition_id>
```
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The tekton pipeline ID.
* `definition_id`: A string in the format `94299034-d45f-4e9a-8ed5-6bd5c7bb7ada`. The definition ID.

# Syntax
```
$ terraform import ibm_tekton_pipeline_definition.tekton_pipeline_definition <pipeline_id>/<definition_id>
```
