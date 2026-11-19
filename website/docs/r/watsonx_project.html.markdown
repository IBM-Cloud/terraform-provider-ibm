---
subcategory: "watsonx"
layout: "ibm"
page_title: "IBM : ibm_watsonx_project"
description: |-
  Manages a watsonx.ai project.
---

# ibm_watsonx_project

Creates a watsonx.ai project backed by Cloud Object Storage and bound to a watsonx.ai Runtime instance.

## Example usage

```terraform
resource "ibm_watsonx_project" "this" {
  region            = "us-south"
  name              = "my-project"
  description       = "Created by Terraform"
  tags              = ["env:dev"]
  cos_instance_id   = module.cos.cos_instance_guid
  cos_crn           = module.cos.cos_instance_crn
  storage_delegated = true
  runtime_name      = ibm_resource_instance.runtime.name
  runtime_guid      = ibm_resource_instance.runtime.guid
  runtime_crn       = ibm_resource_instance.runtime.crn
}
```

## Argument reference

- `region` - (Required, Forces new resource, String) watsonx region.
- `endpoint` - (Optional, Forces new resource, String) Override for the Data Platform API host.
- `name` - (Required, String) Project name, unique in the account.
- `description` - (Optional, String) Description.
- `tags` - (Optional, Forces new resource, List of String) Tags. Create-only in the API.
- `cos_instance_id` - (Required, Forces new resource, String) COS instance GUID.
- `cos_crn` - (Required, Forces new resource, String) COS instance CRN.
- `storage_delegated` - (Optional, Forces new resource, Boolean) COS is delegated. Default `false`.
- `runtime_name`, `runtime_guid`, `runtime_crn` - (Required, String) watsonx.ai Runtime instance.
- `mark_as_sensitive` - (Optional, Forces new resource, Boolean) Restrict data leaving the project. Default `false`.

## Attribute reference

- `id` - Project GUID.
- `location` - API path, `/v2/projects/<guid>`.
- `project_url` - Link to the project in the watsonx UI.

## Import

```
terraform import ibm_watsonx_project.this us-south/<project_guid>
```
