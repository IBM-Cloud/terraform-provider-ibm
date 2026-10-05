---
layout: "ibm"
page_title: "IBM : ibm_rhaii_project"
description: |-
  Get information about a Red Hat AI Inference project.
subcategory: "Red Hat AI Inference"
---

# ibm_rhaii_project

Retrieve information about an existing Red Hat AI Inference (RHAII) project. You can look up the project by its name or by its ID.

For more information, see [Red Hat AI Inference on IBM Cloud](https://cloud.ibm.com/docs/inference) and the [Red Hat AI Inference API](https://cloud.ibm.com/docs/apis/inference).

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "Default"
}

data "ibm_rhaii_project" "by_name" {
  name              = "my-rhaii-project"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_rhaii_project" "by_id" {
  identifier = "917bc95a-fef0-4039-b936-e0b6fb17b721"
}
```

## Argument Reference

You can specify the following arguments for this data source. You must set exactly one of `name` or `identifier`.

* `name` - (Optional, String) The name of the project.
* `identifier` - (Optional, String) The ID (GUID) of the project. Conflicts with `name`, `resource_group_id`, and `location`.
* `resource_group_id` - (Optional, String) The ID of the resource group to search in. Use it when more than one project has the same name.
* `location` - (Optional, String) The region of the project, for example `us-east`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) The unique identifier of the project. This is the CRN of the resource instance.
* `project_id` - (String) The ID of the project. Use this value as `project_id` in the Red Hat AI Inference API.
* `endpoint` - (String) The base URL of the Red Hat AI Inference API for this project, for example `https://us-east.rhai.ibm.com/v1/projects/<project_id>`.
* `guid` - (String) The GUID of the resource instance.
* `crn` - (String) The CRN of the project.
* `service` - (String) The service name of the instance. Always `instructlab`.
* `plan` - (String) The pricing plan of the project.
* `status` - (String) The status of the project.
* `tags` - (Array of Strings) The tags of the project.
* `extensions` - (Map) The extended metadata of the project.
