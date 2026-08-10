---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_route_report_job"
description: |-
  Manages DynamicRouteServerRouteReportJob.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_route_report_job

Create, update, and delete DynamicRouteServerRouteReportJobs with this resource.

## Example Usage

```hcl
resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
  dynamic_route_server_id = "dynamic_route_server_id"
  format = "json"
  name = "my-dynamic-route-server-route-report-1"
  storage_bucket {
		crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
		name = "bucket-27200-lwx4cfvcue"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `format` - (Optional, String) The format used for the route report:`json` - The route report is generated based on the json schema.
  * Constraints: Allowable values are: `json`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `name` - (Optional, String) The name for this dynamic route server route report export job. The name must not be used by another export job for the image. Changing the name will not affect the exported image name, `storage_object.name`, or `storage_href` values.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `storage_bucket` - (Required, List) The Cloud Object Storage bucket of the exported dynamic route server route report object.
Nested schema for **storage_bucket**:
	* `crn` - (Required, String) The CRN of this Cloud Object Storage bucket.
	  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Computed, String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `name` - (Required, String) The globally unique name of this Cloud Object Storage bucket.
	  * Constraints: The maximum length is `63` characters. The minimum length is `3` characters. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Computed, String) The resource type.
	  * Constraints: Allowable values are: `cos_bucket`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the DynamicRouteServerRouteReportJob.
* `completed_at` - (String) The date and time that the dynamic route server route report export job was completed.If absent, the dynamic route server route report export job has not yet completed.
* `created_at` - (String) The date and time that the dynamic route server route report export job was created.
* `href` - (String) The URL for this dynamic route server route report export job.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_dynamic_route_server_route_report_job_id` - (String) The unique identifier for this dynamic route server route report export job.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `json_schema` - (List) JSON schema that defines the structure of the route report content.
Nested schema for **json_schema**:
	* `href` - (String) The canonical URI of the JSON Schema that defines the dynamic route server route report.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `type` - (String) The type of schema document.
	  * Constraints: Allowable values are: `json`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `version` - (String) The version of the route report schema.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_route_report_job`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `started_at` - (String) The date and time that the dynamic route server route report export job started running.If absent, the export job has not yet started.
* `status` - (String) The status of this dynamic route server route report export job:- `deleting`:Dynamic route server route report export job is being deleted- `failed`:Dynamic route server route report export job could not be completed  successfully- `queued`:Dynamic route server route report export job is queued- `running`:Dynamic route server route report export job is in progress- `succeeded`:Dynamic route server route report export job was completed successfullyThe exported route report object is automatically deleted for `failed` jobs.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `deleting`, `failed`, `queued`, `running`, `succeeded`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The reasons for the current status (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `cannot_access_storage_bucket`, `internal_error`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the status reason.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) A link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `storage_href` - (String) The Cloud Object Storage location of the exported dynamic route server route report object. The object at this location will not exist until the job completes successfully. The exported image object is not managed by the IBM VPC service, and may be removed or replaced with a different object by any user or service with IAM authorization to the storage bucket.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^cos:\/\/([^\/?#]*)([^?#]*)$/`.
* `storage_object` - (List) The Cloud Object Storage object for the exported image. This object will not exist untilthe job completes successfully. The exported dynamic route server route report object isnot managed by the IBM VPC service, and may be removed or replaced with a differentobject by any user or service with IAM authorization to the storage bucket.
Nested schema for **storage_object**:
	* `name` - (String) The name of this Cloud Object Storage object. Names are unique within a Cloud Object Storage bucket.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\/!\\-_.*'\\(\\)]*$/`.


## Import

You can import the `ibm_is_dynamic_route_server_route_report_job` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `is_dynamic_route_server_route_report_job_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_route_report_job_id&gt;
</pre>
* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `is_dynamic_route_server_route_report_job_id`: A string in the format `r006-095e9baf-01d4-4e29-986e-20d26606b82a`. The unique identifier for this dynamic route server route report export job.

# Syntax
<pre>
$ terraform import ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job &lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_route_report_job_id&gt;
</pre>
