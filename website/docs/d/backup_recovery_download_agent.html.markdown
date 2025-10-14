---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_download_agent"
description: |-
  Get information about backup_recovery_download_agent
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_download_agent

Provides a read-only data source to retrieve information about a backup_recovery_download_agent. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_download_agent" "backup_recovery_download_agent" {
	platform = "kWindows"
	file_path = "./agent.exe"
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `linux_params` - (Optional, List) Linux agent parameters.
Nested schema for **linux_params**:
	* `package_type` - (Required, String) Specifies the type of installer.
	  * Constraints: Allowable values are: `kScript`, `kRPM`, `kSuseRPM`, `kDEB`, `kPowerPCRPM`.
* `platform` - (Required, String) Specifies the platform for which agent needs to be downloaded.
* `file_path` - (Required, String) Specifies the absolute path for download.
  * Constraints: Allowable values are: `kWindows`, `kLinux`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_download_agent.

