---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline"
description: |-
  Manages cd_tekton_pipeline.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline

Provides a resource for cd_tekton_pipeline. This allows cd_tekton_pipeline to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  worker {
		id = "public"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `pipeline_id` - (Required, String) ID of current instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `enable_notifications` - (Optional, Boolean) Flag whether to enable notifications for this pipeline. When enabled, pipeline run events are published on all slack integration specified channels in the parent toolchain.
  * Constraints: The default value is `false`.
* `enable_partial_cloning` - (Optional, Boolean) Flag whether to enable partial cloning for this pipeline. When partial clone is enabled, only the files contained within the paths specified in definition repositories are read and cloned, this means that symbolic links might not work.
  * Constraints: The default value is `false`.
* `worker` - (Optional, List) Worker object containing worker ID only. If omitted the IBM Managed shared workers are used by default.
Nested scheme for **worker**:
	* `id` - (Required, String) ID of the worker.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]{1,253}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_tekton_pipeline.
* `build_number` - (Integer) The latest pipeline run build number. If this property is absent, the pipeline hasn't had any pipeline runs.
  * Constraints: The minimum value is `1`.
* `created_at` - (String) Standard RFC 3339 Date Time String.
* `definitions` - (List) Definition list.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested scheme for **definitions**:
	* `id` - (String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `source` - (List) Source repository containing the Tekton pipeline definition.
	Nested scheme for **source**:
		* `properties` - (List) Properties of the source, which define the URL of the repository and a branch or tag.
		Nested scheme for **properties**:
			* `branch` - (String) A branch from the repo, specify one of branch or tag only.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
			* `path` - (String) The path to the definition's YAML files.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
			* `tag` - (String) A tag from the repo, specify one of branch or tag only.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,253}$/`.
			* `tool` - (List) Reference to the repository tool, in the parent toolchain, that contains the pipeline definition.
			Nested scheme for **tool**:
				* `id` - (String) ID of the repository tool instance in the parent toolchain.
				  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
			* `url` - (Forces new resource, String) URL of the definition repository.
			  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `type` - (String) The only supported source type is "git", indicating that the source is a git repository.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.
* `enabled` - (Boolean) Flag whether this pipeline is enabled.
  * Constraints: The default value is `true`.
* `name` - (String) String.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,253}[a-zA-Z0-9]$/`.
* `properties` - (List) Tekton pipeline's environment properties.
  * Constraints: The maximum length is `1024` items. The minimum length is `0` items.
