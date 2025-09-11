---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Event Notification Integration'
description: |-
  Manage EN Integration.
---

# ibm_app_config_integration_en

Create or Delete App Configuration and Event Notification services' integration.

## Example usage

```terraform
resource "ibm_app_config_integration_en" "app_config_integration_en" {
  guid = "guid"
  integration_id = "integration_id"
  en_instance_crn = "en_instance_crn"
  en_endpoint = "en_endpoint"
  en_source_name = "en_source_name"
  description = "description"
}
```

## Argument reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `integration_id` - (Required, String) The integration ID.
- `en_instance_crn` - (Required, String) The CRN of Event Notification service.
- `en_endpoint` - (Required, String) The API endpoint of Event Notification service.
- `en_source_name` - (Required, String) The name by which EN source will be created in Event Notifiaction service.
- `description` - (Optional, String) The description of integration between EN and AC service.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `integration_type` - (String) This will be EVENT_NOTIFICATIONS always.
- `created_time` - (Timestamp) The creation time of the feature flag.
- `updated_time` - (Timestamp) The last modified time of the feature flag data.
- `href` - (String) The feature flag URL.

## Import

The `ibm_app_config_integration_en` resource can be imported by using `guid` of the App Configuration instance and `integrationId`. Get the `guid` from the service instance credentials section of the dashboard.

## Syntax

```bash
terraform import ibm_app_config_integration_en.sample  <guid/integrationId>
```

## Example

```bash
terraform import ibm_app_config_integration_en.sample 272111153-c118-4116-8116-b811fbc31132/sample_integration_en
```
