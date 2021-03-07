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

## `ibm_is_dedicated_host_group`
{: #is_dedicated_host_group}

Create, update, or delete an DedicatedHostGroup.
{: shortdesc}

### Sample Terraform code
{: #is_dedicated_host_group-sample}

```
resource "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
  class = "mx2"
  family = "balanced"
  name = "placeholder"
  resource_group = var.is_dedicated_host_group_resource_group
  zone = {"name":"us-south-1"}
}
```

### Input parameters
{: #is_dedicated_host_group-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`class`|String|Optional|The dedicated host profile class for hosts in this group.|No|
|`family`|String|Optional|The dedicated host profile family for hosts in this group.|No|
|`name`|String|Optional|The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.|No|
|`resource_group`|List|Optional|The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used. You can specify one item in this list only.|No|
|`zone`|List|Optional|The zone this dedicated host group will reside in. You can specify one item in this list only.|No|

### Output parameters
{: #is_dedicated_host_group-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the DedicatedHostGroup.|
|`created_at`|String|The date and time that the dedicated host group was created.|
|`crn`|String|The CRN for this dedicated host group.|
|`dedicated_hosts`|List|The dedicated hosts that are in this dedicated host group.|
|`dedicated_hosts.crn`|String|The CRN for this dedicated host.|
|`dedicated_hosts.deleted`|List|If present, this property indicates the referenced resource has been deleted and providessome supplementary information. This list contains only one item.|
|`dedicated_hosts.deleted.more_info`|String|Link to documentation about deleted resources.|
|`dedicated_hosts.href`|String|The URL for this dedicated host.|
|`dedicated_hosts.id`|String|The unique identifier for this dedicated host.|
|`dedicated_hosts.name`|String|The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.|
|`dedicated_hosts.resource_type`|String|The type of resource referenced.|
|`href`|String|The URL for this dedicated host group.|
|`resource_type`|String|The type of resource referenced.|
|`supported_instance_profiles`|List|Array of instance profiles that can be used by instances placed on this dedicated host group.|
|`supported_instance_profiles.href`|String|The URL for this virtual server instance profile.|
|`supported_instance_profiles.name`|String|The globally unique name for this virtual server instance profile.|

### Import
{: #is_dedicated_host_group-import}

`ibm_is_dedicated_host_group` can be imported by ID

```
$ terraform import ibm_is_dedicated_host_group.example sample-id
```

