---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_trigger"
description: |-
  Get information about cd_tekton_pipeline_trigger
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_trigger

Provides a read-only data source to retrieve information about a cd_tekton_pipeline_trigger. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
	pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.pipeline_id
	trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_trigger.
* `cron` - (String) Only needed for timer triggers. CRON expression that indicates when this trigger will activate. Maximum frequency is every 5 minutes. The string is based on UNIX crontab syntax: minute, hour, day of month, month, day of week. Example: The CRON expression 0 *_/2 * * * - translates to - every 2 hours.
  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^[-0-9a-zA-Z,\\*\/ ]{5,253}$/`.
* `enable_events_from_forks` - (Boolean) When enabled, pull request events from forks of the selected repository will trigger a pipeline run.
  * Constraints: The default value is `false`.
* `enabled` - (Boolean) Flag to check if the trigger is enabled.
  * Constraints: The default value is `true`.
* `event_listener` - (String) Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `events` - (List) Either 'events' or 'filter' is required specifically for Git triggers. Stores a list of events that a Git trigger listens to. Choose one or more from 'push', 'pull_request', and 'pull_request_closed'. If SCM repositories use the 'merge request' term, they correspond to the generic term i.e. 'pull request'.
  * Constraints: Allowable list items are: `push`, `pull_request`, `pull_request_closed`. The maximum length is `3` items. The minimum length is `0` items.
* `favorite` - (Boolean) Mark the trigger as a favorite.
  * Constraints: The default value is `false`.
* `filter` - (String) Either 'events' or 'filter' can be used. Stores the CEL (Common Expression Language) expression value which is used for event filtering against the Git webhook payloads.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `href` - (String) API URL for interacting with the trigger. Only included when fetching the list of pipeline triggers.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `max_concurrent_runs` - (Integer) Defines the maximum number of concurrent runs for this trigger. If omitted then the concurrency limit is disabled for this trigger.
* `name` - (String) Trigger name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^([a-zA-Z0-9]{1,2}|[a-zA-Z0-9][0-9a-zA-Z-_.: \/\\(\\)\\[\\]]{1,251}[a-zA-Z0-9])$/`.
* `properties` - (List) Optional trigger properties are used to override or supplement the pipeline properties when triggering a pipeline run.
  * Constraints: The maximum length is `1024` items. The minimum length is `0` items.
Nested schema for **properties**:
	* `enum` - (List) Options for `single_select` property type. Only needed for `single_select` property type.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
	* `href` - (String) API URL for interacting with the trigger property.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `locked` - (Boolean) When true, this property cannot be overridden at runtime. Attempting to override it will result in run requests being rejected. The default is false.
	* `name` - (Forces new resource, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
	* `path` - (String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
	* `type` - (Forces new resource, String) Property type.
	  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
	* `value` - (String) Property value. Any string value is valid.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `secret` - (List) Only needed for Generic Webhook trigger type. The secret is used to start the Generic Webhook trigger.
Nested schema for **secret**:
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
Nested schema for **source**:
	* `properties` - (List) Properties of the source, which define the URL of the repository and a branch or pattern.
	Nested schema for **properties**:
		* `blind_connection` - (Boolean) True if the repository server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
		* `branch` - (String) Name of a branch from the repo. Only one of branch, pattern, or filter should be specified.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `hook_id` - (String) Repository webhook ID. It is generated upon trigger creation.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `pattern` - (String) The pattern of Git branch or tag. You can specify a glob pattern such as '!test' or '*master' to match against multiple tags or branches in the repository.The glob pattern used must conform to Bash 4.3 specifications, see bash documentation for more info: https://www.gnu.org/software/bash/manual/bash.html#Pattern-Matching. Only one of branch, pattern, or filter should be specified.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.:@=$&^\/\\?\\!\\*\\+\\[\\]\\(\\)\\{\\}\\|\\\\]*$/`.
		* `tool` - (List) Reference to the repository tool in the parent toolchain.
		Nested schema for **tool**:
			* `id` - (String) ID of the repository tool instance in the parent toolchain.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Forces new resource, String) URL of the repository to which the trigger is listening.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `type` - (String) The only supported source type is "git", indicating that the source is a git repository.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.
* `tags` - (List) Optional trigger tags array.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `128` items. The minimum length is `0` items.
* `timezone` - (String) Only used for timer triggers. Specify the timezone used for this timer trigger, which will ensure the CRON activates this trigger relative to the specified timezone. If no timezone is specified, the default timezone used is UTC. Valid timezones are those listed in the IANA timezone database, https://www.iana.org/time-zones.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z+_., \/]{1,253}$/`.
* `type` - (String) Trigger type.
  * Constraints: Allowable values are: `manual`, `scm`, `timer`, `generic`.
* `webhook_url` - (String) Webhook URL that can be used to trigger pipeline runs.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `worker` - (List) Details of the worker used to run the trigger.
Nested schema for **worker**:
	* `id` - (String) ID of the worker.
	  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]{1,36}$/`.
	* `name` - (String) Name of the worker. Computed based on the worker ID.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,253}$/`.
	* `type` - (String) Type of the worker. Computed based on the worker ID.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.

