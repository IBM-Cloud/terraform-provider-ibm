---

copyright:
  years: 2023
lastupdated: "2023"

keywords: terraform

subcollection: terraform

---

# Projects API Specification data sources
{: #project-data-sources}

Review the data sources that you can use to retrieve information about your Projects API Specification resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_project`
{: #project}

Retrieve information about project.
{: shortdesc}

### Sample Terraform code
{: #project-sample}

```
data "ibm_project" "project" {
  complete = true
  exclude_configs = true
  id = projectIdLink
}
```

### Input parameters
{: #project-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`complete`|Boolean|Optional|The flag to determine if full metadata should be returned. The default value is `false`.|
|`exclude_configs`|Boolean|Optional|Only return with the active configuration, no drafts. The default value is `false`.|
|`id`|String|Required|The ID of the project, which uniquely identifies it. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|

### Output parameters
{: #project-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`configs`|List|The project configurations. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.name`|String|The configuration name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`configs.labels`|List|A collection of configuration labels. The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.description`|String|A project configuration description. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s).*\\S$/`.|
|`configs.locator_id`|String|The location ID of a Project configuration manual property. The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`configs.type`|String|The type of a Project Config Manual Property. Allowable values are: `terraform_template`, `schematics_blueprint`.|
|`configs.input`|List|The outputs of a Schematics template property. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.input.name`|String|The variable name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`configs.input.type`|String|The variable type. Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.|
|`configs.input.required`|Boolean|Whether the variable is required or not.|
|`configs.output`|List|The outputs of a Schematics template property. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.output.name`|String|The variable name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`configs.output.description`|String|A short explanation of the output value. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`configs.output.value`|List|The output value. The list items must match regular expression `/^(?!\\s).+\\S$/`. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.setting`|List|An optional setting object That is passed to the cart API. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.setting.name`|String|The name of the configuration setting. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`configs.setting.value`|String|The value of a the configuration setting. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`description`|String|A project descriptive text.|
|`metadata`|List|Metadata of the project. This list contains only one item.|
|`metadata.crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`metadata.created_at`|String|A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339.|
|`metadata.cumulative_needs_attention_view`|List|The cumulative list of needs attention items of a project. The maximum length is `10000` items. The minimum length is `0` items.|
|`metadata.cumulative_needs_attention_view.event`|String|The event name.|
|`metadata.cumulative_needs_attention_view.event_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_version`|Integer|The version number of the configuration.|
|`metadata.cumulative_needs_attention_view_err`|String|True to indicate the fetch of needs attention items that failed.|
|`metadata.location`|String|The location of where the project was created.|
|`metadata.resource_group`|String|The resource group of where the project was created.|
|`metadata.state`|String|The project status value. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(CREATING|CREATING_FAILED|UPDATING|UPDATING_FAILED|READY)$/`.|
|`metadata.event_notifications_crn`|String|The CRN of the event notifications instance if one is connected to this project.|
|`name`|String|The project name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|

## `ibm_event_notification`
{: #event_notification}

Retrieve information about Get Event Notifications Integration response.
{: shortdesc}

### Sample Terraform code
{: #event_notification-sample}

```
data "ibm_event_notification" "event_notification" {
  id = "id"
}
```

### Input parameters
{: #event_notification-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`id`|String|Required|The ID of the project, which uniquely identifies it. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|

### Output parameters
{: #event_notification-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`description`|String|A description of the instance of the event. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`enabled`|Boolean|The status of instance of the event.|
|`name`|String|The name of the instance of the event. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`topic_count`|Integer|The topic count of the instance of the event.|
|`topic_names`|List|The topic names of the instance of the event. The list items must match regular expression `/^(?!\\s).+\\S$/`. The maximum length is `10000` items. The minimum length is `0` items.|
|`type`|String|The type of the instance of event. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.|
|`updated_at`|String|A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339.|

