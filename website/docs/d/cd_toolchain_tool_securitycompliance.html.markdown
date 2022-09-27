---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_securitycompliance"
description: |-
  Get information about cd_toolchain_tool_securitycompliance
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_securitycompliance

~> **Beta:** This data source is in Beta, and is subject to change.

Provides a read-only data source for cd_toolchain_tool_securitycompliance. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.toolchain_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_securitycompliance.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `api_key` - (String) The IBM Cloud API key is used to access the Security and Compliance API. You can obtain your API key with 'ibmcloud iam api-key-create' or via the console at https://cloud.ibm.com/iam#/apikeys by clicking **Create API key** (Each API key only can be viewed once).
	  * Constraints: The value must match regular expression `/\\S/`.
	* `evidence_namespace` - (String) The kind of pipeline evidence to be displayed in Security and Compliance Center for this toolchain. The evidence locker will be searched for CD (Continuous Deployment) pipeline evidence, or for CC (Continuous Compliance) pipeline evidence.
	* `evidence_repo_name` - (String) To collect and store evidence for all tasks performed, a Git repository is required as an evidence locker.
	* `location` - (String)
	* `name` - (String) Give this tool integration a name, for example: my-security-compliance.
	* `profile` - (String) Select an existing profile, where a profile is a collection of security controls. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).
	* `scope` - (String) Select an existing scope name to narrow the focus of the validation scan. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-scopes) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).
	* `trigger_info` - (Map) trigger_info.
	* `trigger_scan` - (String) Enabling trigger validation scans provides details for a pipeline task to trigger a scan.
	  * Constraints: Allowable values are: `disabled`, `enabled`.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing the this resource through an API.
	* `ui_href` - (String) URI representing the this resource through the UI.

* `resource_group_id` - (String) Resource group where tool can be found.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

