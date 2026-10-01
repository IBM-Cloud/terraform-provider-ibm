---
layout: "ibm"
page_title: "IBM : ibm_pdr_powervs_workspace"
description: |-
  Get information about pdr_powervs_workspace
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_powervs_workspace

Retrieves the power virtual server workspaces for primary and standby orchestrator based on location id.

## Example Usage

```hcl
data "ibm_pdr_powervs_workspace" "pdr_powervs_workspace" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
	location_id = "us-south"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `dr_standby_workspace_description` - (String) Description of Standby Workspace.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `dr_standby_workspaces` - (List) The list of standby disaster recovery workspaces.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

Nested schema for **dr_standby_workspaces**:
	* `details` - (Object) Details of the DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
		  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-iaas:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
	* `id` - (String) The unique identifier of the standby workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	
	* `location` - (List) Represents a disaster recovery location.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
		* `url` - (String) The URL endpoint to access the DR location.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `name` - (String) The name of the standby workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `status` - (String) The status of the standby workspace.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.    
* `dr_workspace_description` - (String) Description of Workspace.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `dr_workspaces` - (List) The list of primary disaster recovery workspaces.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

Nested schema for **dr_workspaces**:
	* `default` - (Boolean) Indicates if this is the default DR workspace.
	* `details` - (Object) Details of the DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
		  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-iaas:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
	* `id` - (String) The unique identifier of the DR workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `location` - (List) Represents a disaster recovery location.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
		* `url` - (String) The URL endpoint to access the DR location.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
	* `name` - (String) The name of the DR workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `status` - (String) The status of the DR workspace.
	  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

