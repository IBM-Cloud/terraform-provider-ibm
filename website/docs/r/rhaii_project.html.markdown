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
  tags              = ["env:dev"]
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
* `location` - (Optional, Forces new resource, String) The region where the project is created. The default value is `us-east`, which is the only region where the service is available today.
* `plan` - (Optional, String) The pricing plan of the project. The default value is `instructlab-pricing-plan`.
* `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. If you do not set it, the default resource group of the account is used.
* `tags` - (Optional, Array of Strings) The tags that you want to add to the project.
* `parameters` - (Optional, Map) Arbitrary parameters to pass to the service broker. Conflicts with `parameters_json`.
* `parameters_json` - (Optional, String) Arbitrary parameters to pass to the service broker, in JSON string format. Conflicts with `parameters`.
* `service_endpoints` - (Optional, String) The types of the service endpoints. Supported values are `public`, `private`, and `public-and-private`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - (String) The unique identifier of the project. This is the CRN of the resource instance.
* `project_id` - (String) The ID of the project. Use this value as `project_id` in the Red Hat AI Inference API. It is the same value as `guid`.
* `endpoint` - (String) The base URL of the Red Hat AI Inference API for this project, for example `https://us-east.rhai.ibm.com/v1/projects/<project_id>`.
* `guid` - (String) The GUID of the resource instance.
* `crn` - (String) The CRN of the project.
* `service` - (String) The service name of the instance. Always `instructlab`.
* `status` - (String) The status of the project.
* `state` - (String) The current state of the project, for example `active`.
* `dashboard_url` - (String) The relative URL of the project in the IBM Cloud console.
* `account_id` - (String) The ID of the account that owns the project.
* `resource_group_crn` - (String) The CRN of the resource group.
* `resource_plan_id` - (String) The ID of the plan of the project.
* `target_crn` - (String) The deployment CRN in the global catalog.
* `created_at` - (String) The date when the project was created.
* `created_by` - (String) The subject who created the project.
* `last_operation` - (Map) The status of the last operation on the project.
* `plan_history` - (List) The plan history of the project.

The resource also exports the other attributes of the [`ibm_resource_instance`](resource_instance.html) resource.

## Import

You can import the `ibm_rhaii_project` resource by using the CRN of the project.

# Syntax
<pre>
$ terraform import ibm_rhaii_project.project &lt;crn&gt;
</pre>

# Example
```
$ terraform import ibm_rhaii_project.project "crn:v1:bluemix:public:instructlab:us-east:a/<account_id>:<project_id>::"
```
