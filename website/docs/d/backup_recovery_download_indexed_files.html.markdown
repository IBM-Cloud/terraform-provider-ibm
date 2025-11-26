---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_download_indexed_files"
description: |-
  Get information about backup_recovery_download_indexed_files
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_download_indexed_files

Provides a read-only data source to retrieve information about backup_recovery_download_indexed_files. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_download_indexed_files" "backup_recovery_download_indexed_files" {
	snapshots_id = "snapshots_id"
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `file_path` - (Optional, String) Specifies the path to the file to download. If no path is specified and snapshot environment is kVMWare, VMX file for VMware will be downloaded. For other snapshot environments, this field must be specified.
* `length` - (Optional, Integer) Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets).
* `nvram_file` - (Optional, Boolean) Specifies if NVRAM file for VMware should be downloaded.
* `retry_attempt` - (Optional, Integer) Specifies the number of attempts the protection run took to create this file.
* `snapshots_id` - (Required, Forces new resource, String) Specifies the snapshot id to download from.
* `start_offset` - (Optional, Integer) Specifies the start offset of file chunk to be downloaded.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_download_indexed_files.

