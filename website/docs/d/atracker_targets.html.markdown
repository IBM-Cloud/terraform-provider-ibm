---
layout: "ibm"
page_title: "IBM : ibm_atracker_targets"
description: |-
  Get information about atracker_targets
subcategory: "Activity Tracker Event Routing"
---

# ibm_atracker_targets

Provides a read-only data source to retrieve information about atracker_targets. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_atracker_targets" "atracker_targets" {
	name = "a-cos-target-us-south"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `name` - (String) The name of the target resource.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the atracker_targets.
* `targets` - (List) A list of target resources.
Nested scheme for **targets**:
	* `id` - (String) The uuid of the target resource.
	* `name` - (String) The name of the target resource.
	* `crn` - (String) The crn of the target resource.
	* `target_type` - (String) The type of the target.
	  * Constraints: Allowable values are: `cloud_object_storage`, `logdna`, `event_streams`, `cloud_logs`.
	* `encrypt_key` - (String) The encryption key that is used to encrypt events before Activity Tracker services buffer them on storage. This credential is masked in the response.
	* `region` - (String) Included this optional field if you used it to create a target in a different region other than the one you are connected.
	* `cos_endpoint` - (List) Property values for a Cloud Object Storage Endpoint.
	Nested scheme for **cos_endpoint**:
		* `api_key` - (String) The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response. This is required if service_to_service is not enabled.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `bucket` - (String) The bucket name under the Cloud Object Storage instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `endpoint` - (String) The host name of the Cloud Object Storage endpoint.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `service_to_service_enabled` - (Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
		* `target_crn` - (String) The CRN of the Cloud Object Storage instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `logdna_endpoint` - (List) Property values for a LogDNA Endpoint.
	Nested scheme for **logdna_endpoint**:
		* `ingestion_key` - (String) The LogDNA ingestion key is used for routing logs to a specific LogDNA instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `target_crn` - (String) The CRN of the LogDNA instance.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `eventstreams_endpoint` - (List) Property values for Event streams Endpoint.
	Nested scheme for **eventstreams_endpoint**:
		* `api_key` - (String) The user password (api key) for the message hub topic in the Event Streams instance. This is required if service_to_service is not enabled..
			* Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `topic` - (String) The topic name defined under the Event streams instance.
			* Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
		* `brokers` - (List) The list of brokers defined under the Event streams instance and used in the event streams endpoint.
			* Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `service_to_service_enabled` - (Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
		* `target_crn` - (String) The CRN of the Event streams instance.
			* Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `cloudlogs_endpoint` - (List) Property values for an IBM Cloud Logs endpoint.
	Nested scheme for **cloudlogs_endpoint**:
		* `target_crn` - (String) The CRN of the IBM Cloud Logs instance.
			* Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
  * `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
	Nested scheme for **write_status**:
  	* `last_failure` - (String) The timestamp of the failure.
  	* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
  	* `status` - (String) The status such as failed or success.
	* `created_at` - (String) The timestamp of the target creation time.
	* `updated_at` - (String) The timestamp of the target last updated time.
	* `api_version` - (Integer) The API version of the target.
  * `cos_write_status` - **DEPRECATED** (List) The status of the write attempt with the provided cos_endpoint parameters.
	Nested scheme for **cos_write_status**:
		* `status` - (String) The status such as failed or success.
		* `last_failure` - (String) The timestamp of the failure.
		* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
	* `created` - **DEPRECATED** (String) The timestamp of the target creation time.
	* `updated` - **DEPRECATED** (String) The timestamp of the target last updated time.
