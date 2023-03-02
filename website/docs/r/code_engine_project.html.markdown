---
layout: "ibm"
page_title: "IBM : ibm_code_engine_project"
description: |-
  Manages code_engine_project.
subcategory: "Code Engine"
---

# ibm_code_engine_project

Provides a resource for ibm_code_engine_project. This allows ibm_code_engine_project to be created and deleted.

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_code_engine_project" "code_engine_project_instance" {
  name = "<your_project_name>"
  resource_group_id = data.ibm_resource_group.group.id
}
```

## Timeouts

* `Create` The creation of an instance is considered failed when no response is received for 2 minutes.
* `Delete` The deletion of an instance is considered failed when no response is received for 2 minutes.

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Required, Forces new resource, String) The name of the project.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9\\-._: ])+$/`.
* `resource_group_id` - (Required, Forces new resource, String) ID of the resource group for your project deployment. If this field is not defined, the default resource group of the account will be used.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-z0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the code_engine_project.
* `account_id` - (String) An alphanumeric value identifying the account ID.
* `created_at` - (String) The date when the project was created.
* `crn` - (String) The CRN of the project.
* `href` - (String) When you provision a new resource, a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region for your project deployment. Possible values: `au-syd`, `br-sao`, `ca-tor`, `eu-de`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`.
  * Constraints: The maximum length is `48` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z0-9-]+$/`.
* `resource_type` - (String) The type of the project.
  * Constraints: Allowable values are: `project_v2`.
* `status` - (String) The current state of the project. For example, if the project is created and ready to get used, it will return active.
  * Constraints: Possible values are: `active`, `inactive`, `pending_removal`, `hard_deleting`, `hard_deletion_failed`, `hard_deleted`, `deleting`, `deletion_failed`, `soft_deleted`, `preparing`, `creating`, `creation_failed`.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key` argument in the IBM Cloud provider block.

Usage:
```hcl
provider "ibm" {
    ibmcloud_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY` environment variables, representing your IBM Cloud platform API key.

```hcl
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_code_engine_project` resource by using `id`. The ID of the project.

# Syntax
```
$ terraform import ibm_code_engine_project.code_engine_project <id>
```

# Example
```
$ terraform import ibm_code_engine_project.code_engine_project 4e49b3e0-27a8-48d2-a784-c7ee48bb863b
```
