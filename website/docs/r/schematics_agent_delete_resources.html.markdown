---
layout: "ibm"
page_title: "IBM : ibm_schematics_agent_delete_resources"
description: |-
  Manages schematics_agent_delete_resources.
subcategory: "Schematics"
---

# ibm_schematics_agent_delete_resources

~> **Beta:** This resource is in Beta, and is subject to change.

Provides a resource for schematics_agent_delete_resources. This allows schematics_agent_delete_resources to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_agent_delete_resources" "schematics_agent_delete_resources_instance" {
  agent_id = "agent_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `agent_id` - (Required, Forces new resource, String) Agent ID to get the details of agent.
* `force` - (Optional, Boolean) Equivalent to -force options in the command line, default is false.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_agent_delete_resources.
* `agent_version` - (String) Agent version.
* `job_id` - (String) Job Id.
* `log_url` - (String) URL to the full agent resources destroy job logs.
* `status_code` - (String) Final result of the agent resources destroy job.
  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.
* `status_message` - (String) The outcome of the agent resources destroy job, in a formatted log string.
* `updated_at` - (String) The agent resources destroy job updation time.
* `updated_by` - (String) Email address of user who ran the agent resources destroy job.

## Import

You can import the `ibm_schematics_agent_delete_resources` resource by using `agent_id`.
The `agent_id` property can be formed from `agent_id`, and `agent_id` in the following format:

```
<agent_id>/<agent_delete_resources_job_id>
```
* `agent_id`: A string. Agent ID to get the details of agent.
* `agent_delete_resources_job_id`: A string. ID of the recent AgentDestroy Job.

# Syntax
```
$ terraform import ibm_schematics_agent_delete_resources.schematics_agent_delete_resources_instance <agent_id>/<agent_delete_resources_job_id>
```
