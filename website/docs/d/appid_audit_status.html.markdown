---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Audit Status"
description: |-
    Retrieves AppID Audit Status.
---

# ibm_appid_audit_status
Retrieve IBM Cloud AppID Management Services audit status. For more information, see [auditing events for App ID](https://cloud.ibm.com/docs/appid?topic=appid-at-events&interface=api)

## Example usage

```terraform
data "ibm_appid_audit_status" "status" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source (available for graduated tier only).

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` - (Boolean) `true` if auditing is turned on
