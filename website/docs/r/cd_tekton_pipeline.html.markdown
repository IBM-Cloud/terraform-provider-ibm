---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline"
description: |-
  Manages cd_tekton_pipeline.
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline

Provides a resource for cd_tekton_pipeline. This allows cd_tekton_pipeline to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
  worker {
		id = "id"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `enable_partial_cloning` - (Optional, Boolean) Flag whether to enable partial cloning for this pipeline. When partial clone is enabled, only the files contained within the paths specified in definition repositories will be read and cloned. This means symbolic links may not work.
  * Constraints: The default value is `false`.
* `enable_slack_notifications` - (Optional, Boolean) Flag whether to enable slack notifications for this pipeline. When enabled, pipeline run events will be published on all slack integration specified channels in the enclosing toolchain.
  * Constraints: The default value is `false`.
* `worker` - (Optional, List) Worker object containing worker ID only. If omitted the IBM Managed shared workers are used by default.
Nested scheme for **worker**:
	* `id` - (Required, String) ID of the worker.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^(public)|(preview)|([-0-9a-fA-F]{36})$/`.

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
	* `scm_source` - (List) SCM source for Tekton pipeline definition.
	Nested scheme for **scm_source**:
		* `branch` - (String) A branch from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `path` - (String) The path to the definition's yaml files.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `service_instance_id` - (String) ID of the SCM repository service instance.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `tag` - (String) A tag from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,235}$/`.
		* `url` - (Forces new resource, String) URL of the definition repository.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `enabled` - (Boolean) Flag whether this pipeline is enabled.
* `name` - (String) String.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.
* `properties` - (List) Tekton pipeline's environment properties.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested scheme for **properties**:
	* `enum` - (List) Options for `single_select` property type. Only needed when using `single_select` property type.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`. The maximum length is `128` items. The minimum length is `0` items.
	* `name` - (Forces new resource, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
	* `path` - (String) A dot notation path for `integration` type properties to select a value from the tool integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `type` - (String) Property type.
	  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
	* `value` - (String) Property value.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
* `resource_group_id` - (String) ID.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]+$/`.
* `runs_url` - (String) URL for this pipeline showing the list of pipeline runs.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `status` - (String) Pipeline status.
  * Constraints: Allowable values are: `configured`, `configuring`.
* `toolchain` - (List) Toolchain object.
Nested scheme for **toolchain**:
	* `crn` - (String) The CRN for the toolchain that contains the Tekton pipeline.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `id` - (String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `triggers` - (List) Tekton pipeline triggers list.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested scheme for **triggers**:
	* `cron` - (String) Only needed for timer triggers. Cron expression for timer trigger. Maximum frequency is every 5 minutes.
	  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\\*|([0-9]|1[0-9]|2[0-3])|\\*\/([0-9]|1[0-9]|2[0-3])) (\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\\*|([1-9]|1[0-2])|\\*\/([1-9]|1[0-2])) (\\*|([0-6])|\\*\/([0-6]))$/`.
	* `disabled` - (Boolean) Flag whether the trigger is disabled. If omitted the trigger is enabled by default.
	* `event_listener` - (String) Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `events` - (List) Only needed for Git triggers. Events object defines the events to which this Git trigger listens.
	Nested scheme for **events**:
		* `pull_request` - (Boolean) If true, the trigger listens for 'open pull request' or 'update pull request' Git webhook events.
		* `pull_request_closed` - (Boolean) If true, the trigger listens for 'close pull request' Git webhook events.
		* `push` - (Boolean) If true, the trigger listens for 'push' Git webhook events.
	* `href` - (String) API URL for interacting with the trigger.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `max_concurrent_runs` - (Integer) Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit.
	* `name` - (String) Trigger name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.
	* `properties` - (List) Trigger properties.
	  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
	Nested scheme for **properties**:
		* `enum` - (List) Options for `single_select` property type. Only needed for `single_select` property type.
		  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`. The maximum length is `128` items. The minimum length is `0` items.
		* `href` - (String) API URL for interacting with the trigger property.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `name` - (Forces new resource, String) Property name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
		* `path` - (String) A dot notation path for `integration` type properties to select a value from the tool integration. If left blank the full tool integration data will be used.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
		* `type` - (String) Property type.
		  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
		* `value` - (String) Property value. Can be empty and should be omitted for `single_select` property type.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `scm_source` - (List) SCM source repository for a Git trigger. Only needed for Git triggers.
	Nested scheme for **scm_source**:
		* `blind_connection` - (Boolean) True if the repository server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
		* `branch` - (String) Name of a branch from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `hook_id` - (String) ID of the webhook from the repo. Computed upon creation of the trigger.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `pattern` - (String) Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^.{1,235}$/`.
		* `service_instance_id` - (String) ID of the repository service instance.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Forces new resource, String) URL of the repository to which the trigger is listening.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `secret` - (List) Only needed for generic webhook trigger type. Secret used to start generic webhook trigger.
	Nested scheme for **secret**:
		* `algorithm` - (String) Algorithm used for `digest_matches` secret type. Only needed for `digest_matches` secret type.
		  * Constraints: Allowable values are: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160`.
		* `key_name` - (String) Secret name, not needed if type is `internal_validation`.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `source` - (String) Secret location, not needed if secret type is `internal_validation`.
		  * Constraints: Allowable values are: `header`, `payload`, `query`.
		* `type` - (String) Secret type.
		  * Constraints: Allowable values are: `token_matches`, `digest_matches`, `internal_validation`.
		* `value` - (String) Secret value, not needed if secret type is `internal_validation`.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `tags` - (List) Trigger tags array.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`. The maximum length is `128` items. The minimum length is `0` items.
	* `timezone` - (String) Only needed for timer triggers. Timezone for timer trigger.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_., \/]{1,234}$/`.
	* `type` - (String) Trigger type.
	  * Constraints: Allowable values are: .
	* `webhook_url` - (String) Webhook URL that can be used to trigger pipeline runs.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `worker` - (List) Worker used to run the trigger. If not specified the trigger will use the default pipeline worker.
	Nested scheme for **worker**:
		* `id` - (Forces new resource, String) ID of the worker.
		  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^(public)|(preview)|([-0-9a-fA-F]{36})$/`.
		* `name` - (String) Name of the worker. Computed based on the worker ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,235}$/`.
		* `type` - (String) Type of the worker. Computed based on the worker ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
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
