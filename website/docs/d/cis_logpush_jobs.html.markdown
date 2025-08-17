---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_logpush_jobs"
description: |-
  Get information on an IBM Cloud Internet Services logpush jobs.
---

# ibm_cis_logpush_jobs

Retrieve information about an IBM Cloud Internet Services logpush jobs data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_logpush_jobs" "tests" {
  cis_id          = data.ibm_cis.cis.id
  domain_id       = data.ibm_cis_domain.cis_domain.domain_id
  job_id          = data.ibm_cis_domain.logpush_job.job_id
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The Domain ID of the CIS service instance.
- `job_id` - (Optional, String) Logpush Job ID. If `job_id` is not provided you will get the list of jobs. If `job_id` is provided you will get the information of the particular job.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The Logpush Job ID. It is a combination of <`job_id`>,<`cis_id`> attributes concatenated with ":"
- `logpush_job_pack` - (List)
  - `job_id` - (String) The Logpush job ID.
  - `name` - (String) The name of the Logpush job.
  - `enabled` - (Bool) Whether the logpush job is enabled or not.
  - `logpull_options` - (String) Configuration string for the Logpush Job.
  - `destination_conf` - (String) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
  - `dataset` - (String) Dataset to be pulled for the Logpush Job. The values are `http_requests`, `range_events`, `firewall_events`.
  - `frequency` - (String) The frequency at which CIS sends batches of logs to your destination, `high`, `low`.
  - `last_complete` - (String) Records the last time that logs have been successfully pushed.
  - `last_error` - (String) Records the last time the job failed.
  - `error_message` - (String) The last failure.
