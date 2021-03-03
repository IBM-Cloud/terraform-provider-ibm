---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Virtual Private Cloud API data sources
{: #vpc-data-sources}

Review the data sources that you can use to retrieve information about your Virtual Private Cloud API resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_is_dedicated_host_profile`
{: #is_dedicated_host_profile}

Retrieve information about DedicatedHostProfile.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_host_profile-sample}

```
data "ibm_is_dedicated_host_profile" "is_dedicated_host_profile" {
  name = "bc1-4x16"
}
```

### Input parameters
{: #is_dedicated_host_profile-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`name`|String|Optional|The globally unique name for this virtual server instance profile.|

### Output parameters
{: #is_dedicated_host_profile-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`class`|String|The product class this dedicated host profile belongs to.|
|`family`|String|The product family this dedicated host profile belongs toThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.|
|`href`|String|The URL for this dedicated host.|
|`memory`|List| This list contains only one item.|
|`memory.type`|String|The type for this profile field.|
|`memory.value`|Integer|The value for this profile field.|
|`memory.default`|Integer|The default value for this profile field.|
|`memory.max`|Integer|The maximum value for this profile field.|
|`memory.min`|Integer|The minimum value for this profile field.|
|`memory.step`|Integer|The increment step value for this profile field.|
|`memory.values`|List|The permitted values for this profile field.|
|`socket_count`|List| This list contains only one item.|
|`socket_count.type`|String|The type for this profile field.|
|`socket_count.value`|Integer|The value for this profile field.|
|`socket_count.default`|Integer|The default value for this profile field.|
|`socket_count.max`|Integer|The maximum value for this profile field.|
|`socket_count.min`|Integer|The minimum value for this profile field.|
|`socket_count.step`|Integer|The increment step value for this profile field.|
|`socket_count.values`|List|The permitted values for this profile field.|
|`supported_instance_profiles`|List|Array of instance profiles that can be used by instances placed on dedicated hosts with this profile.|
|`supported_instance_profiles.href`|String|The URL for this virtual server instance profile.|
|`supported_instance_profiles.name`|String|The globally unique name for this virtual server instance profile.|
|`vcpu_architecture`|List| This list contains only one item.|
|`vcpu_architecture.type`|String|The type for this profile field.|
|`vcpu_architecture.value`|String|The VCPU architecture for a dedicated host with this profile.|
|`vcpu_count`|List| This list contains only one item.|
|`vcpu_count.type`|String|The type for this profile field.|
|`vcpu_count.value`|Integer|The value for this profile field.|
|`vcpu_count.default`|Integer|The default value for this profile field.|
|`vcpu_count.max`|Integer|The maximum value for this profile field.|
|`vcpu_count.min`|Integer|The minimum value for this profile field.|
|`vcpu_count.step`|Integer|The increment step value for this profile field.|
|`vcpu_count.values`|List|The permitted values for this profile field.|

## `ibm_is_dedicated_host_profiles`
{: #is_dedicated_host_profiles}

Retrieve information about DedicatedHostProfileCollection.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_host_profiles-sample}

```
data "ibm_is_dedicated_host_profiles" "is_dedicated_host_profiles" {
  name = "mx2-host-152x1216"
}
```

### Input parameters
{: #is_dedicated_host_profiles-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`name`|String|Optional|The globally unique name for this dedicated host profile.|

### Output parameters
{: #is_dedicated_host_profiles-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`first`|List|A link to the first page of resources. This list contains only one item.|
|`first.href`|String|The URL for a page of resources.|
|`limit`|Integer|The maximum number of resources that can be returned by the request.|
|`next`|List|A link to the next page of resources. This property is present for all pagesexcept the last page. This list contains only one item.|
|`next.href`|String|The URL for a page of resources.|
|`profiles`|List|Collection of dedicated host profiles.|
|`profiles.class`|String|The product class this dedicated host profile belongs to.|
|`profiles.family`|String|The product family this dedicated host profile belongs toThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.|
|`profiles.href`|String|The URL for this dedicated host.|
|`profiles.memory`|List| This list contains only one item.|
|`profiles.memory.type`|String|The type for this profile field.|
|`profiles.memory.value`|Integer|The value for this profile field.|
|`profiles.memory.default`|Integer|The default value for this profile field.|
|`profiles.memory.max`|Integer|The maximum value for this profile field.|
|`profiles.memory.min`|Integer|The minimum value for this profile field.|
|`profiles.memory.step`|Integer|The increment step value for this profile field.|
|`profiles.memory.values`|List|The permitted values for this profile field.|
|`profiles.name`|String|The globally unique name for this dedicated host profile.|
|`profiles.socket_count`|List| This list contains only one item.|
|`profiles.socket_count.type`|String|The type for this profile field.|
|`profiles.socket_count.value`|Integer|The value for this profile field.|
|`profiles.socket_count.default`|Integer|The default value for this profile field.|
|`profiles.socket_count.max`|Integer|The maximum value for this profile field.|
|`profiles.socket_count.min`|Integer|The minimum value for this profile field.|
|`profiles.socket_count.step`|Integer|The increment step value for this profile field.|
|`profiles.socket_count.values`|List|The permitted values for this profile field.|
|`profiles.supported_instance_profiles`|List|Array of instance profiles that can be used by instances placed on dedicated hosts with this profile.|
|`profiles.supported_instance_profiles.href`|String|The URL for this virtual server instance profile.|
|`profiles.supported_instance_profiles.name`|String|The globally unique name for this virtual server instance profile.|
|`profiles.vcpu_architecture`|List| This list contains only one item.|
|`profiles.vcpu_architecture.type`|String|The type for this profile field.|
|`profiles.vcpu_architecture.value`|String|The VCPU architecture for a dedicated host with this profile.|
|`profiles.vcpu_count`|List| This list contains only one item.|
|`profiles.vcpu_count.type`|String|The type for this profile field.|
|`profiles.vcpu_count.value`|Integer|The value for this profile field.|
|`profiles.vcpu_count.default`|Integer|The default value for this profile field.|
|`profiles.vcpu_count.max`|Integer|The maximum value for this profile field.|
|`profiles.vcpu_count.min`|Integer|The minimum value for this profile field.|
|`profiles.vcpu_count.step`|Integer|The increment step value for this profile field.|
|`profiles.vcpu_count.values`|List|The permitted values for this profile field.|
|`total_count`|Integer|The total number of resources across all pages.|

