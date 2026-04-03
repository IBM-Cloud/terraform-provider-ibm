---
layout: "ibm"
page_title: "IBM : ibm_pha_agent_job_status"
description: |-
  Get information about pha_agent_job_status
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_agent_job_status

Returns the current status of the job associated with a PowerHA agent file download. It indicates whether the download job is in running, completed, or failed, along with relevant metadata such as job ID, Job creation time and last updated time.

## Example Usage

```hcl
data "ibm_pha_agent_job_status" "pha_agent_job_status" {
	accept_language = "en-US"
	if_none_match = "abcdef"
	instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
	job_id = "4235r23r5vdfdf-2323"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier of the provisioned instance.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.
* `job_id` - (Required, Forces new resource, String) Unique ID to track the pha agent file download.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_agent_job_status.
* `bytes_downloaded` - (Integer) Number of bytes downloaded so far.
* `creation_at` - (String) Timestamp when the job was created.
* `file_name` - (String) Name of the file that has been downloaded.
* `last_updated_at` - (String) Timestamp of the last update for this status.
* `service_instance_id` - (String) Identifier of the service instance associated with the deployment.
* `status` - (String) Current status of the deployment (e.g., running, completed, failed).
* `total_bytes` - (Integer) Total size in bytes of the file that has to be downloaded.
* `vm_id` - (String) Identifier of the virtual machine involved in the deployment.

