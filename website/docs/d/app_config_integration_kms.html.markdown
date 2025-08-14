---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Key Management Service Integration'
description: |-
  Get information about Key Management Service Integration
---

# ibm_app_config_integration_kms

Retrieve information about an existing IBM Cloud App Configuration Key Management Service Integration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_app_config_integration_kms" "app_config_integration_kms" {
  guid = "guid"
  integration_id = "integration_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `integration_id` - (Required, String) The Integration ID.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `integration_type` - (String) The type of integration [will be KMS always].
- `kms_instance_crn` - (String) The CRN of integrated KMS instance.
- `kms_endpoint` - (String) The API endpoint for the KMS instance.
- `root_key_id` - (String) The key ID used for encryption.
- `key_status` - (String) The status of usability of key.
- `kms_schema_type` - (String) Type of KMS service used.
- `created_time` - (Timestamp) The creation time of the integration.
- `updated_time` - (Timestamp) The last modified time of the integration.
- `href` - (String) The integration URL.
