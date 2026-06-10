---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_definition"
description: |-
  Manages cd_tekton_pipeline_definition.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_definition

Create, update, and delete cd_tekton_pipeline_definitions with this resource.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  source {
		type = "git"
		properties {
			url = "url"
			branch = "branch"
			tag = "tag"
			path = "path"
			tool {
				id = "id"
			}
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `source` - (Required, List) Source repository containing the Tekton pipeline definition.
Nested schema for **source**:
	* `properties` - (Required, List) Properties of the source, which define the URL of the repository and a branch or tag.
	Nested schema for **properties**:
		* `branch` - (Optional, String) A branch from the repo, specify one of branch or tag only.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `path` - (Required, String) The path to the definition's YAML files.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `tag` - (Optional, String) A tag from the repo, specify one of branch or tag only.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,253}$/`.
		* `tool` - (Optional, List) Reference to the repository tool in the parent toolchain.
		Nested schema for **tool**:
			* `id` - (Computed, String) ID of the repository tool instance in the parent toolchain.
			  * Constraints: Length must be `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Required, Forces new resource, String) URL of the definition repository.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `type` - (Required, String) The only supported source type is "git", indicating that the source is a git repository.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_definition.
* `definition_id` - (String) The aggregated definition ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `href` - (String) API URL for interacting with the definition.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.


## Import

You can import the `ibm_cd_tekton_pipeline_definition` resource by using `id`.
The `id` property can be formed from `pipeline_id`, and `definition_id` in the following format:

<pre>
&lt;pipeline_id&gt;/&lt;definition_id&gt;
</pre>
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The Tekton pipeline ID.
* `definition_id`: A string. The aggregated definition ID.

# Syntax
<pre>
$ terraform import ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition &lt;pipeline_id&gt;/&lt;definition_id&gt;
</pre>
