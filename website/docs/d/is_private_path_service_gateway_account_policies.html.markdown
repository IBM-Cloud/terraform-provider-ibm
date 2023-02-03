---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_account_policies"
description: |-
  Get information about PrivatePathServiceGatewayAccountPolicyCollection
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_account_policies

Provides a read-only data source for PrivatePathServiceGatewayAccountPolicyCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_private_path_service_gateway_account_policies" "is_private_path_service_gateway_account_policies" {
	private_path_service_gateway_id = ibm_is_private_path_service_gateway_account_policy.is_private_path_service_gateway_account_policy.private_path_service_gateway_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `private_path_service_gateway_id` - (Required, Forces new resource, String) The private path service gateway identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the PrivatePathServiceGatewayAccountPolicyCollection.
* `account_policies` - (List) Collection of account policies.
Nested scheme for **account_policies**:
	* `access_policy` - (String) The access policy for the account:- permit: access will be permitted- deny:  access will be denied- review: access will be manually reviewedThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  * Constraints: Allowable values are: `deny`, `permit`, `review`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `account` - (List) The account for this access policy.
	Nested scheme for **account**:
		* `id` - (String)
		  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
		* `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `account`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `created_at` - (String) The date and time that the account policy was created.
	* `href` - (String) The URL for this account policy.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this account policy.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `private_path_service_gateway_account_policy`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `updated_at` - (String) The date and time that the account policy was updated.

* `first` - (List) A link to the first page of resources.
Nested scheme for **first**:
	* `href` - (String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `limit` - (Integer) The maximum number of resources that can be returned by the request.
  * Constraints: The maximum value is `100`. The minimum value is `1`.

* `next` - (List) A link to the next page of resources. This property is present for all pagesexcept the last page.
Nested scheme for **next**:
	* `href` - (String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `total_count` - (Integer) The total number of resources across all pages.
  * Constraints: The minimum value is `0`.

