---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_ce'
description: |-
  Get information about a Code Engine destination
---

# ibm_en_destination_ce

Provides a read-only data source for Cloud Engine destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_destination_ce" "cloudengine_en_destination" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_ce.destination1.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `codengine_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type ibmce.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `url` - (String) URL of code engine project.

  - `verb` - (String) HTTP method of code engine url. Allowable values are: `get`, `post`.

  - `type` - (Optional, String) The code engine destination type . Allowable values are: `application`, `job`.

  - `job_name` - (Optional, String) name of the code engine job.

  - `project_crn` - (Optional, String) CRN of the code engine project.

  - `custom_headers` - (Optional, String) Custom headers (Key-Value pair) for webhook call.

  - `sensitive_headers` - (Optional, array) List of sensitive headers from custom headers.

- `updated_at` - (String) Last updated time.
