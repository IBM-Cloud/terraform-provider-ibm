---
layout: "ibm"
page_title: "IBM : ibm_schematics_policies"
description: |-
  Get information about schematics_policies
subcategory: "Schematics"
---

# ibm_schematics_policies

Provides a read-only data source for schematics_policies. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_policies" "schematics_policies" {
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `policy_kind` - (Optional, String) Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution.
  * Constraints: Allowable values are: `agent_assignment_policy`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_policies.
* `limit` - (Integer) The number of policy records returned.

* `offset` - (Integer) The skipped number of policy records.

* `policies` - (List) The list of Schematics policies.
Nested scheme for **policies**:
	* `account` - (String) The Account id.
	* `created_at` - (String) The policy creation time.
	* `created_by` - (String) The user who created the Policy.
	* `crn` - (String) The policy CRN.
	* `description` - (String) The description of Schematics customization policy.
	* `id` - (String) The system generated Policy Id.
	* `location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
	  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
	* `name` - (String) The name of Schematics customization policy.
	* `policy_kind` - (String) Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution.
	  * Constraints: Allowable values are: `agent_assignment_policy`.
	* `resource_group` - (String) Resource-group name for the Policy.  By default, Policy will be created in Default Resource Group.
	* `state` - (List) User defined status of the Schematics object.
	Nested scheme for **state**:
		* `set_at` - (String) When the User who set the state of the Object.
		* `set_by` - (String) Name of the User who set the state of the Object.
		* `state` - (String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
		  * Constraints: Allowable values are: `draft`, `live`, `locked`, `disable`.
	* `tags` - (List) Tags for the Schematics customization policy.
	* `updated_at` - (String) The policy updation time.

* `total_count` - (Integer) The total number of policy records.

