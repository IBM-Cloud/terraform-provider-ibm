---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_definition"
description: |-
  Get information about cd_tekton_pipeline_definition
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_definition

Provides a read-only data source to retrieve information about a cd_tekton_pipeline_definition. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition" {
	definition_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance.definition_id
	pipeline_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance.pipeline_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `definition_id` - (Required, Forces new resource, String) The definition ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_definition.
* `href` - (String) API URL for interacting with the definition.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `source` - (List) Source repository containing the Tekton pipeline definition.
Nested schema for **source**:
	* `properties` - (List) Properties of the source, which define the URL of the repository and a branch or tag.
	Nested schema for **properties**:
		* `branch` - (String) A branch from the repo, specify one of branch or tag only.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `path` - (String) The path to the definition's YAML files.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `tag` - (String) A tag from the repo, specify one of branch or tag only.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,253}$/`.
		* `tool` - (List) Reference to the repository tool in the parent toolchain.
		Nested schema for **tool**:
			* `id` - (String) ID of the repository tool instance in the parent toolchain.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Forces new resource, String) URL of the definition repository.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `type` - (String) The only supported source type is "git", indicating that the source is a git repository.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.

