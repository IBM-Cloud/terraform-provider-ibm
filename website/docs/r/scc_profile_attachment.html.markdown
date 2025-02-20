---
layout: "ibm"
page_title: "IBM : ibm_scc_profile_attachment"
description: |-
  Manages scc_profile_attachment.
subcategory: "Security and Compliance Center"
---

# ibm_scc_profile_attachment

Create, update, and delete profile attachments with this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

Making a profile attachment using an IBM `ibm_scc_scope`:
```hcl
## Local Variables
locals {
  scc_instance_id               = "f6939361-4f72-47a3-ae5e-0ee77a90ee31"
  ibm_cloud_sample_profile_id   = "623ee808-2fcd-4700-8149-cc5500512ad7"
}

## Datasources

# datasource to obtain information of a profile
data "ibm_scc_profile" "sample_profile_id" {
  instance_id = local.scc_instance_id
  profile_id  = local.ibm_cloud_sample_profile_id
}

## Resources

# resource to create a scope targeting an account
resource "ibm_scc_scope" "scc_personal_account_scope" {
  description = "An scope targeting an account, made using Terraform"
  environment = "ibm-cloud"
  name        = "Terraform sample resource group scope"
  properties  = {
    scope_type  = "account"
    scope_id    = "7379262615a74cb3b9f346408a3e1694"
  }
  instance_id = local.scc_instance_id
}

# resource to create a profile attachment to a predefined profile
resource "ibm_scc_profile_attachment" "cis-profile-attachment-instance" {
  instance_id = local.scc_instance_id
  name        = "tf-demo-profile-attach-demo"
  description = "Sample Profile attachment using Terraform"
  profile_id  = local.ibm_cloud_sample_profile_id

  schedule = "every_7_days"
  status   = "disabled"

  # scope created by the resource ibm_scc_scope
  scope {
    id = ibm_scc_scope.scc_personal_account_scope.scope_id
  }

  # dynamically use the default parameters of a profile if there are any
  dynamic "attachment_parameters" { 
   for_each = data.ibm_scc_profile.sample_profile_id.default_parameters
   content {
      parameter_name         = attachment_parameters.value["parameter_name"]
      parameter_display_name = attachment_parameters.value["parameter_display_name"]
      parameter_type         = attachment_parameters.value["parameter_type"]
      parameter_value        = attachment_parameters.value["parameter_default_value"]
      assessment_type        = attachment_parameters.value["assessment_type"]
      assessment_id          = attachment_parameters.value["assessment_id"]
    }
  }

  notifications {
    enabled = false
    controls {
      failed_control_ids = []
      threshold_limit    = 10
    }
  }
}
```
Making a profile attachment using an IBM `account_id`:

**NOTE** This is considered legacy support and will be deprecated soon.

