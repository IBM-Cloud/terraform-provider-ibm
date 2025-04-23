---
layout: "ibm"
page_title: "IBM : ibm_atracker_targets"
description: |-
  Get information about atracker_targets
subcategory: "Activity Tracker Event Routing"
---

# ibm_atracker_targets

Provides a read-only data source to retrieve information about atracker_targets. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_atracker_targets" "atracker_targets" {
	name = "a-cos-target-us-south"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the target resource.
* `region` - (Optional, String) Limit the query to the specified region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the atracker_targets.
* `targets` - (List) A list of target resources.
Nested schema for **targets**:
	* `api_version` - (Integer) The API version of the target.
	  * Constraints: The maximum value is `2`. The minimum value is `2`.
	* `cloudlogs_endpoint` - (List) Property values for the IBM Cloud Logs endpoint in responses.
	Nested schema for **cloudlogs_endpoint**:
		* `target_crn` - (String) The CRN of the IBM Cloud Logs instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `cos_endpoint` - (List) Property values for a Cloud Object Storage Endpoint in responses.
	Nested schema for **cos_endpoint**:
		* `bucket` - (String) The bucket name under the Cloud Object Storage instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `endpoint` - (String) The host name of the Cloud Object Storage endpoint.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `service_to_service_enabled` - (Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
		* `target_crn` - (String) The CRN of the Cloud Object Storage instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `created_at` - (String) The timestamp of the target creation time.
	* `crn` - (String) The crn of the target resource.
	* `eventstreams_endpoint` - (List) Property values for the Event Streams Endpoint in responses.
	Nested schema for **eventstreams_endpoint**:
		* `api_key` - (String) The user password (api key) for the message hub topic in the Event Streams instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `brokers` - (List) List of broker endpoints.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `service_to_service_enabled` - (Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
		* `target_crn` - (String) The CRN of the Event Streams instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `topic` - (String) The messsage hub topic defined in the Event Streams instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `id` - (String) The uuid of the target resource.
	* `message` - (String) An optional message containing information about the target.
	* `name` - (String) The name of the target resource.
	* `region` - (String) Included this optional field if you used it to create a target in a different region other than the one you are connected.
	* `target_type` - (String) The type of the target.
	  * Constraints: Allowable values are: `cloud_object_storage`, `event_streams`, `cloud_logs`.
	* `updated_at` - (String) The timestamp of the target last updated time.
	* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
	Nested schema for **write_status**:
		* `last_failure` - (String) The timestamp of the failure.
		* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
		* `status` - (String) The status such as failed or success.
