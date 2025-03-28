---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_registration"
description: |-
  Manages Data-Source Connector Registration Request.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_registration

Create Data-Source Connector Registration Requests with this resource.

**Note**
ibm_backup_recovery_connector_registration resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

## Example Usage

```hcl
resource "ibm_backup_recovery_connector_registration" "backup_recovery_connector_registration_instance" {
  registration_token = "registration_token"
  access_token = resource.ibm_backup_recovery_connector_access_token.name.access_token
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connector_id` - (Optional, Forces new resource, Integer) The connector's ID to be used for registration. Two connectors belonging to the same tenant are guaranteed to have different IDs.
* `registration_token` - (Required, Forces new resource, String) The registration token. This can be obtained from ibm_backup_recovery_data_source_connection output or from an existing connection.
* `access_token` - (Required, String) Specify an access token for connector authentication.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Data-Source Connector Registration Request.
