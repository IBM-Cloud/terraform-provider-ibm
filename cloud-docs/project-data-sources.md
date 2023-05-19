---

copyright:
  years: 2023
lastupdated: "2023"

keywords: terraform

subcollection: terraform

---

# Projects API data sources
{: #project-data-sources}

Review the data sources that you can use to retrieve information about your Projects API resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_project`
{: #project}

Retrieve information about Project definition.
{: shortdesc}

### Sample Terraform code
{: #project-sample}

```
data "ibm_project" "project" {
  id = "id"
}
```

### Input parameters
{: #project-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`id`|String|Required|The unique project ID. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|

### Output parameters
{: #project-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`configs`|List|The project configurations. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.id`|String|The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.name`|String|The configuration name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.|
|`configs.labels`|List|A collection of configuration labels. The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.description`|String|The project configuration description. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`configs.authorizations`|List|The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager. This list contains only one item.|
|`configs.authorizations.trusted_profile`|List|The trusted profile for authorizations. This list contains only one item.|
|`configs.authorizations.trusted_profile.id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.authorizations.trusted_profile.target_iam_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.authorizations.method`|String|The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`configs.authorizations.api_key`|String|The IBM Cloud API Key. The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.|
|`configs.compliance_profile`|List|The profile required for compliance. This list contains only one item.|
|`configs.compliance_profile.id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.compliance_profile.instance_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.compliance_profile.instance_location`|String|The location of the compliance instance. The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.|
|`configs.compliance_profile.attachment_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`configs.compliance_profile.profile_name`|String|The name of the compliance profile. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.|
|`configs.locator_id`|String|A dotted value of catalogID.versionID. The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.|
|`configs.type`|String|The type of a project configuration manual property. Allowable values are: `terraform_template`, `schematics_blueprint`.|
|`configs.input`|List|The outputs of a Schematics template property. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.input.name`|String|The variable name. The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.|
|`configs.input.type`|String|The variable type. Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.|
|`configs.input.value`|String|Can be any value - a string, number, boolean, array, or object.|
|`configs.input.required`|Boolean|Whether the variable is required or not.|
|`configs.output`|List|The outputs of a Schematics template property. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.output.name`|String|The variable name. The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.|
|`configs.output.description`|String|A short explanation of the output value. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`configs.output.value`|String|Can be any value - a string, number, boolean, array, or object.|
|`configs.setting`|List|Schematics environment variables to use to deploy the configuration. The maximum length is `10000` items. The minimum length is `0` items.|
|`configs.setting.name`|String|The name of the configuration setting. The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.|
|`configs.setting.value`|String|The value of the configuration setting. The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.|
|`crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`description`|String|A project descriptive text. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`destroy_on_delete`|Boolean|The policy that indicates whether the resources are destroyed or not when a project is deleted. The default value is `true`.|
|`metadata`|List|The metadata of the project. This list contains only one item.|
|`metadata.crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`metadata.created_at`|String|A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.|
|`metadata.cumulative_needs_attention_view`|List|The cumulative list of needs attention items for a project. The maximum length is `10000` items. The minimum length is `0` items.|
|`metadata.cumulative_needs_attention_view.event`|String|The event name. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.cumulative_needs_attention_view.event_id`|String|A unique ID for that individual event. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_id`|String|A unique ID for the configuration. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_version`|Integer|The version number of the configuration.|
|`metadata.cumulative_needs_attention_view_err`|String|True indicates that the fetch of the needs attention items failed. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.location`|String|The IBM Cloud location where a resource is deployed. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.resource_group`|String|The resource group where the project's data and tools are created. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.state`|String|The project status value. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(CREATING|CREATING_FAILED|UPDATING|UPDATING_FAILED|READY)$/`.|
|`metadata.event_notifications_crn`|String|The CRN of the event notifications instance if one is connected to this project. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`name`|String|The project name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.|

