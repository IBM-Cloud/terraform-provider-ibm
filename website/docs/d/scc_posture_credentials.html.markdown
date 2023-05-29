---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_credentials"
description: |-
  Get information about list_credentials
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_credentials

Provides a read-only data source for list_credentials. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_credentials" "list_credentials" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_credentials.
* `credentials` - (List) The details of a credentials.
Nested scheme for **credentials**:
	* `created_at` - (String) The time that the credentials was created in UTC.
	* `created_by` - (String) ID of the user who created the credentials.
	* `description` - (String) Credentials description.
	* `display_fields` - (List) Details the fields on the credential. This will change as per credential type selected.
	Nested scheme for **display_fields**:
		* `auth_url` - (String) auth url of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_arn` - (String) AWS arn value.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_client_id` - (String) AWS client Id.This is mandatory for AWS Cloud.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_client_secret` - (String) AWS client secret.This is mandatory for AWS Cloud.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `aws_region` - (String) AWS region.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_client_id` - (String) Azure client Id. This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_client_secret` - (String) Azure client secret.This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `azure_resource_group` - (String) Azure resource group.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_subscription_id` - (String) Azure subscription Id.This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `database_name` - (String) Database name.This is mandatory for Database Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ibm_api_key` - (String) The IBM Cloud API Key. This is mandatory for IBM Credential Type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ms_365_client_id` - (String) The MS365 client Id.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ms_365_client_secret` - (String) The MS365 client secret.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ms_365_tenant_id` - (String) The MS365 tenantId.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `password` - (String) password of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `pem_data` - (String) The base64 encoded data to associate with the PEM file.
		  * Constraints: The maximum length is `4000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `pem_file_name` - (String) The name of the PEM file.
		  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `project_domain_name` - (String) project domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `project_name` - (String) Project name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `user_domain_name` - (String) user domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `username` - (String) username of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `winrm_authtype` - (String) Kerberos windows auth type.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `winrm_port` - (String) Kerberos windows port.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
		* `winrm_usessl` - (String) Kerberos windows ssl.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `enabled` - (Boolean) Credentials status enabled/disbaled.
	* `id` - (String) Credentials ID.
	* `name` - (String) Credentials name.
	* `purpose` - (String) Purpose for which the credential is created.
	  * Constraints: Allowable values are: `discovery_collection`, `discovery_fact_collection`, `remediation`, `discovery_collection_remediation`, `discovery_fact_collection_remediation`.
	* `type` - (String) Credentials type.
	  * Constraints: Allowable values are: `username_password`, `aws_cloud`, `azure_cloud`, `database`, `kerberos_windows`, `ms_365`, `openstack_cloud`, `ibm_cloud`, `user_name_pem`.
	* `updated_at` - (String) The modified time that the credentials was modified in UTC.
	* `updated_by` - (String) ID of the user who modified the credentials.

* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_credentials is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
