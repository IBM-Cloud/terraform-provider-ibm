---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connectors_metadata"
description: |-
  Get information about backup_recovery_connectors_metadata
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_connectors_metadata

Provides a read-only data source to retrieve information about a backup_recovery_connectors_metadata. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connectors_metadata" "backup_recovery_connectors_metadata" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_connectors_metadata.
* `connector_image_metadata` - (List) Specifies information about the connector images for various platforms.
Nested schema for **connector_image_metadata**:
	* `connector_image_file_list` - (List) Specifies info about connector images for the supported platforms.
	Nested schema for **connector_image_file_list**:
		* `image_type` - (String) Specifies the platform on which the image can be deployed.
		  * Constraints: Allowable values are: `VSI`, `VMware`.
		* `url` - (String) Specifies the URL to access the file.

