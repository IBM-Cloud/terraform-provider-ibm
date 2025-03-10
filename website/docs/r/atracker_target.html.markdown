---
layout: "ibm"
page_title: "IBM : ibm_atracker_target"
description: |-
  Manages atracker_target.
subcategory: "Activity Tracker API Version 2"
---

# ibm_atracker_target

Create, update, and delete atracker_targets with this resource.

## Example Usage

```hcl
resource "ibm_atracker_target" "atracker_target_instance" {
  cloudlogs_endpoint {
		target_crn = "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
  }
  cos_endpoint {
		endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
		target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		bucket = "my-atracker-bucket"
		service_to_service_enabled = true
  }
  name = "my-cos-target"
  target_type = "cloud_object_storage"
  region = "us-south"
}

resource "ibm_atracker_target" "atracker_eventstreams_target" {
  eventstreams_endpoint {
		target_crn = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		brokers = [ "kafka-x:9094" ]
		topic = "my-topic"
		api_key = "xxxxxxxxxxxxxx" // pragma: allowlist secret
		service_to_service_enabled = false
  }
  name = "my-eventstreams-target"
  target_type = "event_streams"
  region = "us-south"
}

resource "ibm_atracker_target" "atracker_cloudlogs_target" {
  cloudlogs_endpoint {
    target_crn = "crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
  }
  name = "my-cloudlogs-target"
  target_type = "cloud_logs"
  region = "us-south"
}

```

## Argument Reference

You can specify the following arguments for this resource.

* `cloudlogs_endpoint` - (Optional, List) Property values for the IBM Cloud Logs endpoint in responses.
Nested schema for **cloudlogs_endpoint**:
	* `target_crn` - (Required, String) The CRN of the IBM Cloud Logs instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
* `cos_endpoint` - (Optional, List) Property values for a Cloud Object Storage Endpoint in responses.
Nested schema for **cos_endpoint**:
	* `bucket` - (Required, String) The bucket name under the Cloud Object Storage instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `endpoint` - (Required, String) The host name of the Cloud Object Storage endpoint.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `service_to_service_enabled` - (Required, Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
	* `target_crn` - (Required, String) The CRN of the Cloud Object Storage instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
* `eventstreams_endpoint` - (Optional, List) Property values for the Event Streams Endpoint in responses.
Nested schema for **eventstreams_endpoint**:
	* `api_key` - (Optional, String) The user password (api key) for the message hub topic in the Event Streams instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `brokers` - (Required, List) List of broker endpoints.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `service_to_service_enabled` - (Optional, Boolean) Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.
	* `target_crn` - (Required, String) The CRN of the Event Streams instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `topic` - (Required, String) The messsage hub topic defined in the Event Streams instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
* `name` - (Required, String) The name of the target resource.
* `region` - (Optional, String) Included this optional field if you used it to create a target in a different region other than the one you are connected.
* `target_type` - (Required, String) The type of the target.
  * Constraints: Allowable values are: `cloud_object_storage`, `event_streams`, `cloud_logs`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the atracker_target.
* `api_version` - (Integer) The API version of the target.
  * Constraints: The maximum value is `2`. The minimum value is `2`.
* `created_at` - (String) The timestamp of the target creation time.
* `crn` - (String) The crn of the target resource.
* `message` - (String) An optional message containing information about the target.
* `updated_at` - (String) The timestamp of the target last updated time.
* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
Nested schema for **write_status**:
	* `last_failure` - (String) The timestamp of the failure.
	* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
	* `status` - (String) The status such as failed or success.


## Import

You can import the `ibm_atracker_target` resource by using `id`. The uuid of the target resource.

# Syntax
<pre>
$ terraform import ibm_atracker_target.atracker_target &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_atracker_target.atracker_target f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6
```
