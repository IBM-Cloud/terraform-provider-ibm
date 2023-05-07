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

## `ibm_project_instance`
{: #project_instance}

Create, update, or delete an Project definition.
{: shortdesc}

### Sample Terraform code
{: #project_instance-sample}

```
resource "ibm_project_instance" "project_instance" {
  description = "A microservice to deploy on top of ACME infrastructure."
  location = "us-south"
  name = "acme-microservice"
  resource_group = "Default"
}
```

### Input parameters
{: #project_instance-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`configs`|List|Optional|The project configurations. The maximum length is `10000` items. The minimum length is `0` items.|No|
|`description`|String|Optional|A project's descriptive text. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.|No|
|`location`|String|Required|The location where the project's data and tools are created. The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.|No|
|`name`|String|Required|The project name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.|No|
|`resource_group`|String|Required|The resource group where the project's data and tools are created. The maximum length is `40` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|No|

### Output parameters
{: #project_instance-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the Project definition.|
|`crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`metadata`|List|The metadata of the project. This list contains only one item.|
|`metadata.crn`|String|An IBM Cloud resource name, which uniquely identifies a resource. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|
|`metadata.created_at`|String|A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.|
|`metadata.cumulative_needs_attention_view`|List|The cumulative list of needs attention items for a project. The maximum length is `10000` items. The minimum length is `0` items.|
|`metadata.cumulative_needs_attention_view.event`|String|The event name. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.cumulative_needs_attention_view.event_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_id`|String|The unique ID of a project. The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.|
|`metadata.cumulative_needs_attention_view.config_version`|Integer|The version number of the configuration.|
|`metadata.cumulative_needs_attention_view_err`|String|\"True\" indicates that the fetch of the needs attention items failed. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.location`|String|The IBM Cloud location where a resource is deployed. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.resource_group`|String|The resource group where the project's data and tools are created. The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.|
|`metadata.state`|String|The project status value. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(CREATING|CREATING_FAILED|UPDATING|UPDATING_FAILED|READY)$/`.|
|`metadata.event_notifications_crn`|String|The CRN of the event notifications instance if one is connected to this project. The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.|

### Import
{: #project_instance-import}

`ibm_project_instance` can be imported by ID

```
$ terraform import ibm_project_instance.example sample-id
```