```hcl
resource "ibm_scc_profile_attachment" "scc_profile_attachment_instance" {
  profile_id = "a0bd1ee2-1ed3-407e-a2f4-ce7a1a38f54d"
  instance_id = "34324315-2edc-23dc-2389-34982389834d"
  name = "profile_attachment_name"
  description = "scc_profile_attachment_description"
  scope {
    environment = "ibm-cloud"	
    properties {
      name = "scope_id"
      value = resource.ibm_scc_control_library.scc_control_library_instance.account_id
    }
    properties {
      name = "scope_type"
      value = "account"
    }
  }
  schedule = "every_30_days"
  status = "enabled"
  notifications {
    enabled = false
    controls {
      failed_control_ids = []
      threshold_limit = 14
    }
  }
  attachment_parameters {
    parameter_value = "22"
    assessment_id = "rule-this-is-a-fake-ruleid"
    parameter_display_name = "Network ACL rule for allowed IPs to SSH port"
    parameter_name = "ssh_port"
    parameter_type = "numeric"
	}
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `profile_id` - (Required, Forces new resource, String) The profile ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `scope` - (List) The scope payload for the multi cloud feature.
  * Constraints: 
    * The maximum length is `8` items. The minimum length is `0` items.
  
  Nested schema for **scope**:
	* `environment` - (String) The environment that relates to this scope.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `properties` - (List) The properties supported for scoping by this environment.
	  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
    Nested schema for **properties**:
      ~> NOTE: Defining the `scope_type` value must be either `account`, `account.resource_group`, `enterprise`, `enterprise.account` and `enterprise.account_group`."
      ~> NOTE: Defining the `scope_id` value will be the id of the `scope_type`(ex. `enterprise.account_group` will be the ID of the account_group within an enterprise)
		* `name` - (Required, String) The name of the property.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `value` - (Required, String) The value of the property.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.;
  * `id` - (Optional, String) The ID of an `ibm_scc_scope` that is prexisiting
    * Constraints: `id` must not be used with `environment` and `properties`
* `notifications` - (Required, List) The configuration for setting up notifications if a scan fails. Requires event_notifications from the instance settings to be setup.

Nested schema for **notifications**:
	* `controls` - (List) The failed controls.
	Nested schema for **controls**:
		* `failed_control_ids` - (List) The failed control IDs.
		  * Constraints: The list items must match regular expression `/^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$|^$/`. The maximum length is `512` items. The minimum length is `0` items.
		* `threshold_limit` - (Integer) The threshold limit.
	* `enabled` - (Boolean) The flag to enable notifications. Set to true to enabled notifications, false to disable
* `attachment_parameters` - (List) The attachment parameters required from the profile that the attachment is targeting. All parameters listed from the profile needs to be set. **NOTE**: All `attachment_parameters` must be defined; use `datasource.ibm_scc_profile` to see all necessary parameters.

Nested schema for **attachment_parameters**:
    * `parameter_name` - (Required, String) The name of the parameter to target.
    * `parameter_display_name` - (Required, String) The display name of the parameter shown in the UI.
    * `parameter_type` - (Required, String) The type of the parameter value.
    * `parameter_value` - (Required, String) The value of the parameter.
    * `assessment_type` - (String) The type of assessment the parameter uses. 
* `schedule` - (String) The schedule of an attachment evaluation.
  * Constraints: Allowable values are: `daily`, `every_7_days`, `every_30_days`.
* `name` - (String) The name of the attachment.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_profile_attachment.
* `profile_attachment_id` - (String) The ID that is associated with the created `profile_attachment`
* `account_id` - (String) The account ID that is associated to the attachment.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `attachment_id` - (String) The ID of the attachment.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `attachment_parameters` - (List) The profile parameters for the attachment.
  * Constraints: The maximum length is `512` items. The minimum length is `0` items.

  Nested schema for **attachment_parameters**:
	* `assessment_id` - (String) The implementation ID of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `assessment_type` - (String) The type of the implementation.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `parameter_display_name` - (String) The parameter display name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
	* `parameter_name` - (String) The parameter name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_]*$/`.
	* `parameter_type` - (String) The parameter type.
	  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`.
	* `parameter_value` - (String) The value of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'"\\s\\-\\[\\]]+$/`.
* `created_by` - (String) The user who created the attachment.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `created_on` - (String) The date when the attachment was created.
* `description` - (String) The description for the attachment.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
* `instance_id` - (String) The instance ID of the account that is associated to the attachment.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$|^$/`.
* `last_scan` - (List) The details of the last scan of an attachment.

  Nested schema for **last_scan**:
	* `id` - (String) The ID of the last scan of an attachment.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
	* `status` - (String) The status of the last scan of an attachment.
	  * Constraints: Allowable values are: `in_progress`, `completed`.
	* `time` - (String) The time when the last scan started.
* `next_scan_time` - (String) The start time of the next scan.
* `status` - (String) The status of an attachment evaluation.
  * Constraints: Allowable values are: `enabled`, `disabled`.
* `updated_by` - (String) The user who updated the attachment.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `updated_on` - (String) The date when the attachment was updated.


## Import

You can import the `ibm_scc_profile_attachment` resource by using `id`.
The `id` property can be formed from `instance_id`, `profiles_id`, and `attachment_id` in the following format:

```bash
<instance_id>/<profile_id>/<attachment_id>
```
* `instance_id`: A string. The instance ID.
* `profile_id`: A string. The profile ID.
* `attachment_id`: A string. The attachment ID.

# Syntax
```bash
$ terraform import ibm_scc_profile_attachment.scc_profile_attachment <instance_id>/<profile_id>/<attachment_id>
```

# Example
```bash
$ terraform import ibm_scc_profile_attachment.scc_profile_attachment 00000000-1111-2222-3333-444444444444/00000000-1111-2222-3333-444444444444/f3517159-889e-4781-819a-89d89b747c85
```
