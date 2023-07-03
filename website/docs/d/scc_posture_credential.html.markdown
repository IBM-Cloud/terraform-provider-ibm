---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_credential"
description: |-
  Get information about credential
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_credential

Provides a read-only data source for credential. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_credential" "credential" {
	credential_id = "credential_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The id for the given API.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `credential_id` - The unique identifier of the credential.
* `created_at` - (Required, String) The time that the credentials was created in UTC.

* `created_by` - (Required, String) ID of the user who created the credentials.

* `description` - (Required, String) Credentials description.

* `display_fields` - (Required, List) Details the fields on the credential. This will change as per credential type selected.
Nested scheme for **display_fields**:
	* `auth_url` - (Optional, String) auth url of the Open Stack cloud.This is mandatory for Open Stack Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `aws_arn` - (Optional, String) AWS arn value.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `aws_client_id` - (Optional, String) AWS client Id.This is mandatory for AWS Cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `aws_client_secret` - (Optional, String) AWS client secret.This is mandatory for AWS Cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `aws_region` - (Optional, String) AWS region.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `azure_client_id` - (Optional, String) Azure client Id. This is mandatory for Azure Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `azure_client_secret` - (Optional, String) Azure client secret.This is mandatory for Azure Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `azure_resource_group` - (Optional, String) Azure resource group.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `azure_subscription_id` - (Optional, String) Azure subscription Id.This is mandatory for Azure Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `database_name` - (Optional, String) Database name.This is mandatory for Database Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `ibm_api_key` - (Optional, String) The IBM Cloud API Key. This is mandatory for IBM Credential Type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `ms_365_client_id` - (Optional, String) The MS365 client Id.This is mandatory for Windows MS365 Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `ms_365_client_secret` - (Optional, String) The MS365 client secret.This is mandatory for Windows MS365 Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `ms_365_tenant_id` - (Optional, String) The MS365 tenantId.This is mandatory for Windows MS365 Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `password` - (Optional, String) password of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `pem_data` - (Optional, String) The base64 encoded data to associate with the PEM file.
	  * Constraints: The maximum length is `4000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `pem_file_name` - (Optional, String) The name of the PEM file.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `project_domain_name` - (Optional, String) project domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `project_name` - (Optional, String) Project name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `user_domain_name` - (Optional, String) user domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `username` - (Optional, String) username of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `winrm_authtype` - (Optional, String) Kerberos windows auth type.This is mandatory for Windows Kerberos Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `winrm_port` - (Optional, String) Kerberos windows port.This is mandatory for Windows Kerberos Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
	* `winrm_usessl` - (Optional, String) Kerberos windows ssl.This is mandatory for Windows Kerberos Credential type.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.

* `enabled` - (Required, Boolean) Credentials status enabled/disbaled.

* `name` - (Required, String) Credentials name.

* `purpose` - (Required, String) Purpose for which the credential is created.
  * Constraints: Allowable values are: `discovery_collection`, `discovery_fact_collection`, `remediation`, `discovery_collection_remediation`, `discovery_fact_collection_remediation`.

* `type` - (Required, String) Credentials type.
  * Constraints: Allowable values are: `username_password`, `aws_cloud`, `azure_cloud`, `database`, `kerberos_windows`, `ms_365`, `openstack_cloud`, `ibm_cloud`, `user_name_pem`.

* `updated_at` - (Required, String) The modified time that the credentials was modified in UTC.

* `updated_by` - (Required, String) ID of the user who modified the credentials.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_credential is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