Nested scheme for **properties**:
	* `enum` - (List) Options for `single_select` property type. Only needed when using `single_select` property type.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
	* `name` - (Forces new resource, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
	* `path` - (String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
	* `type` - (String) Property type.
	  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
	* `value` - (String) Property value. Any string value is valid.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `resource_group` - (List) The ID of the resource group in which the pipeline was created.
Nested scheme for **resource_group**:
	* `id` - (String) ID.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]+$/`.
* `runs_url` - (String) URL for this pipeline showing the list of pipeline runs.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `status` - (String) Pipeline status.
  * Constraints: Allowable values are: `configured`, `configuring`.
* `toolchain` - (List) Toolchain object containing references to the parent toolchain.
Nested scheme for **toolchain**:
	* `crn` - (String) The CRN for the toolchain that contains the Tekton pipeline.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `id` - (String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `triggers` - (List) Tekton pipeline triggers list.
  * Constraints: The maximum length is `1024` items. The minimum length is `0` items.
Nested scheme for **triggers**:
	* `cron` - (String) Only needed for timer triggers. Cron expression that indicates when this trigger will activate. Maximum frequency is every 5 minutes. The string is based on UNIX crontab syntax: minute, hour, day of month, month, day of week. Example: 0 *_/2 * * * - every 2 hours.
	  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\\*|([0-9]|1[0-9]|2[0-3])|\\*\/([0-9]|1[0-9]|2[0-3])) (\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\\*|([1-9]|1[0-2])|\\*\/([1-9]|1[0-2])) (\\*|([0-6])|\\*\/([0-6]))$/`.
	* `enabled` - (Boolean) Flag whether the trigger is enabled.
	  * Constraints: The default value is `true`.
	* `event_listener` - (String) Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
	* `events` - (List) Only needed for Git triggers. List of events to which a Git trigger listens. Choose one or more from: 'push', 'pull_request' and 'pull_request_closed'. For SCM repositories that use 'merge request' events, such events map to the equivalent 'pull request' events.
	  * Constraints: Allowable list items are: `push`, `pull_request`, `pull_request_closed`. The list items must match regular expression `/^[-0-9a-zA-Z_,]+$/`. The maximum length is `3` items. The minimum length is `0` items.
	* `href` - (String) API URL for interacting with the trigger. Only included when fetching the list of pipeline triggers.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `max_concurrent_runs` - (Integer) Defines the maximum number of concurrent runs for this trigger. If omitted then the concurrency limit is disabled for this trigger.
	* `name` - (String) Trigger name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,253}[a-zA-Z0-9]$/`.
	* `properties` - (List) Optional trigger properties used to override or supplement the pipeline properties when triggering a pipeline run.
	  * Constraints: The maximum length is `1024` items. The minimum length is `0` items.
	Nested scheme for **properties**:
		* `enum` - (List) Options for `single_select` property type. Only needed for `single_select` property type.
		  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
		* `href` - (String) API URL for interacting with the trigger property.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `name` - (Forces new resource, String) Property name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `path` - (String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
		* `type` - (String) Property type.
		  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
		* `value` - (String) Property value. Any string value is valid.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `secret` - (List) Only needed for generic webhook trigger type. Secret used to start generic webhook trigger.
	Nested scheme for **secret**:
		* `algorithm` - (String) Algorithm used for `digest_matches` secret type. Only needed for `digest_matches` secret type.
		  * Constraints: Allowable values are: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160`.
		* `key_name` - (String) Secret name, not needed if type is `internal_validation`.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `source` - (String) Secret location, not needed if secret type is `internal_validation`.
		  * Constraints: Allowable values are: `header`, `payload`, `query`.
		* `type` - (String) Secret type.
		  * Constraints: Allowable values are: `token_matches`, `digest_matches`, `internal_validation`.
		* `value` - (String) Secret value, not needed if secret type is `internal_validation`.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `source` - (List) Source repository for a Git trigger. Only required for Git triggers. The referenced repository URL must match the URL of a repository tool integration in the parent toolchain. Obtain the list of integrations from the toolchain API https://cloud.ibm.com/apidocs/toolchain#list-tools.
	Nested scheme for **source**:
		* `properties` - (List) Properties of the source, which define the URL of the repository and a branch or pattern.
		Nested scheme for **properties**:
			* `blind_connection` - (Boolean) True if the repository server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
			* `branch` - (String) Name of a branch from the repo. One of branch or pattern must be specified, but only one or the other.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
			* `hook_id` - (String) ID of the webhook from the repo. Computed upon creation of the trigger.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
			* `pattern` - (String) Git branch or tag pattern to listen to, specify one of branch or pattern only. When specifying a tag to listen to, you can also specify a simple glob pattern such as '!test' or '*master' to match against multiple tags/branches in the repository.
			  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.!*]*$/`.
			* `tool` - (List) Reference to the repository tool in the parent toolchain.
			Nested scheme for **tool**:
				* `id` - (String) ID of the repository tool instance in the parent toolchain.
				  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
			* `url` - (Forces new resource, String) URL of the repository to which the trigger is listening.
			  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `type` - (String) The only supported source type is "git", indicating that the source is a git repository.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.
	* `tags` - (List) Optional trigger tags array.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `128` items. The minimum length is `0` items.
	* `timezone` - (String) Only used for timer triggers. Specify the timezone used for this timer trigger, which will ensure the cron activates this trigger relative to the specified timezone. If no timezone is specified, the default timezone used is UTC. Valid timezones are those listed in the IANA timezone database, https://www.iana.org/time-zones.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z+_., \/]{1,253}$/`.
	* `type` - (String) Trigger type.
	  * Constraints: Allowable values are: `manual`, `scm`, `timer`, `generic`.
	* `webhook_url` - (String) Webhook URL that can be used to trigger pipeline runs.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `worker` - (List) Worker used to run the trigger. If not specified the trigger will use the default pipeline worker.
	Nested scheme for **worker**:
		* `id` - (Forces new resource, String) ID of the worker.
		  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]{1,36}$/`.
		* `name` - (String) Name of the worker. Computed based on the worker ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,253}$/`.
		* `type` - (String) Type of the worker. Computed based on the worker ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `updated_at` - (String) Standard RFC 3339 Date Time String.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_tekton_pipeline` resource by using `id`. UUID.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline.cd_tekton_pipeline <id>
```
