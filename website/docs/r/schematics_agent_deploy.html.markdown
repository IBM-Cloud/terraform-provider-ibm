---
layout: "ibm"
page_title: "IBM : ibm_schematics_agent_deploy"
description: |-
  Manages schematics_agent_deploy.
subcategory: "Schematics"
---

# ibm_schematics_agent_deploy

Provides a resource for schematics_agent_deploy. This allows schematics_agent_deploy to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_agent_deploy" "schematics_agent_deploy_instance" {
  agent_id = "agent_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `agent_id` - (Required, Forces new resource, String) Agent ID to get the details of agent.
* `force` - (Optional, Boolean) Equivalent to -force options in the command line, default is false.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_agent_deploy.
* `agent_version` - (String) Agent version.
* `is_redeployed` - (Boolean) True, when the same version of the agent was redeployed.
* `job_id` - (String) Job Id.
* `log_url` - (String) URL to the full agent deployment job logs.
* `status_code` - (String) Final result of the agent deployment job.
  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.
* `status_message` - (String) The outcome of the agent deployment job, in a formatted log string.
* `updated_at` - (String) The agent deploy job updation time.
* `updated_by` - (String) Email address of user who ran the agent deploy job.

## Import

You can import the `ibm_schematics_agent_deploy` resource by using `agent_id`.
The `agent_id` property can be formed from `agent_id`, and `agent_id` in the following format:

```
<agent_id>/<agent_id>
```
* `agent_id`: A string. Agent ID to get the details of agent.
* `agent_id`: A string. Agent ID to get the details of agent.

# Syntax
```
$ terraform import ibm_schematics_agent_deploy.schematics_agent_deploy <agent_id>/<agent_id>
```
