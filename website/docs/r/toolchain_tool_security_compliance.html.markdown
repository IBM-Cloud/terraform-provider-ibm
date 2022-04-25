---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_security_compliance"
description: |-
  Manages toolchain_tool_security_compliance.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_security_compliance

Provides a resource for toolchain_tool_security_compliance. This allows toolchain_tool_security_compliance to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_security_compliance" "toolchain_tool_security_compliance" {
  parameters {
		name = "name"
		evidence_repo_name = "evidence_repo_name"
		trigger_scan = "disabled"
		location = "IBM Cloud"
		api-key = "api-key"
		scope = "scope"
		profile = "profile"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Arbitrary JSON data.
Nested scheme for **parameters**:
	* `api_key` - (Optional, String) The IBM Cloud API key is used to access the Security and Compliance API. You can obtain your API key with 'ibmcloud iam api-key-create' or via the console at https://cloud.ibm.com/iam#/apikeys by clicking **Create API key** (Each API key only can be viewed once).
	  * Constraints: The value must match regular expression `/\\S/`.
	* `evidence_repo_name` - (Required, String) To collect and store evidence for all tasks performed, a Git repository is required as an evidence locker.
	* `location` - (Optional, String)
	  * Constraints: Allowable values are: `IBM Cloud`, `IBM Cloud (Staging)`, `IBM Cloud (Development)`.
	* `name` - (Required, String) Give this tool integration a name, for example: my-security-compliance.
	* `profile` - (Optional, String) Select an existing profile, where a profile is a collection of security controls. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
	* `scope` - (Optional, String) Select an existing scope name to narrow the focus of the validation scan. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-scopes).
	* `trigger_scan` - (Optional, String) Enabling trigger validation scans provides details for a pipeline task to trigger a scan.
	  * Constraints: Allowable values are: `disabled`, `enabled`.
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_security_compliance.
* `crn` - (Required, String) 
* `get_integration_by_id_response_id` - (Required, String) 
* `href` - (Required, String) 
* `referent` - (Required, List) 
Nested scheme for **referent**:
	* `api_href` - (Optional, String)
	* `ui_href` - (Optional, String)
* `resource_group_id` - (Required, String) 
* `state` - (Required, String) 
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (Required, String) 
* `updated_at` - (Required, String) 

## Import

You can import the `ibm_toolchain_tool_security_compliance` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_security_compliance.toolchain_tool_security_compliance <toolchain_id>/<integration_id>
```
