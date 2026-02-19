---
layout: "ibm"
page_title: "IBM : ibm_is_instance_software_attachment"
description: |-
  Manages InstanceSoftwareAttachment.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_instance_software_attachment

Create, update, and delete InstanceSoftwareAttachments with this resource.

## Example Usage

```hcl
resource "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
  instance_id = "instance_id"
  name = "my-software-attachment"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The virtual server instance identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `name` - (Optional, String) The name for this instance software attachment. The name is unique across all instance software attachments for the instance.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the InstanceSoftwareAttachment.
* `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this instance software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.
Nested schema for **catalog_offering**:
	* `plan` - (List) The billing plan for the catalog offering version associated with this instance softwareattachment.If absent, no billing plan is associated with the catalog offering version (free).
	Nested schema for **plan**:
		* `crn` - (String) The CRN for this[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering version's billing plan.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `version` - (List) The catalog offering version associated with this instance software attachment.
	Nested schema for **version**:
		* `crn` - (String) The CRN for this version of a[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
* `created_at` - (String) The date and time that the instance software attachment was created.
* `entitlement` - (List) The entitlement for the licensed software for this instance software attachment.
Nested schema for **entitlement**:
	* `licensed_software` - (List) The licensed software for this instance software attachment entitlement.
	Nested schema for **licensed_software**:
		* `sku` - (String) The SKU for this licensed software.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~]+$/`.
* `href` - (String) The URL for this instance software attachment.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_instance_software_attachment_id` - (String) The unique identifier for this instance software attachment.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_reasons` - (List) The lifecycle reasons for this instance software attachment (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `failed_registration`: the software instance's registration to Resource Controller,  which includes creation of any required software license(s), has failed. Delete the  instance and provision it again. If the problem persists, contact IBM Support.- `internal_error`: internal error (contact IBM support)- `pending_registration`: the software instance's registration to Resource Controller,  and the creation of any required software license(s), is being processed.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `failed_registration`, `internal_error`, `pending_registration`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the instance software attachment.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `offering_instance` - (List) 
Nested schema for **offering_instance**:
	* `crn` - (String) The CRN for the software offering instance registered with Resource Controller that is associated with the instance software attachment.
	  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-_]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `instance_software_attachment`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.


## Import

You can import the `ibm_is_instance_software_attachment` resource by using `id`.
The `id` property can be formed from `instance_id`, and `is_instance_software_attachment_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;is_instance_software_attachment_id&gt;
</pre>
* `instance_id`: A string. The virtual server instance identifier.
* `is_instance_software_attachment_id`: A string in the format `0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e`. The unique identifier for this instance software attachment.

# Syntax
<pre>
$ terraform import ibm_is_instance_software_attachment.is_instance_software_attachment &lt;instance_id&gt;/&lt;is_instance_software_attachment_id&gt;
</pre>
