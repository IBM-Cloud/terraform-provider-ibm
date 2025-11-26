---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_download_files"
description: |-
  Get information about backup_recovery_download_files
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_download_files

Provides a read-only data source to retrieve information about backup_recovery_download_files. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_download_files" "backup_recovery_download_files" {
	backup_recovery_download_files_id = ibm_backup_recovery.backup_recovery_instance.backup_recovery_id
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `backup_recovery_download_files_id` - (Required, Forces new resource, String) Specifies the id of a Recovery.
  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
* `file_type` - (Optional, String) Specifies the downloaded type, i.e: error, success_files_list.
* `include_tenants` - (Optional, Boolean) Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned.
* `length` - (Optional, Integer) Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets).
* `source_name` - (Optional, String) Specifies the name of the source on which restore is done.
* `start_offset` - (Optional, Integer) Specifies the start offset of file chunk to be downloaded.
* `start_time` - (Optional, String) Specifies the start time of restore task.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_download_files.

