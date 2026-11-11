---
subcategory: "watsonx"
layout: "ibm"
page_title: "IBM : ibm_watsonx_storage_delegation"
description: |-
  Delegates a COS instance to watsonx for KMS encryption.
---

# ibm_watsonx_storage_delegation

Delegates a Cloud Object Storage instance to watsonx so projects and catalogs on it are encrypted with your own key. Create the COS to KMS authorization policy first.

## Example usage

```terraform
resource "ibm_watsonx_storage_delegation" "this" {
  region          = "us-south"
  cos_instance_id = module.cos.cos_instance_guid
  kms_key_crn     = var.cos_kms_key_crn
  depends_on      = [time_sleep.wait_for_authorization_policy]
}
```

## Argument reference

- `region` - (Required, Forces new resource, String) One of `us-south`, `eu-gb`, `eu-de`, `jp-tok`, `au-syd`, `ca-tor`.
- `endpoint` - (Optional, Forces new resource, String) Override for the Data Platform UI host.
- `cos_instance_id` - (Required, Forces new resource, String) COS instance GUID.
- `kms_key_crn` - (Required, String) Root key CRN.
- `catalogs` - (Optional, Boolean) Delegate catalog storage. Default `true`.
- `projects` - (Optional, Boolean) Delegate project storage. Default `true`.

## Import

```
terraform import ibm_watsonx_storage_delegation.this us-south/<cos_instance_guid>
```
