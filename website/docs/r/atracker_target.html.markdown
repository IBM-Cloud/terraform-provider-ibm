---
layout: "ibm"
page_title: "IBM : ibm_atracker_target"
description: |-
  Manages Activity Tracker Target.
subcategory: "Activity Tracker"
---

# ibm_atracker_target

Provides a resource for Activity Tracker Target. This allows Activity Tracker Target to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_atracker_target" "atracker_target" {
  cos_endpoint { 
     endpoint = "endpoint" 
     target_crn = "target_crn" 
     bucket = "bucket" 
     api_key = "api_key"
  }
  name = "my-cos-target"
  target_type = "cloud_object_storage"
  region = "us-south"
}

resource "ibm_atracker_target" "atracker_logdna_target" {
  target_type = "logdna"
  logdna_endpoint {
    target_crn = "crn:v1:bluemix:public:logdna:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
    ingestion_key = "xxxxxxxxxxxxxx"
  }
  name = "my-logdna-target"
  target_type = "logdna"
  region = "us-south"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `cos_endpoint` - (Optional, List) Property values for a Cloud Object Storage Endpoint.
Nested scheme for **cos_endpoint**:
	* `api_key` - (Optional, String) The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response. This is required if service_to_service is not enabled.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `bucket` - (Required, String) The bucket name under the Cloud Object Storage instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `endpoint` - (Required, String) The host name of the Cloud Object Storage endpoint.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `service_to_service_enabled` - (Optional, Boolean) ATracker service is enabled to support service to service authentication. If service to service is enabled then set this flag is true and do not supply apikey.
	* `target_crn` - (Required, String) The CRN of the Cloud Object Storage instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
* `logdna_endpoint` - (Optional, List) Property values for a LogDNA Endpoint.
Nested scheme for **logdna_endpoint**:
	* `ingestion_key` - (Required, String) The LogDNA ingestion key is used for routing logs to a specific LogDNA instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
	* `target_crn` - (Required, String) The CRN of the LogDNA instance.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:\/]+$/`.
* `name` - (Required, String) The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
* `region` - (Optional, String) Include this optional field if you want to create a target in a different region other than the one you are connected.
  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
* `target_type` - (Required, Forces new resource, String) The type of the target. It can be cloud_object_storage or logdna. Based on this type you must include cos_endpoint or logdna_endpoint.
  * Constraints: Allowable values are: `cloud_object_storage`, `logdna`.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.
* `id` - The unique identifier of the atracker_target.
* `api_version` - (Required, Integer) The API version of the target.
* `created_at` - (String) The timestamp of the target creation time.
* `crn` - (Required, String) The crn of the target resource.
* `encryption_key` - (Optional, String) The encryption key that is used to encrypt events before Activity Tracker services buffer them on storage. This credential is masked in the response.
* `updated_at` - (String) The timestamp of the target last updated time.
* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
Nested scheme for **write_status**:
	* `last_failure` - (Optional, String) The timestamp of the failure.
	* `reason_for_last_failure` - (Optional, String) Detailed description of the cause of the failure.
	* `status` - (Required, String) The status such as failed or success.
* `cos_write_status` - **DEPRECATED** (Optional, List) The status of the write attempt with the provided cos_endpoint parameters.
Nested scheme for **cos_write_status**:
	* `status` - (Optional, String) The status such as failed or success.
	* `last_failure` - (Optional, String) The timestamp of the failure.
	* `reason_for_last_failure` - (Optional, String) Detailed description of the cause of the failure.
* `created` - **DEPRECATED** (Optional, String) The timestamp of the target creation time.
* `updated` - **DEPRECATED** (Optional, String) The timestamp of the target last updated time.

## Import

You can import the `ibm_atracker_target` resource by using `id`. The uuid of the target resource.

# Syntax
```
$ terraform import ibm_atracker_target.atracker_target <id>
```

# Example
```
$ terraform import ibm_atracker_target.atracker_target f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6
```
