---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: Org"
description: |-
  Manages IBM organization.
---

# `ibm_org`

Create, update, or delete a Cloud Foundry organization. For more information, about organization, see [Updating orgs and spaces](https://cloud.ibm.com/docs/account?topic=account-orgupdates).


## Example usage
The following example create the `myorg` Cloud Foundry organization and assigns users access to the organization. 


```
resource "ibm_org" "testacc_org" {
    name = "myorg"
    org_quota_definition_guid = "myorgquotaguid"
    auditors = ["auditor@in.ibm.com"]
    managers = ["manager@in.ibm.com"]
    users = ["user@in.ibm.com"]
    billing_managers = ["billingmanager@in.ibm.com"]
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `auditors`(Optional, Sets) The email addresses of the users that you want to assign Cloud Foundry **Auditor** access to. The email address needs to be associated with an IBMID. Auditors have the following permissions within the org: <ul><li>View users and their assigned roles.</li><li>View quota information.</li></ul>.
- `billing_managers`(Optional, Sets) The email addresses of the users that you want to assign the **Billing manager** access to. The email address needs to be associated with an IBMID. Billing managers have the following permissions within the org: <ul><li>View runtime and service usage information on the usage dashboard.</li></ul>.
- `managers`(Optional, Sets) The email addresses of the users that you want to assign Cloud Foundry **Manager** access to. The email address needs to be associated with an IBMID. Managers have the following permissions within the org: <ul><li>Create, view, edit, or delete spaces.</li><li>View usage and quota information.</li><li>Invite users and manage user access. </li><li> Assign roles to users.</li><li>Manage custom domains.</li></ul>.
- `name` - (Required, String) The descriptive name used of the Cloud Foundry organization. The name must be unique in IBM Cloud.
- `org_quota_definition_guid` - (Optional, String) The GUID for the quota that is assigned to organization. The quota sets memory, service, and instance limits for the organization.
- `tags`-(Optional, array of strings) Tags associated with the org.  **Note** Tags are managed locally and not stored on the IBM Cloud service endpoint.
- `users`(Optional, Sets) The email addresses of the users that you want to grant org-level access to. The email address needs to be associated with an IBMID.

**Note**

By default, the user that creates this resource is assigned the **Manager** Cloud Foundry role.  Terraform returns an error when you assign the **Manager** or Cloud Foundry user role to yourself. User information is not persisted in the  Terraform state file to avoid any incorrect information.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The unique identifier of the Cloud Foundry organization.


### Import
The Cloud Foundry organization can be imported by using the `id`.

**Syntax**

```
terraform import ibm_org.myorg <id>
```

**Example**

```
terraform import ibm_org.myorg abde-12345
```

If you bring your current organization in Terraform management, you can perform others operations such as adding a user or assigning users with the required roles.


