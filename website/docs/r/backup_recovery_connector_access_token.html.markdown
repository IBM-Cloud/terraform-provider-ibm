---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_access_token"
description: |-
  Manages backup_recovery_connector_access_token.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_access_token

Create backup_recovery_connector_access_tokens with this resource.

**Note**
ibm_backup_recovery_connector_access_token resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

## Example Usage

```hcl
resource "ibm_backup_recovery_connector_access_token" "backup_recovery_connector_access_token_instance" {
    username = "admin"
    password = "admin"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `domain` - (Optional, Forces new resource, String) Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain.
* `password` - (Optional, Forces new resource, String) Specifies the password of the Cohesity user account.
* `username` - (Optional, Forces new resource, String) Specifies the login name of the Cohesity user.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_connector_access_token.