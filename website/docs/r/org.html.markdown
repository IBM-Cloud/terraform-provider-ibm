---
layout: "ibm"
page_title: "IBM: Org"
sidebar_current: "docs-ibm-resource-org"
description: |-
  Manages IBM Organization.
---

# ibm\_org

Provides a organization resource. This allows organizations to be created, updated, and deleted.

## Example Usage

```hcl
resource "ibm_org" "testacc_org" {
			name = "myorg"
      org_quota_definition_guid = "myorgquota"
			auditors = ["auditor@in.ibm.com"]
			managers = ["manager@in.ibm.com"]
			users = ["user@in.ibm.com"]
      billing_managers = ["billingmanager@in.ibm.com"]
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify a org.
* `org_quota_definition_guid` - (Optional, string) The name of the Org Quota Definition associated with the org.
* `managers` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give a manager role in this org. Users with the manager role  can create, view, edit, or delete spaces within the organization, view the organization's usage and quota, invite users to the organization, manage who has access to the organization and their roles in the organization, and manage custom domains for the organization.
* `users` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want add at the org level. 
* `auditors` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give an auditor role in this org. Users with the auditor role can also view the users in the organization and their assigned roles, and the quota for the organization.
* `billing_managers` - (Optional, set) The email addresses (associated with IBMids) of the users to whom you want to give an billing manager role in this org. Users with the billing managers can view runtime and service usage information for the organization on the Usage Dashboard page.

**NOTE**: By default the newly created org will have your own  email address to the `managers` or `users` field .

* `tags` - (Optional, array of strings) Tags associated with the org instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the organization.  
