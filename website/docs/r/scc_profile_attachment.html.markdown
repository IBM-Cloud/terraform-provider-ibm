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

Making a profile attachment using an IBM `account_id`. **NOTE** This usage will create a scope that cannot be tracked by terraform:
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

Full example of creating a profile attachment using `ibm_scc_scope` with an enterprise scope and fact_provider scope to target a SOC2 profile:
```hcl
# resource to make a provider_type_instance for the scope
resource "ibm_scc_provider_type_instance" "facts_provider_type_instance" {
  name             = "ibm_facts_engine"
  instance_id      = "f740da22-29e0-47fc-b768-73f0d5bb2955" #id of the provider_type_instance
  provider_type_id = "19d0e3c118e9ccb6fa66abb26bfb1e60"     #id of the provider_type
  attributes       = {"source": "scc"}
}

# resource to make an enterprise scope
resource "ibm_scc_scope" "scc_enterprise_scope" {
  description = "This scope targets an IBM account"
  environment = "ibm-cloud"
  instance_id = "f740da22-29e0-47fc-b768-73f0d5bb2955"
  name        = "Sample account Scope done from Terraform"
  properties {
    enterprise_id = "fbf5d5bd0698cef3ab3b488559e08992"
  }
}

# resource to make a facts_provider_type_instance
resource "ibm_scc_scope" "scc_ibm_facts_scope" {
  description = "This scope targets a facts provider type instance"
  environment = "ibm-cloud"
  instance_id = "f740da22-29e0-47fc-b768-73f0d5bb2955"
  name        = "Sample facts Scope done from Terraform"
  properties {
    ibm_facts_api_instance_id = ibm_scc_provider_type_instance.facts_provider_type_instance.provider_type_instance_id
  }
}

# resource to create a profile attachment
resource "ibm_scc_profile_attachment" "scc-profile-attachment-instance" {
  instance_id = "f740da22-29e0-47fc-b768-73f0d5bb2955"
  name        = "tf-demo-ibm-SOC2-demo"
  description = "This attachment targets the SCC IBM Cloud SOC2 profile using Terraform"
  profile_id  = "9487feb1-f1d0-4ca4-86ef-3680368b284a"

  schedule = "every_7_days"
  status   = "enabled"

  # scope ID of the ibm_facts_engine
  scope {
    id = ibm_scc_scope.scc_ibm_facts_scope.scope_id 
  }

  # scope ID of the ibm account
  scope {
    id = ibm_scc_scope.scc_enterprise_scope.scope_id
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-64b9c630-0308-43d0-a7bd-113a7baa1059"
    parameter_name = "server_asset_excluded_role_list"
    parameter_value = "['Ingress', 'Ingress IP']"
    parameter_display_name = "Server asset excluded roles - Example: [\"server\",\"ingress\"]"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-933c1885-1666-4f0d-b9f8-98ea151bad60"
    parameter_name = "qradar_last_detected_on"
    parameter_value = "3"
    parameter_display_name = "Logs last detected within N days"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-933c1885-1666-4f0d-b9f8-98ea151bad60"
    parameter_name = "qradar_last_detected_on_for_warn"
    parameter_value = "7"
    parameter_display_name = "Logs last detected within N days for warning status"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-c6b6b0de-36c6-49cc-93d4-e0970d43aeb5"
    parameter_name = "pce_interval_limit"
    parameter_value = "20"
    parameter_display_name = "PCE creation interval check limit"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-12070ae2-8843-4290-94a6-34f57a84911b"
    parameter_name = "due_date_interval"
    parameter_value = "30"
    parameter_display_name = "Due Date interval Limit"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-db92a4b5-edad-4731-a0c7-94fa65e7ead1"
    parameter_name = "server_asset_excluded_role_list"
    parameter_value = "['ingress', 'ingress ip']"
    parameter_display_name = "Server asset excluded roles - Example: [\"server\",\"ingress\"]"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-37513a2c-f9fb-4c42-9c9f-1d5f4f2d20a2"
    parameter_name = "server_asset_excluded_role_list"
    parameter_value = "['ingress', 'ingress ip']"
    parameter_display_name = "Server asset excluded roles - Example: [\"server\",\"ingress\"]"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-37513a2c-f9fb-4c42-9c9f-1d5f4f2d20a2"
    parameter_name = "patch-due-date-start-interval"
    parameter_value = "0"
    parameter_display_name = "Days in past from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-37513a2c-f9fb-4c42-9c9f-1d5f4f2d20a2"
    parameter_name = "patch-due-date-end-interval"
    parameter_value = "7"
    parameter_display_name = "Days in future from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-3ed5a784-cfd6-4cf2-8efd-7a19d52adcaa"
    parameter_name = "close_category"
    parameter_value = "['successful']"
    parameter_display_name = "Close code for the change"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-3ed5a784-cfd6-4cf2-8efd-7a19d52adcaa"
    parameter_name = "service_environment"
    parameter_value = "['production']"
    parameter_display_name = "Choice of environment"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-de3ebf60-8b94-4d09-8e60-98aefcd747fa"
    parameter_name = "vmt-platform-high-timeline"
    parameter_value = "15"
    parameter_display_name = "Timeline for high severity vulnerability"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-de3ebf60-8b94-4d09-8e60-98aefcd747fa"
    parameter_name = "vmt-platform-due-date-start-interval"
    parameter_value = "0"
    parameter_display_name = "Days in past from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-de3ebf60-8b94-4d09-8e60-98aefcd747fa"
    parameter_name = "vmt-platform-due-date-end-interval"
    parameter_value = "7"
    parameter_display_name = "Days in future from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-c3b77405-bd19-4b1e-a374-722200ac77fc"
    parameter_name = "vmt-platform-medium-timeline"
    parameter_value = "45"
    parameter_display_name = "Timeline for medium severity vulnerability"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-c3b77405-bd19-4b1e-a374-722200ac77fc"
    parameter_name = "remote-vmt-due-date-start-interval"
    parameter_value = "0"
    parameter_display_name = "Days in past from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-c3b77405-bd19-4b1e-a374-722200ac77fc"
    parameter_name = "remote-vmt-due-date-end-interval"
    parameter_value = "7"
    parameter_display_name = "Days in future from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-19d4a5e0-6727-40ad-8a41-13ea522bc83d"
    parameter_name = "vmt-platform-low-timeline"
    parameter_value = "45"
    parameter_display_name = "Timeline for low severity vulnerability"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-19d4a5e0-6727-40ad-8a41-13ea522bc83d"
    parameter_name = "vmt-due-date-start-interval"
    parameter_value = "0"
    parameter_display_name = "Days in past from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-19d4a5e0-6727-40ad-8a41-13ea522bc83d"
    parameter_name = "vmt-due-date-end-interval"
    parameter_value = "7"
    parameter_display_name = "Days in future from current day"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-76d793eb-abef-44ca-be92-c7b8754226d0"
    parameter_name = "vmt-platform-network-scan-frequency"
    parameter_value = "7"
    parameter_display_name = "Scan frequency for platform network scan"
    parameter_type = "numeric"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-7cacfc49-46a3-4aca-9c21-263e0ad40e94"
    parameter_name = "server_asset_excluded_role_list"
    parameter_value = "['ingress', 'ingress ip']"
    parameter_display_name = "Server asset excluded roles - Example: [\"server\",\"ingress\"]"
    parameter_type = "string_list"
  }

  attachment_parameters {
    assessment_type = "automated"
    assessment_id = "rule-903adad2-1989-4ce8-893e-dd43a56ab008"
    parameter_name = "backup_interval"
    parameter_value = "7"
    parameter_display_name = "Expected backup interval to check in days"
    parameter_type = "numeric"
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

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `profile_id` - (Required, Forces new resource, String) The profile ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `scope` - (List) The scope payload for the multi cloud feature.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
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
