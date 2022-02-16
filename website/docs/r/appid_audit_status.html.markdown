---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Audit Status"
description: |-
    Provides AppID Audit Status resource.
---

# ibm_appid_audit_status

Create, update, or reset an IBM Cloud AppID Management Services audit status (available for graduated tier only). For more information, see [auditing events for App ID](https://cloud.ibm.com/docs/appid?topic=appid-at-events&interface=api)

## Example usage

```terraform
resource "ibm_appid_audit_status" "status" {
  tenant_id = var.tenant_id
  is_active = true
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` - (Required, Boolean) `true` if auditing should be turned on

## Import

The `ibm_appid_audit_status` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_audit_status.status <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_audit_status.status 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
