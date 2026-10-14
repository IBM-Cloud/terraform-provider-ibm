---
layout: "ibm"
page_title: "IBM : ibm_pdr_last_operation"
description: |-
  Get information about pdr_last_operation
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_last_operation

Retrieves the status of the last operation performed on the specified service instance, such as provisioning, updating, or deprovisioning.

## Example Usage

```hcl
data "ibm_pdr_last_operation" "pdr_last_operation" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
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
* `crn` - (String) The service instance crn.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `deployment_name` - (String) The name of the service instance deployment.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `is_api_key_expired` - (Boolean) Indicates whether the API key used for the deployment is expired.
* `last_updated_orchestrator_deployment_time` - (String) The deployment time of primary orchestrator VM.
  * Constraints: The maximum length is `1048` characters. The minimum length is `20` characters.
* `last_updated_standby_orchestrator_deployment_time` - (String) The deployment time of StandBy orchestrator VM.
  * Constraints: The maximum length is `1048` characters. The minimum length is `20` characters.
* `mfa_enabled` - (String) Indicated whether multi factor authentication is ennabled or not.
  * Constraints: The maximum length is `5` characters. The minimum length is `4` characters. The value must match regular expression `/^(true|false)$/`.
* `orch_ext_connectivity_status` - (String) Status of standby node addition to the orchestrator cluster.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `orch_standby_node_addition_status` - (String) The status of standby node in the Orchestrator cluster.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `orchestrator_cluster_message` - (String) The current status of the primary orchestrator VM.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `orchestrator_config_status` - (String) The configuration status of the orchestrator cluster.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `orchestrator_ha` - (Boolean) Indicates whether high availability (HA) is enabled for the orchestrator.
* `plan_name` - (String) The name of the DR Automation plan.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_ ]+$/`.
* `primary_description` - (String) Indicates the progress details of primary orchestrator creation.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `primary_error_description` - (String) Capture the error while creating primary orchestrator.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `primary_ip_address` - (String) The IP address of the primary orchestrator VM.
  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
* `primary_orchestrator_status` - (String) The configuration status of the orchestrator cluster.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `recovery_location` - (String) The disaster recovery location associated with the instance.
  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `resource_group` - (String) The resource group to which the service instance belongs.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `standby_description` - (String) Indicates the progress details of primary orchestrator creation.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `standby_error_description` - (String) Capture the error while creating standby orchestrator.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `standby_ip_address` - (String) The IP address of the standby orchestrator VM.
  * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\\-:]+$/`.
* `standby_status` - (String) The current state of the standby orchestrator.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `status` - (String) The current state of the primary orchestrator.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.

