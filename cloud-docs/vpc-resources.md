---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Virtual Private Cloud API resources
{: #vpc-resources}

Create, update, or delete Virtual Private Cloud API resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_is_dedicated_host`
{: #is_dedicated_host}

Create, update, or delete an DedicatedHost.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_host-sample}

```
resource "ibm_is_dedicated_host" "is_dedicated_host" {
  dedicated_host_prototype = {"group":{"id":"0c8eccb4-271c-4518-956c-32bfce5cf83b"},"profile":{"name":"m-62x496"}}
}
```

### Input parameters
{: #is_dedicated_host-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`dedicated_host_prototype`|List|Required|The dedicated host prototype object. You can specify one item in this list only.|No|

### Output parameters
{: #is_dedicated_host-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the DedicatedHost.|
|`available_memory`|Integer|The amount of memory in gibibytes that is currently available for instances.|
|`available_vcpu`|List|The available VCPU for the dedicated host. This list contains only one item.|
|`available_vcpu.architecture`|String|The VCPU architecture.|
|`available_vcpu.count`|Integer|The number of VCPUs assigned.|
|`created_at`|String|The date and time that the dedicated host was created.|
|`crn`|String|The CRN for this dedicated host.|
|`group`|List|The dedicated host group this dedicated host is in. This list contains only one item.|
|`group.crn`|String|The CRN for this dedicated host group.|
|`group.deleted`|List|If present, this property indicates the referenced resource has been deleted and providessome supplementary information. This list contains only one item.|
|`group.deleted.more_info`|String|Link to documentation about deleted resources.|
|`group.href`|String|The URL for this dedicated host group.|
|`group.id`|String|The unique identifier for this dedicated host group.|
|`group.name`|String|The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.|
|`group.resource_type`|String|The type of resource referenced.|
|`href`|String|The URL for this dedicated host.|
|`instance_placement_enabled`|Boolean|If set to true, instances can be placed on this dedicated host.|
|`instances`|List|Array of instances that are allocated to this dedicated host.|
|`instances.crn`|String|The CRN for this virtual server instance.|
|`instances.deleted`|List|If present, this property indicates the referenced resource has been deleted and providessome supplementary information. This list contains only one item.|
|`instances.deleted.more_info`|String|Link to documentation about deleted resources.|
|`instances.href`|String|The URL for this virtual server instance.|
|`instances.id`|String|The unique identifier for this virtual server instance.|
|`instances.name`|String|The user-defined name for this virtual server instance (and default system hostname).|
|`lifecycle_state`|String|The lifecycle state of the dedicated host resource.|
|`memory`|Integer|The total amount of memory in gibibytes for this host.|
|`name`|String|The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.|
|`profile`|List|The profile this dedicated host uses. This list contains only one item.|
|`profile.href`|String|The URL for this dedicated host.|
|`profile.name`|String|The globally unique name for this dedicated host profile.|
|`provisionable`|Boolean|Indicates whether this dedicated host is available for instance creation.|
|`resource_group`|List|The resource group for this dedicated host. This list contains only one item.|
|`resource_group.href`|String|The URL for this resource group.|
|`resource_group.id`|String|The unique identifier for this resource group.|
|`resource_group.name`|String|The user-defined name for this resource group.|
|`resource_type`|String|The type of resource referenced.|
|`socket_count`|Integer|The total number of sockets for this host.|
|`state`|String|The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.|
|`supported_instance_profiles`|List|Array of instance profiles that can be used by instances placed on this dedicated host.|
|`supported_instance_profiles.href`|String|The URL for this virtual server instance profile.|
|`supported_instance_profiles.name`|String|The globally unique name for this virtual server instance profile.|
|`vcpu`|List|The total VCPU of the dedicated host. This list contains only one item.|
|`vcpu.architecture`|String|The VCPU architecture.|
|`vcpu.count`|Integer|The number of VCPUs assigned.|
|`zone`|List|The zone this dedicated host resides in. This list contains only one item.|
|`zone.href`|String|The URL for this zone.|
|`zone.name`|String|The globally unique name for this zone.|

### Import
{: #is_dedicated_host-import}

`ibm_is_dedicated_host` can be imported by ID

```
$ terraform import ibm_is_dedicated_host.example sample-id
```

