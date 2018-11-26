---
layout: "ibm"
page_title: "IBM: Org"
sidebar_current: "docs-ibm-resource-org"
description: |-
  Manages IBM Organization.
---

# ibm\_org

Provides an organization resource. This allows organizations to be created, updated, and deleted.

## Example Usage

```hcl
resource "ibm_org" "testacc_org" {
    name = "myorg"
    org_quota_definition_guid = "myorgquotaguid"
    auditors = ["auditor@in.ibm.com"]
    managers = ["manager@in.ibm.com"]
    users = ["user@in.ibm.com"]
    billing_managers = ["billingmanager@in.ibm.com"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify a org. The org name must be unique in IBM Cloud and cannot be in use by another IBM Cloud user. 
* `org_quota_definition_guid` - (Optional, string) The GUID for the quota associated with the org. The quota sets memory, service, and instance limits for the org.
* `managers` - (Optional, set) The email addresses for users that you want to assign manager access to. The email address needs to be associated with an IBMid. Managers have the following permissions within the org:
  * Can create, view, edit, or delete spaces.
  * Can view usage and quota information.
  * Can invite users and manage user access.
  * Can assign roles to users.
  * Can manage custom domains.
* `users` - (Optional, set) The email addresses for the users that you want to grant org-level access to. The email address needs to be associated with an IBMid. 
* `auditors` - (Optional, set) The email addresses for the users that you want to assign auditor access to. The email address needs to be associated with an IBMid. Auditors have the following permissions within the org:
  * Can view users and their assigned roles.
  * Can view quota information.
* `billing_managers` - (Optional, set) The email addresses for the users that you want to assign billing manager access to. The email address needs to be associated with an IBMid. Billing managers have the following permissions within the org:
  * Can view runtime and service usage information on the usage dashboard.  

**NOTE**: By default the user creating this resource will have the manager role as per the Cloud Foundry API behavior. Terraform will throw error if you add yourself to the manager role or as a user. This information is not persisted in the state file to avoid any spurious diffs.
* `tags` - (Optional, array of strings) Tags associated with the org.  
  **NOTE**: Tags are managed locally and not stored on the IBM Cloud service endpoint.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the org.  


## Import

Org can be imported using the `id`, e.g.

```
$ terraform import ibm_org.myorg abde-12345
```
Once you bring your current org under terraform management, you can perform others operations like adding user or assigning them required roles.
