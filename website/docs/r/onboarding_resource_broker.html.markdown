---
layout: "ibm"
page_title: "IBM : ibm_onboarding_resource_broker"
description: |-
  Manages onboarding_resource_broker.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_resource_broker

**Note - Intended for internal use only. This resource is strictly experimental and subject to change without notice.**

Create, update, and delete onboarding_resource_brokers with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_resource_broker" "onboarding_resource_broker_instance" {
  allow_context_updates = false
  auth_scheme = "bearer"
  auth_username = "apikey"
  broker_url = "https://broker-url-for-my-service.com"
  catalog_type = "service"
  name = "brokername"
  region = "global"
  resource_group_crn = "crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f"
  state = "active"
  type = "provision_through"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `allow_context_updates` - (Optional, Boolean) Whether the resource controller will call the broker for any context changes to the instance. Currently, the only context related change is an instance name update.
* `auth_password` - (Optional, String) The authentication password to reach the broker.
* `auth_scheme` - (Required, String) The supported authentication scheme for the broker.
  * Constraints: Allowable values are: `bearer`, `bearer-crn`.
* `auth_username` - (Optional, String) The authentication username to reach the broker.
  * Constraints: Allowable values are: `apikey`.
* `broker_url` - (Required, String) The URL associated with the broker application.
* `catalog_type` - (Optional, String) To enable the provisioning of your broker, set this parameter value to `service`.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_.-]+$/`.
* `name` - (Required, String) The name of the broker.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
* `region` - (Optional, String) The region where the pricing plan is available.
* `resource_group_crn` - (Optional, String) The cloud resource name of the resource group.
* `state` - (Optional, String) The state of the broker.
  * Constraints: Allowable values are: `active`, `removed`.
* `type` - (Required, String) The type of the provisioning model.
  * Constraints: Allowable values are: `provision_through`, `provision_behind`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_resource_broker.
* `account_id` - (String) The ID of the account in which you manage the broker.
* `created_at` - (String) The time when the service broker was created.
* `created_by` - (List) The details of the user who created this broker.
Nested schema for **created_by**:
	* `user_id` - (String) The ID of the user who dispatched this action.
	* `user_name` - (String) The username of the user who dispatched this action.
* `crn` - (String) The cloud resource name (CRN) of the broker.
* `deleted_at` - (String) The time when the service broker was deleted.
* `deleted_by` - (List) The details of the user who deleted this broker.
Nested schema for **deleted_by**:
	* `user_id` - (String) The ID of the user who dispatched this action.
	* `user_name` - (String) The username of the user who dispatched this action.
* `guid` - (String) The globally unique identifier of the broker.
* `updated_at` - (String) The time when the service broker was updated.
* `updated_by` - (List) The details of the user who updated this broker.
Nested schema for **updated_by**:
	* `user_id` - (String) The ID of the user who dispatched this action.
	* `user_name` - (String) The username of the user who dispatched this action.
* `url` - (String) The URL associated with the broker.


## Import

You can import the `ibm_onboarding_resource_broker` resource by using `id`. The identifier of the broker.

# Syntax
<pre>
$ terraform import ibm_onboarding_resource_broker.onboarding_resource_broker id;
</pre>
