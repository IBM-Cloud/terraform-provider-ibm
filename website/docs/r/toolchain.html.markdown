---
layout: "ibm"
page_title: "IBM : ibm_toolchain"
description: |-
  Manages toolchain.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain

Provides a resource for toolchain. This allows toolchain to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain" "toolchain" {
  description = "A sample toolchain to test the API"
  generator = "API"
  key = "somekey5"
  name = "TestToolchain3"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `container` - (Optional, Forces new resource, List) 
Nested scheme for **container**:
	* `guid` - (Required, String)
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `organization_guid`, `resource_group_id`.
* `description` - (Optional, String) Describes the toolchain.
  * Constraints: The maximum length is `500` characters.
* `generator` - (Required, Forces new resource, String) A description of who generated the toolchain.
  * Constraints: Allowable values are: `API`, `IBM Bluemix DevOps Services`, `IBM Cloud DevOps Services`, `Bluemix`, `IBM Cloud`, `otc_service`.
* `key` - (Optional, Forces new resource, String) <strong>Deprecated: </strong><br><br>Key of this toolchain, can be used when querying for toolchain.
  * Constraints: The maximum length is `200` characters.
* `name` - (Required, String) Toolchain name.
  * Constraints: The maximum length is `128` characters.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain.
* `creator` - (Required, String) 
* `crn` - (Optional, String) CRN for resource group based toolchains.
* `status` - (Optional, List) The status of the toolchain.
Nested scheme for **status**:
	* `detailed_status` - (Required, List) A list of particular status issues related to the toolchain.
	  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
	Nested scheme for **detailed_status**:
		* `details` - (Optional, String) A longer description of the problem, typically message/details would be displayed together in a UI to give the user an understanding of the issue. Should always be included when the 'status' is not 'ok'.
		  * Constraints: The maximum length is `700` characters.
		* `message` - (Optional, String) A short message indicating the problem. Should always be included when the 'status' is not 'ok'.
		  * Constraints: The maximum length is `500` characters.
		* `status` - (Required, String) The status of the particular issue'ok' indicates a normal state, additional messages are possible but unlikely'warning' indicates a possible problem with the toolchain, but usage may continue'error' indicates furthur use of this toolchain would be problematic, user actions should be blocked.
		  * Constraints: Allowable values are: `ok`, `warning`, `error`.
		* `status_line` - (Optional, String) A short status that can be used to mark up a toolchain card or other location. Should always be included when the 'status' is not 'ok'.
		  * Constraints: The maximum length is `300` characters.
* `tags` - (Required, List) 

## Import

You can import the `ibm_toolchain` resource by using `toolchain_guid`. The unique identifier of the toolchain.

# Syntax
```
$ terraform import ibm_toolchain.toolchain <toolchain_guid>
```

# Example
```
$ terraform import ibm_toolchain.toolchain 4ef49b76-768f-4990-a208-7e18d5960af2
```
