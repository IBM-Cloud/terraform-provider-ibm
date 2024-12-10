---
layout: "ibm"
page_title: "IBM : ibm_schematics_agent_prs"
description: |-
  Get information about schematics_agent_prs
subcategory: "Schematics"
---

# ibm_schematics_agent_prs

Provides a read-only data source for schematics_agent_prs. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_agent_prs" "schematics_agent_prs" {
	agent_id = ibm_schematics_agent_prs.schematics_agent_prs.agent_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `agent_id` - (Required, Forces new resource, String) Agent ID to get the details of agent.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_agent_prs.
* `agent_version` - (String) Agent version.

* `job_id` - (String) Job Id.

* `log_url` - (String) URL to the full pre-requisite scanner job logs.

* `status_code` - (String) Final result of the pre-requisite scanner job.
  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.

* `status_message` - (String) The outcome of the pre-requisite scanner job, in a formatted log string.

* `updated_at` - (String) The agent prs job updation time.

* `updated_by` - (String) Email address of user who ran the agent prs job.

