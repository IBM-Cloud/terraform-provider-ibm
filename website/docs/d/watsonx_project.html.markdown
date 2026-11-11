---
subcategory: "watsonx"
layout: "ibm"
page_title: "IBM : ibm_watsonx_project"
description: |-
  Reads a watsonx.ai project.
---

# ibm_watsonx_project (Data Source)

Reads a watsonx.ai project by GUID.

## Example usage

```terraform
data "ibm_watsonx_project" "this" {
  region     = "us-south"
  project_id = ibm_watsonx_project.this.id
}
```

## Argument reference

- `region` - (Required, String) watsonx region.
- `project_id` - (Required, String) Project GUID.
- `endpoint` - (Optional, String) Override for the Data Platform API host.

## Attribute reference

- `name`, `description`, `tags`, `cos_instance_id`, `cos_crn`, `project_url`.
- `project_json` - Full API record. Same content as `api_response` on `data.restapi_object`, so `jsondecode()` callers keep working.
