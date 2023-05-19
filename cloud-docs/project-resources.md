---

copyright:
  years: 2023
lastupdated: "2023"

keywords: terraform

subcollection: terraform

---

# Projects API resources
{: #project-resources}

Create, update, or delete Projects API resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_project`
{: #project}

Create, update, or delete an project.
{: shortdesc}

### Sample Terraform code
{: #project-sample}

```
resource "ibm_project" "project" {
  description = "A microservice to deploy on top of ACME infrastructure."
  location = "us-south"
  name = "acme-microservice"
  resource_group = "Default"
}
```

### Input parameters
{: #project-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`configs`|List|Optional|The project configurations. The maximum length is `10000` items. The minimum length is `0` items.|Yes|
|`description`|String|Optional|A project's descriptive text. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|Yes|
|`destroy_on_delete`|Boolean|Optional|The policy that indicates whether the resources are destroyed or not when a project is deleted. The default value is `true`.|Yes|
|`location`|String|Required|The location where the project's data and tools are created. The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.|Yes|
|`name`|String|Required|The project name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.|Yes|
|`resource_group`|String|Required|The resource group where the project's data and tools are created. The maximum length is `40` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|Yes|

### Output parameters
{: #project-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the project.|
|`crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
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

### Import
{: #project-import}

`ibm_project` can be imported by ID

```
$ terraform import ibm_project.example sample-id
```

## `ibm_project_config`
{: #project_config}

Create, update, or delete an project_config.
{: shortdesc}

### Sample Terraform code
{: #project_config-sample}

```
resource "ibm_project_config" "project_config" {
  description = "Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace."
  labels = ["env:stage","governance:test","build:0"]
  locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
  name = "env-stage"
  project_id = ibm_project.project_instance.id
}
```

### Input parameters
{: #project_config-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`authorizations`|List|Optional|The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager. You can specify one item in this list only.|No|
|`compliance_profile`|List|Optional|The profile required for compliance. You can specify one item in this list only.|No|
|`description`|String|Optional|The project configuration description. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|No|
|`input`|List|Optional|The input values to use to deploy the configuration. The maximum length is `10000` items. The minimum length is `0` items.|No|
|`labels`|List|Optional|A collection of configuration labels. The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.|No|
|`locator_id`|String|Required|A dotted value of catalogID.versionID. The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.|No|
|`name`|String|Required|The configuration name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.|No|
|`project_id`|String|Required|The unique project ID. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|Yes|
|`setting`|List|Optional|Schematics environment variables to use to deploy the configuration. The maximum length is `10000` items. The minimum length is `0` items.|No|

### Output parameters
{: #project_config-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the project_config.|
|`output`|List|The outputs of a Schematics template property. The maximum length is `10000` items. The minimum length is `0` items.|
|`output.name`|String|The variable name. The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.|
|`output.description`|String|A short explanation of the output value. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|
|`output.value`|String|Can be any value - a string, number, boolean, array, or object.|
|`project_config_id`|String|The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`type`|String|The type of a project configuration manual property. Allowable values are: `terraform_template`, `schematics_blueprint`.|

### Import
{: #project_config-import}

`ibm_project_config` can be imported by ID

```
$ terraform import ibm_project_config.example sample-id
```

