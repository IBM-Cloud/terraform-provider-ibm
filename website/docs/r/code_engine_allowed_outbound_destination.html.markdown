---
layout: "ibm"
page_title: "IBM : ibm_code_engine_allowed_outbound_destination"
description: |-
  Manages code_engine_allowed_outbound_destination.
subcategory: "Code Engine"
---

# ibm_code_engine_allowed_outbound_destination

Create, update, and delete code_engine_allowed_outbound_destinations with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
  project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
  type       = "cidr_block"
  name       = "my-cidr-block-1"
  cidr_block = "192.68.3.0/24"
}
```

```hcl
resource "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
  project_id                       = data.ibm_code_engine_project.code_engine_project_instance.project_id
  type                             = "private_path_service_gateway"
  name                             = "my-private-path-service-gateway"
  private_path_service_gateway_crn = "crn:v1:bluemix:public:is:eu-de:a/4329073d16d2f3663f74bfa955259139::private-path-service-gateway:r010-a78455bd-ff11-4d10-8a4f-b664b15079a4"
  isolation_policy                 = "dedicated"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cidr_block` - (Optional, String) The IPv4 address range.
  * Constraints: The maximum length is `18` characters. The minimum length is `9` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.
* `isolation_policy` - (Optional, String) Optional property to specify the isolation policy of the private path service gateway. If set to `shared`, other projects within the same account or enterprise account family can connect to Private Path service, too. If set to `dedicated` the gateway can only be used by a single Code Engine project. If not specified the isolation policy will be set to `shared`.
  * Constraints: The default value is `shared`. Allowable values are: `shared`, `dedicated`.
* `name` - (Required, Forces new resource, String) The name of the allowed outbound destination.
* `private_path_service_gateway_crn` - (Optional, Forces new resource, String) The CRN of the Private Path service.
  * Constraints: The maximum length is `253` characters. The minimum length is `20` characters. The value must match regular expression `/^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:is\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:\\:private-path-service-gateway\\:[\\-a-zA-Z0-9\/.]*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `type` - (Required, Forces new resource, String) Specify the type of the allowed outbound destination. Allowed types are: `cidr_block` and `private_path_service_gateway`.
  * Constraints: The default value is `cidr_block`. Allowable values are: `cidr_block`, `private_path_service_gateway`. The value must match regular expression `/^(cidr_block|private_path_service_gateway)$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_allowed_outbound_destination.
* `entity_tag` - (String) The version of the allowed outbound destination, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `status` - (String) The current status of the outbound destination.
  * Constraints: Allowable values are: `ready`, `failed`, `deploying`.
* `status_details` - (List) 
Nested schema for **status_details**:
	* `endpoint_gateway` - (List) Optional information about the endpoint gateway located in the Code Engine VPC that connects to the private path service gateway.
	Nested schema for **endpoint_gateway**:
		* `account_id` - (String) The account that created the endpoint gateway.
		* `created_at` - (String) The timestamp when the endpoint gateway was created.
		* `ips` - (List) The reserved IPs bound to this endpoint gateway.
		* `name` - (String) The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.
	* `private_path_service_gateway` - (List) Optional information about the private path service gateway that this allowed outbound destination points to.
	Nested schema for **private_path_service_gateway**:
		* `id` - (String) The private path service gateway identifier.
		* `name` - (String) The name of private path service gateway.
		* `service_endpoints` - (List) The fully qualified domain names for this private path service gateway. The domains are used for endpoint gateways to connect to the service and are configured in the VPC for each endpoint gateway.
	* `reason` - (String) Optional information to provide more context in case of a 'failed' or 'deploying' status.
	  * Constraints: Allowable values are: `ready`, `private_path_crn_invalid`, `private_path_not_in_same_region`, `private_path_not_in_same_account_family`, `private_path_not_found`, `private_path_not_published`, `private_path_connection_already_exists`, `private_path_connection_approval_denied`, `private_path_connection_approval_pending`, `deploying`, `failed`.

* `etag` - ETag identifier for code_engine_allowed_outbound_destination.

## Import

You can import the `ibm_code_engine_allowed_outbound_destination` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `allow-all`. The name of the allowed outbound destination.

# Syntax
<pre>
$ terraform import ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination &lt;project_id&gt;/&lt;name&gt;
</pre>
