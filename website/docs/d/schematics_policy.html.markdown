---
layout: "ibm"
page_title: "IBM : ibm_schematics_policy"
description: |-
  Get information about schematics_policy
subcategory: "Schematics"
---

# ibm_schematics_policy

Provides a read-only data source for schematics_policy. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_policy" "schematics_policy" {
	policy_id = "policy_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `policy_id` - (Required, Forces new resource, String) ID to get the details of policy.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_policy.
* `account` - (String) The Account id.

* `created_at` - (String) The policy creation time.

* `created_by` - (String) The user who created the policy.

* `crn` - (String) The policy CRN.

* `description` - (String) The description of Schematics customization policy.

* `id` - (String) The system generated policy Id.

* `kind` - (String) Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution.
  * Constraints: Allowable values are: `agent_assignment_policy`.

* `location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.

* `name` - (String) Name of Schematics customization policy.

* `parameter` - (List) The parameter to tune the Schematics policy.
Nested scheme for **parameter**:
	* `agent_assignment_policy_parameter` - (List) Parameters for the `agent_assignment_policy`.
	Nested scheme for **agent_assignment_policy_parameter**:
		* `selector_ids` - (List) The static selectors of schematics object ids (workspace, action or blueprint) for the Schematics policy.
		* `selector_kind` - (String) Types of schematics object selector.
		  * Constraints: Allowable values are: `ids`, `scoped`.
		* `selector_scope` - (List) The selectors to dynamically list of schematics object ids (workspace, action or blueprint) for the Schematics policy.
		Nested scheme for **selector_scope**:
			* `kind` - (String) Name of the Schematics automation resource.
			  * Constraints: Allowable values are: `workspace`, `action`, `system`, `environment`, `blueprint`.
			* `locations` - (List) The location based selector.
			  * Constraints: Allowable list items are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
			* `resource_groups` - (List) The resource group based selector.
			* `tags` - (List) The tag based selector.

* `resource_group` - (String) The resource group name for the policy.  By default, Policy will be created in `default` Resource Group.

* `scoped_resources` - (List) List of scoped Schematics resources targeted by the policy.
Nested scheme for **scoped_resources**:
	* `id` - (String) Schematics resource Id.
	* `kind` - (String) Name of the Schematics automation resource.
	  * Constraints: Allowable values are: `workspace`, `action`, `system`, `environment`, `blueprint`.

* `state` - (List) User defined status of the Schematics object.
Nested scheme for **state**:
	* `set_at` - (String) When the User who set the state of the Object.
	* `set_by` - (String) Name of the User who set the state of the Object.
	* `state` - (String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
	  * Constraints: Allowable values are: `draft`, `live`, `locked`, `disable`.

* `tags` - (List) Tags for the Schematics customization policy.

* `target` - (List) The objects for the Schematics policy.
Nested scheme for **target**:
	* `selector_ids` - (List) Static selectors of schematics object ids (agent, workspace, action or blueprint) for the Schematics policy.
	* `selector_kind` - (String) Types of schematics object selector.
	  * Constraints: Allowable values are: `ids`, `scoped`.
	* `selector_scope` - (List) Selectors to dynamically list of schematics object ids (agent, workspace, action or blueprint) for the Schematics policy.
	Nested scheme for **selector_scope**:
		* `kind` - (String) Name of the Schematics automation resource.
		  * Constraints: Allowable values are: `workspace`, `action`, `system`, `environment`, `blueprint`.
		* `locations` - (List) The location based selector.
		  * Constraints: Allowable list items are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
		* `resource_groups` - (List) The resource group based selector.
		* `tags` - (List) The tag based selector.

* `updated_at` - (String) The policy updation time.

