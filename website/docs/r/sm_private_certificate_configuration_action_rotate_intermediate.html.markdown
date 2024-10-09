---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_action_rotate_intermediate"
description: |-
  Manages PrivateCertificateConfigurationActionRotateIntermediate.
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_action_rotate_intermediate

Provides a resource for PrivateCertificateConfigurationActionRotateIntermediate. This allows PrivateCertificateConfigurationActionRotateIntermediate to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_private_certificate_configuration_action_rotate_intermediate" "sm_private_certificate_configuration_action_rotate_intermediate_instance" {
  instance_id           = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region                = "us-south"
  name                  = "my_intermediate_configuration_name"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Required, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `name` - (Required, Forces new resource, String) The name that uniquely identifies the intermediate configuration to be rotated.
