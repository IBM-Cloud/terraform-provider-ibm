---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_route_report_job"
description: |-
  Manages IBM Cloud VPC dynamic route server route report job.
---

# ibm_is_dynamic_route_server_route_report_job

Create, update, and delete a route report job for a dynamic route server. A route report job generates a snapshot of the routes learned and advertised by the dynamic route server. The resource waits for the job to reach a terminal status (`succeeded` or `failed`) before the apply completes, and exposes the generated report content as a JSON string once the job has succeeded.

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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_virtual_network_interface" "example" {
  name   = "example-vni"
  subnet = ibm_is_subnet.example.id
}

resource "ibm_is_dynamic_route_server" "example" {
  name      = "example-drs"
  local_asn = 64512
  members {
    name = "example-drs-member"
    virtual_network_interfaces {
      id = ibm_is_virtual_network_interface.example.id
    }
  }
}

resource "ibm_is_dynamic_route_server_route_report_job" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-route-report-job"
  schema_version          = "v1"
}

output "route_report" {
  value = jsondecode(ibm_is_dynamic_route_server_route_report_job.example.content)
}
```

## Timeouts

The `ibm_is_dynamic_route_server_route_report_job` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- `create` - (Default 10 minutes) Used for creating the route report job and waiting for it to reach a terminal status.
- `update` - (Default 10 minutes) Used for updating the route report job.
- `delete` - (Default 10 minutes) Used for deleting the route report job and waiting for it to be removed.

## Argument Reference

You can specify the following arguments for this resource.

- `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
- `name` - (Optional, String) The name for this dynamic route server route report job. The name is unique across all route report jobs for the dynamic route server. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `schema_version` - (Optional, Forces new resource, String) The schema version to use when generating the route report content. If not specified, defaults to `v1`. Allowed values are: `v1`.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the resource, in the format `<dynamic_route_server_id>/<dynamic_route_server_route_report_job_id>`.
- `dynamic_route_server_route_report_job_id` - (String) The unique identifier for this dynamic route server route report job.
- `completed_at` - (String) The date and time that this route report job completed. Empty if the route report job has not yet completed.
- `content` - (String) The route report content as a JSON string. Empty until the job status is `succeeded`. Use `jsondecode()` to work with the structured report.
- `created_at` - (String) The date and time that this dynamic route server route report job was created.
- `format` - (String) The format of the route report. Allowed values are: `json`.
- `href` - (String) The URL for this dynamic route server route report job.
- `resource_type` - (String) The resource type.
- `started_at` - (String) The date and time that this route report job started running. Empty if the route report job has not yet started.
- `status` - (String) The status of the route report job. Allowed values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
- `status_reasons` - (List) The reasons for the current dynamic route server route report status (if any).

  Nested schema for `status_reasons`:
  - `code` - (String) A reason code for this route report status. Allowed values are: `cannot_generate_route_report`, `degraded`, `initializing`, `maintenance`, `route_report_unavailable`, `stopped_by_provider`, `stopped_by_user`.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) A link to documentation about this status reason.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_dynamic_route_server_route_report_job` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `dynamic_route_server_route_report_job_id` in the following format:

```
<dynamic_route_server_id>/<dynamic_route_server_route_report_job_id>
```

- `dynamic_route_server_id`: A string. The dynamic route server identifier.
- `dynamic_route_server_route_report_job_id`: A string. The dynamic route server route report job identifier.

```terraform
import {
  to = ibm_is_dynamic_route_server_route_report_job.example
  id = "<dynamic_route_server_id>/<dynamic_route_server_route_report_job_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_dynamic_route_server_route_report_job.example <dynamic_route_server_id>/<dynamic_route_server_route_report_job_id>
```
