---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Key Management Service Integration'
description: |-
  Manage KMS Integration.
---

# ibm_app_config_integration_kms

Create or Delete App Configuration and Key Management services' integration.

## Example usage

```terraform
resource "ibm_app_config_integration_kms" "app_config_integration_kms" {
  guid = "guid"
  integration_id = "integration_id"
  kms_instance_crn = "kms_instance_crn"
  kms_endpoint = "kms_endpoint"
  root_key_id = "root_key_id"
}
```

## Argument reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `integration_id` - (Required, String) The integration ID.
- `kms_instance_crn` - (Required, String) The CRN of Key Management Service.
- `kms_endpoint` - (Required, String) The API endpoint of Key Management service.
- `root_key_id` - (Required, String) The ID of root key of KMS instance.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `integration_type` - (String) This will be KMS always.
- `kms_schema_type` - (String) The field indicating type of KMS instance used (eg:- KP, HPCP).
- `created_time` - (Timestamp) The creation time of the feature flag.
- `updated_time` - (Timestamp) The last modified time of the feature flag data.
- `href` - (String) The feature flag URL.
