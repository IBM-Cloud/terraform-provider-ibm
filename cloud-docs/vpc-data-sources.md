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

## `ibm_is_dedicated_host`
{: #is_dedicated_host}

Retrieve information about DedicatedHost.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_host-sample}

```
data "ibm_is_dedicated_host" "is_dedicated_host" {
  id = "1e09281b-f177-46fb-baf1-bc152b2e391a"
}
```

### Input parameters
{: #is_dedicated_host-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`id`|String|Optional|The unique identifier for this virtual server instance.|

### Output parameters
{: #is_dedicated_host-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
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

## `ibm_is_dedicated_hosts`
{: #is_dedicated_hosts}

Retrieve information about DedicatedHostCollection.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_hosts-sample}

```
data "ibm_is_dedicated_hosts" "is_dedicated_hosts" {
  id = "1e09281b-f177-46fb-baf1-bc152b2e391a"
}
```

### Input parameters
{: #is_dedicated_hosts-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`id`|String|Optional|The unique identifier for this dedicated host.|

### Output parameters
{: #is_dedicated_hosts-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`dedicated_hosts`|List|Collection of dedicated hosts.|
|`dedicated_hosts.available_memory`|Integer|The amount of memory in gibibytes that is currently available for instances.|
|`dedicated_hosts.available_vcpu`|List|The available VCPU for the dedicated host. This list contains only one item.|
|`dedicated_hosts.available_vcpu.architecture`|String|The VCPU architecture.|
|`dedicated_hosts.available_vcpu.count`|Integer|The number of VCPUs assigned.|
|`dedicated_hosts.created_at`|String|The date and time that the dedicated host was created.|
|`dedicated_hosts.crn`|String|The CRN for this dedicated host.|
|`dedicated_hosts.group`|List|The dedicated host group this dedicated host is in. This list contains only one item.|
|`dedicated_hosts.group.crn`|String|The CRN for this dedicated host group.|
|`dedicated_hosts.group.deleted`|List|If present, this property indicates the referenced resource has been deleted and providessome supplementary information. This list contains only one item.|
|`dedicated_hosts.group.deleted.more_info`|String|Link to documentation about deleted resources.|
|`dedicated_hosts.group.href`|String|The URL for this dedicated host group.|
|`dedicated_hosts.group.id`|String|The unique identifier for this dedicated host group.|
|`dedicated_hosts.group.name`|String|The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.|
|`dedicated_hosts.group.resource_type`|String|The type of resource referenced.|
|`dedicated_hosts.href`|String|The URL for this dedicated host.|
|`dedicated_hosts.id`|String|The unique identifier for this dedicated host.|
|`dedicated_hosts.instance_placement_enabled`|Boolean|If set to true, instances can be placed on this dedicated host.|
|`dedicated_hosts.instances`|List|Array of instances that are allocated to this dedicated host.|
|`dedicated_hosts.instances.crn`|String|The CRN for this virtual server instance.|
|`dedicated_hosts.instances.deleted`|List|If present, this property indicates the referenced resource has been deleted and providessome supplementary information. This list contains only one item.|
|`dedicated_hosts.instances.deleted.more_info`|String|Link to documentation about deleted resources.|
|`dedicated_hosts.instances.href`|String|The URL for this virtual server instance.|
|`dedicated_hosts.instances.id`|String|The unique identifier for this virtual server instance.|
|`dedicated_hosts.instances.name`|String|The user-defined name for this virtual server instance (and default system hostname).|
|`dedicated_hosts.lifecycle_state`|String|The lifecycle state of the dedicated host resource.|
|`dedicated_hosts.memory`|Integer|The total amount of memory in gibibytes for this host.|
|`dedicated_hosts.name`|String|The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.|
|`dedicated_hosts.profile`|List|The profile this dedicated host uses. This list contains only one item.|
|`dedicated_hosts.profile.href`|String|The URL for this dedicated host.|
|`dedicated_hosts.profile.name`|String|The globally unique name for this dedicated host profile.|
|`dedicated_hosts.provisionable`|Boolean|Indicates whether this dedicated host is available for instance creation.|
|`dedicated_hosts.resource_group`|List|The resource group for this dedicated host. This list contains only one item.|
|`dedicated_hosts.resource_group.href`|String|The URL for this resource group.|
|`dedicated_hosts.resource_group.id`|String|The unique identifier for this resource group.|
|`dedicated_hosts.resource_group.name`|String|The user-defined name for this resource group.|
|`dedicated_hosts.resource_type`|String|The type of resource referenced.|
|`dedicated_hosts.socket_count`|Integer|The total number of sockets for this host.|
|`dedicated_hosts.state`|String|The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.|
|`dedicated_hosts.supported_instance_profiles`|List|Array of instance profiles that can be used by instances placed on this dedicated host.|
|`dedicated_hosts.supported_instance_profiles.href`|String|The URL for this virtual server instance profile.|
|`dedicated_hosts.supported_instance_profiles.name`|String|The globally unique name for this virtual server instance profile.|
|`dedicated_hosts.vcpu`|List|The total VCPU of the dedicated host. This list contains only one item.|
|`dedicated_hosts.vcpu.architecture`|String|The VCPU architecture.|
|`dedicated_hosts.vcpu.count`|Integer|The number of VCPUs assigned.|
|`dedicated_hosts.zone`|List|The zone this dedicated host resides in. This list contains only one item.|
|`dedicated_hosts.zone.href`|String|The URL for this zone.|
|`dedicated_hosts.zone.name`|String|The globally unique name for this zone.|
|`first`|List|A link to the first page of resources. This list contains only one item.|
|`first.href`|String|The URL for a page of resources.|
|`limit`|Integer|The maximum number of resources that can be returned by the request.|
|`next`|List|A link to the next page of resources. This property is present for all pagesexcept the last page. This list contains only one item.|
|`next.href`|String|The URL for a page of resources.|
|`total_count`|Integer|The total number of resources across all pages.|

