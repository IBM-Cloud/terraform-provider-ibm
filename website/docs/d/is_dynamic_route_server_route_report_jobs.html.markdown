---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_route_report_jobs"
description: |-
  Get information about IBM Cloud VPC dynamic route server route report jobs.
---

# ibm_is_dynamic_route_server_route_report_jobs

Retrieve the collection of route report jobs for a dynamic route server as a read-only data source. All pages of the collection are fetched.

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
data "ibm_is_dynamic_route_server_route_report_jobs" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  sort                    = "name"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
- `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order. The default value is `-created_at`. Allowed values are: `created_at`, `name`, `-created_at`, `-name`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the data source.
- `route_report_jobs` - (List) Collection of dynamic route server route report jobs.

  Nested schema for `route_report_jobs`:
  - `completed_at` - (String) The date and time that this route report job completed. Absent if the route report job has not yet completed.
  - `created_at` - (String) The date and time that this dynamic route server route report job was created.
  - `format` - (String) The format of the route report. Allowed values are: `json`.
  - `href` - (String) The URL for this dynamic route server route report job.
  - `id` - (String) The unique identifier for this dynamic route server route report job.
  - `name` - (String) The name for this dynamic route server route report job. The name is unique across all route report jobs for the dynamic route server.
  - `resource_type` - (String) The resource type.
  - `schema_version` - (String) The schema version of the route report content. Allowed values are: `v1`.
  - `started_at` - (String) The date and time that this route report job started running. Absent if the route report job has not yet started.
  - `status` - (String) The status of the route report job. Allowed values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
  - `status_reasons` - (List) The reasons for the current dynamic route server route report status (if any).

    Nested schema for `status_reasons`:
    - `code` - (String) A reason code for this route report status. Allowed values are: `cannot_generate_route_report`, `degraded`, `initializing`, `maintenance`, `route_report_unavailable`, `stopped_by_provider`, `stopped_by_user`.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) A link to documentation about this status reason.
