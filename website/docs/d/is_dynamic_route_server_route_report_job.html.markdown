---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_route_report_job"
description: |-
  Get information about an IBM Cloud VPC dynamic route server route report job.
---

# ibm_is_dynamic_route_server_route_report_job

Retrieve information about an existing dynamic route server route report job as a read-only data source. The job can be looked up by its identifier or by its name. Once the job has succeeded, the generated report content is available as a JSON string.

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_dynamic_route_server_route_report_job" "by_id" {
  dynamic_route_server_id                     = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_route_report_job_id = ibm_is_dynamic_route_server_route_report_job.example.dynamic_route_server_route_report_job_id
}

data "ibm_is_dynamic_route_server_route_report_job" "by_name" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-route-report-job"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
- `dynamic_route_server_route_report_job_id` - (Optional, String) The dynamic route server route report job identifier. Exactly one of `dynamic_route_server_route_report_job_id` or `name` must be specified.
- `name` - (Optional, String) The name of the dynamic route server route report job. Exactly one of `dynamic_route_server_route_report_job_id` or `name` must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the data source, in the format `<dynamic_route_server_id>/<dynamic_route_server_route_report_job_id>`.
- `completed_at` - (String) The date and time that this route report job completed. Empty if the route report job has not yet completed.
- `content` - (String) The route report content as a JSON string. Empty until the job status is `succeeded`.
- `created_at` - (String) The date and time that this dynamic route server route report job was created.
- `format` - (String) The format of the route report. Allowed values are: `json`.
- `href` - (String) The URL for this dynamic route server route report job.
- `resource_type` - (String) The resource type.
- `schema_version` - (String) The schema version of the route report content. Allowed values are: `v1`.
- `started_at` - (String) The date and time that this route report job started running. Empty if the route report job has not yet started.
- `status` - (String) The status of the route report job. Allowed values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
- `status_reasons` - (List) The reasons for the current dynamic route server route report status (if any).

  Nested schema for `status_reasons`:
  - `code` - (String) A reason code for this route report status. Allowed values are: `cannot_generate_route_report`, `degraded`, `initializing`, `maintenance`, `route_report_unavailable`, `stopped_by_provider`, `stopped_by_user`.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) A link to documentation about this status reason.
