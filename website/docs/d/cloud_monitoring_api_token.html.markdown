---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM: ibm_cloud_monitoring_api_token"
description: |-
  Get the IBM Cloud Monitoring API token of an IBM Cloud Monitoring instance.
---

# ibm_cloud_monitoring_api_token

Retrieve the API token of an IBM Cloud Monitoring instance. You can use this token to configure tools that call the IBM Cloud Monitoring API directly. The data source calls the `GET /api/token` endpoint of the IBM Cloud Monitoring instance with the IAM access token of the provider. For more information, see [Getting the API token](https://cloud.ibm.com/docs/monitoring?topic=monitoring-api_token).

The token is a secret. It is stored in the Terraform state, so protect your state file.

## Example usage

```terraform
resource "ibm_resource_instance" "monitoring" {
  name     = "my-monitoring"
  service  = "sysdig-monitor"
  plan     = "graduated-tier"
  location = "us-south"
}

data "ibm_cloud_monitoring_api_token" "token" {
  instance_id = ibm_resource_instance.monitoring.guid
  region      = "us-south"
}

output "cloud_monitoring_api_token" {
  value     = data.ibm_cloud_monitoring_api_token.token.value
  sensitive = true
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `instance_id` - (Required, String) The GUID or the CRN of the IBM Cloud Monitoring instance. When you pass a CRN, the GUID is read from it.
- `region` - (Optional, String) The region of the IBM Cloud Monitoring instance, for example `us-south`. Default value is the provider region.
- `endpoint_type` - (Optional, String) The type of IBM Cloud Monitoring endpoint to call. Allowed values are `public` and `private`. Default value is `public`. Use `private` when Terraform runs inside IBM Cloud and must use private endpoints.

You can set the `IBMCLOUD_MONITORING_API_ENDPOINT` environment variable (for example `https://us-south.monitoring.cloud.ibm.com`) to override the endpoint. When it is set, `region` and `endpoint_type` are not used to build the URL.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - (String) The GUID of the IBM Cloud Monitoring instance.
- `value` - (String, Sensitive) The API token of the IBM Cloud Monitoring instance.
