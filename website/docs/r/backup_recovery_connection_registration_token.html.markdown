---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connection_registration_token"
description: |-
  Manages backup_recovery_connection_registration_token.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_connection_registration_token

Create backup_recovery_connection_registration_tokens with this resource.

**Note**
ibm_backup_recovery_connection_registration_token resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_connection_registration_token" "backup_recovery_connection_registration_token_instance" {
  connection_id = ibm_backup_recovery_data_source_connection.backup_recovery_data_source_connection_instance.connectionID
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connection_id` - (Required, Forces new resource, String) Specifies the ID of the connection, connectors belonging to which are to be fetched.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_connection_registration_token.
* `registration_token` - (String) 

### Import
Not Supported