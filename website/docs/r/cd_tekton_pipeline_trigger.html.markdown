---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_trigger"
description: |-
  Manages cd_tekton_pipeline_trigger.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_trigger

Create, update, and delete cd_tekton_pipeline_triggers with this resource.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  event_listener = "pr-listener"
  max_concurrent_runs = 3
  name = "Manual Trigger"
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  type = "manual"
  worker {
    id = "public"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cron` - (Optional, String) Only needed for timer triggers. CRON expression that indicates when this trigger will activate. Maximum frequency is every 5 minutes. The string is based on UNIX crontab syntax: minute, hour, day of month, month, day of week. Example: The CRON expression 0 *_/2 * * * - translates to - every 2 hours.
  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^[-0-9a-zA-Z,\\*\/ ]{5,253}$/`.
* `disable_draft_events` - (Optional, Boolean) Prevent new pipeline runs from being triggered by events from draft pull requests.
  * Constraints: The default value is `false`.
* `enable_events_from_forks` - (Optional, Boolean) When enabled, pull request events from forks of the selected repository will trigger a pipeline run.
  * Constraints: The default value is `false`.
* `enabled` - (Optional, Boolean) Flag to check if the trigger is enabled.
  * Constraints: The default value is `true`.
* `event_listener` - (Required, String) Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `events` - (Optional, List) Either 'events' or 'filter' is required specifically for Git triggers. Stores a list of events that a Git trigger listens to. Choose one or more from 'push', 'pull_request', and 'pull_request_closed'. If SCM repositories use the 'merge request' term, they correspond to the generic term i.e. 'pull request'.
  * Constraints: Allowable list items are: `push`, `pull_request`, `pull_request_closed`. The maximum length is `3` items. The minimum length is `0` items.
* `favorite` - (Optional, Boolean) Mark the trigger as a favorite.
  * Constraints: The default value is `false`.
* `filter` - (Optional, String) Either 'events' or 'filter' can be used. Stores the CEL (Common Expression Language) expression value which is used for event filtering against the Git webhook payloads.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `limit_waiting_runs` - (Optional, Boolean) Flag that will limit the trigger to a maximum of one waiting run. A newly triggered run will cause any other waiting run(s) to be automatically cancelled.
  * Constraints: The default value is `false`.
* `max_concurrent_runs` - (Optional, Integer) Defines the maximum number of concurrent runs for this trigger. If omitted then the concurrency limit is disabled for this trigger.
* `name` - (Required, String) Trigger name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^([a-zA-Z0-9]{1,2}|[a-zA-Z0-9][0-9a-zA-Z-_.: \/\\(\\)\\[\\]]{1,251}[a-zA-Z0-9])$/`.
* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `secret` - (Optional, List) Only needed for Generic Webhook trigger type. The secret is used to start the Generic Webhook trigger.
Nested schema for **secret**:
	* `algorithm` - (Optional, String) Algorithm used for `digest_matches` secret type. Only needed for `digest_matches` secret type.
	  * Constraints: Allowable values are: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160`.
	* `key_name` - (Optional, String) Secret name, not needed if type is `internal_validation`.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
	* `source` - (Optional, String) Secret location, not needed if secret type is `internal_validation`.
	  * Constraints: Allowable values are: `header`, `payload`, `query`.
	* `type` - (Optional, String) Secret type.
	  * Constraints: Allowable values are: `token_matches`, `digest_matches`, `internal_validation`.
	* `value` - (Optional, String) Secret value, not needed if secret type is `internal_validation`.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `source` - (Optional, List) Source repository for a Git trigger. Only required for Git triggers. The referenced repository URL must match the URL of a repository tool integration in the parent toolchain. Obtain the list of integrations from the toolchain API https://cloud.ibm.com/apidocs/toolchain#list-tools.
Nested schema for **source**:
	* `properties` - (Required, List) Properties of the source, which define the URL of the repository and a branch or pattern.
	Nested schema for **properties**:
		* `blind_connection` - (Computed, Boolean) True if the repository server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
		* `branch` - (Optional, String) Name of a branch from the repo. Only one of branch, pattern, or filter should be specified.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `hook_id` - (Computed, String) Repository webhook ID. It is generated upon trigger creation.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
		* `pattern` - (Optional, String) The pattern of Git branch or tag. You can specify a glob pattern such as '!test' or '*master' to match against multiple tags or branches in the repository.The glob pattern used must conform to Bash 4.3 specifications, see bash documentation for more info: https://www.gnu.org/software/bash/manual/bash.html#Pattern-Matching. Only one of branch, pattern, or filter should be specified.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.:@=$&^\/\\?\\!\\*\\+\\[\\]\\(\\)\\{\\}\\|\\\\]*$/`.
		* `tool` - (Required, List) Reference to the repository tool in the parent toolchain.
		Nested schema for **tool**:
			* `id` - (Computed, String) ID of the repository tool instance in the parent toolchain.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Required, Forces new resource, String) URL of the repository to which the trigger is listening.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `type` - (Required, String) The only supported source type is "git", indicating that the source is a git repository.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^git$/`.
* `tags` - (Optional, List) Optional trigger tags array.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `128` items. The minimum length is `0` items.
* `timezone` - (Optional, String) Only used for timer triggers. Specify the timezone used for this timer trigger, which will ensure the CRON activates this trigger relative to the specified timezone. If no timezone is specified, the default timezone used is UTC. Valid timezones are those listed in the IANA timezone database, https://www.iana.org/time-zones.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z+_., \/]{1,253}$/`.
* `type` - (Required, String) Trigger type.
  * Constraints: Allowable values are: .
* `worker` - (Optional, List) Details of the worker used to run the trigger.
Nested schema for **worker**:
	* `id` - (Required, String) ID of the worker.
	  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]{1,36}$/`.
	* `name` - (Computed, String) Name of the worker. Computed based on the worker ID.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,253}$/`.
	* `type` - (Computed, String) Type of the worker. Computed based on the worker ID.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_trigger.
* `href` - (String) API URL for interacting with the trigger. Only included when fetching the list of pipeline triggers.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
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
* `trigger_id` - (String) The Trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `webhook_url` - (String) Webhook URL that can be used to trigger pipeline runs.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.


## Import

You can import the `ibm_cd_tekton_pipeline_trigger` resource by using `id`.
The `id` property can be formed from `pipeline_id`, and `trigger_id` in the following format:

<pre>
&lt;pipeline_id&gt;/&lt;trigger_id&gt;
</pre>
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The Tekton pipeline ID.
* `trigger_id`: A string. The Trigger ID.

# Syntax
<pre>
$ terraform import ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger &lt;pipeline_id&gt;/&lt;trigger_id&gt;
</pre>
