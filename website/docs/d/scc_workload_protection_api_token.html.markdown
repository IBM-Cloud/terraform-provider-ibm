---
subcategory: "Security and Compliance Center"
layout: "ibm"
page_title: "IBM: ibm_scc_workload_protection_api_token"
description: |-
  Get the API token of an IBM Cloud Security and Compliance Center Workload Protection instance.
---

# ibm_scc_workload_protection_api_token

Retrieve the API token of an IBM Cloud Security and Compliance Center Workload Protection instance. You can use this token to configure tools that call the Workload Protection API directly, such as the registry scanner or the host and agent shields. The data source calls the `GET /api/token` endpoint of the Workload Protection instance with the IAM access token of the provider. For more information, see the [IBM Cloud Security and Compliance Center Workload Protection documentation](https://cloud.ibm.com/docs/workload-protection).

The token is a secret. It is stored in the Terraform state, so protect your state file.

## Example usage

```terraform
resource "ibm_resource_instance" "workload_protection" {
  name     = "my-workload-protection"
  service  = "sysdig-secure"
  plan     = "graduated-tier"
  location = "us-south"
}

data "ibm_scc_workload_protection_api_token" "token" {
  instance_id = ibm_resource_instance.workload_protection.guid
  region      = "us-south"
}

output "scc_workload_protection_api_token" {
  value     = data.ibm_scc_workload_protection_api_token.token.value
  sensitive = true
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `instance_id` - (Required, String) The GUID or the CRN of the Security and Compliance Center Workload Protection instance. When you pass a CRN, the GUID is read from it.
- `region` - (Optional, String) The region of the Security and Compliance Center Workload Protection instance, for example `us-south`. Default value is the provider region.
- `endpoint_type` - (Optional, String) The type of Security and Compliance Center Workload Protection endpoint to call. Allowed values are `public` and `private`. When not set, the endpoint follows the provider `visibility` setting, the same as other services.

The data source uses the Security and Compliance Center Workload Protection client that the provider builds from its configuration (region, `visibility`, endpoints file and IAM authentication). You can override the endpoint with the `IBMCLOUD_SCC_WORKLOAD_PROTECTION_API_ENDPOINT` environment variable or the same key in the endpoints file, for example `https://us-south.security-compliance-secure.cloud.ibm.com`. For more information, see [custom service endpoints](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/guides/custom-service-endpoints). When the environment variable is set, it is used as is and `region` and `endpoint_type` do not change the URL.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - (String) The GUID of the Security and Compliance Center Workload Protection instance.
- `value` - (String, Sensitive) The API token of the Security and Compliance Center Workload Protection instance.
