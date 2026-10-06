---
layout: "ibm"
page_title: "IBM : ibm_rhaii_project"
description: |-
  Manages a Red Hat AI Inference project.
subcategory: "Red Hat AI Inference"
---

# ibm_rhaii_project

Create, update, and delete a Red Hat AI Inference (RHAII) project with this resource.

An RHAII project is an instance of the `instructlab` service in the IBM Cloud resource controller. This resource is the Terraform equivalent of the following IBM Cloud CLI command.

```sh
ibmcloud resource service-instance-create <name> instructlab instructlab-pricing-plan us-east -g <resource_group>
```

The provider resolves the service, the plan, and the deployment for the location from the global catalog. You do not need to look up the plan ID or the deployment CRN yourself.

For more information, see [Red Hat AI Inference on IBM Cloud](https://cloud.ibm.com/docs/inference) and the [Red Hat AI Inference API](https://cloud.ibm.com/docs/apis/inference).

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "Default"
}

resource "ibm_rhaii_project" "project" {
  name              = "my-rhaii-project"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["env:dev", "team:ai"]
  access_tags       = ["project:rhaii"]
}

output "rhaii_project_id" {
  value = ibm_rhaii_project.project.project_id
}

output "rhaii_endpoint" {
  value = ibm_rhaii_project.project.endpoint
}
```

## Timeouts

The `ibm_rhaii_project` resource provides the following [Timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating the project.
* `update` - (Default 10 minutes) Used for updating the project.
* `delete` - (Default 10 minutes) Used for deleting the project.

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, String) The name of the project.
* `location` - (Optional, Forces new resource, String) The region where the project is created. The default value is `us-east`, which is the only region where the service is available today. If you set a region where the service is not deployed, the error lists the valid regions.
* `plan` - (Optional, String) The pricing plan of the project. The default value is `instructlab-pricing-plan`. You can set the plan name or the plan ID from the global catalog.
* `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. If you do not set it, the default resource group of the account is used.
* `tags` - (Optional, Array of Strings) The user tags of the project. User tags help you organize and search for resources. They do not control access. Each tag can have up to 128 characters, and can contain letters, numbers, spaces, `_`, `.`, `-` and `:`.
* `access_tags` - (Optional, Array of Strings) The access management tags of the project, in `key:value` format. Use them in IAM access policies to control who can work with the project. The tags must already exist in the account. You can create them with the `ibm_resource_tag` resource or in the console.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - (String) The unique identifier of the project. This is the CRN of the resource instance.
* `project_id` - (String) The ID of the project. Use this value as `project_id` in the Red Hat AI Inference API. It is the same value as `guid`.
* `endpoint` - (String) The base URL of the Red Hat AI Inference API for this project, for example `https://us-east.rhai.ibm.com/v1/projects/<project_id>`.
* `guid` - (String) The GUID of the resource instance.
* `crn` - (String) The CRN of the project.
* `service` - (String) The service name of the project. Always `instructlab`.
* `state` - (String) The state of the project, for example `active`.
* `dashboard_url` - (String) The relative URL of the project in the IBM Cloud console.
* `account_id` - (String) The ID of the account that owns the project.
* `resource_group_crn` - (String) The CRN of the resource group.
* `resource_plan_id` - (String) The catalog ID of the plan of the project.
* `target_crn` - (String) The deployment CRN of the project in the global catalog.
* `created_at` - (String) The date when the project was created.
* `created_by` - (String) The subject who created the project.
* `updated_at` - (String) The date when the project was last updated.
* `updated_by` - (String) The subject who last updated the project.
* `locked` - (Boolean) Whether the project is locked.
* `last_operation` - (List) The last operation on the project.
  Nested schema for **last_operation**:
  * `type` - (String) The type of the operation, for example `create`.
  * `state` - (String) The state of the operation, for example `succeeded`.
  * `async` - (Boolean) Whether the operation is asynchronous.
  * `description` - (String) The description of the operation.

## Import

You can import the `ibm_rhaii_project` resource by using the CRN or the GUID (project ID) of the project. After the import, the ID in the state is always the CRN.

# Syntax
<pre>
$ terraform import ibm_rhaii_project.project &lt;crn&gt;
</pre>

# Example
```
$ terraform import ibm_rhaii_project.project "crn:v1:bluemix:public:instructlab:us-east:a/<account_id>:<project_id>::"
```
